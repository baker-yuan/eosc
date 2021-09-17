package eosc

import (
	"encoding/json"
)

type IData interface {
	UnMarshal(v interface{}) error
	Marshal() ([]byte, error)
}

type StoreValue struct {
	Id         string
	Profession string
	Name       string
	Driver     string
	CreateTime string
	UpdateTime string
	IData      IData
	Sing       string
}

type IStoreEventHandler interface {
	OnInit(vs []StoreValue) error
	OnDel(v StoreValue) error
	OnChange(v StoreValue) error
}

type IStoreRW interface {
	Initialization() error
	All() []StoreValue
	Get(id string) (StoreValue, bool)
	Set(v StoreValue) error
	Del(id string) error
	Reset([]StoreValue) error
}

type IStoreListener interface {
	AddListen(h IStoreEventHandler) error
}

type IStore interface {
	IStoreRW
	GetListener() IStoreListener
}

type IStoreFactory interface {
	Create(params map[string]string) (IStore, error)
}

type BytesData []byte

func (b BytesData) Marshal() ([]byte, error) {
	return b, nil
}

func (b BytesData) UnMarshal(v interface{}) error {
	return json.Unmarshal(b, v)
}

func MarshalBytes(v interface{}) (BytesData, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return BytesData(data), nil
}
