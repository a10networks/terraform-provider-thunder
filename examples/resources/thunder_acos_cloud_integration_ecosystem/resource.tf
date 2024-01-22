provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_cloud_integration_ecosystem" "thunder_acos_cloud_integration_ecosystem" {
  consul {
    host_name = "15"
    port      = 43561
    action    = "disable"
  }
  dummy = 0
  k8s {
    action = "disable"
  }
  oracle {
    host_name        = "2"
    port             = 47037
    compartment_id   = "11"
    tenancy_id       = "59"
    user_id          = "19"
    fingerprint      = "40"
    private_key_path = "21"
    action           = "disable"
  }
}