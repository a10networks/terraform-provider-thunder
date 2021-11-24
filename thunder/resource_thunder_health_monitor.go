package thunder

//Thunder resource HealthMonitor

import (
	"context"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceHealthMonitor() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHealthMonitorCreate,
		UpdateContext: resourceHealthMonitorUpdate,
		ReadContext:   resourceHealthMonitorRead,
		DeleteContext: resourceHealthMonitorDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dsr_l2_strict": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"retry": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"up_retry": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"override_ipv4": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"override_ipv6": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"override_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"passive": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"status_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"passive_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sample_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"strict_retry_on_server_err_resp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_after_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ssl_ciphers": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ssl_ticket": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ssl_ticket_lifetime": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ssl_version": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ssl_dgversion": {
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
			"method": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"icmp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"transparent": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ipv6": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ip": {
										Type:        schema.TypeString,
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
						"tcp": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"method_tcp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tcp_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port_halfopen": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port_send": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_resp": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"port_contains": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"maintenance": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"maintenance_text": {
										Type:        schema.TypeString,
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
						"udp": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"udp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"udp_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"force_up_with_single_healthcheck": {
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
						"http": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"http_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"http_expect": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"http_response_code": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"response_code_regex": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"http_text": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"text_regex": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"http_host": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"http_maintenance_code": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"http_url": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"url_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"maintenance": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"maintenance_text": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"maintenance_text_regex": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"url_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"post_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"post_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"http_postdata": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"http_postfile": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"http_username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"http_password": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"http_password_string": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"http_kerberos_auth": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"http_kerberos_realm": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"http_kerberos_kdc": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_kerberos_hostip": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"http_kerberos_hostipv6": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"http_kerberos_port": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"http_kerberos_portv6": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ftp": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ftp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ftp_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ftp_username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ftp_password": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ftp_password_string": {
										Type:        schema.TypeString,
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
						"snmp": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"snmp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"snmp_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"community": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"oid": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mib": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"asn": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"operation": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"oper_type": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"smtp": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"smtp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"smtp_domain": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"smtp_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"smtp_starttls": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"mail_from": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"rcpt_to": {
										Type:        schema.TypeString,
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
						"dns": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dns_ip_key": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dns_ipv4_addr": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dns_ipv6_addr": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dns_ipv4_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dns_ipv4_expect": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_ipv4_response": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"dns_ipv4_fqdn": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"dns_ipv4_recurse": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dns_ipv4_tcp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dns_ipv6_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dns_ipv6_expect": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_ipv6_response": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"dns_ipv6_fqdn": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"dns_ipv6_recurse": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dns_ipv6_tcp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dns_domain": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dns_domain_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dns_domain_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dns_domain_expect": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_domain_response": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"dns_domain_fqdn": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"dns_domain_ipv4": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"dns_domain_ipv6": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"dns_domain_recurse": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dns_domain_tcp": {
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
						"pop3": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pop3": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"pop3_username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"pop3_password": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"pop3_password_string": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"pop3_port": {
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
						"imap": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"imap": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"imap_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"imap_username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"imap_password": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"imap_password_string": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"pwd_auth": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"imap_plain": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"imap_cram_md5": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"imap_login": {
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
						"sip": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sip": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"register": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"sip_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"expect_response_code": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"sip_tcp": {
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
						"radius": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"radius": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"radius_username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"radius_password_string": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"radius_secret": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"radius_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"radius_expect": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"radius_response_code": {
										Type:        schema.TypeString,
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
						"ldap": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ldap": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ldap_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ldap_security": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ldap_binddn": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ldap_password": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ldap_password_string": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ldap_run_search": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"base_dn": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ldap_query": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"accept_res_ref": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"accept_not_found": {
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
						"rtsp": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rtsp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"rtspurl": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"rtsp_port": {
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
						"database": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"database": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"database_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"db_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"db_username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"db_password": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"db_password_str": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"db_send": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"db_receive": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"db_row": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"db_column": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"db_receive_integer": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"db_row_integer": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"db_column_integer": {
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
						"external": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"external": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ext_program": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"shared_partition_program": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ext_program_shared": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ext_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ext_arguments": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ext_preference": {
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
						"ntp": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ntp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ntp_port": {
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
						"kerberos_kdc": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kerberos_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kinit": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"kinit_pricipal_name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"kinit_password": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"kinit_kdc": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"tcp_only": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"kadmin": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"kadmin_realm": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"kadmin_pricipal_name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"kadmin_password": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"kadmin_server": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"kadmin_kdc": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"kpasswd": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"kpasswd_pricipal_name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"kpasswd_password": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"kpasswd_server": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"kpasswd_kdc": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"https": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"https": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"web_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"disable_sslv2hello": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"https_host": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"sni": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"https_expect": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"https_response_code": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"response_code_regex": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"https_text": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"text_regex": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"https_url": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"url_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"url_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"post_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"post_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"https_postdata": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"https_postfile": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"https_maintenance_code": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"maintenance": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"maintenance_text": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"maintenance_text_regex": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"https_username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"https_server_cert_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"https_password": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"https_password_string": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"https_kerberos_auth": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"https_kerberos_realm": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"https_kerberos_kdc": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"https_kerberos_hostip": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"https_kerberos_hostipv6": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"https_kerberos_port": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"https_kerberos_portv6": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"cert_key_shared": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
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
									"key_pass_phrase": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"key_phrase": {
										Type:        schema.TypeString,
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
						"tacplus": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tacplus": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tacplus_username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"tacplus_password": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tacplus_password_string": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"tacplus_secret": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tacplus_secret_string": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"tacplus_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tacplus_type": {
										Type:        schema.TypeString,
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
						"compound": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"compound": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"rpn_string": {
										Type:        schema.TypeString,
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
				},
			},
		},
	}
}

func resourceHealthMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating HealthMonitor (Inside resourceHealthMonitorCreate) ")
		name1 := d.Get("name").(string)
		data := dataToHealthMonitor(d)
		logger.Println("[INFO] received formatted data from method data to HealthMonitor --")
		d.SetId(name1)
		err := go_thunder.PostHealthMonitor(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceHealthMonitorRead(ctx, d, meta)

	}
	return diags
}

func resourceHealthMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading HealthMonitor (Inside resourceHealthMonitorRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetHealthMonitor(client.Token, name1, client.Host)
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

func resourceHealthMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying HealthMonitor   (Inside resourceHealthMonitorUpdate) ")
		data := dataToHealthMonitor(d)
		logger.Println("[INFO] received formatted data from method data to HealthMonitor ")
		err := go_thunder.PutHealthMonitor(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceHealthMonitorRead(ctx, d, meta)

	}
	return diags
}

func resourceHealthMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceHealthMonitorDelete) " + name1)
		err := go_thunder.DeleteHealthMonitor(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToHealthMonitor(d *schema.ResourceData) go_thunder.HealthMonitor {
	var vc go_thunder.HealthMonitor
	var c go_thunder.HealthMonitorInstance
	c.HealthMonitorInstanceName = d.Get("name").(string)
	c.HealthMonitorInstanceDsrL2Strict = d.Get("dsr_l2_strict").(int)
	c.HealthMonitorInstanceRetry = d.Get("retry").(int)
	c.HealthMonitorInstanceUpRetry = d.Get("up_retry").(int)
	c.HealthMonitorInstanceOverrideIpv4 = d.Get("override_ipv4").(string)
	c.HealthMonitorInstanceOverrideIpv6 = d.Get("override_ipv6").(string)
	c.HealthMonitorInstanceOverridePort = d.Get("override_port").(int)
	c.HealthMonitorInstancePassive = d.Get("passive").(int)
	c.HealthMonitorInstanceStatusCode = d.Get("status_code").(string)
	c.HealthMonitorInstancePassiveInterval = d.Get("passive_interval").(int)
	c.HealthMonitorInstanceSampleThreshold = d.Get("sample_threshold").(int)
	c.HealthMonitorInstanceThreshold = d.Get("threshold").(int)
	c.HealthMonitorInstanceStrictRetryOnServerErrResp = d.Get("strict_retry_on_server_err_resp").(int)
	c.HealthMonitorInstanceDisableAfterDown = d.Get("disable_after_down").(int)
	c.HealthMonitorInstanceInterval = d.Get("interval").(int)
	c.HealthMonitorInstanceTimeout = d.Get("timeout").(int)
	c.HealthMonitorInstanceSslCiphers = d.Get("ssl_ciphers").(string)
	c.HealthMonitorInstanceSslTicket = d.Get("ssl_ticket").(int)
	c.HealthMonitorInstanceSslTicketLifetime = d.Get("ssl_ticket_lifetime").(int)
	c.HealthMonitorInstanceSslVersion = d.Get("ssl_version").(int)
	c.HealthMonitorInstanceSslDgversion = d.Get("ssl_dgversion").(int)
	c.HealthMonitorInstanceUserTag = d.Get("user_tag").(string)

	var obj1 go_thunder.HealthMonitorInstanceMethod
	prefix1 := "method.0."

	var obj1_1 go_thunder.HealthMonitorInstanceMethodIcmp
	prefix1_1 := prefix1 + "icmp.0."
	obj1_1.HealthMonitorInstanceMethodIcmpIcmp = d.Get(prefix1_1 + "icmp").(int)
	obj1_1.HealthMonitorInstanceMethodIcmpTransparent = d.Get(prefix1_1 + "transparent").(int)
	obj1_1.HealthMonitorInstanceMethodIcmpIpv6 = d.Get(prefix1_1 + "ipv6").(string)
	obj1_1.HealthMonitorInstanceMethodIcmpIP = d.Get(prefix1_1 + "ip").(string)

	obj1.HealthMonitorInstanceMethodIcmpIcmp = obj1_1

	var obj1_2 go_thunder.HealthMonitorInstanceMethodTCP
	prefix1_2 := prefix1 + "tcp.0."
	obj1_2.HealthMonitorInstanceMethodTCPMethodTCP = d.Get(prefix1_2 + "method_tcp").(int)
	obj1_2.HealthMonitorInstanceMethodTCPTCPPort = d.Get(prefix1_2 + "tcp_port").(int)
	obj1_2.HealthMonitorInstanceMethodTCPPortHalfopen = d.Get(prefix1_2 + "port_halfopen").(int)
	obj1_2.HealthMonitorInstanceMethodTCPPortSend = d.Get(prefix1_2 + "port_send").(string)

	var obj1_2_1 go_thunder.HealthMonitorInstanceMethodTCPPortResp
	prefix1_2_1 := prefix1_2 + "port_resp.0."
	obj1_2_1.HealthMonitorInstanceMethodTCPPortRespPortContains = d.Get(prefix1_2_1 + "port_contains").(string)

	obj1_2.HealthMonitorInstanceMethodTCPPortRespPortContains = obj1_2_1

	obj1_2.HealthMonitorInstanceMethodTCPMaintenance = d.Get(prefix1_2 + "maintenance").(int)
	obj1_2.HealthMonitorInstanceMethodTCPMaintenanceText = d.Get(prefix1_2 + "maintenance_text").(string)

	obj1.HealthMonitorInstanceMethodTCPMethodTCP = obj1_2

	var obj1_3 go_thunder.HealthMonitorInstanceMethodUDP
	prefix1_3 := prefix1 + "udp.0."
	obj1_3.HealthMonitorInstanceMethodUDPUDP = d.Get(prefix1_3 + "udp").(int)
	obj1_3.HealthMonitorInstanceMethodUDPUDPPort = d.Get(prefix1_3 + "udp_port").(int)
	obj1_3.HealthMonitorInstanceMethodUDPForceUpWithSingleHealthcheck = d.Get(prefix1_3 + "force_up_with_single_healthcheck").(int)

	obj1.HealthMonitorInstanceMethodUDPUDP = obj1_3

	var obj1_4 go_thunder.HealthMonitorInstanceMethodHTTP
	prefix1_4 := prefix1 + "http.0."
	obj1_4.HealthMonitorInstanceMethodHTTPHTTP = d.Get(prefix1_4 + "http").(int)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPPort = d.Get(prefix1_4 + "http_port").(int)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPExpect = d.Get(prefix1_4 + "http_expect").(int)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPResponseCode = d.Get(prefix1_4 + "http_response_code").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPResponseCodeRegex = d.Get(prefix1_4 + "response_code_regex").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPText = d.Get(prefix1_4 + "http_text").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPTextRegex = d.Get(prefix1_4 + "text_regex").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPHost = d.Get(prefix1_4 + "http_host").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPMaintenanceCode = d.Get(prefix1_4 + "http_maintenance_code").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPURL = d.Get(prefix1_4 + "http_url").(int)
	obj1_4.HealthMonitorInstanceMethodHTTPURLType = d.Get(prefix1_4 + "url_type").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPMaintenance = d.Get(prefix1_4 + "maintenance").(int)
	obj1_4.HealthMonitorInstanceMethodHTTPMaintenanceText = d.Get(prefix1_4 + "maintenance_text").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPMaintenanceTextRegex = d.Get(prefix1_4 + "maintenance_text_regex").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPURLPath = d.Get(prefix1_4 + "url_path").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPPostPath = d.Get(prefix1_4 + "post_path").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPPostType = d.Get(prefix1_4 + "post_type").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPPostdata = d.Get(prefix1_4 + "http_postdata").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPPostfile = d.Get(prefix1_4 + "http_postfile").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPUsername = d.Get(prefix1_4 + "http_username").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPPassword = d.Get(prefix1_4 + "http_password").(int)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPPasswordString = d.Get(prefix1_4 + "http_password_string").(string)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPKerberosAuth = d.Get(prefix1_4 + "http_kerberos_auth").(int)
	obj1_4.HealthMonitorInstanceMethodHTTPHTTPKerberosRealm = d.Get(prefix1_4 + "http_kerberos_realm").(string)

	var obj1_4_1 go_thunder.HealthMonitorInstanceMethodHTTPHTTPKerberosKdc
	prefix1_4_1 := prefix1_4 + "http_kerberos_kdc.0."
	obj1_4_1.HealthMonitorInstanceMethodHTTPHTTPKerberosKdcHTTPKerberosHostip = d.Get(prefix1_4_1 + "http_kerberos_hostip").(string)
	obj1_4_1.HealthMonitorInstanceMethodHTTPHTTPKerberosKdcHTTPKerberosHostipv6 = d.Get(prefix1_4_1 + "http_kerberos_hostipv6").(string)
	obj1_4_1.HealthMonitorInstanceMethodHTTPHTTPKerberosKdcHTTPKerberosPort = d.Get(prefix1_4_1 + "http_kerberos_port").(int)
	obj1_4_1.HealthMonitorInstanceMethodHTTPHTTPKerberosKdcHTTPKerberosPortv6 = d.Get(prefix1_4_1 + "http_kerberos_portv6").(int)

	obj1_4.HealthMonitorInstanceMethodHTTPHTTPKerberosKdcHTTPKerberosHostip = obj1_4_1

	obj1.HealthMonitorInstanceMethodHTTPHTTP = obj1_4

	var obj1_5 go_thunder.HealthMonitorInstanceMethodFtp
	prefix1_5 := prefix1 + "ftp.0."
	obj1_5.HealthMonitorInstanceMethodFtpFtp = d.Get(prefix1_5 + "ftp").(int)
	obj1_5.HealthMonitorInstanceMethodFtpFtpPort = d.Get(prefix1_5 + "ftp_port").(int)
	obj1_5.HealthMonitorInstanceMethodFtpFtpUsername = d.Get(prefix1_5 + "ftp_username").(string)
	obj1_5.HealthMonitorInstanceMethodFtpFtpPassword = d.Get(prefix1_5 + "ftp_password").(int)
	obj1_5.HealthMonitorInstanceMethodFtpFtpPasswordString = d.Get(prefix1_5 + "ftp_password_string").(string)

	obj1.HealthMonitorInstanceMethodFtpFtp = obj1_5

	var obj1_6 go_thunder.HealthMonitorInstanceMethodSnmp
	prefix1_6 := prefix1 + "snmp.0."
	obj1_6.HealthMonitorInstanceMethodSnmpSnmp = d.Get(prefix1_6 + "snmp").(int)
	obj1_6.HealthMonitorInstanceMethodSnmpSnmpPort = d.Get(prefix1_6 + "snmp_port").(int)
	obj1_6.HealthMonitorInstanceMethodSnmpCommunity = d.Get(prefix1_6 + "community").(string)

	var obj1_6_1 go_thunder.HealthMonitorInstanceMethodSnmpOid
	prefix1_6_1 := prefix1_6 + "oid.0."
	obj1_6_1.HealthMonitorInstanceMethodSnmpOidMib = d.Get(prefix1_6_1 + "mib").(string)
	obj1_6_1.HealthMonitorInstanceMethodSnmpOidAsn = d.Get(prefix1_6_1 + "asn").(string)

	obj1_6.HealthMonitorInstanceMethodSnmpOidMib = obj1_6_1

	var obj1_6_2 go_thunder.HealthMonitorInstanceMethodSnmpOperation
	prefix1_6_2 := prefix1_6 + "operation.0."
	obj1_6_2.HealthMonitorInstanceMethodSnmpOperationOperType = d.Get(prefix1_6_2 + "oper_type").(string)

	obj1_6.HealthMonitorInstanceMethodSnmpOperationOperType = obj1_6_2

	obj1.HealthMonitorInstanceMethodSnmpSnmp = obj1_6

	var obj1_7 go_thunder.HealthMonitorInstanceMethodSMTP
	prefix1_7 := prefix1 + "smtp.0."
	obj1_7.HealthMonitorInstanceMethodSMTPSMTP = d.Get(prefix1_7 + "smtp").(int)
	obj1_7.HealthMonitorInstanceMethodSMTPSMTPDomain = d.Get(prefix1_7 + "smtp_domain").(string)
	obj1_7.HealthMonitorInstanceMethodSMTPSMTPPort = d.Get(prefix1_7 + "smtp_port").(int)
	obj1_7.HealthMonitorInstanceMethodSMTPSMTPStarttls = d.Get(prefix1_7 + "smtp_starttls").(int)
	obj1_7.HealthMonitorInstanceMethodSMTPMailFrom = d.Get(prefix1_7 + "mail_from").(string)
	obj1_7.HealthMonitorInstanceMethodSMTPRcptTo = d.Get(prefix1_7 + "rcpt_to").(string)

	obj1.HealthMonitorInstanceMethodSMTPSMTP = obj1_7

	var obj1_8 go_thunder.HealthMonitorInstanceMethodDNS
	prefix1_8 := prefix1 + "dns.0."
	obj1_8.HealthMonitorInstanceMethodDNSDNS = d.Get(prefix1_8 + "dns").(int)
	obj1_8.HealthMonitorInstanceMethodDNSDNSIPKey = d.Get(prefix1_8 + "dns_ip_key").(int)
	obj1_8.HealthMonitorInstanceMethodDNSDNSIpv4Addr = d.Get(prefix1_8 + "dns_ipv4_addr").(string)
	obj1_8.HealthMonitorInstanceMethodDNSDNSIpv6Addr = d.Get(prefix1_8 + "dns_ipv6_addr").(string)
	obj1_8.HealthMonitorInstanceMethodDNSDNSIpv4Port = d.Get(prefix1_8 + "dns_ipv4_port").(int)

	var obj1_8_1 go_thunder.HealthMonitorInstanceMethodDNSDNSIpv4Expect
	prefix1_8_1 := prefix1_8 + "dns_ipv4_expect.0."
	obj1_8_1.HealthMonitorInstanceMethodDNSDNSIpv4ExpectDNSIpv4Response = d.Get(prefix1_8_1 + "dns_ipv4_response").(string)
	obj1_8_1.HealthMonitorInstanceMethodDNSDNSIpv4ExpectDNSIpv4Fqdn = d.Get(prefix1_8_1 + "dns_ipv4_fqdn").(string)

	obj1_8.HealthMonitorInstanceMethodDNSDNSIpv4ExpectDNSIpv4Response = obj1_8_1

	obj1_8.HealthMonitorInstanceMethodDNSDNSIpv4Recurse = d.Get(prefix1_8 + "dns_ipv4_recurse").(string)
	obj1_8.HealthMonitorInstanceMethodDNSDNSIpv4TCP = d.Get(prefix1_8 + "dns_ipv4_tcp").(int)
	obj1_8.HealthMonitorInstanceMethodDNSDNSIpv6Port = d.Get(prefix1_8 + "dns_ipv6_port").(int)

	var obj1_8_2 go_thunder.HealthMonitorInstanceMethodDNSDNSIpv6Expect
	prefix1_8_2 := prefix1_8 + "dns_ipv6_expect.0."
	obj1_8_2.HealthMonitorInstanceMethodDNSDNSIpv6ExpectDNSIpv6Response = d.Get(prefix1_8_2 + "dns_ipv6_response").(string)
	obj1_8_2.HealthMonitorInstanceMethodDNSDNSIpv6ExpectDNSIpv6Fqdn = d.Get(prefix1_8_2 + "dns_ipv6_fqdn").(string)

	obj1_8.HealthMonitorInstanceMethodDNSDNSIpv6ExpectDNSIpv6Response = obj1_8_2

	obj1_8.HealthMonitorInstanceMethodDNSDNSIpv6Recurse = d.Get(prefix1_8 + "dns_ipv6_recurse").(string)
	obj1_8.HealthMonitorInstanceMethodDNSDNSIpv6TCP = d.Get(prefix1_8 + "dns_ipv6_tcp").(int)
	obj1_8.HealthMonitorInstanceMethodDNSDNSDomain = d.Get(prefix1_8 + "dns_domain").(string)
	obj1_8.HealthMonitorInstanceMethodDNSDNSDomainPort = d.Get(prefix1_8 + "dns_domain_port").(int)
	obj1_8.HealthMonitorInstanceMethodDNSDNSDomainType = d.Get(prefix1_8 + "dns_domain_type").(string)

	var obj1_8_3 go_thunder.HealthMonitorInstanceMethodDNSDNSDomainExpect
	prefix1_8_3 := prefix1_8 + "dns_domain_expect.0."
	obj1_8_3.HealthMonitorInstanceMethodDNSDNSDomainExpectDNSDomainResponse = d.Get(prefix1_8_3 + "dns_domain_response").(string)
	obj1_8_3.HealthMonitorInstanceMethodDNSDNSDomainExpectDNSDomainFqdn = d.Get(prefix1_8_3 + "dns_domain_fqdn").(string)
	obj1_8_3.HealthMonitorInstanceMethodDNSDNSDomainExpectDNSDomainIpv4 = d.Get(prefix1_8_3 + "dns_domain_ipv4").(string)
	obj1_8_3.HealthMonitorInstanceMethodDNSDNSDomainExpectDNSDomainIpv6 = d.Get(prefix1_8_3 + "dns_domain_ipv6").(string)

	obj1_8.HealthMonitorInstanceMethodDNSDNSDomainExpectDNSDomainResponse = obj1_8_3

	obj1_8.HealthMonitorInstanceMethodDNSDNSDomainRecurse = d.Get(prefix1_8 + "dns_domain_recurse").(string)
	obj1_8.HealthMonitorInstanceMethodDNSDNSDomainTCP = d.Get(prefix1_8 + "dns_domain_tcp").(int)

	obj1.HealthMonitorInstanceMethodDNSDNS = obj1_8

	var obj1_9 go_thunder.HealthMonitorInstanceMethodPop3
	prefix1_9 := prefix1 + "pop3.0."
	obj1_9.HealthMonitorInstanceMethodPop3Pop3 = d.Get(prefix1_9 + "pop3").(int)
	obj1_9.HealthMonitorInstanceMethodPop3Pop3Username = d.Get(prefix1_9 + "pop3_username").(string)
	obj1_9.HealthMonitorInstanceMethodPop3Pop3Password = d.Get(prefix1_9 + "pop3_password").(int)
	obj1_9.HealthMonitorInstanceMethodPop3Pop3PasswordString = d.Get(prefix1_9 + "pop3_password_string").(string)
	obj1_9.HealthMonitorInstanceMethodPop3Pop3Port = d.Get(prefix1_9 + "pop3_port").(int)

	obj1.HealthMonitorInstanceMethodPop3Pop3 = obj1_9

	var obj1_10 go_thunder.HealthMonitorInstanceMethodImap
	prefix1_10 := prefix1 + "imap.0."
	obj1_10.HealthMonitorInstanceMethodImapImap = d.Get(prefix1_10 + "imap").(int)
	obj1_10.HealthMonitorInstanceMethodImapImapPort = d.Get(prefix1_10 + "imap_port").(int)
	obj1_10.HealthMonitorInstanceMethodImapImapUsername = d.Get(prefix1_10 + "imap_username").(string)
	obj1_10.HealthMonitorInstanceMethodImapImapPassword = d.Get(prefix1_10 + "imap_password").(int)
	obj1_10.HealthMonitorInstanceMethodImapImapPasswordString = d.Get(prefix1_10 + "imap_password_string").(string)
	obj1_10.HealthMonitorInstanceMethodImapPwdAuth = d.Get(prefix1_10 + "pwd_auth").(int)
	obj1_10.HealthMonitorInstanceMethodImapImapPlain = d.Get(prefix1_10 + "imap_plain").(int)
	obj1_10.HealthMonitorInstanceMethodImapImapCramMd5 = d.Get(prefix1_10 + "imap_cram_md5").(int)
	obj1_10.HealthMonitorInstanceMethodImapImapLogin = d.Get(prefix1_10 + "imap_login").(int)

	obj1.HealthMonitorInstanceMethodImapImap = obj1_10

	var obj1_11 go_thunder.HealthMonitorInstanceMethodSip
	prefix1_11 := prefix1 + "sip.0."
	obj1_11.HealthMonitorInstanceMethodSipSip = d.Get(prefix1_11 + "sip").(int)
	obj1_11.HealthMonitorInstanceMethodSipRegister = d.Get(prefix1_11 + "register").(int)
	obj1_11.HealthMonitorInstanceMethodSipSipPort = d.Get(prefix1_11 + "sip_port").(int)
	obj1_11.HealthMonitorInstanceMethodSipExpectResponseCode = d.Get(prefix1_11 + "expect_response_code").(string)
	obj1_11.HealthMonitorInstanceMethodSipSipTCP = d.Get(prefix1_11 + "sip_tcp").(int)

	obj1.HealthMonitorInstanceMethodSipSip = obj1_11

	var obj1_12 go_thunder.HealthMonitorInstanceMethodRadius
	prefix1_12 := prefix1 + "radius.0."
	obj1_12.HealthMonitorInstanceMethodRadiusRadius = d.Get(prefix1_12 + "radius").(int)
	obj1_12.HealthMonitorInstanceMethodRadiusRadiusUsername = d.Get(prefix1_12 + "radius_username").(string)
	obj1_12.HealthMonitorInstanceMethodRadiusRadiusPasswordString = d.Get(prefix1_12 + "radius_password_string").(string)
	obj1_12.HealthMonitorInstanceMethodRadiusRadiusSecret = d.Get(prefix1_12 + "radius_secret").(string)
	obj1_12.HealthMonitorInstanceMethodRadiusRadiusPort = d.Get(prefix1_12 + "radius_port").(int)
	obj1_12.HealthMonitorInstanceMethodRadiusRadiusExpect = d.Get(prefix1_12 + "radius_expect").(int)
	obj1_12.HealthMonitorInstanceMethodRadiusRadiusResponseCode = d.Get(prefix1_12 + "radius_response_code").(string)

	obj1.HealthMonitorInstanceMethodRadiusRadius = obj1_12

	var obj1_13 go_thunder.HealthMonitorInstanceMethodLdap
	prefix1_13 := prefix1 + "ldap.0."
	obj1_13.HealthMonitorInstanceMethodLdapLdap = d.Get(prefix1_13 + "ldap").(int)
	obj1_13.HealthMonitorInstanceMethodLdapLdapPort = d.Get(prefix1_13 + "ldap_port").(int)
	obj1_13.HealthMonitorInstanceMethodLdapLdapSecurity = d.Get(prefix1_13 + "ldap_security").(string)
	obj1_13.HealthMonitorInstanceMethodLdapLdapBinddn = d.Get(prefix1_13 + "ldap_binddn").(string)
	obj1_13.HealthMonitorInstanceMethodLdapLdapPassword = d.Get(prefix1_13 + "ldap_password").(int)
	obj1_13.HealthMonitorInstanceMethodLdapLdapPasswordString = d.Get(prefix1_13 + "ldap_password_string").(string)
	obj1_13.HealthMonitorInstanceMethodLdapLdapRunSearch = d.Get(prefix1_13 + "ldap_run_search").(int)
	obj1_13.HealthMonitorInstanceMethodLdapBaseDN = d.Get(prefix1_13 + "base_dn").(string)
	obj1_13.HealthMonitorInstanceMethodLdapLdapQuery = d.Get(prefix1_13 + "ldap_query").(string)
	obj1_13.HealthMonitorInstanceMethodLdapAcceptResRef = d.Get(prefix1_13 + "accept_res_ref").(int)
	obj1_13.HealthMonitorInstanceMethodLdapAcceptNotFound = d.Get(prefix1_13 + "accept_not_found").(int)

	obj1.HealthMonitorInstanceMethodLdapLdap = obj1_13

	var obj1_14 go_thunder.HealthMonitorInstanceMethodRtsp
	prefix1_14 := prefix1 + "rtsp.0."
	obj1_14.HealthMonitorInstanceMethodRtspRtsp = d.Get(prefix1_14 + "rtsp").(int)
	obj1_14.HealthMonitorInstanceMethodRtspRtspurl = d.Get(prefix1_14 + "rtspurl").(string)
	obj1_14.HealthMonitorInstanceMethodRtspRtspPort = d.Get(prefix1_14 + "rtsp_port").(int)

	obj1.HealthMonitorInstanceMethodRtspRtsp = obj1_14

	var obj1_15 go_thunder.HealthMonitorInstanceMethodDatabase
	prefix1_15 := prefix1 + "database.0."
	obj1_15.HealthMonitorInstanceMethodDatabaseDatabase = d.Get(prefix1_15 + "database").(int)
	obj1_15.HealthMonitorInstanceMethodDatabaseDatabaseName = d.Get(prefix1_15 + "database_name").(string)
	obj1_15.HealthMonitorInstanceMethodDatabaseDbName = d.Get(prefix1_15 + "db_name").(string)
	obj1_15.HealthMonitorInstanceMethodDatabaseDbUsername = d.Get(prefix1_15 + "db_username").(string)
	obj1_15.HealthMonitorInstanceMethodDatabaseDbPassword = d.Get(prefix1_15 + "db_password").(int)
	obj1_15.HealthMonitorInstanceMethodDatabaseDbPasswordStr = d.Get(prefix1_15 + "db_password_str").(string)
	obj1_15.HealthMonitorInstanceMethodDatabaseDbSend = d.Get(prefix1_15 + "db_send").(string)
	obj1_15.HealthMonitorInstanceMethodDatabaseDbReceive = d.Get(prefix1_15 + "db_receive").(string)
	obj1_15.HealthMonitorInstanceMethodDatabaseDbRow = d.Get(prefix1_15 + "db_row").(int)
	obj1_15.HealthMonitorInstanceMethodDatabaseDbColumn = d.Get(prefix1_15 + "db_column").(int)
	obj1_15.HealthMonitorInstanceMethodDatabaseDbReceiveInteger = d.Get(prefix1_15 + "db_receive_integer").(int)
	obj1_15.HealthMonitorInstanceMethodDatabaseDbRowInteger = d.Get(prefix1_15 + "db_row_integer").(int)
	obj1_15.HealthMonitorInstanceMethodDatabaseDbColumnInteger = d.Get(prefix1_15 + "db_column_integer").(int)

	obj1.HealthMonitorInstanceMethodDatabaseDatabase = obj1_15

	var obj1_16 go_thunder.HealthMonitorInstanceMethodExternal
	prefix1_16 := prefix1 + "external.0."
	obj1_16.HealthMonitorInstanceMethodExternalExternal = d.Get(prefix1_16 + "external").(int)
	obj1_16.HealthMonitorInstanceMethodExternalExtProgram = d.Get(prefix1_16 + "ext_program").(string)
	obj1_16.HealthMonitorInstanceMethodExternalSharedPartitionProgram = d.Get(prefix1_16 + "shared_partition_program").(int)
	obj1_16.HealthMonitorInstanceMethodExternalExtProgramShared = d.Get(prefix1_16 + "ext_program_shared").(string)
	obj1_16.HealthMonitorInstanceMethodExternalExtPort = d.Get(prefix1_16 + "ext_port").(int)
	obj1_16.HealthMonitorInstanceMethodExternalExtArguments = d.Get(prefix1_16 + "ext_arguments").(string)
	obj1_16.HealthMonitorInstanceMethodExternalExtPreference = d.Get(prefix1_16 + "ext_preference").(int)

	obj1.HealthMonitorInstanceMethodExternalExternal = obj1_16

	var obj1_17 go_thunder.HealthMonitorInstanceMethodNtp
	prefix1_17 := prefix1 + "ntp.0."
	obj1_17.HealthMonitorInstanceMethodNtpNtp = d.Get(prefix1_17 + "ntp").(int)
	obj1_17.HealthMonitorInstanceMethodNtpNtpPort = d.Get(prefix1_17 + "ntp_port").(int)

	obj1.HealthMonitorInstanceMethodNtpNtp = obj1_17

	var obj1_18 go_thunder.HealthMonitorInstanceMethodKerberosKdc
	prefix1_18 := prefix1 + "kerberos_kdc.0."

	var obj1_18_1 go_thunder.HealthMonitorInstanceMethodKerberosKdcKerberosCfg
	prefix1_18_1 := prefix1_18 + "kerberos_cfg.0."
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKinit = d.Get(prefix1_18_1 + "kinit").(int)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKinitPricipalName = d.Get(prefix1_18_1 + "kinit_pricipal_name").(string)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKinitPassword = d.Get(prefix1_18_1 + "kinit_password").(string)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKinitKdc = d.Get(prefix1_18_1 + "kinit_kdc").(string)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgTCPOnly = d.Get(prefix1_18_1 + "tcp_only").(int)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKadmin = d.Get(prefix1_18_1 + "kadmin").(int)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKadminRealm = d.Get(prefix1_18_1 + "kadmin_realm").(string)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKadminPricipalName = d.Get(prefix1_18_1 + "kadmin_pricipal_name").(string)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKadminPassword = d.Get(prefix1_18_1 + "kadmin_password").(string)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKadminServer = d.Get(prefix1_18_1 + "kadmin_server").(string)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKadminKdc = d.Get(prefix1_18_1 + "kadmin_kdc").(string)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKpasswd = d.Get(prefix1_18_1 + "kpasswd").(int)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKpasswdPricipalName = d.Get(prefix1_18_1 + "kpasswd_pricipal_name").(string)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKpasswdPassword = d.Get(prefix1_18_1 + "kpasswd_password").(string)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKpasswdServer = d.Get(prefix1_18_1 + "kpasswd_server").(string)
	obj1_18_1.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKpasswdKdc = d.Get(prefix1_18_1 + "kpasswd_kdc").(string)

	obj1_18.HealthMonitorInstanceMethodKerberosKdcKerberosCfgKinit = obj1_18_1

	obj1.HealthMonitorInstanceMethodKerberosKdcKerberosCfg = obj1_18

	var obj1_19 go_thunder.HealthMonitorInstanceMethodHTTPS
	prefix1_19 := prefix1 + "https.0."
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPS = d.Get(prefix1_19 + "https").(int)
	obj1_19.HealthMonitorInstanceMethodHTTPSWebPort = d.Get(prefix1_19 + "web_port").(int)
	obj1_19.HealthMonitorInstanceMethodHTTPSDisableSslv2Hello = d.Get(prefix1_19 + "disable_sslv2hello").(int)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSHost = d.Get(prefix1_19 + "https_host").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSSni = d.Get(prefix1_19 + "sni").(int)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSExpect = d.Get(prefix1_19 + "https_expect").(int)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSResponseCode = d.Get(prefix1_19 + "https_response_code").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSResponseCodeRegex = d.Get(prefix1_19 + "response_code_regex").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSText = d.Get(prefix1_19 + "https_text").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSTextRegex = d.Get(prefix1_19 + "text_regex").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSURL = d.Get(prefix1_19 + "https_url").(int)
	obj1_19.HealthMonitorInstanceMethodHTTPSURLType = d.Get(prefix1_19 + "url_type").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSURLPath = d.Get(prefix1_19 + "url_path").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSPostPath = d.Get(prefix1_19 + "post_path").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSPostType = d.Get(prefix1_19 + "post_type").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSPostdata = d.Get(prefix1_19 + "https_postdata").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSPostfile = d.Get(prefix1_19 + "https_postfile").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSMaintenanceCode = d.Get(prefix1_19 + "https_maintenance_code").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSMaintenance = d.Get(prefix1_19 + "maintenance").(int)
	obj1_19.HealthMonitorInstanceMethodHTTPSMaintenanceText = d.Get(prefix1_19 + "maintenance_text").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSMaintenanceTextRegex = d.Get(prefix1_19 + "maintenance_text_regex").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSUsername = d.Get(prefix1_19 + "https_username").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSServerCertName = d.Get(prefix1_19 + "https_server_cert_name").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSPassword = d.Get(prefix1_19 + "https_password").(int)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSPasswordString = d.Get(prefix1_19 + "https_password_string").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSKerberosAuth = d.Get(prefix1_19 + "https_kerberos_auth").(int)
	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSKerberosRealm = d.Get(prefix1_19 + "https_kerberos_realm").(string)

	var obj1_19_1 go_thunder.HealthMonitorInstanceMethodHTTPSHTTPSKerberosKdc
	prefix1_19_1 := prefix1_19 + "https_kerberos_kdc.0."
	obj1_19_1.HealthMonitorInstanceMethodHTTPSHTTPSKerberosKdcHTTPSKerberosHostip = d.Get(prefix1_19_1 + "https_kerberos_hostip").(string)
	obj1_19_1.HealthMonitorInstanceMethodHTTPSHTTPSKerberosKdcHTTPSKerberosHostipv6 = d.Get(prefix1_19_1 + "https_kerberos_hostipv6").(string)
	obj1_19_1.HealthMonitorInstanceMethodHTTPSHTTPSKerberosKdcHTTPSKerberosPort = d.Get(prefix1_19_1 + "https_kerberos_port").(int)
	obj1_19_1.HealthMonitorInstanceMethodHTTPSHTTPSKerberosKdcHTTPSKerberosPortv6 = d.Get(prefix1_19_1 + "https_kerberos_portv6").(int)

	obj1_19.HealthMonitorInstanceMethodHTTPSHTTPSKerberosKdcHTTPSKerberosHostip = obj1_19_1

	obj1_19.HealthMonitorInstanceMethodHTTPSCertKeyShared = d.Get(prefix1_19 + "cert_key_shared").(int)
	obj1_19.HealthMonitorInstanceMethodHTTPSCert = d.Get(prefix1_19 + "cert").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSKey = d.Get(prefix1_19 + "key").(string)
	obj1_19.HealthMonitorInstanceMethodHTTPSKeyPassPhrase = d.Get(prefix1_19 + "key_pass_phrase").(int)
	obj1_19.HealthMonitorInstanceMethodHTTPSKeyPhrase = d.Get(prefix1_19 + "key_phrase").(string)

	obj1.HealthMonitorInstanceMethodHTTPSHTTPS = obj1_19

	var obj1_20 go_thunder.HealthMonitorInstanceMethodTacplus
	prefix1_20 := prefix1 + "tacplus.0."
	obj1_20.HealthMonitorInstanceMethodTacplusTacplus = d.Get(prefix1_20 + "tacplus").(int)
	obj1_20.HealthMonitorInstanceMethodTacplusTacplusUsername = d.Get(prefix1_20 + "tacplus_username").(string)
	obj1_20.HealthMonitorInstanceMethodTacplusTacplusPassword = d.Get(prefix1_20 + "tacplus_password").(int)
	obj1_20.HealthMonitorInstanceMethodTacplusTacplusPasswordString = d.Get(prefix1_20 + "tacplus_password_string").(string)
	obj1_20.HealthMonitorInstanceMethodTacplusTacplusSecret = d.Get(prefix1_20 + "tacplus_secret").(int)
	obj1_20.HealthMonitorInstanceMethodTacplusTacplusSecretString = d.Get(prefix1_20 + "tacplus_secret_string").(string)
	obj1_20.HealthMonitorInstanceMethodTacplusTacplusPort = d.Get(prefix1_20 + "tacplus_port").(int)
	obj1_20.HealthMonitorInstanceMethodTacplusTacplusType = d.Get(prefix1_20 + "tacplus_type").(string)

	obj1.HealthMonitorInstanceMethodTacplusTacplus = obj1_20

	var obj1_21 go_thunder.HealthMonitorInstanceMethodCompound
	prefix1_21 := prefix1 + "compound.0."
	obj1_21.HealthMonitorInstanceMethodCompoundCompound = d.Get(prefix1_21 + "compound").(int)
	obj1_21.HealthMonitorInstanceMethodCompoundRpnString = d.Get(prefix1_21 + "rpn_string").(string)

	obj1.HealthMonitorInstanceMethodCompoundCompound = obj1_21

	c.HealthMonitorInstanceMethodIcmp = obj1

	vc.HealthMonitorInstanceName = c
	return vc
}
