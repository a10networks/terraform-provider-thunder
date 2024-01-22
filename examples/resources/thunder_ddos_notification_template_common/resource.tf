provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_notification_template_common" "thunder_ddos_notification_template_common" {

  default_template {
    default_notification_template = "test"
  }
  on_box_gui_notification = "enable"

}