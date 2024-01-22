provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_buffered" "thunder_logging_buffered" {
  buffersize           = 30000
  levelname            = "alert"
  partition_buffersize = 3000
}