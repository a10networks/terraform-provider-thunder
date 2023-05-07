# A10 Networks Terraform Provider Release v1.2.2.

Welcome to GitHub repository for A10’s Registered Terraform Provider for vThunder.

## What is vThunder Terraform Provider?

This collection of sample terraform resources [tf scripts] along with examples will help you get started with ACOS-vThunder AXAPI/v3 cofiguration.
Terraform scripts simplifies appliying cofiguration on vThunder. You can configure or de-configure vthunder settings. 
Terraform provider plugin will only configure vThunder via axapi, It will not install vThunder.

## A10’s vThunder Support Information

Below listed A10’s vThunder vADC (Application Delivery Controller) are tested and supported.
- 64-bit Advanced Core OS (ACOS) version 6.0.0-p1, build 47.
- 64-bit Advanced Core OS (ACOS) version 6.0.0-p2, build 22.

## TTP Support Matrix 

| ACOS | [TTP 1.0.0](https://github.com/a10networks/terraform-provider-thunder/tree/v1.0.0) | [TTP 1.1.0](https://github.com/a10networks/terraform-provider-thunder/tree/v1.1.0) | [TTP 1.2.1](https://github.com/a10networks/terraform-provider-thunder/tree/v1.2.1) | [TTP 1.2.2](https://github.com/a10networks/terraform-provider-thunder/tree/v1.2.2) | 
| :--------: | :-------: | :-------:  | :-------: | :-------: | 
| `5.2.1-p3-build 70` | `Yes` | `No`  | `No` | `No` |
| `5.2.1-p4-build 70` | `Yes` | `No`  | `No` | `No` |
| `5.2.1-p5-build 114`| `No`  | `Yes` | `No` | `No` |
| `5.2.1-p6-build 74` | `No`  | `Yes` | `No` | `No` |
| `6.0.0-p1-build 47` | `No`  | `No`  | `Yes`| `Yes` |
| `6.0.0-p1-build 22` | `No`  | `No`  | `Yes`| `Yes` |

## Release Logs Information

- Support for MAC M1 processors
- Support for ACOS 6.0.0-p1 AXAPIs
- Support for ACOS 6.0.0-p2 AXAPIs
- Fixed Change Optional to Required for partition, slb server, and slb service-group
- Defect Fixtures

# Terraform Provider Thunder

A [Terraform](https://www.terraform.io) provider for A10 Thunder.

- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
dd- Terraform Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

# Maintainers

This provider plugin is maintained by the Cloud team at [A10 Networks](https://www.a10networks.com/).

# Requirements

- [Terraform](https://www.terraform.io/downloads.html) 1.0.x
- [Go](https://golang.org/doc/install) 1.16 (to build the provider plugin)


# Building The Provider 

```sh
$ git clone git@github.com:a10networks/terraform-provider-thunder.git
$ make build
```

# Building The Provider in Go Work-Space

```sh
$ mkdir -p $GOPATH/src/github.com/a10networks/
$ cd  $GOPATH/src/github.com/a10networks/
$ git clone git@github.com:a10networks/terraform-provider-thunder.git
```

Build the provider

```sh
$ make build
```

# Local Plugin Installation

For Thunder plugin Installation without fetching it from A10 Networks Namespace.  

Inside your cloned repo; here x.y.z is version example 0.4.5
1. $ go build -o terraform-provider-thunder
2. $ mkdir -p ~/.terraform.d/plugins/a10networks.com/a10networks/thunder/x.y.z/linux_amd64/
3. $ cp terraform-provider-thunder ~/.terraform.d/plugins/a10networks.com/a10networks/thunder/x.y.z/linux_amd64/
4. create version.tf file like:
```
terraform {
  required_providers {
    thunder = {
      source  = "a10networks.com/a10networks/thunder"
      version = "x.y.z"
    }
  }
}
```

If you face some dependency issue try `$ go mod tidy` or `$ go mod vendor`

*Note:* After cloing you can also run `$ make local` to perform these above step automatically for dummy version 7.7.7 in Linux env.

# Documentation

Documentation for the A10 Thunder Terraform integration is available at [link]
Terraform provider documentation is available at https://registry.terraform.io/providers/a10networks/thunder/latest/docs

A10 Thunder AXAPI support documentation available at https://documentation.a10networks.com/docs/IaC/Terraform/thunder-terraform-provider/1-2-0/

# Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-thunder
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

## Samples
See the examples directory for various LB topologies that can be driven from this terraform provider.

## Bug Reporting and Feature Requests
Please submit bug reports and feature requests via GitHub issues. When reporting bugs, please include the Terraform script that demonstrates the bug and the command output. Stack traces will be helpful. Please ensure any sensitive information is redacted as Issues and Pull Requests are publicly viewable.

## Contact
If you have a question that cannot be submitted via Github Issues, please email support@a10networks.com with "a10-terraform-provider" in the subject line.

