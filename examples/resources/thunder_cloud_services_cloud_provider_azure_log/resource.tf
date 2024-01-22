provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cloud_services_cloud_provider_azure_log" "thunder_cloud_services_cloud_provider_azure_log" {
  action            = "disable"
  active_partitions = "195"
  resource_id       = "181"
  workspace_id      = "4"
}