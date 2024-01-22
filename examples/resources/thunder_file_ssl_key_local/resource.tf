provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_ssl_key_local" "thunder_file_ssl_key_local" {
  action      = "import"
  dst_file    = "test123"
  file_handle = "/mnt/c/Users/PNimbhore/WorkSpace Terraform/sample-certificates/test_crl.crl"
  secured     = 0
}