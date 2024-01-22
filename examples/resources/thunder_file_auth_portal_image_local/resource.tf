provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_auth_portal_image_local" "thunder_file_auth_portal_image_local" {
  action      = "import"
  dst_file    = "test"
  file        = "test"
  file_handle = "/mnt/c/Users/PNimbhore/WorkSpace Terraform/sample-certificates/ECDSA"
}