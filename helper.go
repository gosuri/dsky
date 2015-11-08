package dsky

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"text/template"
	"unicode"

	"github.com/spf13/cobra"
)

// UsageFunc provides usage information for topics and commands
func (cli *CLI) UsageFunc(c *cobra.Command) error {
	var ct *Topic
	for _, topic := range cli.Topics {
		if topic.Name == c.Name() {
			ct = topic
			// Search for subcomands for the topic.
			// apps should have apps:create, apps:destroy
			ct.Commands = make([]*cobra.Command, 0)
			for _, cmd := range c.Parent().Commands() {
				// Conditions for inclusion:
				// 1. Include if the command's name beings the topic name.
				//    Any commands that begin with auth:* will be part of the the auth topic
				// 2. The command's name is not topics'name. Do not include auth command under auth topic
				// 3. Is not a help command
				if strings.HasPrefix(cmd.Name(), ct.Name) && cmd.Name() != ct.Name && cmd.Name() != "help" {
					ct.Commands = append(ct.Commands, cmd)
				}
			}
		}
	}

	// Show help topics for root command
	if c.Name() == cli.Root().Name() {
		sort.Sort(ByName{cli.Topics})
		return tmpl(c.Out(), rootTemplate, cli)
	}

	// Print the regular help
	if err := tmpl(c.Out(), usageTemplate, c); err != nil {
		return err
	}

	// Print the topic help if this is a topic
	if ct != nil && len(ct.Commands) > 0 {
		return tmpl(c.Out(), additionalCmdsTmpl, ct)
	}
	return nil
}

// rpad adds padding to the right of a string
func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(template, s)
}

func trimRightSpace(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) error {
	t := template.New("top")
	t.Funcs(template.FuncMap{
		"trim":           strings.TrimSpace,
		"rpad":           rpad,
		"gt":             cobra.Gt,
		"trimRightSpace": trimRightSpace,
		"eq":             cobra.Eq,
	})
	template.Must(t.Parse(text))
	return t.Execute(w, data)
}

const rootTemplate = `{{$pad := .Topics.TopicNamePadding}}Usage: {{.Root.Use}}
{{ if gt .Topics.Primary 0 }}
Primary help topics, type "{{.Name}} help TOPIC" for more details:
{{range .Topics.Primary }} 
  {{rpad .Name $pad}} {{.Desc}}{{ end }}{{ end }}

{{ if gt .Topics.Additional 0 }}Additional topics:
{{range .Topics.Additional }} 
  {{rpad .Name $pad}} {{.Desc}}{{ end }}{{ end }}

`

const additionalCmdsTmpl = `{{$pad := .NamePadding}}Additional commands, type "{{.RootCmdPath}} COMMAND --help" for more details:
{{range .Commands}}
  {{rpad .Use $pad}} {{.Short}}{{end}}

`

const usageTemplate = `Usage: {{if .Runnable}}{{.UseLine}}{{if .HasFlags}} [options]{{end}}{{end}}{{if .HasSubCommands}}
  {{ .CommandPath}} [command]{{end}}{{if gt .Aliases 0}}

Aliases: {{.NameAndAliases}} {{end}}{{if .HasExample}}

Examples:

  {{ .Example }}{{end}}{{ if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasLocalFlags}}

Options:

{{.LocalFlags.FlagUsages | trimRightSpace}}{{end}}{{ if .HasInheritedFlags}}

General Options:

{{.InheritedFlags.FlagUsages | trimRightSpace}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsHelpCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasSubCommands }}

	Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}

`
