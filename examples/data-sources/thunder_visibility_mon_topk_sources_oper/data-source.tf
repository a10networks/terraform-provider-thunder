provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_mon_topk_sources_oper" "thunder_visibility_mon_topk_sources_oper" {

}
output "get_visibility_mon_topk_sources_oper" {
  value = ["${data.thunder_visibility_mon_topk_sources_oper.thunder_visibility_mon_topk_sources_oper}"]
}