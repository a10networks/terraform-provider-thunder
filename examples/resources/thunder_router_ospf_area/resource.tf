provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_router_ospf_area" "resourceRouterOspfAreaTest" {
  process_id   = 400
  area_ipv4    = "0.0.0.2"
  default_cost = 200
  shortcut     = "enable"
  auth_cfg {
    authentication = 1
    message_digest = 1
  }
  stub_cfg {
    stub       = 0
    no_summary = 0
  }
  filter_lists {
    filter_list     = 1
    acl_name        = "ACL1"
    acl_direction   = "in"
    plist_name      = "PREFIX1"
    plist_direction = "in"
  }
  range_list {
    area_range_prefix = "1.1.1.1/24"
    option            = "advertise"
  }
  virtual_link_list {
    virtual_link_ip_addr        = "3.3.3.3"
    bfd                         = 1
    hello_interval              = 10
    dead_interval               = 40
    retransmit_interval         = 300
    transmit_delay              = 50
    virtual_link_authentication = 1
    virtual_link_auth_type      = "message-digest"
    authentication_key          = "string"
    message_digest_key          = 20
    md5                         = "MD5"
  }
}

