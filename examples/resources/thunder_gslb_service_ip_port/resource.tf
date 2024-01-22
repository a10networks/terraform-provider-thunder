provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_service_ip_port" "thunder_gslb_service_ip_port" {

  node_name                     = "vs2"
  action                        = "enable"
  health_check                  = "ping"
  health_check_protocol_disable = 1
  port_num                      = 30672
  port_proto                    = "tcp"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "60"
}