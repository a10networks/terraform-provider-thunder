provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat64_alg_esp" "thunder_cgnv6_nat64_alg_esp" {
  esp_enable = "enable"
  sampling_enable {
    counters1 = "all"
  }
}