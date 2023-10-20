provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_server" "servers" {
  name               = "web1"
  host               = "10.10.10.1"
  external_ip        = "21.21.21.1"
  use_aam_server     = 0
  ipv6               = "2002::10.10.10.1"
  template_server    = "server_template"
  template_link_cost = "link_template"
  health_check       = "tcp1"
  no_logging         = 0
  conn_limit         = 1
  conn_resume        = 1
  weight             = "301"
  slow_start         = 1
  spoofing_cache     = 1
  stats_data_action  = "stats-data-disable"
  extended_stats     = 1
  alternate_server {
    alternate      = 2
    alternate_name = "web2"
  }
  user_tag = "slb_server"
  sampling_enable {
    counters1 = "total-conn"
  }
  port_list {
    port_number       = 80
    protocol          = "tcp"
    range             = 0
    template_port     = "template_port"
    action            = "enable"
    no_ssl            = 1
    support_http2     = 1
    weight            = 1
    conn_limit        = 1
    no_logging        = 0
    conn_resume       = 1
    stats_data_action = "stats-data-disable"
    extended_stats    = 1
    auth_cfg {
      service_principal_name = "kerb1"
    }
    user_tag = "port"

    sampling_enable {
      counters1 = "curr_req"
    }
  }
  port_list {
    port_number       = 53
    protocol          = "udp"
    range             = 0
    template_port     = "template_port"
    action            = "disable-with-health-check"
    no_ssl            = 1
    health_check      = "tcp1"
    support_http2     = 1
    weight            = 1
    conn_limit        = 1
    no_logging        = 0
    conn_resume       = 1
    stats_data_action = "stats-data-disable"
    extended_stats    = 1
    auth_cfg {
      service_principal_name = "kerb1"
    }
    user_tag = "port"
    sampling_enable {
      counters1 = "curr_req"
    }

  }
}