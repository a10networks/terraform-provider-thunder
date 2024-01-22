provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_common_aflex_table_entry_sync" "thunder_slb_common_aflex_table_entry_sync" {
  aflex_table_entry_sync_enable        = 1
  aflex_table_entry_sync_max_key_len   = 756
  aflex_table_entry_sync_max_value_len = 565
  aflex_table_entry_sync_min_lifetime  = 1
}