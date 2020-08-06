package main

import
(
	"time"
	"github.com/jakewrite/muxinator"
	"github.com/cwyz-dev/Home-Automation/config-management-service/domain"
	"github.com/cwyz-dev/Home-Automation/config-management-service/controller"
	"github.com/cwyz-dev/Home-Automation/config-management-service/service"
)

func main()
{
	config := domain.Config{}
	configService := service.ConfigService
	{
		Config: & config,
		Location: "config.yaml"
	}
	go configService.Watch(time.Second * 60)
	c := controller.Controller
	{
		Config: &config,
	}
	router := mux.inator.NewRouter()
	router.Get("/read/{serviceName}", c.ReadConfig)
	router.ListenAndServe(":8080")
}
