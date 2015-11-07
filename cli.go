package racer

import (
	"strings"

	"github.com/gosuri/racer/ui"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// DefaultCLI is the default CLI for the package
var DefaultCLI = New(nil)

// CLI represent the Racer CLI library
type CLI struct {
	// Topics are a list a of Help Topics
	Topics

	// NoIteractive flag determines the interactive mode of the CLI
	NoInteractive bool

	// NoColor when true does not display colors
	NoColor bool

	ui    ui.UserInterface
	root  *cobra.Command
	flags *pflag.FlagSet
}

// SetRoot registers the command as the root command for the default CLI
func SetRoot(root *cobra.Command) *CLI {
	return DefaultCLI.SetRoot(root)
}

// AddTopic adds a help topic to the default CLI
func AddTopic(name, desc string, primary bool) *CLI {
	return DefaultCLI.AddTopic(name, desc, primary)
}

// AddCommand register a cobra command with the default CLI
func AddCommand(cmd *cobra.Command) *CLI {
	return DefaultCLI.AddCommand(cmd)
}

// Execute executes the DefaultCLI. It binds usage function and runs the root command.
func Execute() error {
	return DefaultCLI.Execute()
}

// Printer returns the printer for the Default CLI
func Printer() *ui.UIPrinter {
	return DefaultCLI.Printer()
}

// Printer returns the printer for the CLI
func (c *CLI) Printer() *ui.UIPrinter {
	return c.UI().Printer()
}

// UI reeturns the Default UI
func UI() ui.UserInterface {
	return DefaultCLI.UI()
}

// New returns an instance of the CLI and registers the root command
func New(root *cobra.Command) *CLI {
	return &CLI{root: root}
}

// AddTopic adds a help topic to the CLI
func (c *CLI) AddTopic(name, desc string, primary bool) *CLI {
	c.Topics = append(c.Topics, &Topic{name, desc, primary, nil, c.Root().CommandPath()})
	return c
}

// SetRoot registers the command as root command
func (c *CLI) SetRoot(root *cobra.Command) *CLI {
	c.root = root
	return c
}

// AddCommand register a cobra command with the CLI
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

func (c *CLI) UI() ui.UserInterface {
	if c.ui == nil {
		c.ui = ui.New()
	}
	c.ui.SetNoInteractive(c.NoInteractive)
	c.ui.SetNoColor(c.NoColor)
	return c.ui
}

func (c *CLI) SetUI(ui ui.UserInterface) *CLI {
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
