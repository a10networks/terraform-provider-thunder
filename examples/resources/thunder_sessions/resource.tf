provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sessions" "thunder_sessions" {
  ext {
  }
  smp {
  }
  smp_table {
  }
}