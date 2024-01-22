provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_monitor" "thunder_visibility_monitor" {
  netflow {
    listening_port          = 80
    template_active_timeout = 31
  }
  primary_monitor    = "xflow"
  monitor_key        = "dest"
  mon_entity_topk    = 1
  source_entity_topk = 1
  secondary_monitor {
    secondary_monitoring_key = "service"
    mon_entity_topk          = 1
    source_entity_topk       = 1
  }
  sflow {
    listening_port = 6343
  }
}