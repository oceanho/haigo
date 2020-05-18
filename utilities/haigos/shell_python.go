package haigos

import (
	"fmt"
	"time"
)

type PythonScriptExecutor struct {
	PyInterpreter string
}

var (
	pythonScriptExecutorEngine     string
	pythonScriptExecutorEngineArgs []string
)

func init() {
	fmt.Printf("PythonScriptExecutor")
	pythonScriptExecutorEngine = "/usr/bin/env"
	pythonScriptExecutorEngineArgs = []string{
		"python",
		"-c",
	}
}

func (PythonScriptExecutor) Engine() string {
	return pythonScriptExecutorEngine
}

func (PythonScriptExecutor) Args() []string {
	return pythonScriptExecutorEngineArgs
}

func (executor PythonScriptExecutor) Execute(scriptCmd ScriptCommand) error {
	return invoker(executor, scriptCmd)
}

func (executor PythonScriptExecutor) ExecuteTimeout(scriptCmd ScriptCommand, timeout time.Duration) error {
	return invokerTimeout(executor, scriptCmd, timeout)
}
