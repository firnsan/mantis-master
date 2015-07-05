package config

import (
	"testing"
	"strings"
	"fmt"
	"os"
)

func TestReadService(t *testing.T) {
	r := strings.NewReader(`{"name":"mysql", "git":"github.com", "cmd":"./mysql", "buildCmd":"aaa"}`)
	service, err:= ReadService(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(service)
}

func TestReadHost(t *testing.T) {
	f,err := os.Open("../hosts_conf/host1.json")
	if err != nil {
		fmt.Println(err)
	}
	host, err := ReadHost(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(host)
}