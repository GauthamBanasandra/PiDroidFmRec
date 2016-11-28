package deviceController

import (
	"../jsonReader"
	"strings"
	"strconv"
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

// Commands for unlocking the phone.
func GetUnlockCommands(ip string, password string) []string {
	return []string{cmdConn + ip, cmdWakeUp, cmdUnlock, cmdEnterPwd + password, cmdEnterKey}
}

// Transforms a list ActionElements defined in the configuration file into commands.
func GetCommands(actions []jsonReader.ActionElement) (commands []string) {
	for _, action := range actions {
		switch cmd := strings.TrimSpace(action.Cmd); cmd {
		case "tap":
			commands = append(commands, cmdPrefix + cmd + " " + strconv.Itoa(action.Input.X1) + " " +
				strconv.Itoa(action.Input.Y1))
		case "text":
			commands = append(commands, cmdPrefix + cmd + " " + action.Input.Text)
		}
	}

	return
}