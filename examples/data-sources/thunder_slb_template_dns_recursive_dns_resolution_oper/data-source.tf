provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_template_dns_recursive_dns_resolution_oper" "thunder_slb_template_dns_recursive_dns_resolution_oper" {

  name = "test"
}

output "get_slb_template_dns_recursive_dns_resolution_oper" {
  value = ["${data.thunder_slb_template_dns_recursive_dns_resolution_oper.thunder_slb_template_dns_recursive_dns_resolution_oper}"]
}