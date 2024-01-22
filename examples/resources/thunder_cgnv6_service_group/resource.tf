provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_service_group" "thunder_cgnv6_service_group" {

  name = "testing"
  member_list {
    name     = "test12345"
    port     = 25499
    user_tag = "33"
    sampling_enable {
      counters1 = "all"
    }
  }
  protocol = "tcp"
  sampling_enable {
    counters1 = "all"
  }
  traffic_replication_mirror_ip_repl = 1
  user_tag                           = "31"
}