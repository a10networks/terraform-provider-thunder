provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_server_ssh" "Slb_Template_Server_Ssh_Test" {
  forward_proxy_enable = 1
  name                 = "SSH"
  user_tag             = "ServerSSH"
}