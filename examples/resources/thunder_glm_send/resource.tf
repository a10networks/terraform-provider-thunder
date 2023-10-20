
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_glm_send"  "GlmSend"  {
    license_request = 1
      
}