provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_hm_dplane_stats" "thunder_slb_hm_dplane_stats" {

}
output "get_slb_hm_dplane_stats" {
  value = ["${data.thunder_slb_hm_dplane_stats.thunder_slb_hm_dplane_stats}"]
}