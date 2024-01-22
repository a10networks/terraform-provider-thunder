provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cloud_services_cloud_provider_azure_metrics" "thunder_cloud_services_cloud_provider_azure_metrics" {
  action                 = "disable"
  active_partitions      = "86"
  cps                    = "disable"
  cpu                    = "disable"
  disk                   = "disable"
  interfaces             = "disable"
  memory                 = "disable"
  packet_drop            = "disable"
  packet_rate            = "disable"
  resource_id            = "252"
  server_down_count      = "disable"
  server_down_percentage = "disable"
  server_error           = "disable"
  sessions               = "disable"
  ssl_cert               = "disable"
  throughput             = "disable"
  tps                    = "disable"
}