provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat64_alg_rtsp" "thunder_cgnv6_nat64_alg_rtsp" {
  rtsp_enable = "enable"
}