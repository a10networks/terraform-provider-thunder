provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_logon_form_based" "thunder_aam_authentication_logon_form_based" {
  challenge_variable = "1"
  cp_page_cfg {
    changepassword_url = "56"
    cp_user_enum       = "changepassword-username-variable"
    cp_user_var        = "49"
    cp_old_pwd_enum    = "changepassword-old-password-variable"
    cp_old_pwd_var     = "44"
    cp_new_pwd_enum    = "changepassword-new-password-variable"
    cp_new_pwd_var     = "35"
    cp_cfm_pwd_enum    = "changepassword-password-confirm-variable"
    cp_cfm_pwd_var     = "40"
  }
  csp_support {
    specificuri         = "37"
    optional_second_uri = "69"
  }
  duration     = 1800
  hsts_timeout = 298757643
  logon_page_cfg {
    action_url                   = "10"
    username_variable            = "44"
    password_variable            = "59"
    passcode_variable            = "60"
    captcha_variable             = "11"
    login_failure_message        = "11"
    authz_failure_message        = "68"
    disable_change_password_link = 0
  }
  name                = "28"
  new_pin_variable    = "41"
  next_token_variable = "1"
  notify_cp_page_cfg {
    notifychangepassword_change_url   = "59"
    notifychangepassword_continue_url = "25"
  }
  retry    = 3
  user_tag = "92"
}