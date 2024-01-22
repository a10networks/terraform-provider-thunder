provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_distributed_forwarding_fw" "thunder_scaleout_distributed_forwarding_fw" {
  fw_value                  = "enable"
  session_offload_direction = "both"
  threshold {
    threshold_value = 51
    protocol_value  = "TCP"
  }
}