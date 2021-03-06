package deviceController

import (
	"fmt"
	"os/exec"
	"os"
	"strings"
	"time"
)

/*
Struct for representing a list of commands.
A channel is included for communicating intermediate outputs.
*/
type CommandList struct {
	Commands   []string
	CmdChannel chan string
}

// Unlocks the Android device given the IP address and password.
func UnlockDevice(ip string, password string) {
	ExecuteCommands(GetUnlockCommands(ip, password))
}

/*
Executes a given command or a list of commands.
Logic is written to infer the type.
 */
func executeCmdShell(cmd interface{}) string {
	var output string

	// Get the type of the parameter passed.
	switch cmd.(type) {
	// If it's just one command.
	case string:
		output, err := exec.Command(prg, strings.Split(cmd.(string), " ")...).CombinedOutput()
		if err != nil {
			os.Stderr.WriteString("error executing command " + err.Error())
		}
		return string(output)

	// If it's a list of commands.
	case CommandList:
		// Get the channel.
		channel := cmd.(CommandList).CmdChannel
		for _, c := range cmd.(CommandList).Commands {
			output, err := exec.Command(prg, strings.Split(c, " ")...).CombinedOutput()
			if err != nil {
				os.Stderr.WriteString("error executing command " + err.Error())
			}
			// Delay between commands.
			time.Sleep(2 * time.Second)
			// Write each output to the channel.
			channel <- string(output)
		}
		close(channel)
	}

	return output
}


// Executes a list of commands.
func ExecuteCommands(commands []string) {
	// Creating a channel to display outputs of each command immediately after executing.
	cmdChannel := make(chan string)
	cmdList := CommandList{commands, cmdChannel}

	// Begin executing commands.
	go executeCmdShell(cmdList)
	for c := range cmdChannel {
		// Print the intermediate outputs.
		fmt.Println(c)
	}

}