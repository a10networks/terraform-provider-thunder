package thunder

//Thunder resource SlbTemplateClientSsl

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateClientSsl() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateClientSslCreate,
		UpdateContext: resourceSlbTemplateClientSslUpdate,
		ReadContext:   resourceSlbTemplateClientSslRead,
		DeleteContext: resourceSlbTemplateClientSslDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ca_certs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_cert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ca_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"client_ocsp": {
							Type:        schema.TypeInt,
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
			"chain_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"chain_cert_shared_str": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dh_type": {
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
			"local_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocsp_stapling": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocspst_ca_cert": {
				Type:        schema.TypeString,
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
			"ocspst_srvr_days": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocspst_srvr_hours": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocspst_srvr_minutes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocspst_srvr_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocspst_sg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ocspst_sg_days": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocspst_sg_hours": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocspst_sg_minutes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocspst_sg_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ssli_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sslilogging": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"client_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"req_ca_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_certificate_request_ca": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_cert_req_ca_shared": {
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
			"crl_certs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crl": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"crl_shared": {
							Type:        schema.TypeInt,
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
			"fp_ca_shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_ca_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_ca_key_shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fp_ca_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_ca_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_ca_key_passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_ca_chain_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_ca_certificate_shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_alt_sign": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fp_alt_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_alt_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_alt_passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_alt_chain_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_alt_shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_trusted_ca_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"forward_proxy_trusted_ca": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"fp_trusted_ca_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"forward_proxy_decrypted_dscp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_decrypted_dscp_bypass": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enable_tls_alert_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"alert_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_verify_cert_fail_action": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"verify_cert_fail_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_cert_revoke_action": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cert_revoke_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_no_shared_cipher_action": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"no_shared_cipher_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_esni_action": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fp_esni_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_cert_unknown_action": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cert_unknown_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_block_message": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cache_persistence_list_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_cert_ext_crldp": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_cert_ext_aia_ocsp": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_cert_ext_aia_ca_issuers": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"notbefore": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"notbeforeday": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"notbeforemonth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"notbeforeyear": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"notafter": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"notafterday": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"notaftermonth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"notafteryear": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_hash_persistence_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_ssl_version": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_ocsp_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_crl_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_cert_cache_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_cert_cache_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_cert_expiry": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"expire_hours": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"handshake_logging_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_selfsign_redir": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_failsafe_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_log_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fp_cert_fetch_natpool_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_pool": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fp_cert_fetch_natpool_name_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_cert_fetch_natpool_precedence": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fp_cert_fetch_autonat": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fp_cert_fetch_autonat_precedence": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_no_sni_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"case_insensitive": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"class_list_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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
			"user_name_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ad_group_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"exception_user_name_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"exception_ad_group_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"exception_sni_cl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"inspect_list_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"inspect_certificate_subject_cl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"inspect_certificate_issuer_cl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"inspect_certificate_san_cl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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
			"bypass_cert_subject_class_list_name": {
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
			"exception_certificate_subject_cl_name": {
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
			"bypass_cert_issuer_class_list_name": {
				Type:        schema.TypeString,
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
			"exception_certificate_issuer_cl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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
			"bypass_cert_san_class_list_name": {
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
			"exception_certificate_san_cl_name": {
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
			"client_auth_case_insensitive": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"client_auth_class_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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
			"forward_proxy_cert_not_ready_action": {
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
						"bypass_trustworthy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bypass_low_risk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bypass_moderate_risk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bypass_suspicious": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bypass_malicious": {
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
			"exception_web_reputation": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_trustworthy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_low_risk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_moderate_risk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_suspicious": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_malicious": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"web_category": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uncategorized": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"real_estate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"computer_and_internet_security": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"financial_services": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"business_and_economy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"computer_and_internet_info": {
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
						"cult_and_occult": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"travel": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"drugs": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"adult_and_pornography": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"home_and_garden": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"military": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"social_network": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dead_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"stock_advice_and_tools": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"training_and_tools": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dating": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"sex_education": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"religion": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"entertainment_and_arts": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"personal_sites_and_blogs": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"legal": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"local_information": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"streaming_media": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"job_search": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"gambling": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"translation": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"reference_and_research": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shareware_and_freeware": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"peer_to_peer": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"marijuana": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"hacking": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"games": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"philosophy_and_politics": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"weapons": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"pay_to_surf": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"hunting_and_fishing": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"society": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"educational_institutions": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"online_greeting_cards": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"sports": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"swimsuits_and_intimate_apparel": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"questionable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"kids": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"hate_and_racism": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"personal_storage": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"violence": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"keyloggers_and_monitoring": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"search_engines": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"internet_portals": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"web_advertisements": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cheating": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"gross": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"web_based_email": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"malware_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"phishing_and_other_fraud": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"proxy_avoid_and_anonymizers": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"spyware_and_adware": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"music": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"government": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"nudity": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"news_and_media": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"illegal": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cdns": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"internet_communications": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bot_nets": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"abortion": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"health_and_medicine": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"confirmed_spam_sources": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"spam_urls": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"unconfirmed_spam_sources": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"open_http_proxies": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dynamic_comment": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"parked_domains": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alcohol_and_tobacco": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"private_ip_addresses": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"image_and_video_search": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"fashion_and_beauty": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"recreation_and_hobbies": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"motor_vehicles": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"web_hosting_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"food_and_dining": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"nudity_artistic": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"illegal_pornography": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"exception_web_category": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_uncategorized": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_real_estate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_computer_and_internet_security": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_financial_services": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_business_and_economy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_computer_and_internet_info": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_auctions": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_shopping": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_cult_and_occult": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_travel": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_drugs": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_adult_and_pornography": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_home_and_garden": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_military": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_social_network": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_dead_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_stock_advice_and_tools": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_training_and_tools": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_dating": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_sex_education": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_religion": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_entertainment_and_arts": {
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
						"exception_local_information": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_streaming_media": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_job_search": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_gambling": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_translation": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_reference_and_research": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_shareware_and_freeware": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_peer_to_peer": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_marijuana": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_hacking": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_games": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_philosophy_and_politics": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_weapons": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_pay_to_surf": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_hunting_and_fishing": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_society": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_educational_institutions": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_online_greeting_cards": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_sports": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_swimsuits_and_intimate_apparel": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_questionable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_kids": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_hate_and_racism": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_personal_storage": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_violence": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_keyloggers_and_monitoring": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_search_engines": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_internet_portals": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_web_advertisements": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_cheating": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_gross": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_web_based_email": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_malware_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_phishing_and_other_fraud": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_proxy_avoid_and_anonymizers": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_spyware_and_adware": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_music": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_government": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_nudity": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_news_and_media": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_illegal": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_cdns": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_internet_communications": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_bot_nets": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_abortion": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_health_and_medicine": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_confirmed_spam_sources": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_spam_urls": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_unconfirmed_spam_sources": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_open_http_proxies": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_dynamic_comment": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_parked_domains": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_alcohol_and_tobacco": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_private_ip_addresses": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_image_and_video_search": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_fashion_and_beauty": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_recreation_and_hobbies": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_motor_vehicles": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_web_hosting_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_food_and_dining": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_nudity_artistic": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exception_illegal_pornography": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"require_web_category": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"client_ipv4_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_ipv4_list_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"client_ipv6_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_ipv6_list_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"server_ipv4_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_ipv4_list_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"server_ipv6_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_ipv6_list_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"exception_client_ipv4_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_client_ipv4_list_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"exception_client_ipv6_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_client_ipv6_list_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"exception_server_ipv4_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_server_ipv4_list_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"exception_server_ipv6_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_server_ipv6_list_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"local_cert_pin_list": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_cert_pin_list_bypass_fail_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"central_cert_pin_list": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_require_sni_cert_matched": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_cipher": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_cipher_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_cipher_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_hsm": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"hsm_type": {
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
			"server_name_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_cert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_chain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_passphrase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_name_alternate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_name_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_cert_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_chain_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_key_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_passphrase_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_name_regex_alternate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_shared_regex": {
							Type:        schema.TypeInt,
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
			"sni_enable_log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sni_bypass_missing_cert": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sni_bypass_expired_cert": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sni_bypass_explicit_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sni_bypass_enable_log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"direct_client_server_auth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"session_cache_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"session_cache_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"session_ticket_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"session_ticket_lifetime": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ssl_false_start_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_sslv3": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"version": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dgversion": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"renegotiation_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sslv2_bypass_service_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"authorization": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"authen_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ldap_base_dn_from_cert": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ldap_search_filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_sg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_sg_dn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auth_sg_filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_username_attribute": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"non_ssl_bypass_service_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"non_ssl_bypass_l4session": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enable_ssli_ftp_alg": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"early_data": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"no_anti_replay": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ja3_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ja3_insert_http_header": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ja3_reject_class_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ja3_reject_max_number_per_host": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ja3_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"certificate_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"passphrase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
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
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbTemplateClientSslCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateClientSsl (Inside resourceSlbTemplateClientSslCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateClientSsl(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateClientSsl --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateClientSsl(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateClientSslRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateClientSslRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateClientSsl (Inside resourceSlbTemplateClientSslRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateClientSsl(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceSlbTemplateClientSslUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplateClientSsl   (Inside resourceSlbTemplateClientSslUpdate) ")
		data := dataToSlbTemplateClientSsl(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateClientSsl ")
		err := go_thunder.PutSlbTemplateClientSsl(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateClientSslRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateClientSslDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateClientSslDelete) " + name1)
		err := go_thunder.DeleteSlbTemplateClientSsl(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplateClientSsl(d *schema.ResourceData) go_thunder.SlbTemplateClientSsl {
	var vc go_thunder.SlbTemplateClientSsl
	var c go_thunder.SlbTemplateClientSslInstance
	c.SlbTemplateClientSslInstanceName = d.Get("name").(string)
	c.SlbTemplateClientSslInstanceAuthUsername = d.Get("auth_username").(string)

	SlbTemplateClientSslInstanceCaCertsCount := d.Get("ca_certs.#").(int)
	c.SlbTemplateClientSslInstanceCaCertsCaCert = make([]go_thunder.SlbTemplateClientSslInstanceCaCerts, 0, SlbTemplateClientSslInstanceCaCertsCount)

	for i := 0; i < SlbTemplateClientSslInstanceCaCertsCount; i++ {
		var obj1 go_thunder.SlbTemplateClientSslInstanceCaCerts
		prefix1 := fmt.Sprintf("ca_certs.%d.", i)
		obj1.SlbTemplateClientSslInstanceCaCertsCaCert = d.Get(prefix1 + "ca_cert").(string)
		obj1.SlbTemplateClientSslInstanceCaCertsCaShared = d.Get(prefix1 + "ca_shared").(int)
		obj1.SlbTemplateClientSslInstanceCaCertsClientOcsp = d.Get(prefix1 + "client_ocsp").(int)
		obj1.SlbTemplateClientSslInstanceCaCertsClientOcspSrvr = d.Get(prefix1 + "client_ocsp_srvr").(string)
		obj1.SlbTemplateClientSslInstanceCaCertsClientOcspSg = d.Get(prefix1 + "client_ocsp_sg").(string)
		c.SlbTemplateClientSslInstanceCaCertsCaCert = append(c.SlbTemplateClientSslInstanceCaCertsCaCert, obj1)
	}

	c.SlbTemplateClientSslInstanceChainCert = d.Get("chain_cert").(string)
	c.SlbTemplateClientSslInstanceChainCertSharedStr = d.Get("chain_cert_shared_str").(string)
	c.SlbTemplateClientSslInstanceDhType = d.Get("dh_type").(string)

	SlbTemplateClientSslInstanceEcListCount := d.Get("ec_list.#").(int)
	c.SlbTemplateClientSslInstanceEcListEc = make([]go_thunder.SlbTemplateClientSslInstanceEcList, 0, SlbTemplateClientSslInstanceEcListCount)

	for i := 0; i < SlbTemplateClientSslInstanceEcListCount; i++ {
		var obj2 go_thunder.SlbTemplateClientSslInstanceEcList
		prefix2 := fmt.Sprintf("ec_list.%d.", i)
		obj2.SlbTemplateClientSslInstanceEcListEc = d.Get(prefix2 + "ec").(string)
		c.SlbTemplateClientSslInstanceEcListEc = append(c.SlbTemplateClientSslInstanceEcListEc, obj2)
	}

	c.SlbTemplateClientSslInstanceLocalLogging = d.Get("local_logging").(int)
	c.SlbTemplateClientSslInstanceOcspStapling = d.Get("ocsp_stapling").(int)
	c.SlbTemplateClientSslInstanceOcspstCaCert = d.Get("ocspst_ca_cert").(string)
	c.SlbTemplateClientSslInstanceOcspstOcsp = d.Get("ocspst_ocsp").(int)
	c.SlbTemplateClientSslInstanceOcspstSrvr = d.Get("ocspst_srvr").(string)
	c.SlbTemplateClientSslInstanceOcspstSrvrDays = d.Get("ocspst_srvr_days").(int)
	c.SlbTemplateClientSslInstanceOcspstSrvrHours = d.Get("ocspst_srvr_hours").(int)
	c.SlbTemplateClientSslInstanceOcspstSrvrMinutes = d.Get("ocspst_srvr_minutes").(int)
	c.SlbTemplateClientSslInstanceOcspstSrvrTimeout = d.Get("ocspst_srvr_timeout").(int)
	c.SlbTemplateClientSslInstanceOcspstSg = d.Get("ocspst_sg").(string)
	c.SlbTemplateClientSslInstanceOcspstSgDays = d.Get("ocspst_sg_days").(int)
	c.SlbTemplateClientSslInstanceOcspstSgHours = d.Get("ocspst_sg_hours").(int)
	c.SlbTemplateClientSslInstanceOcspstSgMinutes = d.Get("ocspst_sg_minutes").(int)
	c.SlbTemplateClientSslInstanceOcspstSgTimeout = d.Get("ocspst_sg_timeout").(int)
	c.SlbTemplateClientSslInstanceSsliLogging = d.Get("ssli_logging").(int)
	c.SlbTemplateClientSslInstanceSslilogging = d.Get("sslilogging").(string)
	c.SlbTemplateClientSslInstanceClientCertificate = d.Get("client_certificate").(string)

	SlbTemplateClientSslInstanceReqCaListsCount := d.Get("req_ca_lists.#").(int)
	c.SlbTemplateClientSslInstanceReqCaListsClientCertificateRequestCA = make([]go_thunder.SlbTemplateClientSslInstanceReqCaLists, 0, SlbTemplateClientSslInstanceReqCaListsCount)

	for i := 0; i < SlbTemplateClientSslInstanceReqCaListsCount; i++ {
		var obj3 go_thunder.SlbTemplateClientSslInstanceReqCaLists
		prefix3 := fmt.Sprintf("req_ca_lists.%d.", i)
		obj3.SlbTemplateClientSslInstanceReqCaListsClientCertificateRequestCA = d.Get(prefix3 + "client_certificate_request_ca").(string)
		obj3.SlbTemplateClientSslInstanceReqCaListsClientCertReqCaShared = d.Get(prefix3 + "client_cert_req_ca_shared").(int)
		c.SlbTemplateClientSslInstanceReqCaListsClientCertificateRequestCA = append(c.SlbTemplateClientSslInstanceReqCaListsClientCertificateRequestCA, obj3)
	}

	c.SlbTemplateClientSslInstanceCloseNotify = d.Get("close_notify").(int)

	SlbTemplateClientSslInstanceCrlCertsCount := d.Get("crl_certs.#").(int)
	c.SlbTemplateClientSslInstanceCrlCertsCrl = make([]go_thunder.SlbTemplateClientSslInstanceCrlCerts, 0, SlbTemplateClientSslInstanceCrlCertsCount)

	for i := 0; i < SlbTemplateClientSslInstanceCrlCertsCount; i++ {
		var obj4 go_thunder.SlbTemplateClientSslInstanceCrlCerts
		prefix4 := fmt.Sprintf("crl_certs.%d.", i)
		obj4.SlbTemplateClientSslInstanceCrlCertsCrl = d.Get(prefix4 + "crl").(string)
		obj4.SlbTemplateClientSslInstanceCrlCertsCrlShared = d.Get(prefix4 + "crl_shared").(int)
		c.SlbTemplateClientSslInstanceCrlCertsCrl = append(c.SlbTemplateClientSslInstanceCrlCertsCrl, obj4)
	}

	c.SlbTemplateClientSslInstanceForwardProxyCaCert = d.Get("forward_proxy_ca_cert").(string)
	c.SlbTemplateClientSslInstanceFpCaShared = d.Get("fp_ca_shared").(int)
	c.SlbTemplateClientSslInstanceForwardProxyCaKey = d.Get("forward_proxy_ca_key").(string)
	c.SlbTemplateClientSslInstanceForwardPassphrase = d.Get("forward_passphrase").(string)
	c.SlbTemplateClientSslInstanceFpCaKeyShared = d.Get("fp_ca_key_shared").(int)
	c.SlbTemplateClientSslInstanceFpCaCertificate = d.Get("fp_ca_certificate").(string)
	c.SlbTemplateClientSslInstanceFpCaKey = d.Get("fp_ca_key").(string)
	c.SlbTemplateClientSslInstanceFpCaKeyPassphrase = d.Get("fp_ca_key_passphrase").(string)
	c.SlbTemplateClientSslInstanceFpCaChainCert = d.Get("fp_ca_chain_cert").(string)
	c.SlbTemplateClientSslInstanceFpCaCertificateShared = d.Get("fp_ca_certificate_shared").(int)
	c.SlbTemplateClientSslInstanceForwardProxyAltSign = d.Get("forward_proxy_alt_sign").(int)
	c.SlbTemplateClientSslInstanceFpAltCert = d.Get("fp_alt_cert").(string)
	c.SlbTemplateClientSslInstanceFpAltKey = d.Get("fp_alt_key").(string)
	c.SlbTemplateClientSslInstanceFpAltPassphrase = d.Get("fp_alt_passphrase").(string)
	c.SlbTemplateClientSslInstanceFpAltChainCert = d.Get("fp_alt_chain_cert").(string)
	c.SlbTemplateClientSslInstanceFpAltShared = d.Get("fp_alt_shared").(int)

	SlbTemplateClientSslInstanceForwardProxyTrustedCaListsCount := d.Get("forward_proxy_trusted_ca_lists.#").(int)
	c.SlbTemplateClientSslInstanceForwardProxyTrustedCaListsForwardProxyTrustedCa = make([]go_thunder.SlbTemplateClientSslInstanceForwardProxyTrustedCaLists, 0, SlbTemplateClientSslInstanceForwardProxyTrustedCaListsCount)

	for i := 0; i < SlbTemplateClientSslInstanceForwardProxyTrustedCaListsCount; i++ {
		var obj5 go_thunder.SlbTemplateClientSslInstanceForwardProxyTrustedCaLists
		prefix5 := fmt.Sprintf("forward_proxy_trusted_ca_lists.%d.", i)
		obj5.SlbTemplateClientSslInstanceForwardProxyTrustedCaListsForwardProxyTrustedCa = d.Get(prefix5 + "forward_proxy_trusted_ca").(string)
		obj5.SlbTemplateClientSslInstanceForwardProxyTrustedCaListsFpTrustedCaShared = d.Get(prefix5 + "fp_trusted_ca_shared").(int)
		c.SlbTemplateClientSslInstanceForwardProxyTrustedCaListsForwardProxyTrustedCa = append(c.SlbTemplateClientSslInstanceForwardProxyTrustedCaListsForwardProxyTrustedCa, obj5)
	}

	c.SlbTemplateClientSslInstanceForwardProxyDecryptedDscp = d.Get("forward_proxy_decrypted_dscp").(int)
	c.SlbTemplateClientSslInstanceForwardProxyDecryptedDscpBypass = d.Get("forward_proxy_decrypted_dscp_bypass").(int)
	c.SlbTemplateClientSslInstanceEnableTLSAlertLogging = d.Get("enable_tls_alert_logging").(int)
	c.SlbTemplateClientSslInstanceAlertType = d.Get("alert_type").(string)
	c.SlbTemplateClientSslInstanceForwardProxyVerifyCertFailAction = d.Get("forward_proxy_verify_cert_fail_action").(int)
	c.SlbTemplateClientSslInstanceVerifyCertFailAction = d.Get("verify_cert_fail_action").(string)
	c.SlbTemplateClientSslInstanceForwardProxyCertRevokeAction = d.Get("forward_proxy_cert_revoke_action").(int)
	c.SlbTemplateClientSslInstanceCertRevokeAction = d.Get("cert_revoke_action").(string)
	c.SlbTemplateClientSslInstanceForwardProxyNoSharedCipherAction = d.Get("forward_proxy_no_shared_cipher_action").(int)
	c.SlbTemplateClientSslInstanceNoSharedCipherAction = d.Get("no_shared_cipher_action").(string)
	c.SlbTemplateClientSslInstanceForwardProxyEsniAction = d.Get("forward_proxy_esni_action").(int)
	c.SlbTemplateClientSslInstanceFpEsniAction = d.Get("fp_esni_action").(string)
	c.SlbTemplateClientSslInstanceForwardProxyCertUnknownAction = d.Get("forward_proxy_cert_unknown_action").(int)
	c.SlbTemplateClientSslInstanceCertUnknownAction = d.Get("cert_unknown_action").(string)
	c.SlbTemplateClientSslInstanceForwardProxyBlockMessage = d.Get("forward_proxy_block_message").(string)
	c.SlbTemplateClientSslInstanceCachePersistenceListName = d.Get("cache_persistence_list_name").(string)
	c.SlbTemplateClientSslInstanceFpCertExtCrldp = d.Get("fp_cert_ext_crldp").(string)
	c.SlbTemplateClientSslInstanceFpCertExtAiaOcsp = d.Get("fp_cert_ext_aia_ocsp").(string)
	c.SlbTemplateClientSslInstanceFpCertExtAiaCaIssuers = d.Get("fp_cert_ext_aia_ca_issuers").(string)
	c.SlbTemplateClientSslInstanceNotbefore = d.Get("notbefore").(int)
	c.SlbTemplateClientSslInstanceNotbeforeday = d.Get("notbeforeday").(int)
	c.SlbTemplateClientSslInstanceNotbeforemonth = d.Get("notbeforemonth").(int)
	c.SlbTemplateClientSslInstanceNotbeforeyear = d.Get("notbeforeyear").(int)
	c.SlbTemplateClientSslInstanceNotafter = d.Get("notafter").(int)
	c.SlbTemplateClientSslInstanceNotafterday = d.Get("notafterday").(int)
	c.SlbTemplateClientSslInstanceNotaftermonth = d.Get("notaftermonth").(int)
	c.SlbTemplateClientSslInstanceNotafteryear = d.Get("notafteryear").(int)
	c.SlbTemplateClientSslInstanceForwardProxyHashPersistenceInterval = d.Get("forward_proxy_hash_persistence_interval").(int)
	c.SlbTemplateClientSslInstanceForwardProxySslVersion = d.Get("forward_proxy_ssl_version").(int)
	c.SlbTemplateClientSslInstanceForwardProxyOcspDisable = d.Get("forward_proxy_ocsp_disable").(int)
	c.SlbTemplateClientSslInstanceForwardProxyCrlDisable = d.Get("forward_proxy_crl_disable").(int)
	c.SlbTemplateClientSslInstanceForwardProxyCertCacheTimeout = d.Get("forward_proxy_cert_cache_timeout").(int)
	c.SlbTemplateClientSslInstanceForwardProxyCertCacheLimit = d.Get("forward_proxy_cert_cache_limit").(int)
	c.SlbTemplateClientSslInstanceForwardProxyCertExpiry = d.Get("forward_proxy_cert_expiry").(int)
	c.SlbTemplateClientSslInstanceExpireHours = d.Get("expire_hours").(int)
	c.SlbTemplateClientSslInstanceForwardProxyEnable = d.Get("forward_proxy_enable").(int)
	c.SlbTemplateClientSslInstanceHandshakeLoggingEnable = d.Get("handshake_logging_enable").(int)
	c.SlbTemplateClientSslInstanceForwardProxySelfsignRedir = d.Get("forward_proxy_selfsign_redir").(int)
	c.SlbTemplateClientSslInstanceForwardProxyFailsafeDisable = d.Get("forward_proxy_failsafe_disable").(int)
	c.SlbTemplateClientSslInstanceForwardProxyLogDisable = d.Get("forward_proxy_log_disable").(int)
	c.SlbTemplateClientSslInstanceFpCertFetchNatpoolName = d.Get("fp_cert_fetch_natpool_name").(string)
	c.SlbTemplateClientSslInstanceSharedPartitionPool = d.Get("shared_partition_pool").(int)
	c.SlbTemplateClientSslInstanceFpCertFetchNatpoolNameShared = d.Get("fp_cert_fetch_natpool_name_shared").(string)
	c.SlbTemplateClientSslInstanceFpCertFetchNatpoolPrecedence = d.Get("fp_cert_fetch_natpool_precedence").(int)
	c.SlbTemplateClientSslInstanceFpCertFetchAutonat = d.Get("fp_cert_fetch_autonat").(string)
	c.SlbTemplateClientSslInstanceFpCertFetchAutonatPrecedence = d.Get("fp_cert_fetch_autonat_precedence").(int)
	c.SlbTemplateClientSslInstanceForwardProxyNoSniAction = d.Get("forward_proxy_no_sni_action").(string)
	c.SlbTemplateClientSslInstanceCaseInsensitive = d.Get("case_insensitive").(int)
	c.SlbTemplateClientSslInstanceClassListName = d.Get("class_list_name").(string)

	SlbTemplateClientSslInstanceMultiClassListCount := d.Get("multi_class_list.#").(int)
	c.SlbTemplateClientSslInstanceMultiClassListMultiClistName = make([]go_thunder.SlbTemplateClientSslInstanceMultiClassList, 0, SlbTemplateClientSslInstanceMultiClassListCount)

	for i := 0; i < SlbTemplateClientSslInstanceMultiClassListCount; i++ {
		var obj6 go_thunder.SlbTemplateClientSslInstanceMultiClassList
		prefix6 := fmt.Sprintf("multi_class_list.%d.", i)
		obj6.SlbTemplateClientSslInstanceMultiClassListMultiClistName = d.Get(prefix6 + "multi_clist_name").(string)
		c.SlbTemplateClientSslInstanceMultiClassListMultiClistName = append(c.SlbTemplateClientSslInstanceMultiClassListMultiClistName, obj6)
	}

	c.SlbTemplateClientSslInstanceUserNameList = d.Get("user_name_list").(string)
	c.SlbTemplateClientSslInstanceAdGroupList = d.Get("ad_group_list").(string)
	c.SlbTemplateClientSslInstanceExceptionUserNameList = d.Get("exception_user_name_list").(string)
	c.SlbTemplateClientSslInstanceExceptionAdGroupList = d.Get("exception_ad_group_list").(string)
	c.SlbTemplateClientSslInstanceExceptionSniClName = d.Get("exception_sni_cl_name").(string)
	c.SlbTemplateClientSslInstanceInspectListName = d.Get("inspect_list_name").(string)
	c.SlbTemplateClientSslInstanceInspectCertificateSubjectClName = d.Get("inspect_certificate_subject_cl_name").(string)
	c.SlbTemplateClientSslInstanceInspectCertificateIssuerClName = d.Get("inspect_certificate_issuer_cl_name").(string)
	c.SlbTemplateClientSslInstanceInspectCertificateSanClName = d.Get("inspect_certificate_san_cl_name").(string)

	SlbTemplateClientSslInstanceContainsListCount := d.Get("contains_list.#").(int)
	c.SlbTemplateClientSslInstanceContainsListContains = make([]go_thunder.SlbTemplateClientSslInstanceContainsList, 0, SlbTemplateClientSslInstanceContainsListCount)

	for i := 0; i < SlbTemplateClientSslInstanceContainsListCount; i++ {
		var obj7 go_thunder.SlbTemplateClientSslInstanceContainsList
		prefix7 := fmt.Sprintf("contains_list.%d.", i)
		obj7.SlbTemplateClientSslInstanceContainsListContains = d.Get(prefix7 + "contains").(string)
		c.SlbTemplateClientSslInstanceContainsListContains = append(c.SlbTemplateClientSslInstanceContainsListContains, obj7)
	}

	SlbTemplateClientSslInstanceEndsWithListCount := d.Get("ends_with_list.#").(int)
	c.SlbTemplateClientSslInstanceEndsWithListEndsWith = make([]go_thunder.SlbTemplateClientSslInstanceEndsWithList, 0, SlbTemplateClientSslInstanceEndsWithListCount)

	for i := 0; i < SlbTemplateClientSslInstanceEndsWithListCount; i++ {
		var obj8 go_thunder.SlbTemplateClientSslInstanceEndsWithList
		prefix8 := fmt.Sprintf("ends_with_list.%d.", i)
		obj8.SlbTemplateClientSslInstanceEndsWithListEndsWith = d.Get(prefix8 + "ends_with").(string)
		c.SlbTemplateClientSslInstanceEndsWithListEndsWith = append(c.SlbTemplateClientSslInstanceEndsWithListEndsWith, obj8)
	}

	SlbTemplateClientSslInstanceEqualsListCount := d.Get("equals_list.#").(int)
	c.SlbTemplateClientSslInstanceEqualsListEquals = make([]go_thunder.SlbTemplateClientSslInstanceEqualsList, 0, SlbTemplateClientSslInstanceEqualsListCount)

	for i := 0; i < SlbTemplateClientSslInstanceEqualsListCount; i++ {
		var obj9 go_thunder.SlbTemplateClientSslInstanceEqualsList
		prefix9 := fmt.Sprintf("equals_list.%d.", i)
		obj9.SlbTemplateClientSslInstanceEqualsListEquals = d.Get(prefix9 + "equals").(string)
		c.SlbTemplateClientSslInstanceEqualsListEquals = append(c.SlbTemplateClientSslInstanceEqualsListEquals, obj9)
	}

	SlbTemplateClientSslInstanceStartsWithListCount := d.Get("starts_with_list.#").(int)
	c.SlbTemplateClientSslInstanceStartsWithListStartsWith = make([]go_thunder.SlbTemplateClientSslInstanceStartsWithList, 0, SlbTemplateClientSslInstanceStartsWithListCount)

	for i := 0; i < SlbTemplateClientSslInstanceStartsWithListCount; i++ {
		var obj10 go_thunder.SlbTemplateClientSslInstanceStartsWithList
		prefix10 := fmt.Sprintf("starts_with_list.%d.", i)
		obj10.SlbTemplateClientSslInstanceStartsWithListStartsWith = d.Get(prefix10 + "starts_with").(string)
		c.SlbTemplateClientSslInstanceStartsWithListStartsWith = append(c.SlbTemplateClientSslInstanceStartsWithListStartsWith, obj10)
	}

	SlbTemplateClientSslInstanceCertificateSubjectContainsListCount := d.Get("certificate_subject_contains_list.#").(int)
	c.SlbTemplateClientSslInstanceCertificateSubjectContainsListCertificateSubjectContains = make([]go_thunder.SlbTemplateClientSslInstanceCertificateSubjectContainsList, 0, SlbTemplateClientSslInstanceCertificateSubjectContainsListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCertificateSubjectContainsListCount; i++ {
		var obj11 go_thunder.SlbTemplateClientSslInstanceCertificateSubjectContainsList
		prefix11 := fmt.Sprintf("certificate_subject_contains_list.%d.", i)
		obj11.SlbTemplateClientSslInstanceCertificateSubjectContainsListCertificateSubjectContains = d.Get(prefix11 + "certificate_subject_contains").(string)
		c.SlbTemplateClientSslInstanceCertificateSubjectContainsListCertificateSubjectContains = append(c.SlbTemplateClientSslInstanceCertificateSubjectContainsListCertificateSubjectContains, obj11)
	}

	c.SlbTemplateClientSslInstanceBypassCertSubjectClassListName = d.Get("bypass_cert_subject_class_list_name").(string)

	SlbTemplateClientSslInstanceBypassCertSubjectMultiClassListCount := d.Get("bypass_cert_subject_multi_class_list.#").(int)
	c.SlbTemplateClientSslInstanceBypassCertSubjectMultiClassListBypassCertSubjectMultiClassListName = make([]go_thunder.SlbTemplateClientSslInstanceBypassCertSubjectMultiClassList, 0, SlbTemplateClientSslInstanceBypassCertSubjectMultiClassListCount)

	for i := 0; i < SlbTemplateClientSslInstanceBypassCertSubjectMultiClassListCount; i++ {
		var obj12 go_thunder.SlbTemplateClientSslInstanceBypassCertSubjectMultiClassList
		prefix12 := fmt.Sprintf("bypass_cert_subject_multi_class_list.%d.", i)
		obj12.SlbTemplateClientSslInstanceBypassCertSubjectMultiClassListBypassCertSubjectMultiClassListName = d.Get(prefix12 + "bypass_cert_subject_multi_class_list_name").(string)
		c.SlbTemplateClientSslInstanceBypassCertSubjectMultiClassListBypassCertSubjectMultiClassListName = append(c.SlbTemplateClientSslInstanceBypassCertSubjectMultiClassListBypassCertSubjectMultiClassListName, obj12)
	}

	c.SlbTemplateClientSslInstanceExceptionCertificateSubjectClName = d.Get("exception_certificate_subject_cl_name").(string)

	SlbTemplateClientSslInstanceCertificateSubjectEndsWithListCount := d.Get("certificate_subject_ends_with_list.#").(int)
	c.SlbTemplateClientSslInstanceCertificateSubjectEndsWithListCertificateSubjectEndsWith = make([]go_thunder.SlbTemplateClientSslInstanceCertificateSubjectEndsWithList, 0, SlbTemplateClientSslInstanceCertificateSubjectEndsWithListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCertificateSubjectEndsWithListCount; i++ {
		var obj13 go_thunder.SlbTemplateClientSslInstanceCertificateSubjectEndsWithList
		prefix13 := fmt.Sprintf("certificate_subject_ends_with_list.%d.", i)
		obj13.SlbTemplateClientSslInstanceCertificateSubjectEndsWithListCertificateSubjectEndsWith = d.Get(prefix13 + "certificate_subject_ends_with").(string)
		c.SlbTemplateClientSslInstanceCertificateSubjectEndsWithListCertificateSubjectEndsWith = append(c.SlbTemplateClientSslInstanceCertificateSubjectEndsWithListCertificateSubjectEndsWith, obj13)
	}

	SlbTemplateClientSslInstanceCertificateSubjectEqualsListCount := d.Get("certificate_subject_equals_list.#").(int)
	c.SlbTemplateClientSslInstanceCertificateSubjectEqualsListCertificateSubjectEquals = make([]go_thunder.SlbTemplateClientSslInstanceCertificateSubjectEqualsList, 0, SlbTemplateClientSslInstanceCertificateSubjectEqualsListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCertificateSubjectEqualsListCount; i++ {
		var obj14 go_thunder.SlbTemplateClientSslInstanceCertificateSubjectEqualsList
		prefix14 := fmt.Sprintf("certificate_subject_equals_list.%d.", i)
		obj14.SlbTemplateClientSslInstanceCertificateSubjectEqualsListCertificateSubjectEquals = d.Get(prefix14 + "certificate_subject_equals").(string)
		c.SlbTemplateClientSslInstanceCertificateSubjectEqualsListCertificateSubjectEquals = append(c.SlbTemplateClientSslInstanceCertificateSubjectEqualsListCertificateSubjectEquals, obj14)
	}

	SlbTemplateClientSslInstanceCertificateSubjectStartsWithListCount := d.Get("certificate_subject_starts_with_list.#").(int)
	c.SlbTemplateClientSslInstanceCertificateSubjectStartsWithListCertificateSubjectStarts = make([]go_thunder.SlbTemplateClientSslInstanceCertificateSubjectStartsWithList, 0, SlbTemplateClientSslInstanceCertificateSubjectStartsWithListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCertificateSubjectStartsWithListCount; i++ {
		var obj15 go_thunder.SlbTemplateClientSslInstanceCertificateSubjectStartsWithList
		prefix15 := fmt.Sprintf("certificate_subject_starts_with_list.%d.", i)
		obj15.SlbTemplateClientSslInstanceCertificateSubjectStartsWithListCertificateSubjectStarts = d.Get(prefix15 + "certificate_subject_starts").(string)
		c.SlbTemplateClientSslInstanceCertificateSubjectStartsWithListCertificateSubjectStarts = append(c.SlbTemplateClientSslInstanceCertificateSubjectStartsWithListCertificateSubjectStarts, obj15)
	}

	SlbTemplateClientSslInstanceCertificateIssuerContainsListCount := d.Get("certificate_issuer_contains_list.#").(int)
	c.SlbTemplateClientSslInstanceCertificateIssuerContainsListCertificateIssuerContains = make([]go_thunder.SlbTemplateClientSslInstanceCertificateIssuerContainsList, 0, SlbTemplateClientSslInstanceCertificateIssuerContainsListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCertificateIssuerContainsListCount; i++ {
		var obj16 go_thunder.SlbTemplateClientSslInstanceCertificateIssuerContainsList
		prefix16 := fmt.Sprintf("certificate_issuer_contains_list.%d.", i)
		obj16.SlbTemplateClientSslInstanceCertificateIssuerContainsListCertificateIssuerContains = d.Get(prefix16 + "certificate_issuer_contains").(string)
		c.SlbTemplateClientSslInstanceCertificateIssuerContainsListCertificateIssuerContains = append(c.SlbTemplateClientSslInstanceCertificateIssuerContainsListCertificateIssuerContains, obj16)
	}

	c.SlbTemplateClientSslInstanceBypassCertIssuerClassListName = d.Get("bypass_cert_issuer_class_list_name").(string)

	SlbTemplateClientSslInstanceBypassCertIssuerMultiClassListCount := d.Get("bypass_cert_issuer_multi_class_list.#").(int)
	c.SlbTemplateClientSslInstanceBypassCertIssuerMultiClassListBypassCertIssuerMultiClassListName = make([]go_thunder.SlbTemplateClientSslInstanceBypassCertIssuerMultiClassList, 0, SlbTemplateClientSslInstanceBypassCertIssuerMultiClassListCount)

	for i := 0; i < SlbTemplateClientSslInstanceBypassCertIssuerMultiClassListCount; i++ {
		var obj17 go_thunder.SlbTemplateClientSslInstanceBypassCertIssuerMultiClassList
		prefix17 := fmt.Sprintf("bypass_cert_issuer_multi_class_list.%d.", i)
		obj17.SlbTemplateClientSslInstanceBypassCertIssuerMultiClassListBypassCertIssuerMultiClassListName = d.Get(prefix17 + "bypass_cert_issuer_multi_class_list_name").(string)
		c.SlbTemplateClientSslInstanceBypassCertIssuerMultiClassListBypassCertIssuerMultiClassListName = append(c.SlbTemplateClientSslInstanceBypassCertIssuerMultiClassListBypassCertIssuerMultiClassListName, obj17)
	}

	c.SlbTemplateClientSslInstanceExceptionCertificateIssuerClName = d.Get("exception_certificate_issuer_cl_name").(string)

	SlbTemplateClientSslInstanceCertificateIssuerEndsWithListCount := d.Get("certificate_issuer_ends_with_list.#").(int)
	c.SlbTemplateClientSslInstanceCertificateIssuerEndsWithListCertificateIssuerEndsWith = make([]go_thunder.SlbTemplateClientSslInstanceCertificateIssuerEndsWithList, 0, SlbTemplateClientSslInstanceCertificateIssuerEndsWithListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCertificateIssuerEndsWithListCount; i++ {
		var obj18 go_thunder.SlbTemplateClientSslInstanceCertificateIssuerEndsWithList
		prefix18 := fmt.Sprintf("certificate_issuer_ends_with_list.%d.", i)
		obj18.SlbTemplateClientSslInstanceCertificateIssuerEndsWithListCertificateIssuerEndsWith = d.Get(prefix18 + "certificate_issuer_ends_with").(string)
		c.SlbTemplateClientSslInstanceCertificateIssuerEndsWithListCertificateIssuerEndsWith = append(c.SlbTemplateClientSslInstanceCertificateIssuerEndsWithListCertificateIssuerEndsWith, obj18)
	}

	SlbTemplateClientSslInstanceCertificateIssuerEqualsListCount := d.Get("certificate_issuer_equals_list.#").(int)
	c.SlbTemplateClientSslInstanceCertificateIssuerEqualsListCertificateIssuerEquals = make([]go_thunder.SlbTemplateClientSslInstanceCertificateIssuerEqualsList, 0, SlbTemplateClientSslInstanceCertificateIssuerEqualsListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCertificateIssuerEqualsListCount; i++ {
		var obj19 go_thunder.SlbTemplateClientSslInstanceCertificateIssuerEqualsList
		prefix19 := fmt.Sprintf("certificate_issuer_equals_list.%d.", i)
		obj19.SlbTemplateClientSslInstanceCertificateIssuerEqualsListCertificateIssuerEquals = d.Get(prefix19 + "certificate_issuer_equals").(string)
		c.SlbTemplateClientSslInstanceCertificateIssuerEqualsListCertificateIssuerEquals = append(c.SlbTemplateClientSslInstanceCertificateIssuerEqualsListCertificateIssuerEquals, obj19)
	}

	SlbTemplateClientSslInstanceCertificateIssuerStartsWithListCount := d.Get("certificate_issuer_starts_with_list.#").(int)
	c.SlbTemplateClientSslInstanceCertificateIssuerStartsWithListCertificateIssuerStarts = make([]go_thunder.SlbTemplateClientSslInstanceCertificateIssuerStartsWithList, 0, SlbTemplateClientSslInstanceCertificateIssuerStartsWithListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCertificateIssuerStartsWithListCount; i++ {
		var obj20 go_thunder.SlbTemplateClientSslInstanceCertificateIssuerStartsWithList
		prefix20 := fmt.Sprintf("certificate_issuer_starts_with_list.%d.", i)
		obj20.SlbTemplateClientSslInstanceCertificateIssuerStartsWithListCertificateIssuerStarts = d.Get(prefix20 + "certificate_issuer_starts").(string)
		c.SlbTemplateClientSslInstanceCertificateIssuerStartsWithListCertificateIssuerStarts = append(c.SlbTemplateClientSslInstanceCertificateIssuerStartsWithListCertificateIssuerStarts, obj20)
	}

	SlbTemplateClientSslInstanceCertificateSanContainsListCount := d.Get("certificate_san_contains_list.#").(int)
	c.SlbTemplateClientSslInstanceCertificateSanContainsListCertificateSanContains = make([]go_thunder.SlbTemplateClientSslInstanceCertificateSanContainsList, 0, SlbTemplateClientSslInstanceCertificateSanContainsListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCertificateSanContainsListCount; i++ {
		var obj21 go_thunder.SlbTemplateClientSslInstanceCertificateSanContainsList
		prefix21 := fmt.Sprintf("certificate_san_contains_list.%d.", i)
		obj21.SlbTemplateClientSslInstanceCertificateSanContainsListCertificateSanContains = d.Get(prefix21 + "certificate_san_contains").(string)
		c.SlbTemplateClientSslInstanceCertificateSanContainsListCertificateSanContains = append(c.SlbTemplateClientSslInstanceCertificateSanContainsListCertificateSanContains, obj21)
	}

	c.SlbTemplateClientSslInstanceBypassCertSanClassListName = d.Get("bypass_cert_san_class_list_name").(string)

	SlbTemplateClientSslInstanceBypassCertSanMultiClassListCount := d.Get("bypass_cert_san_multi_class_list.#").(int)
	c.SlbTemplateClientSslInstanceBypassCertSanMultiClassListBypassCertSanMultiClassListName = make([]go_thunder.SlbTemplateClientSslInstanceBypassCertSanMultiClassList, 0, SlbTemplateClientSslInstanceBypassCertSanMultiClassListCount)

	for i := 0; i < SlbTemplateClientSslInstanceBypassCertSanMultiClassListCount; i++ {
		var obj22 go_thunder.SlbTemplateClientSslInstanceBypassCertSanMultiClassList
		prefix22 := fmt.Sprintf("bypass_cert_san_multi_class_list.%d.", i)
		obj22.SlbTemplateClientSslInstanceBypassCertSanMultiClassListBypassCertSanMultiClassListName = d.Get(prefix22 + "bypass_cert_san_multi_class_list_name").(string)
		c.SlbTemplateClientSslInstanceBypassCertSanMultiClassListBypassCertSanMultiClassListName = append(c.SlbTemplateClientSslInstanceBypassCertSanMultiClassListBypassCertSanMultiClassListName, obj22)
	}

	c.SlbTemplateClientSslInstanceExceptionCertificateSanClName = d.Get("exception_certificate_san_cl_name").(string)

	SlbTemplateClientSslInstanceCertificateSanEndsWithListCount := d.Get("certificate_san_ends_with_list.#").(int)
	c.SlbTemplateClientSslInstanceCertificateSanEndsWithListCertificateSanEndsWith = make([]go_thunder.SlbTemplateClientSslInstanceCertificateSanEndsWithList, 0, SlbTemplateClientSslInstanceCertificateSanEndsWithListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCertificateSanEndsWithListCount; i++ {
		var obj23 go_thunder.SlbTemplateClientSslInstanceCertificateSanEndsWithList
		prefix23 := fmt.Sprintf("certificate_san_ends_with_list.%d.", i)
		obj23.SlbTemplateClientSslInstanceCertificateSanEndsWithListCertificateSanEndsWith = d.Get(prefix23 + "certificate_san_ends_with").(string)
		c.SlbTemplateClientSslInstanceCertificateSanEndsWithListCertificateSanEndsWith = append(c.SlbTemplateClientSslInstanceCertificateSanEndsWithListCertificateSanEndsWith, obj23)
	}

	SlbTemplateClientSslInstanceCertificateSanEqualsListCount := d.Get("certificate_san_equals_list.#").(int)
	c.SlbTemplateClientSslInstanceCertificateSanEqualsListCertificateSanEquals = make([]go_thunder.SlbTemplateClientSslInstanceCertificateSanEqualsList, 0, SlbTemplateClientSslInstanceCertificateSanEqualsListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCertificateSanEqualsListCount; i++ {
		var obj24 go_thunder.SlbTemplateClientSslInstanceCertificateSanEqualsList
		prefix24 := fmt.Sprintf("certificate_san_equals_list.%d.", i)
		obj24.SlbTemplateClientSslInstanceCertificateSanEqualsListCertificateSanEquals = d.Get(prefix24 + "certificate_san_equals").(string)
		c.SlbTemplateClientSslInstanceCertificateSanEqualsListCertificateSanEquals = append(c.SlbTemplateClientSslInstanceCertificateSanEqualsListCertificateSanEquals, obj24)
	}

	SlbTemplateClientSslInstanceCertificateSanStartsWithListCount := d.Get("certificate_san_starts_with_list.#").(int)
	c.SlbTemplateClientSslInstanceCertificateSanStartsWithListCertificateSanStarts = make([]go_thunder.SlbTemplateClientSslInstanceCertificateSanStartsWithList, 0, SlbTemplateClientSslInstanceCertificateSanStartsWithListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCertificateSanStartsWithListCount; i++ {
		var obj25 go_thunder.SlbTemplateClientSslInstanceCertificateSanStartsWithList
		prefix25 := fmt.Sprintf("certificate_san_starts_with_list.%d.", i)
		obj25.SlbTemplateClientSslInstanceCertificateSanStartsWithListCertificateSanStarts = d.Get(prefix25 + "certificate_san_starts").(string)
		c.SlbTemplateClientSslInstanceCertificateSanStartsWithListCertificateSanStarts = append(c.SlbTemplateClientSslInstanceCertificateSanStartsWithListCertificateSanStarts, obj25)
	}

	c.SlbTemplateClientSslInstanceClientAuthCaseInsensitive = d.Get("client_auth_case_insensitive").(int)
	c.SlbTemplateClientSslInstanceClientAuthClassList = d.Get("client_auth_class_list").(string)

	SlbTemplateClientSslInstanceClientAuthContainsListCount := d.Get("client_auth_contains_list.#").(int)
	c.SlbTemplateClientSslInstanceClientAuthContainsListClientAuthContains = make([]go_thunder.SlbTemplateClientSslInstanceClientAuthContainsList, 0, SlbTemplateClientSslInstanceClientAuthContainsListCount)

	for i := 0; i < SlbTemplateClientSslInstanceClientAuthContainsListCount; i++ {
		var obj26 go_thunder.SlbTemplateClientSslInstanceClientAuthContainsList
		prefix26 := fmt.Sprintf("client_auth_contains_list.%d.", i)
		obj26.SlbTemplateClientSslInstanceClientAuthContainsListClientAuthContains = d.Get(prefix26 + "client_auth_contains").(string)
		c.SlbTemplateClientSslInstanceClientAuthContainsListClientAuthContains = append(c.SlbTemplateClientSslInstanceClientAuthContainsListClientAuthContains, obj26)
	}

	SlbTemplateClientSslInstanceClientAuthEndsWithListCount := d.Get("client_auth_ends_with_list.#").(int)
	c.SlbTemplateClientSslInstanceClientAuthEndsWithListClientAuthEndsWith = make([]go_thunder.SlbTemplateClientSslInstanceClientAuthEndsWithList, 0, SlbTemplateClientSslInstanceClientAuthEndsWithListCount)

	for i := 0; i < SlbTemplateClientSslInstanceClientAuthEndsWithListCount; i++ {
		var obj27 go_thunder.SlbTemplateClientSslInstanceClientAuthEndsWithList
		prefix27 := fmt.Sprintf("client_auth_ends_with_list.%d.", i)
		obj27.SlbTemplateClientSslInstanceClientAuthEndsWithListClientAuthEndsWith = d.Get(prefix27 + "client_auth_ends_with").(string)
		c.SlbTemplateClientSslInstanceClientAuthEndsWithListClientAuthEndsWith = append(c.SlbTemplateClientSslInstanceClientAuthEndsWithListClientAuthEndsWith, obj27)
	}

	SlbTemplateClientSslInstanceClientAuthEqualsListCount := d.Get("client_auth_equals_list.#").(int)
	c.SlbTemplateClientSslInstanceClientAuthEqualsListClientAuthEquals = make([]go_thunder.SlbTemplateClientSslInstanceClientAuthEqualsList, 0, SlbTemplateClientSslInstanceClientAuthEqualsListCount)

	for i := 0; i < SlbTemplateClientSslInstanceClientAuthEqualsListCount; i++ {
		var obj28 go_thunder.SlbTemplateClientSslInstanceClientAuthEqualsList
		prefix28 := fmt.Sprintf("client_auth_equals_list.%d.", i)
		obj28.SlbTemplateClientSslInstanceClientAuthEqualsListClientAuthEquals = d.Get(prefix28 + "client_auth_equals").(string)
		c.SlbTemplateClientSslInstanceClientAuthEqualsListClientAuthEquals = append(c.SlbTemplateClientSslInstanceClientAuthEqualsListClientAuthEquals, obj28)
	}

	SlbTemplateClientSslInstanceClientAuthStartsWithListCount := d.Get("client_auth_starts_with_list.#").(int)
	c.SlbTemplateClientSslInstanceClientAuthStartsWithListClientAuthStartsWith = make([]go_thunder.SlbTemplateClientSslInstanceClientAuthStartsWithList, 0, SlbTemplateClientSslInstanceClientAuthStartsWithListCount)

	for i := 0; i < SlbTemplateClientSslInstanceClientAuthStartsWithListCount; i++ {
		var obj29 go_thunder.SlbTemplateClientSslInstanceClientAuthStartsWithList
		prefix29 := fmt.Sprintf("client_auth_starts_with_list.%d.", i)
		obj29.SlbTemplateClientSslInstanceClientAuthStartsWithListClientAuthStartsWith = d.Get(prefix29 + "client_auth_starts_with").(string)
		c.SlbTemplateClientSslInstanceClientAuthStartsWithListClientAuthStartsWith = append(c.SlbTemplateClientSslInstanceClientAuthStartsWithListClientAuthStartsWith, obj29)
	}

	c.SlbTemplateClientSslInstanceForwardProxyCertNotReadyAction = d.Get("forward_proxy_cert_not_ready_action").(string)

	var obj30 go_thunder.SlbTemplateClientSslInstanceWebReputation
	prefix30 := "web_reputation.0."
	obj30.SlbTemplateClientSslInstanceWebReputationBypassTrustworthy = d.Get(prefix30 + "bypass_trustworthy").(int)
	obj30.SlbTemplateClientSslInstanceWebReputationBypassLowRisk = d.Get(prefix30 + "bypass_low_risk").(int)
	obj30.SlbTemplateClientSslInstanceWebReputationBypassModerateRisk = d.Get(prefix30 + "bypass_moderate_risk").(int)
	obj30.SlbTemplateClientSslInstanceWebReputationBypassSuspicious = d.Get(prefix30 + "bypass_suspicious").(int)
	obj30.SlbTemplateClientSslInstanceWebReputationBypassMalicious = d.Get(prefix30 + "bypass_malicious").(int)
	obj30.SlbTemplateClientSslInstanceWebReputationBypassThreshold = d.Get(prefix30 + "bypass_threshold").(int)

	c.SlbTemplateClientSslInstanceWebReputationBypassTrustworthy = obj30

	var obj31 go_thunder.SlbTemplateClientSslInstanceExceptionWebReputation
	prefix31 := "exception_web_reputation.0."
	obj31.SlbTemplateClientSslInstanceExceptionWebReputationExceptionTrustworthy = d.Get(prefix31 + "exception_trustworthy").(int)
	obj31.SlbTemplateClientSslInstanceExceptionWebReputationExceptionLowRisk = d.Get(prefix31 + "exception_low_risk").(int)
	obj31.SlbTemplateClientSslInstanceExceptionWebReputationExceptionModerateRisk = d.Get(prefix31 + "exception_moderate_risk").(int)
	obj31.SlbTemplateClientSslInstanceExceptionWebReputationExceptionSuspicious = d.Get(prefix31 + "exception_suspicious").(int)
	obj31.SlbTemplateClientSslInstanceExceptionWebReputationExceptionMalicious = d.Get(prefix31 + "exception_malicious").(int)
	obj31.SlbTemplateClientSslInstanceExceptionWebReputationExceptionThreshold = d.Get(prefix31 + "exception_threshold").(int)

	c.SlbTemplateClientSslInstanceExceptionWebReputationExceptionTrustworthy = obj31

	var obj32 go_thunder.SlbTemplateClientSslInstanceWebCategory
	prefix32 := "web_category.0."
	obj32.SlbTemplateClientSslInstanceWebCategoryUncategorized = d.Get(prefix32 + "uncategorized").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryRealEstate = d.Get(prefix32 + "real_estate").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryComputerAndInternetSecurity = d.Get(prefix32 + "computer_and_internet_security").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryFinancialServices = d.Get(prefix32 + "financial_services").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryBusinessAndEconomy = d.Get(prefix32 + "business_and_economy").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryComputerAndInternetInfo = d.Get(prefix32 + "computer_and_internet_info").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryAuctions = d.Get(prefix32 + "auctions").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryShopping = d.Get(prefix32 + "shopping").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryCultAndOccult = d.Get(prefix32 + "cult_and_occult").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryTravel = d.Get(prefix32 + "travel").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryDrugs = d.Get(prefix32 + "drugs").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryAdultAndPornography = d.Get(prefix32 + "adult_and_pornography").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryHomeAndGarden = d.Get(prefix32 + "home_and_garden").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryMilitary = d.Get(prefix32 + "military").(int)
	obj32.SlbTemplateClientSslInstanceWebCategorySocialNetwork = d.Get(prefix32 + "social_network").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryDeadSites = d.Get(prefix32 + "dead_sites").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryStockAdviceAndTools = d.Get(prefix32 + "stock_advice_and_tools").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryTrainingAndTools = d.Get(prefix32 + "training_and_tools").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryDating = d.Get(prefix32 + "dating").(int)
	obj32.SlbTemplateClientSslInstanceWebCategorySexEducation = d.Get(prefix32 + "sex_education").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryReligion = d.Get(prefix32 + "religion").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryEntertainmentAndArts = d.Get(prefix32 + "entertainment_and_arts").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryPersonalSitesAndBlogs = d.Get(prefix32 + "personal_sites_and_blogs").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryLegal = d.Get(prefix32 + "legal").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryLocalInformation = d.Get(prefix32 + "local_information").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryStreamingMedia = d.Get(prefix32 + "streaming_media").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryJobSearch = d.Get(prefix32 + "job_search").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryGambling = d.Get(prefix32 + "gambling").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryTranslation = d.Get(prefix32 + "translation").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryReferenceAndResearch = d.Get(prefix32 + "reference_and_research").(int)
	obj32.SlbTemplateClientSslInstanceWebCategorySharewareAndFreeware = d.Get(prefix32 + "shareware_and_freeware").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryPeerToPeer = d.Get(prefix32 + "peer_to_peer").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryMarijuana = d.Get(prefix32 + "marijuana").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryHacking = d.Get(prefix32 + "hacking").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryGames = d.Get(prefix32 + "games").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryPhilosophyAndPolitics = d.Get(prefix32 + "philosophy_and_politics").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryWeapons = d.Get(prefix32 + "weapons").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryPayToSurf = d.Get(prefix32 + "pay_to_surf").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryHuntingAndFishing = d.Get(prefix32 + "hunting_and_fishing").(int)
	obj32.SlbTemplateClientSslInstanceWebCategorySociety = d.Get(prefix32 + "society").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryEducationalInstitutions = d.Get(prefix32 + "educational_institutions").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryOnlineGreetingCards = d.Get(prefix32 + "online_greeting_cards").(int)
	obj32.SlbTemplateClientSslInstanceWebCategorySports = d.Get(prefix32 + "sports").(int)
	obj32.SlbTemplateClientSslInstanceWebCategorySwimsuitsAndIntimateApparel = d.Get(prefix32 + "swimsuits_and_intimate_apparel").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryQuestionable = d.Get(prefix32 + "questionable").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryKids = d.Get(prefix32 + "kids").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryHateAndRacism = d.Get(prefix32 + "hate_and_racism").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryPersonalStorage = d.Get(prefix32 + "personal_storage").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryViolence = d.Get(prefix32 + "violence").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryKeyloggersAndMonitoring = d.Get(prefix32 + "keyloggers_and_monitoring").(int)
	obj32.SlbTemplateClientSslInstanceWebCategorySearchEngines = d.Get(prefix32 + "search_engines").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryInternetPortals = d.Get(prefix32 + "internet_portals").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryWebAdvertisements = d.Get(prefix32 + "web_advertisements").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryCheating = d.Get(prefix32 + "cheating").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryGross = d.Get(prefix32 + "gross").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryWebBasedEmail = d.Get(prefix32 + "web_based_email").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryMalwareSites = d.Get(prefix32 + "malware_sites").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryPhishingAndOtherFraud = d.Get(prefix32 + "phishing_and_other_fraud").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryProxyAvoidAndAnonymizers = d.Get(prefix32 + "proxy_avoid_and_anonymizers").(int)
	obj32.SlbTemplateClientSslInstanceWebCategorySpywareAndAdware = d.Get(prefix32 + "spyware_and_adware").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryMusic = d.Get(prefix32 + "music").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryGovernment = d.Get(prefix32 + "government").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryNudity = d.Get(prefix32 + "nudity").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryNewsAndMedia = d.Get(prefix32 + "news_and_media").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryIllegal = d.Get(prefix32 + "illegal").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryCdns = d.Get(prefix32 + "cdns").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryInternetCommunications = d.Get(prefix32 + "internet_communications").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryBotNets = d.Get(prefix32 + "bot_nets").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryAbortion = d.Get(prefix32 + "abortion").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryHealthAndMedicine = d.Get(prefix32 + "health_and_medicine").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryConfirmedSpamSources = d.Get(prefix32 + "confirmed_spam_sources").(int)
	obj32.SlbTemplateClientSslInstanceWebCategorySpamUrls = d.Get(prefix32 + "spam_urls").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryUnconfirmedSpamSources = d.Get(prefix32 + "unconfirmed_spam_sources").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryOpenHTTPProxies = d.Get(prefix32 + "open_http_proxies").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryDynamicComment = d.Get(prefix32 + "dynamic_comment").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryParkedDomains = d.Get(prefix32 + "parked_domains").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryAlcoholAndTobacco = d.Get(prefix32 + "alcohol_and_tobacco").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryPrivateIPAddresses = d.Get(prefix32 + "private_ip_addresses").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryImageAndVideoSearch = d.Get(prefix32 + "image_and_video_search").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryFashionAndBeauty = d.Get(prefix32 + "fashion_and_beauty").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryRecreationAndHobbies = d.Get(prefix32 + "recreation_and_hobbies").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryMotorVehicles = d.Get(prefix32 + "motor_vehicles").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryWebHostingSites = d.Get(prefix32 + "web_hosting_sites").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryFoodAndDining = d.Get(prefix32 + "food_and_dining").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryNudityArtistic = d.Get(prefix32 + "nudity_artistic").(int)
	obj32.SlbTemplateClientSslInstanceWebCategoryIllegalPornography = d.Get(prefix32 + "illegal_pornography").(int)

	c.SlbTemplateClientSslInstanceWebCategoryUncategorized = obj32

	var obj33 go_thunder.SlbTemplateClientSslInstanceExceptionWebCategory
	prefix33 := "exception_web_category.0."
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionUncategorized = d.Get(prefix33 + "exception_uncategorized").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionRealEstate = d.Get(prefix33 + "exception_real_estate").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionComputerAndInternetSecurity = d.Get(prefix33 + "exception_computer_and_internet_security").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionFinancialServices = d.Get(prefix33 + "exception_financial_services").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionBusinessAndEconomy = d.Get(prefix33 + "exception_business_and_economy").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionComputerAndInternetInfo = d.Get(prefix33 + "exception_computer_and_internet_info").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionAuctions = d.Get(prefix33 + "exception_auctions").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionShopping = d.Get(prefix33 + "exception_shopping").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionCultAndOccult = d.Get(prefix33 + "exception_cult_and_occult").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionTravel = d.Get(prefix33 + "exception_travel").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionDrugs = d.Get(prefix33 + "exception_drugs").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionAdultAndPornography = d.Get(prefix33 + "exception_adult_and_pornography").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionHomeAndGarden = d.Get(prefix33 + "exception_home_and_garden").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionMilitary = d.Get(prefix33 + "exception_military").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSocialNetwork = d.Get(prefix33 + "exception_social_network").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionDeadSites = d.Get(prefix33 + "exception_dead_sites").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionStockAdviceAndTools = d.Get(prefix33 + "exception_stock_advice_and_tools").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionTrainingAndTools = d.Get(prefix33 + "exception_training_and_tools").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionDating = d.Get(prefix33 + "exception_dating").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSexEducation = d.Get(prefix33 + "exception_sex_education").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionReligion = d.Get(prefix33 + "exception_religion").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionEntertainmentAndArts = d.Get(prefix33 + "exception_entertainment_and_arts").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPersonalSitesAndBlogs = d.Get(prefix33 + "exception_personal_sites_and_blogs").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionLegal = d.Get(prefix33 + "exception_legal").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionLocalInformation = d.Get(prefix33 + "exception_local_information").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionStreamingMedia = d.Get(prefix33 + "exception_streaming_media").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionJobSearch = d.Get(prefix33 + "exception_job_search").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionGambling = d.Get(prefix33 + "exception_gambling").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionTranslation = d.Get(prefix33 + "exception_translation").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionReferenceAndResearch = d.Get(prefix33 + "exception_reference_and_research").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSharewareAndFreeware = d.Get(prefix33 + "exception_shareware_and_freeware").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPeerToPeer = d.Get(prefix33 + "exception_peer_to_peer").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionMarijuana = d.Get(prefix33 + "exception_marijuana").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionHacking = d.Get(prefix33 + "exception_hacking").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionGames = d.Get(prefix33 + "exception_games").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPhilosophyAndPolitics = d.Get(prefix33 + "exception_philosophy_and_politics").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionWeapons = d.Get(prefix33 + "exception_weapons").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPayToSurf = d.Get(prefix33 + "exception_pay_to_surf").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionHuntingAndFishing = d.Get(prefix33 + "exception_hunting_and_fishing").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSociety = d.Get(prefix33 + "exception_society").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionEducationalInstitutions = d.Get(prefix33 + "exception_educational_institutions").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionOnlineGreetingCards = d.Get(prefix33 + "exception_online_greeting_cards").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSports = d.Get(prefix33 + "exception_sports").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSwimsuitsAndIntimateApparel = d.Get(prefix33 + "exception_swimsuits_and_intimate_apparel").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionQuestionable = d.Get(prefix33 + "exception_questionable").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionKids = d.Get(prefix33 + "exception_kids").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionHateAndRacism = d.Get(prefix33 + "exception_hate_and_racism").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPersonalStorage = d.Get(prefix33 + "exception_personal_storage").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionViolence = d.Get(prefix33 + "exception_violence").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionKeyloggersAndMonitoring = d.Get(prefix33 + "exception_keyloggers_and_monitoring").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSearchEngines = d.Get(prefix33 + "exception_search_engines").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionInternetPortals = d.Get(prefix33 + "exception_internet_portals").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionWebAdvertisements = d.Get(prefix33 + "exception_web_advertisements").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionCheating = d.Get(prefix33 + "exception_cheating").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionGross = d.Get(prefix33 + "exception_gross").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionWebBasedEmail = d.Get(prefix33 + "exception_web_based_email").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionMalwareSites = d.Get(prefix33 + "exception_malware_sites").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPhishingAndOtherFraud = d.Get(prefix33 + "exception_phishing_and_other_fraud").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionProxyAvoidAndAnonymizers = d.Get(prefix33 + "exception_proxy_avoid_and_anonymizers").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSpywareAndAdware = d.Get(prefix33 + "exception_spyware_and_adware").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionMusic = d.Get(prefix33 + "exception_music").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionGovernment = d.Get(prefix33 + "exception_government").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionNudity = d.Get(prefix33 + "exception_nudity").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionNewsAndMedia = d.Get(prefix33 + "exception_news_and_media").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionIllegal = d.Get(prefix33 + "exception_illegal").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionCdns = d.Get(prefix33 + "exception_cdns").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionInternetCommunications = d.Get(prefix33 + "exception_internet_communications").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionBotNets = d.Get(prefix33 + "exception_bot_nets").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionAbortion = d.Get(prefix33 + "exception_abortion").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionHealthAndMedicine = d.Get(prefix33 + "exception_health_and_medicine").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionConfirmedSpamSources = d.Get(prefix33 + "exception_confirmed_spam_sources").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSpamUrls = d.Get(prefix33 + "exception_spam_urls").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionUnconfirmedSpamSources = d.Get(prefix33 + "exception_unconfirmed_spam_sources").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionOpenHTTPProxies = d.Get(prefix33 + "exception_open_http_proxies").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionDynamicComment = d.Get(prefix33 + "exception_dynamic_comment").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionParkedDomains = d.Get(prefix33 + "exception_parked_domains").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionAlcoholAndTobacco = d.Get(prefix33 + "exception_alcohol_and_tobacco").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPrivateIPAddresses = d.Get(prefix33 + "exception_private_ip_addresses").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionImageAndVideoSearch = d.Get(prefix33 + "exception_image_and_video_search").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionFashionAndBeauty = d.Get(prefix33 + "exception_fashion_and_beauty").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionRecreationAndHobbies = d.Get(prefix33 + "exception_recreation_and_hobbies").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionMotorVehicles = d.Get(prefix33 + "exception_motor_vehicles").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionWebHostingSites = d.Get(prefix33 + "exception_web_hosting_sites").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionFoodAndDining = d.Get(prefix33 + "exception_food_and_dining").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionNudityArtistic = d.Get(prefix33 + "exception_nudity_artistic").(int)
	obj33.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionIllegalPornography = d.Get(prefix33 + "exception_illegal_pornography").(int)

	c.SlbTemplateClientSslInstanceExceptionWebCategoryExceptionUncategorized = obj33

	c.SlbTemplateClientSslInstanceRequireWebCategory = d.Get("require_web_category").(int)

	SlbTemplateClientSslInstanceClientIpv4ListCount := d.Get("client_ipv4_list.#").(int)
	c.SlbTemplateClientSslInstanceClientIpv4ListClientIpv4ListName = make([]go_thunder.SlbTemplateClientSslInstanceClientIpv4List, 0, SlbTemplateClientSslInstanceClientIpv4ListCount)

	for i := 0; i < SlbTemplateClientSslInstanceClientIpv4ListCount; i++ {
		var obj34 go_thunder.SlbTemplateClientSslInstanceClientIpv4List
		prefix34 := fmt.Sprintf("client_ipv4_list.%d.", i)
		obj34.SlbTemplateClientSslInstanceClientIpv4ListClientIpv4ListName = d.Get(prefix34 + "client_ipv4_list_name").(string)
		c.SlbTemplateClientSslInstanceClientIpv4ListClientIpv4ListName = append(c.SlbTemplateClientSslInstanceClientIpv4ListClientIpv4ListName, obj34)
	}

	SlbTemplateClientSslInstanceClientIpv6ListCount := d.Get("client_ipv6_list.#").(int)
	c.SlbTemplateClientSslInstanceClientIpv6ListClientIpv6ListName = make([]go_thunder.SlbTemplateClientSslInstanceClientIpv6List, 0, SlbTemplateClientSslInstanceClientIpv6ListCount)

	for i := 0; i < SlbTemplateClientSslInstanceClientIpv6ListCount; i++ {
		var obj35 go_thunder.SlbTemplateClientSslInstanceClientIpv6List
		prefix35 := fmt.Sprintf("client_ipv6_list.%d.", i)
		obj35.SlbTemplateClientSslInstanceClientIpv6ListClientIpv6ListName = d.Get(prefix35 + "client_ipv6_list_name").(string)
		c.SlbTemplateClientSslInstanceClientIpv6ListClientIpv6ListName = append(c.SlbTemplateClientSslInstanceClientIpv6ListClientIpv6ListName, obj35)
	}

	SlbTemplateClientSslInstanceServerIpv4ListCount := d.Get("server_ipv4_list.#").(int)
	c.SlbTemplateClientSslInstanceServerIpv4ListServerIpv4ListName = make([]go_thunder.SlbTemplateClientSslInstanceServerIpv4List, 0, SlbTemplateClientSslInstanceServerIpv4ListCount)

	for i := 0; i < SlbTemplateClientSslInstanceServerIpv4ListCount; i++ {
		var obj36 go_thunder.SlbTemplateClientSslInstanceServerIpv4List
		prefix36 := fmt.Sprintf("server_ipv4_list.%d.", i)
		obj36.SlbTemplateClientSslInstanceServerIpv4ListServerIpv4ListName = d.Get(prefix36 + "server_ipv4_list_name").(string)
		c.SlbTemplateClientSslInstanceServerIpv4ListServerIpv4ListName = append(c.SlbTemplateClientSslInstanceServerIpv4ListServerIpv4ListName, obj36)
	}

	SlbTemplateClientSslInstanceServerIpv6ListCount := d.Get("server_ipv6_list.#").(int)
	c.SlbTemplateClientSslInstanceServerIpv6ListServerIpv6ListName = make([]go_thunder.SlbTemplateClientSslInstanceServerIpv6List, 0, SlbTemplateClientSslInstanceServerIpv6ListCount)

	for i := 0; i < SlbTemplateClientSslInstanceServerIpv6ListCount; i++ {
		var obj37 go_thunder.SlbTemplateClientSslInstanceServerIpv6List
		prefix37 := fmt.Sprintf("server_ipv6_list.%d.", i)
		obj37.SlbTemplateClientSslInstanceServerIpv6ListServerIpv6ListName = d.Get(prefix37 + "server_ipv6_list_name").(string)
		c.SlbTemplateClientSslInstanceServerIpv6ListServerIpv6ListName = append(c.SlbTemplateClientSslInstanceServerIpv6ListServerIpv6ListName, obj37)
	}

	SlbTemplateClientSslInstanceExceptionClientIpv4ListCount := d.Get("exception_client_ipv4_list.#").(int)
	c.SlbTemplateClientSslInstanceExceptionClientIpv4ListExceptionClientIpv4ListName = make([]go_thunder.SlbTemplateClientSslInstanceExceptionClientIpv4List, 0, SlbTemplateClientSslInstanceExceptionClientIpv4ListCount)

	for i := 0; i < SlbTemplateClientSslInstanceExceptionClientIpv4ListCount; i++ {
		var obj38 go_thunder.SlbTemplateClientSslInstanceExceptionClientIpv4List
		prefix38 := fmt.Sprintf("exception_client_ipv4_list.%d.", i)
		obj38.SlbTemplateClientSslInstanceExceptionClientIpv4ListExceptionClientIpv4ListName = d.Get(prefix38 + "exception_client_ipv4_list_name").(string)
		c.SlbTemplateClientSslInstanceExceptionClientIpv4ListExceptionClientIpv4ListName = append(c.SlbTemplateClientSslInstanceExceptionClientIpv4ListExceptionClientIpv4ListName, obj38)
	}

	SlbTemplateClientSslInstanceExceptionClientIpv6ListCount := d.Get("exception_client_ipv6_list.#").(int)
	c.SlbTemplateClientSslInstanceExceptionClientIpv6ListExceptionClientIpv6ListName = make([]go_thunder.SlbTemplateClientSslInstanceExceptionClientIpv6List, 0, SlbTemplateClientSslInstanceExceptionClientIpv6ListCount)

	for i := 0; i < SlbTemplateClientSslInstanceExceptionClientIpv6ListCount; i++ {
		var obj39 go_thunder.SlbTemplateClientSslInstanceExceptionClientIpv6List
		prefix39 := fmt.Sprintf("exception_client_ipv6_list.%d.", i)
		obj39.SlbTemplateClientSslInstanceExceptionClientIpv6ListExceptionClientIpv6ListName = d.Get(prefix39 + "exception_client_ipv6_list_name").(string)
		c.SlbTemplateClientSslInstanceExceptionClientIpv6ListExceptionClientIpv6ListName = append(c.SlbTemplateClientSslInstanceExceptionClientIpv6ListExceptionClientIpv6ListName, obj39)
	}

	SlbTemplateClientSslInstanceExceptionServerIpv4ListCount := d.Get("exception_server_ipv4_list.#").(int)
	c.SlbTemplateClientSslInstanceExceptionServerIpv4ListExceptionServerIpv4ListName = make([]go_thunder.SlbTemplateClientSslInstanceExceptionServerIpv4List, 0, SlbTemplateClientSslInstanceExceptionServerIpv4ListCount)

	for i := 0; i < SlbTemplateClientSslInstanceExceptionServerIpv4ListCount; i++ {
		var obj40 go_thunder.SlbTemplateClientSslInstanceExceptionServerIpv4List
		prefix40 := fmt.Sprintf("exception_server_ipv4_list.%d.", i)
		obj40.SlbTemplateClientSslInstanceExceptionServerIpv4ListExceptionServerIpv4ListName = d.Get(prefix40 + "exception_server_ipv4_list_name").(string)
		c.SlbTemplateClientSslInstanceExceptionServerIpv4ListExceptionServerIpv4ListName = append(c.SlbTemplateClientSslInstanceExceptionServerIpv4ListExceptionServerIpv4ListName, obj40)
	}

	SlbTemplateClientSslInstanceExceptionServerIpv6ListCount := d.Get("exception_server_ipv6_list.#").(int)
	c.SlbTemplateClientSslInstanceExceptionServerIpv6ListExceptionServerIpv6ListName = make([]go_thunder.SlbTemplateClientSslInstanceExceptionServerIpv6List, 0, SlbTemplateClientSslInstanceExceptionServerIpv6ListCount)

	for i := 0; i < SlbTemplateClientSslInstanceExceptionServerIpv6ListCount; i++ {
		var obj41 go_thunder.SlbTemplateClientSslInstanceExceptionServerIpv6List
		prefix41 := fmt.Sprintf("exception_server_ipv6_list.%d.", i)
		obj41.SlbTemplateClientSslInstanceExceptionServerIpv6ListExceptionServerIpv6ListName = d.Get(prefix41 + "exception_server_ipv6_list_name").(string)
		c.SlbTemplateClientSslInstanceExceptionServerIpv6ListExceptionServerIpv6ListName = append(c.SlbTemplateClientSslInstanceExceptionServerIpv6ListExceptionServerIpv6ListName, obj41)
	}

	var obj42 go_thunder.SlbTemplateClientSslInstanceLocalCertPinList
	prefix42 := "local_cert_pin_list.0."
	obj42.SlbTemplateClientSslInstanceLocalCertPinListLocalCertPinListBypassFailCount = d.Get(prefix42 + "local_cert_pin_list_bypass_fail_count").(int)

	c.SlbTemplateClientSslInstanceLocalCertPinListLocalCertPinListBypassFailCount = obj42

	c.SlbTemplateClientSslInstanceCentralCertPinList = d.Get("central_cert_pin_list").(int)
	c.SlbTemplateClientSslInstanceForwardProxyRequireSniCertMatched = d.Get("forward_proxy_require_sni_cert_matched").(string)
	c.SlbTemplateClientSslInstanceTemplateCipher = d.Get("template_cipher").(string)
	c.SlbTemplateClientSslInstanceSharedPartitionCipherTemplate = d.Get("shared_partition_cipher_template").(int)
	c.SlbTemplateClientSslInstanceTemplateCipherShared = d.Get("template_cipher_shared").(string)
	c.SlbTemplateClientSslInstanceTemplateHsm = d.Get("template_hsm").(string)
	c.SlbTemplateClientSslInstanceHsmType = d.Get("hsm_type").(string)

	SlbTemplateClientSslInstanceCipherWithoutPrioListCount := d.Get("cipher_without_prio_list.#").(int)
	c.SlbTemplateClientSslInstanceCipherWithoutPrioListCipherWoPrio = make([]go_thunder.SlbTemplateClientSslInstanceCipherWithoutPrioList, 0, SlbTemplateClientSslInstanceCipherWithoutPrioListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCipherWithoutPrioListCount; i++ {
		var obj43 go_thunder.SlbTemplateClientSslInstanceCipherWithoutPrioList
		prefix43 := fmt.Sprintf("cipher_without_prio_list.%d.", i)
		obj43.SlbTemplateClientSslInstanceCipherWithoutPrioListCipherWoPrio = d.Get(prefix43 + "cipher_wo_prio").(string)
		c.SlbTemplateClientSslInstanceCipherWithoutPrioListCipherWoPrio = append(c.SlbTemplateClientSslInstanceCipherWithoutPrioListCipherWoPrio, obj43)
	}

	SlbTemplateClientSslInstanceServerNameListCount := d.Get("server_name_list.#").(int)
	c.SlbTemplateClientSslInstanceServerNameListServerName = make([]go_thunder.SlbTemplateClientSslInstanceServerNameList, 0, SlbTemplateClientSslInstanceServerNameListCount)

	for i := 0; i < SlbTemplateClientSslInstanceServerNameListCount; i++ {
		var obj44 go_thunder.SlbTemplateClientSslInstanceServerNameList
		prefix44 := fmt.Sprintf("server_name_list.%d.", i)
		obj44.SlbTemplateClientSslInstanceServerNameListServerName = d.Get(prefix44 + "server_name").(string)
		obj44.SlbTemplateClientSslInstanceServerNameListServerCert = d.Get(prefix44 + "server_cert").(string)
		obj44.SlbTemplateClientSslInstanceServerNameListServerChain = d.Get(prefix44 + "server_chain").(string)
		obj44.SlbTemplateClientSslInstanceServerNameListServerKey = d.Get(prefix44 + "server_key").(string)
		obj44.SlbTemplateClientSslInstanceServerNameListServerPassphrase = d.Get(prefix44 + "server_passphrase").(string)
		obj44.SlbTemplateClientSslInstanceServerNameListServerNameAlternate = d.Get(prefix44 + "server_name_alternate").(int)
		obj44.SlbTemplateClientSslInstanceServerNameListServerShared = d.Get(prefix44 + "server_shared").(int)
		obj44.SlbTemplateClientSslInstanceServerNameListServerNameRegex = d.Get(prefix44 + "server_name_regex").(string)
		obj44.SlbTemplateClientSslInstanceServerNameListServerCertRegex = d.Get(prefix44 + "server_cert_regex").(string)
		obj44.SlbTemplateClientSslInstanceServerNameListServerChainRegex = d.Get(prefix44 + "server_chain_regex").(string)
		obj44.SlbTemplateClientSslInstanceServerNameListServerKeyRegex = d.Get(prefix44 + "server_key_regex").(string)
		obj44.SlbTemplateClientSslInstanceServerNameListServerPassphraseRegex = d.Get(prefix44 + "server_passphrase_regex").(string)
		obj44.SlbTemplateClientSslInstanceServerNameListServerNameRegexAlternate = d.Get(prefix44 + "server_name_regex_alternate").(int)
		obj44.SlbTemplateClientSslInstanceServerNameListServerSharedRegex = d.Get(prefix44 + "server_shared_regex").(int)
		c.SlbTemplateClientSslInstanceServerNameListServerName = append(c.SlbTemplateClientSslInstanceServerNameListServerName, obj44)
	}

	c.SlbTemplateClientSslInstanceServerNameAutoMap = d.Get("server_name_auto_map").(int)
	c.SlbTemplateClientSslInstanceSniEnableLog = d.Get("sni_enable_log").(int)
	c.SlbTemplateClientSslInstanceSniBypassMissingCert = d.Get("sni_bypass_missing_cert").(int)
	c.SlbTemplateClientSslInstanceSniBypassExpiredCert = d.Get("sni_bypass_expired_cert").(int)
	c.SlbTemplateClientSslInstanceSniBypassExplicitList = d.Get("sni_bypass_explicit_list").(string)
	c.SlbTemplateClientSslInstanceSniBypassEnableLog = d.Get("sni_bypass_enable_log").(int)
	c.SlbTemplateClientSslInstanceDirectClientServerAuth = d.Get("direct_client_server_auth").(int)
	c.SlbTemplateClientSslInstanceSessionCacheSize = d.Get("session_cache_size").(int)
	c.SlbTemplateClientSslInstanceSessionCacheTimeout = d.Get("session_cache_timeout").(int)
	c.SlbTemplateClientSslInstanceSessionTicketDisable = d.Get("session_ticket_disable").(int)
	c.SlbTemplateClientSslInstanceSessionTicketLifetime = d.Get("session_ticket_lifetime").(int)
	c.SlbTemplateClientSslInstanceSslFalseStartDisable = d.Get("ssl_false_start_disable").(int)
	c.SlbTemplateClientSslInstanceDisableSslv3 = d.Get("disable_sslv3").(int)
	c.SlbTemplateClientSslInstanceVersion = d.Get("version").(int)
	c.SlbTemplateClientSslInstanceDgversion = d.Get("dgversion").(int)
	c.SlbTemplateClientSslInstanceRenegotiationDisable = d.Get("renegotiation_disable").(int)
	c.SlbTemplateClientSslInstanceSslv2BypassServiceGroup = d.Get("sslv2_bypass_service_group").(string)
	c.SlbTemplateClientSslInstanceAuthorization = d.Get("authorization").(int)
	c.SlbTemplateClientSslInstanceAuthenName = d.Get("authen_name").(string)
	c.SlbTemplateClientSslInstanceLdapBaseDnFromCert = d.Get("ldap_base_dn_from_cert").(int)
	c.SlbTemplateClientSslInstanceLdapSearchFilter = d.Get("ldap_search_filter").(string)
	c.SlbTemplateClientSslInstanceAuthSg = d.Get("auth_sg").(string)
	c.SlbTemplateClientSslInstanceAuthSgDn = d.Get("auth_sg_dn").(int)
	c.SlbTemplateClientSslInstanceAuthSgFilter = d.Get("auth_sg_filter").(string)
	c.SlbTemplateClientSslInstanceAuthUsernameAttribute = d.Get("auth_username_attribute").(string)
	c.SlbTemplateClientSslInstanceNonSslBypassServiceGroup = d.Get("non_ssl_bypass_service_group").(string)
	c.SlbTemplateClientSslInstanceNonSslBypassL4Session = d.Get("non_ssl_bypass_l4session").(int)
	c.SlbTemplateClientSslInstanceEnableSsliFtpAlg = d.Get("enable_ssli_ftp_alg").(int)
	c.SlbTemplateClientSslInstanceEarlyData = d.Get("early_data").(int)
	c.SlbTemplateClientSslInstanceNoAntiReplay = d.Get("no_anti_replay").(int)
	c.SlbTemplateClientSslInstanceJa3Enable = d.Get("ja3_enable").(int)
	c.SlbTemplateClientSslInstanceJa3InsertHTTPHeader = d.Get("ja3_insert_http_header").(string)
	c.SlbTemplateClientSslInstanceJa3RejectClassList = d.Get("ja3_reject_class_list").(string)
	c.SlbTemplateClientSslInstanceJa3RejectMaxNumberPerHost = d.Get("ja3_reject_max_number_per_host").(int)
	c.SlbTemplateClientSslInstanceJa3TTL = d.Get("ja3_ttl").(int)
	c.SlbTemplateClientSslInstanceUserTag = d.Get("user_tag").(string)

	SlbTemplateClientSslInstanceCertificateListCount := d.Get("certificate_list.#").(int)
	c.SlbTemplateClientSslInstanceCertificateListCert = make([]go_thunder.SlbTemplateClientSslInstanceCertificateList, 0, SlbTemplateClientSslInstanceCertificateListCount)

	for i := 0; i < SlbTemplateClientSslInstanceCertificateListCount; i++ {
		var obj45 go_thunder.SlbTemplateClientSslInstanceCertificateList
		prefix45 := fmt.Sprintf("certificate_list.%d.", i)
		obj45.SlbTemplateClientSslInstanceCertificateListCert = d.Get(prefix45 + "cert").(string)
		obj45.SlbTemplateClientSslInstanceCertificateListKey = d.Get(prefix45 + "key").(string)
		obj45.SlbTemplateClientSslInstanceCertificateListPassphrase = d.Get(prefix45 + "passphrase").(string)
		obj45.SlbTemplateClientSslInstanceCertificateListChainCert = d.Get(prefix45 + "chain_cert").(string)
		obj45.SlbTemplateClientSslInstanceCertificateListShared = d.Get(prefix45 + "shared").(int)
		c.SlbTemplateClientSslInstanceCertificateListCert = append(c.SlbTemplateClientSslInstanceCertificateListCert, obj45)
	}

	vc.SlbTemplateClientSslInstanceName = c
	return vc
}
