package racer

import (
	"strings"

	"github.com/gosuri/racer/ui"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	DefaultCLI = New(nil)
)

type CLI struct {
	Topics
	NoInteractive bool

	ui    ui.UI
	root  *cobra.Command
	flags *pflag.FlagSet
}

func SetRoot(root *cobra.Command) *CLI {
	return DefaultCLI.SetRoot(root)
}

func AddTopic(name, desc string, primary bool) *CLI {
	return DefaultCLI.AddTopic(name, desc, primary)
}

func AddCommand(cmd *cobra.Command) *CLI {
	return DefaultCLI.AddCommand(cmd)
}

func Execute() error {
	return DefaultCLI.Execute()
}

func Printer() *ui.Printer {
	return DefaultCLI.Printer()
}

func (c *CLI) Printer() *ui.Printer {
	return c.UI().Printer()
}

func UI() ui.UI {
	return DefaultCLI.UI()
}

func New(root *cobra.Command) *CLI {
	return &CLI{root: root}
}

func (c *CLI) AddTopic(name, desc string, primary bool) *CLI {
	c.Topics = append(c.Topics, &Topic{name, desc, primary, nil, c.Root().CommandPath()})
	return c
}

func (c *CLI) SetRoot(root *cobra.Command) *CLI {
	c.root = root
	return c
}

func (c *CLI) AddCommand(cmd *cobra.Command) *CLI {
	c.Root().AddCommand(cmd)
	return c
}

// Name returns the first part of the cmd.Use
func (c *CLI) Name() string {
	return strings.Split(c.Root().Use, " ")[0]
}

// Root returns the command associated with the cli
func (c *CLI) Root() *cobra.Command {
	if c.root == nil {
		c.root = &cobra.Command{}
	}
	return c.root
}

func (c *CLI) UI() ui.UI {
	if c.ui == nil {
		c.ui = ui.NewStdUI()
	}
	c.ui.SetNoInteractive(c.NoInteractive)
	return c.ui
}

func (c *CLI) SetUI(ui ui.UI) *CLI {
	c.ui = ui
	return c
}

// Execute binds usage function and runs the root command
func (c *CLI) Execute() error {
	return c.Bind().Root().Execute()
}

// Bind binds the usage function and global flags to the command
func (c *CLI) Bind() *CLI {
	c.Root().PersistentFlags().BoolVarP(&c.NoInteractive, "no-interactive", "", c.NoInteractive, "Disable interactive mode")
	c.Root().SetUsageFunc(c.UsageFunc)
	return c
}
