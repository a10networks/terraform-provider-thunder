provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_monitor_secondary_monitor" "thunder_visibility_monitor_secondary_monitor" {
  secondary_monitoring_key = "service"
  mon_entity_topk          = 1
  source_entity_topk       = 1

}