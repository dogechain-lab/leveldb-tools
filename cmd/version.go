package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/dogechain-lab/leveldb-tools/internal/params"
	"github.com/dogechain-lab/leveldb-tools/internal/version"
	"github.com/urfave/cli/v2"
)

const (
	clientIdentifier = "ldb"
)

var versionCmd = &cli.Command{
	Action:    printVersion,
	Name:      "version",
	Aliases:   []string{"v"},
	Usage:     "Print version numbers",
	ArgsUsage: " ",
	Description: `
The output of this command is supposed to be machine-readable.
`,
}

func printVersion(ctx *cli.Context) error {
	git, _ := version.VCS()

	fmt.Println(strings.Title(clientIdentifier))
	fmt.Println("Version:", params.VersionWithMeta)
	if git.Commit != "" {
		fmt.Println("Git Commit:", git.Commit)
	}
	if git.Date != "" {
		fmt.Println("Git Commit Date:", git.Date)
	}
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Printf("GOPATH=%s\n", os.Getenv("GOPATH"))
	fmt.Printf("GOROOT=%s\n", runtime.GOROOT())
	return nil
}
