provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_geo_location_local" "thunder_file_geo_location_local" {
  action      = "import"
  dst_file    = "test123"
  file        = "test123"
  file_handle = "/mnt/c/Users/PNimbhore/WorkSpace Terraform/sample-certificates/test2.csv"
  user_tag    = "tt"
}