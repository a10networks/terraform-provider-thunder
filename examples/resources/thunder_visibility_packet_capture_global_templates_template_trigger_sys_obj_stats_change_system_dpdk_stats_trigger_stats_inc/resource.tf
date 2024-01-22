provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_dpdk_stats_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_dpdk_stats_trigger_stats_inc" {

  name                = "test"
  err_pkt_drop        = 1
  io_ring_drop        = 1
  io_rx_que_drop      = 1
  io_tx_que_drop      = 1
  pkt_drop            = 1
  pkt_lnk_down_drop   = 1
  rx_align_err        = 1
  rx_crc_err          = 1
  rx_csum_offload_err = 1
  rx_err              = 1
  rx_frame_err        = 1
  rx_len_err          = 1
  rx_long_len_err     = 1
  rx_miss_err         = 1
  rx_no_buff_err      = 1
  rx_over_err         = 1
  rx_short_len_err    = 1
  tx_abort_err        = 1
  tx_carrier_err      = 1
  tx_drop             = 1
  tx_err              = 1
  tx_fifo_err         = 1
  tx_hbeat_err        = 1
  tx_windows_err      = 1
  w_link_down_drop    = 1
  w_ring_drop         = 1
  w_tx_que_drop       = 1
}