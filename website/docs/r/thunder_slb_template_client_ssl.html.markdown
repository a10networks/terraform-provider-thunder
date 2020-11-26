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
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_client_ssl" "Slb_Template_Client_Ssl_Test" {
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
cert = "string"
cert_shared_str = "string"
cert_alternate = "string"
cert_alt_partition_shared = 0
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
        client_certificate_Request_CA =  "string" 
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
forward_encrypted = "Unknown Type: encrypted"
fp_ca_key_shared = 0
forward_proxy_alt_sign = 0
fp_alt_cert = "string"
fp_alt_key = "string"
fp_alt_passphrase = "string"
fp_alt_encrypted = "Unknown Type: encrypted"
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
        }
require_web_category = 0
forward_proxy_require_sni_cert_matched = "string"
key = "string"
key_passphrase = "string"
key_encrypted = "Unknown Type: encrypted"
key_shared_str = "string"
key_shared_passphrase = "string"
key_shared_encrypted = "Unknown Type: encrypted"
key_alternate = "string"
key_alt_passphrase = "string"
key_alt_encrypted = "Unknown Type: encrypted"
key_alt_partition_shared = 0
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
        server_encrypted =  "Unknown Type: encrypted" 
        server_name_alternate =  0 
        server_shared =  0 
        server_name_regex =  "string" 
        server_cert_regex =  "string" 
        server_chain_regex =  "string" 
        server_key_regex =  "string" 
        server_passphrase_regex =  "string" 
        server_encrypted_regex =  "Unknown Type: encrypted" 
        server_name_regex_alternate =  0 
        server_shared_regex =  0 
        }
server_name_auto_map = 0
sni_enable_log = 0
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
early_data = 0
no_anti_replay = 0
uuid = "string"
user_tag = "string"
certificate-list {   
        cert =  "string" 
        key =  "string" 
        passphrase =  "string" 
        key_encrypted =  "Unknown Type: encrypted" 
        chain_cert =  "string" 
        shared =  0 
        uuid =  "string" 
        }
 
}


```

## Argument Reference

* `ad_group_list` - Forward proxy bypass if ad-group matches class-list
* `alert_type` - ‘fatal’: Log fatal alerts;
* `auth_sg` - Specify authorization LDAP service group
* `auth_sg_dn` - Use Subject DN as LDAP search base DN
* `auth_sg_filter` - Specify LDAP search filter
* `auth_username` - Specify the Username Field in the Client Certificate(If multi-fields are specificed, prior one has higher priority)
* `auth_username_attribute` - Specify attribute name of username for client SSL authorization
* `authen_name` - Specify authorization LDAP server name
* `authorization` - Specify LDAP server for client SSL authorizaiton
* `bypass_cert_issuer_class_list_name` - Class List Name
* `bypass_cert_san_class_list_name` - Class List Name
* `bypass_cert_subject_class_list_name` - Class List Name
* `cache_persistence_list_name` - Class List Name
* `case_insensitive` - Case insensitive forward proxy bypass
* `cert_alternate` - Specify the second certificate (Certificate Name)
* `cert_revoke_action` - ‘bypass’: bypass SSLi processing; ‘continue’: continue the connection; ‘drop’: close the connection;
* `cert_shared_str` - Certificate Name
* `cert_str` - Certificate Name
* `cert_unknown_action` - ‘bypass’: bypass SSLi processing; ‘continue’: continue the connection; ‘drop’: close the connection;
* `chain_cert` - Chain Certificate (Chain Certificate Name)
* `chain_cert_shared_str` - Chain Certificate Name
* `class_list_name` - Class List Name
* `client_auth_case_insensitive` - Case insensitive forward proxy client auth bypass
* `client_auth_class_list` - Forward proxy client auth bypass if SNI string matches class-list (Class List Name)
* `client_certificate` - ‘Ignore’: Don’t request client certificate; ‘Require’: Require client certificate; ‘Request’: Request client certificate;
* `close_notify` - Send close notification when terminate connection
* `dgversion` - Lower TLS/SSL version can be downgraded
* `dh_type` - ‘1024’: 1024; ‘1024-dsa’: 1024-dsa; ‘2048’: 2048;
* `direct_client_server_auth` - Let backend server does SSL client authentication directly
* `disable_sslv3` - Reject Client requests for SSL version 3
* `enable_tls_alert_logging` - Enable TLS alert logging
* `exception_ad_group_list` - Exceptions to forward proxy bypass if ad-group matches class-list
* `exception_certificate_issuer_cl_name` - Exceptions to forward-proxy-bypass
* `exception_certificate_san_cl_name` - Exceptions to forward-proxy-bypass
* `exception_certificate_subject_cl_name` - Exceptions to forward-proxy-bypass
* `exception_sni_cl_name` - Exceptions to forward-proxy-bypass
* `exception_user_name_list` - Exceptions to forward proxy bypass if user-name matches class-list
* `expire_hours` - Certificate lifetime in hours
* `forward_encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `forward_passphrase` - Password Phrase
* `forward_proxy_alt_sign` - Forward proxy alternate signing cert and key
* `forward_proxy_block_message` - Message to be included on the block page (Message, enclose in quotes if spaces are present)
* `forward_proxy_ca_cert` - CA Certificate for forward proxy (SSL forward proxy CA Certificate Name)
* `forward_proxy_ca_key` - CA Private Key for forward proxy (SSL forward proxy CA Key Name)
* `forward_proxy_cert_cache_limit` - Certificate cache size limit, default is 524288 (set to 0 for unlimited size)
* `forward_proxy_cert_cache_timeout` - Certificate cache timeout, default is 1 hour (seconds, set to 0 for never timeout)
* `forward_proxy_cert_expiry` - Adjust certificate expiry relative to the time when it is created on the device
* `forward_proxy_cert_not_ready_action` - ‘bypass’: bypass the connection; ‘reset’: reset the connection;
* `forward_proxy_cert_revoke_action` - Action taken if a certificate is irreversibly revoked, bypass SSLi processing by default
* `forward_proxy_cert_unknown_action` - Action taken if a certificate revocation status is unknown, bypass SSLi processing by default
* `forward_proxy_crl_disable` - Disable Certificate Revocation List checking for forward proxy
* `forward_proxy_decrypted_dscp` - Apply a DSCP to decrypted and bypassed traffic (DSCP to apply to decrypted traffic)
* `forward_proxy_decrypted_dscp_bypass` - DSCP to apply to bypassed traffic
* `forward_proxy_enable` - Enable SSL forward proxy
* `forward_proxy_failsafe_disable` - Disable Failsafe for SSL forward proxy
* `forward_proxy_log_disable` - Disable SSL forward proxy logging
* `forward_proxy_no_shared_cipher_action` - Action taken if handshake fails due to no shared ciper, close the connection by default
* `forward_proxy_no_sni_action` - ‘intercept’: intercept in no SNI case; ‘bypass’: bypass in no SNI case; ‘reset’: reset in no SNI case;
* `forward_proxy_ocsp_disable` - Disable ocsp-stapling for forward proxy
* `forward_proxy_selfsign_redir` - Redirect connections to pages with self signed certs to a warning page
* `forward_proxy_ssl_version` - TLS/SSL version, default is TLS1.2 (TLS/SSL version: 31-TLSv1.0, 32-TLSv1.1 and 33-TLSv1.2)
* `forward_proxy_verify_cert_fail_action` - Action taken if certificate verification fails, close the connection by default
* `fp_alt_cert` - CA Certificate for forward proxy alternate signing (Certificate name)
* `fp_alt_encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `fp_alt_key` - CA Private Key for forward proxy alternate signing (Key name)
* `fp_alt_passphrase` - Password Phrase
* `fp_cert_ext_aia_ca_issuers` - CA Issuers (Authority Information Access URI)
* `fp_cert_ext_aia_ocsp` - OCSP (Authority Information Access URI)
* `fp_cert_ext_crldp` - CRL Distribution Point (CRL Distribution Point URI)
* `fp_cert_fetch_autonat` - ‘auto’: Configure auto NAT for server certificate fetching;
* `fp_cert_fetch_autonat_precedence` - Set this NAT pool as higher precedence than other source NAT like configued under template policy
* `fp_cert_fetch_natpool_name` - Specify NAT pool or pool group
* `fp_cert_fetch_natpool_name_shared` - Specify NAT pool or pool group
* `fp_cert_fetch_natpool_precedence` - Set this NAT pool as higher precedence than other source NAT like configued under template policy
* `handshake_logging_enable` - Enable SSL handshake logging
* `hsm_type` - ‘thales-embed’: Thales embed key; ‘thales-hwcrhk’: Thales hwcrhk Key;
* `inspect_certificate_issuer_cl_name` - Forward proxy Inspect if Certificate issuer matches class-list
* `inspect_certificate_san_cl_name` - Forward proxy Inspect if Certificate Subject Alternative Name matches class-list
* `inspect_certificate_subject_cl_name` - Forward proxy Inspect if Certificate Subject matches class-list
* `inspect_list_name` - Class List Name
* `key_alt_encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `key_alt_passphrase` - Password Phrase
* `key_alternate` - Specify the second private key (Key Name)
* `key_encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `key_passphrase` - Password Phrase
* `key_shared_encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `key_shared_passphrase` - Password Phrase
* `key_shared_str` - Key Name
* `key_str` - Key Name
* `ldap_base_dn_from_cert` - Use Subject DN as LDAP search base DN
* `ldap_search_filter` - Specify LDAP search filter
* `local_logging` - Enable local logging
* `name` - Client SSL Template Name
* `no_shared_cipher_action` - ‘bypass’: bypass SSLi processing; ‘drop’: close the connection;
* `non_ssl_bypass_l4session` - Handle the non-ssl session as L4 for performance optimization
* `non_ssl_bypass_service_group` - Service Group for Bypass non-ssl traffic (Service Group Name)
* `notafter` - notAfter date
* `notafterday` - Day
* `notaftermonth` - Month
* `notafteryear` - Year
* `notbefore` - notBefore date
* `notbeforeday` - Day
* `notbeforemonth` - Month
* `notbeforeyear` - Year
* `ocsp_stapling` - Config OCSP stapling support
* `ocspst_ca_cert` - CA certificate
* `ocspst_ocsp` - Specify OCSP Authentication
* `ocspst_sg` - Specify authentication service group
* `ocspst_sg_days` - Specify update period, in days
* `ocspst_sg_hours` - Specify update period, in hours
* `ocspst_sg_minutes` - Specify update period, in minutes
* `ocspst_sg_timeout` - Specify retry timeout (Default is 30 mins)
* `ocspst_srvr` - Specify OCSP authentication server
* `ocspst_srvr_days` - Specify update period, in days
* `ocspst_srvr_hours` - Specify update period, in hours
* `ocspst_srvr_minutes` - Specify update period, in minutes
* `ocspst_srvr_timeout` - Specify retry timeout (Default is 30 mins)
* `renegotiation_disable` - Disable SSL renegotiation
* `require_web_category` - Wait for web category to be resolved before taking bypass decision
* `server_name_auto_map` - Enable automatic mapping of server name indication in Client hello extension
* `session_cache_size` - Session Cache Size (Specify 0 to disable Session ID reuse.)
* `session_cache_timeout` - Session Cache Timeout (Timeout value, in seconds)
* `session_ticket_lifetime` - Session ticket lieftime in seconds from stateless session resumption (Lifetime value in seconds)
* `shared_partition_cipher_template` - Reference a cipher template from shared partition
* `shared_partition_pool` - Reference a NAT pool or pool group from shared partition
* `sni_enable_log` - Enable logging of sni-auto-map failures. Disable by default
* `ssl_false_start_disable` - disable SSL False Start
* `ssli_logging` - SSLi logging level, default is error logging only
* `sslilogging` - ‘disable’: Disable all logging; ‘all’: enable all logging(error, info);
* `sslv2_bypass_service_group` - Service Group for Bypass SSLV2 (Service Group Name)
* `template_cipher` - Cipher Template (Cipher Config Name)
* `template_cipher_shared` - Cipher Template Name
* `template_hsm` - HSM Template (HSM Template Name)
* `user_name_list` - Forward proxy bypass if user-name matches class-list
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `verify_cert_fail_action` - ‘bypass’: bypass SSLi processing; ‘continue’: continue the connection; ‘drop’: close the connection;
* `version` - TLS/SSL version, default is the highest number supported (TLS/SSL version: 30-SSLv3.0, 31-TLSv1.0, 32-TLSv1.1 and 33-TLSv1.2)
* `bypass_cert_subject_multi_class_list_name` - Class List Name
* `certificate_san_contains` - Forward proxy bypass if Certificate SAN contains another string
* `equals` - Forward proxy bypass if SNI string equals another string
* `forward_proxy_trusted_ca` - Forward proxy trusted CA file (CA file name)
* `ec` - ‘secp256r1’: X9_62_prime256v1; ‘secp384r1’: secp384r1;
* `contains` - Forward proxy bypass if SNI string contains another string
* `ends_with` - Forward proxy bypass if SNI string ends with another string
* `ca_cert` - CA Certificate (CA Certificate Name)
* `ca_shared` - CA Certificate Partition Shared
* `client_ocsp` - Specify ocsp authentication server(s) for client certificate verification
* `client_ocsp_sg` - Specify service-group (Service group name)
* `client_ocsp_srvr` - Specify authentication server
* `client_auth_contains` - Forward proxy bypass if SNI string contains another string
* `certificate_subject_contains` - Forward proxy bypass if Certificate Subject contains another string
* `client_certificate_Request_CA` - Send CA lists in certificate request (CA Certificate Name)
* `certificate_subject_starts` - Forward proxy bypass if Certificate Subject starts with another string
* `counters1` - ‘all’: all; ‘real-estate’: real estate category; ‘computer-and-internet-security’: computer and internet security category; ‘financial-services’: financial services category; ‘business-and-economy’: business and economy category; ‘computer-and-internet-info’: computer and internet info category; ‘auctions’: auctions category; ‘shopping’: shopping category; ‘cult-and-occult’: cult and occult category; ‘travel’: travel category; ‘drugs’: drugs category; ‘adult-and-pornography’: adult and pornography category; ‘home-and-garden’: home and garden category; ‘military’: military category; ‘social-network’: social network category; ‘dead-sites’: dead sites category; ‘stock-advice-and-tools’: stock advice and tools category; ‘training-and-tools’: training and tools category; ‘dating’: dating category; ‘sex-education’: sex education category; ‘religion’: religion category; ‘entertainment-and-arts’: entertainment and arts category; ‘personal-sites-and-blogs’: personal sites and blogs category; ‘legal’: legal category; ‘local-information’: local information category; ‘streaming-media’: streaming media category; ‘job-search’: job search category; ‘gambling’: gambling category; ‘translation’: translation category; ‘reference-and-research’: reference and research category; ‘shareware-and-freeware’: shareware and freeware category; ‘peer-to-peer’: peer to peer category; ‘marijuana’: marijuana category; ‘hacking’: hacking category; ‘games’: games category; ‘philosophy-and-politics’: philosophy and politics category; ‘weapons’: weapons category; ‘pay-to-surf’: pay to surf category; ‘hunting-and-fishing’: hunting and fishing category; ‘society’: society category; ‘educational-institutions’: educational institutions category; ‘online-greeting-cards’: online greeting cards category; ‘sports’: sports category; ‘swimsuits-and-intimate-apparel’: swimsuits and intimate apparel category; ‘questionable’: questionable category; ‘kids’: kids category; ‘hate-and-racism’: hate and racism category; ‘personal-storage’: personal storage category; ‘violence’: violence category; ‘keyloggers-and-monitoring’: keyloggers and monitoring category; ‘search-engines’: search engines category; ‘internet-portals’: internet portals category; ‘web-advertisements’: web advertisements category; ‘cheating’: cheating category; ‘gross’: gross category; ‘web-based-email’: web based email category; ‘malware-sites’: malware sites category; ‘phishing-and-other-fraud’: phishing and other fraud category; ‘proxy-avoid-and-anonymizers’: proxy avoid and anonymizers category; ‘spyware-and-adware’: spyware and adware category; ‘music’: music category; ‘government’: government category; ‘nudity’: nudity category; ‘news-and-media’: news and media category; ‘illegal’: illegal category; ‘CDNs’: content delivery networks category; ‘internet-communications’: internet communications category; ‘bot-nets’: bot nets category; ‘abortion’: abortion category; ‘health-and-medicine’: health and medicine category; ‘confirmed-SPAM-sources’: confirmed SPAM sources category; ‘SPAM-URLs’: SPAM URLs category; ‘unconfirmed-SPAM-sources’: unconfirmed SPAM sources category; ‘open-HTTP-proxies’: open HTTP proxies category; ‘dynamic-comment’: dynamic comment category; ‘parked-domains’: parked domains category; ‘alcohol-and-tobacco’: alcohol and tobacco category; ‘private-IP-addresses’: private IP addresses category; ‘image-and-video-search’: image and video search category; ‘fashion-and-beauty’: fashion and beauty category; ‘recreation-and-hobbies’: recreation and hobbies category; ‘motor-vehicles’: motor vehicles category; ‘web-hosting-sites’: web hosting sites category; ‘food-and-dining’: food and dining category; ‘uncategorised’: uncategorised; ‘other-category’: other category;
* `bypass_cert_issuer_multi_class_list_name` - Class List Name
* `client_auth_equals` - Forward proxy bypass if SNI string equals another string
* `certificate_issuer_equals` - Forward proxy bypass if Certificate issuer equals another string
* `certificate_san_ends_with` - Forward proxy bypass if Certificate SAN ends with another string
* `crl` - Certificate Revocation Lists (Certificate Revocation Lists file name)
* `crl_shared` - Certificate Revocation Lists Partition Shared
* `multi_clist_name` - Class List Name
* `certificate_issuer_ends_with` - Forward proxy bypass if Certificate issuer ends with another string
* `abortion` - Category Abortion
* `adult_and_pornography` - Category Adult and Pornography
* `alcohol_and_tobacco` - Category Alcohol and Tobacco
* `auctions` - Category Auctions
* `bot_nets` - Category Bot Nets
* `business_and_economy` - Category Business and Economy
* `cdns` - Category CDNs
* `cheating` - Category Cheating
* `computer_and_internet_info` - Category Computer and Internet Info
* `computer_and_internet_security` - Category Computer and Internet Security
* `confirmed_spam_sources` - Category Confirmed SPAM Sources
* `cult_and_occult` - Category Cult and Occult
* `dating` - Category Dating
* `dead_sites` - Category Dead Sites (db Ops only)
* `drugs` - Category Abused Drugs
* `dynamic_comment` - Category Dynamic Comment
* `educational_institutions` - Category Educational Institutions
* `entertainment_and_arts` - Category Entertainment and Arts
* `fashion_and_beauty` - Category Fashion and Beauty
* `financial_services` - Category Financial Services
* `food_and_dining` - Category Food and Dining
* `gambling` - Category Gambling
* `games` - Category Games
* `government` - Category Government
* `gross` - Category Gross
* `hacking` - Category Hacking
* `hate_and_racism` - Category Hate and Racism
* `health_and_medicine` - Category Health and Medicine
* `home_and_garden` - Category Home and Garden
* `hunting_and_fishing` - Category Hunting and Fishing
* `illegal` - Category Illegal
* `image_and_video_search` - Category Image and Video Search
* `internet_communications` - Category Internet Communications
* `internet_portals` - Category Internet Portals
* `job_search` - Category Job Search
* `keyloggers_and_monitoring` - Category Keyloggers and Monitoring
* `kids` - Category Kids
* `legal` - Category Legal
* `local_information` - Category Local Information
* `malware_sites` - Category Malware Sites
* `marijuana` - Category Marijuana
* `military` - Category Military
* `motor_vehicles` - Category Motor Vehicles
* `music` - Category Music
* `news_and_media` - Category News and Media
* `nudity` - Category Nudity
* `online_greeting_cards` - Category Online Greeting cards
* `open_http_proxies` - Category Open HTTP Proxies
* `parked_domains` - Category Parked Domains
* `pay_to_surf` - Category Pay to Surf
* `peer_to_peer` - Category Peer to Peer
* `personal_sites_and_blogs` - Category Personal sites and Blogs
* `personal_storage` - Category Personal Storage
* `philosophy_and_politics` - Category Philosophy and Political Advocacy
* `phishing_and_other_fraud` - Category Phishing and Other Frauds
* `private_ip_addresses` - Category Private IP Addresses
* `proxy_avoid_and_anonymizers` - Category Proxy Avoid and Anonymizers
* `questionable` - Category Questionable
* `real_estate` - Category Real Estate
* `recreation_and_hobbies` - Category Recreation and Hobbies
* `reference_and_research` - Category Reference and Research
* `religion` - Category Religion
* `search_engines` - Category Search Engines
* `sex_education` - Category Sex Education
* `shareware_and_freeware` - Category Shareware and Freeware
* `shopping` - Category Shopping
* `social_network` - Category Social Network
* `society` - Category Society
* `spam_urls` - Category SPAM URLs
* `sports` - Category Sports
* `spyware_and_adware` - Category Spyware and Adware
* `stock_advice_and_tools` - Category Stock Advice and Tools
* `streaming_media` - Category Streaming Media
* `swimsuits_and_intimate_apparel` - Category Swimsuits and Intimate Apparel
* `training_and_tools` - Category Training and Tools
* `translation` - Category Translation
* `travel` - Category Travel
* `uncategorized` - Uncategorized URLs
* `unconfirmed_spam_sources` - Category Unconfirmed SPAM Sources
* `violence` - Category Violence
* `weapons` - Category Weapons
* `web_advertisements` - Category Web Advertisements
* `web_based_email` - Category Web based email
* `web_hosting_sites` - Category Web Hosting Sites
* `certificate_san_equals` - Forward proxy bypass if Certificate SAN equals another string
* `certificate_issuer_contains` - Forward proxy bypass if Certificate  issuer contains another string (Certificate issuer)
* `client_auth_starts_with` - Forward proxy bypass if SNI string starts with another string
* `certificate_subject_ends_with` - Forward proxy bypass if Certificate Subject ends with another string
* `bypass_cert_san_multi_class_list_name` - Class List Name
* `server_cert` - Server Certificate associated to SNI (Server Certificate Name)
* `server_cert_regex` - Server Certificate associated to SNI regex (Server Certificate Name)
* `server_chain` - Server Certificate Chain associated to SNI (Server Certificate Chain Name)
* `server_chain_regex` - Server Certificate Chain associated to SNI regex (Server Certificate Chain Name)
* `server_encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `server_encrypted_regex` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `server_key` - Server Private Key associated to SNI (Server Private Key Name)
* `server_key_regex` - Server Private Key associated to SNI regex (Server Private Key Name)
* `server_name` - Server name indication in Client hello extension (Server name String)
* `server_name_alternate` - Specific the second certifcate
* `server_name_regex` - Server name indication in Client hello extension with regular expression (Server name String with regex)
* `server_name_regex_alternate` - Specific the second certifcate
* `server_passphrase` - help Password Phrase
* `server_passphrase_regex` - help Password Phrase
* `server_shared` - Server Name Partition Shared
* `server_shared_regex` - Server Name Partition Shared
* `certificate_issuer_starts` - Forward proxy bypass if Certificate issuer starts with another string
* `certificate_san_starts` - Forward proxy bypass if Certificate SAN starts with another string
* `client_auth_ends_with` - Forward proxy bypass if SNI string ends with another string
* `certificate_subject_equals` - Forward proxy bypass if Certificate Subject equals another string
* `cipher_wo_prio` - ‘SSL3_RSA_DES_192_CBC3_SHA’: SSL3_RSA_DES_192_CBC3_SHA; ‘SSL3_RSA_RC4_128_MD5’: SSL3_RSA_RC4_128_MD5; ‘SSL3_RSA_RC4_128_SHA’: SSL3_RSA_RC4_128_SHA; ‘TLS1_RSA_AES_128_SHA’: TLS1_RSA_AES_128_SHA; ‘TLS1_RSA_AES_256_SHA’: TLS1_RSA_AES_256_SHA; ‘TLS1_RSA_AES_128_SHA256’: TLS1_RSA_AES_128_SHA256; ‘TLS1_RSA_AES_256_SHA256’: TLS1_RSA_AES_256_SHA256; ‘TLS1_DHE_RSA_AES_128_GCM_SHA256’: TLS1_DHE_RSA_AES_128_GCM_SHA256; ‘TLS1_DHE_RSA_AES_128_SHA’: TLS1_DHE_RSA_AES_128_SHA; ‘TLS1_DHE_RSA_AES_128_SHA256’: TLS1_DHE_RSA_AES_128_SHA256; ‘TLS1_DHE_RSA_AES_256_GCM_SHA384’: TLS1_DHE_RSA_AES_256_GCM_SHA384; ‘TLS1_DHE_RSA_AES_256_SHA’: TLS1_DHE_RSA_AES_256_SHA; ‘TLS1_DHE_RSA_AES_256_SHA256’: TLS1_DHE_RSA_AES_256_SHA256; ‘TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256’: TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256; ‘TLS1_ECDHE_ECDSA_AES_128_SHA’: TLS1_ECDHE_ECDSA_AES_128_SHA; ‘TLS1_ECDHE_ECDSA_AES_128_SHA256’: TLS1_ECDHE_ECDSA_AES_128_SHA256; ‘TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384’: TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384; ‘TLS1_ECDHE_ECDSA_AES_256_SHA’: TLS1_ECDHE_ECDSA_AES_256_SHA; ‘TLS1_ECDHE_RSA_AES_128_GCM_SHA256’: TLS1_ECDHE_RSA_AES_128_GCM_SHA256; ‘TLS1_ECDHE_RSA_AES_128_SHA’: TLS1_ECDHE_RSA_AES_128_SHA; ‘TLS1_ECDHE_RSA_AES_128_SHA256’: TLS1_ECDHE_RSA_AES_128_SHA256; ‘TLS1_ECDHE_RSA_AES_256_GCM_SHA384’: TLS1_ECDHE_RSA_AES_256_GCM_SHA384; ‘TLS1_ECDHE_RSA_AES_256_SHA’: TLS1_ECDHE_RSA_AES_256_SHA; ‘TLS1_RSA_AES_128_GCM_SHA256’: TLS1_RSA_AES_128_GCM_SHA256; ‘TLS1_RSA_AES_256_GCM_SHA384’: TLS1_RSA_AES_256_GCM_SHA384;
* `starts_with` - Forward proxy bypass if SNI string starts with another string
