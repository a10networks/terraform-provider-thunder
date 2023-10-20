provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_partition" "partition" {
  user_tag         = "tag1"
  partition_name   = "part8"
  application_type = "adc"
  id1              = 8
}