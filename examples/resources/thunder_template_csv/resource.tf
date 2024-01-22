provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_csv" "thunder_template_csv" {
  csv_name    = "test"
  delim_char  = ","
  ipv6_enable = 1
  multiple_fields {
    field    = 17
    csv_type = "ip-from"
  }
  user_tag = "62"
}