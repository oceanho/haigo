package haigos

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

type ScriptExecutor interface {
	Engine() string
	Args() []string
	Execute(scriptCmd ScriptCommand) error
	ExecuteTimeout(scriptCmd ScriptCommand, timeout time.Duration) error
}

type ScriptCommand struct {
	RunAs     string
	WorkDir   string
	Cmd       string
	Args      []string
	LogWriter io.Writer
	ErrWriter io.Writer
}

func invokerTimeout(executor ScriptExecutor, scriptCmd ScriptCommand, duration time.Duration) error {
	return nil
}

func invoker(executor ScriptExecutor, scriptCmd ScriptCommand) error {
	var cmdArgs []string
	argLen := len(executor.Args())
	if argLen != 0 {
		cmdArgs = make([]string, argLen)
		copy(cmdArgs, executor.Args())
		cmdArgs = append(cmdArgs, scriptCmd.Cmd)
		for _, v := range scriptCmd.Args {
			cmdArgs = append(cmdArgs, v)
		}
	} else {
		cmdArgs = make([]string, len(scriptCmd.Args))
		cmdArgs = append(cmdArgs, scriptCmd.Cmd)
		copy(cmdArgs, executor.Args())
	}
	cn := executor.Engine()
	cmd := exec.Command(cn, cmdArgs...)
	if scriptCmd.WorkDir != "" {
		cmd.Dir = scriptCmd.WorkDir
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Bind StdoutPipe fail, cmd(%s),args (%s),err %v", cn, strings.Join(cmdArgs, " "), err.Error())
		return err
	}
	defer stdout.Close()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Bind StderrPipe fail, cmd(%s),args (%s),err %v", cn, strings.Join(cmdArgs, " "), err.Error())
		return err
	}
	defer stderr.Close()
	if err := cmd.Start(); err != nil {
		fmt.Printf("Start fail, cmd(%s),args (%s),err %v", cn, strings.Join(cmdArgs, " "), err.Error())
		return err
	}
	//cmd.Stdout = scriptCmd.LogWriter
	//cmd.Stderr = scriptCmd.ErrWriter
	// TODO(Ocean): should be considerate memory size Limit.
	if scriptCmd.LogWriter != nil {
		bytes, _ := ioutil.ReadAll(stdout)
		scriptCmd.LogWriter.Write(bytes)
	}
	if scriptCmd.ErrWriter != nil {
		bytes, _ := ioutil.ReadAll(stderr)
		scriptCmd.ErrWriter.Write(bytes)
	}
	return nil
}
