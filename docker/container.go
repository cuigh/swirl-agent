package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type ContainerListArgs struct {
	// created|restarting|running|removing|paused|exited|dead
	Filter    string `bind:"filter"`
	Name      string `bind:"name"`
	PageIndex int    `bind:"page"`
	PageSize  int    `bind:"size"`
}

// ContainerList return containers on the host.
func ContainerList(args *ContainerListArgs) (containers []types.Container, totalCount int, err error) {
	var (
		ctx context.Context
		cli *client.Client
	)

	ctx, cli, err = mgr.Client()
	if err != nil {
		return
	}

	opts := types.ContainerListOptions{Filters: filters.NewArgs()}
	if args.Filter == "" {
		opts.All = true
	} else {
		opts.Filters.Add("status", args.Filter)
	}
	if args.Name != "" {
		opts.Filters.Add("name", args.Name)
	}

	containers, err = cli.ContainerList(ctx, opts)
	if err == nil {
		//sort.Slice(containers, func(i, j int) bool {
		//	return containers[i] < containers[j].Description.Hostname
		//})
		totalCount = len(containers)
		start, end := page(totalCount, args.PageIndex, args.PageSize)
		containers = containers[start:end]
	}
	return
}

// ContainerInspect return detail information of a container.
func ContainerInspect(id string) (container types.ContainerJSON, err error) {
	var (
		ctx context.Context
		cli *client.Client
	)

	ctx, cli, err = mgr.Client()
	if err == nil {
		container, err = cli.ContainerInspect(ctx, id)
	}
	return
}
