provider "thunder" {
    address  = var.dut
    username = var.username
    password = var.password
    partition = var.partition_name # Into which partition. The default value is "shared".
}
