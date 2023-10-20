provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_alg_dns" "test_thunder_fw_alg_dns" {
  default_port_disable = "default-port-disable"
}