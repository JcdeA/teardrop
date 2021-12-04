package provisioning

import (
	"fmt"

	"os"
	"os/exec"
	"strings"

	"github.com/fosshostorg/teardrop/models"
)

func runCmd(cmd string) error {
	command := exec.Command("bash", "...")
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}

func InstallHashiUp() error {
	return runCmd(`/root/teardrop/provisioning/scripts/setup.sh`)
}

func InstallServer(vms []models.VM) error {
	var retryJoin string
	for _, vm := range vms {
		retryJoin = retryJoin + strings.Split(vm.Address, "/")[0]
	}

	err := runCmd(fmt.Sprintf(`
		hashi-up consul install \
			--version 1.10.4 \
			--local \
			--server \
			--bootstrap-expect %v\
			--client-addr [::] \
			--advertise-addr "{{ GetInterfaceIP \"eth1\" }}" \
			--connect \
			--retry-join "%v"`, len(vms), retryJoin))
	if err != nil {
		return err
	}

	err = runCmd(fmt.Sprintf(`
		hashi-up nomad install \
			--version 1.2.2 \
			--local \
			--server \
			--bootstrap-expect %v \
			--advertise "{{ GetInterfaceIP \"eth1\" }}"
		`, len(vms)))
	if err != nil {
		return err
	}

	err = runCmd(`systemctl start consul && systemctl start nomad`)
	if err != nil {
		return err
	}
	return nil
}
