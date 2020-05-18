package haigos

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestBashScriptExecutor_Execute(t *testing.T) {
	logFile, err := os.Create("tmp.txt") //io.Writer()
	assert.Nil(t, err)
	defer logFile.Close()
	logWriter := io.MultiWriter(logFile)
	bash := new(BashScriptExecutor)
	cmdLsRootDir := ScriptCommand{
		RunAs: "root",
		Cmd:   "ls",
		Args: []string{
			"/",
		},
		LogWriter: logWriter,
		ErrWriter: logWriter,
	}
	err = bash.Execute(cmdLsRootDir)
	t.Logf("%v", err)
}
