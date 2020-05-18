package haigos

import (
	"time"
)

type BashScriptExecutor struct {
}

var (
	bashScriptExecutorEngine     string
	bashScriptExecutorEngineArgs []string
)

func init() {
	bashScriptExecutorEngine = "/bin/bash"
	bashScriptExecutorEngineArgs = []string{
		"-c",
	}
}

func (BashScriptExecutor) Engine() string {
	return bashScriptExecutorEngine
}

func (BashScriptExecutor) Args() []string {
	return bashScriptExecutorEngineArgs
}

func (executor BashScriptExecutor) Execute(scriptCmd ScriptCommand) error {
	return invoker(executor, scriptCmd)
}
func (executor BashScriptExecutor) ExecuteTimeout(scriptCmd ScriptCommand, timeout time.Duration) error {
	return invokerTimeout(executor, scriptCmd, timeout)
}
