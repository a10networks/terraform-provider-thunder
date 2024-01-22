provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cloud_services_meta_data" "thunder_cloud_services_meta_data" {
  action                = "enable"
  prevent_admin_passwd  = 1
  prevent_admin_ssh_key = 1
  prevent_autofill      = 1
  prevent_blob          = 1
  prevent_cloud_service = 1
  prevent_license       = 1
  prevent_user_ops      = 1
  prevent_webservice    = 1
}