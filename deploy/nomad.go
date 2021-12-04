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

func (n *NomadClient) NewDeployment(run models.Run, runId int) {
	config := make(map[string]interface{})
	config["image"] = run.Project.ContainerImage

	resources := &nomad.Resources{
		CPU:      pointer.ToInt(1024),
		MemoryMB: pointer.ToInt(128),
	}
	task := &nomad.Task{
		Name:      run.Name,
		Driver:    "docker",
		Config:    config,
		Resources: resources,
	}
	group := &nomad.TaskGroup{
		Name:  pointer.ToString(run.Name),
		Count: pointer.ToInt(1),
		Tasks: []*nomad.Task{task},
	}
	job := &nomad.Job{
		ID:          pointer.ToString(fmt.Sprint(run.Id)),
		Name:        pointer.ToString(run.Name),
		Region:      pointer.ToString("global"),
		Priority:    pointer.ToInt(20),
		Datacenters: []string{"dc1"},
		Type:        pointer.ToString("batch"),
		TaskGroups:  []*nomad.TaskGroup{group},
	}
	if _, _, err := n.Jobs().Validate(job, nil); err != nil {
		fmt.Printf("Nomad job validation failed. Error: %v\n", err)
	}
	n.Jobs().Register(job, nil)

}
