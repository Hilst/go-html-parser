package main

import (
	"github.com/qri-io/jsonpointer"
)

type DataMap struct {
	data    JSON
	arrays  map[string][]interface{}
	strings map[string]string
}

type JSON = map[string]interface{}

func NewDataMap(data JSON) DataMap {
	return DataMap{
		data:    data,
		arrays:  make(map[string][]interface{}),
		strings: make(map[string]string),
	}
}

func (d *DataMap) String(key string) string {
	return d.strings[key]
}

func (d *DataMap) setString(pointer string, key string, data JSON) {
	poss := d.get(pointer, data).(string)
	d.strings[key] = poss
}

func (d *DataMap) Array(key string) []interface{} {
	return d.arrays[key]
}

func (d *DataMap) setArray(pointer string, key string, data JSON) {
	poss := d.get(pointer, data).([]interface{})
	d.arrays[key] = poss
}

func (d *DataMap) get(pointer string, data JSON) interface{} {
	ptr, err := jsonpointer.Parse(pointer)
	d.check(err)
	got, err := ptr.Eval(data)
	d.check(err)
	return got
}

func (d *DataMap) check(err error) {
	if err != nil {
		panic(err)
	}
}

func (d *DataMap) Prepare(pointer string, key string, typeDesc string, data JSON) string {
	key = "var:" + key
	switch typeDesc {
	case "ARRAY":
		d.setArray(pointer, key, data)
	case "STRING":
		d.setString(pointer, key, data)
	}
	return ""
}
