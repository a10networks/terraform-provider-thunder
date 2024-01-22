provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_dns64_virtualserver_oper" "thunder_cgnv6_dns64_virtualserver_oper" {
  oper {
    state          = "All Up"
    conn_rate_unit = "100ms"
    ip_address     = "10.10.10.10"
  }
  port_list {
    port_number = 32768
    protocol    = "dns-udp"
    oper {
      state = "All Up"
      http_hits_list {
      }
      http_vport_cpu_list {
      }
    }
  }

}
output "get_cgnv6_dns64_virtualserver_oper" {
  value = ["${data.thunder_cgnv6_dns64_virtualserver_oper.thunder_cgnv6_dns64_virtualserver_oper}"]
}