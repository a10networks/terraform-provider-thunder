provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_hc" "thunder_debug_hc" {
  anomaly        = 1
  app_svc_id     = "102"
  error          = 1
  metrics        = 1
  object_uuid    = "54"
  per_connection = 1
  per_request    = 1
  registration   = 1
  uri            = "13"
}