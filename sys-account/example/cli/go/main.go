package main

import (
	"github.com/amplify-cms/sys-share/sys-core/service/logging/zaplog"
	"github.com/spf13/cobra"

	pkg "github.com/amplify-cms/sys-share/sys-account/service/go/pkg"
)

var rootCmd = &cobra.Command{
	Use:   "auth",
	Short: "auth cli",
}

func main() {
	zl := zaplog.NewZapLogger(zaplog.DEBUG, "sys-account-example", true, "")
	zl.InitLogger(nil)
	rootCmd.AddCommand(pkg.NewSysShareProxyClient().CobraCommand())
	if err := rootCmd.Execute(); err != nil {
		zl.Fatalf("command failed: %v", err.Error())
	}
}
