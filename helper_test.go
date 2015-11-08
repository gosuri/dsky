package dsky

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func fakeCLI() *CLI {
	root := &cobra.Command{
		Short: "root short",
		Use:   "root",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cli := New(root)

	cmd1 := &cobra.Command{
		Use:   "cmd1",
		Short: "cmd1-short",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("inside cmd with args: %v\n", args)
		},
	}
	cli.AddCommand(cmd1).AddTopic("cmd1", "cmd1-topic", true)

	cmd2 := &cobra.Command{
		Use:   "cmd2",
		Short: "cmd2-short",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("inside cmd with args: %v\n", args)
		},
	}
	cli.AddCommand(cmd2).AddTopic("cmd2", "cmd2-topic", false)

	return cli.Bind()
}

func TestUsageFunc(t *testing.T) {
	got := fakeCLI().Root().UsageString()
	for _, s := range []string{"cmd1-topic", "cmd2-topic", "Usage: root"} {
		if !strings.Contains(got, s) {
			t.Fatalf("expected %s, got:\n%v\n", s, got)
		}
	}
}
