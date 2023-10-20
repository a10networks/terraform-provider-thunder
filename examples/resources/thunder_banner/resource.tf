provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_banner" "bannerTest" {
  exec_banner_cfg {
    exec        = 1
    exec_banner = "Sample Exec Banner"
  }
  login_banner_cfg {
    login        = 1
    login_banner = "Sample Login Banner"
  }

}

