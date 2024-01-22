provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_glm_proxy_server" "thunder_glm_proxy_server" {
  password = 0
  port     = 10091
  username = "root"
}