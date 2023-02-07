---
subcategory: ""
page_title: "Usage of file_bw_list - Thunder Provider"
description: |-
    Usage of file_bw_list.
---
 
# Usage of file_bw_list

In this provdier, we can use resource `thunder_file_bw_list` to do the operation of thunder CLI command `import bw-list`.
The imported file will be stored as a file-type bw-list.
You can check those files by

```
TH9049(config)#sh bw-list 
Name              Url                               Size(Byte)   Date           
--
TEST              Local                                     23   Feb/15 22:57:10
TEST_2            Local                                     23   Feb/15 23:52:02
Total: 2
```

# Example of Bw-list Files

## Format
```
1.2.0.0/15
3.7.10.0/24
```

```terraform
provider "thunder" {
    address  = var.dut9049
    username = var.username
    password = var.password
}

provider "thunder" {
    alias = "L3V_A"
    address  = var.dut9049
    username = var.username
    password = var.password
    partition = var.l3v_1
}

resource "thunder_file_bw_list" "BWLIST_1" {
    name = "BWLIST_1"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/BWLIST_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_bw_list" "L3V_BWLIST_1" {
    provider = thunder.L3V_A
    name = "BWLIST_1"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/BWLIST_1.txt"
    use_mgmt_port = 1
}
```
