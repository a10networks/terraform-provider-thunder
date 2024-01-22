provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_topn_cgnv6_nat_pool_topn_tmpl_metrics" "thunder_visibility_topn_cgnv6_nat_pool_topn_tmpl_metrics" {

  tcp_total = 1
  name      = "test"
}