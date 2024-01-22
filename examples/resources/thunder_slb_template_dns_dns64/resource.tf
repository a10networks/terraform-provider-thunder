provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dns_dns64" "thunder_slb_template_dns_dns64" {

  name                    = "test"
  cache                   = 1
  change_query            = 1
  enable                  = 1
  parallel_query          = 0
  retry                   = 3
  single_response_disable = 1
  timeout                 = 1
}