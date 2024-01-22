provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_ssl_cert_local" "thunder_file_ssl_cert_local" {
  action              = "import"
  certificate_type    = "pem"
  dst_file            = "test123"
  file                = "test123"
  file_handle         = "/mnt/c/Users/PNimbhore/WorkSpace Terraform/sample-certificates/test123.pem"
  pfx_password        = "47"
  pfx_password_export = "82"
}