provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dns_query_type_filter" "thunder_slb_template_dns_query_type_filter" {

  name = "test"
  query_type {
    str_query_type = "A"
    num_query_type = 60849
  }
  query_type_action = "allow"
}