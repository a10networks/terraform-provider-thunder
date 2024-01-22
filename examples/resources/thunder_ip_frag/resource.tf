provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_frag" "thunder_ip_frag" {
  buff                       = 200000
  max_packets_per_reassembly = 10
  max_reassembly_sessions    = 1
  sampling_enable {
    counters1 = "all"
  }
  cpu_threshold {
    high = 100
    low  = 50
  }
}