provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_unnumbered" "thunder_ipv6_unnumbered" {
  use_source_ipv6 {
    update_source_ipv6 = "363:c5cb:6f65:3854:82f7:4d31:3623:ec50"
  }
}