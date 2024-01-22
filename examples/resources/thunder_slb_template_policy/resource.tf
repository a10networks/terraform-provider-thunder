provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_policy" "thunder_slb_template_policy" {

  name = "test-policy"
  class_list {
    name = "test"
    lid_list {
      lidnum = 517
    }
  }
  interval           = 131
  over_limit         = 1
  over_limit_logging = 1
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "66"
}