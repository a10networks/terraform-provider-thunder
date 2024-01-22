provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_policy_class_list_lid" "thunder_slb_template_policy_class_list_lid" {

  name     = "test-policy"
  lidnum   = 515
  user_tag = "68"
}