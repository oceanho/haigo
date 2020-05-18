package haigo

import "os"

var _debugMode = false

func init() {
	debugEnv := os.Getenv("haigo_debug")
	if debugEnv != "" {
		_debugMode = true
	}
}

func DebugMode(debugMode bool) {
	_debugMode = debugMode
}

func IsDebugMode() bool {
	return _debugMode
}
