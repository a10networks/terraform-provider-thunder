provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_web_category_proxy_server" "resourceWebCategoryProxyServerTest" {
  proxy_host    = "192.168.50.10"
  http_port     = 80
  https_port    = 443
  auth_type     = "ntlm"
  domain        = "NTLMdomain"
  username      = "testing"
  password      = 1
  secret_string = "testSecret"
}