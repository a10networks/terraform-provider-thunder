
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}

resource  "thunder_admin_password"   "admin_password"  {
  password_in_module = "a11"

}