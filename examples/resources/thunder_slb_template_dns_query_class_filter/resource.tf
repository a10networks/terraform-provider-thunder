provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dns_query_class_filter" "thunder_slb_template_dns_query_class_filter" {

  name = "test"
  query_class {
    str_query_class = "INTERNET"
    num_query_class = 34791
  }
  query_class_action = "allow"
}