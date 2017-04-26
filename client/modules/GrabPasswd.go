package modules

import "os/exec"

/*
Will grab the passwd file
*/

func GrabPasswd() string {
	cmd, err := exec.Command("cat /etc/shadow").Output()
	if err != nil {
		panic(err)
	}
	return string(cmd)
}
