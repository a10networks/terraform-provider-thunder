provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_policy_forward_policy_action" "thunder_slb_template_policy_forward_policy_action" {

  name               = "test-policy"
  drop_message       = "513"
  drop_response_code = 350
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "69"
}