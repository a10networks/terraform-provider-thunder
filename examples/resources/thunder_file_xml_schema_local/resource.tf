provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_xml_schema_local" "thunder_file_xml_schema_local" {
  action      = "create"
  dst_file    = "test123"
  file        = "tezt"
  file_handle = "/mnt/c/Users/PNimbhore/WorkSpace Terraform/sample-certificates/syslog.txt"
}