provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_dns" "thunder_health_monitor_method_dns" {

  name       = "tf_test"
  dns        = 1
  dns_domain = "216"
  dns_domain_expect {
    dns_domain_response = "11"
  }
  dns_domain_port    = 34623
  dns_domain_recurse = "enabled"
  dns_domain_tcp     = 0
  dns_domain_type    = "A"
}