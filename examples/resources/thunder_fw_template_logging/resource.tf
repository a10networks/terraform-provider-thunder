provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_template_logging" "thunder_fw_template_logging" {
  name              = "test"
  resolution        = "seconds"
  include_dest_fqdn = 1
  merged_style      = 1
  log {
    http_requests = "host"
  }
  include_radius_attribute {
    framed_ipv6_prefix = 1
    prefix_length      = "32"
  }
  include_http {
    header_cfg {
      http_header = "cookie"
      max_length  = 101
    }
  }
  rule {
    rule_http_requests {
      dest_port {
        dest_port_number = 2
      }
    }
  }
  facility = "user"
  severity = "emergency"
  format   = "ascii"
  user_tag = "test"
  source_address {
    ip   = "10.10.10.10"
    ipv6 = "2001:db8:0:200::1"
  }
}