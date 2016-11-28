package deviceController

import (
	"../jsonReader"
	"fmt"
	"os/exec"
	"os"
	"strings"
)

const (
	prg = "adb"
	cmdPrefix = "shell input "
	cmdWakeUp = cmdPrefix + "keyevent 26"
	cmdUnlock = cmdPrefix + "keyevent 82"
	cmdEnterKey = cmdPrefix + "keyevent 66"
	cmdEnterPwd = cmdPrefix + "text "
	cmdConn = "connect "
)

type CommandList struct {
	Commands   []string
	CmdChannel chan string
}

func UnlockDevice(ip string, password string) {
	cmdChannel := make(chan string)
	commands := CommandList{[]string{cmdConn + ip, cmdWakeUp, cmdUnlock, cmdEnterPwd + password, cmdEnterKey},
		cmdChannel}

	go executeCmdShell(commands)
	for c := range cmdChannel {
		fmt.Println(c)
	}
}

func executeCmdShell(cmd interface{}) string {
	var output string

	switch cmd.(type) {
	case string:
		output, err := exec.Command(prg, strings.Split(cmd.(string), " ")...).CombinedOutput()
		if err != nil {
			os.Stderr.WriteString("error executing command " + err.Error())
		}
		return string(output)

	case CommandList:
		channel := cmd.(CommandList).CmdChannel
		for _, c := range cmd.(CommandList).Commands {
			output, err := exec.Command(prg, strings.Split(c, " ")...).CombinedOutput()
			if err != nil {
				os.Stderr.WriteString("error executing command " + err.Error())
			}
			channel <- string(output)
		}
		close(channel)
	}

	return output
}

func ExecuteActions(actions []jsonReader.ActionElement) {

}