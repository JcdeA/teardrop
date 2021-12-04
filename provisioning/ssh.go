package provisioning

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"

	"github.com/fosshostorg/teardrop/models"
	"golang.org/x/crypto/ssh"
)

func SshVM(vm models.VM, cmd string) error {

	var session *ssh.Session

	var client *ssh.Client

	auth := []ssh.AuthMethod{
		ssh.Password(vm.Password),
	}

	sshConfig := &ssh.ClientConfig{
		User: "root",
		Auth: auth,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			// TODO: fix this thing to properly check host
			return nil
		},
		Timeout: time.Second * 30,
	}
	client, err := ssh.Dial("tcp6", fmt.Sprintf("[%v]:22", strings.Split(vm.Address, "/")[0]), sshConfig)
	if err != nil {
		return err
	}

	session, err = client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	err = session.Run(cmd)

	return err
}

func CopyKeys(vm models.VM, keys []string) error {
	for _, key := range keys {
		// try to ssh into the vm before copying key
		_, err := exec.Command("ssh", "-oBatchMode=yes", fmt.Sprintf("root@%v", strings.Split(vm.Address, "/")[0])).Output()
		if err != nil {
			println(err.Error())
			cmd := fmt.Sprintf("mkdir -p ~/.ssh && echo \"%v\" >> ~/.ssh/authorized_keys", key)
			err := SshVM(vm, cmd)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
