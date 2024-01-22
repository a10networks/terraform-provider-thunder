provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_notification_template_api_authentication" "thunder_ddos_notification_template_api_authentication" {
  name           = "test"
  api_key        = 1
  api_key_string = "14"

}