provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_dns" "thunder_system_dns" {
  recursive_nameserver {
    server_list {
      ipv4_addr = "10.10.10.10"
      v4_desc   = "23"
    }
  }
  sampling_enable {
    counters1 = "all"
  }
}