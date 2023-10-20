provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_service_partition" "test"{
  partition_name = "svcpartition"
  application_type = "adc"
  id1 = 22
}