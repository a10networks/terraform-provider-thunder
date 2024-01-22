provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_monitor" "thunder_monitor" {
  buffer_drop  = 4000
  buffer_usage = 3081497
  conn_type0   = 32767
  conn_type1   = 32767
  conn_type2   = 32767
  conn_type3   = 32767
  conn_type4   = 32767
  ctrl_cpu     = 90
  data_cpu     = 90
  disk         = 85
  memory       = 95
  smp_type0    = 32767
  smp_type1    = 32767
  smp_type2    = 32767
  smp_type3    = 32767
  smp_type4    = 32767
  warn_temp    = 61
}