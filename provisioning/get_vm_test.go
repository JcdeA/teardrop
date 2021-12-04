package provisioning_test

import (
	"os"
	"testing"

	"github.com/fosshostorg/teardrop/provisioning"
	"github.com/joho/godotenv"
)

func TestGetVMs(t *testing.T) {
	err := godotenv.Load("../.env.testing")
	if err != nil {
		panic("failed to load env vars")
	}
	ac := provisioning.NewA64Client(os.Getenv("AARCH64_APIKEY"))
	vms, err := ac.GetVMs()
	if err != nil {
		t.Error(err)
	}
	for _, vm := range vms {
		err := provisioning.CopyKeys(vm, []string{"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAII4bZze3gL2J7AiL7AT4dUMg/vZY71Hd9DW92Wmmcn/2 jcde@jcde.xyz"})
		if err != nil {
			t.Log(err)
		}
	}

}
