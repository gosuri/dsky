package main

import (
	"fmt"
	"github.com/gosuri/racer"
	"github.com/gosuri/racer/ui"
	"github.com/spf13/cobra"
)

type Cluster struct {
	Name, Datacenter string
	Size             int
}

var clusters = []Cluster{
	{"dev", "sfo1", 1},
	{"prod", "sfo2", 3},
}

type App struct {
	Name, Size string
	Deployed   bool
}

var apps = []App{
	{"frontend", "7MB", true},
	{"api", "6MB", false},
}

func main() {
	root := &cobra.Command{
		Short: "Utility to manage your clusters and applications on ovrclk",
		Use:   "ovrclk COMMAND [<args>..] [options]",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	racer.SetRoot(root)

	racer.AddTopic("apps", "create, deploy and manage applications", true)

	appsCmd := &cobra.Command{
		Use:   "apps",
		Short: "List all apps",
		Run: func(cmd *cobra.Command, args []string) {
			racer.Printer().AddTitle("Apps")
			table := ui.NewTable("NAME", "SIZE", "DEPLOYED")
			for _, a := range apps {
				table.AddRow(a.Name, a.Size, a.Deployed)
			}
			racer.Printer().Add(table).Print()
		},
	}
	racer.AddCommand(appsCmd)

	appsInfo := &cobra.Command{
		Use:     "apps:info",
		Short:   "Display info for the app",
		Long:    "Show detailed app information. For more info visit http://example.com/foo",
		Aliases: []string{"info"},
		Example: "ovrclk apps:info -a foo",
		Run: func(cmd *cobra.Command, args []string) {
			app := apps[0]
			racer.Printer().AddTitle(fmt.Sprintf("%s (app)", app.Name))
			table := ui.NewTable()
			table.AddRow("Name:", app.Name)
			table.AddRow("Size:", app.Size)
			table.AddRow("Deployed:", app.Deployed)
			racer.Printer().Add(table).Print()
		},
	}
	var acinfo string
	appsInfo.Flags().StringVarP(&acinfo, "cluster", "c", "", "Show cluster info")
	racer.AddCommand(appsInfo)

	newApp := &App{}
	appsCreate := &cobra.Command{
		Use:     "apps:create",
		Short:   "Create an app",
		Aliases: []string{"create"},
		Run: func(cmd *cobra.Command, args []string) {
			racer.UI().Asker().AskMissingString(&newApp.Name, "Application Name: ")
			if newApp.Name == "" {
				fmt.Println("Error: app name is required")
				return
			}

			fmt.Printf("=> creating app (%s) \n", newApp.Name)
			racer.UI().Printer().AddTitle(fmt.Sprintf("%s (app)", newApp.Name))
			table := new(ui.Table)
			table.AddRow("Name:", newApp.Name)
			table.AddRow("Size:", newApp.Size)
			table.AddRow("Deployed:", newApp.Deployed)
			racer.Printer().Add(table).Print()
		},
	}
	appsCreate.Flags().StringVarP(&newApp.Name, "name", "a", "", "App name")
	racer.AddCommand(appsCreate)

	racer.AddTopic("clusters", "create, teardown and manage clusters", false)
	clusters := &cobra.Command{
		Use:   "clusters",
		Short: "Manage clusters",
		Run: func(cmd *cobra.Command, args []string) {
			racer.Printer().AddTitle("Clusters")
			table := ui.NewTable("NAME", "DATACENTER", "SIZE")
			for _, c := range clusters {
				table.AddRow(c.Name, c.Datacenter, c.Size)
			}
			racer.Printer().Add(table).Print()
		},
	}
	racer.AddCommand(clusters)

	clusterLaunch := &cobra.Command{
		Use:   "clusters:launch",
		Short: "Launch a cluster",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("inside clusters:launch with Args: %v\n", args)
		},
	}
	var dc string
	clusterLaunch.Flags().StringVarP(&dc, "datacenter", "d", "", "datacenter for the cluster")
	racer.AddCommand(clusterLaunch)

	// Global flags
	var host string
	root.PersistentFlags().StringVarP(&host, "server", "s", "", "The address and port of the ovrclk API server")

	if err := racer.Execute(); err != nil {
		fmt.Println(err)
	}
}
