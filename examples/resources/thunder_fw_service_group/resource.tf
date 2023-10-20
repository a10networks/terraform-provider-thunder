provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_service_group" "test_thunder_fw_service_group" {
  name = "sg1"
  protocol = "tcp"
  health_check = "ping"
  traffic_replication_mirror_ip_repl = 1
  user_tag = "test"
  sampling_enable{
      counters1 = "all"
    }
}