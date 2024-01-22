provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_sampling" "thunder_sflow_sampling" {
  eth_list {
    eth_start = 2
    eth_end   = 2
  }
}