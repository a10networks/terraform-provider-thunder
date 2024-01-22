provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_spanning_tree_mode_mstp" "thunder_network_spanning_tree_mode_mstp" {

  action = 1
  instance_list {
    instance_start = 2809
    priority       = 32768
    user_tag       = "48"
  }
  priority = 32768

}