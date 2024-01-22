provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_portal_logon" "thunder_aam_authentication_portal_logon" {

  name       = "default-portal"
  action_url = "16"
  background {
    bgfile  = "4"
    bgstyle = "tile"
  }
  username_var = "8"

}