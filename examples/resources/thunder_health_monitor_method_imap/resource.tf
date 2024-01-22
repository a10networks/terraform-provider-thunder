provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_imap" "thunder_health_monitor_method_imap" {

  name                 = "tf_test"
  imap                 = 1
  imap_cram_md5        = 1
  imap_login           = 1
  imap_password        = 1
  imap_password_string = "a10net"
  imap_plain           = 1
  imap_port            = 14321
  imap_username        = "a11"
  pwd_auth             = 1
}