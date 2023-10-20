provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_cloud_services_cloud_provider_vmware_metrics" "test" {
  action= "enable"
  vrops_host= "enable"
  vrops_username= "enable"
  vrops_password= "enable"
  resource_id= "enable"
  adaptor_key= "enable"
  cpu= "enable"
  memory= "enable"
  disk= "enable"
  throughput= "enable"
  interfaces= "enable"
  cps= "enable"
  tps= "enable"
  server_down_count= "enable"
  server_down_percentage= "enable"
  ssl_cert= "enable"
  server_error= "enable"
  sessions= "enable"
  packet_drop= "enable"
  packet_rate= "enable"
}