provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_partition_shared_vlan" "thunder_partition_shared_vlan" {

  partition_name = "test"
  vlan           = 3

}