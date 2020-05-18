package haigos

import (
	"fmt"
	"github.com/oceanho/haigo"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var debugMode = false

func init() {
	if haigo.IsDebugMode() {
		debugMode = true
	} else {
		isDebug := os.Getenv("haigo_utilities_haigos_debug")
		if isDebug != "" {
			debugMode = true
		}
	}
}

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
	var args = make([]string, 0)
	for _, v := range executor.Args() {
		args = append(args, v)
	}
	args = append(args, fmt.Sprintf("%s %s",
		scriptCmd.Cmd, strings.Join(scriptCmd.Args, " ")))
	cn := executor.Engine()
	cmd := exec.Command(cn, args...)
	if scriptCmd.WorkDir != "" {
		cmd.Dir = scriptCmd.WorkDir
	}
	if debugMode {
		fmt.Printf("invoker command\n")
		fmt.Printf("Cmd: %s\n", cn)
		fmt.Printf("Args: %s\n", strings.Join(args, " "))
		fmt.Printf("WorkDir: %s\n", cmd.Dir)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("get StdoutPipe fail, cmd(%s),args (%s),err %v", cn, strings.Join(args, " "), err.Error())
		return err
	}
	defer stdout.Close()
	if scriptCmd.LogWriter != nil {
		go output(stdout, scriptCmd.LogWriter)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("get StderrPipe fail, cmd(%s),args (%s),err %v", cn, strings.Join(args, " "), err.Error())
		return err
	}
	defer stderr.Close()
	if scriptCmd.ErrWriter != nil {
		go output(stderr, scriptCmd.ErrWriter)
	}
	if err := cmd.Start(); err != nil {
		fmt.Printf("start fail, cmd(%s),args (%s),err %v", cn, strings.Join(args, " "), err.Error())
		return err
	}
	err = cmd.Wait()
	return err
}

func output(reader io.ReadCloser, writer io.Writer) {
	//exit := false
	//go func() {
	//	wg.Wait()
	//	exit = true
	//}()
	buf := make([]byte, 512)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF || n < 1 || exit {
			break
		}
		writer.Write(buf[0:n])
	}
}
