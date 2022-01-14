package deploy

import (
	"fmt"

	"github.com/AlekSi/pointer"
	"github.com/fosshostorg/teardrop/models"
	nomad "github.com/hashicorp/nomad/api"
)

type NomadClient struct {
	*nomad.Client
}

func NewNomadClient(address string) (*NomadClient, error) {
	client, err := nomad.NewClient(&nomad.Config{Address: address})
	if err != nil {
		return nil, err
	}
	return &NomadClient{client}, nil
}

func (n *NomadClient) NewDeployment(run models.Run, runId int) error {
	config := models.Map{
		"image": run.Project.ContainerImage,
		"ports": []string{"http"},
	}

	resources := &nomad.Resources{
		CPU:      pointer.ToInt(5),
		MemoryMB: pointer.ToInt(128),
	}
	task := &nomad.Task{
		Name:      "testgroup",
		Driver:    "docker",
		Config:    config,
		Resources: resources,
	}
	group := &nomad.TaskGroup{
		Name:  pointer.ToString(run.Name),
		Count: pointer.ToInt(1),
		Tasks: []*nomad.Task{task},
		Networks: []*nomad.NetworkResource{
			{
				DynamicPorts: []nomad.Port{
					{Label: "http", To: 80},
				},
			},
		},
	}
	job := &nomad.Job{
		Name:        pointer.ToString(run.Name),
		Region:      pointer.ToString("global"),
		Priority:    pointer.ToInt(20),
		Datacenters: []string{"bos"},
		Type:        pointer.ToString("batch"),
		TaskGroups:  []*nomad.TaskGroup{group},
		ID:          pointer.ToString(fmt.Sprintf("%d", run.Id)),
	}
	if _, _, err := n.Jobs().Validate(job, nil); err != nil {
		fmt.Printf("Nomad job validation failed. Error: %v\n", err)
	}
	_, _, err := n.Jobs().Register(job, nil)
	if err != nil {
		return err
	}
	return nil

}
