provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ng_waf" "thunder_ng_waf" {
  cpu {
  }
  custom_page {
  }
  custom_signals {
  }
  stats_list {
    name     = "a10"
    user_tag = "75"
  }
  status {
  }
}