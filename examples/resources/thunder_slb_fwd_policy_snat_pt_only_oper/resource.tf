provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_fwd_policy_snat_pt_only_oper" "thunder_slb_fwd_policy_snat_pt_only_oper" {

}
output "get_slb_fwd_policy_snat_pt_only_oper" {
  value = ["${data.thunder_slb_fwd_policy_snat_pt_only_oper.thunder_slb_fwd_policy_snat_pt_only_oper}"]
}