provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_mirror_port" "thunder_mirror_port" {
  ethernet     = 2
  mirror_dir   = "both"
  mirror_index = 2
}