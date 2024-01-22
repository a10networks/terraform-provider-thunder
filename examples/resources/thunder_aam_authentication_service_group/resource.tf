provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_service_group" "thunder_aam_authentication_service_group" {

  lb_method = "round-robin"
  member_list {
    name            = "test"
    port            = 1812
    member_state    = "enable"
    member_priority = 5
    user_tag        = "6"
    sampling_enable {
      counters1 = "all"
    }
  }
  name     = "test"
  protocol = "udp"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "test"
}