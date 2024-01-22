provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_polling" "thunder_sflow_polling" {
  a10_proprietary {
    export_deprecated_counters = 0
  }
  ddos {
    toggle                  = "enable"
    compatibility3_0        = 0
    address_byte_order_host = 0
    compatibility2_9        = 0
    dns_cache_zone_stats    = 0
    enable_anomaly_stats    = 0
    dyn_entry_stats         = 0
  }
  ethernet_list {
    start = 2
  }
  http {
    toggle = "enable"
  }
}