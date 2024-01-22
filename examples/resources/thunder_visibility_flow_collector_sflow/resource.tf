provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_flow_collector_sflow" "thunder_visibility_flow_collector_sflow" {
  sampling_enable {
    counters1 = "all"
  }
}