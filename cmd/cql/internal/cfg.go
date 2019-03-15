package internal

import (
	"github.com/CovenantSQL/CovenantSQL/client"
	"github.com/CovenantSQL/CovenantSQL/conf"
	"github.com/CovenantSQL/CovenantSQL/crypto/asymmetric"
	"github.com/CovenantSQL/CovenantSQL/utils"
)

// These are general flags used by console and other commands.
var (
	configFile string
	password   string

	CmdName string
)

// AddCommonFlags adds the flags common to all commands.
func AddCommonFlags(cmd *Command) {
	cmd.Flag.StringVar(&configFile, "config", "~/.cql/config.yaml", "Config file for covenantsql")
	cmd.Flag.StringVar(&password, "password", "", "Master key password for covenantsql")

	// Undocumented, unstable debugging flags.
	cmd.Flag.BoolVar(&asymmetric.BypassSignature, "bypass-signature", false,
		"Disable signature sign and verify, for testing")
}

func configInit() {
	configFile = utils.HomeDirExpand(configFile)

	// init covenantsql driver
	if err := client.Init(configFile, []byte(password)); err != nil {
		ConsoleLog.WithError(err).Error("init covenantsql client failed")
		SetExitStatus(1)
		Exit()
		return
	}

	// TODO(leventeliu): discover more specific confirmation duration from config. We don't have
	// enough informations from config to do that currently, so just use a fixed and long enough
	// duration.
	WaitTxConfirmationMaxDuration = 20 * conf.GConf.BPPeriod
}
