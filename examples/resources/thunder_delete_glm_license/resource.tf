provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_delete_glm_license" "thunder_delete_glm_license" {
  a10_ti        = 0
  ipsec_vpn     = 0
  ngwaf         = 0
  qosmos        = 0
  secure_gaming = 0
  threatstop    = 0
  webroot       = 0
  webroot_ti    = 0
}