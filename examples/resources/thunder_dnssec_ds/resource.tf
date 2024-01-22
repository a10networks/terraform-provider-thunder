provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_dnssec_ds" "thunder_dnssec_ds" {
  ds_delete = 1
  zone_name = "test"
}