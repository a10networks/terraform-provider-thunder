provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_cloud_integration_ecosystem_oracle" "thunder_acos_cloud_integration_ecosystem_oracle" {
  action                = "disable"
  compartment_id        = "90"
  fingerprint           = "82"
  health_check_interval = "5"
  host_name             = "75"
  port                  = 20091
  private_key_path      = "15"
  tenancy_id            = "36"
  user_id               = "72"
}