package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Service struct {
	Name string `json:"name"`
	Git string `json:"git"`
	Path string `json:"path,omitempty"`
	Cmd string `json:"cmd"`
	BuildCmd string `json:"buildCmd"`
}

type Instance struct {
	Service string `json:"service"`
	Name string `json:"name"`
	Path string `json:"path"`
	Cmd string `json:"cmd"`
}

type Host struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Instances []Instance `json:"instances"`
}

func ReadService(r io.Reader) (*Service, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	ret := new(Service)
	err = json.Unmarshal(b, ret)
	return ret, err
}

func ReadHost(r io.Reader) (*Host, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	ret := new(Host)
	err = json.Unmarshal(b, ret)
	return ret, err
}


