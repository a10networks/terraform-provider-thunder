---
layout: "thunder"
page_title: "thunder: thunder_slb_template_client_ssl"
sidebar_current: "docs-thunder-resource-slb-template-client-ssl"
description: |-
    Provides details about thunder slb template client ssl resource for A10
---

# thunder\_slb\_template\_client\_ssl

`thunder_slb_template_client_ssl` Provides details about thunder slb template client ssl
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_client_ssl" "resourceSlbTemplateClientSslTest" {
	name = "string"
auth_username = "string"
ca-certs {   
	ca_cert =  "string" 
	ca_shared =  0 
	client_ocsp =  0 
	client_ocsp_srvr =  "string" 
	client_ocsp_sg =  "string" 
	}
chain_cert = "string"
chain_cert_shared_str = "string"
dh_type = "string"
ec-list {   
	ec =  "string" 
	}
local_logging = 0
ocsp_stapling = 0
ocspst_ca_cert = "string"
ocspst_ocsp = 0
ocspst_srvr = "string"
ocspst_srvr_days = 0
ocspst_srvr_hours = 0
ocspst_srvr_minutes = 0
ocspst_srvr_timeout = 0
ocspst_sg = "string"
ocspst_sg_days = 0
ocspst_sg_hours = 0
ocspst_sg_minutes = 0
ocspst_sg_timeout = 0
ssli_logging = 0
sslilogging = "string"
client_certificate = "string"
req-ca-lists {   
	client_certificate_request_ca =  "string" 
	client_cert_req_ca_shared =  0 
	}
close_notify = 0
crl-certs {   
	crl =  "string" 
	crl_shared =  0 
	}
forward_proxy_ca_cert = "string"
fp_ca_shared = 0
forward_proxy_ca_key = "string"
forward_passphrase = "string"
fp_ca_key_shared = 0
fp_ca_certificate = "string"
fp_ca_key = "string"
fp_ca_key_passphrase = "string"
fp_ca_chain_cert = "string"
fp_ca_certificate_shared = 0
forward_proxy_alt_sign = 0
fp_alt_cert = "string"
fp_alt_key = "string"
fp_alt_passphrase = "string"
fp_alt_chain_cert = "string"
fp_alt_shared = 0
forward-proxy-trusted-ca-lists {   
	forward_proxy_trusted_ca =  "string" 
	fp_trusted_ca_shared =  0 
	}
forward_proxy_decrypted_dscp = 0
forward_proxy_decrypted_dscp_bypass = 0
enable_tls_alert_logging = 0
alert_type = "string"
forward_proxy_verify_cert_fail_action = 0
verify_cert_fail_action = "string"
forward_proxy_cert_revoke_action = 0
cert_revoke_action = "string"
forward_proxy_no_shared_cipher_action = 0
no_shared_cipher_action = "string"
forward_proxy_esni_action = 0
fp_esni_action = "string"
forward_proxy_cert_unknown_action = 0
cert_unknown_action = "string"
forward_proxy_block_message = "string"
cache_persistence_list_name = "string"
fp_cert_ext_crldp = "string"
fp_cert_ext_aia_ocsp = "string"
fp_cert_ext_aia_ca_issuers = "string"
notbefore = 0
notbeforeday = 0
notbeforemonth = 0
notbeforeyear = 0
notafter = 0
notafterday = 0
notaftermonth = 0
notafteryear = 0
forward_proxy_hash_persistence_interval = 0
forward_proxy_ssl_version = 0
forward_proxy_ocsp_disable = 0
forward_proxy_crl_disable = 0
forward_proxy_cert_cache_timeout = 0
forward_proxy_cert_cache_limit = 0
forward_proxy_cert_expiry = 0
expire_hours = 0
forward_proxy_enable = 0
handshake_logging_enable = 0
forward_proxy_selfsign_redir = 0
forward_proxy_failsafe_disable = 0
forward_proxy_log_disable = 0
fp_cert_fetch_natpool_name = "string"
shared_partition_pool = 0
fp_cert_fetch_natpool_name_shared = "string"
fp_cert_fetch_natpool_precedence = 0
fp_cert_fetch_autonat = "string"
fp_cert_fetch_autonat_precedence = 0
forward_proxy_no_sni_action = "string"
case_insensitive = 0
class_list_name = "string"
multi-class-list {   
	multi_clist_name =  "string" 
	}
user_name_list = "string"
ad_group_list = "string"
exception_user_name_list = "string"
exception_ad_group_list = "string"
exception_sni_cl_name = "string"
inspect_list_name = "string"
inspect_certificate_subject_cl_name = "string"
inspect_certificate_issuer_cl_name = "string"
inspect_certificate_san_cl_name = "string"
contains-list {   
	contains =  "string" 
	}
ends-with-list {   
	ends_with =  "string" 
	}
equals-list {   
	equals =  "string" 
	}
starts-with-list {   
	starts_with =  "string" 
	}
certificate-subject-contains-list {   
	certificate_subject_contains =  "string" 
	}
bypass_cert_subject_class_list_name = "string"
bypass-cert-subject-multi-class-list {   
	bypass_cert_subject_multi_class_list_name =  "string" 
	}
exception_certificate_subject_cl_name = "string"
certificate-subject-ends-with-list {   
	certificate_subject_ends_with =  "string" 
	}
certificate-subject-equals-list {   
	certificate_subject_equals =  "string" 
	}
certificate-subject-starts-with-list {   
	certificate_subject_starts =  "string" 
	}
certificate-issuer-contains-list {   
	certificate_issuer_contains =  "string" 
	}
bypass_cert_issuer_class_list_name = "string"
bypass-cert-issuer-multi-class-list {   
	bypass_cert_issuer_multi_class_list_name =  "string" 
	}
exception_certificate_issuer_cl_name = "string"
certificate-issuer-ends-with-list {   
	certificate_issuer_ends_with =  "string" 
	}
certificate-issuer-equals-list {   
	certificate_issuer_equals =  "string" 
	}
certificate-issuer-starts-with-list {   
	certificate_issuer_starts =  "string" 
	}
certificate-san-contains-list {   
	certificate_san_contains =  "string" 
	}
bypass_cert_san_class_list_name = "string"
bypass-cert-san-multi-class-list {   
	bypass_cert_san_multi_class_list_name =  "string" 
	}
exception_certificate_san_cl_name = "string"
certificate-san-ends-with-list {   
	certificate_san_ends_with =  "string" 
	}
certificate-san-equals-list {   
	certificate_san_equals =  "string" 
	}
certificate-san-starts-with-list {   
	certificate_san_starts =  "string" 
	}
client_auth_case_insensitive = 0
client_auth_class_list = "string"
client-auth-contains-list {   
	client_auth_contains =  "string" 
	}
client-auth-ends-with-list {   
	client_auth_ends_with =  "string" 
	}
client-auth-equals-list {   
	client_auth_equals =  "string" 
	}
client-auth-starts-with-list {   
	client_auth_starts_with =  "string" 
	}
forward_proxy_cert_not_ready_action = "string"
web_reputation {  
 	bypass_trustworthy =  0 
	bypass_low_risk =  0 
	bypass_moderate_risk =  0 
	bypass_suspicious =  0 
	bypass_malicious =  0 
	bypass_threshold =  0 
	}
exception_web_reputation {  
 	exception_trustworthy =  0 
	exception_low_risk =  0 
	exception_moderate_risk =  0 
	exception_suspicious =  0 
	exception_malicious =  0 
	exception_threshold =  0 
	}
web_category {  
 	uncategorized =  0 
	real_estate =  0 
	computer_and_internet_security =  0 
	financial_services =  0 
	business_and_economy =  0 
	computer_and_internet_info =  0 
	auctions =  0 
	shopping =  0 
	cult_and_occult =  0 
	travel =  0 
	drugs =  0 
	adult_and_pornography =  0 
	home_and_garden =  0 
	military =  0 
	social_network =  0 
	dead_sites =  0 
	stock_advice_and_tools =  0 
	training_and_tools =  0 
	dating =  0 
	sex_education =  0 
	religion =  0 
	entertainment_and_arts =  0 
	personal_sites_and_blogs =  0 
	legal =  0 
	local_information =  0 
	streaming_media =  0 
	job_search =  0 
	gambling =  0 
	translation =  0 
	reference_and_research =  0 
	shareware_and_freeware =  0 
	peer_to_peer =  0 
	marijuana =  0 
	hacking =  0 
	games =  0 
	philosophy_and_politics =  0 
	weapons =  0 
	pay_to_surf =  0 
	hunting_and_fishing =  0 
	society =  0 
	educational_institutions =  0 
	online_greeting_cards =  0 
	sports =  0 
	swimsuits_and_intimate_apparel =  0 
	questionable =  0 
	kids =  0 
	hate_and_racism =  0 
	personal_storage =  0 
	violence =  0 
	keyloggers_and_monitoring =  0 
	search_engines =  0 
	internet_portals =  0 
	web_advertisements =  0 
	cheating =  0 
	gross =  0 
	web_based_email =  0 
	malware_sites =  0 
	phishing_and_other_fraud =  0 
	proxy_avoid_and_anonymizers =  0 
	spyware_and_adware =  0 
	music =  0 
	government =  0 
	nudity =  0 
	news_and_media =  0 
	illegal =  0 
	cdns =  0 
	internet_communications =  0 
	bot_nets =  0 
	abortion =  0 
	health_and_medicine =  0 
	confirmed_spam_sources =  0 
	spam_urls =  0 
	unconfirmed_spam_sources =  0 
	open_http_proxies =  0 
	dynamic_comment =  0 
	parked_domains =  0 
	alcohol_and_tobacco =  0 
	private_ip_addresses =  0 
	image_and_video_search =  0 
	fashion_and_beauty =  0 
	recreation_and_hobbies =  0 
	motor_vehicles =  0 
	web_hosting_sites =  0 
	food_and_dining =  0 
	nudity_artistic =  0 
	illegal_pornography =  0 
	}
exception_web_category {  
 	exception_uncategorized =  0 
	exception_real_estate =  0 
	exception_computer_and_internet_security =  0 
	exception_financial_services =  0 
	exception_business_and_economy =  0 
	exception_computer_and_internet_info =  0 
	exception_auctions =  0 
	exception_shopping =  0 
	exception_cult_and_occult =  0 
	exception_travel =  0 
	exception_drugs =  0 
	exception_adult_and_pornography =  0 
	exception_home_and_garden =  0 
	exception_military =  0 
	exception_social_network =  0 
	exception_dead_sites =  0 
	exception_stock_advice_and_tools =  0 
	exception_training_and_tools =  0 
	exception_dating =  0 
	exception_sex_education =  0 
	exception_religion =  0 
	exception_entertainment_and_arts =  0 
	exception_personal_sites_and_blogs =  0 
	exception_legal =  0 
	exception_local_information =  0 
	exception_streaming_media =  0 
	exception_job_search =  0 
	exception_gambling =  0 
	exception_translation =  0 
	exception_reference_and_research =  0 
	exception_shareware_and_freeware =  0 
	exception_peer_to_peer =  0 
	exception_marijuana =  0 
	exception_hacking =  0 
	exception_games =  0 
	exception_philosophy_and_politics =  0 
	exception_weapons =  0 
	exception_pay_to_surf =  0 
	exception_hunting_and_fishing =  0 
	exception_society =  0 
	exception_educational_institutions =  0 
	exception_online_greeting_cards =  0 
	exception_sports =  0 
	exception_swimsuits_and_intimate_apparel =  0 
	exception_questionable =  0 
	exception_kids =  0 
	exception_hate_and_racism =  0 
	exception_personal_storage =  0 
	exception_violence =  0 
	exception_keyloggers_and_monitoring =  0 
	exception_search_engines =  0 
	exception_internet_portals =  0 
	exception_web_advertisements =  0 
	exception_cheating =  0 
	exception_gross =  0 
	exception_web_based_email =  0 
	exception_malware_sites =  0 
	exception_phishing_and_other_fraud =  0 
	exception_proxy_avoid_and_anonymizers =  0 
	exception_spyware_and_adware =  0 
	exception_music =  0 
	exception_government =  0 
	exception_nudity =  0 
	exception_news_and_media =  0 
	exception_illegal =  0 
	exception_cdns =  0 
	exception_internet_communications =  0 
	exception_bot_nets =  0 
	exception_abortion =  0 
	exception_health_and_medicine =  0 
	exception_confirmed_spam_sources =  0 
	exception_spam_urls =  0 
	exception_unconfirmed_spam_sources =  0 
	exception_open_http_proxies =  0 
	exception_dynamic_comment =  0 
	exception_parked_domains =  0 
	exception_alcohol_and_tobacco =  0 
	exception_private_ip_addresses =  0 
	exception_image_and_video_search =  0 
	exception_fashion_and_beauty =  0 
	exception_recreation_and_hobbies =  0 
	exception_motor_vehicles =  0 
	exception_web_hosting_sites =  0 
	exception_food_and_dining =  0 
	exception_nudity_artistic =  0 
	exception_illegal_pornography =  0 
	}
require_web_category = 0
client-ipv4-list {   
	client_ipv4_list_name =  "string" 
	}
client-ipv6-list {   
	client_ipv6_list_name =  "string" 
	}
server-ipv4-list {   
	server_ipv4_list_name =  "string" 
	}
server-ipv6-list {   
	server_ipv6_list_name =  "string" 
	}
exception-client-ipv4-list {   
	exception_client_ipv4_list_name =  "string" 
	}
exception-client-ipv6-list {   
	exception_client_ipv6_list_name =  "string" 
	}
exception-server-ipv4-list {   
	exception_server_ipv4_list_name =  "string" 
	}
exception-server-ipv6-list {   
	exception_server_ipv6_list_name =  "string" 
	}
local_cert_pin_list {  
 	local_cert_pin_list_bypass_fail_count =  0 
	}
central_cert_pin_list = 0
forward_proxy_require_sni_cert_matched = "string"
template_cipher = "string"
shared_partition_cipher_template = 0
template_cipher_shared = "string"
template_hsm = "string"
hsm_type = "string"
cipher-without-prio-list {   
	cipher_wo_prio =  "string" 
	}
server-name-list {   
	server_name =  "string" 
	server_cert =  "string" 
	server_chain =  "string" 
	server_key =  "string" 
	server_passphrase =  "string" 
	server_name_alternate =  0 
	server_shared =  0 
	server_name_regex =  "string" 
	server_cert_regex =  "string" 
	server_chain_regex =  "string" 
	server_key_regex =  "string" 
	server_passphrase_regex =  "string" 
	server_name_regex_alternate =  0 
	server_shared_regex =  0 
	}
server_name_auto_map = 0
sni_enable_log = 0
sni_bypass_missing_cert = 0
sni_bypass_expired_cert = 0
sni_bypass_explicit_list = "string"
sni_bypass_enable_log = 0
direct_client_server_auth = 0
session_cache_size = 0
session_cache_timeout = 0
session_ticket_disable = 0
session_ticket_lifetime = 0
ssl_false_start_disable = 0
disable_sslv3 = 0
version = 0
dgversion = 0
renegotiation_disable = 0
sslv2_bypass_service_group = "string"
authorization = 0
authen_name = "string"
ldap_base_dn_from_cert = 0
ldap_search_filter = "string"
auth_sg = "string"
auth_sg_dn = 0
auth_sg_filter = "string"
auth_username_attribute = "string"
non_ssl_bypass_service_group = "string"
non_ssl_bypass_l4session = 0
enable_ssli_ftp_alg = 0
early_data = 0
no_anti_replay = 0
ja3_enable = 0
ja3_insert_http_header = "string"
ja3_reject_class_list = "string"
ja3_reject_max_number_per_host = 0
ja3_ttl = 0
uuid = "string"
user_tag = "string"
certificate-list {   
	cert =  "string" 
	key =  "string" 
	passphrase =  "string" 
	chain_cert =  "string" 
	shared =  0 
	uuid =  "string" 
	}
 
}

```

## Argument Reference

* `client-ssl` - Client SSL Template
* `name` - Client SSL Template Name
* `auth-username` - Specify the Username Field in the Client Certificate(If multi-fields are specificed, prior one has higher priority)
* `ca-cert` - CA Certificate (CA Certificate Name)
* `ca-shared` - CA Certificate Partition Shared
* `client-ocsp` - Specify ocsp authentication server(s) for client certificate verification
* `client-ocsp-srvr` - Specify authentication server
* `client-ocsp-sg` - Specify service-group (Service group name)
* `chain-cert` - Chain Certificate (Chain Certificate Name)
* `chain-cert-shared-str` - Chain Certificate Name
* `dh-type` - '1024': 1024; '1024-dsa': 1024-dsa; '2048': 2048;
* `ec` - 'secp256r1': X9_62_prime256v1; 'secp384r1': secp384r1;
* `local-logging` - Enable local logging
* `ocsp-stapling` - Config OCSP stapling support
* `ocspst-ca-cert` - CA certificate
* `ocspst-ocsp` - Specify OCSP Authentication
* `ocspst-srvr` - Specify OCSP authentication server
* `ocspst-srvr-days` - Specify update period, in days
* `ocspst-srvr-hours` - Specify update period, in hours
* `ocspst-srvr-minutes` - Specify update period, in minutes
* `ocspst-srvr-timeout` - Specify retry timeout (Default is 30 mins)
* `ocspst-sg` - Specify authentication service group
* `ocspst-sg-days` - Specify update period, in days
* `ocspst-sg-hours` - Specify update period, in hours
* `ocspst-sg-minutes` - Specify update period, in minutes
* `ocspst-sg-timeout` - Specify retry timeout (Default is 30 mins)
* `ssli-logging` - SSLi logging level, default is error logging only
* `sslilogging` - 'disable': Disable all logging; 'all': enable all logging(error, info);
* `client-certificate` - 'Ignore': Don't request client certificate; 'Require': Require client certificate; 'Request': Request client certificate;
* `client-certificate-Request-CA` - Send CA lists in certificate request (CA Certificate Name)
* `client-cert-req-ca-shared` - CA Certificate Partition Shared
* `close-notify` - Send close notification when terminate connection
* `crl` - Certificate Revocation Lists (Certificate Revocation Lists file name)
* `crl-shared` - Certificate Revocation Lists Partition Shared
* `forward-proxy-ca-cert` - CA Certificate for forward proxy (SSL forward proxy CA Certificate Name)
* `fp-ca-shared` - CA Certificate Partition Shared
* `forward-proxy-ca-key` - CA Private Key for forward proxy (SSL forward proxy CA Key Name)
* `forward-passphrase` - Password Phrase
* `forward-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `fp-ca-key-shared` - CA Private Key Partition Shared
* `fp-ca-certificate` - CA Certificate for forward proxy (SSL forward proxy CA Certificate Name)
* `fp-ca-key` - CA Private Key for forward proxy (SSL forward proxy CA Key Name)
* `fp-ca-key-passphrase` - Password Phrase
* `fp-ca-key-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `fp-ca-chain-cert` - Chain Certificate (Chain Certificate Name)
* `fp-ca-certificate-shared` - CA Private Key Partition Shared
* `forward-proxy-alt-sign` - Forward proxy alternate signing cert and key
* `fp-alt-cert` - CA Certificate for forward proxy alternate signing (Certificate name)
* `fp-alt-key` - CA Private Key for forward proxy alternate signing (Key name)
* `fp-alt-passphrase` - Password Phrase
* `fp-alt-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `fp-alt-chain-cert` - Chain Certificate (Chain Certificate Name)
* `fp-alt-shared` - Alternate CA Certificate and Private Key Partition Shared
* `forward-proxy-trusted-ca` - Forward proxy trusted CA file (CA file name)
* `fp-trusted-ca-shared` - Trusted CA Certificate Partition Shared
* `forward-proxy-decrypted-dscp` - Apply a DSCP to decrypted and bypassed traffic (DSCP to apply to decrypted traffic)
* `forward-proxy-decrypted-dscp-bypass` - DSCP to apply to bypassed traffic
* `enable-tls-alert-logging` - Enable TLS alert logging
* `alert-type` - 'fatal': Log fatal alerts;
* `forward-proxy-verify-cert-fail-action` - Action taken if certificate verification fails, close the connection by default
* `verify-cert-fail-action` - 'bypass': bypass SSLi processing; 'continue': continue the connection; 'drop': close the connection; 'block': block the connection with a warning page;
* `forward-proxy-cert-revoke-action` - Action taken if a certificate is irreversibly revoked, bypass SSLi processing by default
* `cert-revoke-action` - 'bypass': bypass SSLi processing; 'continue': continue the connection; 'drop': close the connection; 'block': block the connection with a warning page;
* `forward-proxy-no-shared-cipher-action` - Action taken if handshake fails due to no shared ciper, close the connection by default
* `no-shared-cipher-action` - 'bypass': bypass SSLi processing; 'drop': close the connection;
* `forward-proxy-esni-action` - Action taken if receiving encrypted server name indication extension in client hello MSG, bypass the connection by default
* `fp-esni-action` - 'bypass': bypass SSLi processing; 'drop': close the connection;
* `forward-proxy-cert-unknown-action` - Action taken if a certificate revocation status is unknown, bypass SSLi processing by default
* `cert-unknown-action` - 'bypass': bypass SSLi processing; 'continue': continue the connection; 'drop': close the connection; 'block': block the connection with a warning page;
* `forward-proxy-block-message` - Message to be included on the block page (Message, enclose in quotes if spaces are present)
* `cache-persistence-list-name` - Class List Name
* `fp-cert-ext-crldp` - CRL Distribution Point (CRL Distribution Point URI)
* `fp-cert-ext-aia-ocsp` - OCSP (Authority Information Access URI)
* `fp-cert-ext-aia-ca-issuers` - CA Issuers (Authority Information Access URI)
* `notbefore` - notBefore date
* `notbeforeday` - Day
* `notbeforemonth` - Month
* `notbeforeyear` - Year
* `notafter` - notAfter date
* `notafterday` - Day
* `notaftermonth` - Month
* `notafteryear` - Year
* `forward-proxy-hash-persistence-interval` - Set the time interval to save the hash persistence certs (Interval value, in minutes)
* `forward-proxy-ssl-version` - TLS/SSL version, default is TLS1.2 (TLS/SSL version: 31-TLSv1.0, 32-TLSv1.1, 33-TLSv1.2 and 34-TLSv1.3)
* `forward-proxy-ocsp-disable` - Disable ocsp-stapling for forward proxy
* `forward-proxy-crl-disable` - Disable Certificate Revocation List checking for forward proxy
* `forward-proxy-cert-cache-timeout` - Certificate cache timeout, default is 1 hour (seconds, set to 0 for never timeout)
* `forward-proxy-cert-cache-limit` - Certificate cache size limit, default is 524288 (set to 0 for unlimited size)
* `forward-proxy-cert-expiry` - Adjust certificate expiry relative to the time when it is created on the device
* `expire-hours` - Certificate lifetime in hours
* `forward-proxy-enable` - Enable SSL forward proxy
* `handshake-logging-enable` - Enable SSL handshake logging
* `forward-proxy-selfsign-redir` - Redirect connections to pages with self signed certs to a warning page
* `forward-proxy-failsafe-disable` - Disable Failsafe for SSL forward proxy
* `forward-proxy-log-disable` - Disable SSL forward proxy logging
* `fp-cert-fetch-natpool-name` - Specify NAT pool or pool group
* `shared-partition-pool` - Reference a NAT pool or pool group from shared partition
* `fp-cert-fetch-natpool-name-shared` - Specify NAT pool or pool group
* `fp-cert-fetch-natpool-precedence` - Set this NAT pool as higher precedence than other source NAT like configued under template policy
* `fp-cert-fetch-autonat` - 'auto': Configure auto NAT for server certificate fetching;
* `fp-cert-fetch-autonat-precedence` - Set this NAT pool as higher precedence than other source NAT like configued under template policy
* `forward-proxy-no-sni-action` - 'intercept': intercept in no SNI case; 'bypass': bypass in no SNI case; 'reset': reset in no SNI case;
* `case-insensitive` - Case insensitive forward proxy bypass
* `class-list-name` - Class List Name
* `multi-clist-name` - Class List Name
* `user-name-list` - Forward proxy bypass if user-name matches class-list
* `ad-group-list` - Forward proxy bypass if ad-group matches class-list
* `exception-user-name-list` - Exceptions to forward proxy bypass if user-name matches class-list
* `exception-ad-group-list` - Exceptions to forward proxy bypass if ad-group matches class-list
* `exception-sni-cl-name` - Exceptions to forward-proxy-bypass
* `inspect-list-name` - Class List Name
* `inspect-certificate-subject-cl-name` - Forward proxy Inspect if Certificate Subject matches class-list
* `inspect-certificate-issuer-cl-name` - Forward proxy Inspect if Certificate issuer matches class-list
* `inspect-certificate-san-cl-name` - Forward proxy Inspect if Certificate Subject Alternative Name matches class-list
* `contains` - Forward proxy bypass if SNI string contains another string
* `ends-with` - Forward proxy bypass if SNI string ends with another string
* `equals` - Forward proxy bypass if SNI string equals another string
* `starts-with` - Forward proxy bypass if SNI string starts with another string
* `certificate-subject-contains` - Forward proxy bypass if Certificate Subject contains another string
* `bypass-cert-subject-class-list-name` - Class List Name
* `bypass-cert-subject-multi-class-list-name` - Class List Name
* `exception-certificate-subject-cl-name` - Exceptions to forward-proxy-bypass
* `certificate-subject-ends-with` - Forward proxy bypass if Certificate Subject ends with another string
* `certificate-subject-equals` - Forward proxy bypass if Certificate Subject equals another string
* `certificate-subject-starts` - Forward proxy bypass if Certificate Subject starts with another string
* `certificate-issuer-contains` - Forward proxy bypass if Certificate  issuer contains another string (Certificate issuer)
* `bypass-cert-issuer-class-list-name` - Class List Name
* `bypass-cert-issuer-multi-class-list-name` - Class List Name
* `exception-certificate-issuer-cl-name` - Exceptions to forward-proxy-bypass
* `certificate-issuer-ends-with` - Forward proxy bypass if Certificate issuer ends with another string
* `certificate-issuer-equals` - Forward proxy bypass if Certificate issuer equals another string
* `certificate-issuer-starts` - Forward proxy bypass if Certificate issuer starts with another string
* `certificate-san-contains` - Forward proxy bypass if Certificate SAN contains another string
* `bypass-cert-san-class-list-name` - Class List Name
* `bypass-cert-san-multi-class-list-name` - Class List Name
* `exception-certificate-san-cl-name` - Exceptions to forward-proxy-bypass
* `certificate-san-ends-with` - Forward proxy bypass if Certificate SAN ends with another string
* `certificate-san-equals` - Forward proxy bypass if Certificate SAN equals another string
* `certificate-san-starts` - Forward proxy bypass if Certificate SAN starts with another string
* `client-auth-case-insensitive` - Case insensitive forward proxy client auth bypass
* `client-auth-class-list` - Forward proxy client auth bypass if SNI string matches class-list (Class List Name)
* `client-auth-contains` - Forward proxy bypass if SNI string contains another string
* `client-auth-ends-with` - Forward proxy bypass if SNI string ends with another string
* `client-auth-equals` - Forward proxy bypass if SNI string equals another string
* `client-auth-starts-with` - Forward proxy bypass if SNI string starts with another string
* `forward-proxy-cert-not-ready-action` - 'bypass': bypass the connection; 'reset': reset the connection; 'intercept': wait for cert and then inspect the connection;
* `bypass-trustworthy` - Bypass when reputation score is greater than or equal to 81
* `bypass-low-risk` - Bypass when reputation score is greater than or equal to 61
* `bypass-moderate-risk` - Bypass when reputation score is greater than or equal to 41
* `bypass-suspicious` - Bypass when reputation score is greater than or equal to 21
* `bypass-malicious` - Bypass when reputation score is greater than or equal to 1
* `bypass-threshold` - Bypass when reputation score is greater than or equal to the customized score (1-100)
* `exception-trustworthy` - Intercept when reputation score is less than or equal to 100
* `exception-low-risk` - Intercept when reputation score is less than or equal to 80
* `exception-moderate-risk` - Intercept when reputation score is less than or equal to 60
* `exception-suspicious` - Intercept when reputation score is less than or equal to 40
* `exception-malicious` - Intercept when reputation score is less than or equal to 20
* `exception-threshold` - Intercept when reputation score is less than or equal to a customized value (1-100)
* `uncategorized` - Uncategorized URLs
* `real-estate` - Category Real Estate
* `computer-and-internet-security` - Category Computer and Internet Security
* `financial-services` - Category Financial Services
* `business-and-economy` - Category Business and Economy
* `computer-and-internet-info` - Category Computer and Internet Info
* `auctions` - Category Auctions
* `shopping` - Category Shopping
* `cult-and-occult` - Category Cult and Occult
* `travel` - Category Travel
* `drugs` - Category Abused Drugs
* `adult-and-pornography` - Category Adult and Pornography
* `home-and-garden` - Category Home and Garden
* `military` - Category Military
* `social-network` - Category Social Network
* `dead-sites` - Category Dead Sites (db Ops only)
* `stock-advice-and-tools` - Category Stock Advice and Tools
* `training-and-tools` - Category Training and Tools
* `dating` - Category Dating
* `sex-education` - Category Sex Education
* `religion` - Category Religion
* `entertainment-and-arts` - Category Entertainment and Arts
* `personal-sites-and-blogs` - Category Personal sites and Blogs
* `legal` - Category Legal
* `local-information` - Category Local Information
* `streaming-media` - Category Streaming Media
* `job-search` - Category Job Search
* `gambling` - Category Gambling
* `translation` - Category Translation
* `reference-and-research` - Category Reference and Research
* `shareware-and-freeware` - Category Shareware and Freeware
* `peer-to-peer` - Category Peer to Peer
* `marijuana` - Category Marijuana
* `hacking` - Category Hacking
* `games` - Category Games
* `philosophy-and-politics` - Category Philosophy and Political Advocacy
* `weapons` - Category Weapons
* `pay-to-surf` - Category Pay to Surf
* `hunting-and-fishing` - Category Hunting and Fishing
* `society` - Category Society
* `educational-institutions` - Category Educational Institutions
* `online-greeting-cards` - Category Online Greeting cards
* `sports` - Category Sports
* `swimsuits-and-intimate-apparel` - Category Swimsuits and Intimate Apparel
* `questionable` - Category Questionable
* `kids` - Category Kids
* `hate-and-racism` - Category Hate and Racism
* `personal-storage` - Category Personal Storage
* `violence` - Category Violence
* `keyloggers-and-monitoring` - Category Keyloggers and Monitoring
* `search-engines` - Category Search Engines
* `internet-portals` - Category Internet Portals
* `web-advertisements` - Category Web Advertisements
* `cheating` - Category Cheating
* `gross` - Category Gross
* `web-based-email` - Category Web based email
* `malware-sites` - Category Malware Sites
* `phishing-and-other-fraud` - Category Phishing and Other Frauds
* `proxy-avoid-and-anonymizers` - Category Proxy Avoid and Anonymizers
* `spyware-and-adware` - Category Spyware and Adware
* `music` - Category Music
* `government` - Category Government
* `nudity` - Category Nudity
* `news-and-media` - Category News and Media
* `illegal` - Category Illegal
* `cdns` - Category CDNs
* `internet-communications` - Category Internet Communications
* `bot-nets` - Category Bot Nets
* `abortion` - Category Abortion
* `health-and-medicine` - Category Health and Medicine
* `confirmed-spam-sources` - Category Confirmed SPAM Sources
* `spam-urls` - Category SPAM URLs
* `unconfirmed-spam-sources` - Category Unconfirmed SPAM Sources
* `open-http-proxies` - Category Open HTTP Proxies
* `dynamic-comment` - Category Dynamic Comment
* `parked-domains` - Category Parked Domains
* `alcohol-and-tobacco` - Category Alcohol and Tobacco
* `private-ip-addresses` - Category Private IP Addresses
* `image-and-video-search` - Category Image and Video Search
* `fashion-and-beauty` - Category Fashion and Beauty
* `recreation-and-hobbies` - Category Recreation and Hobbies
* `motor-vehicles` - Category Motor Vehicles
* `web-hosting-sites` - Category Web Hosting Sites
* `food-and-dining` - Category Food and Dining
* `nudity-artistic` - Category Nudity join Entertainment and Arts
* `illegal-pornography` - Category Illegal join Adult and Pornography
* `exception-uncategorized` - Uncategorized URLs
* `exception-real-estate` - Category Real Estate
* `exception-computer-and-internet-security` - Category Computer and Internet Security
* `exception-financial-services` - Category Financial Services
* `exception-business-and-economy` - Category Business and Economy
* `exception-computer-and-internet-info` - Category Computer and Internet Info
* `exception-auctions` - Category Auctions
* `exception-shopping` - Category Shopping
* `exception-cult-and-occult` - Category Cult and Occult
* `exception-travel` - Category Travel
* `exception-drugs` - Category Abused Drugs
* `exception-adult-and-pornography` - Category Adult and Pornography
* `exception-home-and-garden` - Category Home and Garden
* `exception-military` - Category Military
* `exception-social-network` - Category Social Network
* `exception-dead-sites` - Category Dead Sites (db Ops only)
* `exception-stock-advice-and-tools` - Category Stock Advice and Tools
* `exception-training-and-tools` - Category Training and Tools
* `exception-dating` - Category Dating
* `exception-sex-education` - Category Sex Education
* `exception-religion` - Category Religion
* `exception-entertainment-and-arts` - Category Entertainment and Arts
* `exception-personal-sites-and-blogs` - Category Personal sites and Blogs
* `exception-legal` - Category Legal
* `exception-local-information` - Category Local Information
* `exception-streaming-media` - Category Streaming Media
* `exception-job-search` - Category Job Search
* `exception-gambling` - Category Gambling
* `exception-translation` - Category Translation
* `exception-reference-and-research` - Category Reference and Research
* `exception-shareware-and-freeware` - Category Shareware and Freeware
* `exception-peer-to-peer` - Category Peer to Peer
* `exception-marijuana` - Category Marijuana
* `exception-hacking` - Category Hacking
* `exception-games` - Category Games
* `exception-philosophy-and-politics` - Category Philosophy and Political Advocacy
* `exception-weapons` - Category Weapons
* `exception-pay-to-surf` - Category Pay to Surf
* `exception-hunting-and-fishing` - Category Hunting and Fishing
* `exception-society` - Category Society
* `exception-educational-institutions` - Category Educational Institutions
* `exception-online-greeting-cards` - Category Online Greeting cards
* `exception-sports` - Category Sports
* `exception-swimsuits-and-intimate-apparel` - Category Swimsuits and Intimate Apparel
* `exception-questionable` - Category Questionable
* `exception-kids` - Category Kids
* `exception-hate-and-racism` - Category Hate and Racism
* `exception-personal-storage` - Category Personal Storage
* `exception-violence` - Category Violence
* `exception-keyloggers-and-monitoring` - Category Keyloggers and Monitoring
* `exception-search-engines` - Category Search Engines
* `exception-internet-portals` - Category Internet Portals
* `exception-web-advertisements` - Category Web Advertisements
* `exception-cheating` - Category Cheating
* `exception-gross` - Category Gross
* `exception-web-based-email` - Category Web based email
* `exception-malware-sites` - Category Malware Sites
* `exception-phishing-and-other-fraud` - Category Phishing and Other Frauds
* `exception-proxy-avoid-and-anonymizers` - Category Proxy Avoid and Anonymizers
* `exception-spyware-and-adware` - Category Spyware and Adware
* `exception-music` - Category Music
* `exception-government` - Category Government
* `exception-nudity` - Category Nudity
* `exception-news-and-media` - Category News and Media
* `exception-illegal` - Category Illegal
* `exception-cdns` - Category CDNs
* `exception-internet-communications` - Category Internet Communications
* `exception-bot-nets` - Category Bot Nets
* `exception-abortion` - Category Abortion
* `exception-health-and-medicine` - Category Health and Medicine
* `exception-confirmed-spam-sources` - Category Confirmed SPAM Sources
* `exception-spam-urls` - Category SPAM URLs
* `exception-unconfirmed-spam-sources` - Category Unconfirmed SPAM Sources
* `exception-open-http-proxies` - Category Open HTTP Proxies
* `exception-dynamic-comment` - Category Dynamic Comment
* `exception-parked-domains` - Category Parked Domains
* `exception-alcohol-and-tobacco` - Category Alcohol and Tobacco
* `exception-private-ip-addresses` - Category Private IP Addresses
* `exception-image-and-video-search` - Category Image and Video Search
* `exception-fashion-and-beauty` - Category Fashion and Beauty
* `exception-recreation-and-hobbies` - Category Recreation and Hobbies
* `exception-motor-vehicles` - Category Motor Vehicles
* `exception-web-hosting-sites` - Category Web Hosting Sites
* `exception-food-and-dining` - Category Food and Dining
* `exception-nudity-artistic` - Category Nudity join Entertainment and Arts
* `exception-illegal-pornography` - Category Illegal join Adult and Pornography
* `require-web-category` - Wait for web category to be resolved before taking bypass decision
* `client-ipv4-list-name` - IPV4 client class-list name
* `client-ipv6-list-name` - IPV6 client class-list name
* `server-ipv4-list-name` - IPV4 server class-list name
* `server-ipv6-list-name` - IPV6 server class-list name
* `exception-client-ipv4-list-name` - IPV4 exception client class-list name
* `exception-client-ipv6-list-name` - IPV6 exception client class-list name
* `exception-server-ipv4-list-name` - IPV4 exception server class-list name
* `exception-server-ipv6-list-name` - IPV6 exception server class-list name
* `local-cert-pin-list-bypass-fail-count` - Set the connection fail count as bypass criteria (Bypass when connection failure count is greater than the criteria (1-65536))
* `central-cert-pin-list` - Forward proxy bypass if SNI string is contained in central updated cert-pinning-candidate list
* `forward-proxy-require-sni-cert-matched` - 'no-match-action-inspect': Inspected if not matched; 'no-match-action-drop': Dropped if not matched;
* `template-cipher` - Cipher Template Name
* `shared-partition-cipher-template` - Reference a cipher template from shared partition
* `template-cipher-shared` - Cipher Template Name
* `template-hsm` - HSM Template (HSM Template Name)
* `hsm-type` - 'thales-embed': Thales embed key; 'thales-hwcrhk': Thales hwcrhk Key;
* `cipher-wo-prio` - 'SSL3_RSA_DES_192_CBC3_SHA': TLS_RSA_WITH_3DES_EDE_CBC_SHA (0x000A); 'SSL3_RSA_RC4_128_MD5': TLS_RSA_WITH_RC4_128_MD5 (0x0004); 'SSL3_RSA_RC4_128_SHA': TLS_RSA_WITH_RC4_128_SHA (0x0005); 'TLS1_RSA_AES_128_SHA': TLS_RSA_WITH_AES_128_CBC_SHA (0x002F); 'TLS1_RSA_AES_256_SHA': TLS_RSA_WITH_AES_256_CBC_SHA (0x0035); 'TLS1_RSA_AES_128_SHA256': TLS_RSA_WITH_AES_128_CBC_SHA256 (0x003C); 'TLS1_RSA_AES_256_SHA256': TLS_RSA_WITH_AES_256_CBC_SHA256 (0x003D); 'TLS1_DHE_RSA_AES_128_GCM_SHA256': TLS_DHE_RSA_WITH_AES_128_GCM_SHA256 (0x009E); 'TLS1_DHE_RSA_AES_128_SHA': TLS_DHE_RSA_WITH_AES_128_CBC_SHA (0x0033); 'TLS1_DHE_RSA_AES_128_SHA256': TLS_DHE_RSA_WITH_AES_128_CBC_SHA256 (0x0067); 'TLS1_DHE_RSA_AES_256_GCM_SHA384': TLS_DHE_RSA_WITH_AES_256_GCM_SHA384 (0x009F); 'TLS1_DHE_RSA_AES_256_SHA': TLS_DHE_RSA_WITH_AES_256_CBC_SHA (0x0039); 'TLS1_DHE_RSA_AES_256_SHA256': TLS_DHE_RSA_WITH_AES_256_CBC_SHA256 (0x006B); 'TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256': TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 (0xC02B); 'TLS1_ECDHE_ECDSA_AES_128_SHA': TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA (0xC009); 'TLS1_ECDHE_ECDSA_AES_128_SHA256': TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256 (0xC023); 'TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384': TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 (0xC02C); 'TLS1_ECDHE_ECDSA_AES_256_SHA': TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA (0xC00A); 'TLS1_ECDHE_RSA_AES_128_GCM_SHA256': TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 (0xC02F); 'TLS1_ECDHE_RSA_AES_128_SHA': TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA (0xC013); 'TLS1_ECDHE_RSA_AES_128_SHA256': TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256 (0xC027); 'TLS1_ECDHE_RSA_AES_256_GCM_SHA384': TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384 (0xC030); 'TLS1_ECDHE_RSA_AES_256_SHA': TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA (0xC014); 'TLS1_RSA_AES_128_GCM_SHA256': TLS_RSA_WITH_AES_128_GCM_SHA256 (0x009C); 'TLS1_RSA_AES_256_GCM_SHA384': TLS_RSA_WITH_AES_256_GCM_SHA384 (0x009D); 'TLS1_ECDHE_RSA_AES_256_SHA384': TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384 (0xC028); 'TLS1_ECDHE_ECDSA_AES_256_SHA384': TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384 (0xC024); 'TLS1_ECDHE_RSA_CHACHA20_POLY1305_SHA256': TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256 (0xCCA8); 'TLS1_ECDHE_ECDSA_CHACHA20_POLY1305_SHA256': TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 (0xCCA9); 'TLS1_DHE_RSA_CHACHA20_POLY1305_SHA256': TLS_DHE_RSA_WITH_CHACHA20_POLY1305_SHA256 (0xCCAA);
* `server-name` - Server name indication in Client hello extension (Server name String)
* `server-cert` - Server Certificate associated to SNI (Server Certificate Name)
* `server-chain` - Server Certificate Chain associated to SNI (Server Certificate Chain Name)
* `server-key` - Server Private Key associated to SNI (Server Private Key Name)
* `server-passphrase` - help Password Phrase
* `server-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `server-name-alternate` - Specific the second certifcate
* `server-shared` - Server Name Partition Shared
* `server-name-regex` - Server name indication in Client hello extension with regular expression (Server name String with regex)
* `server-cert-regex` - Server Certificate associated to SNI regex (Server Certificate Name)
* `server-chain-regex` - Server Certificate Chain associated to SNI regex (Server Certificate Chain Name)
* `server-key-regex` - Server Private Key associated to SNI regex (Server Private Key Name)
* `server-passphrase-regex` - help Password Phrase
* `server-encrypted-regex` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `server-name-regex-alternate` - Specific the second certifcate
* `server-shared-regex` - Server Name Partition Shared
* `server-name-auto-map` - Enable automatic mapping of server name indication in Client hello extension
* `sni-enable-log` - Enable logging of sni-auto-map failures. Disable by default
* `sni-bypass-missing-cert` - Bypass when missing cert/key
* `sni-bypass-expired-cert` - Bypass when certificate expired
* `sni-bypass-explicit-list` - Bypass when matched explicit bypass list (Specify class list name)
* `sni-bypass-enable-log` - Enable logging when bypass event happens, disabled by default
* `direct-client-server-auth` - Let backend server does SSL client authentication directly
* `session-cache-size` - Session Cache Size (Maximum cache size. Default value 0 (Session ID reuse disabled))
* `session-cache-timeout` - Session Cache Timeout (Timeout value, in seconds. Default value 0 (Session cache timeout disabled))
* `session-ticket-disable` - Disable client side session ticket support
* `session-ticket-lifetime` - Session ticket lifetime in seconds from stateless session resumption (Lifetime value in seconds. Default value 0 (Session ticket lifetime is 7200 seconds))
* `ssl-false-start-disable` - disable SSL False Start
* `disable-sslv3` - Reject Client requests for SSL version 3
* `version` - TLS/SSL version, default is the highest number supported (TLS/SSL version: 30-SSLv3.0, 31-TLSv1.0, 32-TLSv1.1, 33-TLSv1.2 and 34-TLSv1.3)
* `dgversion` - Lower TLS/SSL version can be downgraded
* `renegotiation-disable` - Disable SSL renegotiation
* `sslv2-bypass-service-group` - Service Group for Bypass SSLV2 (Service Group Name)
* `authorization` - Specify LDAP server for client SSL authorizaiton
* `authen-name` - Specify authorization LDAP server name
* `ldap-base-dn-from-cert` - Use Subject DN as LDAP search base DN
* `ldap-search-filter` - Specify LDAP search filter
* `auth-sg` - Specify authorization LDAP service group
* `auth-sg-dn` - Use Subject DN as LDAP search base DN
* `auth-sg-filter` - Specify LDAP search filter
* `auth-username-attribute` - Specify attribute name of username for client SSL authorization
* `non-ssl-bypass-service-group` - Service Group for Bypass non-ssl traffic (Service Group Name)
* `non-ssl-bypass-l4session` - Handle the non-ssl session as L4 for performance optimization
* `enable-ssli-ftp-alg` - Enable SSLi FTP over TLS support at which port
* `early-data` - Enable TLS 1.3 early data (0-RTT)
* `no-anti-replay` - Disable anti-replay protection for TLS 1.3 early data (0-RTT data)
* `ja3-enable` - Enable JA3 features
* `ja3-insert-http-header` - Insert the JA3 hash into this request as a HTTP header (HTTP Header Name)
* `ja3-reject-class-list` - Drop request if the JA3 hash matches this class-list (type string-case-insensitive) (Class-List Name)
* `ja3-reject-max-number-per-host` - Drop request if numbers of JA3 of this client address exceeded
* `ja3-ttl` - seconds to keep each JA3 record
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `cert` - Certificate Name
* `key` - Server Private Key (Key Name)
* `passphrase` - Password Phrase
* `key-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `shared` - Server Certificate and Key Partition Shared

