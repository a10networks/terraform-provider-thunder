provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dns_response_rate_limiting" "thunder_slb_template_dns_response_rate_limiting" {

  name                 = "test"
  action               = "rate-limit"
  enable_log           = 0
  filter_response_rate = 10
  match_subnet         = "255.255.255.255"
  match_subnet_v6      = 128
  response_rate        = 5
  rrl_class_list_list {
    name     = "class-test"
    user_tag = "43"
    lid_list {
      lidnum              = 758
      lid_response_rate   = 5
      lid_slip_rate       = 5
      lid_match_subnet    = "255.255.255.255"
      lid_match_subnet_v6 = 128
      lid_window          = 1
      lid_src_ip_only     = 0
      lid_enable_log      = 0
      lid_action          = "rate-limit"
      user_tag            = "16"
    }
  }
  src_ip_only = 0
  tc_rate     = 8
  window      = 1
}