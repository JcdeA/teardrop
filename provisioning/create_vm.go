package provisioning

import (
	"encoding/json"

	"github.com/natesales/aarch64-client-go"
)

func (ac *A64Client) BatchProvisionVMs(pop string, master bool, batchNumber int) ([]aarch64.VM, error) {
	var vms []aarch64.VM

	for i := 0; i < batchNumber; i++ {

		var err error
		var vmResponse aarch64.APIResponse
		var vm *aarch64.VM

		if master {
			vmResponse, err = ac.CreateVM("", pop, "61a37b0b1082096ea965fe97", "v1.medium", "debian")
		} else {
			vmResponse, err = ac.CreateVM("", pop, "61a37b0b1082096ea965fe97", "v1.medium", "debian")
		}

		if err != nil {
			return nil, err
		}
		marshalled, err := json.Marshal(vmResponse.Data)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(marshalled, vm)

		vms = append(vms, *vm)
	}
	return vms, nil
}
