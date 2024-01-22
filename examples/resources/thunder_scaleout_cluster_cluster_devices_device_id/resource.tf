provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_cluster_devices_device_id" "thunder_scaleout_cluster_cluster_devices_device_id" {

  action     = "enable"
  device_id  = 12
  cluster_id = 2
}