package deviceController

const (
	prg = "adb"
	cmdPrefix = "shell input "
	cmdWakeUp = cmdPrefix + "keyevent 26"
	cmdUnlock = cmdPrefix + "keyevent 82"
	cmdEnterKey = cmdPrefix + "keyevent 66"
	cmdEnterPwd = cmdPrefix + "text "
	cmdConn = "connect "
)