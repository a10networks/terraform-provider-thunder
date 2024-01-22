provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_csr" "FileCsr" {
  name         = "FILECSR1"
  cert_type    = "rsa"
  common_name  = "cert"
  country      = "USA"
  digest       = "sha1"
  email        = "user@a10networks.com"
  key_size     = "1024"
  organization = "A10"
  valid_days   = 100
}