package cmd

import (
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	currentDBFlag = "CURRENT_DB"

	directoryFlag = &cli.PathFlag{
		Name:     "directory",
		Aliases:  []string{"dir", "d"},
		EnvVars:  []string{"LDB_DIRECTORY"},
		Required: true,
	}
)

// Expands a file path
// 1. replace tilde with users home dir
// 2. expands embedded environment variables
// 3. cleans the path, e.g. /a/b/../c -> /a/c
// Note, it has limitations, e.g. ~someuser/tmp will not be expanded
func expandPath(p string) string {
	if strings.HasPrefix(p, "~/") || strings.HasPrefix(p, "~\\") {
		if home := homeDir(); home != "" {
			p = home + p[1:]
		}
	}
	return path.Clean(os.ExpandEnv(p))
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}
