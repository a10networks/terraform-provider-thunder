provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_monitored_entity" "thunder_visibility_monitored_entity" {

}