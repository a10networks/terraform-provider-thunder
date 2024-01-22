provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_backup_periodic" "thunder_backup_periodic" {
  day           = 5
  encrypt       = 0
  fixed_nat     = 0
  log           = 0
  system        = 1
  use_mgmt_port = 1
  remote_file   = "scp://suraj:a10@10.64.3.184/home/suraj/testing.txt"
}