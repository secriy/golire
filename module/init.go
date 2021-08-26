package module

import "github.com/secriy/golire/util"

// Log set the logger level of current package.
func Log(level string) {
	util.SetLevelString(level)
}
