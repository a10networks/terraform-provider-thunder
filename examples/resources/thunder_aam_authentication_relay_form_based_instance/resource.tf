provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_relay_form_based_instance" "thunder_aam_authentication_relay_form_based_instance" {

  name = "test"
  request_uri_list {
    match_type              = "equals"
    uri                     = "39"
    user_variable           = "63"
    password_variable       = "30"
    domain_variable         = "4"
    other_variables         = "68"
    max_packet_collect_size = 1048576
    cookie {
      cookie_value {
        cookie_value = "20"
      }
    }
    action_uri = "52"
    user_tag   = "test"
  }
  sampling_enable {
    counters1 = "all"
  }

}