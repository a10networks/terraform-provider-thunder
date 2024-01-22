provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_glm_enterprise_ha_host" "thunder_glm_enterprise_ha_host" {
  host_entry = "10.10.10.10"
}