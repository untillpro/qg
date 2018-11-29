package utils

import (
	"errors"
	"io"
	"log"
	"os/exec"
)

// https://github.com/b4b4r07/go-pipe/blob/master/README.md

// PipedExec allows to execute commands in pipe
type PipedExec struct {
	cmds []*pipedCmd
}

// Stderr redirection
const (
	StderrRedirectNone = iota
	StderrRedirectStdout
	StderrRedirectNull
)

type pipedCmd struct {
	stderrRedirection int
	cmd               *exec.Cmd
}

func (Self *PipedExec) command(name string, stderrRedirection int, args ...string) *PipedExec {
	cmd := exec.Command(name, args...)
	lastIdx := len(Self.cmds) - 1
	if lastIdx > -1 {
		var err error
		cmd.Stdin, err = Self.cmds[lastIdx].cmd.StdoutPipe()
		PanicIfError(err)
	}
	Self.cmds = append(Self.cmds, &pipedCmd{stderrRedirection, cmd})
	return Self
}

// Command adds a command to a pipe
func (Self *PipedExec) Command(name string, args ...string) *PipedExec {
	return Self.command(name, StderrRedirectNone, args...)
}

// Wd sets working directory for the last command
func (Self *PipedExec) Wd(wd string) *PipedExec {
	pipedCmd := Self.cmds[len(Self.cmds)-1]
	pipedCmd.cmd.Dir = wd
	return Self
}

// Run starts the pipe
func (Self *PipedExec) Run(out io.Writer, err io.Writer) error {
	for _, cmd := range Self.cmds {
		if cmd.stderrRedirection == StderrRedirectNone {
			cmd.cmd.Stderr = err
		}
		cmd.cmd.Stdout = out
	}
	lastIdx := len(Self.cmds) - 1
	if lastIdx < 0 {
		return errors.New("Empty command list")
	}

	for _, cmd := range Self.cmds {
		log.Println(cmd.cmd.Path, cmd.cmd.Args)
		err := cmd.cmd.Start()
		if nil != err {
			panic("Failed:" + cmd.cmd.Path)
		}
		err = cmd.cmd.Wait()
		if nil != err {
			return err
		}
	}

	return Self.cmds[lastIdx].cmd.Wait()

}
