provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_service_group_member" "thunder_aam_authentication_service_group_member" {

  member_priority = 16
  member_state    = "enable"
  name            = "test"
  port            = 1812
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "test1"

}