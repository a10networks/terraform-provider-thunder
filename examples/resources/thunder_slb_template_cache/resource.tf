provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_cache" "testcache" {
  name                    = "cache1"
  user_tag                = "cache_testing"
  accept_reload_req       = 1
  age                     = 1
  replacement_policy      = "LFU"
  remove_cookies          = 1
  verify_host             = 1
  default_policy_nocache  = 1
  disable_insert_age      = 1
  disable_insert_via      = 1
  max_cache_size          = 1
  max_content_size        = 0
  min_content_size        = 0
  packet_capture_template = "chache11"
  logging                 = "log1"
  local_uri_policy {
    local_uri = "/abc"
  }
  uri_policy {
    uri          = "/list"
    cache_action = "nocache"
    cache_value  = 0
  }
  sampling_enable {
    counters1 = "all"
  }
}