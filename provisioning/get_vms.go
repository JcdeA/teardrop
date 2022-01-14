package provisioning

import (
	"errors"
	"strings"

	"github.com/fosshostorg/teardrop/models"
)

func (ac *A64Client) GetVMs() ([]models.VM, error) {
	projects, err := ac.Projects()
	if err != nil {
		return nil, err
	}

	for _, p := range projects.Projects {
		if p.Name == "Teardrop" {
			vms := []models.VM{}

			for _, vm := range p.VMs {
				var role models.VMRole
				vmRoleString := string(strings.Split(vm.Hostname, "-")[1][0]) // s or c
				if vmRoleString == "s" {
					role = models.Server
				} else if vmRoleString == "c" {
					role = models.Client
				}
				vms = append(vms, models.VM{VM: vm, Role: role})
			}
			return vms, nil
		}
	}
	return nil, errors.New("no VMs found in Teardrop AArch64 project")
}
