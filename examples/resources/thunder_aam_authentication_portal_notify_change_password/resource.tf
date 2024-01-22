provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_portal_notify_change_password" "thunder_aam_authentication_portal_notify_change_password" {

  name = "default-portal"
  background {
    bgfile  = "29"
    bgstyle = "tile"
  }
  cfm_pwd_cfg {
    confirm_password = 1
    cfm_text         = "24"
    cfm_font         = 1
    cfm_face         = "Arial"
    cfm_size         = 3
    cfm_color        = 1
    cfm_color_name   = "black"
  }
  change_text          = "14"
  change_url           = "6"
  confirm_password_var = "29"
  continue_text        = "44"
  continue_url         = "54"
  new_password_var     = "15"
  new_pwd_cfg {
    new_password   = 1
    new_text       = "47"
    new_font       = 1
    new_face       = "Arial"
    new_size       = 3
    new_color      = 1
    new_color_name = "black"
  }
  old_password_var = "1"
  old_pwd_cfg {
    old_password   = 1
    old_text       = "24"
    old_font       = 1
    old_face       = "Arial"
    old_size       = 3
    old_color      = 1
    old_color_name = "black"
  }
  username_cfg {
    username        = 1
    user_text       = "9"
    user_font       = 1
    user_face       = "Arial"
    user_size       = 3
    user_color      = 1
    user_color_name = "black"
  }
  username_var = "56"

}