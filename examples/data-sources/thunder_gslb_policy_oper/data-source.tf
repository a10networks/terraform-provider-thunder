provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_policy_oper" "thunder_gslb_policy_oper" {

  name = "test"

}
output "get_gslb_policy_oper" {
  value = ["${data.thunder_gslb_policy_oper.thunder_gslb_policy_oper}"]
}