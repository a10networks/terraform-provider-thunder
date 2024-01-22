provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_policy_forward_policy" "thunder_slb_template_policy_forward_policy" {

  name           = "test-policy"
  acos_event_log = 0
  san_filtering {
    ssli_url_filtering_san = "enable-san"
  }
}