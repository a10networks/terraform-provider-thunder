provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_topn" "thunder_visibility_topn" {
  cgnv6_nat_pool_topn_tmpl_list {
    name     = "test"
    user_tag = "test"
    metrics {
      tcp_total = 1
    }
  }
  gtp_apn_prefix_topn_tmpl_list {
    name     = "test"
    user_tag = "117"
    metrics {
      uplink_bytes                        = 1
      downlink_bytes                      = 1
      uplink_pkts                         = 1
      downlink_pkts                       = 1
      gtp_v0_c_tunnel_created             = 1
      gtp_v0_c_tunnel_half_open           = 1
      gtp_v0_c_tunnel_half_closed         = 1
      gtp_v0_c_tunnel_closed              = 1
      gtp_v0_c_tunnel_deleted             = 1
      gtp_v0_c_half_open_tunnel_closed    = 1
      gtp_v1_c_tunnel_created             = 1
      gtp_v1_c_tunnel_half_open           = 1
      gtp_v1_c_tunnel_half_closed         = 1
      gtp_v1_c_tunnel_closed              = 1
      gtp_v1_c_tunnel_deleted             = 1
      gtp_v1_c_half_open_tunnel_closed    = 1
      gtp_v2_c_tunnel_created             = 1
      gtp_v2_c_tunnel_half_open           = 1
      gtp_v2_c_tunnel_half_closed         = 1
      gtp_v2_c_tunnel_closed              = 1
      gtp_v2_c_tunnel_deleted             = 1
      gtp_v2_c_half_open_tunnel_closed    = 1
      gtp_u_tunnel_created                = 1
      gtp_u_tunnel_deleted                = 1
      gtp_v0_c_update_pdp_resp_unsuccess  = 1
      gtp_v1_c_update_pdp_resp_unsuccess  = 1
      gtp_v2_c_mod_bearer_resp_unsuccess  = 1
      gtp_v0_c_create_pdp_resp_unsuccess  = 1
      gtp_v1_c_create_pdp_resp_unsuccess  = 1
      gtp_v2_c_create_sess_resp_unsuccess = 1
      gtp_v2_c_piggyback_message          = 1
      gtp_path_management_message         = 1
      gtp_v0_c_tunnel_deleted_restart     = 1
    }
  }
  gtp_network_element_topn_tmpl_list {
    name     = "test"
    user_tag = "106"
    metrics {
      uplink_bytes                        = 1
      downlink_bytes                      = 1
      uplink_pkts                         = 1
      downlink_pkts                       = 1
      gtp_v0_c_tunnel_created             = 1
      gtp_v0_c_tunnel_half_open           = 1
      gtp_v0_c_tunnel_half_closed         = 1
      gtp_v0_c_tunnel_closed              = 1
      gtp_v0_c_tunnel_deleted             = 1
      gtp_v0_c_half_open_tunnel_closed    = 1
      gtp_v1_c_tunnel_created             = 1
      gtp_v1_c_tunnel_half_open           = 1
      gtp_v1_c_tunnel_half_closed         = 1
      gtp_v1_c_tunnel_closed              = 1
      gtp_v1_c_tunnel_deleted             = 1
      gtp_v1_c_half_open_tunnel_closed    = 1
      gtp_v2_c_tunnel_created             = 1
      gtp_v2_c_tunnel_half_open           = 1
      gtp_v2_c_tunnel_half_closed         = 1
      gtp_v2_c_tunnel_closed              = 1
      gtp_v2_c_tunnel_deleted             = 1
      gtp_v2_c_half_open_tunnel_closed    = 1
      gtp_u_tunnel_created                = 1
      gtp_u_tunnel_deleted                = 1
      gtp_v0_c_update_pdp_resp_unsuccess  = 1
      gtp_v1_c_update_pdp_resp_unsuccess  = 1
      gtp_v2_c_mod_bearer_resp_unsuccess  = 1
      gtp_v0_c_create_pdp_resp_unsuccess  = 1
      gtp_v1_c_create_pdp_resp_unsuccess  = 1
      gtp_v2_c_create_sess_resp_unsuccess = 1
      gtp_v2_c_piggyback_message          = 1
      gtp_path_management_message         = 1
      gtp_v0_c_tunnel_deleted_restart     = 1
      gtp_v1_c_tunnel_deleted_restart     = 1
      gtp_v2_c_tunnel_deleted_restart     = 1
      gtp_v0_c_reserved_message_allow     = 1
      gtp_v1_c_reserved_message_allow     = 1
      gtp_v2_c_reserved_message_allow     = 1
      drop_vld_reserved_field_set         = 1
      drop_vld_tunnel_id_flag             = 1
      drop_vld_invalid_flow_label_v0      = 1
      drop_vld_invalid_teid               = 1
      drop_vld_unsupported_message_type   = 1
    }
  }

}