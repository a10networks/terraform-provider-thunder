provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_distributed_forwarding_cgn" "thunder_scaleout_distributed_forwarding_cgn" {
  cgn_value = "enable"
}