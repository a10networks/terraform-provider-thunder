provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_cloud_services_cloud_provider_azure_metrics" "test" {
  action =  "enable"
  client_id =  "enable"
  secret_id =  "enable"
  tenant_id =  "enable"
  resource_id =  "enable"
  location =  "enable"
  cpu =  "enable"
  memory =  "enable"
  disk =  "enable"
  throughput =  "enable"
  interfaces =  "enable"
  cps =  "enable"
  tps =  "enable"
  server_down_count =  "enable"
  server_down_percentage =  "enable"
  ssl_cert =  "enable"
  server_error =  "enable"
  sessions =  "enable"
  packet_drop =  "enable"
  packet_rate =  "enable"
}