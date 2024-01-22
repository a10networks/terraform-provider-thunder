provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_capture_config" "thunder_ddos_dst_entry_capture_config" {

  mode           = "drop"
  name           = "10"
  dst_entry_name = "test"

}