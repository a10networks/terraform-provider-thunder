provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_template_logging" "thunder_cgnv6_template_logging" {
  name                    = "test"
  batched_logging_disable = 1
  custom {
    custom_header = "use-syslog-header"

    custom_time_stamp_format = "19"
  }
  facility            = "local0"
  format              = "default"
  include_destination = 1
  include_http {
    header_cfg {
      http_header       = "cookie"
      max_length        = 100
      custom_max_length = 100
    }
    l4_session_info = 1
    method          = 1
    request_number  = 1
    file_extension  = 1
  }
  include_inside_user_mac    = 1
  include_partition_name     = 1
  include_port_block_account = 1
  include_radius_attribute {
    attr_cfg {
      attr       = "imei"
      attr_event = "http-requests"
    }
    no_quote               = 0
    insert_if_not_existing = 0
    zero_in_custom_attr    = 0
  }
  include_session_byte_count = 0
  resolution                 = "seconds"


  user_tag = "83"
}