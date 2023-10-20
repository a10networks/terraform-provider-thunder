provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_router_ospf_default_information" "RouterOspfDefaultInformation" {
  originate   = 1
  always      = 1
  metric      = 40
  metric_type = 1
  route_map   = "MAP"
  process_id  = 400
}
