package main

import (
	"github.com/subhroacharjee/appname/internal/server"
	"github.com/subhroacharjee/appname/pkg/controllers"
	"github.com/subhroacharjee/appname/utils/container"
)

func run() error {
	container.InitDependencies()
	return container.Container.Invoke(func(s *server.Server) error {
		controllers.InitRouter(s.G)
		return s.Run()
	})
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
