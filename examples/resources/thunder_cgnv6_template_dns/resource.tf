provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_template_dns" "thunder_cgnv6_template_dns" {
  name           = "test1"
  default_policy = "cache"
  drop           = 0
  max_cache_size = 12
  period         = 2763
  user_tag       = "88"
}