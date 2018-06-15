package controller

import (
	"github.com/cuigh/auxo/data"
	"github.com/cuigh/auxo/net/web"
	"github.com/cuigh/swirl-agent/docker"
)

// ContainerController is a controller of docker container
type ContainerController struct {
	List   web.HandlerFunc `path:"/"`
	Detail web.HandlerFunc `path:"/:id/detail"`
}

// Container creates an instance of ContainerController
func Container() (c *ContainerController) {
	return &ContainerController{
		List:   containerList,
		Detail: containerDetail,
	}
}

func containerList(ctx web.Context) error {
	args := &docker.ContainerListArgs{}
	err := ctx.Bind(args)
	if err != nil {
		return err
	}
	if args.PageSize <= 0 {
		args.PageSize = 25
	}
	if args.PageIndex <= 0 {
		args.PageIndex = 1
	}

	containers, totalCount, err := docker.ContainerList(args)
	if err != nil {
		return err
	}

	return ctx.JSON(data.Map{
		"count": totalCount,
		"list":  containers,
	})
}

func containerDetail(ctx web.Context) error {
	id := ctx.P("id")
	container, err := docker.ContainerInspect(id)
	if err != nil {
		return err
	}

	return ctx.JSON(container)
}
