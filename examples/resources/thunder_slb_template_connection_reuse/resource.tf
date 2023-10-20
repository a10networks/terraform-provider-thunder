provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_connection_reuse" "conn-reuse" {
  name              = "conn_reuse"
  limit_per_server  = 65535
  timeout           = 3600
  keep_alive_conn   = 1
  preopen           = 1
  num_conn_per_port = 1
  user_tag          = "connreuse"
} 