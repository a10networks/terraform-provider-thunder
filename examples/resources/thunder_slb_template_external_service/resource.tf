provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_external_service" "test_thunder_slb_template_external_service" {
  name = "ext_svc"
  type = "url-filter"
  request_header_forward_list {
    request_header_forward = "req_head"
  }
  failure_action = "drop"
  timeout        = 40
  action         = "drop"
  service_group  = "sg1"
  user_tag       = "test"
}