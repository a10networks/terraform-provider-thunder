provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource  "thunder_gslb_template_csv"   "test_thunder_gslb_template_csv"  {
    csv_name = "test"
    delim_char = "d"
    ipv6_enable = 1
    user_tag = "usetTest"
  }