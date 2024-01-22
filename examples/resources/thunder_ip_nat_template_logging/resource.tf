provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_template_logging" "thunder_ip_nat_template_logging" {
  name                = "a11"
  facility            = "local0"
  include_destination = 1
  include_rip_rport   = 1
  log {
    port_mappings = "creation"
  }
  severity {
    severity_string = "debug"
  }
  source_port {
    source_port_num = 524
  }
  user_tag = "119"
}