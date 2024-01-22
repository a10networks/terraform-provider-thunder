provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_ca_cert" "thunder_file_ca_cert" {
  name             = "CACERT"
  protocol         = "http"
  host             = "10.64.3.190"
  path             = "/test123.pem"
  use_mgmt_port    = 1
  certificate_type = "pem"
}