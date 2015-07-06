package main

import (
	"fmt"
	// "flag"
	//"os"
	"strings"
	"errors"
	"log"
	"github.com/docopt/docopt-go"
	"github.com/firnsan/mantis-master/service"
)


func parseHost(hosts []string) (map[string][]string, error) {
	if len(hosts) == 0 {
		return nil, errors.New("nil hosts")
	}
	var m = make(map[string][]string)

	for _, appAndHost := range hosts {

		var idx int
		var host string // map's key
		var apps string

		var appArr []string
		var a []string // map's value


		idx = strings.Index(appAndHost, "@")
		if idx == -1 { // host1
			if strings.Index(appAndHost, ",") != -1 {
				return nil, errors.New("invalid host: " + appAndHost)
			}
			host = appAndHost
		} else if idx != -1 { // app1@host1
			apps = appAndHost[0:idx]
			host = appAndHost[idx+1:]

			if host == "" {
				return nil, errors.New("nil host")
			}

			appArr = strings.Split(apps, ",")

			for _, app := range appArr {
				app = strings.Trim(app, " ")
				if app == "" {
					continue
				}
				a = append(a, app)
			}
		}
		if _, alreadyExist := m[host]; alreadyExist {
			return nil, errors.New("duplicative host: " + host)
		}
		m[host] = a
	}

	return m, nil
}

func main() {

	usage := `Mantis Cmd

Usage:
	mantis_cmd run <host>... [--autoUpdate]
	mantis_cmd run <instance@host>... [--autoUpdate]
	mantis_cmd run <instance@host> [--service=<service>|--cmd=<command>|--path=<path>]
	mantis_cmd deploy <service@host>...

`

	arguments, _ := docopt.Parse(usage, nil, true, "Mantis Cmd 0.1", false)
	fmt.Println(arguments)

	if arguments["run"].(bool) {
		m, err := parseHost(arguments["<host>"].([]string))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m)

		for host, apps := range m {
			var insts []service.Instance
			for _, app := range apps {
				insts = append(insts,
					service.Instance{Name:app})
			}
			service.RunInstances(host, insts)
		}
	}
	if arguments["deploy"].(bool) {
		m, err := parseHost(arguments["<service@host>"].([]string))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m)

		for host, services := range m {
			if len(services) == 0 { // 没有指定该host的service
				log.Println("no services for " + host)
				continue
			}
			var servs []service.Service
			for _, serv := range services {
				servs = append(servs,
					service.Service{Name:serv})
			}
			service.DeployServices(host, servs)
		}

	}

}