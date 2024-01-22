provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_sctp" "thunder_template_sctp" {
  name           = "test"
  checksum_check = "enable"
  log {
    payload_proto_filtering = 1
  }
  permit_payload_protocol {
    permit_config_id {
      protocol_id = 34
    }
    permit_config_name {
      protocol_name = "iua"
    }
  }
  sctp_half_open_idle_timeout = 4
  sctp_idle_timeout           = 5
  user_tag                    = "8"
}