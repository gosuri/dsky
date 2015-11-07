# racer [![GoDoc](https://godoc.org/github.com/gosuri/racer?status.svg)](https://godoc.org/github.com/gosuri/racer) [![Build Status](https://travis-ci.org/gosuri/racer.svg?branch=master)](https://travis-ci.org/gosuri/racer)

Racer is a UI helper library for [Cobra](https://github.com/spf13/cobra), a library for creating powerful modern CLI applications. It powers almost all of [OvrClk](http://ovrclk.com)'s tools.

## Features

Racer currently supports [Cobra 0.0.9](https://github.com/spf13/cobra) and adds the below additionaly functionality:

* Command organization based on help topic `foo:bar`
* Tabled Output formatting helpers
* Interactive UI components - Normal Prompts, Password Prompts, Progress Bar

Cobra provides:

* Fully posix compliant flags (including short & long versions)
* Nested sub commands
* Global, local and cascading flags
* Easy generation of applications & commands with `cobra create appname` & `cobra add cmdname`
* Intelligent suggestions (`app srver`.. did you mean `app server`)
* Automatic help generation for commands and flags
* Automatic detailed help for `app help [command]`
* Automatic help flag recognition of `-h`, `--help`, etc.
* Automatically generated bash autocomplete for your application
* Automatically generated man pages for your application
* Command aliases so you can change things without breaking them
* The flexibilty to define your own help, usage, etc
* Optional tight integration with [viper](http://github.com/spf13/viper) for 12 factor apps

## Help Topics

Racer organizes help based on a namespace, the below is the output for OvrClk command utility built using racer:

```sh
$ ovrclk help

Utility to manage your clusters and applications on ovrclk

Usage: ovrclk COMMAND [<args>..] [options]

Primary help topics, type "ovrclk help TOPIC" for more details:

  apps        create, deploy and manage applications
  auth        login, logout and display vpn and token
  clusters    launch, teardown and manage clusters
  config      set and unset config variables for apps
  logs        display application and storage logs
  releases    deploy, revert and manage app releases
  storage     attach, detach and manage storage appliances
  users       add, remove and manage users

Additional topics:

  datacenters display datacenters
  regions     display regions
  version     display version

$ ovrclk clusters help

List all clusters

Usage: ovrclk clusters [options]

Options:

  -h, --help=false: help for clusters

General Options:

  -s, --server="": The address and port of the ovrclk API server
  -t, --token="": Bearer token for authenticating with the API server

Additional commands, type "ovrclk COMMAND --help" for more details:

  clusters:up [SELECTOR]                   Launch a new cluster. The selector could either be a region or a datacenter name
  clusters:info --cluster=<name>           Display info for the cluster
  clusters:down --cluster=<name>           Teardown the cluster
  clusters:rename NEWNAME --cluster=<name> Rename the cluster
  clusters:update --cluster=<name>         Update the cluster
```

## UI Components

### Progress Bar

```go
count := 5000
bar := racer.UI().NewProgressBar(count)
for i := 0; i < count; i++ {
  bar.Increment()
  time.Sleep(time.Millisecond)
}
bar.FinishPrint("The End!")
```

```sh
963 / 5000 [=====================>---------------------------------] 39.26 % 3s
```

### Password Prompter

```go
var password string
racer.UI().Prompter().PromptHiddenString(&password, "Password: ")
```

## More Examples

See examples under [/examples](examples) for full examples
