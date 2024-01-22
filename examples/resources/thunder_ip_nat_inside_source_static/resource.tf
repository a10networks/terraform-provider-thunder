provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_inside_source_static" "thunder_ip_nat_inside_source_static" {
  action      = "enable"
  nat_address = "10.10.10.10"
  src_address = "10.10.10.10"

}