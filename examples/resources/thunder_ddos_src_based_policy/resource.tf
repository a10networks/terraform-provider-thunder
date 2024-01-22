provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_based_policy" "thunder_ddos_src_based_policy" {
  name = "test"
  policy_class_list_list {
    class_list_name = "10"
  }
  user_tag = "106"
}