provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_policy_class_list" "thunder_slb_template_policy_class_list" {

  name = "test-policy"
}