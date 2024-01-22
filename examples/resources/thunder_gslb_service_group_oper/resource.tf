provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_service_group_oper" "thunder_gslb_service_group_oper" {
  oper {
    session_list {
      hits             = hits
      last_second_hits = last_second_hits
      update           = update
      aging            = aging
    }
    matched        = matched
    total_sessions = total_sessions
  }
  service_group_name = "30"
}
output "get_gslb_service_group_oper" {
  value = ["${data.thunder_gslb_service_group_oper.thunder_gslb_service_group_oper}"]
}