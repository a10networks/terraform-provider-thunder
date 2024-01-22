provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_authentication" "thunder_authentication" {
  enable_cfg {
    enable_auth_type = "local"
  }
  login_cfg {
    privilege_mode = 1
    local          = 1
  }
  mode_cfg {
    mode = "single"
  }
  multiple_auth_reject = 1
  type_cfg {
    type = "local"
  }
}