provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_netflow_monitor_record" "thunder_netflow_monitor_record" {

  name                 = "a11"
  dslite               = 1
  nat44                = 1
  nat64                = 1
  netflow_v5           = 1
  netflow_v5_ext       = 1
  port_batch_dslite    = "both"
  port_batch_nat44     = "both"
  port_batch_nat64     = "both"
  port_batch_v2_dslite = "both"
  port_batch_v2_nat44  = "both"
  port_batch_v2_nat64  = "both"
  port_mapping_dslite  = "both"
  port_mapping_nat44   = "both"
  port_mapping_nat64   = "both"
  sesn_event_dslite    = "both"
  sesn_event_fw4       = "both"
  sesn_event_fw6       = "both"
  sesn_event_nat44     = "both"
  sesn_event_nat64     = "both"
}