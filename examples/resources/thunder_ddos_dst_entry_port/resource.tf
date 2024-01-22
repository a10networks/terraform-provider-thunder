provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_port" "thunder_ddos_dst_entry_port" {

  capture_config {
    capture_config_name = "10"
    capture_config_mode = "drop"
  }
  dst_entry_name   = "test"
  port_num         = 22
  protocol         = "tcp"
  detection_enable = 1
  enable_top_k     = 1
  topk_num_records = 22
  deny             = 1
  sflow {
    polling {
      sflow_tcp {
        sflow_tcp_basic = 1
      }
    }
  }

  user_tag = "abcd"

}