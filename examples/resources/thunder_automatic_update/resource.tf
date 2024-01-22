provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_automatic_update" "thunder_automatic_update" {
  config_list {
    feature_name = "ca-bundle"
    schedule     = 1
    weekly       = 1
    week_day     = "Wednesday"
    week_time    = "11:11"
  }
  proxy_server {
    proxy_host    = "85"
    https_port    = 57158
    auth_type     = "basic"
    username      = "69"
    password      = 1
    secret_string = "asd"
  }
  use_mgmt_port = 1
}