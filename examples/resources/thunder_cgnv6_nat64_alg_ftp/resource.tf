provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat64_alg_ftp" "thunder_cgnv6_nat64_alg_ftp" {
  ftp_enable         = "disable"
  trans_eprt_to_port = "disable"
  trans_epsv_to_pasv = "disable"
  trans_lprt_to_port = "disable"
  trans_lpsv_to_pasv = "disable"
  xlat_no_trans_pasv = "enable"
}