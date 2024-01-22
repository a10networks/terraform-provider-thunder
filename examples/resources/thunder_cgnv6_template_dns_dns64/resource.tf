provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_template_dns_dns64" "thunder_cgnv6_template_dns_dns64" {

  name                    = "test1"
  answer_only_disable     = 1
  auth_data               = 1
  cache                   = 1
  change_query            = 1
  compress_disable        = 1
  deep_check_qr           = 1
  deep_check_rr_disable   = 1
  drop_cname_disable      = 1
  edns_append             = 1
  enable                  = 1
  fast_append             = 1
  ignore_rcode3_disable   = 1
  max_qr_length           = 128
  parallel_query          = 1
  passive_query_disable   = 1
  retry                   = 3
  single_response_disable = 1
  timeout                 = 1
  trans_ptr               = 1
  trans_ptr_query         = 1
  ttl                     = 780355973
}