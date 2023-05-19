// Issue 142
// Generate pending taint on variadic param,
// resolve during compose.
// Warning should be generated.

package main

import (
	"net/http"
	"os/exec"
)

func HandleReq(req *http.Request) {
	vals := req.URL.Query()
	cmdName := vals.Get("cmd")
	RunCmd(cmdName)
}

func RunCmd(name ...string) {
	cmd := exec.Command(name[0])
	cmd.Run()
}
