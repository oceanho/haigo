package haigos

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestPythonScriptExecutor_Execute(t *testing.T) {
	logFile, err := os.Create("tmp-py.txt") //io.Writer()
	assert.Nil(t, err)
	defer logFile.Close()
	logWriter := io.MultiWriter(logFile)
	bash := new(PythonScriptExecutor)
	cmdLsRootDir := ScriptCommand{
		RunAs:     "root",
		Cmd:       "import os; print(os.path)",
		LogWriter: logWriter,
		ErrWriter: logWriter,
	}
	err = bash.Execute(cmdLsRootDir)
	t.Logf("%v", err)
}
