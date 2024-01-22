provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_sip" "thunder_slb_template_sip" {
  name                    = "sip-test"
  alg_dest_nat            = 1
  alg_source_nat          = 1
  call_id_persist_disable = 1
  client_keep_alive       = 1
  client_request_header {
    client_request_header_erase = "34"
    client_request_erase_all    = 0
  }
  client_response_header {
    client_response_header_erase = "33"
    client_response_erase_all    = 0
  }
  dialog_aware          = 1
  drop_when_client_fail = 1
  drop_when_server_fail = 1
  exclude_translation {
    translation_value = "body"
  }
  failed_client_selection     = 1
  failed_server_selection     = 1
  insert_client_ip            = 0
  interval                    = 30
  keep_server_ip_if_match_acl = 0
  pstn_gw                     = "pstn"
  server_keep_alive           = 0
  server_request_header {
    server_request_header_erase = "37"
    server_request_erase_all    = 0
  }
  server_response_header {
    server_response_header_erase = "34"
    server_response_erase_all    = 0
  }
  server_selection_per_request = 1
  smp_call_id_rtp_session      = 0
  timeout                      = 30
  user_tag                     = "68"
}