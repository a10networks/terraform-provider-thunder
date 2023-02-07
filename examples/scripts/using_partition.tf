provider "thunder" {
  address  = "${var.address}"
  username = "${var.username}"
  password = "${var.password}"
  partition = "${var.partition}"
}

resource "thunder_server" "test" {   
  name   = "real-server"
  host   = "10.10.10.1"
  port_list {
    port_number = 80
    protocol    = "tcp"
  }
}