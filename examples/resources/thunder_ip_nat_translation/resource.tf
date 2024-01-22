provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_translation" "thunder_ip_nat_translation" {

  icmp_timeout {
    icmp_timeout = "fast"
  }
  ignore_tcp_msl = 1
  service_timeout_list {
    service_type = "tcp"
    port         = 57460
    timeout_type = "age"
    timeout_val  = 7756
  }
  tcp_timeout = 200
}