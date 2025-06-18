package server

import (
	"emoLog/utils"
	"github.com/hashicorp/consul/api"
	"os"
)

func RegisterToConsul() {

	config := api.DefaultConfig()
	// grpc等待时间过长 不使用
	env := os.Getenv("ENV")
	isProduction := env == "production"
	var ip string
	var addr string
	if !isProduction {
		addr = os.Getenv("ADDR")

		if len(addr) != 0 {
			config.Address = addr
		}
		// 本地开发环境，生产环境需要配置
		result, err := utils.GetLocalIP()
		if err != nil {
			ip = "127.0.0.1"
		} else {
			ip = result
		}
	}
	client, _ := api.NewClient(config)

	reg := &api.AgentServiceRegistration{
		ID:      "emo-service",
		Name:    "emo-service",
		Port:    50052,
		Address: ip,
		Check: &api.AgentServiceCheck{
			GRPC:     ip + ":50052",
			Interval: "10s",
		},
	}

	if err := client.Agent().ServiceRegister(reg); err != nil {
		return
	}
}
