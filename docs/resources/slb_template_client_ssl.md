---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_template_client_ssl Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_slb_template_client_ssl (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **ad_group_list** (String)
- **alert_type** (String)
- **auth_sg** (String)
- **auth_sg_dn** (Number)
- **auth_sg_filter** (String)
- **auth_username** (String)
- **auth_username_attribute** (String)
- **authen_name** (String)
- **authorization** (Number)
- **bypass_cert_issuer_class_list_name** (String)
- **bypass_cert_issuer_multi_class_list** (Block List) (see [below for nested schema](#nestedblock--bypass_cert_issuer_multi_class_list))
- **bypass_cert_san_class_list_name** (String)
- **bypass_cert_san_multi_class_list** (Block List) (see [below for nested schema](#nestedblock--bypass_cert_san_multi_class_list))
- **bypass_cert_subject_class_list_name** (String)
- **bypass_cert_subject_multi_class_list** (Block List) (see [below for nested schema](#nestedblock--bypass_cert_subject_multi_class_list))
- **ca_certs** (Block List) (see [below for nested schema](#nestedblock--ca_certs))
- **cache_persistence_list_name** (String)
- **case_insensitive** (Number)
- **central_cert_pin_list** (Number)
- **cert_revoke_action** (String)
- **cert_unknown_action** (String)
- **certificate_issuer_contains_list** (Block List) (see [below for nested schema](#nestedblock--certificate_issuer_contains_list))
- **certificate_issuer_ends_with_list** (Block List) (see [below for nested schema](#nestedblock--certificate_issuer_ends_with_list))
- **certificate_issuer_equals_list** (Block List) (see [below for nested schema](#nestedblock--certificate_issuer_equals_list))
- **certificate_issuer_starts_with_list** (Block List) (see [below for nested schema](#nestedblock--certificate_issuer_starts_with_list))
- **certificate_list** (Block List) (see [below for nested schema](#nestedblock--certificate_list))
- **certificate_san_contains_list** (Block List) (see [below for nested schema](#nestedblock--certificate_san_contains_list))
- **certificate_san_ends_with_list** (Block List) (see [below for nested schema](#nestedblock--certificate_san_ends_with_list))
- **certificate_san_equals_list** (Block List) (see [below for nested schema](#nestedblock--certificate_san_equals_list))
- **certificate_san_starts_with_list** (Block List) (see [below for nested schema](#nestedblock--certificate_san_starts_with_list))
- **certificate_subject_contains_list** (Block List) (see [below for nested schema](#nestedblock--certificate_subject_contains_list))
- **certificate_subject_ends_with_list** (Block List) (see [below for nested schema](#nestedblock--certificate_subject_ends_with_list))
- **certificate_subject_equals_list** (Block List) (see [below for nested schema](#nestedblock--certificate_subject_equals_list))
- **certificate_subject_starts_with_list** (Block List) (see [below for nested schema](#nestedblock--certificate_subject_starts_with_list))
- **chain_cert** (String)
- **chain_cert_shared_str** (String)
- **cipher_without_prio_list** (Block List) (see [below for nested schema](#nestedblock--cipher_without_prio_list))
- **class_list_name** (String)
- **client_auth_case_insensitive** (Number)
- **client_auth_class_list** (String)
- **client_auth_contains_list** (Block List) (see [below for nested schema](#nestedblock--client_auth_contains_list))
- **client_auth_ends_with_list** (Block List) (see [below for nested schema](#nestedblock--client_auth_ends_with_list))
- **client_auth_equals_list** (Block List) (see [below for nested schema](#nestedblock--client_auth_equals_list))
- **client_auth_starts_with_list** (Block List) (see [below for nested schema](#nestedblock--client_auth_starts_with_list))
- **client_certificate** (String)
- **client_ipv4_list** (Block List) (see [below for nested schema](#nestedblock--client_ipv4_list))
- **client_ipv6_list** (Block List) (see [below for nested schema](#nestedblock--client_ipv6_list))
- **close_notify** (Number)
- **contains_list** (Block List) (see [below for nested schema](#nestedblock--contains_list))
- **crl_certs** (Block List) (see [below for nested schema](#nestedblock--crl_certs))
- **dgversion** (Number)
- **dh_type** (String)
- **direct_client_server_auth** (Number)
- **disable_sslv3** (Number)
- **early_data** (Number)
- **ec_list** (Block List) (see [below for nested schema](#nestedblock--ec_list))
- **enable_ssli_ftp_alg** (Number)
- **enable_tls_alert_logging** (Number)
- **ends_with_list** (Block List) (see [below for nested schema](#nestedblock--ends_with_list))
- **equals_list** (Block List) (see [below for nested schema](#nestedblock--equals_list))
- **exception_ad_group_list** (String)
- **exception_certificate_issuer_cl_name** (String)
- **exception_certificate_san_cl_name** (String)
- **exception_certificate_subject_cl_name** (String)
- **exception_client_ipv4_list** (Block List) (see [below for nested schema](#nestedblock--exception_client_ipv4_list))
- **exception_client_ipv6_list** (Block List) (see [below for nested schema](#nestedblock--exception_client_ipv6_list))
- **exception_server_ipv4_list** (Block List) (see [below for nested schema](#nestedblock--exception_server_ipv4_list))
- **exception_server_ipv6_list** (Block List) (see [below for nested schema](#nestedblock--exception_server_ipv6_list))
- **exception_sni_cl_name** (String)
- **exception_user_name_list** (String)
- **exception_web_category** (Block List, Max: 1) (see [below for nested schema](#nestedblock--exception_web_category))
- **exception_web_reputation** (Block List, Max: 1) (see [below for nested schema](#nestedblock--exception_web_reputation))
- **expire_hours** (Number)
- **forward_passphrase** (String)
- **forward_proxy_alt_sign** (Number)
- **forward_proxy_block_message** (String)
- **forward_proxy_ca_cert** (String)
- **forward_proxy_ca_key** (String)
- **forward_proxy_cert_cache_limit** (Number)
- **forward_proxy_cert_cache_timeout** (Number)
- **forward_proxy_cert_expiry** (Number)
- **forward_proxy_cert_not_ready_action** (String)
- **forward_proxy_cert_revoke_action** (Number)
- **forward_proxy_cert_unknown_action** (Number)
- **forward_proxy_crl_disable** (Number)
- **forward_proxy_decrypted_dscp** (Number)
- **forward_proxy_decrypted_dscp_bypass** (Number)
- **forward_proxy_enable** (Number)
- **forward_proxy_esni_action** (Number)
- **forward_proxy_failsafe_disable** (Number)
- **forward_proxy_hash_persistence_interval** (Number)
- **forward_proxy_log_disable** (Number)
- **forward_proxy_no_shared_cipher_action** (Number)
- **forward_proxy_no_sni_action** (String)
- **forward_proxy_ocsp_disable** (Number)
- **forward_proxy_require_sni_cert_matched** (String)
- **forward_proxy_selfsign_redir** (Number)
- **forward_proxy_ssl_version** (Number)
- **forward_proxy_trusted_ca_lists** (Block List) (see [below for nested schema](#nestedblock--forward_proxy_trusted_ca_lists))
- **forward_proxy_verify_cert_fail_action** (Number)
- **fp_alt_cert** (String)
- **fp_alt_chain_cert** (String)
- **fp_alt_key** (String)
- **fp_alt_passphrase** (String)
- **fp_alt_shared** (Number)
- **fp_ca_certificate** (String)
- **fp_ca_certificate_shared** (Number)
- **fp_ca_chain_cert** (String)
- **fp_ca_key** (String)
- **fp_ca_key_passphrase** (String)
- **fp_ca_key_shared** (Number)
- **fp_ca_shared** (Number)
- **fp_cert_ext_aia_ca_issuers** (String)
- **fp_cert_ext_aia_ocsp** (String)
- **fp_cert_ext_crldp** (String)
- **fp_cert_fetch_autonat** (String)
- **fp_cert_fetch_autonat_precedence** (Number)
- **fp_cert_fetch_natpool_name** (String)
- **fp_cert_fetch_natpool_name_shared** (String)
- **fp_cert_fetch_natpool_precedence** (Number)
- **fp_esni_action** (String)
- **handshake_logging_enable** (Number)
- **hsm_type** (String)
- **id** (String) The ID of this resource.
- **inspect_certificate_issuer_cl_name** (String)
- **inspect_certificate_san_cl_name** (String)
- **inspect_certificate_subject_cl_name** (String)
- **inspect_list_name** (String)
- **ja3_enable** (Number)
- **ja3_insert_http_header** (String)
- **ja3_reject_class_list** (String)
- **ja3_reject_max_number_per_host** (Number)
- **ja3_ttl** (Number)
- **ldap_base_dn_from_cert** (Number)
- **ldap_search_filter** (String)
- **local_cert_pin_list** (Block List, Max: 1) (see [below for nested schema](#nestedblock--local_cert_pin_list))
- **local_logging** (Number)
- **multi_class_list** (Block List) (see [below for nested schema](#nestedblock--multi_class_list))
- **name** (String)
- **no_anti_replay** (Number)
- **no_shared_cipher_action** (String)
- **non_ssl_bypass_l4session** (Number)
- **non_ssl_bypass_service_group** (String)
- **notafter** (Number)
- **notafterday** (Number)
- **notaftermonth** (Number)
- **notafteryear** (Number)
- **notbefore** (Number)
- **notbeforeday** (Number)
- **notbeforemonth** (Number)
- **notbeforeyear** (Number)
- **ocsp_stapling** (Number)
- **ocspst_ca_cert** (String)
- **ocspst_ocsp** (Number)
- **ocspst_sg** (String)
- **ocspst_sg_days** (Number)
- **ocspst_sg_hours** (Number)
- **ocspst_sg_minutes** (Number)
- **ocspst_sg_timeout** (Number)
- **ocspst_srvr** (String)
- **ocspst_srvr_days** (Number)
- **ocspst_srvr_hours** (Number)
- **ocspst_srvr_minutes** (Number)
- **ocspst_srvr_timeout** (Number)
- **renegotiation_disable** (Number)
- **req_ca_lists** (Block List) (see [below for nested schema](#nestedblock--req_ca_lists))
- **require_web_category** (Number)
- **server_ipv4_list** (Block List) (see [below for nested schema](#nestedblock--server_ipv4_list))
- **server_ipv6_list** (Block List) (see [below for nested schema](#nestedblock--server_ipv6_list))
- **server_name_auto_map** (Number)
- **server_name_list** (Block List) (see [below for nested schema](#nestedblock--server_name_list))
- **session_cache_size** (Number)
- **session_cache_timeout** (Number)
- **session_ticket_disable** (Number)
- **session_ticket_lifetime** (Number)
- **shared_partition_cipher_template** (Number)
- **shared_partition_pool** (Number)
- **sni_bypass_enable_log** (Number)
- **sni_bypass_expired_cert** (Number)
- **sni_bypass_explicit_list** (String)
- **sni_bypass_missing_cert** (Number)
- **sni_enable_log** (Number)
- **ssl_false_start_disable** (Number)
- **ssli_logging** (Number)
- **sslilogging** (String)
- **sslv2_bypass_service_group** (String)
- **starts_with_list** (Block List) (see [below for nested schema](#nestedblock--starts_with_list))
- **template_cipher** (String)
- **template_cipher_shared** (String)
- **template_hsm** (String)
- **user_name_list** (String)
- **user_tag** (String)
- **uuid** (String)
- **verify_cert_fail_action** (String)
- **version** (Number)
- **web_category** (Block List, Max: 1) (see [below for nested schema](#nestedblock--web_category))
- **web_reputation** (Block List, Max: 1) (see [below for nested schema](#nestedblock--web_reputation))

<a id="nestedblock--bypass_cert_issuer_multi_class_list"></a>
### Nested Schema for `bypass_cert_issuer_multi_class_list`

Optional:

- **bypass_cert_issuer_multi_class_list_name** (String)


<a id="nestedblock--bypass_cert_san_multi_class_list"></a>
### Nested Schema for `bypass_cert_san_multi_class_list`

Optional:

- **bypass_cert_san_multi_class_list_name** (String)


<a id="nestedblock--bypass_cert_subject_multi_class_list"></a>
### Nested Schema for `bypass_cert_subject_multi_class_list`

Optional:

- **bypass_cert_subject_multi_class_list_name** (String)


<a id="nestedblock--ca_certs"></a>
### Nested Schema for `ca_certs`

Optional:

- **ca_cert** (String)
- **ca_shared** (Number)
- **client_ocsp** (Number)
- **client_ocsp_sg** (String)
- **client_ocsp_srvr** (String)


<a id="nestedblock--certificate_issuer_contains_list"></a>
### Nested Schema for `certificate_issuer_contains_list`

Optional:

- **certificate_issuer_contains** (String)


<a id="nestedblock--certificate_issuer_ends_with_list"></a>
### Nested Schema for `certificate_issuer_ends_with_list`

Optional:

- **certificate_issuer_ends_with** (String)


<a id="nestedblock--certificate_issuer_equals_list"></a>
### Nested Schema for `certificate_issuer_equals_list`

Optional:

- **certificate_issuer_equals** (String)


<a id="nestedblock--certificate_issuer_starts_with_list"></a>
### Nested Schema for `certificate_issuer_starts_with_list`

Optional:

- **certificate_issuer_starts** (String)


<a id="nestedblock--certificate_list"></a>
### Nested Schema for `certificate_list`

Optional:

- **cert** (String)
- **chain_cert** (String)
- **key** (String)
- **passphrase** (String)
- **shared** (Number)
- **uuid** (String)


<a id="nestedblock--certificate_san_contains_list"></a>
### Nested Schema for `certificate_san_contains_list`

Optional:

- **certificate_san_contains** (String)


<a id="nestedblock--certificate_san_ends_with_list"></a>
### Nested Schema for `certificate_san_ends_with_list`

Optional:

- **certificate_san_ends_with** (String)


<a id="nestedblock--certificate_san_equals_list"></a>
### Nested Schema for `certificate_san_equals_list`

Optional:

- **certificate_san_equals** (String)


<a id="nestedblock--certificate_san_starts_with_list"></a>
### Nested Schema for `certificate_san_starts_with_list`

Optional:

- **certificate_san_starts** (String)


<a id="nestedblock--certificate_subject_contains_list"></a>
### Nested Schema for `certificate_subject_contains_list`

Optional:

- **certificate_subject_contains** (String)


<a id="nestedblock--certificate_subject_ends_with_list"></a>
### Nested Schema for `certificate_subject_ends_with_list`

Optional:

- **certificate_subject_ends_with** (String)


<a id="nestedblock--certificate_subject_equals_list"></a>
### Nested Schema for `certificate_subject_equals_list`

Optional:

- **certificate_subject_equals** (String)


<a id="nestedblock--certificate_subject_starts_with_list"></a>
### Nested Schema for `certificate_subject_starts_with_list`

Optional:

- **certificate_subject_starts** (String)


<a id="nestedblock--cipher_without_prio_list"></a>
### Nested Schema for `cipher_without_prio_list`

Optional:

- **cipher_wo_prio** (String)


<a id="nestedblock--client_auth_contains_list"></a>
### Nested Schema for `client_auth_contains_list`

Optional:

- **client_auth_contains** (String)


<a id="nestedblock--client_auth_ends_with_list"></a>
### Nested Schema for `client_auth_ends_with_list`

Optional:

- **client_auth_ends_with** (String)


<a id="nestedblock--client_auth_equals_list"></a>
### Nested Schema for `client_auth_equals_list`

Optional:

- **client_auth_equals** (String)


<a id="nestedblock--client_auth_starts_with_list"></a>
### Nested Schema for `client_auth_starts_with_list`

Optional:

- **client_auth_starts_with** (String)


<a id="nestedblock--client_ipv4_list"></a>
### Nested Schema for `client_ipv4_list`

Optional:

- **client_ipv4_list_name** (String)


<a id="nestedblock--client_ipv6_list"></a>
### Nested Schema for `client_ipv6_list`

Optional:

- **client_ipv6_list_name** (String)


<a id="nestedblock--contains_list"></a>
### Nested Schema for `contains_list`

Optional:

- **contains** (String)


<a id="nestedblock--crl_certs"></a>
### Nested Schema for `crl_certs`

Optional:

- **crl** (String)
- **crl_shared** (Number)


<a id="nestedblock--ec_list"></a>
### Nested Schema for `ec_list`

Optional:

- **ec** (String)


<a id="nestedblock--ends_with_list"></a>
### Nested Schema for `ends_with_list`

Optional:

- **ends_with** (String)


<a id="nestedblock--equals_list"></a>
### Nested Schema for `equals_list`

Optional:

- **equals** (String)


<a id="nestedblock--exception_client_ipv4_list"></a>
### Nested Schema for `exception_client_ipv4_list`

Optional:

- **exception_client_ipv4_list_name** (String)


<a id="nestedblock--exception_client_ipv6_list"></a>
### Nested Schema for `exception_client_ipv6_list`

Optional:

- **exception_client_ipv6_list_name** (String)


<a id="nestedblock--exception_server_ipv4_list"></a>
### Nested Schema for `exception_server_ipv4_list`

Optional:

- **exception_server_ipv4_list_name** (String)


<a id="nestedblock--exception_server_ipv6_list"></a>
### Nested Schema for `exception_server_ipv6_list`

Optional:

- **exception_server_ipv6_list_name** (String)


<a id="nestedblock--exception_web_category"></a>
### Nested Schema for `exception_web_category`

Optional:

- **exception_abortion** (Number)
- **exception_adult_and_pornography** (Number)
- **exception_alcohol_and_tobacco** (Number)
- **exception_auctions** (Number)
- **exception_bot_nets** (Number)
- **exception_business_and_economy** (Number)
- **exception_cdns** (Number)
- **exception_cheating** (Number)
- **exception_computer_and_internet_info** (Number)
- **exception_computer_and_internet_security** (Number)
- **exception_confirmed_spam_sources** (Number)
- **exception_cult_and_occult** (Number)
- **exception_dating** (Number)
- **exception_dead_sites** (Number)
- **exception_drugs** (Number)
- **exception_dynamic_comment** (Number)
- **exception_educational_institutions** (Number)
- **exception_entertainment_and_arts** (Number)
- **exception_fashion_and_beauty** (Number)
- **exception_financial_services** (Number)
- **exception_food_and_dining** (Number)
- **exception_gambling** (Number)
- **exception_games** (Number)
- **exception_government** (Number)
- **exception_gross** (Number)
- **exception_hacking** (Number)
- **exception_hate_and_racism** (Number)
- **exception_health_and_medicine** (Number)
- **exception_home_and_garden** (Number)
- **exception_hunting_and_fishing** (Number)
- **exception_illegal** (Number)
- **exception_illegal_pornography** (Number)
- **exception_image_and_video_search** (Number)
- **exception_internet_communications** (Number)
- **exception_internet_portals** (Number)
- **exception_job_search** (Number)
- **exception_keyloggers_and_monitoring** (Number)
- **exception_kids** (Number)
- **exception_legal** (Number)
- **exception_local_information** (Number)
- **exception_malware_sites** (Number)
- **exception_marijuana** (Number)
- **exception_military** (Number)
- **exception_motor_vehicles** (Number)
- **exception_music** (Number)
- **exception_news_and_media** (Number)
- **exception_nudity** (Number)
- **exception_nudity_artistic** (Number)
- **exception_online_greeting_cards** (Number)
- **exception_open_http_proxies** (Number)
- **exception_parked_domains** (Number)
- **exception_pay_to_surf** (Number)
- **exception_peer_to_peer** (Number)
- **exception_personal_sites_and_blogs** (Number)
- **exception_personal_storage** (Number)
- **exception_philosophy_and_politics** (Number)
- **exception_phishing_and_other_fraud** (Number)
- **exception_private_ip_addresses** (Number)
- **exception_proxy_avoid_and_anonymizers** (Number)
- **exception_questionable** (Number)
- **exception_real_estate** (Number)
- **exception_recreation_and_hobbies** (Number)
- **exception_reference_and_research** (Number)
- **exception_religion** (Number)
- **exception_search_engines** (Number)
- **exception_sex_education** (Number)
- **exception_shareware_and_freeware** (Number)
- **exception_shopping** (Number)
- **exception_social_network** (Number)
- **exception_society** (Number)
- **exception_spam_urls** (Number)
- **exception_sports** (Number)
- **exception_spyware_and_adware** (Number)
- **exception_stock_advice_and_tools** (Number)
- **exception_streaming_media** (Number)
- **exception_swimsuits_and_intimate_apparel** (Number)
- **exception_training_and_tools** (Number)
- **exception_translation** (Number)
- **exception_travel** (Number)
- **exception_uncategorized** (Number)
- **exception_unconfirmed_spam_sources** (Number)
- **exception_violence** (Number)
- **exception_weapons** (Number)
- **exception_web_advertisements** (Number)
- **exception_web_based_email** (Number)
- **exception_web_hosting_sites** (Number)


<a id="nestedblock--exception_web_reputation"></a>
### Nested Schema for `exception_web_reputation`

Optional:

- **exception_low_risk** (Number)
- **exception_malicious** (Number)
- **exception_moderate_risk** (Number)
- **exception_suspicious** (Number)
- **exception_threshold** (Number)
- **exception_trustworthy** (Number)


<a id="nestedblock--forward_proxy_trusted_ca_lists"></a>
### Nested Schema for `forward_proxy_trusted_ca_lists`

Optional:

- **forward_proxy_trusted_ca** (String)
- **fp_trusted_ca_shared** (Number)


<a id="nestedblock--local_cert_pin_list"></a>
### Nested Schema for `local_cert_pin_list`

Optional:

- **local_cert_pin_list_bypass_fail_count** (Number)


<a id="nestedblock--multi_class_list"></a>
### Nested Schema for `multi_class_list`

Optional:

- **multi_clist_name** (String)


<a id="nestedblock--req_ca_lists"></a>
### Nested Schema for `req_ca_lists`

Optional:

- **client_cert_req_ca_shared** (Number)
- **client_certificate_request_ca** (String)


<a id="nestedblock--server_ipv4_list"></a>
### Nested Schema for `server_ipv4_list`

Optional:

- **server_ipv4_list_name** (String)


<a id="nestedblock--server_ipv6_list"></a>
### Nested Schema for `server_ipv6_list`

Optional:

- **server_ipv6_list_name** (String)


<a id="nestedblock--server_name_list"></a>
### Nested Schema for `server_name_list`

Optional:

- **server_cert** (String)
- **server_cert_regex** (String)
- **server_chain** (String)
- **server_chain_regex** (String)
- **server_key** (String)
- **server_key_regex** (String)
- **server_name** (String)
- **server_name_alternate** (Number)
- **server_name_regex** (String)
- **server_name_regex_alternate** (Number)
- **server_passphrase** (String)
- **server_passphrase_regex** (String)
- **server_shared** (Number)
- **server_shared_regex** (Number)


<a id="nestedblock--starts_with_list"></a>
### Nested Schema for `starts_with_list`

Optional:

- **starts_with** (String)


<a id="nestedblock--web_category"></a>
### Nested Schema for `web_category`

Optional:

- **abortion** (Number)
- **adult_and_pornography** (Number)
- **alcohol_and_tobacco** (Number)
- **auctions** (Number)
- **bot_nets** (Number)
- **business_and_economy** (Number)
- **cdns** (Number)
- **cheating** (Number)
- **computer_and_internet_info** (Number)
- **computer_and_internet_security** (Number)
- **confirmed_spam_sources** (Number)
- **cult_and_occult** (Number)
- **dating** (Number)
- **dead_sites** (Number)
- **drugs** (Number)
- **dynamic_comment** (Number)
- **educational_institutions** (Number)
- **entertainment_and_arts** (Number)
- **fashion_and_beauty** (Number)
- **financial_services** (Number)
- **food_and_dining** (Number)
- **gambling** (Number)
- **games** (Number)
- **government** (Number)
- **gross** (Number)
- **hacking** (Number)
- **hate_and_racism** (Number)
- **health_and_medicine** (Number)
- **home_and_garden** (Number)
- **hunting_and_fishing** (Number)
- **illegal** (Number)
- **illegal_pornography** (Number)
- **image_and_video_search** (Number)
- **internet_communications** (Number)
- **internet_portals** (Number)
- **job_search** (Number)
- **keyloggers_and_monitoring** (Number)
- **kids** (Number)
- **legal** (Number)
- **local_information** (Number)
- **malware_sites** (Number)
- **marijuana** (Number)
- **military** (Number)
- **motor_vehicles** (Number)
- **music** (Number)
- **news_and_media** (Number)
- **nudity** (Number)
- **nudity_artistic** (Number)
- **online_greeting_cards** (Number)
- **open_http_proxies** (Number)
- **parked_domains** (Number)
- **pay_to_surf** (Number)
- **peer_to_peer** (Number)
- **personal_sites_and_blogs** (Number)
- **personal_storage** (Number)
- **philosophy_and_politics** (Number)
- **phishing_and_other_fraud** (Number)
- **private_ip_addresses** (Number)
- **proxy_avoid_and_anonymizers** (Number)
- **questionable** (Number)
- **real_estate** (Number)
- **recreation_and_hobbies** (Number)
- **reference_and_research** (Number)
- **religion** (Number)
- **search_engines** (Number)
- **sex_education** (Number)
- **shareware_and_freeware** (Number)
- **shopping** (Number)
- **social_network** (Number)
- **society** (Number)
- **spam_urls** (Number)
- **sports** (Number)
- **spyware_and_adware** (Number)
- **stock_advice_and_tools** (Number)
- **streaming_media** (Number)
- **swimsuits_and_intimate_apparel** (Number)
- **training_and_tools** (Number)
- **translation** (Number)
- **travel** (Number)
- **uncategorized** (Number)
- **unconfirmed_spam_sources** (Number)
- **violence** (Number)
- **weapons** (Number)
- **web_advertisements** (Number)
- **web_based_email** (Number)
- **web_hosting_sites** (Number)


<a id="nestedblock--web_reputation"></a>
### Nested Schema for `web_reputation`

Optional:

- **bypass_low_risk** (Number)
- **bypass_malicious** (Number)
- **bypass_moderate_risk** (Number)
- **bypass_suspicious** (Number)
- **bypass_threshold** (Number)
- **bypass_trustworthy** (Number)

