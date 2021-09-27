package process_worker

import (
	"github.com/eolinker/eosc"
	"github.com/eolinker/eosc/log"
)

var _ IProfession = (*Profession)(nil)

type IProfession interface {
	GetDriver(name string) (eosc.IProfessionDriver, bool)
}
type Profession struct {
	*eosc.ProfessionConfig

	drivers ITypedDrivers
}

func (p *Profession) GetDriver(name string) (eosc.IProfessionDriver, bool) {
	return p.drivers.Get(name)
}

func NewProfession(c *eosc.ProfessionConfig) *Profession {
	ds := NewTypedDrivers()
	for _, d := range c.Drivers {
		df, b := eosc.DefaultProfessionDriverRegister.GetProfessionDriver(d.Id)
		if !b {
			log.Warn("driver not exist:", d.Id)
			continue
		}
		driver, err := df.Create(c.Name, d.Name, d.Label, d.Desc, d.Params)
		if err != nil {
			log.Warnf("create driver %s of %s :%v", d.Id, c.Name, err)
			continue
		}
		ds.data.Set(d.Name, driver)
	}
	return &Profession{
		ProfessionConfig: c,
		drivers:          ds,
	}
}
