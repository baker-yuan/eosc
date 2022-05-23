package process_admin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eolinker/eosc"
	"github.com/eolinker/eosc/extends"
	"github.com/eolinker/eosc/log"
	"github.com/eolinker/eosc/utils/schema"
	"strings"
	"sync"
)

var (
	ErrorExtenderNotWork         = errors.New("not work")
	ErrorInnerExtenderCantChange = errors.New("is inner")
	ErrorNotExist                = errors.New("not exist")
	ErrorDuplicatePath           = errors.New("path duplicate")
	ErrorNotMatch                = errors.New("not match profession")
)

type ExtenderProject struct {
	group   string
	project string
	version string
	renders eosc.IUntyped
	isWork  bool
}

func (ep *ExtenderProject) toInfo() []*ExtenderItem {
	response := make([]*ExtenderItem, 0, ep.renders.Count())
	for _, name := range ep.renders.Keys() {
		response = append(response, &ExtenderItem{
			ExtenderItemInfo: ExtenderItemInfo{
				Group:   ep.group,
				Project: ep.project,
				Name:    name,
				Version: ep.version,
			},
			Id: fmt.Sprint(ep.group, ":", ep.version, ":", name),
		})
	}
	return response
}

type ExtenderData struct {
	Versions map[string]string
	Infos    map[string]*ExtenderProject
	history  map[string]bool
	locker   sync.RWMutex
}

func NewExtenderData(conf map[string][]byte) *ExtenderData {
	vs := make(map[string]string)
	for k, v := range conf {
		vs[k] = string(v)
	}
	for group, project := range extends.GetInners() {
		vs[toProject(group, project)] = extends.InNertVersion
	}

	ed := &ExtenderData{
		Versions: vs,
		Infos:    make(map[string]*ExtenderProject),
		history:  map[string]bool{},
		locker:   sync.RWMutex{},
	}
	ed.init()
	return ed
}
func (e *ExtenderData) IsWork() bool {
	e.locker.RLock()
	defer e.locker.RUnlock()
	for k, version := range e.Versions {
		id := idVersion(k, version)
		if info, has := e.Infos[id]; !has || !info.isWork {
			return has
		}
	}
	return true
}
func (e *ExtenderData) init() {

	for k, v := range e.Versions {
		group, project := readProject(k)
		e.load(group, project, v)
	}

}
func (e *ExtenderData) Delete(group, project string) (*ExtenderProject, error) {
	e.locker.RLock()
	defer e.locker.RUnlock()
	if extends.IsInner(group, project) {
		return nil, ErrorInnerExtenderCantChange
	}

	name := toProject(group, project)
	v, has := e.Versions[name]
	if !has {
		return nil, ErrorNotExist
	}
	extenderProject, _ := e.load(group, project, v)
	delete(e.Versions, name)
	return extenderProject, nil
}

func (e *ExtenderData) getVersion(group, project string) (version string, has bool) {
	if extends.IsInner(group, project) {
		return extends.InNertVersion, true
	}
	v, has := e.Versions[toProject(group, project)]

	return v, has

}

func (e *ExtenderData) setVersion(group, project, version string) {
	id := toProject(group, project)
	e.Versions[id] = version
}
func (e *ExtenderData) SetVersion(group, project, version string) (*ExtenderProject, error) {
	log.Debug("SetVersion:", group, ":", project, ":", version)
	e.locker.Lock()
	defer e.locker.Unlock()

	if extends.IsInner(group, project) {
		return nil, fmt.Errorf("%s:%s %w", group, project, ErrorInnerExtenderCantChange)
	}

	load, err := e.load(group, project, version)
	if err != nil {
		return nil, err
	}
	if !load.isWork {
		return nil, fmt.Errorf("%s:%s:%s %w", group, project, version, ErrorExtenderNotWork)
	}
	e.setVersion(group, project, version)
	return load, nil
}

func (e *ExtenderData) load(group, project, version string) (*ExtenderProject, error) {
	id := toVersion(group, project, version)

	if e.history[toProject(group, project)] {
		info, has := e.Infos[id]
		if has {
			return info, nil
		}
		extendsInfo, err := extends.CheckExtender(group, project, version)

		if err != nil {
			info = &ExtenderProject{
				renders: nil,
				isWork:  false,
			}
		} else {
			renders := eosc.NewUntyped()
			for _, pi := range extendsInfo.Plugins {
				var render *schema.Schema
				json.Unmarshal([]byte(pi.Render), render)
				renders.Set(pi.Name, render)
			}
			info = &ExtenderProject{
				isWork:  true,
				renders: renders,
			}
		}

		e.Infos[id] = info
		return info, err
	}
	info := &ExtenderProject{
		renders: nil,
		isWork:  false,
	}
	register, err := extends.ReadExtenderProject(group, project, version)
	if err == nil {

		ds := register.All()
		info.renders = eosc.NewUntyped()
		for name, d := range ds {
			info.renders.Set(name, d.Render())
		}
		info.isWork = true
	}
	e.history[toProject(group, project)] = true
	e.Infos[id] = info
	return info, err
}

type ExtenderItemInfo struct {
	Group   string `json:"group"`
	Project string `json:"project"`
	Name    string `json:"name"`
	Version string `json:"version"`
}
type ExtenderItem struct {
	ExtenderItemInfo
	Id string `json:"id"`
}
type ExtenderItemRender struct {
	ExtenderItemInfo
	Render interface{} `json:"render"`
}

func (e *ExtenderData) List() []*ExtenderItem {
	rs := make([]*ExtenderItem, 0, len(e.Versions))
	e.locker.RLock()
	defer e.locker.RUnlock()
	for k, version := range e.Versions {
		id := idVersion(k, version)
		info, has := e.Infos[id]
		if has && info.isWork {
			group, project := readProject(k)
			for _, name := range info.renders.Keys() {
				rs = append(rs, &ExtenderItem{
					Id: fmt.Sprint(group, ":", project, ":", name),
					ExtenderItemInfo: ExtenderItemInfo{
						Group:   group,
						Project: project,
						Name:    name,
						Version: version,
					},
				})
			}
		}
	}
	return rs
}
func (e *ExtenderData) GetRender(group, project, name string) (*ExtenderItemRender, bool) {

	e.locker.RLock()
	defer e.locker.RUnlock()
	version, has := e.getVersion(group, project)
	if !has {
		return nil, false
	}

	projectInfo, hasInfo := e.Infos[toVersion(group, project, version)]
	if !hasInfo || !projectInfo.isWork {
		return &ExtenderItemRender{
			ExtenderItemInfo: ExtenderItemInfo{
				Group:   group,
				Project: project,
				Name:    name,
				Version: version,
			},
			Render: nil,
		}, true
	}
	render, hasRender := projectInfo.renders.Get(name)
	if !hasRender {
		return &ExtenderItemRender{
			ExtenderItemInfo: ExtenderItemInfo{
				Group:   group,
				Project: project,
				Name:    name,
				Version: version,
			},
			Render: nil,
		}, true
	}
	return &ExtenderItemRender{
		ExtenderItemInfo: ExtenderItemInfo{
			Group:   group,
			Project: project,
			Name:    name,
			Version: version,
		},
		Render: render,
	}, true
}

func (e *ExtenderData) GetInfo(group, project string) ([]*ExtenderItemInfo, bool) {

	e.locker.RLock()
	defer e.locker.RUnlock()

	version, has := e.getVersion(group, project)
	if !has {
		return nil, false
	}
	var rs []*ExtenderItemInfo
	projectInfo, hasInfo := e.Infos[toVersion(group, project, version)]
	if !hasInfo || !projectInfo.isWork {
		return rs, true
	}

	names := projectInfo.renders.Keys()
	rs = make([]*ExtenderItemInfo, 0, len(names))
	for _, name := range names {
		rs = append(rs, &ExtenderItemInfo{
			Group:   group,
			Project: project,
			Name:    name,
			Version: version,
		})
	}
	return rs, true
}
func toProject(group, project string) string {
	return fmt.Sprint(group, ":", project)
}
func idVersion(id, version string) string {
	return fmt.Sprint(id, ":", version)
}
func toVersion(group, project, version string) string {
	return fmt.Sprint(group, ":", project, ":", version)
}
func readProject(id string) (group, project string) {
	vs := strings.Split(id, ":")
	group = vs[0]
	if len(vs) > 1 {
		project = vs[1]
	}

	return
}
