provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_common_quic" "thunder_slb_common_quic" {
  cid_len          = 4
  cpu_offset       = 12
  enable_hash      = 0
  enable_signature = 0
  quic_lb_offset   = 8
  signature        = "122"
  signature_len    = 3
  signature_offset = 4
}