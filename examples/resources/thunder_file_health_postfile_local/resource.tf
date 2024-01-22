provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_health_postfile_local" "thunder_file_health_postfile_local" {
  action      = "create"
  dst_file    = "test"
  file        = "test"
  file_handle = "/mnt/c/Users/PNimbhore/WorkSpace Terraform/sample-certificates/glm.txt"
}