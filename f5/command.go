package f5

type BashCommand struct {
	Command         string `json:"command"`
	UtilCommandArgs string `json:"utilCmdArgs"`
}

type BashCommandResult struct {
	Kind            string `json:"kind"`
	Command         string `json:"command"`
	UtilCommandArgs string `json:"utilCmdArgs"`
	CommandResult   string `json:"commandResult"`
}

func (f *Device) Run(command string) (error, *BashCommandResult) {
	u := f.Proto + "://" + f.Hostname + "/mgmt/tm/util/bash"
	b := BashCommand{Command: "run", UtilCommandArgs: "-c \"" + command + "\""}
	r := BashCommandResult{}
	err, _ := f.sendRequest(u, POST, &b, &r)
	if err != nil {
		return err, nil
	}
	return nil, &r
}
