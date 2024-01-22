provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_ds_lite_alg_ftp" "thunder_cgnv6_ds_lite_alg_ftp" {
  ftp_enable = "disable"
}