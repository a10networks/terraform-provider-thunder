provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dns_logging_response_type_type" "thunder_slb_template_dns_logging_response_type_type" {

  name                  = "test"
  caa_type_limit_num    = 0
  digest                = 0
  dnskey_type_limit_num = 0
  ds_type_limit_num     = 0
  length_limit_flag     = 0
  naptr_type_limit_num  = 0
  opt_type_limit_num    = 0
  other_data            = 0
  public_key            = 0
  rdata_field           = 0
  response_type_name    = "TXT"
  rrsig_type_limit_num  = 0
  service_field         = 0
  signature             = 0
  tsig_type_limit_num   = 0
  txt_data              = 0
  txt_type_limit_num    = 0
  user_tag              = "108"
  value_field           = 0
}