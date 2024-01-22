provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_common_cert_pinning_candidate_list_feedback_opt_in" "thunder_slb_common_cert_pinning_candidate_list_feedback_opt_in" {
  daily    = 1
  enable   = 1
  schedule = 1
  day_time = "12:45"
}