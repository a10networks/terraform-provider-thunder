provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_axdebug_filter_config" "thunder_axdebug_filter_config" {
  dst              = 1
  dst_ip           = 1
  dst_ipv4_address = "10.10.10.10"
  number           = 22
  port             = 1
  port_num_max     = 32744
  port_num_min     = 18940
  user_tag         = "35"

}