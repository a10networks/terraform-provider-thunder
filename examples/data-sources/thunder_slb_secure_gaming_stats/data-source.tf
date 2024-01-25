provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_secure_gaming_stats" "thunder_slb_secure_gaming_stats" {
}
output "get_slb_secure_gaming_stats" {
  value = ["${data.thunder_slb_secure_gaming_stats.thunder_slb_secure_gaming_stats}"]
}