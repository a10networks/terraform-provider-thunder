provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_sctp_permit_payload_protocol_protocol_name" "thunder_cgnv6_sctp_permit_payload_protocol_protocol_name" {
  protocol = "iua"
}