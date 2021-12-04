package provisioning

import (
	"errors"
	"strings"

	"github.com/fosshostorg/teardrop/models"
)

func (ac *A64Client) GetVMs() ([]models.VM, error) {
	projects, _ := ac.Projects()
	for _, p := range projects.Projects {
		if p.Name == "Fosshost - Teardrop" {
			vms := []models.VM{}

			for _, vm := range p.VMs {
				var role models.VMRole
				if strings.HasPrefix(vm.Hostname, "master") {
					role = models.Master
				} else if strings.HasPrefix(vm.Hostname, "worker") {
					role = models.Worker

				}
				vms = append(vms, models.VM{VM: vm, Role: role})
			}
			return vms, nil
		}
	}
	return nil, errors.New("no VMs found in Teardrop AArch64 project")
}
