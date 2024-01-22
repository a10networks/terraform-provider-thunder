provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_startup_config_local" "thunder_file_startup_config_local" {
  action      = "import"
  file        = "test123"
  file_handle = "/mnt/c/Users/PNimbhore/WorkSpace Terraform/sample-certificates/startup.txt"
}