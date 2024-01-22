provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_flow_collector_netflow_template" "thunder_visibility_flow_collector_netflow_template" {
  sampling_enable {
    counters1 = "all"
  }
}