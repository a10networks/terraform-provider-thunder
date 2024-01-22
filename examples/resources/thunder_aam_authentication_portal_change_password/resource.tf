provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_portal_change_password" "thunder_aam_authentication_portal_change_password" {

  name       = "default-portal"
  action_url = "56"
  background {
    bgfile  = "43"
    bgstyle = "tile"
  }
  cfm_pwd_cfg {
    confirm_password = 1
    cfm_text         = "60"
    cfm_font         = 1
    cfm_face         = "Arial"
    cfm_size         = 3
    cfm_color        = 1
    cfm_color_name   = "black"
  }
  confirm_password_var = "40"
  username_var         = "61"

}