provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_cache" "thunder_slb_template_cache" {
  name                   = "temp1"
  accept_reload_req      = 1
  age                    = 36001
  default_policy_nocache = 1
  disable_insert_age     = 1
  disable_insert_via     = 1
  max_cache_size         = 232
  max_content_size       = 81292011
  min_content_size       = 5324222
  remove_cookies         = 0
  replacement_policy     = "LFU"
  sampling_enable {
    counters1 = "all"
  }
  user_tag    = "121"
  verify_host = 1
}