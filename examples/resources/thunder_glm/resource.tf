provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_glm" "Glm_Test" {
  use_mgmt_port = 1
  burst         = 0
  interval      = 1
  token         = "a10incedo"
  enterprise    = "test"
  proxy_server {
    username      = "saleem123"
    host          = "192.168.20.1"
    password      = 1
    port          = 6666
    secret_string = "stringa"
  }
  appliance_name     = "pune"
  enable_requests    = 1
  allocate_bandwidth = 50
  port               = 5555
}