provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_token_authentication_secret_salt" "thunder_ddos_token_authentication_secret_salt" {

  current_salt  = 4191031059
  previous_salt = 176240879

}