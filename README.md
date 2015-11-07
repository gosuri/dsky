# racer

Racer is a helper library for [Cobra](https://github.com/spf13/cobra), a library for creating powerful modern CLI applications. It powers almost all of [OvrClk](http://ovrclk.com)'s tools.

## Features

Racer currently supports [Cobra 0.0.9](https://github.com/spf13/cobra) and adds the below additionaly functionality:

* Topic namespaced subcommand `foo:bar` for easy.
* Automatic help generation, provides `-h` `--help` flags.
* Tabled Output formatting helpers
* Progress Bar
* Interactive User Inputs

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

## Examples

See examples under [/examples](examples) for full examples
