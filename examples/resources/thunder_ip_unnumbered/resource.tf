provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_unnumbered" "thunder_ip_unnumbered" {

  use_source_ip {
    update_source_ip = "10.10.10.10"
  }
}