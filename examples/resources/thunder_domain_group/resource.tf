provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_domain_group" "thunder_domain_group" {
  name = "test_domain_group"
  domain_rate_list {
    dummy_name = "configuration"
    user_tag   = "test"
    domain_list_list {
      name            = "test"
      per_suffix_rate = 3
      user_tag        = "test"
    }
  }
  user_tag = "test"

}