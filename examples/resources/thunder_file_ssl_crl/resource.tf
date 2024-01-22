provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_ssl_crl" "FileSslCrl" {
  name          = "FILE_SSL_CRL"
  protocol      = "scp"
  host          = "10.64.3.186"
  path          = "/home/server1/test_crl.crl"
  password      = "a10"
  username      = "server1"
  overwrite     = 1
  use_mgmt_port = 1
}