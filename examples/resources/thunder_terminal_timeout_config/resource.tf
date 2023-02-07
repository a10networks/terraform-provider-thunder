provider "thunder" {
    address  = var.dut9049
    username = var.username
    password = var.password
}

resource "thunder_terminal" "thunder_terminal_test"{
    idle_timeout = 30
    length = 100
    width = 120
    editing = 1
    gslb_cfg {
      gslb_prompt = 1
      symbol = 1
      group_role = 1
    }
    history_cfg {
      enable = 1
      size = 110
    }
    prompt_cfg {
      hostname = 1
      prompt = 1
    }
    auto_size = 1
  }