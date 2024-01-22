provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dns_recursive_dns_resolution_lookup_order" "thunder_slb_template_dns_recursive_dns_resolution_lookup_order" {

  name = "test"
  query_type {
    num_query_type = 3046
    order          = "ipv4-precede-ipv6"
  }
}