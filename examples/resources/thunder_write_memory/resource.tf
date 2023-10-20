
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_write_memory"   "WriteMemory"{
        destination =  "primary"
        partition = "shared"
  }
