provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_cpu_packet_prio_support" "thunder_system_cpu_packet_prio_support" {
  disable = 0
}