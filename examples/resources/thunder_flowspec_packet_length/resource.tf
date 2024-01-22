provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_flowspec_packet_length" "thunder_flowspec_packet_length" {

  name                    = "test"
  length                  = 40
  packet_length_attribute = "eq"
}