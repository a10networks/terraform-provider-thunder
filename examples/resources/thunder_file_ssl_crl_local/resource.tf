provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_ssl_crl_local" "thunder_file_ssl_crl_local" {
  action      = "import"
  file        = "test123"
  file_handle = "/mnt/c/Users/PNimbhore/WorkSpace Terraform/sample-certificates/test_crl.crl"
}