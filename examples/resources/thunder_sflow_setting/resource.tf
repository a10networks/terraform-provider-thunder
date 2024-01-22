provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_setting" "thunder_sflow_setting" {
  append_mapping_info         = 1
  counter_polling_interval    = 20
  default_counter_polling_mtu = 1500
  local_collection            = "enable"
  management_link_utilization = 5733
  packet_sampling_rate        = 5733
  source_ip_use_mgmt          = 1
}