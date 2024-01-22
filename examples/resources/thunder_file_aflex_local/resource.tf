provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_aflex_local" "thunder_file_aflex_local" {
  action      = "import"
  dst_file    = "test"
  file        = "test"
  file_handle = "/mnt/c/Users/PNimbhore/WorkSpace Terraform/sample-certificates/class-list-ac1-a10"
  skip_backup = 0
}