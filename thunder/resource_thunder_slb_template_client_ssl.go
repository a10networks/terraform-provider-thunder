package thunder

//Thunder resource SlbTemplateClientSSL

import (
	"fmt"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateClientSSL() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateClientSSLCreate,
		Update: resourceSlbTemplateClientSSLUpdate,
		Read:   resourceSlbTemplateClientSSLRead,
		Delete: resourceSlbTemplateClientSSLDelete,
		Schema: map[string]*schema.Schema{
			"auth_username_attribute": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_verify_cert_fail_action": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fp_cert_ext_crldp": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_cert_not_ready_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"no_anti_replay": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ocspst_sg_minutes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_failsafe_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enable_tls_alert_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"authorization": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ldap_base_dn_from_cert": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_ocsp_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocspst_ocsp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocspst_srvr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_cipher_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cipher_without_prio_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cipher_wo_prio": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"exception_ad_group_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_require_sni_cert_matched": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_ca_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ec_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ec": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ca_certs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_ocsp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ca_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ca_cert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_ocsp_srvr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_ocsp_sg": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"notbeforeday": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"session_ticket_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auth_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_cert_cache_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fp_cert_fetch_natpool_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"exception_sni_cl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"client_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"client_auth_ends_with_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_auth_ends_with": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"client_auth_equals_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_auth_equals": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"server_name_auto_map": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_pool": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_hsm": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"version": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_alt_sign": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_no_shared_cipher_action": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"certificate_subject_equals_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_subject_equals": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"forward_proxy_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_cert_unknown_action": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"notafteryear": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"key_shared_encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ocspst_srvr_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_crl_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"certificate_issuer_contains_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_issuer_contains": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"forward_proxy_ca_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ssl_false_start_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fp_alt_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"key_alt_encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cert_alternate": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_selfsign_redir": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"certificate_subject_starts_with_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_subject_starts": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"forward_proxy_ssl_version": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auth_sg_filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ocspst_sg_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"exception_web_category": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_motor_vehicles": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_home_and_garden": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_shopping": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_real_estate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_peer_to_peer": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_health_and_medicine": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_society": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_proxy_avoid_and_anonymizers": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_military": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_business_and_economy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_personal_sites_and_blogs": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_legal": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_philosophy_and_politics": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_abortion": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_private_ip_addresses": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_local_information": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_online_greeting_cards": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_drugs": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_phishing_and_other_fraud": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_job_search": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_web_advertisements": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_web_based_email": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_computer_and_internet_info": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_cult_and_occult": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_shareware_and_freeware": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_spam_urls": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_games": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_hunting_and_fishing": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_streaming_media": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_unconfirmed_spam_sources": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_adult_and_pornography": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_kids": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_bot_nets": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_violence": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_entertainment_and_arts": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_gross": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_cdns": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_open_http_proxies": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_internet_portals": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_translation": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_gambling": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_internet_communications": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_marijuana": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_malware_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_web_hosting_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_alcohol_and_tobacco": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_training_and_tools": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_cheating": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_dating": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_weapons": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_hacking": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_dynamic_comment": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_food_and_dining": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_news_and_media": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_social_network": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_recreation_and_hobbies": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_stock_advice_and_tools": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_educational_institutions": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_sex_education": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_financial_services": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_dead_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_uncategorized": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_personal_storage": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_fashion_and_beauty": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_search_engines": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_nudity": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_questionable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_hate_and_racism": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_keyloggers_and_monitoring": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_religion": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_confirmed_spam_sources": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_music": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_spyware_and_adware": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_image_and_video_search": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_sports": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_reference_and_research": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_auctions": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_parked_domains": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_computer_and_internet_security": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_illegal": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_pay_to_surf": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_travel": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_government": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_swimsuits_and_intimate_apparel": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"notbefore": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_cipher": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"client_auth_class_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"bypass_cert_san_multi_class_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_cert_san_multi_class_list_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"server_name_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_cert_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_chain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_name_alternate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_chain_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_passphrase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_name_regex_alternate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_cert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_encrypted_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_name_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_passphrase_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_key_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_shared_regex": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_encrypted": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"session_cache_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"exception_certificate_subject_cl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"exception_web_reputation": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_suspicious": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_moderate_risk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_malicious": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_low_risk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_trustworthy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"non_ssl_bypass_service_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"certificate_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"chain_cert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"passphrase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"key_encrypted": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"fp_cert_fetch_autonat": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"bypass_cert_san_class_list_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_ca_shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"bypass_cert_subject_class_list_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"inspect_certificate_san_cl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_decrypted_dscp_bypass": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_cert_revoke_action": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"session_ticket_lifetime": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"key_alternate": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"certificate_san_contains_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_san_contains": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"fp_cert_ext_aia_ca_issuers": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cache_persistence_list_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"authen_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"exception_certificate_san_cl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_cert_expiry": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ldap_search_filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ocspst_srvr_days": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"key_passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ocspst_sg_hours": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocspst_sg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sni_enable_log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"direct_client_server_auth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"exception_certificate_issuer_cl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"web_reputation": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_malicious": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bypass_moderate_risk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bypass_low_risk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bypass_suspicious": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bypass_trustworthy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bypass_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ocspst_srvr_minutes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_log_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sslv2_bypass_service_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_cert_ext_aia_ocsp": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"certificate_san_ends_with_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_san_ends_with": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"non_ssl_bypass_l4session": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"class_list_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_cert_fetch_natpool_name_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cert_alt_partition_shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"local_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_trusted_ca_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fp_trusted_ca_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"forward_proxy_trusted_ca": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"notaftermonth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fp_cert_fetch_natpool_precedence": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_sslv3": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fp_alt_shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"crl_certs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crl_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"crl": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"fp_cert_fetch_autonat_precedence": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fp_ca_key_shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bypass_cert_issuer_multi_class_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_cert_issuer_multi_class_list_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"client_auth_contains_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_auth_contains": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"certificate_san_starts_with_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_san_starts": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"verify_cert_fail_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"notbeforeyear": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_no_sni_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"web_category": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phishing_and_other_fraud": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"job_search": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"keyloggers_and_monitoring": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"questionable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"abortion": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"malware_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dynamic_comment": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"web_hosting_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"gambling": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"unconfirmed_spam_sources": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"marijuana": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"adult_and_pornography": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alcohol_and_tobacco": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"personal_sites_and_blogs": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"hacking": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"music": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"reference_and_research": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"health_and_medicine": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"legal": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"games": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"military": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bot_nets": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"hunting_and_fishing": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"auctions": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shopping": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dead_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"business_and_economy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"philosophy_and_politics": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"hate_and_racism": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"gross": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"internet_portals": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"recreation_and_hobbies": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"proxy_avoid_and_anonymizers": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"internet_communications": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"computer_and_internet_security": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"personal_storage": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cheating": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"illegal": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"computer_and_internet_info": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"motor_vehicles": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"online_greeting_cards": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"spyware_and_adware": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"stock_advice_and_tools": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"sex_education": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"news_and_media": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"weapons": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"swimsuits_and_intimate_apparel": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"web_advertisements": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"pay_to_surf": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"home_and_garden": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"real_estate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"entertainment_and_arts": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"nudity": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"spam_urls": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cdns": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"confirmed_spam_sources": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"educational_institutions": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cult_and_occult": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"society": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"image_and_video_search": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"government": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"social_network": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"web_based_email": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dating": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"parked_domains": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"kids": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"violence": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"streaming_media": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"sports": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"drugs": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"search_engines": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"open_http_proxies": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"local_information": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uncategorized": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shareware_and_freeware": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"religion": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"peer_to_peer": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"financial_services": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"translation": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"food_and_dining": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"private_ip_addresses": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"travel": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"training_and_tools": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"fashion_and_beauty": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"close_notify": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"client_auth_case_insensitive": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"equals_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"equals": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"early_data": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocsp_stapling": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"certificate_issuer_equals_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_issuer_equals": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ocspst_sg_days": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"expire_hours": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"notbeforemonth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"req_ca_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_cert_req_ca_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"client_certificate_request_ca": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"certificate_san_equals_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_san_equals": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"multi_class_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"multi_clist_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"fp_alt_encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"inspect_certificate_issuer_cl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sslilogging": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"alert_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"handshake_logging_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"certificate_subject_contains_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_subject_contains": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"dgversion": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"no_shared_cipher_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"key_alt_partition_shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocspst_ca_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"case_insensitive": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"inspect_list_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_sg_dn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_name_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ends_with_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ends_with": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"certificate_issuer_starts_with_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_issuer_starts": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"inspect_certificate_subject_cl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ocspst_srvr_hours": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"starts_with_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"starts_with": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"contains_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"contains": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"key_shared_str": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_alt_passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"key_encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"exception_user_name_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"chain_cert_shared_str": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"chain_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cert_unknown_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_cipher_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cert_revoke_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_cert_cache_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"key_shared_passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"require_web_category": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"notafter": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cert_shared_str": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"key_alt_passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"session_cache_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_block_message": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"certificate_subject_ends_with_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_subject_ends_with": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"auth_sg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"notafterday": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fp_alt_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"client_auth_starts_with_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_auth_starts_with": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"hsm_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"renegotiation_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dh_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"bypass_cert_subject_multi_class_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_cert_subject_multi_class_list_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"certificate_issuer_ends_with_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_issuer_ends_with": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"bypass_cert_issuer_class_list_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ad_group_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ssli_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_decrypted_dscp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateClientSSLCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateClientSSL (Inside resourceSlbTemplateClientSSLCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateClientSSL(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateClientSSL --")
		d.SetId(name)
		go_thunder.PostSlbTemplateClientSSL(client.Token, data, client.Host)

		return resourceSlbTemplateClientSSLRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateClientSSLRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateClientSSL (Inside resourceSlbTemplateClientSSLRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTemplateClientSSL(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateClientSSLUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateClientSSL   (Inside resourceSlbTemplateClientSSLUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateClientSSL(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateClientSSL ")
		d.SetId(name)
		go_thunder.PutSlbTemplateClientSSL(client.Token, name, data, client.Host)

		return resourceSlbTemplateClientSSLRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateClientSSLDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateClientSSLDelete) " + name)
		err := go_thunder.DeleteSlbTemplateClientSSL(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateClientSSL(d *schema.ResourceData) go_thunder.ClientSSL {
	var vc go_thunder.ClientSSL
	var c go_thunder.ClientSSLInstance

	BypassCertSubjectMultiClassListCount := d.Get("bypass_cert_subject_multi_class_list.#").(int)
	c.BypassCertSubjectMultiClassListName = make([]go_thunder.BypassCertSubjectMultiClassList, 0, BypassCertSubjectMultiClassListCount)

	for i := 0; i < BypassCertSubjectMultiClassListCount; i++ {
		var obj1 go_thunder.BypassCertSubjectMultiClassList
		prefix := fmt.Sprintf("bypass_cert_subject_multi_class_list.%d.", i)
		obj1.BypassCertSubjectMultiClassListName = d.Get(prefix + "bypass_cert_subject_multi_class_list_name").(string)
		c.BypassCertSubjectMultiClassListName = append(c.BypassCertSubjectMultiClassListName, obj1)
	}

	c.VerifyCertFailAction = d.Get("verify_cert_fail_action").(string)
	c.InspectCertificateIssuerClName = d.Get("inspect_certificate_issuer_cl_name").(string)

	CertificateSanContainsListCount := d.Get("certificate_san_contains_list.#").(int)
	c.CertificateSanContains = make([]go_thunder.CertificateSanContainsList, 0, CertificateSanContainsListCount)

	for i := 0; i < CertificateSanContainsListCount; i++ {
		var obj2 go_thunder.CertificateSanContainsList
		prefix := fmt.Sprintf("certificate_san_contains_list.%d.", i)
		obj2.CertificateSanContains = d.Get(prefix + "certificate_san_contains").(string)
		c.CertificateSanContains = append(c.CertificateSanContains, obj2)
	}

	c.ForwardProxyBlockMessage = d.Get("forward_proxy_block_message").(string)
	c.DirectClientServerAuth = d.Get("direct_client_server_auth").(int)
	c.OcspstSgHours = d.Get("ocspst_sg_hours").(int)
	c.NoSharedCipherAction = d.Get("no_shared_cipher_action").(string)
	c.FpCertFetchAutonat = d.Get("fp_cert_fetch_autonat").(string)

	EqualsListCount := d.Get("equals_list.#").(int)
	c.Equals = make([]go_thunder.EqualsList, 0, EqualsListCount)

	for i := 0; i < EqualsListCount; i++ {
		var obj3 go_thunder.EqualsList
		prefix := fmt.Sprintf("equals_list.%d.", i)
		obj3.Equals = d.Get(prefix + "equals").(string)
		c.Equals = append(c.Equals, obj3)
	}

	c.ExceptionCertificateSubjectClName = d.Get("exception_certificate_subject_cl_name").(string)

	ForwardProxyTrustedCaListsCount := d.Get("forward_proxy_trusted_ca_lists.#").(int)
	c.ForwardProxyTrustedCa = make([]go_thunder.ForwardProxyTrustedCaLists, 0, ForwardProxyTrustedCaListsCount)

	for i := 0; i < ForwardProxyTrustedCaListsCount; i++ {
		var obj4 go_thunder.ForwardProxyTrustedCaLists
		prefix := fmt.Sprintf("forward_proxy_trusted_ca_lists.%d.", i)
		obj4.ForwardProxyTrustedCa = d.Get(prefix + "forward_proxy_trusted_ca").(string)
		c.ForwardProxyTrustedCa = append(c.ForwardProxyTrustedCa, obj4)
	}

	c.TemplateCipherShared = d.Get("template_cipher_shared").(string)
	c.ForwardProxyCaCert = d.Get("forward_proxy_ca_cert").(string)
	c.SslFalseStartDisable = d.Get("ssl_false_start_disable").(int)
	c.Dgversion = d.Get("dgversion").(int)
	c.ClientAuthClassList = d.Get("client_auth_class_list").(string)

	CertificateListCount := d.Get("certificate_list.#").(int)
	c.Passphrase = make([]go_thunder.SlbTemplateCertificateList, 0, CertificateListCount)

	for i := 0; i < CertificateListCount; i++ {
		var obj100 go_thunder.SlbTemplateCertificateList
		prefix100 := fmt.Sprintf("certificate_list.%d.", i)
		obj100.ChainCert = d.Get(prefix100 + "chain_cert").(string)
		obj100.Key = d.Get(prefix100 + "key").(string)
		obj100.Cert = d.Get(prefix100 + "cert").(string)
		obj100.Passphrase = d.Get(prefix100 + "passphrase").(string)
		obj100.KeyEncrypted = d.Get(prefix100 + "key_encrypted").(string)
		c.Passphrase = append(c.Passphrase, obj100)
	}

	c.KeyEncrypted = d.Get("key_encrypted").(string)
	c.Notafteryear = d.Get("notafteryear").(int)
	c.ForwardProxyAltSign = d.Get("forward_proxy_alt_sign").(int)
	c.TemplateHsm = d.Get("template_hsm").(string)
	c.ForwardPassphrase = d.Get("forward_passphrase").(string)
	c.ExceptionCertificateIssuerClName = d.Get("exception_certificate_issuer_cl_name").(string)

	ContainsListCount := d.Get("contains_list.#").(int)
	c.Contains = make([]go_thunder.ContainsList, 0, ContainsListCount)

	for i := 0; i < ContainsListCount; i++ {
		var obj5 go_thunder.ContainsList
		prefix := fmt.Sprintf("contains_list.%d.", i)
		obj5.Contains = d.Get(prefix + "contains").(string)
		c.Contains = append(c.Contains, obj5)
	}

	c.ForwardProxyCaKey = d.Get("forward_proxy_ca_key").(string)
	c.Notbefore = d.Get("notbefore").(int)

	EndsWithListCount := d.Get("ends_with_list.#").(int)
	c.EndsWith = make([]go_thunder.EndsWithList, 0, EndsWithListCount)

	for i := 0; i < EndsWithListCount; i++ {
		var obj6 go_thunder.EndsWithList
		prefix := fmt.Sprintf("ends_with_list.%d.", i)
		obj6.EndsWith = d.Get(prefix + "ends_with").(string)
		c.EndsWith = append(c.EndsWith, obj6)
	}

	c.BypassCertSubjectClassListName = d.Get("bypass_cert_subject_class_list_name").(string)
	c.Notafter = d.Get("notafter").(int)
	c.ClassListName = d.Get("class_list_name").(string)
	c.OcspstOcsp = d.Get("ocspst_ocsp").(int)
	c.Notbeforeday = d.Get("notbeforeday").(int)
	c.KeyAltPassphrase = d.Get("key_alt_passphrase").(string)
	c.ForwardProxySslVersion = d.Get("forward_proxy_ssl_version").(int)

	CaCertsCount := d.Get("ca_certs.#").(int)
	c.ClientOcspSrvr = make([]go_thunder.CaCerts2, 0, CaCertsCount)

	for i := 0; i < CaCertsCount; i++ {
		var obj7 go_thunder.CaCerts2
		prefix := fmt.Sprintf("ca_certs.%d.", i)
		obj7.ClientOcspSrvr = d.Get(prefix + "client_ocsp_srvr").(string)
		obj7.CaCert = d.Get(prefix + "ca_cert").(string)
		obj7.ClientOcspSg = d.Get(prefix + "client_ocsp_sg").(string)
		obj7.ClientOcsp = d.Get(prefix + "client_ocsp").(int)
		obj7.CaShared = d.Get(prefix + "ca_shared").(int)
		c.ClientOcspSrvr = append(c.ClientOcspSrvr, obj7)
	}

	c.ForwardProxyCrlDisable = d.Get("forward_proxy_crl_disable").(int)

	ClientAuthContainsListCount := d.Get("client_auth_contains_list.#").(int)
	c.ClientAuthContains = make([]go_thunder.ClientAuthContainsList, 0, ClientAuthContainsListCount)

	for i := 0; i < ClientAuthContainsListCount; i++ {
		var obj8 go_thunder.ClientAuthContainsList
		prefix := fmt.Sprintf("client_auth_contains_list.%d.", i)
		obj8.ClientAuthContains = d.Get(prefix + "client_auth_contains").(string)
		c.ClientAuthContains = append(c.ClientAuthContains, obj8)
	}

	CertificateSubjectContainsListCount := d.Get("certificate_subject_contains_list.#").(int)
	c.CertificateSubjectContains = make([]go_thunder.CertificateSubjectContainsList, 0, CertificateSubjectContainsListCount)

	for i := 0; i < CertificateSubjectContainsListCount; i++ {
		var obj9 go_thunder.CertificateSubjectContainsList
		prefix := fmt.Sprintf("certificate_subject_contains_list.%d.", i)
		obj9.CertificateSubjectContains = d.Get(prefix + "certificate_subject_contains").(string)
		c.CertificateSubjectContains = append(c.CertificateSubjectContains, obj9)
	}

	c.Name = d.Get("name").(string)
	c.ForwardProxyCertRevokeAction = d.Get("forward_proxy_cert_revoke_action").(int)
	c.FpCertExtAiaOcsp = d.Get("fp_cert_ext_aia_ocsp").(string)

	ReqCaListsCount := d.Get("req_ca_lists.#").(int)
	c.ClientCertificateRequestCA = make([]go_thunder.ReqCaLists, 0, ReqCaListsCount)

	for i := 0; i < ReqCaListsCount; i++ {
		var obj10 go_thunder.ReqCaLists
		prefix := fmt.Sprintf("req_ca_lists.%d.", i)
		obj10.ClientCertificateRequestCA = d.Get(prefix + "client_certificate_request_ca").(string)
		c.ClientCertificateRequestCA = append(c.ClientCertificateRequestCA, obj10)
	}

	c.UserTag = d.Get("user_tag").(string)
	c.CertUnknownAction = d.Get("cert_unknown_action").(string)

	// SamplingEnableCount := d.Get("sampling_enable.#").(int)
	// c.Counters1 = make([]go_thunder.SamplingEnable5, 0, SamplingEnableCount)

	// for i := 0; i < SamplingEnableCount; i++ {
	// 	var obj11 go_thunder.SamplingEnable5
	// 	prefix := fmt.Sprintf("sampling_enable.%d.", i)
	// 	obj11.Counters1 = d.Get(prefix + "counters1").(string)
	// 	c.Counters1 = append(c.Counters1, obj11)
	// }

	c.RenegotiationDisable = d.Get("renegotiation_disable").(int)
	c.ExceptionAdGroupList = d.Get("exception_ad_group_list").(string)
	c.KeyAlternate = d.Get("key_alternate").(string)
	c.FpAltKey = d.Get("fp_alt_key").(string)
	c.ServerNameAutoMap = d.Get("server_name_auto_map").(int)
	c.DisableSslv3 = d.Get("disable_sslv3").(int)

	BypassCertIssuerMultiClassListCount := d.Get("bypass_cert_issuer_multi_class_list.#").(int)
	c.BypassCertIssuerMultiClassListName = make([]go_thunder.BypassCertIssuerMultiClassList, 0, BypassCertIssuerMultiClassListCount)

	for i := 0; i < BypassCertIssuerMultiClassListCount; i++ {
		var obj12 go_thunder.BypassCertIssuerMultiClassList
		prefix := fmt.Sprintf("bypass_cert_issuer_multi_class_list.%d.", i)
		obj12.BypassCertIssuerMultiClassListName = d.Get(prefix + "bypass_cert_issuer_multi_class_list_name").(string)
		c.BypassCertIssuerMultiClassListName = append(c.BypassCertIssuerMultiClassListName, obj12)
	}

	ClientAuthEqualsListCount := d.Get("client_auth_equals_list.#").(int)
	c.ClientAuthEquals = make([]go_thunder.ClientAuthEqualsList, 0, ClientAuthEqualsListCount)

	for i := 0; i < ClientAuthEqualsListCount; i++ {
		var obj13 go_thunder.ClientAuthEqualsList
		prefix := fmt.Sprintf("client_auth_equals_list.%d.", i)
		obj13.ClientAuthEquals = d.Get(prefix + "client_auth_equals").(string)
		c.ClientAuthEquals = append(c.ClientAuthEquals, obj13)
	}

	c.ForwardProxyNoSniAction = d.Get("forward_proxy_no_sni_action").(string)

	CertificateIssuerEqualsListCount := d.Get("certificate_issuer_equals_list.#").(int)
	c.CertificateIssuerEquals = make([]go_thunder.CertificateIssuerEqualsList, 0, CertificateIssuerEqualsListCount)

	for i := 0; i < CertificateIssuerEqualsListCount; i++ {
		var obj14 go_thunder.CertificateIssuerEqualsList
		prefix := fmt.Sprintf("certificate_issuer_equals_list.%d.", i)
		obj14.CertificateIssuerEquals = d.Get(prefix + "certificate_issuer_equals").(string)
		c.CertificateIssuerEquals = append(c.CertificateIssuerEquals, obj14)
	}

	c.FpAltPassphrase = d.Get("fp_alt_passphrase").(string)

	CertificateSubjectStartsWithListCount := d.Get("certificate_subject_starts_with_list.#").(int)
	c.CertificateSubjectStarts = make([]go_thunder.CertificateSubjectStartsWithList, 0, CertificateSubjectStartsWithListCount)

	for i := 0; i < CertificateSubjectStartsWithListCount; i++ {
		var obj15 go_thunder.CertificateSubjectStartsWithList
		prefix := fmt.Sprintf("certificate_subject_starts_with_list.%d.", i)
		obj15.CertificateSubjectStarts = d.Get(prefix + "certificate_subject_starts").(string)
		c.CertificateSubjectStarts = append(c.CertificateSubjectStarts, obj15)
	}

	CertificateSanEndsWithListCount := d.Get("certificate_san_ends_with_list.#").(int)
	c.CertificateSanEndsWith = make([]go_thunder.CertificateSanEndsWithList, 0, CertificateSanEndsWithListCount)

	for i := 0; i < CertificateSanEndsWithListCount; i++ {
		var obj16 go_thunder.CertificateSanEndsWithList
		prefix := fmt.Sprintf("certificate_san_ends_with_list.%d.", i)
		obj16.CertificateSanEndsWith = d.Get(prefix + "certificate_san_ends_with").(string)
		c.CertificateSanEndsWith = append(c.CertificateSanEndsWith, obj16)
	}

	c.ForwardProxyCertCacheTimeout = d.Get("forward_proxy_cert_cache_timeout").(int)
	c.FpCertFetchNatpoolNameShared = d.Get("fp_cert_fetch_natpool_name_shared").(string)

	CrlCertsCount := d.Get("crl_certs.#").(int)
	c.CrlShared = make([]go_thunder.CrlCerts2, 0, CrlCertsCount)

	for i := 0; i < CrlCertsCount; i++ {
		var obj17 go_thunder.CrlCerts2
		prefix := fmt.Sprintf("crl_certs.%d.", i)
		obj17.CrlShared = d.Get(prefix + "crl_shared").(int)
		obj17.Crl = d.Get(prefix + "crl").(string)
		c.CrlShared = append(c.CrlShared, obj17)
	}

	c.Notafterday = d.Get("notafterday").(int)
	c.OcspstSrvrHours = d.Get("ocspst_srvr_hours").(int)
	c.LocalLogging = d.Get("local_logging").(int)
	c.FpCertFetchAutonatPrecedence = d.Get("fp_cert_fetch_autonat_precedence").(int)
	//c.CertStr = d.Get("cert_str").(string)
	c.CertSharedStr = d.Get("cert_shared_str").(string)
	c.CertRevokeAction = d.Get("cert_revoke_action").(string)
	c.Version = d.Get("version").(int)

	MultiClassListCount := d.Get("multi_class_list.#").(int)
	c.MultiClistName = make([]go_thunder.MultiClassList, 0, MultiClassListCount)

	for i := 0; i < MultiClassListCount; i++ {
		var obj18 go_thunder.MultiClassList
		prefix := fmt.Sprintf("multi_class_list.%d.", i)
		obj18.MultiClistName = d.Get(prefix + "multi_clist_name").(string)
		c.MultiClistName = append(c.MultiClistName, obj18)
	}

	c.UserNameList = d.Get("user_name_list").(string)
	c.SessionTicketLifetime = d.Get("session_ticket_lifetime").(int)

	CertificateIssuerEndsWithListCount := d.Get("certificate_issuer_ends_with_list.#").(int)
	c.CertificateIssuerEndsWith = make([]go_thunder.CertificateIssuerEndsWithList, 0, CertificateIssuerEndsWithListCount)

	for i := 0; i < CertificateIssuerEndsWithListCount; i++ {
		var obj19 go_thunder.CertificateIssuerEndsWithList
		prefix := fmt.Sprintf("certificate_issuer_ends_with_list.%d.", i)
		obj19.CertificateIssuerEndsWith = d.Get(prefix + "certificate_issuer_ends_with").(string)
		c.CertificateIssuerEndsWith = append(c.CertificateIssuerEndsWith, obj19)
	}

	c.SsliLogging = d.Get("ssli_logging").(int)
	c.SessionCacheSize = d.Get("session_cache_size").(int)
	c.HandshakeLoggingEnable = d.Get("handshake_logging_enable").(int)
	c.NonSslBypassServiceGroup = d.Get("non_ssl_bypass_service_group").(string)
	c.ForwardProxyFailsafeDisable = d.Get("forward_proxy_failsafe_disable").(int)
	c.SessionCacheTimeout = d.Get("session_cache_timeout").(int)
	c.Sslv2BypassServiceGroup = d.Get("sslv2_bypass_service_group").(string)
	c.ForwardProxyDecryptedDscp = d.Get("forward_proxy_decrypted_dscp").(int)
	c.AuthSg = d.Get("auth_sg").(string)
	c.OcspstCaCert = d.Get("ocspst_ca_cert").(string)
	c.ForwardProxySelfsignRedir = d.Get("forward_proxy_selfsign_redir").(int)
	c.AuthSgDn = d.Get("auth_sg_dn").(int)
	c.HsmType = d.Get("hsm_type").(string)
	c.ForwardProxyLogDisable = d.Get("forward_proxy_log_disable").(int)
	c.FpAltEncrypted = d.Get("fp_alt_encrypted").(string)
	c.InspectCertificateSanClName = d.Get("inspect_certificate_san_cl_name").(string)

	var obj20 go_thunder.WebCategory
	prefix := "web_category.0."
	obj20.PhilosophyAndPolitics = d.Get(prefix + "philosophy_and_politics").(int)
	obj20.StockAdviceAndTools = d.Get(prefix + "stock_advice_and_tools").(int)
	obj20.NewsAndMedia = d.Get(prefix + "news_and_media").(int)
	obj20.BusinessAndEconomy = d.Get(prefix + "business_and_economy").(int)
	obj20.PeerToPeer = d.Get(prefix + "peer_to_peer").(int)
	obj20.PhishingAndOtherFraud = d.Get(prefix + "phishing_and_other_fraud").(int)
	obj20.Nudity = d.Get(prefix + "nudity").(int)
	obj20.Weapons = d.Get(prefix + "weapons").(int)
	obj20.HealthAndMedicine = d.Get(prefix + "health_and_medicine").(int)
	obj20.Marijuana = d.Get(prefix + "marijuana").(int)
	obj20.HomeAndGarden = d.Get(prefix + "home_and_garden").(int)
	obj20.CultAndOccult = d.Get(prefix + "cult_and_occult").(int)
	obj20.Society = d.Get(prefix + "society").(int)
	obj20.PersonalStorage = d.Get(prefix + "personal_storage").(int)
	obj20.ComputerAndInternetSecurity = d.Get(prefix + "computer_and_internet_security").(int)
	obj20.FoodAndDining = d.Get(prefix + "food_and_dining").(int)
	obj20.MotorVehicles = d.Get(prefix + "motor_vehicles").(int)
	obj20.SwimsuitsAndIntimateApparel = d.Get(prefix + "swimsuits_and_intimate_apparel").(int)
	obj20.DeadSites = d.Get(prefix + "dead_sites").(int)
	obj20.Translation = d.Get(prefix + "translation").(int)
	obj20.ProxyAvoidAndAnonymizers = d.Get(prefix + "proxy_avoid_and_anonymizers").(int)
	obj20.FinancialServices = d.Get(prefix + "financial_services").(int)
	obj20.Gross = d.Get(prefix + "gross").(int)
	obj20.Cheating = d.Get(prefix + "cheating").(int)
	obj20.EntertainmentAndArts = d.Get(prefix + "entertainment_and_arts").(int)
	obj20.SexEducation = d.Get(prefix + "sex_education").(int)
	obj20.Illegal = d.Get(prefix + "illegal").(int)
	obj20.Travel = d.Get(prefix + "travel").(int)
	obj20.Cdns = d.Get(prefix + "cdns").(int)
	obj20.LocalInformation = d.Get(prefix + "local_information").(int)
	obj20.Legal = d.Get(prefix + "legal").(int)
	obj20.Sports = d.Get(prefix + "sports").(int)
	obj20.BotNets = d.Get(prefix + "bot_nets").(int)
	obj20.Religion = d.Get(prefix + "religion").(int)
	obj20.PrivateIPAddresses = d.Get(prefix + "private_ip_addresses").(int)
	obj20.Music = d.Get(prefix + "music").(int)
	obj20.HateAndRacism = d.Get(prefix + "hate_and_racism").(int)
	obj20.OpenHTTPProxies = d.Get(prefix + "open_http_proxies").(int)
	obj20.InternetCommunications = d.Get(prefix + "internet_communications").(int)
	obj20.SharewareAndFreeware = d.Get(prefix + "shareware_and_freeware").(int)
	obj20.Dating = d.Get(prefix + "dating").(int)
	obj20.SpywareAndAdware = d.Get(prefix + "spyware_and_adware").(int)
	obj20.Uncategorized = d.Get(prefix + "uncategorized").(int)
	obj20.Questionable = d.Get(prefix + "questionable").(int)
	obj20.ReferenceAndResearch = d.Get(prefix + "reference_and_research").(int)
	obj20.WebAdvertisements = d.Get(prefix + "web_advertisements").(int)
	obj20.StreamingMedia = d.Get(prefix + "streaming_media").(int)
	obj20.SocialNetwork = d.Get(prefix + "social_network").(int)
	obj20.Government = d.Get(prefix + "government").(int)
	obj20.Drugs = d.Get(prefix + "drugs").(int)
	obj20.WebHostingSites = d.Get(prefix + "web_hosting_sites").(int)
	obj20.MalwareSites = d.Get(prefix + "malware_sites").(int)
	obj20.PayToSurf = d.Get(prefix + "pay_to_surf").(int)
	obj20.SpamUrls = d.Get(prefix + "spam_urls").(int)
	obj20.Kids = d.Get(prefix + "kids").(int)
	obj20.Gambling = d.Get(prefix + "gambling").(int)
	obj20.OnlineGreetingCards = d.Get(prefix + "online_greeting_cards").(int)
	obj20.ConfirmedSpamSources = d.Get(prefix + "confirmed_spam_sources").(int)
	obj20.ImageAndVideoSearch = d.Get(prefix + "image_and_video_search").(int)
	obj20.EducationalInstitutions = d.Get(prefix + "educational_institutions").(int)
	obj20.KeyloggersAndMonitoring = d.Get(prefix + "keyloggers_and_monitoring").(int)
	obj20.HuntingAndFishing = d.Get(prefix + "hunting_and_fishing").(int)
	obj20.SearchEngines = d.Get(prefix + "search_engines").(int)
	obj20.FashionAndBeauty = d.Get(prefix + "fashion_and_beauty").(int)
	obj20.DynamicComment = d.Get(prefix + "dynamic_comment").(int)
	obj20.ComputerAndInternetInfo = d.Get(prefix + "computer_and_internet_info").(int)
	obj20.RealEstate = d.Get(prefix + "real_estate").(int)
	obj20.InternetPortals = d.Get(prefix + "internet_portals").(int)
	obj20.Shopping = d.Get(prefix + "shopping").(int)
	obj20.Violence = d.Get(prefix + "violence").(int)
	obj20.Abortion = d.Get(prefix + "abortion").(int)
	obj20.TrainingAndTools = d.Get(prefix + "training_and_tools").(int)
	obj20.WebBasedEmail = d.Get(prefix + "web_based_email").(int)
	obj20.PersonalSitesAndBlogs = d.Get(prefix + "personal_sites_and_blogs").(int)
	obj20.UnconfirmedSpamSources = d.Get(prefix + "unconfirmed_spam_sources").(int)
	obj20.Games = d.Get(prefix + "games").(int)
	obj20.ParkedDomains = d.Get(prefix + "parked_domains").(int)
	obj20.Auctions = d.Get(prefix + "auctions").(int)
	obj20.JobSearch = d.Get(prefix + "job_search").(int)
	obj20.RecreationAndHobbies = d.Get(prefix + "recreation_and_hobbies").(int)
	obj20.Hacking = d.Get(prefix + "hacking").(int)
	obj20.AlcoholAndTobacco = d.Get(prefix + "alcohol_and_tobacco").(int)
	obj20.AdultAndPornography = d.Get(prefix + "adult_and_pornography").(int)
	obj20.Military = d.Get(prefix + "military").(int)
	c.PhilosophyAndPolitics = obj20

	CertificateSanEqualsListCount := d.Get("certificate_san_equals_list.#").(int)
	c.CertificateSanEquals = make([]go_thunder.CertificateSanEqualsList, 0, CertificateSanEqualsListCount)

	for i := 0; i < CertificateSanEqualsListCount; i++ {
		var obj21 go_thunder.CertificateSanEqualsList
		prefix := fmt.Sprintf("certificate_san_equals_list.%d.", i)
		obj21.CertificateSanEquals = d.Get(prefix + "certificate_san_equals").(string)
		c.CertificateSanEquals = append(c.CertificateSanEquals, obj21)
	}

	c.TemplateCipher = d.Get("template_cipher").(string)
	c.Notbeforemonth = d.Get("notbeforemonth").(int)
	c.BypassCertSanClassListName = d.Get("bypass_cert_san_class_list_name").(string)
	c.ChainCert = d.Get("chain_cert").(string)
	c.ForwardProxyCertUnknownAction = d.Get("forward_proxy_cert_unknown_action").(int)
	c.ExceptionCertificateSanClName = d.Get("exception_certificate_san_cl_name").(string)
	c.OcspstSg = d.Get("ocspst_sg").(string)
	c.KeyAltEncrypted = d.Get("key_alt_encrypted").(string)
	c.FpCertExtAiaCaIssuers = d.Get("fp_cert_ext_aia_ca_issuers").(string)
	c.AuthenName = d.Get("authen_name").(string)
	c.ExpireHours = d.Get("expire_hours").(int)
	c.ClientAuthCaseInsensitive = d.Get("client_auth_case_insensitive").(int)
	c.OcspStapling = d.Get("ocsp_stapling").(int)
	c.Notbeforeyear = d.Get("notbeforeyear").(int)
	c.ForwardEncrypted = d.Get("forward_encrypted").(string)
	c.SniEnableLog = d.Get("sni_enable_log").(int)
	c.KeySharedStr = d.Get("key_shared_str").(string)
	c.Notaftermonth = d.Get("notaftermonth").(int)
	c.CachePersistenceListName = d.Get("cache_persistence_list_name").(string)
	c.OcspstSgTimeout = d.Get("ocspst_sg_timeout").(int)
	c.KeyPassphrase = d.Get("key_passphrase").(string)
	c.OcspstSrvr = d.Get("ocspst_srvr").(string)
	c.OcspstSrvrMinutes = d.Get("ocspst_srvr_minutes").(int)

	CertificateIssuerContainsListCount := d.Get("certificate_issuer_contains_list.#").(int)
	c.CertificateIssuerContains = make([]go_thunder.CertificateIssuerContainsList, 0, CertificateIssuerContainsListCount)

	for i := 0; i < CertificateIssuerContainsListCount; i++ {
		var obj22 go_thunder.CertificateIssuerContainsList
		prefix := fmt.Sprintf("certificate_issuer_contains_list.%d.", i)
		obj22.CertificateIssuerContains = d.Get(prefix + "certificate_issuer_contains").(string)
		c.CertificateIssuerContains = append(c.CertificateIssuerContains, obj22)
	}

	c.RequireWebCategory = d.Get("require_web_category").(int)

	BypassCertSanMultiClassListCount := d.Get("bypass_cert_san_multi_class_list.#").(int)
	c.BypassCertSanMultiClassListName = make([]go_thunder.BypassCertSanMultiClassList, 0, BypassCertSanMultiClassListCount)

	for i := 0; i < BypassCertSanMultiClassListCount; i++ {
		var obj23 go_thunder.BypassCertSanMultiClassList
		prefix := fmt.Sprintf("bypass_cert_san_multi_class_list.%d.", i)
		obj23.BypassCertSanMultiClassListName = d.Get(prefix + "bypass_cert_san_multi_class_list_name").(string)
		c.BypassCertSanMultiClassListName = append(c.BypassCertSanMultiClassListName, obj23)
	}

	ClientAuthStartsWithListCount := d.Get("client_auth_starts_with_list.#").(int)
	c.ClientAuthStartsWith = make([]go_thunder.ClientAuthStartsWithList, 0, ClientAuthStartsWithListCount)

	for i := 0; i < ClientAuthStartsWithListCount; i++ {
		var obj24 go_thunder.ClientAuthStartsWithList
		prefix := fmt.Sprintf("client_auth_starts_with_list.%d.", i)
		obj24.ClientAuthStartsWith = d.Get(prefix + "client_auth_starts_with").(string)
		c.ClientAuthStartsWith = append(c.ClientAuthStartsWith, obj24)
	}

	CertificateSubjectEndsWithListCount := d.Get("certificate_subject_ends_with_list.#").(int)
	c.CertificateSubjectEndsWith = make([]go_thunder.CertificateSubjectEndsWithList, 0, CertificateSubjectEndsWithListCount)

	for i := 0; i < CertificateSubjectEndsWithListCount; i++ {
		var obj25 go_thunder.CertificateSubjectEndsWithList
		prefix := fmt.Sprintf("certificate_subject_ends_with_list.%d.", i)
		obj25.CertificateSubjectEndsWith = d.Get(prefix + "certificate_subject_ends_with").(string)
		c.CertificateSubjectEndsWith = append(c.CertificateSubjectEndsWith, obj25)
	}

	c.Authorization = d.Get("authorization").(int)
	c.ForwardProxyVerifyCertFailAction = d.Get("forward_proxy_verify_cert_fail_action").(int)
	c.OcspstSrvrDays = d.Get("ocspst_srvr_days").(int)

	EcListCount := d.Get("ec_list.#").(int)
	c.Ec = make([]go_thunder.EcList2, 0, EcListCount)

	for i := 0; i < EcListCount; i++ {
		var obj26 go_thunder.EcList2
		prefix := fmt.Sprintf("ec_list.%d.", i)
		obj26.Ec = d.Get(prefix + "ec").(string)
		c.Ec = append(c.Ec, obj26)
	}

	c.ForwardProxyDecryptedDscpBypass = d.Get("forward_proxy_decrypted_dscp_bypass").(int)
	c.AlertType = d.Get("alert_type").(string)
	c.ForwardProxyCertNotReadyAction = d.Get("forward_proxy_cert_not_ready_action").(string)

	ServerNameListCount := d.Get("server_name_list.#").(int)
	c.ServerShared = make([]go_thunder.ServerNameList, 0, ServerNameListCount)

	for i := 0; i < ServerNameListCount; i++ {
		var obj27 go_thunder.ServerNameList
		prefix := fmt.Sprintf("server_name_list.%d.", i)
		obj27.ServerShared = d.Get(prefix + "server_shared").(int)
		obj27.ServerPassphraseRegex = d.Get(prefix + "server_passphrase_regex").(string)
		obj27.ServerChain = d.Get(prefix + "server_chain").(string)
		obj27.ServerCertRegex = d.Get(prefix + "server_cert_regex").(string)
		obj27.ServerName = d.Get(prefix + "server_name").(string)
		obj27.ServerKeyRegex = d.Get(prefix + "server_key_regex").(string)
		obj27.ServerNameRegexAlternate = d.Get(prefix + "server_name_regex_alternate").(int)
		obj27.ServerEncryptedRegex = d.Get(prefix + "server_encrypted_regex").(string)
		obj27.ServerSharedRegex = d.Get(prefix + "server_shared_regex").(int)
		obj27.ServerNameRegex = d.Get(prefix + "server_name_regex").(string)
		obj27.ServerPassphrase = d.Get(prefix + "server_passphrase").(string)
		obj27.ServerKey = d.Get(prefix + "server_key").(string)
		obj27.ServerChainRegex = d.Get(prefix + "server_chain_regex").(string)
		obj27.ServerNameAlternate = d.Get(prefix + "server_name_alternate").(int)
		obj27.ServerEncrypted = d.Get(prefix + "server_encrypted").(string)
		obj27.ServerCert = d.Get(prefix + "server_cert").(string)
		c.ServerShared = append(c.ServerShared, obj27)
	}

	c.BypassCertIssuerClassListName = d.Get("bypass_cert_issuer_class_list_name").(string)
	c.FpCertExtCrldp = d.Get("fp_cert_ext_crldp").(string)
	c.SharedPartitionCipherTemplate = d.Get("shared_partition_cipher_template").(int)
	c.FpCertFetchNatpoolPrecedence = d.Get("fp_cert_fetch_natpool_precedence").(int)
	c.CertAlternate = d.Get("cert_alternate").(string)
	c.ForwardProxyCertCacheLimit = d.Get("forward_proxy_cert_cache_limit").(int)
	c.NonSslBypassL4Session = d.Get("non_ssl_bypass_l4session").(int)

	CertificateIssuerStartsWithListCount := d.Get("certificate_issuer_starts_with_list.#").(int)
	c.CertificateIssuerStarts = make([]go_thunder.CertificateIssuerStartsWithList, 0, CertificateIssuerStartsWithListCount)

	for i := 0; i < CertificateIssuerStartsWithListCount; i++ {
		var obj28 go_thunder.CertificateIssuerStartsWithList
		prefix := fmt.Sprintf("certificate_issuer_starts_with_list.%d.", i)
		obj28.CertificateIssuerStarts = d.Get(prefix + "certificate_issuer_starts").(string)
		c.CertificateIssuerStarts = append(c.CertificateIssuerStarts, obj28)
	}

	CertificateSanStartsWithListCount := d.Get("certificate_san_starts_with_list.#").(int)
	c.CertificateSanStarts = make([]go_thunder.CertificateSanStartsWithList, 0, CertificateSanStartsWithListCount)

	for i := 0; i < CertificateSanStartsWithListCount; i++ {
		var obj29 go_thunder.CertificateSanStartsWithList
		prefix := fmt.Sprintf("certificate_san_starts_with_list.%d.", i)
		obj29.CertificateSanStarts = d.Get(prefix + "certificate_san_starts").(string)
		c.CertificateSanStarts = append(c.CertificateSanStarts, obj29)
	}

	ClientAuthEndsWithListCount := d.Get("client_auth_ends_with_list.#").(int)
	c.ClientAuthEndsWith = make([]go_thunder.ClientAuthEndsWithList, 0, ClientAuthEndsWithListCount)

	for i := 0; i < ClientAuthEndsWithListCount; i++ {
		var obj30 go_thunder.ClientAuthEndsWithList
		prefix := fmt.Sprintf("client_auth_ends_with_list.%d.", i)
		obj30.ClientAuthEndsWith = d.Get(prefix + "client_auth_ends_with").(string)
		c.ClientAuthEndsWith = append(c.ClientAuthEndsWith, obj30)
	}

	c.CloseNotify = d.Get("close_notify").(int)
	c.ForwardProxyNoSharedCipherAction = d.Get("forward_proxy_no_shared_cipher_action").(int)
	c.ForwardProxyOcspDisable = d.Get("forward_proxy_ocsp_disable").(int)
	c.Sslilogging = d.Get("sslilogging").(string)
	c.AuthUsername = d.Get("auth_username").(string)
	c.ExceptionUserNameList = d.Get("exception_user_name_list").(string)
	c.OcspstSgDays = d.Get("ocspst_sg_days").(int)
	//c.KeyStr = d.Get("key_str").(string)
	c.InspectListName = d.Get("inspect_list_name").(string)
	c.AuthUsernameAttribute = d.Get("auth_username_attribute").(string)
	c.FpCertFetchNatpoolName = d.Get("fp_cert_fetch_natpool_name").(string)
	c.ExceptionSniClName = d.Get("exception_sni_cl_name").(string)
	c.InspectCertificateSubjectClName = d.Get("inspect_certificate_subject_cl_name").(string)
	c.LdapBaseDnFromCert = d.Get("ldap_base_dn_from_cert").(int)
	c.AdGroupList = d.Get("ad_group_list").(string)
	c.ClientCertificate = d.Get("client_certificate").(string)
	c.ForwardProxyCertExpiry = d.Get("forward_proxy_cert_expiry").(int)
	c.ForwardProxyEnable = d.Get("forward_proxy_enable").(int)
	c.SharedPartitionPool = d.Get("shared_partition_pool").(int)
	c.LdapSearchFilter = d.Get("ldap_search_filter").(string)
	c.KeySharedEncrypted = d.Get("key_shared_encrypted").(string)
	c.AuthSgFilter = d.Get("auth_sg_filter").(string)
	c.OcspstSrvrTimeout = d.Get("ocspst_srvr_timeout").(int)

	CertificateSubjectEqualsListCount := d.Get("certificate_subject_equals_list.#").(int)
	c.CertificateSubjectEquals = make([]go_thunder.CertificateSubjectEqualsList, 0, CertificateSubjectEqualsListCount)

	for i := 0; i < CertificateSubjectEqualsListCount; i++ {
		var obj31 go_thunder.CertificateSubjectEqualsList
		prefix := fmt.Sprintf("certificate_subject_equals_list.%d.", i)
		obj31.CertificateSubjectEquals = d.Get(prefix + "certificate_subject_equals").(string)
		c.CertificateSubjectEquals = append(c.CertificateSubjectEquals, obj31)
	}

	c.ChainCertSharedStr = d.Get("chain_cert_shared_str").(string)
	c.EnableTLSAlertLogging = d.Get("enable_tls_alert_logging").(int)
	c.DhType = d.Get("dh_type").(string)
	c.FpAltCert = d.Get("fp_alt_cert").(string)
	c.Key = d.Get("key").(string)
	c.Cert = d.Get("cert").(string)
	c.CaseInsensitive = d.Get("case_insensitive").(int)

	CipherWithoutPrioListCount := d.Get("cipher_without_prio_list.#").(int)
	c.CipherWoPrio = make([]go_thunder.CipherWithoutPrioList2, 0, CipherWithoutPrioListCount)

	for i := 0; i < CipherWithoutPrioListCount; i++ {
		var obj32 go_thunder.CipherWithoutPrioList2
		prefix := fmt.Sprintf("cipher_without_prio_list.%d.", i)
		obj32.CipherWoPrio = d.Get(prefix + "cipher_wo_prio").(string)
		c.CipherWoPrio = append(c.CipherWoPrio, obj32)
	}

	c.OcspstSgMinutes = d.Get("ocspst_sg_minutes").(int)

	StartsWithListCount := d.Get("starts_with_list.#").(int)
	c.StartsWith = make([]go_thunder.StartsWithList, 0, StartsWithListCount)

	for i := 0; i < StartsWithListCount; i++ {
		var obj33 go_thunder.StartsWithList
		prefix := fmt.Sprintf("starts_with_list.%d.", i)
		obj33.StartsWith = d.Get(prefix + "starts_with").(string)
		c.StartsWith = append(c.StartsWith, obj33)
	}

	c.KeySharedPassphrase = d.Get("key_shared_passphrase").(string)

	vc.UUID = c
	return vc
}
