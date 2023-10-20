provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_common_buffer_threshold" "test_thunder_slb_common_buffer_threshold" {
  buff_thresh                = 1
  buff_thresh_hw_buff        = 200
  buff_thresh_relieve_thresh = 100
  buff_thresh_sys_buff_high  = 30
  buff_thresh_sys_buff_low   = 20
}

