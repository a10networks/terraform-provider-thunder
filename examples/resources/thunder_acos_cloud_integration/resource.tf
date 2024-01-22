provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_cloud_integration" "thunder_acos_cloud_integration" {
  dummy = 0
  ecosystem {
    dummy = 0
    consul {
      host_name = "125"
      port      = 62058
      action    = "disable"
    }
    oracle {
      host_name        = "18"
      port             = 20567
      compartment_id   = "73"
      tenancy_id       = "80"
      user_id          = "56"
      fingerprint      = "61"
      private_key_path = "96"
      action           = "disable"
    }
    k8s {
      action = "disable"
    }
  }
}