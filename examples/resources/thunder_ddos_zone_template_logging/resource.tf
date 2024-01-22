provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_logging" "thunder_ddos_zone_template_logging" {

  enable_action_logging = 1
  log_format_cef        = 1
  log_format_custom     = "376"
  logging_tmpl_name     = "32"
  use_obj_name          = 1
  user_tag              = "82"

}