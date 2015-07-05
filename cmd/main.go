package main

import (
	"fmt"
	// "flag"
	//"os"
	"strings"
	"errors"
	"log"
	"github.com/docopt/docopt-go"
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
	mantis_cmd run <host>... [--autoUpdate|--autoBuild]
	mantis_cmd run <app@host>... [--autoUpdate|--autoBuild]
	mantis_cmd run <app@host> [--service=<service>|--cmd=<command>|--path=<path>]

`

	arguments, _ := docopt.Parse(usage, nil, true, "Mantis Cmd 0.1", false)
	fmt.Println(arguments)

	if arguments["run"].(bool) {
		hosts, err := parseHost(arguments["<host>"].([]string))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hosts)
	}
}