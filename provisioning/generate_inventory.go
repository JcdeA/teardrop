package provisioning

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (ac *A64Client) GenerateInventory() error {
	vms, err := ac.GetVMs()
	if err != nil {
		return err
	}

	masters := make(fiber.Map)
	for _, vm := range vms {
		if strings.Contains(vm.Hostname, "master") {
			masters[vm.Hostname] = fiber.Map{
				"ansible_host": strings.Split(vm.Address, "/")[0],
				"ansible_user": "root",

				"nomad_node_role": "server",
			}
		}

	}

	workers := make(fiber.Map)
	for _, vm := range vms {
		if strings.Contains(vm.Hostname, "worker") {
			workers[vm.Hostname] = fiber.Map{
				"ansible_host":    strings.Split(vm.Address, "/")[0],
				"ansible_user":    "root",
				"nomad_node_role": "client",
			}
		}
	}

	i := fiber.Map{
		"all": fiber.Map{
			"children": fiber.Map{
				"workers": fiber.Map{
					"hosts": workers,
				},
				"masters": fiber.Map{
					"hosts": masters,
				},
			},
		},
	}
	ba, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	os.WriteFile("/home/jcde/teardrop/provisioning/ansible/hosts.json", ba, 0644)
	println("generated inventory file")
	return nil
}
