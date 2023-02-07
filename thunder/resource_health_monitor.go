package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceHealthMonitor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor`: Define the Health Monitor object\n\n",
		CreateContext: resourceHealthMonitorCreate,
		UpdateContext: resourceHealthMonitorUpdate,
		ReadContext:   resourceHealthMonitorRead,
		DeleteContext: resourceHealthMonitorDelete,
		Schema: map[string]*schema.Schema{
			"default_state_up": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Initial health state will default to UP",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"disable_after_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the target if health check failed",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"dsr_l2_strict": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable strict L2dsr health-check",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the Healthcheck Interval (Interval Value, in seconds (default 5))",
				ValidateFunc: validation.IntBetween(1, 180),
			},
			"method": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"icmp": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "ICMP type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"transparent": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply transparent mode",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ipv6": {
										Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address of destination behind monitored node",
										ValidateFunc:  validation.IsIPv6Address,
										ConflictsWith: []string{"method.0.icmp.0.ip"},
									},
									"ip": {
										Type: schema.TypeString, Optional: true, Description: "Specify IPv4 address of destination behind monitored node",
										ValidateFunc:  validation.IsIPv4Address,
										ConflictsWith: []string{"method.0.icmp.0.ipv6"},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"tcp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"method_tcp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"tcp_port": {
										Type: schema.TypeInt, Optional: true, Description: "Specify TCP port (Specify port number)",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"port_halfopen": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set TCP SYN check",
										ValidateFunc:  validation.IntBetween(0, 1),
										ConflictsWith: []string{"method.0.tcp.0.port_send"},
									},
									"port_send": {
										Type: schema.TypeString, Optional: true, Description: "Send a string to server (Specify the string)",
										ValidateFunc:  validation.StringLenBetween(1, 511),
										ConflictsWith: []string{"method.0.tcp.0.port_halfopen"},
									},
									"port_resp": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"port_contains": {
													Type: schema.TypeString, Optional: true, Description: "Mark server up if response string contains string (Specify the string)",
													ValidateFunc: validation.StringLenBetween(1, 127),
												},
											},
										},
									},
									"maintenance": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify response text for maintenance",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"maintenance_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify text for maintenance",
										ValidateFunc: validation.StringLenBetween(1, 127),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"udp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"udp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "UDP type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"udp_port": {
										Type: schema.TypeInt, Optional: true, Description: "Specify UDP port (Specify port number)",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"force_up_with_single_healthcheck": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force Up with no response at the first time",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"http": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "HTTP type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"http_port": {
										Type: schema.TypeInt, Optional: true, Default: 80, Description: "Specify HTTP Port (Specify port number (default 80))",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"http_expect": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify what you expect from the response message",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"http_response_code": {
										Type: schema.TypeString, Optional: true, Description: "Specify response code range (e.g. 200,400-430) (Format is xx,xx-xx (xx between [100, 899]))",
										ValidateFunc:  validation.StringLenBetween(1, 31),
										ConflictsWith: []string{"method.0.http.0.http_text"},
									},
									"response_code_regex": {
										Type: schema.TypeString, Optional: true, Description: "Specify response code range with Regex (code with Regex, such as [2-5][0-9][0-9])",
										ValidateFunc: validation.StringLenBetween(1, 127),
									},
									"http_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify text expected",
										ValidateFunc:  validation.StringLenBetween(1, 127),
										ConflictsWith: []string{"method.0.http.0.http_response_code", "method.0.http.0.text_regex"},
									},
									"text_regex": {
										Type: schema.TypeString, Optional: true, Description: "Specify text expected  with Regex",
										ValidateFunc:  validation.StringLenBetween(1, 127),
										ConflictsWith: []string{"method.0.http.0.http_text"},
									},
									"http_host": {
										Type: schema.TypeString, Optional: true, Description: "Specify \"Host:\" header used in request (enclose IPv6 address in [])",
										ValidateFunc: validation.StringLenBetween(1, 63),
									},
									"http_maintenance_code": {
										Type: schema.TypeString, Optional: true, Description: "Specify response code for maintenance (Format is xx,xx-xx (xx between [100, 899]))",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"http_url": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify URL string, default is GET /",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"url_type": {
										Type: schema.TypeString, Optional: true, Description: "'GET': HTTP GET method; 'POST': HTTP POST method; 'HEAD': HTTP HEAD method;",
										ValidateFunc: validation.StringInSlice([]string{"GET", "POST", "HEAD"}, false),
									},
									"maintenance": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify response text for maintenance",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"maintenance_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify text for maintenance",
										ValidateFunc:  validation.StringLenBetween(1, 127),
										ConflictsWith: []string{"method.0.http.0.maintenance_text_regex"},
									},
									"maintenance_text_regex": {
										Type: schema.TypeString, Optional: true, Description: "Specify Regex text for maintenance",
										ValidateFunc:  validation.StringLenBetween(1, 127),
										ConflictsWith: []string{"method.0.http.0.maintenance_text"},
									},
									"url_path": {
										Type: schema.TypeString, Optional: true, Description: "Specify URL path, default is \"/\"",
										ValidateFunc: validation.StringLenBetween(1, 500),
									},
									"post_path": {
										Type: schema.TypeString, Optional: true, Description: "Specify URL path, default is \"/\"",
										ValidateFunc: validation.StringLenBetween(1, 500),
									},
									"post_type": {
										Type: schema.TypeString, Optional: true, Description: "'postdata': Specify the HTTP post data; 'postfile': Specify the HTTP post data;",
										ValidateFunc: validation.StringInSlice([]string{"postdata", "postfile"}, false),
									},
									"http_postdata": {
										Type: schema.TypeString, Optional: true, Description: "Specify the HTTP post data (Input post data here)",
										ValidateFunc: validation.StringLenBetween(1, 255),
									},
									"http_postfile": {
										Type: schema.TypeString, Optional: true, Description: "Specify the HTTP post data (Input post data file name here)",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"http_username": {
										Type: schema.TypeString, Optional: true, Description: "Specify the username",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"http_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"http_password_string": {
										Type: schema.TypeString, Optional: true, Description: "Specify password, '' means empty password",
									},
									//omit encrypted field: 'http_encrypted'
									"http_kerberos_auth": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Http Kerberos Auth",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"http_kerberos_realm": {
										Type: schema.TypeString, Optional: true, Description: "Specify realm of Kerberos server",
										ValidateFunc: validation.StringLenBetween(1, 63),
									},
									"http_kerberos_kdc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_kerberos_hostip": {
													Type: schema.TypeString, Optional: true, Description: "Kdc's hostname(length:1-31) or IP address",
													ValidateFunc:  validation.StringLenBetween(1, 31),
													ConflictsWith: []string{"method.0.http.0.http_kerberos_kdc.0.http_kerberos_hostipv6"},
												},
												"http_kerberos_hostipv6": {
													Type: schema.TypeString, Optional: true, Description: "Server's IPV6 address",
													ValidateFunc:  validation.IsIPv6Address,
													ConflictsWith: []string{"method.0.http.0.http_kerberos_kdc.0.http_kerberos_hostip"},
												},
												"http_kerberos_port": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the kdc port",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"http_kerberos_portv6": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the kdc port",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ftp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ftp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "FTP type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ftp_port": {
										Type: schema.TypeInt, Optional: true, Default: 21, Description: "Specify FTP port (Specify port number, default is 21)",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"ftp_username": {
										Type: schema.TypeString, Optional: true, Description: "Specify the username",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"ftp_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ftp_password_string": {
										Type: schema.TypeString, Optional: true, Description: "Specify the user password, '' means empty password",
									},
									//omit encrypted field: 'ftp_encrypted'
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"snmp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"snmp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "SNMP type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"snmp_port": {
										Type: schema.TypeInt, Optional: true, Default: 161, Description: "Specify SNMP port, default is 161 (Port Number)",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"community": {
										Type: schema.TypeString, Optional: true, Default: "public", Description: "Specify SNMP community, default is \"public\" (Community String)",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"oid": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mib": {
													Type: schema.TypeString, Optional: true, Description: "'sysDescr': The MIB-2 OID of system description, 1.1.0; 'sysUpTime': The MIB-2 OID of system up time, 1.3.0; 'sysName': The MIB-2 OID of system nume, 1.5.0;",
													ValidateFunc:  validation.StringInSlice([]string{"sysDescr", "sysUpTime", "sysName"}, false),
													ConflictsWith: []string{"method.0.snmp.0.oid.0.asn"},
												},
												"asn": {
													Type: schema.TypeString, Optional: true, Description: "Specify the format in ASN.1 style",
													ValidateFunc:  validation.StringLenBetween(1, 128),
													ConflictsWith: []string{"method.0.snmp.0.oid.0.mib"},
												},
											},
										},
									},
									"operation": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"oper_type": {
													Type: schema.TypeString, Optional: true, Description: "'getnext': Get-Next-Request command; 'get': Get-Request command;",
													ValidateFunc: validation.StringInSlice([]string{"getnext", "get"}, false),
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"smtp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"smtp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "SMTP type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"smtp_domain": {
										Type: schema.TypeString, Optional: true, Description: "Specify domain name of 'helo' command",
										ValidateFunc: validation.StringLenBetween(1, 254),
									},
									"smtp_port": {
										Type: schema.TypeInt, Optional: true, Default: 25, Description: "Specify SMTP port, default is 25 (Port Number (default 25))",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"smtp_starttls": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Check the STARTTLS support at helo response",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"mail_from": {
										Type: schema.TypeString, Optional: true, Description: "Specify SMTP Sender",
										ValidateFunc: validation.StringLenBetween(1, 255),
									},
									"rcpt_to": {
										Type: schema.TypeString, Optional: true, Description: "Specify SMTP Receiver",
										ValidateFunc: validation.StringLenBetween(1, 255),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"dns": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "DNS type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"dns_ip_key": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reverse DNS lookup (Specify IPv4 or IPv6 address)",
										ValidateFunc:  validation.IntBetween(0, 1),
										ConflictsWith: []string{"method.0.dns.0.dns_domain"},
									},
									"dns_ipv4_addr": {
										Type: schema.TypeString, Optional: true, Description: "Specify IPv4 address",
										ValidateFunc:  validation.IsIPv4Address,
										ConflictsWith: []string{"method.0.dns.0.dns_ipv6_addr"},
									},
									"dns_ipv6_addr": {
										Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address",
										ValidateFunc:  validation.IsIPv6Address,
										ConflictsWith: []string{"method.0.dns.0.dns_ipv4_addr"},
									},
									"dns_ipv4_port": {
										Type: schema.TypeInt, Optional: true, Default: 53, Description: "Specify DNS port, default is 53 (DNS Port(default 53))",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"dns_ipv4_expect": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_ipv4_response": {
													Type: schema.TypeString, Optional: true, Description: "Specify response code range (e.g. 0,1-5) (Format is xx,xx-xx (xx between [0,15]))",
													ValidateFunc:  validation.StringLenBetween(1, 31),
													ConflictsWith: []string{"method.0.dns.0.dns_ipv4_expect.0.dns_ipv4_fqdn"},
												},
												"dns_ipv4_fqdn": {
													Type: schema.TypeString, Optional: true, Description: "Specify fully qualified domain name expected in DNS response answer",
													ValidateFunc:  validation.StringLenBetween(1, 255),
													ConflictsWith: []string{"method.0.dns.0.dns_ipv4_expect.0.dns_ipv4_response"},
												},
											},
										},
									},
									"dns_ipv4_recurse": {
										Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'enabled': Set the recursion bit; 'disabled': Clear the recursion bit;",
										ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
									},
									"dns_ipv4_tcp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure DNS transport over TCP, default is UDP",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"dns_ipv6_port": {
										Type: schema.TypeInt, Optional: true, Default: 53, Description: "Specify DNS port, default is 53 (DNS Port(default 53))",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"dns_ipv6_expect": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_ipv6_response": {
													Type: schema.TypeString, Optional: true, Description: "Specify response code range (e.g. 0,1-5) (Format is xx,xx-xx (xx between [0,15]))",
													ValidateFunc:  validation.StringLenBetween(1, 31),
													ConflictsWith: []string{"method.0.dns.0.dns_ipv6_expect.0.dns_ipv6_fqdn"},
												},
												"dns_ipv6_fqdn": {
													Type: schema.TypeString, Optional: true, Description: "Specify fully qualified domain name expected in DNS response answer",
													ValidateFunc:  validation.StringLenBetween(1, 255),
													ConflictsWith: []string{"method.0.dns.0.dns_ipv6_expect.0.dns_ipv6_response"},
												},
											},
										},
									},
									"dns_ipv6_recurse": {
										Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'enabled': Set the recursion bit; 'disabled': Clear the recursion bit;",
										ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
									},
									"dns_ipv6_tcp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure DNS transport over TCP, default is UDP",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"dns_domain": {
										Type: schema.TypeString, Optional: true, Description: "Specify fully qualified domain name of the host",
										ValidateFunc:  validation.StringLenBetween(1, 255),
										ConflictsWith: []string{"method.0.dns.0.dns_ip_key"},
									},
									"dns_domain_port": {
										Type: schema.TypeInt, Optional: true, Default: 53, Description: "Specify DNS port, default is 53 (DNS Port(default 53))",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"dns_domain_type": {
										Type: schema.TypeString, Optional: true, Default: "A", Description: "'A': Used for storing Ipv4 address (default); 'CNAME': Canonical name for a DNS alias; 'SOA': Start of authority; 'PTR': Domain name pointer; 'MX': Mail exchanger; 'TXT': Text string; 'AAAA': Used for storing Ipv6 128-bits address;",
										ValidateFunc: validation.StringInSlice([]string{"A", "CNAME", "SOA", "PTR", "MX", "TXT", "AAAA"}, false),
									},
									"dns_domain_expect": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_domain_response": {
													Type: schema.TypeString, Optional: true, Description: "Specify response code range (e.g. 0,1-5) (Format is xx,xx-xx (xx between [0,15]))",
													ValidateFunc:  validation.StringLenBetween(1, 31),
													ConflictsWith: []string{"method.0.dns.0.dns_domain_expect.0.dns_domain_fqdn", "method.0.dns.0.dns_domain_expect.0.dns_domain_ipv4", "method.0.dns.0.dns_domain_expect.0.dns_domain_ipv6"},
												},
												"dns_domain_fqdn": {
													Type: schema.TypeString, Optional: true, Description: "Specify fully qualified domain name expected in DNS response answer",
													ValidateFunc:  validation.StringLenBetween(1, 255),
													ConflictsWith: []string{"method.0.dns.0.dns_domain_expect.0.dns_domain_response", "method.0.dns.0.dns_domain_expect.0.dns_domain_ipv4", "method.0.dns.0.dns_domain_expect.0.dns_domain_ipv6"},
												},
												"dns_domain_ipv4": {
													Type: schema.TypeString, Optional: true, Description: "Specify expected resolved IPv4 address",
													ValidateFunc:  validation.IsIPv4Address,
													ConflictsWith: []string{"method.0.dns.0.dns_domain_expect.0.dns_domain_response", "method.0.dns.0.dns_domain_expect.0.dns_domain_fqdn", "method.0.dns.0.dns_domain_expect.0.dns_domain_ipv6"},
												},
												"dns_domain_ipv6": {
													Type: schema.TypeString, Optional: true, Description: "Specify expected resolved IPv6 address",
													ValidateFunc:  validation.IsIPv6Address,
													ConflictsWith: []string{"method.0.dns.0.dns_domain_expect.0.dns_domain_response", "method.0.dns.0.dns_domain_expect.0.dns_domain_fqdn", "method.0.dns.0.dns_domain_expect.0.dns_domain_ipv4"},
												},
											},
										},
									},
									"dns_domain_recurse": {
										Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'enabled': Set the recursion bit; 'disabled': Clear the recursion bit;",
										ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
									},
									"dns_domain_tcp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure DNS transport over TCP, default is UDP",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"pop3": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pop3": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "POP3 type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"pop3_username": {
										Type: schema.TypeString, Optional: true, Description: "Specify the username",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"pop3_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"pop3_password_string": {
										Type: schema.TypeString, Optional: true, Description: "Specify the user password, '' means empty password",
									},
									//omit encrypted field: 'pop3_encrypted'
									"pop3_port": {
										Type: schema.TypeInt, Optional: true, Default: 110, Description: "Specify the POP3 port, default is 110 (Port Number (default 110))",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"imap": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"imap": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "IMAP type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"imap_port": {
										Type: schema.TypeInt, Optional: true, Default: 143, Description: "Specify the IMAP port, default is 143 (Port Number (default 143))",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"imap_username": {
										Type: schema.TypeString, Optional: true, Description: "Specify the username",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"imap_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"imap_password_string": {
										Type: schema.TypeString, Optional: true, Description: "Configure password, '' means empty password",
									},
									//omit encrypted field: 'imap_encrypted'
									"pwd_auth": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the Authentication method",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"imap_plain": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Plain text",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"imap_cram_md5": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Challenge-response authentication mechanism",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"imap_login": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Simple login",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"sip": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sip": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"register": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send SIP REGISTER message, default is to send OPTION message",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"sip_port": {
										Type: schema.TypeInt, Optional: true, Default: 5060, Description: "Specify the SIP port, default is 5060 (Port Number)",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"expect_response_code": {
										Type: schema.TypeString, Optional: true, Description: "Specify accepted response codes (e.g. 200, 400-430, any) (Format is xxx,xxx-xxx,any (xxx between [100,899]))",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"sip_tcp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use TCP for transmission, default is UDP",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"radius": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"radius": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "RADIUS type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"radius_username": {
										Type: schema.TypeString, Optional: true, Description: "Specify the username",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"radius_password_string": {
										Type: schema.TypeString, Optional: true, Description: "Configure password, '' means empty password",
									},
									//omit encrypted field: 'radius_encrypted'
									"radius_secret": {
										Type: schema.TypeString, Optional: true, Description: "Configure shared secret of RADIUS server",
									},
									//omit encrypted field: 'radius_secret_encrypted'
									"radius_port": {
										Type: schema.TypeInt, Optional: true, Default: 1812, Description: "Specify the RADIUS port, default is 1812 (Port number (default 1812))",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"radius_expect": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify what you expect from the response message",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"radius_response_code": {
										Type: schema.TypeString, Optional: true, Description: "Specify response code range (e.g. 2,4-7) (Format is xx,xx-xx (xx between [1, 13]))",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ldap": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ldap": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "LDAP type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ldap_port": {
										Type: schema.TypeInt, Optional: true, Default: 389, Description: "Specify the LDAP port (Speciry port number, default is 389, or 636 if LDAP over SSL)",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"ldap_security": {
										Type: schema.TypeString, Optional: true, Description: "'overssl': Set LDAP over SSL; 'StartTLS': LDAP switch to TLS;",
										ValidateFunc: validation.StringInSlice([]string{"overssl", "StartTLS"}, false),
									},
									"ldap_binddn": {
										Type: schema.TypeString, Optional: true, Description: "Specify the distinguished name for bindRequest (LDAP DN distinguished name)",
										ValidateFunc: validation.StringLenBetween(1, 127),
									},
									"ldap_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ldap_password_string": {
										Type: schema.TypeString, Optional: true, Description: "Configure password, '' means empty password",
									},
									//omit encrypted field: 'ldap_encrypted'
									"ldap_run_search": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify a query to be executed",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"basedn": {
										Type: schema.TypeString, Optional: true, Description: "Specify LDAP DN distinguished name",
										ValidateFunc: validation.StringLenBetween(1, 127),
									},
									"ldap_query": {
										Type: schema.TypeString, Optional: true, Description: "LDAP query to be excuted",
										ValidateFunc: validation.StringLenBetween(1, 255),
									},
									"acceptresref": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mark server up on receiving a search result reference response",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"acceptnotfound": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mark server up on receiving a not-found response",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"rtsp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rtsp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "RTSP type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"rtspurl": {
										Type: schema.TypeString, Optional: true, Description: "Specify URL string (Specify the path on the server)",
										ValidateFunc: validation.StringLenBetween(1, 127),
									},
									"rtsp_port": {
										Type: schema.TypeInt, Optional: true, Default: 554, Description: "Specify RTSP port, default is 554 (Port Number (default 554))",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"database": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"database": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "DATABASE type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"database_name": {
										Type: schema.TypeString, Optional: true, Description: "'mssql': Specify MSSQL database; 'mysql': Specify MySQL database; 'oracle': Specify Oracle database; 'postgresql': Specify PostgreSQL database;",
										ValidateFunc: validation.StringInSlice([]string{"mssql", "mysql", "oracle", "postgresql"}, false),
									},
									"db_name": {
										Type: schema.TypeString, Optional: true, Description: "Specify the database name",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"db_username": {
										Type: schema.TypeString, Optional: true, Description: "Specify the username",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"db_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"db_password_str": {
										Type: schema.TypeString, Optional: true, Description: "Configure password",
									},
									//omit encrypted field: 'db_encrypted'
									"db_send": {
										Type: schema.TypeString, Optional: true, Description: "Specify the SQL query",
										ValidateFunc: validation.StringLenBetween(1, 127),
									},
									"db_receive": {
										Type: schema.TypeString, Optional: true, Description: "Specify the response string",
										ValidateFunc:  validation.StringLenBetween(1, 31),
										ConflictsWith: []string{"method.0.database.0.db_receive_integer"},
									},
									"db_row": {
										Type: schema.TypeInt, Optional: true, Description: "Specify the row number for receiving",
										ValidateFunc: validation.IntBetween(1, 10),
									},
									"db_column": {
										Type: schema.TypeInt, Optional: true, Description: "Specify the column number for receiving",
										ValidateFunc: validation.IntBetween(1, 10),
									},
									"db_receive_integer": {
										Type: schema.TypeInt, Optional: true, Description: "Specify the response integer",
										ValidateFunc:  validation.IntBetween(0, 2147483647),
										ConflictsWith: []string{"method.0.database.0.db_receive"},
									},
									"db_row_integer": {
										Type: schema.TypeInt, Optional: true, Description: "Specify the row number for receiving",
										ValidateFunc: validation.IntBetween(1, 10),
									},
									"db_column_integer": {
										Type: schema.TypeInt, Optional: true, Description: "Specify the column number for receiving",
										ValidateFunc: validation.IntBetween(1, 10),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"external": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"external": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "EXTERNAL type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ext_program": {
										Type: schema.TypeString, Optional: true, Description: "Specify external application (Program name)",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"shared_partition_program": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "external application from shared partition",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ext_program_shared": {
										Type: schema.TypeString, Optional: true, Description: "Specify external application (Program name)",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"ext_port": {
										Type: schema.TypeInt, Optional: true, Description: "Specify the server port (Port Number)",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"ext_arguments": {
										Type: schema.TypeString, Optional: true, Description: "Specify external application's arguments (Application arguments)",
										ValidateFunc: validation.StringLenBetween(1, 127),
									},
									"ext_preference": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Get server's perference",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ntp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ntp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "NTP type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ntp_port": {
										Type: schema.TypeInt, Optional: true, Default: 123, Description: "Specify the NTP port, default is 123 (Port Number (default 123))",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"kerberos_kdc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kerberos_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kinit": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Kerberos KDC",
													ValidateFunc:  validation.IntBetween(0, 1),
													ConflictsWith: []string{"method.0.kerberos_kdc.0.kerberos_cfg.0.kadmin", "method.0.kerberos_kdc.0.kerberos_cfg.0.kpasswd"},
												},
												"kinit_pricipal_name": {
													Type: schema.TypeString, Optional: true, Description: "Specify the principal name",
													ValidateFunc: validation.StringLenBetween(1, 127),
												},
												"kinit_password": {
													Type: schema.TypeString, Optional: true, Description: "Password",
												},
												//omit encrypted field: 'kinit_encrypted'
												"kinit_kdc": {
													Type: schema.TypeString, Optional: true, Description: "Specify the kdc server, host|ip [:port]",
													ValidateFunc: validation.StringLenBetween(1, 63),
												},
												"tcp_only": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the kerberos tcp only",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"kadmin": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Kerberos admin",
													ValidateFunc:  validation.IntBetween(0, 1),
													ConflictsWith: []string{"method.0.kerberos_kdc.0.kerberos_cfg.0.kinit", "method.0.kerberos_kdc.0.kerberos_cfg.0.kpasswd"},
												},
												"kadmin_realm": {
													Type: schema.TypeString, Optional: true, Description: "Specify the realm",
													ValidateFunc: validation.StringLenBetween(1, 63),
												},
												"kadmin_pricipal_name": {
													Type: schema.TypeString, Optional: true, Description: "Specify the principal name",
													ValidateFunc: validation.StringLenBetween(1, 127),
												},
												"kadmin_password": {
													Type: schema.TypeString, Optional: true, Description: "Password",
												},
												//omit encrypted field: 'kadmin_encrypted'
												"kadmin_server": {
													Type: schema.TypeString, Optional: true, Description: "Specify the admin server, host|ip [:port]",
													ValidateFunc: validation.StringLenBetween(1, 63),
												},
												"kadmin_kdc": {
													Type: schema.TypeString, Optional: true, Description: "Specify the kdc server, host|ip [:port]",
													ValidateFunc: validation.StringLenBetween(1, 63),
												},
												"kpasswd": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Kerberos change passwd",
													ValidateFunc:  validation.IntBetween(0, 1),
													ConflictsWith: []string{"method.0.kerberos_kdc.0.kerberos_cfg.0.kinit", "method.0.kerberos_kdc.0.kerberos_cfg.0.kadmin"},
												},
												"kpasswd_pricipal_name": {
													Type: schema.TypeString, Optional: true, Description: "Specify the principal name",
													ValidateFunc: validation.StringLenBetween(1, 127),
												},
												"kpasswd_password": {
													Type: schema.TypeString, Optional: true, Description: "Password",
												},
												//omit encrypted field: 'kpasswd_encrypted'
												"kpasswd_server": {
													Type: schema.TypeString, Optional: true, Description: "Specify the Kerberos password server, host|ip [:port]",
													ValidateFunc: validation.StringLenBetween(1, 63),
												},
												"kpasswd_kdc": {
													Type: schema.TypeString, Optional: true, Description: "Specify the kdc server, host|ip [:port]",
													ValidateFunc: validation.StringLenBetween(1, 63),
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"https": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"https": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "HTTPS type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"web_port": {
										Type: schema.TypeInt, Optional: true, Default: 443, Description: "Specify HTTPS port (Port Number (default 443))",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"disable_sslv2hello": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable SSLv2Hello for HTTPs",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"https_host": {
										Type: schema.TypeString, Optional: true, Description: "Specify \"Host:\" header used in request (enclose IPv6 address in [])",
										ValidateFunc: validation.StringLenBetween(1, 63),
									},
									"sni": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Server Name Indication for HTTPs",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"https_expect": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify what you expect from the response message",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"https_response_code": {
										Type: schema.TypeString, Optional: true, Description: "Specify response code range (e.g. 200,400-430) (Format is xx,xx-xx (xx between [100, 899])",
										ValidateFunc:  validation.StringLenBetween(1, 31),
										ConflictsWith: []string{"method.0.https.0.https_text"},
									},
									"response_code_regex": {
										Type: schema.TypeString, Optional: true, Description: "Specify response code range with Regex (code with Regex, such as [2-5][0-9][0-9])",
										ValidateFunc: validation.StringLenBetween(1, 127),
									},
									"https_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify text expected",
										ValidateFunc:  validation.StringLenBetween(1, 127),
										ConflictsWith: []string{"method.0.https.0.https_response_code", "method.0.https.0.text_regex"},
									},
									"text_regex": {
										Type: schema.TypeString, Optional: true, Description: "Specify text expected  with Regex",
										ValidateFunc:  validation.StringLenBetween(1, 127),
										ConflictsWith: []string{"method.0.https.0.https_text"},
									},
									"https_url": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify URL string, default is GET /",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"url_type": {
										Type: schema.TypeString, Optional: true, Description: "'GET': HTTP GET method; 'POST': HTTP POST method; 'HEAD': HTTP HEAD method;",
										ValidateFunc: validation.StringInSlice([]string{"GET", "POST", "HEAD"}, false),
									},
									"url_path": {
										Type: schema.TypeString, Optional: true, Description: "Specify URL path, default is \"/\"",
										ValidateFunc: validation.StringLenBetween(1, 500),
									},
									"post_path": {
										Type: schema.TypeString, Optional: true, Description: "Specify URL path, default is \"/\"",
										ValidateFunc: validation.StringLenBetween(1, 500),
									},
									"post_type": {
										Type: schema.TypeString, Optional: true, Description: "'postdata': Specify the HTTP post data; 'postfile': Specify the HTTP post data;",
										ValidateFunc: validation.StringInSlice([]string{"postdata", "postfile"}, false),
									},
									"https_postdata": {
										Type: schema.TypeString, Optional: true, Description: "Specify the HTTP post data (Input post data here)",
										ValidateFunc: validation.StringLenBetween(1, 255),
									},
									"https_postfile": {
										Type: schema.TypeString, Optional: true, Description: "Specify the HTTP post data (Input post data file name here)",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"https_maintenance_code": {
										Type: schema.TypeString, Optional: true, Description: "Specify response code for maintenance (Format is xx,xx-xx (xx between [100, 899])",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"maintenance": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify response text for maintenance",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"maintenance_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify text for maintenance",
										ValidateFunc:  validation.StringLenBetween(1, 127),
										ConflictsWith: []string{"method.0.https.0.maintenance_text_regex"},
									},
									"maintenance_text_regex": {
										Type: schema.TypeString, Optional: true, Description: "Specify Regex text for maintenance",
										ValidateFunc:  validation.StringLenBetween(1, 127),
										ConflictsWith: []string{"method.0.https.0.maintenance_text"},
									},
									"https_username": {
										Type: schema.TypeString, Optional: true, Description: "Specify the username",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"https_server_cert_name": {
										Type: schema.TypeString, Optional: true, Description: "Expect Server Cert commonName",
										ValidateFunc: validation.StringLenBetween(1, 63),
									},
									"https_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"https_password_string": {
										Type: schema.TypeString, Optional: true, Description: "Configure password, '' means empty password",
									},
									//omit encrypted field: 'https_encrypted'
									"https_kerberos_auth": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Https Kerberos Auth",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"https_kerberos_realm": {
										Type: schema.TypeString, Optional: true, Description: "Specify realm of Kerberos server",
										ValidateFunc: validation.StringLenBetween(1, 63),
									},
									"https_kerberos_kdc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"https_kerberos_hostip": {
													Type: schema.TypeString, Optional: true, Description: "Kdc's hostname(length:1-31) or IP address",
													ValidateFunc:  validation.StringLenBetween(1, 31),
													ConflictsWith: []string{"method.0.https.0.https_kerberos_kdc.0.https_kerberos_hostipv6"},
												},
												"https_kerberos_hostipv6": {
													Type: schema.TypeString, Optional: true, Description: "Server's IPV6 address",
													ValidateFunc:  validation.IsIPv6Address,
													ConflictsWith: []string{"method.0.https.0.https_kerberos_kdc.0.https_kerberos_hostip"},
												},
												"https_kerberos_port": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the kdc port",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"https_kerberos_portv6": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the kdc port",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
											},
										},
									},
									"cert_key_shared": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Select shared partition",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"cert": {
										Type: schema.TypeString, Optional: true, Description: "Specify client certificate (Certificate name)",
										ValidateFunc: validation.StringLenBetween(1, 255),
									},
									"key": {
										Type: schema.TypeString, Optional: true, Description: "Specify client private key (Key name)",
										ValidateFunc: validation.StringLenBetween(1, 255),
									},
									"key_pass_phrase": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Client private key password phrase",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"key_phrase": {
										Type: schema.TypeString, Optional: true, Description: "Password Phrase",
									},
									//omit encrypted field: 'https_key_encrypted'
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"tacplus": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tacplus": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "TACACS+ type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"tacplus_username": {
										Type: schema.TypeString, Optional: true, Description: "Specify the username",
										ValidateFunc: validation.StringLenBetween(1, 31),
									},
									"tacplus_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"tacplus_password_string": {
										Type: schema.TypeString, Optional: true, Description: "Configure password, '' means empty password",
									},
									//omit encrypted field: 'tacplus_encrypted'
									"tacplus_secret": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the shared secret of TACACS+ server",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"tacplus_secret_string": {
										Type: schema.TypeString, Optional: true, Description: "Shared Crypto Key",
									},
									//omit encrypted field: 'secret_encrypted'
									"tacplus_port": {
										Type: schema.TypeInt, Optional: true, Default: 49, Description: "Specify the TACACS+ port, default 49 (Port number (default 49))",
										ValidateFunc: validation.IntBetween(1, 65534),
									},
									"tacplus_type": {
										Type: schema.TypeString, Optional: true, Default: "inbound-ascii-login", Description: "'inbound-ascii-login': Specify Inbound ASCII Login type;",
										ValidateFunc: validation.StringInSlice([]string{"inbound-ascii-login"}, false),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"compound": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"compound": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Compound type",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"rpn_string": {
										Type: schema.TypeString, Optional: true, Description: "Enter Reverse Polish Notation, example: sub hm1 sub hm2 and",
										ValidateFunc: validation.StringLenBetween(1, 512),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Monitor Name",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"override_ipv4": {
				Type: schema.TypeString, Optional: true, Description: "Override implicitly inherited IPv4 address from target",
				ValidateFunc: validation.IsIPv4Address,
			},
			"override_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Override implicitly inherited IPv6 address from target",
				ValidateFunc: validation.IsIPv6Address,
			},
			"override_port": {
				Type: schema.TypeInt, Optional: true, Description: "Override implicitly inherited port from target (Port number (1-65534))",
				ValidateFunc: validation.IntBetween(1, 65534),
			},
			"passive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify passive mode",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"passive_interval": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Interval to do manual health checking while in passive mode (Specify value in seconds (Default is 10 s))",
				ValidateFunc: validation.IntBetween(1, 180),
			},
			"retry": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the Healthcheck Retries (Retry Count (default 3))",
				ValidateFunc: validation.IntBetween(1, 10),
			},
			"sample_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 50, Description: "Number of samples in one epoch above which passive HC is enabled. If below or equal to the threshold, passive HC is disabled (Specify number of samples in one second (Default is 50). If the number of samples is 0, no action is taken)",
				ValidateFunc: validation.IntBetween(1, 10000),
			},
			"ssl_ciphers": {
				Type: schema.TypeString, Optional: true, Default: "DEFAULT", Description: "Specify OpenSSL Cipher Suite name(s) for Health check (OpenSSL Cipher Suite(s) (Eg: AES128-SHA256), if the cipher is invalid, would give information at HM down reason)",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"ssl_dgversion": {
				Type: schema.TypeInt, Optional: true, Default: 31, Description: "Lower TLS/SSL version can be downgraded",
				ValidateFunc: validation.IntBetween(31, 34),
			},
			"ssl_ticket": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL-Ticket Session Resumption",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"ssl_ticket_lifetime": {
				Type: schema.TypeInt, Optional: true, Description: "SSL-Ticket lifetime (seconds)",
				ValidateFunc: validation.IntBetween(1, 604800),
			},
			"ssl_version": {
				Type: schema.TypeInt, Optional: true, Default: 34, Description: "TLS/SSL version (TLS/SSL version: 31-TLSv1.0, 32-TLSv1.1, 33-TLSv1.2 and 34-TLSv1.3)",
				ValidateFunc: validation.IntBetween(31, 34),
			},
			"status_code": {
				Type: schema.TypeString, Optional: true, Description: "'status-code-2xx': Enable passive mode with 2xx http status code; 'status-code-non-5xx': Enable passive mode with non-5xx http status code;",
				ValidateFunc: validation.StringInSlice([]string{"status-code-2xx", "status-code-non-5xx"}, false),
			},
			"strict_retry_on_server_err_resp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Require strictly retry",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"threshold": {
				Type: schema.TypeInt, Optional: true, Default: 75, Description: "Threshold percentage above which passive mode is enabled (Specify percentage (Default is 75%))",
				ValidateFunc: validation.IntBetween(0, 100),
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the Healthcheck Timeout (Timeout Value, in seconds(default 5), Timeout should be less than or equal to interval)",
				ValidateFunc: validation.IntBetween(1, 180),
			},
			"up_retry": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the Healthcheck Retries before declaring target up (Up-retry count (default 1))",
				ValidateFunc: validation.IntBetween(1, 10),
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}

func resourceHealthMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectHealthMonitorMethod(d []interface{}) edpt.HealthMonitorMethod {
	var ret edpt.HealthMonitorMethod
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Icmp = getObjectHealthMonitorMethodIcmp(in["icmp"].([]interface{}))
		ret.Tcp = getObjectHealthMonitorMethodTcp(in["tcp"].([]interface{}))
		ret.Udp = getObjectHealthMonitorMethodUdp(in["udp"].([]interface{}))
		ret.Http = getObjectHealthMonitorMethodHttp(in["http"].([]interface{}))
		ret.Ftp = getObjectHealthMonitorMethodFtp(in["ftp"].([]interface{}))
		ret.Snmp = getObjectHealthMonitorMethodSnmp(in["snmp"].([]interface{}))
		ret.Smtp = getObjectHealthMonitorMethodSmtp(in["smtp"].([]interface{}))
		ret.Dns = getObjectHealthMonitorMethodDns(in["dns"].([]interface{}))
		ret.Pop3 = getObjectHealthMonitorMethodPop3(in["pop3"].([]interface{}))
		ret.Imap = getObjectHealthMonitorMethodImap(in["imap"].([]interface{}))
		ret.Sip = getObjectHealthMonitorMethodSip(in["sip"].([]interface{}))
		ret.Radius = getObjectHealthMonitorMethodRadius(in["radius"].([]interface{}))
		ret.Ldap = getObjectHealthMonitorMethodLdap(in["ldap"].([]interface{}))
		ret.Rtsp = getObjectHealthMonitorMethodRtsp(in["rtsp"].([]interface{}))
		ret.Database = getObjectHealthMonitorMethodDatabase(in["database"].([]interface{}))
		ret.External = getObjectHealthMonitorMethodExternal(in["external"].([]interface{}))
		ret.Ntp = getObjectHealthMonitorMethodNtp(in["ntp"].([]interface{}))
		ret.KerberosKdc = getObjectHealthMonitorMethodKerberosKdc(in["kerberos_kdc"].([]interface{}))
		ret.Https = getObjectHealthMonitorMethodHttps(in["https"].([]interface{}))
		ret.Tacplus = getObjectHealthMonitorMethodTacplus(in["tacplus"].([]interface{}))
		ret.Compound = getObjectHealthMonitorMethodCompound(in["compound"].([]interface{}))
	}
	return ret
}

func getObjectHealthMonitorMethodIcmp(d []interface{}) edpt.HealthMonitorMethodIcmp {
	var ret edpt.HealthMonitorMethodIcmp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Icmp = in["icmp"].(int)
		ret.Transparent = in["transparent"].(int)
		ret.Ipv6 = in["ipv6"].(string)
		ret.Ip = in["ip"].(string)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodTcp(d []interface{}) edpt.HealthMonitorMethodTcp {
	var ret edpt.HealthMonitorMethodTcp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.MethodTcp = in["method_tcp"].(int)
		ret.TcpPort = in["tcp_port"].(int)
		ret.PortHalfopen = in["port_halfopen"].(int)
		ret.PortSend = in["port_send"].(string)
		ret.PortResp = getObjectHealthMonitorMethodTcpPortResp(in["port_resp"].([]interface{}))
		ret.Maintenance = in["maintenance"].(int)
		ret.MaintenanceText = in["maintenance_text"].(string)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodTcpPortResp(d []interface{}) edpt.HealthMonitorMethodTcpPortResp {
	var ret edpt.HealthMonitorMethodTcpPortResp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.PortContains = in["port_contains"].(string)
	}
	return ret
}

func getObjectHealthMonitorMethodUdp(d []interface{}) edpt.HealthMonitorMethodUdp {
	var ret edpt.HealthMonitorMethodUdp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Udp = in["udp"].(int)
		ret.UdpPort = in["udp_port"].(int)
		ret.ForceUpWithSingleHealthcheck = in["force_up_with_single_healthcheck"].(int)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodHttp(d []interface{}) edpt.HealthMonitorMethodHttp {
	var ret edpt.HealthMonitorMethodHttp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Http = in["http"].(int)
		ret.HttpPort = in["http_port"].(int)
		ret.HttpExpect = in["http_expect"].(int)
		ret.HttpResponseCode = in["http_response_code"].(string)
		ret.ResponseCodeRegex = in["response_code_regex"].(string)
		ret.HttpText = in["http_text"].(string)
		ret.TextRegex = in["text_regex"].(string)
		ret.HttpHost = in["http_host"].(string)
		ret.HttpMaintenanceCode = in["http_maintenance_code"].(string)
		ret.HttpUrl = in["http_url"].(int)
		ret.UrlType = in["url_type"].(string)
		ret.Maintenance = in["maintenance"].(int)
		ret.MaintenanceText = in["maintenance_text"].(string)
		ret.MaintenanceTextRegex = in["maintenance_text_regex"].(string)
		ret.UrlPath = in["url_path"].(string)
		ret.PostPath = in["post_path"].(string)
		ret.PostType = in["post_type"].(string)
		ret.HttpPostdata = in["http_postdata"].(string)
		ret.HttpPostfile = in["http_postfile"].(string)
		ret.HttpUsername = in["http_username"].(string)
		ret.HttpPassword = in["http_password"].(int)
		ret.HttpPasswordString = in["http_password_string"].(string)
		//omit http_encrypted
		ret.HttpKerberosAuth = in["http_kerberos_auth"].(int)
		ret.HttpKerberosRealm = in["http_kerberos_realm"].(string)
		ret.HttpKerberosKdc = getObjectHealthMonitorMethodHttpHttpKerberosKdc(in["http_kerberos_kdc"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodHttpHttpKerberosKdc(d []interface{}) edpt.HealthMonitorMethodHttpHttpKerberosKdc {
	var ret edpt.HealthMonitorMethodHttpHttpKerberosKdc
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.HttpKerberosHostip = in["http_kerberos_hostip"].(string)
		ret.HttpKerberosHostipv6 = in["http_kerberos_hostipv6"].(string)
		ret.HttpKerberosPort = in["http_kerberos_port"].(int)
		ret.HttpKerberosPortv6 = in["http_kerberos_portv6"].(int)
	}
	return ret
}

func getObjectHealthMonitorMethodFtp(d []interface{}) edpt.HealthMonitorMethodFtp {
	var ret edpt.HealthMonitorMethodFtp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ftp = in["ftp"].(int)
		ret.FtpPort = in["ftp_port"].(int)
		ret.FtpUsername = in["ftp_username"].(string)
		ret.FtpPassword = in["ftp_password"].(int)
		ret.FtpPasswordString = in["ftp_password_string"].(string)
		//omit ftp_encrypted
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodSnmp(d []interface{}) edpt.HealthMonitorMethodSnmp {
	var ret edpt.HealthMonitorMethodSnmp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Snmp = in["snmp"].(int)
		ret.SnmpPort = in["snmp_port"].(int)
		ret.Community = in["community"].(string)
		ret.Oid = getObjectHealthMonitorMethodSnmpOid(in["oid"].([]interface{}))
		ret.Operation = getObjectHealthMonitorMethodSnmpOperation(in["operation"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodSnmpOid(d []interface{}) edpt.HealthMonitorMethodSnmpOid {
	var ret edpt.HealthMonitorMethodSnmpOid
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Mib = in["mib"].(string)
		ret.Asn = in["asn"].(string)
	}
	return ret
}

func getObjectHealthMonitorMethodSnmpOperation(d []interface{}) edpt.HealthMonitorMethodSnmpOperation {
	var ret edpt.HealthMonitorMethodSnmpOperation
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.OperType = in["oper_type"].(string)
	}
	return ret
}

func getObjectHealthMonitorMethodSmtp(d []interface{}) edpt.HealthMonitorMethodSmtp {
	var ret edpt.HealthMonitorMethodSmtp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Smtp = in["smtp"].(int)
		ret.SmtpDomain = in["smtp_domain"].(string)
		ret.SmtpPort = in["smtp_port"].(int)
		ret.SmtpStarttls = in["smtp_starttls"].(int)
		ret.MailFrom = in["mail_from"].(string)
		ret.RcptTo = in["rcpt_to"].(string)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodDns(d []interface{}) edpt.HealthMonitorMethodDns {
	var ret edpt.HealthMonitorMethodDns
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Dns = in["dns"].(int)
		ret.DnsIpKey = in["dns_ip_key"].(int)
		ret.DnsIpv4Addr = in["dns_ipv4_addr"].(string)
		ret.DnsIpv6Addr = in["dns_ipv6_addr"].(string)
		ret.DnsIpv4Port = in["dns_ipv4_port"].(int)
		ret.DnsIpv4Expect = getObjectHealthMonitorMethodDnsDnsIpv4Expect(in["dns_ipv4_expect"].([]interface{}))
		ret.DnsIpv4Recurse = in["dns_ipv4_recurse"].(string)
		ret.DnsIpv4Tcp = in["dns_ipv4_tcp"].(int)
		ret.DnsIpv6Port = in["dns_ipv6_port"].(int)
		ret.DnsIpv6Expect = getObjectHealthMonitorMethodDnsDnsIpv6Expect(in["dns_ipv6_expect"].([]interface{}))
		ret.DnsIpv6Recurse = in["dns_ipv6_recurse"].(string)
		ret.DnsIpv6Tcp = in["dns_ipv6_tcp"].(int)
		ret.DnsDomain = in["dns_domain"].(string)
		ret.DnsDomainPort = in["dns_domain_port"].(int)
		ret.DnsDomainType = in["dns_domain_type"].(string)
		ret.DnsDomainExpect = getObjectHealthMonitorMethodDnsDnsDomainExpect(in["dns_domain_expect"].([]interface{}))
		ret.DnsDomainRecurse = in["dns_domain_recurse"].(string)
		ret.DnsDomainTcp = in["dns_domain_tcp"].(int)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodDnsDnsIpv4Expect(d []interface{}) edpt.HealthMonitorMethodDnsDnsIpv4Expect {
	var ret edpt.HealthMonitorMethodDnsDnsIpv4Expect
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.DnsIpv4Response = in["dns_ipv4_response"].(string)
		ret.DnsIpv4Fqdn = in["dns_ipv4_fqdn"].(string)
	}
	return ret
}

func getObjectHealthMonitorMethodDnsDnsIpv6Expect(d []interface{}) edpt.HealthMonitorMethodDnsDnsIpv6Expect {
	var ret edpt.HealthMonitorMethodDnsDnsIpv6Expect
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.DnsIpv6Response = in["dns_ipv6_response"].(string)
		ret.DnsIpv6Fqdn = in["dns_ipv6_fqdn"].(string)
	}
	return ret
}

func getObjectHealthMonitorMethodDnsDnsDomainExpect(d []interface{}) edpt.HealthMonitorMethodDnsDnsDomainExpect {
	var ret edpt.HealthMonitorMethodDnsDnsDomainExpect
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.DnsDomainResponse = in["dns_domain_response"].(string)
		ret.DnsDomainFqdn = in["dns_domain_fqdn"].(string)
		ret.DnsDomainIpv4 = in["dns_domain_ipv4"].(string)
		ret.DnsDomainIpv6 = in["dns_domain_ipv6"].(string)
	}
	return ret
}

func getObjectHealthMonitorMethodPop3(d []interface{}) edpt.HealthMonitorMethodPop3 {
	var ret edpt.HealthMonitorMethodPop3
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Pop3 = in["pop3"].(int)
		ret.Pop3Username = in["pop3_username"].(string)
		ret.Pop3Password = in["pop3_password"].(int)
		ret.Pop3PasswordString = in["pop3_password_string"].(string)
		//omit pop3_encrypted
		ret.Pop3Port = in["pop3_port"].(int)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodImap(d []interface{}) edpt.HealthMonitorMethodImap {
	var ret edpt.HealthMonitorMethodImap
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Imap = in["imap"].(int)
		ret.ImapPort = in["imap_port"].(int)
		ret.ImapUsername = in["imap_username"].(string)
		ret.ImapPassword = in["imap_password"].(int)
		ret.ImapPasswordString = in["imap_password_string"].(string)
		//omit imap_encrypted
		ret.PwdAuth = in["pwd_auth"].(int)
		ret.ImapPlain = in["imap_plain"].(int)
		ret.ImapCramMd5 = in["imap_cram_md5"].(int)
		ret.ImapLogin = in["imap_login"].(int)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodSip(d []interface{}) edpt.HealthMonitorMethodSip {
	var ret edpt.HealthMonitorMethodSip
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Sip = in["sip"].(int)
		ret.Register = in["register"].(int)
		ret.SipPort = in["sip_port"].(int)
		ret.ExpectResponseCode = in["expect_response_code"].(string)
		ret.SipTcp = in["sip_tcp"].(int)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodRadius(d []interface{}) edpt.HealthMonitorMethodRadius {
	var ret edpt.HealthMonitorMethodRadius
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Radius = in["radius"].(int)
		ret.RadiusUsername = in["radius_username"].(string)
		ret.RadiusPasswordString = in["radius_password_string"].(string)
		//omit radius_encrypted
		ret.RadiusSecret = in["radius_secret"].(string)
		//omit radius_secret_encrypted
		ret.RadiusPort = in["radius_port"].(int)
		ret.RadiusExpect = in["radius_expect"].(int)
		ret.RadiusResponseCode = in["radius_response_code"].(string)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodLdap(d []interface{}) edpt.HealthMonitorMethodLdap {
	var ret edpt.HealthMonitorMethodLdap
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ldap = in["ldap"].(int)
		ret.LdapPort = in["ldap_port"].(int)
		ret.LdapSecurity = in["ldap_security"].(string)
		ret.LdapBinddn = in["ldap_binddn"].(string)
		ret.LdapPassword = in["ldap_password"].(int)
		ret.LdapPasswordString = in["ldap_password_string"].(string)
		//omit ldap_encrypted
		ret.LdapRunSearch = in["ldap_run_search"].(int)
		ret.Basedn = in["basedn"].(string)
		ret.LdapQuery = in["ldap_query"].(string)
		ret.Acceptresref = in["acceptresref"].(int)
		ret.Acceptnotfound = in["acceptnotfound"].(int)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodRtsp(d []interface{}) edpt.HealthMonitorMethodRtsp {
	var ret edpt.HealthMonitorMethodRtsp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Rtsp = in["rtsp"].(int)
		ret.Rtspurl = in["rtspurl"].(string)
		ret.RtspPort = in["rtsp_port"].(int)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodDatabase(d []interface{}) edpt.HealthMonitorMethodDatabase {
	var ret edpt.HealthMonitorMethodDatabase
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Database = in["database"].(int)
		ret.DatabaseName = in["database_name"].(string)
		ret.DbName = in["db_name"].(string)
		ret.DbUsername = in["db_username"].(string)
		ret.DbPassword = in["db_password"].(int)
		ret.DbPasswordStr = in["db_password_str"].(string)
		//omit db_encrypted
		ret.DbSend = in["db_send"].(string)
		ret.DbReceive = in["db_receive"].(string)
		ret.DbRow = in["db_row"].(int)
		ret.DbColumn = in["db_column"].(int)
		ret.DbReceiveInteger = in["db_receive_integer"].(int)
		ret.DbRowInteger = in["db_row_integer"].(int)
		ret.DbColumnInteger = in["db_column_integer"].(int)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodExternal(d []interface{}) edpt.HealthMonitorMethodExternal {
	var ret edpt.HealthMonitorMethodExternal
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.External = in["external"].(int)
		ret.ExtProgram = in["ext_program"].(string)
		ret.SharedPartitionProgram = in["shared_partition_program"].(int)
		ret.ExtProgramShared = in["ext_program_shared"].(string)
		ret.ExtPort = in["ext_port"].(int)
		ret.ExtArguments = in["ext_arguments"].(string)
		ret.ExtPreference = in["ext_preference"].(int)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodNtp(d []interface{}) edpt.HealthMonitorMethodNtp {
	var ret edpt.HealthMonitorMethodNtp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ntp = in["ntp"].(int)
		ret.NtpPort = in["ntp_port"].(int)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodKerberosKdc(d []interface{}) edpt.HealthMonitorMethodKerberosKdc {
	var ret edpt.HealthMonitorMethodKerberosKdc
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.KerberosCfg = getObjectHealthMonitorMethodKerberosKdcKerberosCfg(in["kerberos_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodKerberosKdcKerberosCfg(d []interface{}) edpt.HealthMonitorMethodKerberosKdcKerberosCfg {
	var ret edpt.HealthMonitorMethodKerberosKdcKerberosCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Kinit = in["kinit"].(int)
		ret.KinitPricipalName = in["kinit_pricipal_name"].(string)
		ret.KinitPassword = in["kinit_password"].(string)
		//omit kinit_encrypted
		ret.KinitKdc = in["kinit_kdc"].(string)
		ret.TcpOnly = in["tcp_only"].(int)
		ret.Kadmin = in["kadmin"].(int)
		ret.KadminRealm = in["kadmin_realm"].(string)
		ret.KadminPricipalName = in["kadmin_pricipal_name"].(string)
		ret.KadminPassword = in["kadmin_password"].(string)
		//omit kadmin_encrypted
		ret.KadminServer = in["kadmin_server"].(string)
		ret.KadminKdc = in["kadmin_kdc"].(string)
		ret.Kpasswd = in["kpasswd"].(int)
		ret.KpasswdPricipalName = in["kpasswd_pricipal_name"].(string)
		ret.KpasswdPassword = in["kpasswd_password"].(string)
		//omit kpasswd_encrypted
		ret.KpasswdServer = in["kpasswd_server"].(string)
		ret.KpasswdKdc = in["kpasswd_kdc"].(string)
	}
	return ret
}

func getObjectHealthMonitorMethodHttps(d []interface{}) edpt.HealthMonitorMethodHttps {
	var ret edpt.HealthMonitorMethodHttps
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Https = in["https"].(int)
		ret.WebPort = in["web_port"].(int)
		ret.DisableSslv2hello = in["disable_sslv2hello"].(int)
		ret.HttpsHost = in["https_host"].(string)
		ret.Sni = in["sni"].(int)
		ret.HttpsExpect = in["https_expect"].(int)
		ret.HttpsResponseCode = in["https_response_code"].(string)
		ret.ResponseCodeRegex = in["response_code_regex"].(string)
		ret.HttpsText = in["https_text"].(string)
		ret.TextRegex = in["text_regex"].(string)
		ret.HttpsUrl = in["https_url"].(int)
		ret.UrlType = in["url_type"].(string)
		ret.UrlPath = in["url_path"].(string)
		ret.PostPath = in["post_path"].(string)
		ret.PostType = in["post_type"].(string)
		ret.HttpsPostdata = in["https_postdata"].(string)
		ret.HttpsPostfile = in["https_postfile"].(string)
		ret.HttpsMaintenanceCode = in["https_maintenance_code"].(string)
		ret.Maintenance = in["maintenance"].(int)
		ret.MaintenanceText = in["maintenance_text"].(string)
		ret.MaintenanceTextRegex = in["maintenance_text_regex"].(string)
		ret.HttpsUsername = in["https_username"].(string)
		ret.HttpsServerCertName = in["https_server_cert_name"].(string)
		ret.HttpsPassword = in["https_password"].(int)
		ret.HttpsPasswordString = in["https_password_string"].(string)
		//omit https_encrypted
		ret.HttpsKerberosAuth = in["https_kerberos_auth"].(int)
		ret.HttpsKerberosRealm = in["https_kerberos_realm"].(string)
		ret.HttpsKerberosKdc = getObjectHealthMonitorMethodHttpsHttpsKerberosKdc(in["https_kerberos_kdc"].([]interface{}))
		ret.CertKeyShared = in["cert_key_shared"].(int)
		ret.Cert = in["cert"].(string)
		ret.Key = in["key"].(string)
		ret.KeyPassPhrase = in["key_pass_phrase"].(int)
		ret.KeyPhrase = in["key_phrase"].(string)
		//omit https_key_encrypted
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodHttpsHttpsKerberosKdc(d []interface{}) edpt.HealthMonitorMethodHttpsHttpsKerberosKdc {
	var ret edpt.HealthMonitorMethodHttpsHttpsKerberosKdc
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.HttpsKerberosHostip = in["https_kerberos_hostip"].(string)
		ret.HttpsKerberosHostipv6 = in["https_kerberos_hostipv6"].(string)
		ret.HttpsKerberosPort = in["https_kerberos_port"].(int)
		ret.HttpsKerberosPortv6 = in["https_kerberos_portv6"].(int)
	}
	return ret
}

func getObjectHealthMonitorMethodTacplus(d []interface{}) edpt.HealthMonitorMethodTacplus {
	var ret edpt.HealthMonitorMethodTacplus
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Tacplus = in["tacplus"].(int)
		ret.TacplusUsername = in["tacplus_username"].(string)
		ret.TacplusPassword = in["tacplus_password"].(int)
		ret.TacplusPasswordString = in["tacplus_password_string"].(string)
		//omit tacplus_encrypted
		ret.TacplusSecret = in["tacplus_secret"].(int)
		ret.TacplusSecretString = in["tacplus_secret_string"].(string)
		//omit secret_encrypted
		ret.TacplusPort = in["tacplus_port"].(int)
		ret.TacplusType = in["tacplus_type"].(string)
		//omit uuid
	}
	return ret
}

func getObjectHealthMonitorMethodCompound(d []interface{}) edpt.HealthMonitorMethodCompound {
	var ret edpt.HealthMonitorMethodCompound
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Compound = in["compound"].(int)
		ret.RpnString = in["rpn_string"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointHealthMonitor(d *schema.ResourceData) edpt.HealthMonitor {
	var ret edpt.HealthMonitor
	ret.Inst.DefaultStateUp = d.Get("default_state_up").(int)
	ret.Inst.DisableAfterDown = d.Get("disable_after_down").(int)
	ret.Inst.DsrL2Strict = d.Get("dsr_l2_strict").(int)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.Method = getObjectHealthMonitorMethod(d.Get("method").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OverrideIpv4 = d.Get("override_ipv4").(string)
	ret.Inst.OverrideIpv6 = d.Get("override_ipv6").(string)
	ret.Inst.OverridePort = d.Get("override_port").(int)
	ret.Inst.Passive = d.Get("passive").(int)
	ret.Inst.PassiveInterval = d.Get("passive_interval").(int)
	ret.Inst.Retry = d.Get("retry").(int)
	ret.Inst.SampleThreshold = d.Get("sample_threshold").(int)
	ret.Inst.SslCiphers = d.Get("ssl_ciphers").(string)
	ret.Inst.SslDgversion = d.Get("ssl_dgversion").(int)
	ret.Inst.SslTicket = d.Get("ssl_ticket").(int)
	ret.Inst.SslTicketLifetime = d.Get("ssl_ticket_lifetime").(int)
	ret.Inst.SslVersion = d.Get("ssl_version").(int)
	ret.Inst.StatusCode = d.Get("status_code").(string)
	ret.Inst.StrictRetryOnServerErrResp = d.Get("strict_retry_on_server_err_resp").(int)
	ret.Inst.Threshold = d.Get("threshold").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.UpRetry = d.Get("up_retry").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
