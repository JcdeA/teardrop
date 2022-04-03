/*
Copyright © 2021 Jeeho Ahn and Ayane Satomi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"os"

	"github.com/fosshostorg/teardrop/provisioning"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.testing")
	c := provisioning.NewA64Client(os.Getenv("AARCH64_APIKEY"))
	vms, _ := c.GetVMs()
	for _, vm := range vms {
		println("lmao")
		provisioning.CopyKeys(vm, []string{"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAII4bZze3gL2J7AiL7AT4dUMg/vZY71Hd9DW92Wmmcn/2",
			"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIJYokzEKkIYgmaA2rejRP68ZR+f+/R+PCfB+Olmn+2wR",
			"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIL0ciASK3DsPi9Nl52YbhoJTSYrk5s+Y1JjAL2eKuxqE",
			"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICIt5CNW7Ozr1ITLRwimORCMdoA94YmtGLD7bWNtxc/N"})
	}

	//api.StartAPI()

}
