provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_mssql" "thunder_slb_mssql" {
  sampling_enable {
    counters1 = "all"
  }
}