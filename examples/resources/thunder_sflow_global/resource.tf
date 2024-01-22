provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_global" "thunder_sflow_global" {
  sampling_enable {
    counters1 = "all"
  }
}