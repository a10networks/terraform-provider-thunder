provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_portal_logon_fail" "thunder_aam_authentication_portal_logon_fail" {

  name = "default-portal"
  background {
    bgfile  = "53"
    bgstyle = "tile"
  }
  fail_msg_cfg {
    fail_msg        = 1
    fail_text       = "36"
    fail_font       = 1
    fail_face       = "Arial"
    fail_size       = 3
    fail_color      = 1
    fail_color_name = "black"
  }
  title_cfg {
    title            = 1
    title_text       = "53"
    title_font       = 1
    title_face       = "Arial"
    title_size       = 5
    title_color      = 1
    title_color_name = "black"
  }

}