
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_terminal"   "Terminal"  {
    auto_size = 0
    editing = 1
    gslb_cfg {
      gslb_prompt = 0
    }
    history_cfg {
      enable = 1
      size = 100
    }
    idle_timeout= 0
    length = 110
    prompt_cfg {
      prompt = 0
    }
    width = 100
   
  
}