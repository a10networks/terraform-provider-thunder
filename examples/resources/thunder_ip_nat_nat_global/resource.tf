provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_nat_global" "thunder_ip_nat_nat_global" {
  sampling_enable {
    counters1 = "all"
  }
}