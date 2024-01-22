provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_inside_source_class_list" "thunder_ip_nat_inside_source_class_list" {
}