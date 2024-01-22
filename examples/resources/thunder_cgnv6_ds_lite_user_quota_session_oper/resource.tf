provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_ds_lite_user_quota_session_oper" "thunder_cgnv6_ds_lite_user_quota_session_oper" {

}
output "get_cgnv6_ds_lite_user_quota_session_oper" {
  value = ["${data.thunder_cgnv6_ds_lite_user_quota_session_oper.thunder_cgnv6_ds_lite_user_quota_session_oper}"]
}