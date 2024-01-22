provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_sixrd_fragmentation_inbound" "thunder_cgnv6_sixrd_fragmentation_inbound" {

  action = "drop"

}