provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_action_list" "thunder_ddos_action_list" {
  name = "DDOS_ACTION_LIST"
  action {
    action = "ignore"
  }
  capture_config = "10"
  user_tag       = "test_user_tag"
  zone_template {
    logging = "testLog"
    encap   = "testEncap"
  }
}