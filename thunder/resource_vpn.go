package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpn() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vpn`: VPN Commands\n\n__PLACEHOLDER__",
		CreateContext: resourceVpnCreate,
		UpdateContext: resourceVpnUpdate,
		ReadContext:   resourceVpnRead,
		DeleteContext: resourceVpnDelete,

		Schema: map[string]*schema.Schema{
			"asymmetric_flow_support": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Support asymmetric flows pass through IPsec tunnel",
			},
			"crl": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"default": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"enable_vpn_metrics": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable exporting vpn statstics to Harmony",
			},
			"error": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"errordump": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"extended_matching": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable session extended matching for packet comes from IPsec tunnel",
			},
			"fragment_after_encap": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Fragment after adding IPsec headers",
			},
			"group_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ike_acc_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IKE Acceleration by Cavium Nitrox card",
			},
			"ike_gateway_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "IKE-gateway name",
						},
						"ike_version": {
							Type: schema.TypeString, Optional: true, Default: "v2", Description: "'v1': IKEv1 key exchange; 'v2': IKEv2 key exchange;",
						},
						"mode": {
							Type: schema.TypeString, Optional: true, Default: "main", Description: "'main': Negotiate Main mode (Default); 'aggressive': Negotiate Aggressive mode;",
						},
						"auth_method": {
							Type: schema.TypeString, Optional: true, Default: "preshare-key", Description: "'preshare-key': Authenticate the remote gateway using a pre-shared key (Default); 'rsa-signature': Authenticate the remote gateway using an RSA certificate; 'ecdsa-signature': Authenticate the remote gateway using an ECDSA certificate; 'eap-radius': Authenticate the remote gateway using an EAP Radius server; 'eap-tls': Authenticate the remote gateway using EAP TLS;",
						},
						"preshare_key_value": {
							Type: schema.TypeString, Optional: true, Description: "pre-shared key",
						},
						"hash": {
							Type: schema.TypeString, Optional: true, Description: "'sha256': Secure Hash Algorithm 256; 'sha384': Secure Hash Algorithm 384; 'sha512': Secure Hash Algorithm 512;",
						},
						"interface_management": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "only handle traffic on management interface, share partition only",
						},
						"key": {
							Type: schema.TypeString, Optional: true, Description: "Private Key",
						},
						"key_passphrase": {
							Type: schema.TypeString, Optional: true, Description: "Private Key Pass Phrase",
						},
						"vrid": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Default VRRP-A vrid",
									},
									"vrid_num": {
										Type: schema.TypeInt, Optional: true, Description: "Specify ha VRRP-A vrid",
									},
								},
							},
						},
						"local_cert": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"local_cert_name": {
										Type: schema.TypeString, Optional: true, Description: "Certificate File Name",
									},
								},
							},
						},
						"remote_ca_cert": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"remote_cert_name": {
										Type: schema.TypeString, Optional: true, Description: "Remote CA certificate DN (C=, ST=, L=, O=, CN=) without emailAddress",
									},
								},
							},
						},
						"local_id": {
							Type: schema.TypeString, Optional: true, Description: "Local Gateway Identity",
						},
						"remote_id": {
							Type: schema.TypeString, Optional: true, Description: "Remote Gateway Identity",
						},
						"enc_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"encryption": {
										Type: schema.TypeString, Optional: true, Description: "'des': Data Encryption Standard algorithm; '3des': Triple Data Encryption Standard algorithm; 'aes-128': Advanced Encryption Standard algorithm CBC Mode(key size: 128 bits); 'aes-192': Advanced Encryption Standard algorithm CBC Mode(key size: 192 bits); 'aes-256': Advanced Encryption Standard algorithm CBC Mode(key size: 256 bits); 'aes-gcm-128': Advanced Encryption Standard algorithm Galois/Counter Mode(key size: 128 bits, ICV size: 16 bytes), only for IKEv2; 'aes-gcm-192': Advanced Encryption Standard algorithm Galois/Counter Mode(key size: 192 bits, ICV size: 16 bytes), only for IKEv2; 'aes-gcm-256': Advanced Encryption Standard algorithm Galois/Counter Mode(key size: 256 bits, ICV size: 16 bytes), only for IKEv2; 'null': No encryption algorithm, only for IKEv2;",
									},
									"hash": {
										Type: schema.TypeString, Optional: true, Description: "'md5': MD5 Dessage-Digest Algorithm; 'sha1': Secure Hash Algorithm 1; 'sha256': Secure Hash Algorithm 256; 'sha384': Secure Hash Algorithm 384; 'sha512': Secure Hash Algorithm 512;",
									},
									"prf": {
										Type: schema.TypeString, Optional: true, Description: "'md5': MD5 Dessage-Digest Algorithm; 'sha1': Secure Hash Algorithm 1; 'sha256': Secure Hash Algorithm 256; 'sha384': Secure Hash Algorithm 384; 'sha512': Secure Hash Algorithm 512;",
									},
									"priority": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Prioritizes (1-10) security protocol, least value has highest priority",
									},
									"gcm_priority": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Prioritizes (1-10) security protocol, least value has highest priority",
									},
								},
							},
						},
						"dh_group": {
							Type: schema.TypeString, Optional: true, Default: "1", Description: "'1': Diffie-Hellman group 1 - 768-bit(Default); '2': Diffie-Hellman group 2 - 1024-bit; '5': Diffie-Hellman group 5 - 1536-bit; '14': Diffie-Hellman group 14 - 2048-bit; '15': Diffie-Hellman group 15 - 3072-bit; '16': Diffie-Hellman group 16 - 4096-bit; '18': Diffie-Hellman group 18 - 8192-bit; '19': Diffie-Hellman group 19 - 256-bit Elliptic Curve; '20': Diffie-Hellman group 20 - 384-bit Elliptic Curve;",
						},
						"local_address": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"local_ip": {
										Type: schema.TypeString, Optional: true, Description: "Ipv4 address",
									},
									"local_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "Ipv6 address",
									},
								},
							},
						},
						"remote_address": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"remote_ip": {
										Type: schema.TypeString, Optional: true, Description: "Ipv4 address",
									},
									"dns": {
										Type: schema.TypeString, Optional: true, Description: "Remote IP based on Domain name",
									},
									"remote_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "Ipv6 address",
									},
								},
							},
						},
						"lifetime": {
							Type: schema.TypeInt, Optional: true, Default: 86400, Description: "IKE SA age in seconds",
						},
						"fragment_size": {
							Type: schema.TypeInt, Optional: true, Description: "Enable IKE message fragment and set fragment size",
						},
						"nat_traversal": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"dpd": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interval": {
										Type: schema.TypeInt, Optional: true, Description: "Interval time in seconds",
									},
									"retry": {
										Type: schema.TypeInt, Optional: true, Description: "Retry times",
									},
								},
							},
						},
						"disable_rekey": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable initiating rekey",
						},
						"configuration_payload": {
							Type: schema.TypeString, Optional: true, Description: "'dhcp': Enable DHCP configuration-payload; 'radius': Enable RADIUS configuration-payload;",
						},
						"dhcp_server": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pri": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dhcp_pri_ipv4": {
													Type: schema.TypeString, Optional: true, Description: "Primary DHCP Server IP Address",
												},
											},
										},
									},
									"sec": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dhcp_sec_ipv4": {
													Type: schema.TypeString, Optional: true, Description: "Secondary DHCP Server IP Address",
												},
											},
										},
									},
								},
							},
						},
						"radius_server": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"radius_pri": {
										Type: schema.TypeString, Optional: true, Description: "Primary RADIUS Authentication Server",
									},
									"radius_sec": {
										Type: schema.TypeString, Optional: true, Description: "Secondary RADIUS Authentication Server",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'v2-init-rekey': Initiate Rekey; 'v2-rsp-rekey': Respond Rekey; 'v2-child-sa-rekey': Child SA Rekey; 'v2-in-invalid': Incoming Invalid; 'v2-in-invalid-spi': Incoming Invalid SPI; 'v2-in-init-req': Incoming Init Request; 'v2-in-init-rsp': Incoming Init Response; 'v2-out-init-req': Outgoing Init Request; 'v2-out-init-rsp': Outgoing Init Response; 'v2-in-auth-req': Incoming Auth Request; 'v2-in-auth-rsp': Incoming Auth Response; 'v2-out-auth-req': Outgoing Auth Request; 'v2-out-auth-rsp': Outgoing Auth Response; 'v2-in-create-child-req': Incoming Create Child Request; 'v2-in-create-child-rsp': Incoming Create Child Response; 'v2-out-create-child-req': Outgoing Create Child Request; 'v2-out-create-child-rsp': Outgoing Create Child Response; 'v2-in-info-req': Incoming Info Request; 'v2-in-info-rsp': Incoming Info Response; 'v2-out-info-req': Outgoing Info Request; 'v2-out-info-rsp': Outgoing Info Response; 'v1-in-id-prot-req': Incoming ID Protection Request; 'v1-in-id-prot-rsp': Incoming ID Protection Response; 'v1-out-id-prot-req': Outgoing ID Protection Request; 'v1-out-id-prot-rsp': Outgoing ID Protection Response; 'v1-in-auth-only-req': Incoming Auth Only Request; 'v1-in-auth-only-rsp': Incoming Auth Only Response; 'v1-out-auth-only-req': Outgoing Auth Only Request; 'v1-out-auth-only-rsp': Outgoing Auth Only Response; 'v1-in-aggressive-req': Incoming Aggressive Request; 'v1-in-aggressive-rsp': Incoming Aggressive Response; 'v1-out-aggressive-req': Outgoing Aggressive Request; 'v1-out-aggressive-rsp': Outgoing Aggressive Response; 'v1-in-info-v1-req': Incoming Info Request; 'v1-in-info-v1-rsp': Incoming Info Response; 'v1-out-info-v1-req': Outgoing Info Request; 'v1-out-info-v1-rsp': Outgoing Info Response; 'v1-in-transaction-req': Incoming Transaction Request; 'v1-in-transaction-rsp': Incoming Transaction Response; 'v1-out-transaction-req': Outgoing Transaction Request; 'v1-out-transaction-rsp': Outgoing Transaction Response; 'v1-in-quick-mode-req': Incoming Quick Mode Request; 'v1-in-quick-mode-rsp': Incoming Quick Mode Response; 'v1-out-quick-mode-req': Outgoing Quick Mode Request; 'v1-out-quick-mode-rsp': Outgoing Quick Mode Response; 'v1-in-new-group-mode-req': Incoming New Group Mode Request; 'v1-in-new-group-mode-rsp': Incoming New Group Mode Response; 'v1-out-new-group-mode-req': Outgoing New Group Mode Request; 'v1-out-new-group-mode-rsp': Outgoing New Group Mode Response; 'v1-child-sa-invalid-spi': Invalid SPI for Child SAs; 'v2-child-sa-invalid-spi': Invalid SPI for Child SAs; 'ike-current-version': IKE version;",
									},
								},
							},
						},
					},
				},
			},
			"ike_logging_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IKE negotiation logging",
			},
			"ike_sa": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ike_sa_brief": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ike_sa_clients": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ike_sa_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 600, Description: "Timeout IKE-SA in connecting state in seconds (default 600s)",
			},
			"ike_stats_by_gw": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ike_stats_global": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'v2-init-rekey': Initiate Rekey; 'v2-rsp-rekey': Respond Rekey; 'v2-child-sa-rekey': Child SA Rekey; 'v2-in-invalid': Incoming Invalid; 'v2-in-invalid-spi': Incoming Invalid SPI; 'v2-in-init-req': Incoming Init Request; 'v2-in-init-rsp': Incoming Init Response; 'v2-out-init-req': Outgoing Init Request; 'v2-out-init-rsp': Outgoing Init Response; 'v2-in-auth-req': Incoming Auth Request; 'v2-in-auth-rsp': Incoming Auth Response; 'v2-out-auth-req': Outgoing Auth Request; 'v2-out-auth-rsp': Outgoing Auth Response; 'v2-in-create-child-req': Incoming Create Child Request; 'v2-in-create-child-rsp': Incoming Create Child Response; 'v2-out-create-child-req': Outgoing Create Child Request; 'v2-out-create-child-rsp': Outgoing Create Child Response; 'v2-in-info-req': Incoming Info Request; 'v2-in-info-rsp': Incoming Info Response; 'v2-out-info-req': Outgoing Info Request; 'v2-out-info-rsp': Outgoing Info Response; 'v1-in-id-prot-req': Incoming ID Protection Request; 'v1-in-id-prot-rsp': Incoming ID Protection Response; 'v1-out-id-prot-req': Outgoing ID Protection Request; 'v1-out-id-prot-rsp': Outgoing ID Protection Response; 'v1-in-auth-only-req': Incoming Auth Only Request; 'v1-in-auth-only-rsp': Incoming Auth Only Response; 'v1-out-auth-only-req': Outgoing Auth Only Request; 'v1-out-auth-only-rsp': Outgoing Auth Only Response; 'v1-in-aggressive-req': Incoming Aggressive Request; 'v1-in-aggressive-rsp': Incoming Aggressive Response; 'v1-out-aggressive-req': Outgoing Aggressive Request; 'v1-out-aggressive-rsp': Outgoing Aggressive Response; 'v1-in-info-v1-req': Incoming Info Request; 'v1-in-info-v1-rsp': Incoming Info Response; 'v1-out-info-v1-req': Outgoing Info Request; 'v1-out-info-v1-rsp': Outgoing Info Response; 'v1-in-transaction-req': Incoming Transaction Request; 'v1-in-transaction-rsp': Incoming Transaction Response; 'v1-out-transaction-req': Outgoing Transaction Request; 'v1-out-transaction-rsp': Outgoing Transaction Response; 'v1-in-quick-mode-req': Incoming Quick Mode Request; 'v1-in-quick-mode-rsp': Incoming Quick Mode Response; 'v1-out-quick-mode-req': Outgoing Quick Mode Request; 'v1-out-quick-mode-rsp': Outgoing Quick Mode Response; 'v1-in-new-group-mode-req': Incoming New Group Mode Request; 'v1-in-new-group-mode-rsp': Incoming New Group Mode Response; 'v1-out-new-group-mode-req': Outgoing New Group Mode Request; 'v1-out-new-group-mode-rsp': Outgoing New Group Mode Response;",
									},
								},
							},
						},
					},
				},
			},
			"ipsec_cipher_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable cipher check, IPsec SA cipher must weaker than IKE gateway cipher, and DES/3DES/MD5/null will not work.",
			},
			"ipsec_error_dump": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Support record the error ipsec cavium information in dump file",
			},
			"ipsec_group_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Group name",
						},
						"ipsecgroup_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipsec": {
										Type: schema.TypeString, Optional: true, Description: "specify a name to group active/backup tunnels",
									},
									"priority": {
										Type: schema.TypeInt, Optional: true, Description: "Highest priority value is the active tunnel",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"ipsec_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "IPsec name",
						},
						"mode": {
							Type: schema.TypeString, Optional: true, Default: "tunnel", Description: "'tunnel': Encapsulating the packet in IPsec tunnel mode (Default);",
						},
						"dscp": {
							Type: schema.TypeString, Optional: true, Description: "'default': Default dscp (000000); 'af11': AF11 (001010); 'af12': AF12 (001100); 'af13': AF13 (001110); 'af21': AF21 (010010); 'af22': AF22 (010100); 'af23': AF23 (010110); 'af31': AF31 (011010); 'af32': AF32 (011100); 'af33': AF33 (011110); 'af41': AF41 (100010); 'af42': AF42 (100100); 'af43': AF43 (100110); 'cs1': CS1 (001000); 'cs2': CS2 (010000); 'cs3': CS3 (011000); 'cs4': CS4 (100000); 'cs5': CS5 (101000); 'cs6': CS6 (110000); 'cs7': CS7 (111000); 'ef': EF (101110); '0': 000000; '1': 000001; '2': 000010; '3': 000011; '4': 000100; '5': 000101; '6': 000110; '7': 000111; '8': 001000; '9': 001001; '10': 001010; '11': 001011; '12': 001100; '13': 001101; '14': 001110; '15': 001111; '16': 010000; '17': 010001; '18': 010010; '19': 010011; '20': 010100; '21': 010101; '22': 010110; '23': 010111; '24': 011000; '25': 011001; '26': 011010; '27': 011011; '28': 011100; '29': 011101; '30': 011110; '31': 011111; '32': 100000; '33': 100001; '34': 100010; '35': 100011; '36': 100100; '37': 100101; '38': 100110; '39': 100111; '40': 101000; '41': 101001; '42': 101010; '43': 101011; '44': 101100; '45': 101101; '46': 101110; '47': 101111; '48': 110000; '49': 110001; '50': 110010; '51': 110011; '52': 110100; '53': 110101; '54': 110110; '55': 110111; '56': 111000; '57': 111001; '58': 111010; '59': 111011; '60': 111100; '61': 111101; '62': 111110; '63': 111111;",
						},
						"proto": {
							Type: schema.TypeString, Optional: true, Default: "esp", Description: "'esp': Encapsulating security protocol (Default);",
						},
						"dh_group": {
							Type: schema.TypeString, Optional: true, Default: "0", Description: "'0': Diffie-Hellman group 0 (Default); '1': Diffie-Hellman group 1 - 768-bits; '2': Diffie-Hellman group 2 - 1024-bits; '5': Diffie-Hellman group 5 - 1536-bits; '14': Diffie-Hellman group 14 - 2048-bits; '15': Diffie-Hellman group 15 - 3072-bits; '16': Diffie-Hellman group 16 - 4096-bits; '18': Diffie-Hellman group 18 - 8192-bits; '19': Diffie-Hellman group 19 - 256-bit Elliptic Curve; '20': Diffie-Hellman group 20 - 384-bit Elliptic Curve;",
						},
						"enc_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"encryption": {
										Type: schema.TypeString, Optional: true, Description: "'des': Data Encryption Standard algorithm; '3des': Triple Data Encryption Standard algorithm; 'aes-128': Advanced Encryption Standard algorithm CBC Mode(key size: 128 bits); 'aes-192': Advanced Encryption Standard algorithm CBC Mode(key size: 192 bits); 'aes-256': Advanced Encryption Standard algorithm CBC Mode(key size: 256 bits); 'aes-gcm-128': Advanced Encryption Standard algorithm Galois/Counter Mode(key size: 128 bits, ICV size: 16 bytes); 'aes-gcm-192': Advanced Encryption Standard algorithm Galois/Counter Mode(key size: 192 bits, ICV size: 16 bytes); 'aes-gcm-256': Advanced Encryption Standard algorithm Galois/Counter Mode(key size: 256 bits, ICV size: 16 bytes); 'null': No encryption algorithm;",
									},
									"hash": {
										Type: schema.TypeString, Optional: true, Description: "'md5': MD5 Dessage-Digest Algorithm; 'sha1': Secure Hash Algorithm 1; 'sha256': Secure Hash Algorithm 256; 'sha384': Secure Hash Algorithm 384; 'sha512': Secure Hash Algorithm 512; 'null': No hash algorithm;",
									},
									"priority": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Prioritizes (1-10) security protocol, least value has highest priority",
									},
									"gcm_priority": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Prioritizes (1-10) security protocol, least value has highest priority",
									},
								},
							},
						},
						"lifetime": {
							Type: schema.TypeInt, Optional: true, Default: 28800, Description: "IPsec SA age in seconds",
						},
						"lifebytes": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPsec SA age in megabytes (0 indicates unlimited bytes)",
						},
						"anti_replay_window": {
							Type: schema.TypeString, Optional: true, Default: "0", Description: "'0': Disable Anti-Replay Window Check; '32': Window size of 32; '64': Window size of 64; '128': Window size of 128; '256': Window size of 256; '512': Window size of 512; '1024': Window size of 1024; '2048': Window size of 2048; '3072': Window size of 3072; '4096': Window size of 4096; '8192': Window size of 8192;",
						},
						"up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Initiates SA negotiation to bring the IPsec connection up",
						},
						"sequence_number_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not use incremental sequence number in the ESP header",
						},
						"traffic_selector": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"local": {
													Type: schema.TypeString, Optional: true, Description: "Local Traffic Selector",
												},
												"local_netmask": {
													Type: schema.TypeString, Optional: true, Description: "IPv4 Address Network Mask",
												},
												"local_port": {
													Type: schema.TypeInt, Optional: true, Description: "Port Number",
												},
												"remote_ipv4_assigned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remote IP address assigned",
												},
												"remote_ip": {
													Type: schema.TypeString, Optional: true, Description: "IPv4 Address",
												},
												"remote_netmask": {
													Type: schema.TypeString, Optional: true, Description: "IPv4 Address Network Mask",
												},
												"remote_port": {
													Type: schema.TypeInt, Optional: true, Description: "Port Number",
												},
												"protocol": {
													Type: schema.TypeInt, Optional: true, Description: "IP Protocol Number (0-255)",
												},
											},
										},
									},
									"ipv6": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"localv6": {
													Type: schema.TypeString, Optional: true, Description: "Local Traffic Selector",
												},
												"local_portv6": {
													Type: schema.TypeInt, Optional: true, Description: "Port Number",
												},
												"remote_ipv6_assigned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remote IPv6 address assigned",
												},
												"remote_ipv6": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 Address",
												},
												"remote_portv6": {
													Type: schema.TypeInt, Optional: true, Description: "Port Number",
												},
												"protocolv6": {
													Type: schema.TypeInt, Optional: true, Description: "IP Protocol Number (0-255)",
												},
											},
										},
									},
								},
							},
						},
						"enforce_traffic_selector": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enforce Traffic Selector",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'packets-encrypted': Encrypted Packets; 'packets-decrypted': Decrypted Packets; 'anti-replay-num': Anti-Replay Failure; 'rekey-num': Rekey Times; 'packets-err-inactive': Inactive Error; 'packets-err-encryption': Encryption Error; 'packets-err-pad-check': Pad Check Error; 'packets-err-pkt-sanity': Packets Sanity Error; 'packets-err-icv-check': ICV Check Error; 'packets-err-lifetime-lifebytes': Lifetime Lifebytes Error; 'bytes-encrypted': Encrypted Bytes; 'bytes-decrypted': Decrypted Bytes; 'prefrag-success': Pre-frag Success; 'prefrag-error': Pre-frag Error; 'cavium-bytes-encrypted': CAVIUM Encrypted Bytes; 'cavium-bytes-decrypted': CAVIUM Decrypted Bytes; 'cavium-packets-encrypted': CAVIUM Encrypted Packets; 'cavium-packets-decrypted': CAVIUM Decrypted Packets; 'qat-bytes-encrypted': QAT Encrypted Bytes; 'qat-bytes-decrypted': QAT Decrypted Bytes; 'qat-packets-encrypted': QAT Encrypted Packets; 'qat-packets-decrypted': QAT Decrypted Packets; 'tunnel-intf-down': Packet dropped: Tunnel Interface Down; 'pkt-fail-prep-to-send': Packet dropped: Failed in prepare to send; 'no-next-hop': Packet dropped: No next hop; 'invalid-tunnel-id': Packet dropped: Invalid tunnel ID; 'no-tunnel-found': Packet dropped: No tunnel found; 'pkt-fail-to-send': Packet dropped: Failed to send; 'frag-after-encap-frag-packets': Frag-after-encap Fragment Generated; 'frag-received': Fragment Received; 'sequence-num': Sequence Number; 'sequence-num-rollover': Sequence Number Rollover; 'packets-err-nh-check': Next Header Check Error;",
									},
								},
							},
						},
						"bind_tunnel": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface index",
									},
									"next_hop": {
										Type: schema.TypeString, Optional: true, Description: "IPsec Next Hop IP Address",
									},
									"next_hop_v6": {
										Type: schema.TypeString, Optional: true, Description: "IPsec Next Hop IPv6 Address",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ipsec_gateway": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ike_gateway": {
										Type: schema.TypeString, Optional: true, Description: "Gateway to use for IPsec SA",
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
			"ipsec_mgmt_default_policy_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop MGMT traffic that is not match ipsec tunnel, share partition only",
			},
			"ipsec_sa": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ipsec_sa_by_gw": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ipsec_sa_clients": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ipsec_sa_stats_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'packets-encrypted': Encrypted Packets; 'packets-decrypted': Decrypted Packets; 'anti-replay-num': Anti-Replay Failure; 'rekey-num': Rekey Times; 'packets-err-inactive': Inactive Error; 'packets-err-encryption': Encryption Error; 'packets-err-pad-check': Pad Check Error; 'packets-err-pkt-sanity': Packets Sanity Error; 'packets-err-icv-check': ICV Check Error; 'packets-err-lifetime-lifebytes': Lifetime Lifebytes Error; 'bytes-encrypted': Encrypted Bytes; 'bytes-decrypted': Decrypted Bytes; 'prefrag-success': Pre-frag Success; 'prefrag-error': Pre-frag Error; 'cavium-bytes-encrypted': CAVIUM Encrypted Bytes; 'cavium-bytes-decrypted': CAVIUM Decrypted Bytes; 'cavium-packets-encrypted': CAVIUM Encrypted Packets; 'cavium-packets-decrypted': CAVIUM Decrypted Packets; 'qat-bytes-encrypted': QAT Encrypted Bytes; 'qat-bytes-decrypted': QAT Decrypted Bytes; 'qat-packets-encrypted': QAT Encrypted Packets; 'qat-packets-decrypted': QAT Decrypted Packets; 'tunnel-intf-down': Packet dropped: Tunnel Interface Down; 'pkt-fail-prep-to-send': Packet dropped: Failed in prepare to send; 'no-next-hop': Packet dropped: No next hop; 'invalid-tunnel-id': Packet dropped: Invalid tunnel ID; 'no-tunnel-found': Packet dropped: No tunnel found; 'pkt-fail-to-send': Packet dropped: Failed to send; 'frag-after-encap-frag-packets': Frag-after-encap Fragment Generated; 'frag-received': Fragment Received; 'sequence-num': Sequence Number; 'sequence-num-rollover': Sequence Number Rollover; 'packets-err-nh-check': Next Header Check Error;",
									},
								},
							},
						},
					},
				},
			},
			"jumbo_fragment": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Support IKE jumbo fragment packet",
			},
			"log": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"nat_traversal_flow_affinity": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Choose IPsec UDP source port based on port of inner flow (only for A10 to A10)",
			},
			"ocsp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"revocation_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Revocation name",
						},
						"ca": {
							Type: schema.TypeString, Optional: true, Description: "Certificate Authority file name",
						},
						"crl": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crl_pri": {
										Type: schema.TypeString, Optional: true, Description: "Primary CRL URL (http://www.example.com/ocsp) (only .der filetypes)",
									},
									"crl_sec": {
										Type: schema.TypeString, Optional: true, Description: "Secondary CRL URL (http://www.example.com/ocsp) (only .der filetypes)",
									},
								},
							},
						},
						"ocsp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ocsp_pri": {
										Type: schema.TypeString, Optional: true, Description: "Primary OCSP Authentication Server",
									},
									"ocsp_sec": {
										Type: schema.TypeString, Optional: true, Description: "Secondary OCSP Authentication Server",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'passthrough': passthrough; 'ha-standby-drop': ha-standby-drop;",
						},
					},
				},
			},
			"signature_authentication": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable use of different hash algorithms for signature authentication in IKEv2",
			},
			"stateful_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "VPN module will work in stateful mode and create sessions",
			},
			"tcp_mss_adjust_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable TCP MSS adjustment in SYN packet",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVpnCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpn(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnRead(ctx, d, meta)
	}
	return diags
}

func resourceVpnUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpn(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnRead(ctx, d, meta)
	}
	return diags
}
func resourceVpnDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpn(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVpnRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpn(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVpnCrl3614(d []interface{}) edpt.VpnCrl3614 {

	var ret edpt.VpnCrl3614
	return ret
}

func getObjectVpnDefault3615(d []interface{}) edpt.VpnDefault3615 {

	var ret edpt.VpnDefault3615
	return ret
}

func getObjectVpnError3616(d []interface{}) edpt.VpnError3616 {

	var ret edpt.VpnError3616
	return ret
}

func getObjectVpnErrordump3617(d []interface{}) edpt.VpnErrordump3617 {

	var ret edpt.VpnErrordump3617
	return ret
}

func getObjectVpnGroupList3618(d []interface{}) edpt.VpnGroupList3618 {

	var ret edpt.VpnGroupList3618
	return ret
}

func getSliceVpnIkeGatewayList(d []interface{}) []edpt.VpnIkeGatewayList {

	count1 := len(d)
	ret := make([]edpt.VpnIkeGatewayList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIkeGatewayList
		oi.Name = in["name"].(string)
		oi.IkeVersion = in["ike_version"].(string)
		oi.Mode = in["mode"].(string)
		oi.AuthMethod = in["auth_method"].(string)
		oi.PreshareKeyValue = in["preshare_key_value"].(string)
		//omit preshare_key_encrypted
		oi.Hash = in["hash"].(string)
		oi.InterfaceManagement = in["interface_management"].(int)
		oi.Key = in["key"].(string)
		oi.KeyPassphrase = in["key_passphrase"].(string)
		//omit key_passphrase_encrypted
		oi.Vrid = getObjectVpnIkeGatewayListVrid(in["vrid"].([]interface{}))
		oi.LocalCert = getObjectVpnIkeGatewayListLocalCert(in["local_cert"].([]interface{}))
		oi.RemoteCaCert = getObjectVpnIkeGatewayListRemoteCaCert(in["remote_ca_cert"].([]interface{}))
		oi.LocalId = in["local_id"].(string)
		oi.RemoteId = in["remote_id"].(string)
		oi.EncCfg = getSliceVpnIkeGatewayListEncCfg(in["enc_cfg"].([]interface{}))
		oi.DhGroup = in["dh_group"].(string)
		oi.LocalAddress = getObjectVpnIkeGatewayListLocalAddress(in["local_address"].([]interface{}))
		oi.RemoteAddress = getObjectVpnIkeGatewayListRemoteAddress(in["remote_address"].([]interface{}))
		oi.Lifetime = in["lifetime"].(int)
		oi.FragmentSize = in["fragment_size"].(int)
		oi.NatTraversal = in["nat_traversal"].(int)
		oi.Dpd = getObjectVpnIkeGatewayListDpd(in["dpd"].([]interface{}))
		oi.DisableRekey = in["disable_rekey"].(int)
		oi.ConfigurationPayload = in["configuration_payload"].(string)
		oi.DhcpServer = getObjectVpnIkeGatewayListDhcpServer(in["dhcp_server"].([]interface{}))
		oi.RadiusServer = getObjectVpnIkeGatewayListRadiusServer(in["radius_server"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceVpnIkeGatewayListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnIkeGatewayListVrid(d []interface{}) edpt.VpnIkeGatewayListVrid {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayListVrid
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Default = in["default"].(int)
		ret.VridNum = in["vrid_num"].(int)
	}
	return ret
}

func getObjectVpnIkeGatewayListLocalCert(d []interface{}) edpt.VpnIkeGatewayListLocalCert {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayListLocalCert
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LocalCertName = in["local_cert_name"].(string)
	}
	return ret
}

func getObjectVpnIkeGatewayListRemoteCaCert(d []interface{}) edpt.VpnIkeGatewayListRemoteCaCert {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayListRemoteCaCert
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RemoteCertName = in["remote_cert_name"].(string)
	}
	return ret
}

func getSliceVpnIkeGatewayListEncCfg(d []interface{}) []edpt.VpnIkeGatewayListEncCfg {

	count1 := len(d)
	ret := make([]edpt.VpnIkeGatewayListEncCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIkeGatewayListEncCfg
		oi.Encryption = in["encryption"].(string)
		oi.Hash = in["hash"].(string)
		oi.Prf = in["prf"].(string)
		oi.Priority = in["priority"].(int)
		oi.Gcm_priority = in["gcm_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnIkeGatewayListLocalAddress(d []interface{}) edpt.VpnIkeGatewayListLocalAddress {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayListLocalAddress
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LocalIp = in["local_ip"].(string)
		ret.LocalIpv6 = in["local_ipv6"].(string)
	}
	return ret
}

func getObjectVpnIkeGatewayListRemoteAddress(d []interface{}) edpt.VpnIkeGatewayListRemoteAddress {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayListRemoteAddress
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RemoteIp = in["remote_ip"].(string)
		ret.Dns = in["dns"].(string)
		ret.RemoteIpv6 = in["remote_ipv6"].(string)
	}
	return ret
}

func getObjectVpnIkeGatewayListDpd(d []interface{}) edpt.VpnIkeGatewayListDpd {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayListDpd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.Retry = in["retry"].(int)
	}
	return ret
}

func getObjectVpnIkeGatewayListDhcpServer(d []interface{}) edpt.VpnIkeGatewayListDhcpServer {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayListDhcpServer
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Pri = getObjectVpnIkeGatewayListDhcpServerPri(in["pri"].([]interface{}))
		ret.Sec = getObjectVpnIkeGatewayListDhcpServerSec(in["sec"].([]interface{}))
	}
	return ret
}

func getObjectVpnIkeGatewayListDhcpServerPri(d []interface{}) edpt.VpnIkeGatewayListDhcpServerPri {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayListDhcpServerPri
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DhcpPriIpv4 = in["dhcp_pri_ipv4"].(string)
	}
	return ret
}

func getObjectVpnIkeGatewayListDhcpServerSec(d []interface{}) edpt.VpnIkeGatewayListDhcpServerSec {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayListDhcpServerSec
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DhcpSecIpv4 = in["dhcp_sec_ipv4"].(string)
	}
	return ret
}

func getObjectVpnIkeGatewayListRadiusServer(d []interface{}) edpt.VpnIkeGatewayListRadiusServer {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayListRadiusServer
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RadiusPri = in["radius_pri"].(string)
		ret.RadiusSec = in["radius_sec"].(string)
	}
	return ret
}

func getSliceVpnIkeGatewayListSamplingEnable(d []interface{}) []edpt.VpnIkeGatewayListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VpnIkeGatewayListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIkeGatewayListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnIkeSa3619(d []interface{}) edpt.VpnIkeSa3619 {

	var ret edpt.VpnIkeSa3619
	return ret
}

func getObjectVpnIkeSaBrief3620(d []interface{}) edpt.VpnIkeSaBrief3620 {

	var ret edpt.VpnIkeSaBrief3620
	return ret
}

func getObjectVpnIkeSaClients3621(d []interface{}) edpt.VpnIkeSaClients3621 {

	var ret edpt.VpnIkeSaClients3621
	return ret
}

func getObjectVpnIkeStatsByGw3622(d []interface{}) edpt.VpnIkeStatsByGw3622 {

	var ret edpt.VpnIkeStatsByGw3622
	return ret
}

func getObjectVpnIkeStatsGlobal3623(d []interface{}) edpt.VpnIkeStatsGlobal3623 {

	count1 := len(d)
	var ret edpt.VpnIkeStatsGlobal3623
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceVpnIkeStatsGlobalSamplingEnable3624(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceVpnIkeStatsGlobalSamplingEnable3624(d []interface{}) []edpt.VpnIkeStatsGlobalSamplingEnable3624 {

	count1 := len(d)
	ret := make([]edpt.VpnIkeStatsGlobalSamplingEnable3624, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIkeStatsGlobalSamplingEnable3624
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVpnIpsecGroupList(d []interface{}) []edpt.VpnIpsecGroupList {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecGroupList
		oi.Name = in["name"].(string)
		oi.IpsecgroupCfg = getSliceVpnIpsecGroupListIpsecgroupCfg(in["ipsecgroup_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVpnIpsecGroupListIpsecgroupCfg(d []interface{}) []edpt.VpnIpsecGroupListIpsecgroupCfg {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecGroupListIpsecgroupCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecGroupListIpsecgroupCfg
		oi.Ipsec = in["ipsec"].(string)
		oi.Priority = in["priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVpnIpsecList(d []interface{}) []edpt.VpnIpsecList {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecList
		oi.Name = in["name"].(string)
		oi.Mode = in["mode"].(string)
		oi.Dscp = in["dscp"].(string)
		oi.Proto = in["proto"].(string)
		oi.DhGroup = in["dh_group"].(string)
		oi.EncCfg = getSliceVpnIpsecListEncCfg(in["enc_cfg"].([]interface{}))
		oi.Lifetime = in["lifetime"].(int)
		oi.Lifebytes = in["lifebytes"].(int)
		oi.AntiReplayWindow = in["anti_replay_window"].(string)
		oi.Up = in["up"].(int)
		oi.SequenceNumberDisable = in["sequence_number_disable"].(int)
		oi.TrafficSelector = getObjectVpnIpsecListTrafficSelector(in["traffic_selector"].([]interface{}))
		oi.EnforceTrafficSelector = in["enforce_traffic_selector"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceVpnIpsecListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.BindTunnel = getObjectVpnIpsecListBindTunnel(in["bind_tunnel"].([]interface{}))
		oi.IpsecGateway = getObjectVpnIpsecListIpsecGateway(in["ipsec_gateway"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVpnIpsecListEncCfg(d []interface{}) []edpt.VpnIpsecListEncCfg {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecListEncCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecListEncCfg
		oi.Encryption = in["encryption"].(string)
		oi.Hash = in["hash"].(string)
		oi.Priority = in["priority"].(int)
		oi.Gcm_priority = in["gcm_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnIpsecListTrafficSelector(d []interface{}) edpt.VpnIpsecListTrafficSelector {

	count1 := len(d)
	var ret edpt.VpnIpsecListTrafficSelector
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4 = getObjectVpnIpsecListTrafficSelectorIpv4(in["ipv4"].([]interface{}))
		ret.Ipv6 = getObjectVpnIpsecListTrafficSelectorIpv6(in["ipv6"].([]interface{}))
	}
	return ret
}

func getObjectVpnIpsecListTrafficSelectorIpv4(d []interface{}) edpt.VpnIpsecListTrafficSelectorIpv4 {

	count1 := len(d)
	var ret edpt.VpnIpsecListTrafficSelectorIpv4
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Local = in["local"].(string)
		ret.Local_netmask = in["local_netmask"].(string)
		ret.Local_port = in["local_port"].(int)
		ret.RemoteIpv4Assigned = in["remote_ipv4_assigned"].(int)
		ret.RemoteIp = in["remote_ip"].(string)
		ret.Remote_netmask = in["remote_netmask"].(string)
		ret.Remote_port = in["remote_port"].(int)
		ret.Protocol = in["protocol"].(int)
	}
	return ret
}

func getObjectVpnIpsecListTrafficSelectorIpv6(d []interface{}) edpt.VpnIpsecListTrafficSelectorIpv6 {

	count1 := len(d)
	var ret edpt.VpnIpsecListTrafficSelectorIpv6
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Localv6 = in["localv6"].(string)
		ret.Local_portv6 = in["local_portv6"].(int)
		ret.RemoteIpv6Assigned = in["remote_ipv6_assigned"].(int)
		ret.RemoteIpv6 = in["remote_ipv6"].(string)
		ret.Remote_portv6 = in["remote_portv6"].(int)
		ret.Protocolv6 = in["protocolv6"].(int)
	}
	return ret
}

func getSliceVpnIpsecListSamplingEnable(d []interface{}) []edpt.VpnIpsecListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnIpsecListBindTunnel(d []interface{}) edpt.VpnIpsecListBindTunnel {

	count1 := len(d)
	var ret edpt.VpnIpsecListBindTunnel
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tunnel = in["tunnel"].(int)
		ret.NextHop = in["next_hop"].(string)
		ret.NextHopV6 = in["next_hop_v6"].(string)
		//omit uuid
	}
	return ret
}

func getObjectVpnIpsecListIpsecGateway(d []interface{}) edpt.VpnIpsecListIpsecGateway {

	count1 := len(d)
	var ret edpt.VpnIpsecListIpsecGateway
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IkeGateway = in["ike_gateway"].(string)
		//omit uuid
	}
	return ret
}

func getObjectVpnIpsecSa3625(d []interface{}) edpt.VpnIpsecSa3625 {

	var ret edpt.VpnIpsecSa3625
	return ret
}

func getObjectVpnIpsecSaByGw3626(d []interface{}) edpt.VpnIpsecSaByGw3626 {

	var ret edpt.VpnIpsecSaByGw3626
	return ret
}

func getObjectVpnIpsecSaClients3627(d []interface{}) edpt.VpnIpsecSaClients3627 {

	var ret edpt.VpnIpsecSaClients3627
	return ret
}

func getSliceVpnIpsecSaStatsList(d []interface{}) []edpt.VpnIpsecSaStatsList {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecSaStatsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecSaStatsList
		oi.SamplingEnable = getSliceVpnIpsecSaStatsListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVpnIpsecSaStatsListSamplingEnable(d []interface{}) []edpt.VpnIpsecSaStatsListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecSaStatsListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecSaStatsListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnLog3628(d []interface{}) edpt.VpnLog3628 {

	var ret edpt.VpnLog3628
	return ret
}

func getObjectVpnOcsp3629(d []interface{}) edpt.VpnOcsp3629 {

	var ret edpt.VpnOcsp3629
	return ret
}

func getSliceVpnRevocationList(d []interface{}) []edpt.VpnRevocationList {

	count1 := len(d)
	ret := make([]edpt.VpnRevocationList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnRevocationList
		oi.Name = in["name"].(string)
		oi.Ca = in["ca"].(string)
		oi.Crl = getObjectVpnRevocationListCrl(in["crl"].([]interface{}))
		oi.Ocsp = getObjectVpnRevocationListOcsp(in["ocsp"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnRevocationListCrl(d []interface{}) edpt.VpnRevocationListCrl {

	count1 := len(d)
	var ret edpt.VpnRevocationListCrl
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CrlPri = in["crl_pri"].(string)
		ret.CrlSec = in["crl_sec"].(string)
	}
	return ret
}

func getObjectVpnRevocationListOcsp(d []interface{}) edpt.VpnRevocationListOcsp {

	count1 := len(d)
	var ret edpt.VpnRevocationListOcsp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OcspPri = in["ocsp_pri"].(string)
		ret.OcspSec = in["ocsp_sec"].(string)
	}
	return ret
}

func getSliceVpnSamplingEnable(d []interface{}) []edpt.VpnSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VpnSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpn(d *schema.ResourceData) edpt.Vpn {
	var ret edpt.Vpn
	ret.Inst.AsymmetricFlowSupport = d.Get("asymmetric_flow_support").(int)
	ret.Inst.Crl = getObjectVpnCrl3614(d.Get("crl").([]interface{}))
	ret.Inst.Default = getObjectVpnDefault3615(d.Get("default").([]interface{}))
	ret.Inst.EnableVpnMetrics = d.Get("enable_vpn_metrics").(int)
	ret.Inst.Error = getObjectVpnError3616(d.Get("error").([]interface{}))
	ret.Inst.Errordump = getObjectVpnErrordump3617(d.Get("errordump").([]interface{}))
	ret.Inst.ExtendedMatching = d.Get("extended_matching").(int)
	ret.Inst.FragmentAfterEncap = d.Get("fragment_after_encap").(int)
	ret.Inst.GroupList = getObjectVpnGroupList3618(d.Get("group_list").([]interface{}))
	ret.Inst.IkeAccEnable = d.Get("ike_acc_enable").(int)
	ret.Inst.IkeGatewayList = getSliceVpnIkeGatewayList(d.Get("ike_gateway_list").([]interface{}))
	ret.Inst.IkeLoggingEnable = d.Get("ike_logging_enable").(int)
	ret.Inst.IkeSa = getObjectVpnIkeSa3619(d.Get("ike_sa").([]interface{}))
	ret.Inst.IkeSaBrief = getObjectVpnIkeSaBrief3620(d.Get("ike_sa_brief").([]interface{}))
	ret.Inst.IkeSaClients = getObjectVpnIkeSaClients3621(d.Get("ike_sa_clients").([]interface{}))
	ret.Inst.IkeSaTimeout = d.Get("ike_sa_timeout").(int)
	ret.Inst.IkeStatsByGw = getObjectVpnIkeStatsByGw3622(d.Get("ike_stats_by_gw").([]interface{}))
	ret.Inst.IkeStatsGlobal = getObjectVpnIkeStatsGlobal3623(d.Get("ike_stats_global").([]interface{}))
	ret.Inst.IpsecCipherCheck = d.Get("ipsec_cipher_check").(int)
	ret.Inst.IpsecErrorDump = d.Get("ipsec_error_dump").(int)
	ret.Inst.IpsecGroupList = getSliceVpnIpsecGroupList(d.Get("ipsec_group_list").([]interface{}))
	ret.Inst.IpsecList = getSliceVpnIpsecList(d.Get("ipsec_list").([]interface{}))
	ret.Inst.IpsecMgmtDefaultPolicyDrop = d.Get("ipsec_mgmt_default_policy_drop").(int)
	ret.Inst.IpsecSa = getObjectVpnIpsecSa3625(d.Get("ipsec_sa").([]interface{}))
	ret.Inst.IpsecSaByGw = getObjectVpnIpsecSaByGw3626(d.Get("ipsec_sa_by_gw").([]interface{}))
	ret.Inst.IpsecSaClients = getObjectVpnIpsecSaClients3627(d.Get("ipsec_sa_clients").([]interface{}))
	ret.Inst.IpsecSaStatsList = getSliceVpnIpsecSaStatsList(d.Get("ipsec_sa_stats_list").([]interface{}))
	ret.Inst.JumboFragment = d.Get("jumbo_fragment").(int)
	ret.Inst.Log = getObjectVpnLog3628(d.Get("log").([]interface{}))
	ret.Inst.NatTraversalFlowAffinity = d.Get("nat_traversal_flow_affinity").(int)
	ret.Inst.Ocsp = getObjectVpnOcsp3629(d.Get("ocsp").([]interface{}))
	ret.Inst.RevocationList = getSliceVpnRevocationList(d.Get("revocation_list").([]interface{}))
	ret.Inst.SamplingEnable = getSliceVpnSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SignatureAuthentication = d.Get("signature_authentication").(int)
	ret.Inst.StatefulMode = d.Get("stateful_mode").(int)
	ret.Inst.TcpMssAdjustDisable = d.Get("tcp_mss_adjust_disable").(int)
	//omit uuid
	return ret
}
