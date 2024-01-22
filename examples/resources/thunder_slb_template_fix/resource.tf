provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_fix" "test_thunder_slb_template_fix" {
  name             = "test_fix"
  logging          = "init"
  insert_client_ip = 1
  tag_switching {
    switching_type = "sender-comp-id"
    equals         = "test"
    service_group  = "sg1"
  }
  user_tag = "test"
}