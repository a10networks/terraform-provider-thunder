provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ipv6_frag" "testname" {
  sampling_enable {
    counters1 = "session-inserted"
  }
  frag_timeout = 16000
}