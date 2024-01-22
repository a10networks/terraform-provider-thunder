provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_polling_ddos" "thunder_sflow_polling_ddos" {
  address_byte_order_host = 0
  compatibility2_9        = 0
  compatibility3_0        = 0
  dns_cache_zone_stats    = 0
  dyn_entry_stats         = 0
  enable_anomaly_stats    = 0
  toggle                  = "enable"
}