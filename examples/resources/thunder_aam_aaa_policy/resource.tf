provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_aaa_policy" "thunder_aam_aaa_policy" {
  name = "test"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "54"
}