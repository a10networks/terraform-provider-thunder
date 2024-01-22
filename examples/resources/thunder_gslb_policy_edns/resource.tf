provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_policy_edns" "thunder_gslb_policy_edns" {
  name                     = "test"
  client_subnet_geographic = 1

}