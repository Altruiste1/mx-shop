package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.30.166:8500"
	//cfg.Address = address
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	client.Agent().ServiceRegister(registration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Address = address

	check := &api.AgentServiceCheck{
		HTTP:                           "http://192.168.30.194:8021/health",
		Timeout:                        "5s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "10s",
	}
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func main() {
	_ = Register("http://192.168.30.194", 8021, "user-web", []string{"test"}, "01")
	AllService()
}

func AllService() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.30.166:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter("")
	if err != nil {
		panic(err)
	}
	for key, value := range data {
		fmt.Println(key, value)
	}
}
