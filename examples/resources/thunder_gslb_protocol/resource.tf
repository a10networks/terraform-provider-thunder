provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_gslb_protocol" "thunder_gslb_protocol_test" {
  auto_detect        = 1
  msg_format_acos_2x = 1
  status_interval    = 1000
  use_mgmt_port      = 1
  secure {
    action = "enable"
  }
  limit {
    ardt_query    = 100
    ardt_response = 1100
    ardt_session  = 1100
    conn_response = 1000
    response      = 1000
    message       = 1000

  }
  enable_list {
    type = "controller"
  }
}
