package docker

import (
	"context"
	"os"
	"sync"

	"github.com/cuigh/auxo/log"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

const (
	defaultAPIVersion = "1.32"
)

var mgr = &manager{}

type manager struct {
	client *client.Client
	locker sync.Mutex
	logger log.Logger
}

func (m *manager) Do(fn func(ctx context.Context, cli *client.Client) error) (err error) {
	ctx, cli, err := m.Client()
	if err != nil {
		return err
	}
	return fn(ctx, cli)
}

func (m *manager) Client() (ctx context.Context, cli *client.Client, err error) {
	if m.client == nil {
		m.locker.Lock()
		defer m.locker.Unlock()

		if m.client == nil {
			if apiVersion := os.Getenv("DOCKER_API_VERSION"); apiVersion == "" {
				os.Setenv("DOCKER_API_VERSION", defaultAPIVersion)
			}
			m.client, err = client.NewClientWithOpts(client.FromEnv)
			if err != nil {
				return
			}
		}
	}
	return context.TODO(), m.client, nil
}

func (m *manager) Logger() log.Logger {
	if m.logger == nil {
		m.locker.Lock()
		defer m.locker.Unlock()

		if m.logger == nil {
			m.logger = log.Get("docker")
		}
	}
	return m.logger
}

func version(v uint64) swarm.Version {
	return swarm.Version{Index: v}
}

func page(count, pageIndex, pageSize int) (start, end int) {
	start = pageSize * (pageIndex - 1)
	end = pageSize * pageIndex
	if count < start {
		start, end = 0, 0
	} else if count < end {
		end = count
	}
	return
}
