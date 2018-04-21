package main

import (
	"github.com/untillpro/qg/git"
	"github.com/untillpro/qg/vcs"
)

var cfgStatus vcs.CfgStatus

// Status shows status of current folder
func Status() {
	git.Status(cfgStatus)
}
