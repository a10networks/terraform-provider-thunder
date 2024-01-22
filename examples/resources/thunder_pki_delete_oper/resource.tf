provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_pki_delete_oper" "thunder_pki_delete_oper" {
}
output "get_pki_delete_oper" {
  value = ["${data.thunder_pki_delete_oper.thunder_pki_delete_oper}"]
}