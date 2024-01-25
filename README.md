# A10’s Thunder Terraform Provider Introduction.

Welcome TTP 1.4.1 Latest Version.

Thunder Terraform Provider is a custom plugin to configure thunder using terraform as a IaC tool and terraform scripts[.tf] simplifies applying configuration on thunder. You can configure or de-configure thunder settings.

This plugin contains several configurations of thunder which can be applied via out of box examples provided.
Terraform provider plugin will only configure thunder via axapi, It will not install thunder.

## Support Matrix

| ACOS | [TTP 1.0.0](https://github.com/a10networks/terraform-provider-thunder/tree/v1.0.0) | [TTP 1.1.0](https://github.com/a10networks/terraform-provider-thunder/tree/v1.1.0) | [TTP 1.2.1](https://github.com/a10networks/terraform-provider-thunder/tree/v1.2.1) | [TTP 1.2.2](https://github.com/a10networks/terraform-provider-thunder/tree/v1.2.2) | [TTP 1.3.0](https://github.com/a10networks/terraform-provider-thunder/tree/v1.3.0) | [TTP 1.4.1](https://github.com/a10networks/terraform-provider-thunder/tree/v1.4.1) |
| :--------: | :-------: | :-------:  | :-------: | :-------: | :-------: | :-------: |
| `ACOS version 6.0.2-p1` | `No`  | `No`  | `No`| `No` | `No` | `Yes` |
| `ACOS version 6.0.1` | `No`  | `No`  | `Yes`| `Yes` | `Yes` | `No` |
| `ACOS version 6.0.0-p2` | `No`  | `No`  | `Yes`| `Yes` | `Yes` | `No` |
| `ACOS version 6.0.0-p1` | `No`  | `No`  | `Yes`| `Yes` | `Yes` | `No` |
| `ACOS version 5.2.1-p6` | `No`  | `Yes` | `No` | `No` | `No` | `No` |
| `ACOS version 5.2.1-p5`| `No`  | `Yes` | `No` | `No` | `No` | `No` |
| `ACOS version 5.2.1-p4` | `Yes` | `No`  | `No` | `No` | `No` | `No` |
| `ACOS version 5.2.1-p3` | `Yes` | `No`  | `No` | `No` | `No` | `No` |

## Release Logs

### TTP-1.4.1
- Support for ACOS v6.0.2-p1
- Supports total 2712 resources with examples

### TTP-1.3.0
- Support for ACOS v6.0.1
- Jenkins and Tekton integration examples
- Additional resources:

  1. Cloud Services Cloud Provider AWS Log
  2. Cloud Services Cloud Provider Azure Log
  3. Cloud Services Cloud Provider VMware Log
  4. Cloud Services Cloud Provider AWS Metrics
  5. Cloud Services Cloud Provider Azure Metrics
  6. Cloud Services Cloud Provider VMware Metrics
  7. Service Partition
  8. System App Performance Stats
  9. System Throughput Stats
  10. System Memory Oper
  11. System Hardware Oper
  12. Interface Available Eth List Oper
  13. Sessions Oper
  14. Plat CPU Packet Oper
  15. SLB HTTP Proxy Oper
  16. SLB SSL Stats Oper
  17. SLB Server Oper
  18. Clock Show Oper
  19. Syslog Oper

### TTP-1.2.2
- Support for MAC M1 processors
- Support for ACOS 6.0.0-p1 and ACOS 6.0.0-p2
- Fixed Change Optional to Required for partition, slb server, and slb service-group
- Defect Fixtures

### TTP-1.2.1
- Support for ACOS 6.0.0-p1
- Defect Fixtures
### TTP-1.1.0
- Extended Support for GSLB AXAPIs
- Extended Support for Geo-Location AXAPIs
- Extended Support for Password Change AXAPIs
- Defect Fixtures

## How it works
   1. Install terraform on your local OS, Please refer below sections for more details.
   2. Search required terraform configuration from examples. In case not found create a new one, Please refer below sections for more details.
   3. Execute terraform scripts to apply thunder configuration, Please refer below sections for more details.
   4. Verify thunder configuration after terraform apply, Please refer below sections for more details.

## How to Install Terraform on Ubuntu:

    To install Terraform on Ubuntu, perform the following steps:

    1. Run the following commands to download and install Terraform 1.5.6.
        a.	wget https://releases.hashicorp.com/terraform/1.5.6/terraform_1.5.6_linux_amd64.zip
        b.	unzip terraform_1.5.6_linux_amd64.zip
        c.	mv terraform /usr/local/bin/
    2. Verify installation using below command:
        a.	terraform -version

    For more information, please visit : https://www.terraform.io/downloads.html

## How to Install Terraform on Windows:

    To install Terraform on Windows, perform the following steps:

    1. Download windows installable from:
        https://developer.hashicorp.com/terraform/downloads
        https://releases.hashicorp.com/terraform/1.5.6/terraform_1.5.6_windows_386.zip
        https://releases.hashicorp.com/terraform/1.5.6/terraform_1.5.6_windows_amd64.zip
    2. Extract at C:/Terraform1.5.6
    3. Update environment variable 'path and add 'C:/Terraform1.5.6/'

    For more information, please visit: https://developer.hashicorp.com/terraform/downloads


## How to Install Terraform on MacOS:

    To install Terraform on MacOS, perform the following steps:
    Run the following commands to download and install the latest version of Terraform:

    1. brew tap hashicorp/tap
    2. brew install hashicorp/tap/terraform

    For more information, please visit: https://developer.hashicorp.com/terraform/downloads

## How to Search Thunder Configuration

To search for a Thunder Configuration in the existing examples, perform the following steps:

  1. Search the required Terraform configuration script directory navigate to examples > resources directory.

      **Example:**

      If you want to apply the bgp router configuration on Thunder, search for the thunder_router_bgp directory under the resources directory.

  2. Open the Terraform script from the directory.

      **Example:**

      Open resource.tf script under the thunder_router_bgp directory.

  3. Update the Thunder IP address and login credentials depending on the type of Thunder installed and review the Thunder configurations in the Terraform script.

```
provider "thunder" {
    address = var.dut9049
    username = var.username
    password = var.password
}

resource "thunder_router_bgp" "bgp1" {
    as_number = 101
    neighbor {
        ipv4_neighbor_list {
            neighbor_ipv4 = "10.1.1.104"
            activate =  1
            nbr_remote_as = 104
            allowas_in =  1
            allowas_in_count =  10
            graceful_restart =  1
        }
    }
}

resource "thunder_router_bgp" "bgp2" {
    provider = thunder.L3V_A
    as_number = 201
    neighbor {
        ipv4_neighbor_list {
            neighbor_ipv4 = "10.1.1.204"
            activate =  1
            nbr_remote_as = 204
            allowas_in =  1
            allowas_in_count =  10
            graceful_restart =  1
        }
    }
}

```

4. Add, modify, or remove the Thunder configuration parameters and their corresponding values as appropriate.

5. Save the changes.

6. Identify the compatible Terraform Provider version for your installed ACOS version from Support Matrix Section.

7. Go to [A10 Networks Terraform Thunder Provider](https://registry.terraform.io/providers/a10networks/thunder/latest/docs).

8. Select your Terraform Provider version from the Latest Version drop-down.

9. Expand Resources from the left-panel to select the required Thunder configuration resource.

10. Copy and paste the following Terraform Provider configuration from the USE PROVIDER drop-down to your Terraform script.

```
terraform {
  required_providers {
    thunder = {
      source = "a10networks/thunder"
      version = "1.4.1"
    }
  }
}
```

11. Save the Terraform script.


## How to create new terraform thunder configuration example.

  Here are step-by-step instructions for creating a terraform thunder configuration example.
  For example if you want to apply bgp router configuration on thunder and which doesn't exist in examples.

1. Create a new directory to house your terraform configuration files.

```
  mkdir thunder_router_bgp
  cd thunder_router_bgp
```

2. Create a `.tf` file, such as `thunder_router_bgp.tf`, in your "thunder_router_bgp" directory. In this file, define the  Thunder ROUTER BGP configurations using the A10Networks Thunder Provider. Refer to the official documentation: https://registry.terraform.io/providers/a10networks/thunder/latest/docs/resources/router_bgp  for the required resource and parameters.

   Here is basic example:
  ```
    terraform {
      required_providers {
        thunder = {
          source = "a10networks/thunder"
          version = "1.4.1" # Replace with your desired provider version
        }
      }
    }

    provider "a10networks" {
      host          = "10.10.10.10" # Replace with your desired your thunder device ip
      username      = "admin"    # Replace with your desired your thunder user name
      password      = "password"    # # Replace with your password
    }

    resource "thunder_router_bgp" "thunder_router_bgp" {
    as_number = 101      # Replace with your desired bgp number
    neighbor {
      ipv4_neighbor_list {
        neighbor_ipv4 = "10.1.1.104"   # Replace with your desired neighbor_ipv4 address
        activate =  1                  # Replace with your desired activate value
        nbr_remote_as = 104            # Replace with your desired nbr_remote_as value
        allowas_in =  1                # Replace with your desired allowas_in value
        allowas_in_count =  10         # Replace with your desired allowas_in_count value
        graceful_restart =  1          # Replace with your desired graceful_restart value
          }
      }
    }

  ```
  Adjust the BGP configuration parameters as needed.

## How to execute TTP Configuration from Terraform CLI

To apply the Thunder Terraform configurations using Terraform CLI, perform the following steps:

1. Go to your required Terraform configuration script directory from the command prompt.

2. If the Terraform Provider is not initialize, run the following command to initialize the working directory and download the A10 Networks Terraform Thunder Provider.
```
terraform init
```

If Terraform Provider is already initialize with an older version, run the following command to upgrade the Terraform Provider to the latest version:

```
terraform init -upgrade
```

3. Create an execution plan and ensure that it matches your desired configurations.

```
terraform plan
```

4. Apply the specified configuration on your Thunder device.

```
terraform apply
```


## How to execute TTP Configuration from Jenkins.

Terraform scripts can be executed via Jenkins pipeline.

Please refer the `/devops/jenkins/Jenkins Pipeline Configuration Guide.pdf` file to configure Jenkins pipeline.

This comprehensive document offers step-by-step instructions and best practices for automating TTP (Thunder Terraform Provider) configurations, ensuring efficient management within your Jenkins environment.

## How to execute TTP Configuration from Teckton.

Terraform scripts can be executed via Tekton pipeline.

Please refer the `/devops/tekton/Tekton Pipeline Configuration Guide.pdf` file to configure Tekton pipeline.

This comprehensive document offers step-by-step instructions and best practices for automating TTP (Thunder Terraform Provider) configurations, ensuring efficient management within your Tekton environment.

## How to verify on Thunder

  To verify the applied configurations, follow below steps:

  1. SSH into the Thunder device using your username and password.
  2. Once connected, enter the following commands:

     1. `enable`

        ![image](https://github.com/smundhe-a10/terraform-provider-thunder/assets/107971633/7e532cee-fa8e-4af7-aa50-da56a24dd4c3)

     3. `show running-config`

        ![image](https://github.com/smundhe-a10/terraform-provider-thunder/assets/107971633/ae37e53d-c650-43f0-b71f-2416f4e5d65a)



## How to contribute

If you have created a new example, please save the Terraform file with a resource-specific name, such as "thunder_bgp.tf."

1. Clone the repository.
2. Copy the newly created file and place it under the /examples/resource directory.
3. Create a MR against the master branch.

### Pre-requisite

- [Terraform](https://www.terraform.io/downloads.html) 1.0.x
- [Go](https://golang.org/doc/install) 1.16 (to build the provider plugin)


### How to Build The Provider

```sh
$ git clone git@github.com:a10networks/terraform-provider-thunder.git
$ cd terraform-provider-thunder
$ make build
```


### Local Plugin Installation

After building new provider we need to install and test the plugin.

We need to create a new version please follow below step.

  1. $ cd terraform-provider-thunder
  2. $ go build -o terraform-provider-thunder
  3. $ mkdir -p ~/.terraform.d/plugins/a10networks.com/a10networks/thunder/7.7.7/linux_amd64/
  4. $ cp terraform-provider-thunder ~/.terraform.d/plugins/a10networks.com/a10networks/thunder/7.7.7/linux_amd64/
  5. create version.tf file like:
  ```
  terraform {
    required_providers {
      thunder = {
        source  = "a10networks.com/a10networks/thunder"
        version = "7.7.7"
      }
    }
  }
  ```

If you face some dependency issue try `$ go mod tidy` or `$ go mod vendor`

*Note:* After cloning you can also run `$ make local` to perform these above step automatically for dummy version 7.7.7 in Linux env. Else if you have already executed manual steps please ignore this.

 If raise MR request for contributing.

## Documentation

Terraform provider documentation is available at https://registry.terraform.io/providers/a10networks/thunder/latest/docs

A10 Thunder AXAPI support documentation available at https://documentation.a10networks.com/docs/IaC/Terraform/thunder-terraform-provider/1-4-1/

## Report a Issue

Please raise issue in github repository.
Please include the Terraform script that demonstrates the bug and the command output and stack traces will be helpful.

## Support
Please reach out at support@a10networks.com with "a10-terraform-provider" in the subject line.
