provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_jwt_authorization" "thunder_aam_jwt_authorization" {
  exp_claim_requried = 0
  jwt_cache_enable   = 0
  jwt_exp_default    = 41378
  jwt_forwarding     = 0
  log_level          = "0"
  name               = "test"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "108"
}