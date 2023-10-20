provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_sip" "sips" {
  name                    = "sip1"
  alg_source_nat          = 1
  timeout                 = 1
  alg_dest_nat            = 1
  call_id_persist_disable = 1
  client_keep_alive       = 1
  client_request_header {
    client_request_header_erase = "siptest.com"
    client_request_erase_all    = 1
  }
  client_request_header {
    client_request_header_insert    = "siptest:insert"
    insert_condition_client_request = "insert-if-not-exist"
  }
  client_response_header {
    client_response_header_insert    = "siptestresp:insert"
    insert_condition_client_response = "insert-if-not-exist"
  }
  dialog_aware = 1
  exclude_translation {
    translation_value = "start-line"
  }
  exclude_translation {
    translation_value = "header"
    header_string     = "sipheader"
  }
  exclude_translation {
    translation_value = "body"
  }
  failed_client_selection     = 1
  drop_when_client_fail       = 1
  failed_server_selection     = 1
  drop_when_server_fail       = 1
  insert_client_ip            = 1
  keep_server_ip_if_match_acl = 1
  acl_name_value              = "acltest"
  pstn_gw                     = "pstn1"
  server_keep_alive           = 1
  interval                    = 250
  server_request_header {
    server_request_header_erase = "siptest.com"
    server_request_erase_all    = 1
  }
  server_request_header {
    server_request_header_insert    = "siptest1:insert"
    insert_condition_server_request = "insert-if-not-exist"
  }
  server_response_header {
    server_response_header_erase = "siptest2.com"
    server_response_erase_all    = 1
  }
  server_response_header {
    server_response_header_insert    = "testing:123"
    insert_condition_server_response = "insert-if-not-exist"
  }
  smp_call_id_rtp_session      = 1
  server_selection_per_request = 1
  user_tag                     = "testingsip"
}