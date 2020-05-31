package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"wedog/pkg/config"
	"wedog/pkg/httpServer"
)

var configFile = flag.String("config", "D:\\var\\www\\http\\GoProjects\\src\\wedog\\etc\\config.json", "配置文件路径")

func main() {
	flag.Parse()

	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("data", string(data))
	var conf config.Config
	err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("conf:%+v\n", conf)

	s := httpServer.NewServer(conf)
	s.Run()
}
