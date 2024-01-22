provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_policy_connection_load" "thunder_gslb_policy_connection_load" {
  name                       = "test"
  connection_load_enable     = 1
  connection_load_fail_break = 1
  limit                      = 1
  connection_load_limit      = 128313
  connection_load_interval   = 3
  connection_load_samples    = 6

}