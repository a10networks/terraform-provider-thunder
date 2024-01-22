provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_captcha" "thunder_aam_authentication_captcha" {

  instance_list {
    name                       = "56"
    secret_key                 = 0
    url                        = "59"
    method                     = "POST"
    timeout                    = 10
    secret_key_param_name      = "60"
    token_param_name           = "9"
    resp_result_field_name     = "15"
    resp_error_code_field_name = "24"
  }

}