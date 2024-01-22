provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_based_policy_policy_class_list" "thunder_ddos_src_based_policy_policy_class_list" {
  name            = "test"
  class_list_name = "10"
}