provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_domain_group_domain_rate" "thunder_domain_group_domain_rate" {
  name       = "test_domain_group"
  dummy_name = "configuration"
  user_tag   = "test"
  domain_list_list {
    name            = "test"
    per_suffix_rate = 3
    user_tag        = "test"
  }

}