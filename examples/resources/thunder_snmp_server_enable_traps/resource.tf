provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_enable_traps" "Snmp_Server_Enable_Traps_Test" {
  lldp = 1
  all  = 1
  slb_change {
    all                    = 1
    resource_usage_warning = 1
    ssl_cert_change        = 1
  }
  lsn {
    all                                = 1
    fixed_nat_port_mapping_file_change = 1
    per_ip_port_usage_threshold        = 1
    total_port_usage_threshold         = 1
  }
  vrrp_a {
    active  = 1
    standby = 1
    all     = 1
  }
  snmp {
    linkup   = 1
    all      = 1
    linkdown = 1
  }
  system {
    all           = 1
    data_cpu_high = 1
    power         = 1
    high_disk_use = 1
  }
  ssl {
    server_certificate_error = 1
  }
  vcs {
    state_change = 1
  }
  routing {
    bgp {
      bgpestablishednotification = 1

      bgpbackwardtransnotification = 1
    }
    isis {
      isisauthenticationfailure      = 1
      isisprotocolssupportedmismatch = 1
      isisrejectedadjacency          = 1
      isismaxareaaddressesmismatch   = 1
    }
    ospf {
      ospflsdboverflow            = 1
      uuid                        = 1
      ospfnbrstatechange          = 1
      ospfifstatechange           = 1
      ospfvirtnbrstatechange      = 1
      ospflsdbapproachingoverflow = 1
      ospfifauthfailure           = 1

    }
  }
  gslb {
    all        = 1
    group      = 1
    zone       = 1
    site       = 1
    service_ip = 1
  }
  slb {
    all                      = 1
    server_down              = 1
    vip_port_connratelimit   = 1
    server_selection_failure = 1
    service_group_down       = 1
    server_conn_limit        = 1
  }
  network {
    trunk_port_threshold = 1
  }
}