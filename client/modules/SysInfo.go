package modules

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func windowsSysInfo() (string, string, string, string) {
	//has not been tested
	avProducts := ""

	directory, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		panic(err)
	}

	versionCheck := exec.Command("CMD", "/C", "ver")
	version, _ := versionCheck.Output()
	getUserCommand := string("echo " + "%" + "USERNAME" + "%")
	userCheck := exec.Command("CMD", "/C", getUserCommand)
	user, _ := userCheck.Output()

	// TODO: Find a way to eumerate all AV products on the system

	return directory, string(user), string(version), avProducts
}

func linuxSysInfo() (string, string, string, string) {
	// has not been tested
	avProducts := ""

	directory, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		panic(err)
	}

	versionCheck := exec.Command("uname", "-a")
	version, _ := versionCheck.Output()

	userCheck := exec.Command("whoami")
	user, _ := userCheck.Output()

	return directory, string(version), string(user), avProducts
}

func macSysInfo() (string, string, string, string) {
	// Works for MacOS
	r, _ := regexp.Compile(".*ProductVersion:\\s*([^\n\r]*)")

	avProducts := ""

	directory, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		panic(err)
	}

	versionCheck := exec.Command("sw_vers")
	version, _ := versionCheck.Output()

	versionString := r.FindString(string(version))
	s := strings.Split(versionString, ":")
	s[1] = strings.TrimSpace(s[1])

	userCheck := exec.Command("whoami")
	user, _ := userCheck.Output()

	userS := strings.TrimSpace(string(user))

	return directory, s[1], userS, avProducts
}

func GetSysInfo() (string, string, string, string) {
	if runtime.GOOS == "windows" {
		directory, version, user, avProducts := windowsSysInfo()
	}
	if runtime.GOOS == "linux" {
		directory, version, user, avProducts := linuxSysInfo()
	}
	if runtime.GOOS == "darwin" {
		directory, version, user, avProducts := macSysInfo()
	}
	return directory, user, version, avProducts

}
