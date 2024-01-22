provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_client_ssh" "Slb_Template_Client_Ssh_Test" {
  user_tag              = "ClientSSH"
  forward_proxy_enable  = 1
  passphrase            = "clientkey"
  forward_proxy_hostkey = "SLBCERT"
  name                  = "CLientSSH"
}