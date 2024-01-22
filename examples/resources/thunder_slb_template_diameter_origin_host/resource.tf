provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_diameter_origin_host" "thunder_slb_template_diameter_origin_host" {

  name             = "temp1"
  origin_host_name = "159"
}