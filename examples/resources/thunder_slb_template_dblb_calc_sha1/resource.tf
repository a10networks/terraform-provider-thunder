provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dblb_calc_sha1" "thunder_slb_template_dblb_calc_sha1" {

  name       = "temp1"
  sha1_value = "a43"
}