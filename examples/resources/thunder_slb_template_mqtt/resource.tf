provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_mqtt" "test_thunder_slb_template_mqtt" {
  name                  = "test"
  clientid_hash_persist = 1
  clientid_hash_offset  = 23
  clientid_hash_first   = 6
  clientid_hash_last    = 5
  user_tag              = "testing"
}