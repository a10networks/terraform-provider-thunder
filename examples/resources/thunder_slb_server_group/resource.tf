provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_server_group" "thunder_slb_server_group" {

  name = "server123"
  member_list {
    name = "test-server1"
  }
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "41"
}