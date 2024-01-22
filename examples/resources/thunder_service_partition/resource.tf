provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_service_partition" "thunder_service_partition" {
  application_type = "adc"
  follow_vrid      = 0
  id1              = 1
  partition_name   = "14"
  user_tag         = "52"
}