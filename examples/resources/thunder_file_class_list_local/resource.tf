provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_class_list_local" "thunder_file_class_list_local" {
  action      = "import"
  dst_file    = "test132"
  file        = "test123"
  file_handle = "/mnt/c/Users/PNimbhore/WorkSpace Terraform/sample-certificates/class-list-ac1-a10"
  user_tag    = "27"
}