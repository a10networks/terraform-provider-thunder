provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_persist_cookie" "persist-cookie" {
  name                  = "persistcookie1"
  cookie_name           = "testcookie123"
  dont_honor_conn_rules = 1
  expire                = 31536
  insert_always         = 1
  samesite              = "none"
  domain                = "test.com"
  path                  = "/tmp/cookie"
  pass_thru             = 1
  secure                = 1
  httponly              = 1
  match_type            = 1
  service_group         = 1
  encrypt_level         = 0
  pass_phrase           = "password"
  user_tag              = "persistcookie1"
}