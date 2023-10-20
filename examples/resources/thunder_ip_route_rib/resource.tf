provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ip_route_rib" "rib_1" {
  ip_dest_addr = "0.0.0.0"
  ip_mask      = "/0"
  ip_nexthop_ipv4 {
    ip_next_hop         = "10.102.156.1"
    distance_nexthop_ip = 18
  }
}
resource "thunder_ip_route_rib" "rib_2" {
  ip_dest_addr = "3.4.0.0"
  ip_mask      = "/17"
  ip_nexthop_partition {
    partition_name                = "L3V_A"
    description_nexthop_partition = "timestamp"
  }
}
resource "thunder_ip_route_rib" "rib_3" {
  ip_dest_addr = "7.7.0.0"
  ip_mask      = "/16"
  ip_nexthop_lif {
    lif                     = "Lif1"
    description_nexthop_lif = "Lif_Descrition"
  }
}