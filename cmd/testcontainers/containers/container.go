package containers

import (
	"fmt"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"strings"
)

func IsRunning(pool dockertest.Pool, imagename string) bool {
	dockerContainers, _ := pool.Client.ListContainers(docker.ListContainersOptions{
		All: false,
	})

	for _, dockerContainer := range dockerContainers {
		for _, name := range dockerContainer.Names {
			if strings.Contains(name, imagename) {
				fmt.Println(fmt.Sprintf("%s image is running..", dockerContainer.Image))
				return true
			}
		}
	}

	return false
}
