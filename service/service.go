package service

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/json"
	// "bytes"
	"strings"
	"github.com/firnsan/mantis-master/config"
)


type Service struct {
	Git string `json:"git"`
	BuildCmd string `json:"buildCmd"`
}

type Instance struct {
	Service string `json:"service"`
	Name string `json:"name"`
	Path string `json:"path"`
	Cmd string `json:"cmd"`
}

func readHost(host string) (*config.Host, error) {
	f,err := os.Open("../hosts_conf/"+ host +".json")
	if err != nil {
		log.Printf("fail to read %s's conf: %s", host, err)
		return nil, err
	}
	ret, err := config.ReadHost(f)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

func deployService(method string, serv Service) {
}

func DeployServices(host string, serv []string) {
}

func runInstance(method string, instance *Instance) error{
	fmt.Println(method)
	buf, _ := json.Marshal(*instance)
	fmt.Println(string(buf[:]))

	r := strings.NewReader(string(buf[:]))
	resp, _ := http.Post(method, "text/json", r)

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
