provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_capture_config" "thunder_visibility_packet_capture_capture_config" {
  name                             = "test"
  disable                          = 1
  concurrent_captures              = 33
  concurrent_captures_age          = 4
  number_of_packets_per_conn       = 22
  packet_length                    = 130
  file_size                        = 33
  file_count                       = 11
  number_of_packets_per_capture    = 20
  number_of_packets_total          = 22
  enable_continuous_global_capture = 1
  disable_auto_merge               = 1
  keep_pcap_files_after_merge      = 1
  user_tag                         = "37"
}