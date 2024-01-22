provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_local_log_app_fw" "thunder_logging_local_log_app_fw" {
  dot_plot {
  }
  top_n {
  }
}