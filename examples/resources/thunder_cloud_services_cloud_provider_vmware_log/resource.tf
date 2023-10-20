provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_cloud_services_cloud_provider_vmware_log" "test" {
  action = "enable"
  vrli_host = "test"
}