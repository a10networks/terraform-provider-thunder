provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_terminal" "thunder_terminal" {

  auto_size = 1
  editing   = 1
  gslb_cfg {
    gslb_prompt = 1
    disable     = 1
    group_role  = 1
    symbol      = 1
  }
  history_cfg {
    size = 256
  }
  idle_timeout = 15
  prompt_cfg {
    prompt    = 1
    ha_status = 1
    vcs_cfg {
      vcs_status = 1
    }
  }

}