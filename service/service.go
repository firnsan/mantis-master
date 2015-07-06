package service

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/json"
	// "bytes"
	"errors"
	"strings"
	"github.com/firnsan/mantis-master/config"
)


type Service struct {
	Name string `json:"name"`
	Git string `json:"git"`
	Path string `json:"path"`
	Cmd string `json:cmd`
	BuildCmd string `json:"buildCmd"`
}

type Instance struct {
	Service string `json:"service"`
	Name string `json:"name"`
	Path string `json:"path"`
	Cmd string `json:"cmd"`
}

func readHost(host string) (*config.Host, error) {
	f ,err := os.Open("../host_conf/" + host +".json")
	if err != nil {
		msg := fmt.Sprintf("fail to read %s's conf: %s", host, err)
		log.Println(msg)
		return nil, errors.New(msg)
	}
	ret, err := config.ReadHost(f)
	if err != nil {
		msg := fmt.Sprintf("fail to parse %s's conf: %s", host, err)
		log.Println(msg)
		return nil, errors.New(msg)
	}
	return ret, nil
}

func readService(serv string) (*config.Service, error) {
	f, err := os.Open("../service_conf/" + serv + ".json")
	if err != nil {
		msg := fmt.Sprintf("fail to read %s's conf: %s", serv, err)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	ret, err := config.ReadService(f)
	if err != nil {
		msg := fmt.Sprintf("fail to parse %s's conf: %s", serv, err)
		log.Println(msg)
		return nil, errors.New(msg)
	}
	return ret, nil
}


func deployService(method string, serv *Service) error {
	buf, _ := json.Marshal(serv)
	fmt.Println(string(buf[:]))

	r := strings.NewReader(string(buf[:]))
	resp, _ := http.Post(method, "text/json", r)

	resp.Body.Close()

	return nil
}

func DeployServices(host string, servs []Service) error {
	if len(servs) == 0 {
		return nil
	}
	hostInfo, err := readHost(host)
	if err != nil {
		return err
	}
	method := "http://" + hostInfo.Address + "/service/run"
	// 需要修改servs数组，不能使用for-range
	for i := 0; i<len(servs); i++ {
		serv := &servs[i]
		servInfo, err := readService(serv.Name)
		if err != nil {
			continue
		}
		serv.Git = servInfo.Git
		deployService(method, serv)
	}

	return nil
}

func runInstance(method string, instance *Instance) error{
	// fmt.Println(method)
	buf, _ := json.Marshal(*instance)
	fmt.Println(string(buf[:]))

	r := strings.NewReader(string(buf[:]))
	resp, _ := http.Post(method, "text/json", r)

	// TODO: 当请求失败时，不应该调用Close
	resp.Body.Close()
	
	return nil
}

func RunInstances(host string, instances []Instance) error {
	hostInfo, err := readHost(host)
	if err != nil {
		return err
	}
	method := "http://" + hostInfo.Address + "/instance/run"

	if len(instances) == 0 { //启动host配置的instance
		for _, inst := range hostInfo.Instances {
			instance := new(Instance)
			instance.Service = inst.Service
			instance.Name = inst.Name
			instance.Path = inst.Path // TODO:考虑从service读取path
			instance.Cmd = inst.Cmd

			runInstance(method, instance)
		}
	}

	return nil
}

func RunTempInstance(host string, instance Instance) {
}
