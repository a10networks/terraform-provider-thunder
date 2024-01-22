provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_captcha_instance" "thunder_aam_authentication_captcha_instance" {

  method                     = "POST"
  name                       = "7"
  resp_error_code_field_name = "18"
  resp_result_field_name     = "30"
  secret_key                 = 0
  secret_key_param_name      = "11"
  timeout                    = 10
  token_param_name           = "7"
  url                        = "24"
}