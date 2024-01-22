provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_unnumbered_use_source_ip" "thunder_ip_unnumbered_use_source_ip" {
  update_source_ip = "10.10.10.10"
}