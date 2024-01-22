provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_database" "thunder_health_monitor_method_database" {

  name            = "tf_test"
  database        = 1
  database_name   = "mssql"
  db_column       = 10
  db_name         = "test_db11"
  db_password     = 1
  db_password_str = "root"
  db_receive      = "3"
  db_row          = 9
  db_send         = "91"
  db_username     = "kp11"
}