package provisioning

import (
	"github.com/fosshostorg/teardrop/models"
)

func (ac *A64Client) BatchSetupVMs(vms []models.VM) error {
	for _, vm := range vms {
		err := CopyKeys(vm, []string{"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAII4bZze3gL2J7AiL7AT4dUMg/vZY71Hd9DW92Wmmcn/2 jcde@jcde.xyz"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (ac *A64Client) SetupVM(vm models.VM) error {
	err := CopyKeys(vm, []string{"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAII4bZze3gL2J7AiL7AT4dUMg/vZY71Hd9DW92Wmmcn/2 jcde@jcde.xyz"})
	if err != nil {
		return err
	}

	return nil
}
