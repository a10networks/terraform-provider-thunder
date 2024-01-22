provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_monitored_entity_detail_debug_oper" "thunder_visibility_monitored_entity_detail_debug_oper" {

}
output "get_visibility_monitored_entity_detail_debug_oper" {
  value = ["${data.thunder_visibility_monitored_entity_detail_debug_oper.thunder_visibility_monitored_entity_detail_debug_oper}"]
}