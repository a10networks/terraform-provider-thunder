

provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_reboot"   "Reboot"  {
    
    all = 1
  
  }
  
