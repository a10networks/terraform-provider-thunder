package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateClientSsl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_client_ssl`: Client SSL Template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateClientSslCreate,
		UpdateContext: resourceSlbTemplateClientSslUpdate,
		ReadContext:   resourceSlbTemplateClientSslRead,
		DeleteContext: resourceSlbTemplateClientSslDelete,

		Schema: map[string]*schema.Schema{
			"ad_group_list": {
				Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if ad-group matches class-list",
			},
			"alert_type": {
				Type: schema.TypeString, Optional: true, Description: "'fatal': Log fatal alerts;",
			},
			"auth_sg": {
				Type: schema.TypeString, Optional: true, Description: "Specify authorization LDAP service group",
			},
			"auth_sg_dn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Subject DN as LDAP search base DN",
			},
			"auth_sg_filter": {
				Type: schema.TypeString, Optional: true, Description: "Specify LDAP search filter",
			},
			"auth_username": {
				Type: schema.TypeString, Optional: true, Description: "Specify the Username Field in the Client Certificate(If multi-fields are specificed, prior one has higher priority)",
			},
			"auth_username_attribute": {
				Type: schema.TypeString, Optional: true, Description: "Specify attribute name of username for client SSL authorization",
			},
			"authen_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify authorization LDAP server name",
			},
			"authorization": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify LDAP server for client SSL authorizaiton",
			},
			"bypass_cert_issuer_class_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Class List Name",
			},
			"bypass_cert_issuer_multi_class_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_cert_issuer_multi_class_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Class List Name",
						},
					},
				},
			},
			"bypass_cert_san_class_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Class List Name",
			},
			"bypass_cert_san_multi_class_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_cert_san_multi_class_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Class List Name",
						},
					},
				},
			},
			"bypass_cert_subject_class_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Class List Name",
			},
			"bypass_cert_subject_multi_class_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_cert_subject_multi_class_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Class List Name",
						},
					},
				},
			},
			"ca_certs": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_cert": {
							Type: schema.TypeString, Optional: true, Description: "CA Certificate (CA Certificate Name)",
						},
						"ca_shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "CA Certificate Partition Shared",
						},
						"client_ocsp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify ocsp authentication server(s) for client certificate verification",
						},
						"client_ocsp_srvr": {
							Type: schema.TypeString, Optional: true, Description: "Specify authentication server",
						},
						"client_ocsp_sg": {
							Type: schema.TypeString, Optional: true, Description: "Specify service-group (Service group name)",
						},
					},
				},
			},
			"cache_persistence_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Class List Name",
			},
			"case_insensitive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Case insensitive forward proxy bypass",
			},
			"central_cert_pin_list": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward proxy bypass if SNI string is contained in central updated cert-pinning-candidate list",
			},
			"cert_revoke_action": {
				Type: schema.TypeString, Optional: true, Default: "bypass", Description: "'bypass': bypass SSLi processing; 'continue': continue the connection; 'drop': close the connection; 'block': block the connection with a warning page;",
			},
			"cert_unknown_action": {
				Type: schema.TypeString, Optional: true, Default: "bypass", Description: "'bypass': bypass SSLi processing; 'continue': continue the connection; 'drop': close the connection; 'block': block the connection with a warning page;",
			},
			"certificate_issuer_contains_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_issuer_contains": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if Certificate  issuer contains another string (Certificate issuer)",
						},
					},
				},
			},
			"certificate_issuer_ends_with_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_issuer_ends_with": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if Certificate issuer ends with another string",
						},
					},
				},
			},
			"certificate_issuer_equals_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_issuer_equals": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if Certificate issuer equals another string",
						},
					},
				},
			},
			"certificate_issuer_starts_with_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_issuer_starts": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if Certificate issuer starts with another string",
						},
					},
				},
			},
			"certificate_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert": {
							Type: schema.TypeString, Required: true, Description: "Certificate Name",
						},
						"key": {
							Type: schema.TypeString, Optional: true, Description: "Server Private Key (Key Name)",
						},
						"passphrase": {
							Type: schema.TypeString, Optional: true, Description: "Password Phrase",
						},
						"chain_cert": {
							Type: schema.TypeString, Optional: true, Description: "Chain Certificate (Chain Certificate Name)",
						},
						"shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Server Certificate and Key Partition Shared",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"certificate_san_contains_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_san_contains": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if Certificate SAN contains another string",
						},
					},
				},
			},
			"certificate_san_ends_with_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_san_ends_with": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if Certificate SAN ends with another string",
						},
					},
				},
			},
			"certificate_san_equals_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_san_equals": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if Certificate SAN equals another string",
						},
					},
				},
			},
			"certificate_san_starts_with_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_san_starts": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if Certificate SAN starts with another string",
						},
					},
				},
			},
			"certificate_subject_contains_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_subject_contains": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if Certificate Subject contains another string",
						},
					},
				},
			},
			"certificate_subject_ends_with_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_subject_ends_with": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if Certificate Subject ends with another string",
						},
					},
				},
			},
			"certificate_subject_equals_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_subject_equals": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if Certificate Subject equals another string",
						},
					},
				},
			},
			"certificate_subject_starts_with_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_subject_starts": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if Certificate Subject starts with another string",
						},
					},
				},
			},
			"chain_cert": {
				Type: schema.TypeString, Optional: true, Description: "Chain Certificate Name",
			},
			"chain_cert_shared_str": {
				Type: schema.TypeString, Optional: true, Description: "Chain Certificate Name",
			},
			"cipher_without_prio_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cipher_wo_prio": {
							Type: schema.TypeString, Optional: true, Description: "'SSL3_RSA_DES_192_CBC3_SHA': TLS_RSA_WITH_3DES_EDE_CBC_SHA (0x000A); 'SSL3_RSA_RC4_128_MD5': TLS_RSA_WITH_RC4_128_MD5 (0x0004); 'SSL3_RSA_RC4_128_SHA': TLS_RSA_WITH_RC4_128_SHA (0x0005); 'TLS1_RSA_AES_128_SHA': TLS_RSA_WITH_AES_128_CBC_SHA (0x002F); 'TLS1_RSA_AES_256_SHA': TLS_RSA_WITH_AES_256_CBC_SHA (0x0035); 'TLS1_RSA_AES_128_SHA256': TLS_RSA_WITH_AES_128_CBC_SHA256 (0x003C); 'TLS1_RSA_AES_256_SHA256': TLS_RSA_WITH_AES_256_CBC_SHA256 (0x003D); 'TLS1_DHE_RSA_AES_128_GCM_SHA256': TLS_DHE_RSA_WITH_AES_128_GCM_SHA256 (0x009E); 'TLS1_DHE_RSA_AES_128_SHA': TLS_DHE_RSA_WITH_AES_128_CBC_SHA (0x0033); 'TLS1_DHE_RSA_AES_128_SHA256': TLS_DHE_RSA_WITH_AES_128_CBC_SHA256 (0x0067); 'TLS1_DHE_RSA_AES_256_GCM_SHA384': TLS_DHE_RSA_WITH_AES_256_GCM_SHA384 (0x009F); 'TLS1_DHE_RSA_AES_256_SHA': TLS_DHE_RSA_WITH_AES_256_CBC_SHA (0x0039); 'TLS1_DHE_RSA_AES_256_SHA256': TLS_DHE_RSA_WITH_AES_256_CBC_SHA256 (0x006B); 'TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256': TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 (0xC02B); 'TLS1_ECDHE_ECDSA_AES_128_SHA': TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA (0xC009); 'TLS1_ECDHE_ECDSA_AES_128_SHA256': TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256 (0xC023); 'TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384': TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 (0xC02C); 'TLS1_ECDHE_ECDSA_AES_256_SHA': TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA (0xC00A); 'TLS1_ECDHE_RSA_AES_128_GCM_SHA256': TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 (0xC02F); 'TLS1_ECDHE_RSA_AES_128_SHA': TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA (0xC013); 'TLS1_ECDHE_RSA_AES_128_SHA256': TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256 (0xC027); 'TLS1_ECDHE_RSA_AES_256_GCM_SHA384': TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384 (0xC030); 'TLS1_ECDHE_RSA_AES_256_SHA': TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA (0xC014); 'TLS1_RSA_AES_128_GCM_SHA256': TLS_RSA_WITH_AES_128_GCM_SHA256 (0x009C); 'TLS1_RSA_AES_256_GCM_SHA384': TLS_RSA_WITH_AES_256_GCM_SHA384 (0x009D); 'TLS1_ECDHE_RSA_AES_256_SHA384': TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384 (0xC028); 'TLS1_ECDHE_ECDSA_AES_256_SHA384': TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384 (0xC024); 'TLS1_ECDHE_RSA_CHACHA20_POLY1305_SHA256': TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256 (0xCCA8); 'TLS1_ECDHE_ECDSA_CHACHA20_POLY1305_SHA256': TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 (0xCCA9); 'TLS1_DHE_RSA_CHACHA20_POLY1305_SHA256': TLS_DHE_RSA_WITH_CHACHA20_POLY1305_SHA256 (0xCCAA);",
						},
					},
				},
			},
			"class_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Class List Name",
			},
			"client_auth_case_insensitive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Case insensitive forward proxy client auth bypass",
			},
			"client_auth_class_list": {
				Type: schema.TypeString, Optional: true, Description: "Forward proxy client auth bypass if SNI string matches class-list (Class List Name)",
			},
			"client_auth_contains_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_auth_contains": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if SNI string contains another string",
						},
					},
				},
			},
			"client_auth_ends_with_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_auth_ends_with": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if SNI string ends with another string",
						},
					},
				},
			},
			"client_auth_equals_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_auth_equals": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if SNI string equals another string",
						},
					},
				},
			},
			"client_auth_starts_with_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_auth_starts_with": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if SNI string starts with another string",
						},
					},
				},
			},
			"client_certificate": {
				Type: schema.TypeString, Optional: true, Default: "Ignore", Description: "'Ignore': Don't request client certificate; 'Require': Require client certificate; 'Request': Request client certificate;",
			},
			"client_ipv4_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_ipv4_list_name": {
							Type: schema.TypeString, Optional: true, Description: "IPV4 client class-list name",
						},
					},
				},
			},
			"client_ipv6_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_ipv6_list_name": {
							Type: schema.TypeString, Optional: true, Description: "IPV6 client class-list name",
						},
					},
				},
			},
			"close_notify": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send close notification when terminate connection",
			},
			"contains_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"contains": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if SNI string contains another string",
						},
					},
				},
			},
			"crl_certs": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crl": {
							Type: schema.TypeString, Optional: true, Description: "Certificate Revocation Lists (Certificate Revocation Lists file name)",
						},
						"crl_shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Certificate Revocation Lists Partition Shared",
						},
					},
				},
			},
			"dgversion": {
				Type: schema.TypeInt, Optional: true, Default: 31, Description: "Lower TLS/SSL version can be downgraded",
			},
			"dh_type": {
				Type: schema.TypeString, Optional: true, Description: "'1024': 1024; '1024-dsa': 1024-dsa; '2048': 2048;",
			},
			"direct_client_server_auth": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Let backend server does SSL client authentication directly",
			},
			"disable_sslv3": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reject Client requests for SSL version 3",
			},
			"early_data": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable TLS 1.3 early data (0-RTT)",
			},
			"ec_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ec": {
							Type: schema.TypeString, Optional: true, Description: "'secp256r1': X9_62_prime256v1; 'secp384r1': secp384r1;",
						},
					},
				},
			},
			"enable_ssli_ftp_alg": {
				Type: schema.TypeInt, Optional: true, Description: "Enable SSLi FTP over TLS support at which port",
			},
			"enable_tls_alert_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable TLS alert logging",
			},
			"ends_with_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ends_with": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if SNI string ends with another string",
						},
					},
				},
			},
			"equals_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"equals": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if SNI string equals another string",
						},
					},
				},
			},
			"exception_ad_group_list": {
				Type: schema.TypeString, Optional: true, Description: "Exceptions to forward proxy bypass if ad-group matches class-list",
			},
			"exception_certificate_issuer_cl_name": {
				Type: schema.TypeString, Optional: true, Description: "Exceptions to forward-proxy-bypass",
			},
			"exception_certificate_san_cl_name": {
				Type: schema.TypeString, Optional: true, Description: "Exceptions to forward-proxy-bypass",
			},
			"exception_certificate_subject_cl_name": {
				Type: schema.TypeString, Optional: true, Description: "Exceptions to forward-proxy-bypass",
			},
			"exception_client_ipv4_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_client_ipv4_list_name": {
							Type: schema.TypeString, Optional: true, Description: "IPV4 exception client class-list name",
						},
					},
				},
			},
			"exception_client_ipv6_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_client_ipv6_list_name": {
							Type: schema.TypeString, Optional: true, Description: "IPV6 exception client class-list name",
						},
					},
				},
			},
			"exception_server_ipv4_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_server_ipv4_list_name": {
							Type: schema.TypeString, Optional: true, Description: "IPV4 exception server class-list name",
						},
					},
				},
			},
			"exception_server_ipv6_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_server_ipv6_list_name": {
							Type: schema.TypeString, Optional: true, Description: "IPV6 exception server class-list name",
						},
					},
				},
			},
			"exception_sni_cl_name": {
				Type: schema.TypeString, Optional: true, Description: "Exceptions to forward-proxy-bypass",
			},
			"exception_user_name_list": {
				Type: schema.TypeString, Optional: true, Description: "Exceptions to forward proxy bypass if user-name matches class-list",
			},
			"exception_web_category": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_uncategorized": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Uncategorized URLs",
						},
						"exception_real_estate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Real Estate",
						},
						"exception_computer_and_internet_security": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Computer and Internet Security",
						},
						"exception_financial_services": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Financial Services",
						},
						"exception_business_and_economy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Business and Economy",
						},
						"exception_computer_and_internet_info": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Computer and Internet Info",
						},
						"exception_auctions": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Auctions",
						},
						"exception_shopping": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Shopping",
						},
						"exception_cult_and_occult": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Cult and Occult",
						},
						"exception_travel": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Travel",
						},
						"exception_drugs": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Abused Drugs",
						},
						"exception_adult_and_pornography": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Adult and Pornography",
						},
						"exception_home_and_garden": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Home and Garden",
						},
						"exception_military": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Military",
						},
						"exception_social_network": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Social Network",
						},
						"exception_dead_sites": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Dead Sites (db Ops only)",
						},
						"exception_stock_advice_and_tools": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Stock Advice and Tools",
						},
						"exception_training_and_tools": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Training and Tools",
						},
						"exception_dating": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Dating",
						},
						"exception_sex_education": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Sex Education",
						},
						"exception_religion": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Religion",
						},
						"exception_entertainment_and_arts": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Entertainment and Arts",
						},
						"exception_personal_sites_and_blogs": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Personal sites and Blogs",
						},
						"exception_legal": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Legal",
						},
						"exception_local_information": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Local Information",
						},
						"exception_streaming_media": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Streaming Media",
						},
						"exception_job_search": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Job Search",
						},
						"exception_gambling": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Gambling",
						},
						"exception_translation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Translation",
						},
						"exception_reference_and_research": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Reference and Research",
						},
						"exception_shareware_and_freeware": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Shareware and Freeware",
						},
						"exception_peer_to_peer": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Peer to Peer",
						},
						"exception_marijuana": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Marijuana",
						},
						"exception_hacking": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Hacking",
						},
						"exception_games": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Games",
						},
						"exception_philosophy_and_politics": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Philosophy and Political Advocacy",
						},
						"exception_weapons": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Weapons",
						},
						"exception_pay_to_surf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Pay to Surf",
						},
						"exception_hunting_and_fishing": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Hunting and Fishing",
						},
						"exception_society": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Society",
						},
						"exception_educational_institutions": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Educational Institutions",
						},
						"exception_online_greeting_cards": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Online Greeting cards",
						},
						"exception_sports": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Sports",
						},
						"exception_swimsuits_and_intimate_apparel": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Swimsuits and Intimate Apparel",
						},
						"exception_questionable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Questionable",
						},
						"exception_kids": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Kids",
						},
						"exception_hate_and_racism": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Hate and Racism",
						},
						"exception_personal_storage": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Personal Storage",
						},
						"exception_violence": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Violence",
						},
						"exception_keyloggers_and_monitoring": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Keyloggers and Monitoring",
						},
						"exception_search_engines": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Search Engines",
						},
						"exception_internet_portals": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Internet Portals",
						},
						"exception_web_advertisements": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Web Advertisements",
						},
						"exception_cheating": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Cheating",
						},
						"exception_gross": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Gross",
						},
						"exception_web_based_email": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Web based email",
						},
						"exception_malware_sites": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Malware Sites",
						},
						"exception_phishing_and_other_fraud": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Phishing and Other Frauds",
						},
						"exception_proxy_avoid_and_anonymizers": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Proxy Avoid and Anonymizers",
						},
						"exception_spyware_and_adware": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Spyware and Adware",
						},
						"exception_music": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Music",
						},
						"exception_government": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Government",
						},
						"exception_nudity": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Nudity",
						},
						"exception_news_and_media": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category News and Media",
						},
						"exception_illegal": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Illegal",
						},
						"exception_cdns": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category CDNs",
						},
						"exception_internet_communications": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Internet Communications",
						},
						"exception_bot_nets": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Bot Nets",
						},
						"exception_abortion": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Abortion",
						},
						"exception_health_and_medicine": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Health and Medicine",
						},
						"exception_spam_urls": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category SPAM URLs",
						},
						"exception_dynamically_generated_content": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dynamically Generated Content",
						},
						"exception_parked_domains": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Parked Domains",
						},
						"exception_alcohol_and_tobacco": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Alcohol and Tobacco",
						},
						"exception_image_and_video_search": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Image and Video Search",
						},
						"exception_fashion_and_beauty": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Fashion and Beauty",
						},
						"exception_recreation_and_hobbies": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Recreation and Hobbies",
						},
						"exception_motor_vehicles": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Motor Vehicles",
						},
						"exception_web_hosting_sites": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Web Hosting Sites",
						},
						"exception_nudity_artistic": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Nudity join Entertainment and Arts",
						},
						"exception_illegal_pornography": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Illegal join Adult and Pornography",
						},
					},
				},
			},
			"exception_web_reputation": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_trustworthy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Intercept when reputation score is less than or equal to 100",
						},
						"exception_low_risk": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Intercept when reputation score is less than or equal to 80",
						},
						"exception_moderate_risk": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Intercept when reputation score is less than or equal to 60",
						},
						"exception_suspicious": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Intercept when reputation score is less than or equal to 40",
						},
						"exception_malicious": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Intercept when reputation score is less than or equal to 20",
						},
						"exception_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Intercept when reputation score is less than or equal to a customized value (1-100)",
						},
					},
				},
			},
			"expire_hours": {
				Type: schema.TypeInt, Optional: true, Description: "Certificate lifetime in hours",
			},
			"forward_passphrase": {
				Type: schema.TypeString, Optional: true, Description: "Password Phrase",
			},
			"forward_proxy_alt_sign": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward proxy alternate signing cert and key",
			},
			"forward_proxy_block_message": {
				Type: schema.TypeString, Optional: true, Description: "Message to be included on the block page (Message, enclose in quotes if spaces are present)",
			},
			"forward_proxy_ca_cert": {
				Type: schema.TypeString, Optional: true, Description: "CA Certificate for forward proxy (SSL forward proxy CA Certificate Name)",
			},
			"forward_proxy_ca_key": {
				Type: schema.TypeString, Optional: true, Description: "CA Private Key for forward proxy (SSL forward proxy CA Key Name)",
			},
			"forward_proxy_cert_cache_limit": {
				Type: schema.TypeInt, Optional: true, Default: 524288, Description: "Certificate cache size limit, default is 524288 (set to 0 for unlimited size)",
			},
			"forward_proxy_cert_cache_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 3600, Description: "Certificate cache timeout, default is 1 hour (seconds, set to 0 for never timeout)",
			},
			"forward_proxy_cert_expiry": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Adjust certificate expiry relative to the time when it is created on the device",
			},
			"forward_proxy_cert_not_ready_action": {
				Type: schema.TypeString, Optional: true, Default: "bypass", Description: "'bypass': bypass the connection; 'reset': reset the connection; 'intercept': wait for cert and then inspect the connection;",
			},
			"forward_proxy_cert_revoke_action": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Action taken if a certificate is irreversibly revoked, bypass SSLi processing by default",
			},
			"forward_proxy_cert_unknown_action": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Action taken if a certificate revocation status is unknown, bypass SSLi processing by default",
			},
			"forward_proxy_crl_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Certificate Revocation List checking for forward proxy",
			},
			"forward_proxy_decrypted_dscp": {
				Type: schema.TypeInt, Optional: true, Description: "Apply a DSCP to decrypted and bypassed traffic (DSCP to apply to decrypted traffic)",
			},
			"forward_proxy_decrypted_dscp_bypass": {
				Type: schema.TypeInt, Optional: true, Description: "DSCP to apply to bypassed traffic",
			},
			"forward_proxy_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL forward proxy",
			},
			"forward_proxy_esni_action": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Action taken if receiving encrypted server name indication extension in client hello MSG, bypass the connection by default",
			},
			"forward_proxy_failsafe_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Failsafe for SSL forward proxy",
			},
			"forward_proxy_hash_persistence_interval": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Set the time interval to save the hash persistence certs (Interval value, in minutes)",
			},
			"forward_proxy_log_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable SSL forward proxy logging",
			},
			"forward_proxy_no_shared_cipher_action": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Action taken if handshake fails due to no shared ciper, close the connection by default",
			},
			"forward_proxy_no_sni_action": {
				Type: schema.TypeString, Optional: true, Default: "intercept", Description: "'intercept': intercept in no SNI case; 'bypass': bypass in no SNI case; 'reset': reset in no SNI case;",
			},
			"forward_proxy_ocsp_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable ocsp-stapling for forward proxy",
			},
			"forward_proxy_require_sni_cert_matched": {
				Type: schema.TypeString, Optional: true, Description: "'no-match-action-inspect': Inspected if not matched; 'no-match-action-drop': Dropped if not matched;",
			},
			"forward_proxy_selfsign_redir": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Redirect connections to pages with self signed certs to a warning page",
			},
			"forward_proxy_ssl_version": {
				Type: schema.TypeInt, Optional: true, Default: 33, Description: "TLS/SSL version, default is TLS1.2 (TLS/SSL version: 31-TLSv1.0, 32-TLSv1.1, 33-TLSv1.2 and 34-TLSv1.3)",
			},
			"forward_proxy_trusted_ca_lists": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"forward_proxy_trusted_ca": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy trusted CA file (CA file name)",
						},
						"fp_trusted_ca_shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Trusted CA Certificate Partition Shared",
						},
					},
				},
			},
			"forward_proxy_verify_cert_fail_action": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Action taken if certificate verification fails, close the connection by default",
			},
			"fp_alt_cert": {
				Type: schema.TypeString, Optional: true, Description: "CA Certificate for forward proxy alternate signing (Certificate name)",
			},
			"fp_alt_chain_cert": {
				Type: schema.TypeString, Optional: true, Description: "Chain Certificate (Chain Certificate Name)",
			},
			"fp_alt_key": {
				Type: schema.TypeString, Optional: true, Description: "CA Private Key for forward proxy alternate signing (Key name)",
			},
			"fp_alt_passphrase": {
				Type: schema.TypeString, Optional: true, Description: "Password Phrase",
			},
			"fp_alt_shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Alternate CA Certificate and Private Key Partition Shared",
			},
			"fp_ca_certificate": {
				Type: schema.TypeString, Optional: true, Description: "CA Certificate for forward proxy (SSL forward proxy CA Certificate Name)",
			},
			"fp_ca_certificate_shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "CA Private Key Partition Shared",
			},
			"fp_ca_chain_cert": {
				Type: schema.TypeString, Optional: true, Description: "Chain Certificate (Chain Certificate Name)",
			},
			"fp_ca_key": {
				Type: schema.TypeString, Optional: true, Description: "CA Private Key for forward proxy (SSL forward proxy CA Key Name)",
			},
			"fp_ca_key_passphrase": {
				Type: schema.TypeString, Optional: true, Description: "Password Phrase",
			},
			"fp_ca_key_shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "CA Private Key Partition Shared",
			},
			"fp_ca_shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "CA Certificate Partition Shared",
			},
			"fp_cert_ext_aia_ca_issuers": {
				Type: schema.TypeString, Optional: true, Description: "CA Issuers (Authority Information Access URI)",
			},
			"fp_cert_ext_aia_ocsp": {
				Type: schema.TypeString, Optional: true, Description: "OCSP (Authority Information Access URI)",
			},
			"fp_cert_ext_crldp": {
				Type: schema.TypeString, Optional: true, Description: "CRL Distribution Point (CRL Distribution Point URI)",
			},
			"fp_cert_fetch_autonat": {
				Type: schema.TypeString, Optional: true, Description: "'auto': Configure auto NAT for server certificate fetching;",
			},
			"fp_cert_fetch_autonat_precedence": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set this NAT pool as higher precedence than other source NAT like configued under template policy",
			},
			"fp_cert_fetch_natpool_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify NAT pool or pool group",
			},
			"fp_cert_fetch_natpool_name_shared": {
				Type: schema.TypeString, Optional: true, Description: "Specify NAT pool or pool group",
			},
			"fp_cert_fetch_natpool_precedence": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set this NAT pool as higher precedence than other source NAT like configued under template policy",
			},
			"fp_esni_action": {
				Type: schema.TypeString, Optional: true, Default: "bypass", Description: "'bypass': bypass SSLi processing; 'drop': close the connection;",
			},
			"handshake_logging_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL handshake logging",
			},
			"hsm_type": {
				Type: schema.TypeString, Optional: true, Description: "'thales-embed': Thales embed key; 'thales-hwcrhk': Thales hwcrhk Key;",
			},
			"inspect_certificate_issuer_cl_name": {
				Type: schema.TypeString, Optional: true, Description: "Forward proxy Inspect if Certificate issuer matches class-list",
			},
			"inspect_certificate_san_cl_name": {
				Type: schema.TypeString, Optional: true, Description: "Forward proxy Inspect if Certificate Subject Alternative Name matches class-list",
			},
			"inspect_certificate_subject_cl_name": {
				Type: schema.TypeString, Optional: true, Description: "Forward proxy Inspect if Certificate Subject matches class-list",
			},
			"inspect_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Class List Name",
			},
			"ja3_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable JA3 features",
			},
			"ja3_insert_http_header": {
				Type: schema.TypeString, Optional: true, Description: "Insert the JA3 hash into this request as a HTTP header (HTTP Header Name)",
			},
			"ja3_reject_class_list": {
				Type: schema.TypeString, Optional: true, Description: "Drop request if the JA3 hash matches this class-list (type string-case-insensitive) (Class-List Name)",
			},
			"ja3_reject_max_number_per_host": {
				Type: schema.TypeInt, Optional: true, Description: "Drop request if numbers of JA3 of this client address exceeded",
			},
			"ja3_ttl": {
				Type: schema.TypeInt, Optional: true, Default: 600, Description: "seconds to keep each JA3 record",
			},
			"ldap_base_dn_from_cert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Subject DN as LDAP search base DN",
			},
			"ldap_search_filter": {
				Type: schema.TypeString, Optional: true, Description: "Specify LDAP search filter",
			},
			"local_cert_pin_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_cert_pin_list_bypass_fail_count": {
							Type: schema.TypeInt, Optional: true, Description: "Set the connection fail count as bypass criteria (Bypass when connection failure count is greater than the criteria (1-65536))",
						},
					},
				},
			},
			"local_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable local logging",
			},
			"multi_class_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"multi_clist_name": {
							Type: schema.TypeString, Optional: true, Description: "Class List Name",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Client SSL Template Name",
			},
			"no_anti_replay": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable anti-replay protection for TLS 1.3 early data (0-RTT data)",
			},
			"no_shared_cipher_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'bypass': bypass SSLi processing; 'drop': close the connection;",
			},
			"non_ssl_bypass_l4session": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Handle the non-ssl session as L4 for performance optimization",
			},
			"non_ssl_bypass_service_group": {
				Type: schema.TypeString, Optional: true, Description: "Service Group for Bypass non-ssl traffic (Service Group Name)",
			},
			"notafter": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "notAfter date",
			},
			"notafterday": {
				Type: schema.TypeInt, Optional: true, Description: "Day",
			},
			"notaftermonth": {
				Type: schema.TypeInt, Optional: true, Description: "Month",
			},
			"notafteryear": {
				Type: schema.TypeInt, Optional: true, Description: "Year",
			},
			"notbefore": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "notBefore date",
			},
			"notbeforeday": {
				Type: schema.TypeInt, Optional: true, Description: "Day",
			},
			"notbeforemonth": {
				Type: schema.TypeInt, Optional: true, Description: "Month",
			},
			"notbeforeyear": {
				Type: schema.TypeInt, Optional: true, Description: "Year",
			},
			"ocsp_stapling": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Config OCSP stapling support",
			},
			"ocspst_ca_cert": {
				Type: schema.TypeString, Optional: true, Description: "CA certificate",
			},
			"ocspst_ocsp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OCSP Authentication",
			},
			"ocspst_sg": {
				Type: schema.TypeString, Optional: true, Description: "Specify authentication service group",
			},
			"ocspst_sg_days": {
				Type: schema.TypeInt, Optional: true, Description: "Specify update period, in days",
			},
			"ocspst_sg_hours": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify update period, in hours",
			},
			"ocspst_sg_minutes": {
				Type: schema.TypeInt, Optional: true, Description: "Specify update period, in minutes",
			},
			"ocspst_sg_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Specify retry timeout (Default is 30 mins)",
			},
			"ocspst_srvr": {
				Type: schema.TypeString, Optional: true, Description: "Specify OCSP authentication server",
			},
			"ocspst_srvr_days": {
				Type: schema.TypeInt, Optional: true, Description: "Specify update period, in days",
			},
			"ocspst_srvr_hours": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify update period, in hours",
			},
			"ocspst_srvr_minutes": {
				Type: schema.TypeInt, Optional: true, Description: "Specify update period, in minutes",
			},
			"ocspst_srvr_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Specify retry timeout (Default is 30 mins)",
			},
			"renegotiation_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable SSL renegotiation",
			},
			"req_ca_lists": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_certificate_request_ca": {
							Type: schema.TypeString, Optional: true, Description: "Send CA lists in certificate request (CA Certificate Name)",
						},
						"client_cert_req_ca_shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "CA Certificate Partition Shared",
						},
					},
				},
			},
			"require_web_category": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Wait for web category to be resolved before taking bypass decision",
			},
			"server_ipv4_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_ipv4_list_name": {
							Type: schema.TypeString, Optional: true, Description: "IPV4 server class-list name",
						},
					},
				},
			},
			"server_ipv6_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_ipv6_list_name": {
							Type: schema.TypeString, Optional: true, Description: "IPV6 server class-list name",
						},
					},
				},
			},
			"server_name_auto_map": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic mapping of server name indication in Client hello extension",
			},
			"server_name_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_name": {
							Type: schema.TypeString, Optional: true, Description: "Server name indication in Client hello extension (Server name String)",
						},
						"server_cert": {
							Type: schema.TypeString, Optional: true, Description: "Server Certificate associated to SNI (Server Certificate Name)",
						},
						"server_chain": {
							Type: schema.TypeString, Optional: true, Description: "Server Certificate Chain associated to SNI (Server Certificate Chain Name)",
						},
						"server_key": {
							Type: schema.TypeString, Optional: true, Description: "Server Private Key associated to SNI (Server Private Key Name)",
						},
						"server_passphrase": {
							Type: schema.TypeString, Optional: true, Description: "help Password Phrase",
						},
						"server_name_alternate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specific the second certifcate",
						},
						"server_shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Server Name Partition Shared",
						},
						"sni_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Template associated to SNI",
						},
						"sni_template_client_ssl": {
							Type: schema.TypeString, Optional: true, Description: "Client SSL Template Name",
						},
						"sni_shared_partition_client_ssl_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a Client SSL template from shared partition",
						},
						"sni_template_client_ssl_shared_name": {
							Type: schema.TypeString, Optional: true, Description: "Client SSL Template Name",
						},
						"server_name_regex": {
							Type: schema.TypeString, Optional: true, Description: "Server name indication in Client hello extension with regular expression (Server name String with regex)",
						},
						"server_cert_regex": {
							Type: schema.TypeString, Optional: true, Description: "Server Certificate associated to SNI regex (Server Certificate Name)",
						},
						"server_chain_regex": {
							Type: schema.TypeString, Optional: true, Description: "Server Certificate Chain associated to SNI regex (Server Certificate Chain Name)",
						},
						"server_key_regex": {
							Type: schema.TypeString, Optional: true, Description: "Server Private Key associated to SNI regex (Server Private Key Name)",
						},
						"server_passphrase_regex": {
							Type: schema.TypeString, Optional: true, Description: "help Password Phrase",
						},
						"server_name_regex_alternate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specific the second certifcate",
						},
						"server_shared_regex": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Server Name Partition Shared",
						},
						"sni_regex_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Template associated to SNI regex",
						},
						"sni_regex_template_client_ssl": {
							Type: schema.TypeString, Optional: true, Description: "Client SSL Template Name",
						},
						"sni_regex_shared_partition_client_ssl_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a Client SSL template from shared partition",
						},
						"sni_regex_template_client_ssl_shared_name": {
							Type: schema.TypeString, Optional: true, Description: "Client SSL Template Name",
						},
					},
				},
			},
			"session_cache_size": {
				Type: schema.TypeInt, Optional: true, Description: "Session Cache Size (Maximum cache size. Default value 0 (Session ID reuse disabled))",
			},
			"session_cache_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Session Cache Timeout (Timeout value, in seconds. Default value 0 (Session cache timeout disabled))",
			},
			"session_ticket_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable client side session ticket support",
			},
			"session_ticket_lifetime": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Session ticket lifetime in seconds from stateless session resumption (Lifetime value in seconds. Default value 0 (Session ticket lifetime is 7200 seconds))",
			},
			"shared_partition_cipher_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a cipher template from shared partition",
			},
			"shared_partition_pool": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a NAT pool or pool group from shared partition",
			},
			"sni_bypass_enable_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging when bypass event happens, disabled by default",
			},
			"sni_bypass_expired_cert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bypass when certificate expired",
			},
			"sni_bypass_explicit_list": {
				Type: schema.TypeString, Optional: true, Description: "Bypass when matched explicit bypass list (Specify class list name)",
			},
			"sni_bypass_missing_cert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bypass when missing cert/key",
			},
			"sni_enable_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging of sni-auto-map failures. Disable by default",
			},
			"ssl_false_start_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "disable SSL False Start",
			},
			"ssli_inbound_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable inbound SSLi",
			},
			"ssli_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SSLi logging level, default is error logging only",
			},
			"sslilogging": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable all logging; 'all': enable all logging(error, info);",
			},
			"sslv2_bypass_service_group": {
				Type: schema.TypeString, Optional: true, Description: "Service Group for Bypass SSLV2 (Service Group Name)",
			},
			"starts_with_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"starts_with": {
							Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if SNI string starts with another string",
						},
					},
				},
			},
			"template_cipher": {
				Type: schema.TypeString, Optional: true, Description: "Cipher Template Name",
			},
			"template_cipher_shared": {
				Type: schema.TypeString, Optional: true, Description: "Cipher Template Name",
			},
			"template_hsm": {
				Type: schema.TypeString, Optional: true, Description: "HSM Template (HSM Template Name)",
			},
			"user_name_list": {
				Type: schema.TypeString, Optional: true, Description: "Forward proxy bypass if user-name matches class-list",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"verify_cert_fail_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'bypass': bypass SSLi processing; 'continue': continue the connection; 'drop': close the connection; 'block': block the connection with a warning page;",
			},
			"version": {
				Type: schema.TypeInt, Optional: true, Description: "TLS/SSL version, default is the highest number supported (TLS/SSL version: 30-SSLv3.0, 31-TLSv1.0, 32-TLSv1.1, 33-TLSv1.2 and 34-TLSv1.3)",
			},
			"web_category": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uncategorized": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Uncategorized URLs",
						},
						"real_estate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Real Estate",
						},
						"computer_and_internet_security": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Computer and Internet Security",
						},
						"financial_services": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Financial Services",
						},
						"business_and_economy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Business and Economy",
						},
						"computer_and_internet_info": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Computer and Internet Info",
						},
						"auctions": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Auctions",
						},
						"shopping": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Shopping",
						},
						"cult_and_occult": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Cult and Occult",
						},
						"travel": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Travel",
						},
						"drugs": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Abused Drugs",
						},
						"adult_and_pornography": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Adult and Pornography",
						},
						"home_and_garden": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Home and Garden",
						},
						"military": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Military",
						},
						"social_network": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Social Network",
						},
						"dead_sites": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Dead Sites (db Ops only)",
						},
						"stock_advice_and_tools": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Stock Advice and Tools",
						},
						"training_and_tools": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Training and Tools",
						},
						"dating": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Dating",
						},
						"sex_education": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Sex Education",
						},
						"religion": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Religion",
						},
						"entertainment_and_arts": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Entertainment and Arts",
						},
						"personal_sites_and_blogs": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Personal sites and Blogs",
						},
						"legal": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Legal",
						},
						"local_information": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Local Information",
						},
						"streaming_media": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Streaming Media",
						},
						"job_search": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Job Search",
						},
						"gambling": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Gambling",
						},
						"translation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Translation",
						},
						"reference_and_research": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Reference and Research",
						},
						"shareware_and_freeware": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Shareware and Freeware",
						},
						"peer_to_peer": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Peer to Peer",
						},
						"marijuana": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Marijuana",
						},
						"hacking": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Hacking",
						},
						"games": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Games",
						},
						"philosophy_and_politics": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Philosophy and Political Advocacy",
						},
						"weapons": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Weapons",
						},
						"pay_to_surf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Pay to Surf",
						},
						"hunting_and_fishing": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Hunting and Fishing",
						},
						"society": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Society",
						},
						"educational_institutions": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Educational Institutions",
						},
						"online_greeting_cards": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Online Greeting cards",
						},
						"sports": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Sports",
						},
						"swimsuits_and_intimate_apparel": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Swimsuits and Intimate Apparel",
						},
						"questionable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Questionable",
						},
						"kids": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Kids",
						},
						"hate_and_racism": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Hate and Racism",
						},
						"personal_storage": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Personal Storage",
						},
						"violence": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Violence",
						},
						"keyloggers_and_monitoring": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Keyloggers and Monitoring",
						},
						"search_engines": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Search Engines",
						},
						"internet_portals": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Internet Portals",
						},
						"web_advertisements": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Web Advertisements",
						},
						"cheating": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Cheating",
						},
						"gross": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Gross",
						},
						"web_based_email": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Web based email",
						},
						"malware_sites": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Malware Sites",
						},
						"phishing_and_other_fraud": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Phishing and Other Frauds",
						},
						"proxy_avoid_and_anonymizers": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Proxy Avoid and Anonymizers",
						},
						"spyware_and_adware": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Spyware and Adware",
						},
						"music": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Music",
						},
						"government": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Government",
						},
						"nudity": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Nudity",
						},
						"news_and_media": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category News and Media",
						},
						"illegal": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Illegal",
						},
						"cdns": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category CDNs",
						},
						"internet_communications": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Internet Communications",
						},
						"bot_nets": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Bot Nets",
						},
						"abortion": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Abortion",
						},
						"health_and_medicine": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Health and Medicine",
						},
						"spam_urls": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category SPAM URLs",
						},
						"dynamically_generated_content": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dynamically Generated Content",
						},
						"parked_domains": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Parked Domains",
						},
						"alcohol_and_tobacco": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Alcohol and Tobacco",
						},
						"image_and_video_search": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Image and Video Search",
						},
						"fashion_and_beauty": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Fashion and Beauty",
						},
						"recreation_and_hobbies": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Recreation and Hobbies",
						},
						"motor_vehicles": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Motor Vehicles",
						},
						"web_hosting_sites": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Web Hosting Sites",
						},
						"nudity_artistic": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Nudity join Entertainment and Arts",
						},
						"illegal_pornography": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Illegal join Adult and Pornography",
						},
					},
				},
			},
			"web_reputation": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_trustworthy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bypass when reputation score is greater than or equal to 81",
						},
						"bypass_low_risk": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bypass when reputation score is greater than or equal to 61",
						},
						"bypass_moderate_risk": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bypass when reputation score is greater than or equal to 41",
						},
						"bypass_suspicious": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bypass when reputation score is greater than or equal to 21",
						},
						"bypass_malicious": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bypass when reputation score is greater than or equal to 1",
						},
						"bypass_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass when reputation score is greater than or equal to the customized score (1-100)",
						},
					},
				},
			},
		},
	}
}
func resourceSlbTemplateClientSslCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSslCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSsl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateClientSslRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateClientSslUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSslUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSsl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateClientSslRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateClientSslDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSslDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSsl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateClientSslRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSslRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSsl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateClientSslBypassCertIssuerMultiClassList(d []interface{}) []edpt.SlbTemplateClientSslBypassCertIssuerMultiClassList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslBypassCertIssuerMultiClassList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslBypassCertIssuerMultiClassList
		oi.BypassCertIssuerMultiClassListName = in["bypass_cert_issuer_multi_class_list_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslBypassCertSanMultiClassList(d []interface{}) []edpt.SlbTemplateClientSslBypassCertSanMultiClassList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslBypassCertSanMultiClassList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslBypassCertSanMultiClassList
		oi.BypassCertSanMultiClassListName = in["bypass_cert_san_multi_class_list_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslBypassCertSubjectMultiClassList(d []interface{}) []edpt.SlbTemplateClientSslBypassCertSubjectMultiClassList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslBypassCertSubjectMultiClassList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslBypassCertSubjectMultiClassList
		oi.BypassCertSubjectMultiClassListName = in["bypass_cert_subject_multi_class_list_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCaCerts(d []interface{}) []edpt.SlbTemplateClientSslCaCerts {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCaCerts, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCaCerts
		oi.CaCert = in["ca_cert"].(string)
		oi.CaShared = in["ca_shared"].(int)
		oi.ClientOcsp = in["client_ocsp"].(int)
		oi.ClientOcspSrvr = in["client_ocsp_srvr"].(string)
		oi.ClientOcspSg = in["client_ocsp_sg"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCertificateIssuerContainsList(d []interface{}) []edpt.SlbTemplateClientSslCertificateIssuerContainsList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCertificateIssuerContainsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCertificateIssuerContainsList
		oi.CertificateIssuerContains = in["certificate_issuer_contains"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCertificateIssuerEndsWithList(d []interface{}) []edpt.SlbTemplateClientSslCertificateIssuerEndsWithList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCertificateIssuerEndsWithList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCertificateIssuerEndsWithList
		oi.CertificateIssuerEndsWith = in["certificate_issuer_ends_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCertificateIssuerEqualsList(d []interface{}) []edpt.SlbTemplateClientSslCertificateIssuerEqualsList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCertificateIssuerEqualsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCertificateIssuerEqualsList
		oi.CertificateIssuerEquals = in["certificate_issuer_equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCertificateIssuerStartsWithList(d []interface{}) []edpt.SlbTemplateClientSslCertificateIssuerStartsWithList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCertificateIssuerStartsWithList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCertificateIssuerStartsWithList
		oi.CertificateIssuerStarts = in["certificate_issuer_starts"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCertificateList(d []interface{}) []edpt.SlbTemplateClientSslCertificateList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCertificateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCertificateList
		oi.Cert = in["cert"].(string)
		oi.Key = in["key"].(string)
		oi.Passphrase = in["passphrase"].(string)
		//omit key_encrypted
		oi.ChainCert = in["chain_cert"].(string)
		oi.Shared = in["shared"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCertificateSanContainsList(d []interface{}) []edpt.SlbTemplateClientSslCertificateSanContainsList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCertificateSanContainsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCertificateSanContainsList
		oi.CertificateSanContains = in["certificate_san_contains"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCertificateSanEndsWithList(d []interface{}) []edpt.SlbTemplateClientSslCertificateSanEndsWithList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCertificateSanEndsWithList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCertificateSanEndsWithList
		oi.CertificateSanEndsWith = in["certificate_san_ends_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCertificateSanEqualsList(d []interface{}) []edpt.SlbTemplateClientSslCertificateSanEqualsList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCertificateSanEqualsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCertificateSanEqualsList
		oi.CertificateSanEquals = in["certificate_san_equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCertificateSanStartsWithList(d []interface{}) []edpt.SlbTemplateClientSslCertificateSanStartsWithList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCertificateSanStartsWithList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCertificateSanStartsWithList
		oi.CertificateSanStarts = in["certificate_san_starts"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCertificateSubjectContainsList(d []interface{}) []edpt.SlbTemplateClientSslCertificateSubjectContainsList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCertificateSubjectContainsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCertificateSubjectContainsList
		oi.CertificateSubjectContains = in["certificate_subject_contains"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCertificateSubjectEndsWithList(d []interface{}) []edpt.SlbTemplateClientSslCertificateSubjectEndsWithList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCertificateSubjectEndsWithList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCertificateSubjectEndsWithList
		oi.CertificateSubjectEndsWith = in["certificate_subject_ends_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCertificateSubjectEqualsList(d []interface{}) []edpt.SlbTemplateClientSslCertificateSubjectEqualsList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCertificateSubjectEqualsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCertificateSubjectEqualsList
		oi.CertificateSubjectEquals = in["certificate_subject_equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCertificateSubjectStartsWithList(d []interface{}) []edpt.SlbTemplateClientSslCertificateSubjectStartsWithList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCertificateSubjectStartsWithList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCertificateSubjectStartsWithList
		oi.CertificateSubjectStarts = in["certificate_subject_starts"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCipherWithoutPrioList(d []interface{}) []edpt.SlbTemplateClientSslCipherWithoutPrioList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCipherWithoutPrioList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCipherWithoutPrioList
		oi.CipherWoPrio = in["cipher_wo_prio"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslClientAuthContainsList(d []interface{}) []edpt.SlbTemplateClientSslClientAuthContainsList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslClientAuthContainsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslClientAuthContainsList
		oi.ClientAuthContains = in["client_auth_contains"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslClientAuthEndsWithList(d []interface{}) []edpt.SlbTemplateClientSslClientAuthEndsWithList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslClientAuthEndsWithList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslClientAuthEndsWithList
		oi.ClientAuthEndsWith = in["client_auth_ends_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslClientAuthEqualsList(d []interface{}) []edpt.SlbTemplateClientSslClientAuthEqualsList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslClientAuthEqualsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslClientAuthEqualsList
		oi.ClientAuthEquals = in["client_auth_equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslClientAuthStartsWithList(d []interface{}) []edpt.SlbTemplateClientSslClientAuthStartsWithList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslClientAuthStartsWithList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslClientAuthStartsWithList
		oi.ClientAuthStartsWith = in["client_auth_starts_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslClientIpv4List(d []interface{}) []edpt.SlbTemplateClientSslClientIpv4List {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslClientIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslClientIpv4List
		oi.ClientIpv4ListName = in["client_ipv4_list_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslClientIpv6List(d []interface{}) []edpt.SlbTemplateClientSslClientIpv6List {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslClientIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslClientIpv6List
		oi.ClientIpv6ListName = in["client_ipv6_list_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslContainsList(d []interface{}) []edpt.SlbTemplateClientSslContainsList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslContainsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslContainsList
		oi.Contains = in["contains"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslCrlCerts(d []interface{}) []edpt.SlbTemplateClientSslCrlCerts {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslCrlCerts, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslCrlCerts
		oi.Crl = in["crl"].(string)
		oi.CrlShared = in["crl_shared"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslEcList(d []interface{}) []edpt.SlbTemplateClientSslEcList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslEcList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslEcList
		oi.Ec = in["ec"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslEndsWithList(d []interface{}) []edpt.SlbTemplateClientSslEndsWithList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslEndsWithList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslEndsWithList
		oi.EndsWith = in["ends_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslEqualsList(d []interface{}) []edpt.SlbTemplateClientSslEqualsList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslEqualsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslEqualsList
		oi.Equals = in["equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslExceptionClientIpv4List(d []interface{}) []edpt.SlbTemplateClientSslExceptionClientIpv4List {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslExceptionClientIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslExceptionClientIpv4List
		oi.ExceptionClientIpv4ListName = in["exception_client_ipv4_list_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslExceptionClientIpv6List(d []interface{}) []edpt.SlbTemplateClientSslExceptionClientIpv6List {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslExceptionClientIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslExceptionClientIpv6List
		oi.ExceptionClientIpv6ListName = in["exception_client_ipv6_list_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslExceptionServerIpv4List(d []interface{}) []edpt.SlbTemplateClientSslExceptionServerIpv4List {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslExceptionServerIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslExceptionServerIpv4List
		oi.ExceptionServerIpv4ListName = in["exception_server_ipv4_list_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslExceptionServerIpv6List(d []interface{}) []edpt.SlbTemplateClientSslExceptionServerIpv6List {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslExceptionServerIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslExceptionServerIpv6List
		oi.ExceptionServerIpv6ListName = in["exception_server_ipv6_list_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateClientSslExceptionWebCategory(d []interface{}) edpt.SlbTemplateClientSslExceptionWebCategory {

	count1 := len(d)
	var ret edpt.SlbTemplateClientSslExceptionWebCategory
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ExceptionUncategorized = in["exception_uncategorized"].(int)
		ret.ExceptionRealEstate = in["exception_real_estate"].(int)
		ret.ExceptionComputerAndInternetSecurity = in["exception_computer_and_internet_security"].(int)
		ret.ExceptionFinancialServices = in["exception_financial_services"].(int)
		ret.ExceptionBusinessAndEconomy = in["exception_business_and_economy"].(int)
		ret.ExceptionComputerAndInternetInfo = in["exception_computer_and_internet_info"].(int)
		ret.ExceptionAuctions = in["exception_auctions"].(int)
		ret.ExceptionShopping = in["exception_shopping"].(int)
		ret.ExceptionCultAndOccult = in["exception_cult_and_occult"].(int)
		ret.ExceptionTravel = in["exception_travel"].(int)
		ret.ExceptionDrugs = in["exception_drugs"].(int)
		ret.ExceptionAdultAndPornography = in["exception_adult_and_pornography"].(int)
		ret.ExceptionHomeAndGarden = in["exception_home_and_garden"].(int)
		ret.ExceptionMilitary = in["exception_military"].(int)
		ret.ExceptionSocialNetwork = in["exception_social_network"].(int)
		ret.ExceptionDeadSites = in["exception_dead_sites"].(int)
		ret.ExceptionStockAdviceAndTools = in["exception_stock_advice_and_tools"].(int)
		ret.ExceptionTrainingAndTools = in["exception_training_and_tools"].(int)
		ret.ExceptionDating = in["exception_dating"].(int)
		ret.ExceptionSexEducation = in["exception_sex_education"].(int)
		ret.ExceptionReligion = in["exception_religion"].(int)
		ret.ExceptionEntertainmentAndArts = in["exception_entertainment_and_arts"].(int)
		ret.ExceptionPersonalSitesAndBlogs = in["exception_personal_sites_and_blogs"].(int)
		ret.ExceptionLegal = in["exception_legal"].(int)
		ret.ExceptionLocalInformation = in["exception_local_information"].(int)
		ret.ExceptionStreamingMedia = in["exception_streaming_media"].(int)
		ret.ExceptionJobSearch = in["exception_job_search"].(int)
		ret.ExceptionGambling = in["exception_gambling"].(int)
		ret.ExceptionTranslation = in["exception_translation"].(int)
		ret.ExceptionReferenceAndResearch = in["exception_reference_and_research"].(int)
		ret.ExceptionSharewareAndFreeware = in["exception_shareware_and_freeware"].(int)
		ret.ExceptionPeerToPeer = in["exception_peer_to_peer"].(int)
		ret.ExceptionMarijuana = in["exception_marijuana"].(int)
		ret.ExceptionHacking = in["exception_hacking"].(int)
		ret.ExceptionGames = in["exception_games"].(int)
		ret.ExceptionPhilosophyAndPolitics = in["exception_philosophy_and_politics"].(int)
		ret.ExceptionWeapons = in["exception_weapons"].(int)
		ret.ExceptionPayToSurf = in["exception_pay_to_surf"].(int)
		ret.ExceptionHuntingAndFishing = in["exception_hunting_and_fishing"].(int)
		ret.ExceptionSociety = in["exception_society"].(int)
		ret.ExceptionEducationalInstitutions = in["exception_educational_institutions"].(int)
		ret.ExceptionOnlineGreetingCards = in["exception_online_greeting_cards"].(int)
		ret.ExceptionSports = in["exception_sports"].(int)
		ret.ExceptionSwimsuitsAndIntimateApparel = in["exception_swimsuits_and_intimate_apparel"].(int)
		ret.ExceptionQuestionable = in["exception_questionable"].(int)
		ret.ExceptionKids = in["exception_kids"].(int)
		ret.ExceptionHateAndRacism = in["exception_hate_and_racism"].(int)
		ret.ExceptionPersonalStorage = in["exception_personal_storage"].(int)
		ret.ExceptionViolence = in["exception_violence"].(int)
		ret.ExceptionKeyloggersAndMonitoring = in["exception_keyloggers_and_monitoring"].(int)
		ret.ExceptionSearchEngines = in["exception_search_engines"].(int)
		ret.ExceptionInternetPortals = in["exception_internet_portals"].(int)
		ret.ExceptionWebAdvertisements = in["exception_web_advertisements"].(int)
		ret.ExceptionCheating = in["exception_cheating"].(int)
		ret.ExceptionGross = in["exception_gross"].(int)
		ret.ExceptionWebBasedEmail = in["exception_web_based_email"].(int)
		ret.ExceptionMalwareSites = in["exception_malware_sites"].(int)
		ret.ExceptionPhishingAndOtherFraud = in["exception_phishing_and_other_fraud"].(int)
		ret.ExceptionProxyAvoidAndAnonymizers = in["exception_proxy_avoid_and_anonymizers"].(int)
		ret.ExceptionSpywareAndAdware = in["exception_spyware_and_adware"].(int)
		ret.ExceptionMusic = in["exception_music"].(int)
		ret.ExceptionGovernment = in["exception_government"].(int)
		ret.ExceptionNudity = in["exception_nudity"].(int)
		ret.ExceptionNewsAndMedia = in["exception_news_and_media"].(int)
		ret.ExceptionIllegal = in["exception_illegal"].(int)
		ret.ExceptionCdns = in["exception_cdns"].(int)
		ret.ExceptionInternetCommunications = in["exception_internet_communications"].(int)
		ret.ExceptionBotNets = in["exception_bot_nets"].(int)
		ret.ExceptionAbortion = in["exception_abortion"].(int)
		ret.ExceptionHealthAndMedicine = in["exception_health_and_medicine"].(int)
		ret.ExceptionSpamUrls = in["exception_spam_urls"].(int)
		ret.ExceptionDynamicallyGeneratedContent = in["exception_dynamically_generated_content"].(int)
		ret.ExceptionParkedDomains = in["exception_parked_domains"].(int)
		ret.ExceptionAlcoholAndTobacco = in["exception_alcohol_and_tobacco"].(int)
		ret.ExceptionImageAndVideoSearch = in["exception_image_and_video_search"].(int)
		ret.ExceptionFashionAndBeauty = in["exception_fashion_and_beauty"].(int)
		ret.ExceptionRecreationAndHobbies = in["exception_recreation_and_hobbies"].(int)
		ret.ExceptionMotorVehicles = in["exception_motor_vehicles"].(int)
		ret.ExceptionWebHostingSites = in["exception_web_hosting_sites"].(int)
		ret.ExceptionNudityArtistic = in["exception_nudity_artistic"].(int)
		ret.ExceptionIllegalPornography = in["exception_illegal_pornography"].(int)
	}
	return ret
}

func getObjectSlbTemplateClientSslExceptionWebReputation(d []interface{}) edpt.SlbTemplateClientSslExceptionWebReputation {

	count1 := len(d)
	var ret edpt.SlbTemplateClientSslExceptionWebReputation
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ExceptionTrustworthy = in["exception_trustworthy"].(int)
		ret.ExceptionLowRisk = in["exception_low_risk"].(int)
		ret.ExceptionModerateRisk = in["exception_moderate_risk"].(int)
		ret.ExceptionSuspicious = in["exception_suspicious"].(int)
		ret.ExceptionMalicious = in["exception_malicious"].(int)
		ret.ExceptionThreshold = in["exception_threshold"].(int)
	}
	return ret
}

func getSliceSlbTemplateClientSslForwardProxyTrustedCaLists(d []interface{}) []edpt.SlbTemplateClientSslForwardProxyTrustedCaLists {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslForwardProxyTrustedCaLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslForwardProxyTrustedCaLists
		oi.ForwardProxyTrustedCa = in["forward_proxy_trusted_ca"].(string)
		oi.FpTrustedCaShared = in["fp_trusted_ca_shared"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateClientSslLocalCertPinList(d []interface{}) edpt.SlbTemplateClientSslLocalCertPinList {

	count1 := len(d)
	var ret edpt.SlbTemplateClientSslLocalCertPinList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LocalCertPinListBypassFailCount = in["local_cert_pin_list_bypass_fail_count"].(int)
	}
	return ret
}

func getSliceSlbTemplateClientSslMultiClassList(d []interface{}) []edpt.SlbTemplateClientSslMultiClassList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslMultiClassList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslMultiClassList
		oi.MultiClistName = in["multi_clist_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslReqCaLists(d []interface{}) []edpt.SlbTemplateClientSslReqCaLists {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslReqCaLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslReqCaLists
		oi.ClientCertificateRequestCa = in["client_certificate_request_ca"].(string)
		oi.ClientCertReqCaShared = in["client_cert_req_ca_shared"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslServerIpv4List(d []interface{}) []edpt.SlbTemplateClientSslServerIpv4List {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslServerIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslServerIpv4List
		oi.ServerIpv4ListName = in["server_ipv4_list_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslServerIpv6List(d []interface{}) []edpt.SlbTemplateClientSslServerIpv6List {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslServerIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslServerIpv6List
		oi.ServerIpv6ListName = in["server_ipv6_list_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslServerNameList(d []interface{}) []edpt.SlbTemplateClientSslServerNameList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslServerNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslServerNameList
		oi.ServerName = in["server_name"].(string)
		oi.ServerCert = in["server_cert"].(string)
		oi.ServerChain = in["server_chain"].(string)
		oi.ServerKey = in["server_key"].(string)
		oi.ServerPassphrase = in["server_passphrase"].(string)
		//omit server_encrypted
		oi.ServerNameAlternate = in["server_name_alternate"].(int)
		oi.ServerShared = in["server_shared"].(int)
		oi.SniTemplate = in["sni_template"].(int)
		oi.SniTemplateClientSsl = in["sni_template_client_ssl"].(string)
		oi.SniSharedPartitionClientSslTemplate = in["sni_shared_partition_client_ssl_template"].(int)
		oi.SniTemplateClientSslSharedName = in["sni_template_client_ssl_shared_name"].(string)
		oi.ServerNameRegex = in["server_name_regex"].(string)
		oi.ServerCertRegex = in["server_cert_regex"].(string)
		oi.ServerChainRegex = in["server_chain_regex"].(string)
		oi.ServerKeyRegex = in["server_key_regex"].(string)
		oi.ServerPassphraseRegex = in["server_passphrase_regex"].(string)
		//omit server_encrypted_regex
		oi.ServerNameRegexAlternate = in["server_name_regex_alternate"].(int)
		oi.ServerSharedRegex = in["server_shared_regex"].(int)
		oi.SniRegexTemplate = in["sni_regex_template"].(int)
		oi.SniRegexTemplateClientSsl = in["sni_regex_template_client_ssl"].(string)
		oi.SniRegexSharedPartitionClientSslTemplate = in["sni_regex_shared_partition_client_ssl_template"].(int)
		oi.SniRegexTemplateClientSslSharedName = in["sni_regex_template_client_ssl_shared_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateClientSslStartsWithList(d []interface{}) []edpt.SlbTemplateClientSslStartsWithList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslStartsWithList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslStartsWithList
		oi.StartsWith = in["starts_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateClientSslWebCategory(d []interface{}) edpt.SlbTemplateClientSslWebCategory {

	count1 := len(d)
	var ret edpt.SlbTemplateClientSslWebCategory
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Uncategorized = in["uncategorized"].(int)
		ret.RealEstate = in["real_estate"].(int)
		ret.ComputerAndInternetSecurity = in["computer_and_internet_security"].(int)
		ret.FinancialServices = in["financial_services"].(int)
		ret.BusinessAndEconomy = in["business_and_economy"].(int)
		ret.ComputerAndInternetInfo = in["computer_and_internet_info"].(int)
		ret.Auctions = in["auctions"].(int)
		ret.Shopping = in["shopping"].(int)
		ret.CultAndOccult = in["cult_and_occult"].(int)
		ret.Travel = in["travel"].(int)
		ret.Drugs = in["drugs"].(int)
		ret.AdultAndPornography = in["adult_and_pornography"].(int)
		ret.HomeAndGarden = in["home_and_garden"].(int)
		ret.Military = in["military"].(int)
		ret.SocialNetwork = in["social_network"].(int)
		ret.DeadSites = in["dead_sites"].(int)
		ret.StockAdviceAndTools = in["stock_advice_and_tools"].(int)
		ret.TrainingAndTools = in["training_and_tools"].(int)
		ret.Dating = in["dating"].(int)
		ret.SexEducation = in["sex_education"].(int)
		ret.Religion = in["religion"].(int)
		ret.EntertainmentAndArts = in["entertainment_and_arts"].(int)
		ret.PersonalSitesAndBlogs = in["personal_sites_and_blogs"].(int)
		ret.Legal = in["legal"].(int)
		ret.LocalInformation = in["local_information"].(int)
		ret.StreamingMedia = in["streaming_media"].(int)
		ret.JobSearch = in["job_search"].(int)
		ret.Gambling = in["gambling"].(int)
		ret.Translation = in["translation"].(int)
		ret.ReferenceAndResearch = in["reference_and_research"].(int)
		ret.SharewareAndFreeware = in["shareware_and_freeware"].(int)
		ret.PeerToPeer = in["peer_to_peer"].(int)
		ret.Marijuana = in["marijuana"].(int)
		ret.Hacking = in["hacking"].(int)
		ret.Games = in["games"].(int)
		ret.PhilosophyAndPolitics = in["philosophy_and_politics"].(int)
		ret.Weapons = in["weapons"].(int)
		ret.PayToSurf = in["pay_to_surf"].(int)
		ret.HuntingAndFishing = in["hunting_and_fishing"].(int)
		ret.Society = in["society"].(int)
		ret.EducationalInstitutions = in["educational_institutions"].(int)
		ret.OnlineGreetingCards = in["online_greeting_cards"].(int)
		ret.Sports = in["sports"].(int)
		ret.SwimsuitsAndIntimateApparel = in["swimsuits_and_intimate_apparel"].(int)
		ret.Questionable = in["questionable"].(int)
		ret.Kids = in["kids"].(int)
		ret.HateAndRacism = in["hate_and_racism"].(int)
		ret.PersonalStorage = in["personal_storage"].(int)
		ret.Violence = in["violence"].(int)
		ret.KeyloggersAndMonitoring = in["keyloggers_and_monitoring"].(int)
		ret.SearchEngines = in["search_engines"].(int)
		ret.InternetPortals = in["internet_portals"].(int)
		ret.WebAdvertisements = in["web_advertisements"].(int)
		ret.Cheating = in["cheating"].(int)
		ret.Gross = in["gross"].(int)
		ret.WebBasedEmail = in["web_based_email"].(int)
		ret.MalwareSites = in["malware_sites"].(int)
		ret.PhishingAndOtherFraud = in["phishing_and_other_fraud"].(int)
		ret.ProxyAvoidAndAnonymizers = in["proxy_avoid_and_anonymizers"].(int)
		ret.SpywareAndAdware = in["spyware_and_adware"].(int)
		ret.Music = in["music"].(int)
		ret.Government = in["government"].(int)
		ret.Nudity = in["nudity"].(int)
		ret.NewsAndMedia = in["news_and_media"].(int)
		ret.Illegal = in["illegal"].(int)
		ret.Cdns = in["cdns"].(int)
		ret.InternetCommunications = in["internet_communications"].(int)
		ret.BotNets = in["bot_nets"].(int)
		ret.Abortion = in["abortion"].(int)
		ret.HealthAndMedicine = in["health_and_medicine"].(int)
		ret.SpamUrls = in["spam_urls"].(int)
		ret.DynamicallyGeneratedContent = in["dynamically_generated_content"].(int)
		ret.ParkedDomains = in["parked_domains"].(int)
		ret.AlcoholAndTobacco = in["alcohol_and_tobacco"].(int)
		ret.ImageAndVideoSearch = in["image_and_video_search"].(int)
		ret.FashionAndBeauty = in["fashion_and_beauty"].(int)
		ret.RecreationAndHobbies = in["recreation_and_hobbies"].(int)
		ret.MotorVehicles = in["motor_vehicles"].(int)
		ret.WebHostingSites = in["web_hosting_sites"].(int)
		ret.NudityArtistic = in["nudity_artistic"].(int)
		ret.IllegalPornography = in["illegal_pornography"].(int)
	}
	return ret
}

func getObjectSlbTemplateClientSslWebReputation(d []interface{}) edpt.SlbTemplateClientSslWebReputation {

	count1 := len(d)
	var ret edpt.SlbTemplateClientSslWebReputation
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BypassTrustworthy = in["bypass_trustworthy"].(int)
		ret.BypassLowRisk = in["bypass_low_risk"].(int)
		ret.BypassModerateRisk = in["bypass_moderate_risk"].(int)
		ret.BypassSuspicious = in["bypass_suspicious"].(int)
		ret.BypassMalicious = in["bypass_malicious"].(int)
		ret.BypassThreshold = in["bypass_threshold"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplateClientSsl(d *schema.ResourceData) edpt.SlbTemplateClientSsl {
	var ret edpt.SlbTemplateClientSsl
	ret.Inst.AdGroupList = d.Get("ad_group_list").(string)
	ret.Inst.AlertType = d.Get("alert_type").(string)
	ret.Inst.AuthSg = d.Get("auth_sg").(string)
	ret.Inst.AuthSgDn = d.Get("auth_sg_dn").(int)
	ret.Inst.AuthSgFilter = d.Get("auth_sg_filter").(string)
	ret.Inst.AuthUsername = d.Get("auth_username").(string)
	ret.Inst.AuthUsernameAttribute = d.Get("auth_username_attribute").(string)
	ret.Inst.AuthenName = d.Get("authen_name").(string)
	ret.Inst.Authorization = d.Get("authorization").(int)
	ret.Inst.BypassCertIssuerClassListName = d.Get("bypass_cert_issuer_class_list_name").(string)
	ret.Inst.BypassCertIssuerMultiClassList = getSliceSlbTemplateClientSslBypassCertIssuerMultiClassList(d.Get("bypass_cert_issuer_multi_class_list").([]interface{}))
	ret.Inst.BypassCertSanClassListName = d.Get("bypass_cert_san_class_list_name").(string)
	ret.Inst.BypassCertSanMultiClassList = getSliceSlbTemplateClientSslBypassCertSanMultiClassList(d.Get("bypass_cert_san_multi_class_list").([]interface{}))
	ret.Inst.BypassCertSubjectClassListName = d.Get("bypass_cert_subject_class_list_name").(string)
	ret.Inst.BypassCertSubjectMultiClassList = getSliceSlbTemplateClientSslBypassCertSubjectMultiClassList(d.Get("bypass_cert_subject_multi_class_list").([]interface{}))
	ret.Inst.CaCerts = getSliceSlbTemplateClientSslCaCerts(d.Get("ca_certs").([]interface{}))
	ret.Inst.CachePersistenceListName = d.Get("cache_persistence_list_name").(string)
	ret.Inst.CaseInsensitive = d.Get("case_insensitive").(int)
	ret.Inst.CentralCertPinList = d.Get("central_cert_pin_list").(int)
	ret.Inst.CertRevokeAction = d.Get("cert_revoke_action").(string)
	ret.Inst.CertUnknownAction = d.Get("cert_unknown_action").(string)
	ret.Inst.CertificateIssuerContainsList = getSliceSlbTemplateClientSslCertificateIssuerContainsList(d.Get("certificate_issuer_contains_list").([]interface{}))
	ret.Inst.CertificateIssuerEndsWithList = getSliceSlbTemplateClientSslCertificateIssuerEndsWithList(d.Get("certificate_issuer_ends_with_list").([]interface{}))
	ret.Inst.CertificateIssuerEqualsList = getSliceSlbTemplateClientSslCertificateIssuerEqualsList(d.Get("certificate_issuer_equals_list").([]interface{}))
	ret.Inst.CertificateIssuerStartsWithList = getSliceSlbTemplateClientSslCertificateIssuerStartsWithList(d.Get("certificate_issuer_starts_with_list").([]interface{}))
	ret.Inst.CertificateList = getSliceSlbTemplateClientSslCertificateList(d.Get("certificate_list").([]interface{}))
	ret.Inst.CertificateSanContainsList = getSliceSlbTemplateClientSslCertificateSanContainsList(d.Get("certificate_san_contains_list").([]interface{}))
	ret.Inst.CertificateSanEndsWithList = getSliceSlbTemplateClientSslCertificateSanEndsWithList(d.Get("certificate_san_ends_with_list").([]interface{}))
	ret.Inst.CertificateSanEqualsList = getSliceSlbTemplateClientSslCertificateSanEqualsList(d.Get("certificate_san_equals_list").([]interface{}))
	ret.Inst.CertificateSanStartsWithList = getSliceSlbTemplateClientSslCertificateSanStartsWithList(d.Get("certificate_san_starts_with_list").([]interface{}))
	ret.Inst.CertificateSubjectContainsList = getSliceSlbTemplateClientSslCertificateSubjectContainsList(d.Get("certificate_subject_contains_list").([]interface{}))
	ret.Inst.CertificateSubjectEndsWithList = getSliceSlbTemplateClientSslCertificateSubjectEndsWithList(d.Get("certificate_subject_ends_with_list").([]interface{}))
	ret.Inst.CertificateSubjectEqualsList = getSliceSlbTemplateClientSslCertificateSubjectEqualsList(d.Get("certificate_subject_equals_list").([]interface{}))
	ret.Inst.CertificateSubjectStartsWithList = getSliceSlbTemplateClientSslCertificateSubjectStartsWithList(d.Get("certificate_subject_starts_with_list").([]interface{}))
	ret.Inst.ChainCert = d.Get("chain_cert").(string)
	ret.Inst.ChainCertSharedStr = d.Get("chain_cert_shared_str").(string)
	ret.Inst.CipherWithoutPrioList = getSliceSlbTemplateClientSslCipherWithoutPrioList(d.Get("cipher_without_prio_list").([]interface{}))
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	ret.Inst.ClientAuthCaseInsensitive = d.Get("client_auth_case_insensitive").(int)
	ret.Inst.ClientAuthClassList = d.Get("client_auth_class_list").(string)
	ret.Inst.ClientAuthContainsList = getSliceSlbTemplateClientSslClientAuthContainsList(d.Get("client_auth_contains_list").([]interface{}))
	ret.Inst.ClientAuthEndsWithList = getSliceSlbTemplateClientSslClientAuthEndsWithList(d.Get("client_auth_ends_with_list").([]interface{}))
	ret.Inst.ClientAuthEqualsList = getSliceSlbTemplateClientSslClientAuthEqualsList(d.Get("client_auth_equals_list").([]interface{}))
	ret.Inst.ClientAuthStartsWithList = getSliceSlbTemplateClientSslClientAuthStartsWithList(d.Get("client_auth_starts_with_list").([]interface{}))
	ret.Inst.ClientCertificate = d.Get("client_certificate").(string)
	ret.Inst.ClientIpv4List = getSliceSlbTemplateClientSslClientIpv4List(d.Get("client_ipv4_list").([]interface{}))
	ret.Inst.ClientIpv6List = getSliceSlbTemplateClientSslClientIpv6List(d.Get("client_ipv6_list").([]interface{}))
	ret.Inst.CloseNotify = d.Get("close_notify").(int)
	ret.Inst.ContainsList = getSliceSlbTemplateClientSslContainsList(d.Get("contains_list").([]interface{}))
	ret.Inst.CrlCerts = getSliceSlbTemplateClientSslCrlCerts(d.Get("crl_certs").([]interface{}))
	ret.Inst.Dgversion = d.Get("dgversion").(int)
	ret.Inst.DhType = d.Get("dh_type").(string)
	ret.Inst.DirectClientServerAuth = d.Get("direct_client_server_auth").(int)
	ret.Inst.DisableSslv3 = d.Get("disable_sslv3").(int)
	ret.Inst.EarlyData = d.Get("early_data").(int)
	ret.Inst.EcList = getSliceSlbTemplateClientSslEcList(d.Get("ec_list").([]interface{}))
	ret.Inst.EnableSsliFtpAlg = d.Get("enable_ssli_ftp_alg").(int)
	ret.Inst.EnableTlsAlertLogging = d.Get("enable_tls_alert_logging").(int)
	ret.Inst.EndsWithList = getSliceSlbTemplateClientSslEndsWithList(d.Get("ends_with_list").([]interface{}))
	ret.Inst.EqualsList = getSliceSlbTemplateClientSslEqualsList(d.Get("equals_list").([]interface{}))
	ret.Inst.ExceptionAdGroupList = d.Get("exception_ad_group_list").(string)
	ret.Inst.ExceptionCertificateIssuerClName = d.Get("exception_certificate_issuer_cl_name").(string)
	ret.Inst.ExceptionCertificateSanClName = d.Get("exception_certificate_san_cl_name").(string)
	ret.Inst.ExceptionCertificateSubjectClName = d.Get("exception_certificate_subject_cl_name").(string)
	ret.Inst.ExceptionClientIpv4List = getSliceSlbTemplateClientSslExceptionClientIpv4List(d.Get("exception_client_ipv4_list").([]interface{}))
	ret.Inst.ExceptionClientIpv6List = getSliceSlbTemplateClientSslExceptionClientIpv6List(d.Get("exception_client_ipv6_list").([]interface{}))
	ret.Inst.ExceptionServerIpv4List = getSliceSlbTemplateClientSslExceptionServerIpv4List(d.Get("exception_server_ipv4_list").([]interface{}))
	ret.Inst.ExceptionServerIpv6List = getSliceSlbTemplateClientSslExceptionServerIpv6List(d.Get("exception_server_ipv6_list").([]interface{}))
	ret.Inst.ExceptionSniClName = d.Get("exception_sni_cl_name").(string)
	ret.Inst.ExceptionUserNameList = d.Get("exception_user_name_list").(string)
	ret.Inst.ExceptionWebCategory = getObjectSlbTemplateClientSslExceptionWebCategory(d.Get("exception_web_category").([]interface{}))
	ret.Inst.ExceptionWebReputation = getObjectSlbTemplateClientSslExceptionWebReputation(d.Get("exception_web_reputation").([]interface{}))
	ret.Inst.ExpireHours = d.Get("expire_hours").(int)
	//omit forward_encrypted
	ret.Inst.ForwardPassphrase = d.Get("forward_passphrase").(string)
	ret.Inst.ForwardProxyAltSign = d.Get("forward_proxy_alt_sign").(int)
	ret.Inst.ForwardProxyBlockMessage = d.Get("forward_proxy_block_message").(string)
	ret.Inst.ForwardProxyCaCert = d.Get("forward_proxy_ca_cert").(string)
	ret.Inst.ForwardProxyCaKey = d.Get("forward_proxy_ca_key").(string)
	ret.Inst.ForwardProxyCertCacheLimit = d.Get("forward_proxy_cert_cache_limit").(int)
	ret.Inst.ForwardProxyCertCacheTimeout = d.Get("forward_proxy_cert_cache_timeout").(int)
	ret.Inst.ForwardProxyCertExpiry = d.Get("forward_proxy_cert_expiry").(int)
	ret.Inst.ForwardProxyCertNotReadyAction = d.Get("forward_proxy_cert_not_ready_action").(string)
	ret.Inst.ForwardProxyCertRevokeAction = d.Get("forward_proxy_cert_revoke_action").(int)
	ret.Inst.ForwardProxyCertUnknownAction = d.Get("forward_proxy_cert_unknown_action").(int)
	ret.Inst.ForwardProxyCrlDisable = d.Get("forward_proxy_crl_disable").(int)
	ret.Inst.ForwardProxyDecryptedDscp = d.Get("forward_proxy_decrypted_dscp").(int)
	ret.Inst.ForwardProxyDecryptedDscpBypass = d.Get("forward_proxy_decrypted_dscp_bypass").(int)
	ret.Inst.ForwardProxyEnable = d.Get("forward_proxy_enable").(int)
	ret.Inst.ForwardProxyEsniAction = d.Get("forward_proxy_esni_action").(int)
	ret.Inst.ForwardProxyFailsafeDisable = d.Get("forward_proxy_failsafe_disable").(int)
	ret.Inst.ForwardProxyHashPersistenceInterval = d.Get("forward_proxy_hash_persistence_interval").(int)
	ret.Inst.ForwardProxyLogDisable = d.Get("forward_proxy_log_disable").(int)
	ret.Inst.ForwardProxyNoSharedCipherAction = d.Get("forward_proxy_no_shared_cipher_action").(int)
	ret.Inst.ForwardProxyNoSniAction = d.Get("forward_proxy_no_sni_action").(string)
	ret.Inst.ForwardProxyOcspDisable = d.Get("forward_proxy_ocsp_disable").(int)
	ret.Inst.ForwardProxyRequireSniCertMatched = d.Get("forward_proxy_require_sni_cert_matched").(string)
	ret.Inst.ForwardProxySelfsignRedir = d.Get("forward_proxy_selfsign_redir").(int)
	ret.Inst.ForwardProxySslVersion = d.Get("forward_proxy_ssl_version").(int)
	ret.Inst.ForwardProxyTrustedCaLists = getSliceSlbTemplateClientSslForwardProxyTrustedCaLists(d.Get("forward_proxy_trusted_ca_lists").([]interface{}))
	ret.Inst.ForwardProxyVerifyCertFailAction = d.Get("forward_proxy_verify_cert_fail_action").(int)
	ret.Inst.FpAltCert = d.Get("fp_alt_cert").(string)
	ret.Inst.FpAltChainCert = d.Get("fp_alt_chain_cert").(string)
	//omit fp_alt_encrypted
	ret.Inst.FpAltKey = d.Get("fp_alt_key").(string)
	ret.Inst.FpAltPassphrase = d.Get("fp_alt_passphrase").(string)
	ret.Inst.FpAltShared = d.Get("fp_alt_shared").(int)
	ret.Inst.FpCaCertificate = d.Get("fp_ca_certificate").(string)
	ret.Inst.FpCaCertificateShared = d.Get("fp_ca_certificate_shared").(int)
	ret.Inst.FpCaChainCert = d.Get("fp_ca_chain_cert").(string)
	ret.Inst.FpCaKey = d.Get("fp_ca_key").(string)
	//omit fp_ca_key_encrypted
	ret.Inst.FpCaKeyPassphrase = d.Get("fp_ca_key_passphrase").(string)
	ret.Inst.FpCaKeyShared = d.Get("fp_ca_key_shared").(int)
	ret.Inst.FpCaShared = d.Get("fp_ca_shared").(int)
	ret.Inst.FpCertExtAiaCaIssuers = d.Get("fp_cert_ext_aia_ca_issuers").(string)
	ret.Inst.FpCertExtAiaOcsp = d.Get("fp_cert_ext_aia_ocsp").(string)
	ret.Inst.FpCertExtCrldp = d.Get("fp_cert_ext_crldp").(string)
	ret.Inst.FpCertFetchAutonat = d.Get("fp_cert_fetch_autonat").(string)
	ret.Inst.FpCertFetchAutonatPrecedence = d.Get("fp_cert_fetch_autonat_precedence").(int)
	ret.Inst.FpCertFetchNatpoolName = d.Get("fp_cert_fetch_natpool_name").(string)
	ret.Inst.FpCertFetchNatpoolNameShared = d.Get("fp_cert_fetch_natpool_name_shared").(string)
	ret.Inst.FpCertFetchNatpoolPrecedence = d.Get("fp_cert_fetch_natpool_precedence").(int)
	ret.Inst.FpEsniAction = d.Get("fp_esni_action").(string)
	ret.Inst.HandshakeLoggingEnable = d.Get("handshake_logging_enable").(int)
	ret.Inst.HsmType = d.Get("hsm_type").(string)
	ret.Inst.InspectCertificateIssuerClName = d.Get("inspect_certificate_issuer_cl_name").(string)
	ret.Inst.InspectCertificateSanClName = d.Get("inspect_certificate_san_cl_name").(string)
	ret.Inst.InspectCertificateSubjectClName = d.Get("inspect_certificate_subject_cl_name").(string)
	ret.Inst.InspectListName = d.Get("inspect_list_name").(string)
	ret.Inst.Ja3Enable = d.Get("ja3_enable").(int)
	ret.Inst.Ja3InsertHttpHeader = d.Get("ja3_insert_http_header").(string)
	ret.Inst.Ja3RejectClassList = d.Get("ja3_reject_class_list").(string)
	ret.Inst.Ja3RejectMaxNumberPerHost = d.Get("ja3_reject_max_number_per_host").(int)
	ret.Inst.Ja3Ttl = d.Get("ja3_ttl").(int)
	ret.Inst.LdapBaseDnFromCert = d.Get("ldap_base_dn_from_cert").(int)
	ret.Inst.LdapSearchFilter = d.Get("ldap_search_filter").(string)
	ret.Inst.LocalCertPinList = getObjectSlbTemplateClientSslLocalCertPinList(d.Get("local_cert_pin_list").([]interface{}))
	ret.Inst.LocalLogging = d.Get("local_logging").(int)
	ret.Inst.MultiClassList = getSliceSlbTemplateClientSslMultiClassList(d.Get("multi_class_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NoAntiReplay = d.Get("no_anti_replay").(int)
	ret.Inst.NoSharedCipherAction = d.Get("no_shared_cipher_action").(string)
	ret.Inst.NonSslBypassL4session = d.Get("non_ssl_bypass_l4session").(int)
	ret.Inst.NonSslBypassServiceGroup = d.Get("non_ssl_bypass_service_group").(string)
	ret.Inst.Notafter = d.Get("notafter").(int)
	ret.Inst.Notafterday = d.Get("notafterday").(int)
	ret.Inst.Notaftermonth = d.Get("notaftermonth").(int)
	ret.Inst.Notafteryear = d.Get("notafteryear").(int)
	ret.Inst.Notbefore = d.Get("notbefore").(int)
	ret.Inst.Notbeforeday = d.Get("notbeforeday").(int)
	ret.Inst.Notbeforemonth = d.Get("notbeforemonth").(int)
	ret.Inst.Notbeforeyear = d.Get("notbeforeyear").(int)
	ret.Inst.OcspStapling = d.Get("ocsp_stapling").(int)
	ret.Inst.OcspstCaCert = d.Get("ocspst_ca_cert").(string)
	ret.Inst.OcspstOcsp = d.Get("ocspst_ocsp").(int)
	ret.Inst.OcspstSg = d.Get("ocspst_sg").(string)
	ret.Inst.OcspstSgDays = d.Get("ocspst_sg_days").(int)
	ret.Inst.OcspstSgHours = d.Get("ocspst_sg_hours").(int)
	ret.Inst.OcspstSgMinutes = d.Get("ocspst_sg_minutes").(int)
	ret.Inst.OcspstSgTimeout = d.Get("ocspst_sg_timeout").(int)
	ret.Inst.OcspstSrvr = d.Get("ocspst_srvr").(string)
	ret.Inst.OcspstSrvrDays = d.Get("ocspst_srvr_days").(int)
	ret.Inst.OcspstSrvrHours = d.Get("ocspst_srvr_hours").(int)
	ret.Inst.OcspstSrvrMinutes = d.Get("ocspst_srvr_minutes").(int)
	ret.Inst.OcspstSrvrTimeout = d.Get("ocspst_srvr_timeout").(int)
	ret.Inst.RenegotiationDisable = d.Get("renegotiation_disable").(int)
	ret.Inst.ReqCaLists = getSliceSlbTemplateClientSslReqCaLists(d.Get("req_ca_lists").([]interface{}))
	ret.Inst.RequireWebCategory = d.Get("require_web_category").(int)
	ret.Inst.ServerIpv4List = getSliceSlbTemplateClientSslServerIpv4List(d.Get("server_ipv4_list").([]interface{}))
	ret.Inst.ServerIpv6List = getSliceSlbTemplateClientSslServerIpv6List(d.Get("server_ipv6_list").([]interface{}))
	ret.Inst.ServerNameAutoMap = d.Get("server_name_auto_map").(int)
	ret.Inst.ServerNameList = getSliceSlbTemplateClientSslServerNameList(d.Get("server_name_list").([]interface{}))
	ret.Inst.SessionCacheSize = d.Get("session_cache_size").(int)
	ret.Inst.SessionCacheTimeout = d.Get("session_cache_timeout").(int)
	ret.Inst.SessionTicketDisable = d.Get("session_ticket_disable").(int)
	ret.Inst.SessionTicketLifetime = d.Get("session_ticket_lifetime").(int)
	ret.Inst.SharedPartitionCipherTemplate = d.Get("shared_partition_cipher_template").(int)
	ret.Inst.SharedPartitionPool = d.Get("shared_partition_pool").(int)
	ret.Inst.SniBypassEnableLog = d.Get("sni_bypass_enable_log").(int)
	ret.Inst.SniBypassExpiredCert = d.Get("sni_bypass_expired_cert").(int)
	ret.Inst.SniBypassExplicitList = d.Get("sni_bypass_explicit_list").(string)
	ret.Inst.SniBypassMissingCert = d.Get("sni_bypass_missing_cert").(int)
	ret.Inst.SniEnableLog = d.Get("sni_enable_log").(int)
	ret.Inst.SslFalseStartDisable = d.Get("ssl_false_start_disable").(int)
	ret.Inst.SsliInboundEnable = d.Get("ssli_inbound_enable").(int)
	ret.Inst.SsliLogging = d.Get("ssli_logging").(int)
	ret.Inst.Sslilogging = d.Get("sslilogging").(string)
	ret.Inst.Sslv2BypassServiceGroup = d.Get("sslv2_bypass_service_group").(string)
	ret.Inst.StartsWithList = getSliceSlbTemplateClientSslStartsWithList(d.Get("starts_with_list").([]interface{}))
	ret.Inst.TemplateCipher = d.Get("template_cipher").(string)
	ret.Inst.TemplateCipherShared = d.Get("template_cipher_shared").(string)
	ret.Inst.TemplateHsm = d.Get("template_hsm").(string)
	ret.Inst.UserNameList = d.Get("user_name_list").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VerifyCertFailAction = d.Get("verify_cert_fail_action").(string)
	ret.Inst.Version = d.Get("version").(int)
	ret.Inst.WebCategory = getObjectSlbTemplateClientSslWebCategory(d.Get("web_category").([]interface{}))
	ret.Inst.WebReputation = getObjectSlbTemplateClientSslWebReputation(d.Get("web_reputation").([]interface{}))
	return ret
}
