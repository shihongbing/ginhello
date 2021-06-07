package safe

import (
	"bytes"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func CmdPing(host string) (result bool,err error) {
	sysType := runtime.GOOS
	if sysType == "windows"{
		cmd := exec.Command("cmd","/c","ping -a -n 1"+"host")
		var out bytes.Buffer
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(out.String(),"TTL"){
			result = true
		}
	}
	return result, err
}
