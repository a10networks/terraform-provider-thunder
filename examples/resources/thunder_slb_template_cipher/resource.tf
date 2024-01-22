provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_cipher" "thunder_slb_template_cipher" {
  name = "temp1"
  cipher_cfg {
    cipher_suite = "SSL3_RSA_DES_192_CBC3_SHA"
    priority     = 1
  }
  user_tag = "79"
}