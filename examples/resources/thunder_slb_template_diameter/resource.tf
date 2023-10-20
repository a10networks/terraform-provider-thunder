provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_diameter" "test_thunder_slb_template_diameter" {
  name                       = "testing_diameter"
  customize_cea              = 1
  avp_code                   = 122
  avp_string                 = "test"
  service_group_name         = "sg1"
  dwr_time                   = 100
  idle_timeout               = 5
  multiple_origin_host       = 1
  origin_realm               = "test_realm"
  product_name               = "test_product"
  vendor_id                  = 12345
  session_age                = 456
  dwr_up_retry               = 2
  terminate_on_cca_t         = 1
  forward_unknown_session_id = 1
  forward_to_latest_server   = 1
  load_balance_on_session_id = 1
  message_code_list {
    message_code = 123
  }
  avp_list {
    avp       = 123
    int32     = 12
    mandatory = 1
  }
  origin_host {
    origin_host_name = "test_host"
  }
  user_tag = "tet_user"

}
