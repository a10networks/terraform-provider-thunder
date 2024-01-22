package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIkeGateway() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vpn_ike_gateway`: IKE-gateway settings\n\n__PLACEHOLDER__",
		CreateContext: resourceVpnIkeGatewayCreate,
		UpdateContext: resourceVpnIkeGatewayUpdate,
		ReadContext:   resourceVpnIkeGatewayRead,
		DeleteContext: resourceVpnIkeGatewayDelete,

		Schema: map[string]*schema.Schema{
			"auth_method": {
				Type: schema.TypeString, Optional: true, Default: "preshare-key", Description: "'preshare-key': Authenticate the remote gateway using a pre-shared key (Default); 'rsa-signature': Authenticate the remote gateway using an RSA certificate; 'ecdsa-signature': Authenticate the remote gateway using an ECDSA certificate; 'eap-radius': Authenticate the remote gateway using an EAP Radius server; 'eap-tls': Authenticate the remote gateway using EAP TLS;",
			},
			"configuration_payload": {
				Type: schema.TypeString, Optional: true, Description: "'dhcp': Enable DHCP configuration-payload; 'radius': Enable RADIUS configuration-payload;",
			},
			"dh_group": {
				Type: schema.TypeString, Optional: true, Default: "1", Description: "'1': Diffie-Hellman group 1 - 768-bit(Default); '2': Diffie-Hellman group 2 - 1024-bit; '5': Diffie-Hellman group 5 - 1536-bit; '14': Diffie-Hellman group 14 - 2048-bit; '15': Diffie-Hellman group 15 - 3072-bit; '16': Diffie-Hellman group 16 - 4096-bit; '18': Diffie-Hellman group 18 - 8192-bit; '19': Diffie-Hellman group 19 - 256-bit Elliptic Curve; '20': Diffie-Hellman group 20 - 384-bit Elliptic Curve;",
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
			"disable_rekey": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable initiating rekey",
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
			"fragment_size": {
				Type: schema.TypeInt, Optional: true, Description: "Enable IKE message fragment and set fragment size",
			},
			"hash": {
				Type: schema.TypeString, Optional: true, Description: "'sha256': Secure Hash Algorithm 256; 'sha384': Secure Hash Algorithm 384; 'sha512': Secure Hash Algorithm 512;",
			},
			"ike_version": {
				Type: schema.TypeString, Optional: true, Default: "v2", Description: "'v1': IKEv1 key exchange; 'v2': IKEv2 key exchange;",
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
			"lifetime": {
				Type: schema.TypeInt, Optional: true, Default: 86400, Description: "IKE SA age in seconds",
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
			"local_id": {
				Type: schema.TypeString, Optional: true, Description: "Local Gateway Identity",
			},
			"mode": {
				Type: schema.TypeString, Optional: true, Default: "main", Description: "'main': Negotiate Main mode (Default); 'aggressive': Negotiate Aggressive mode;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "IKE-gateway name",
			},
			"nat_traversal": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"preshare_key_value": {
				Type: schema.TypeString, Optional: true, Description: "pre-shared key",
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
			"remote_id": {
				Type: schema.TypeString, Optional: true, Description: "Remote Gateway Identity",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
		},
	}
}
func resourceVpnIkeGatewayCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIkeGatewayCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIkeGateway(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnIkeGatewayRead(ctx, d, meta)
	}
	return diags
}

func resourceVpnIkeGatewayUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIkeGatewayUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIkeGateway(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnIkeGatewayRead(ctx, d, meta)
	}
	return diags
}
func resourceVpnIkeGatewayDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIkeGatewayDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIkeGateway(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVpnIkeGatewayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIkeGatewayRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIkeGateway(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVpnIkeGatewayDhcpServer(d []interface{}) edpt.VpnIkeGatewayDhcpServer {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayDhcpServer
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Pri = getObjectVpnIkeGatewayDhcpServerPri(in["pri"].([]interface{}))
		ret.Sec = getObjectVpnIkeGatewayDhcpServerSec(in["sec"].([]interface{}))
	}
	return ret
}

func getObjectVpnIkeGatewayDhcpServerPri(d []interface{}) edpt.VpnIkeGatewayDhcpServerPri {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayDhcpServerPri
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DhcpPriIpv4 = in["dhcp_pri_ipv4"].(string)
	}
	return ret
}

func getObjectVpnIkeGatewayDhcpServerSec(d []interface{}) edpt.VpnIkeGatewayDhcpServerSec {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayDhcpServerSec
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DhcpSecIpv4 = in["dhcp_sec_ipv4"].(string)
	}
	return ret
}

func getObjectVpnIkeGatewayDpd(d []interface{}) edpt.VpnIkeGatewayDpd {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayDpd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.Retry = in["retry"].(int)
	}
	return ret
}

func getSliceVpnIkeGatewayEncCfg(d []interface{}) []edpt.VpnIkeGatewayEncCfg {

	count1 := len(d)
	ret := make([]edpt.VpnIkeGatewayEncCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIkeGatewayEncCfg
		oi.Encryption = in["encryption"].(string)
		oi.Hash = in["hash"].(string)
		oi.Prf = in["prf"].(string)
		oi.Priority = in["priority"].(int)
		oi.Gcm_priority = in["gcm_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnIkeGatewayLocalAddress(d []interface{}) edpt.VpnIkeGatewayLocalAddress {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayLocalAddress
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LocalIp = in["local_ip"].(string)
		ret.LocalIpv6 = in["local_ipv6"].(string)
	}
	return ret
}

func getObjectVpnIkeGatewayLocalCert(d []interface{}) edpt.VpnIkeGatewayLocalCert {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayLocalCert
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LocalCertName = in["local_cert_name"].(string)
	}
	return ret
}

func getObjectVpnIkeGatewayRadiusServer(d []interface{}) edpt.VpnIkeGatewayRadiusServer {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayRadiusServer
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RadiusPri = in["radius_pri"].(string)
		ret.RadiusSec = in["radius_sec"].(string)
	}
	return ret
}

func getObjectVpnIkeGatewayRemoteAddress(d []interface{}) edpt.VpnIkeGatewayRemoteAddress {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayRemoteAddress
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RemoteIp = in["remote_ip"].(string)
		ret.Dns = in["dns"].(string)
		ret.RemoteIpv6 = in["remote_ipv6"].(string)
	}
	return ret
}

func getObjectVpnIkeGatewayRemoteCaCert(d []interface{}) edpt.VpnIkeGatewayRemoteCaCert {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayRemoteCaCert
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RemoteCertName = in["remote_cert_name"].(string)
	}
	return ret
}

func getSliceVpnIkeGatewaySamplingEnable(d []interface{}) []edpt.VpnIkeGatewaySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VpnIkeGatewaySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIkeGatewaySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnIkeGatewayVrid(d []interface{}) edpt.VpnIkeGatewayVrid {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayVrid
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Default = in["default"].(int)
		ret.VridNum = in["vrid_num"].(int)
	}
	return ret
}

func dataToEndpointVpnIkeGateway(d *schema.ResourceData) edpt.VpnIkeGateway {
	var ret edpt.VpnIkeGateway
	ret.Inst.AuthMethod = d.Get("auth_method").(string)
	ret.Inst.ConfigurationPayload = d.Get("configuration_payload").(string)
	ret.Inst.DhGroup = d.Get("dh_group").(string)
	ret.Inst.DhcpServer = getObjectVpnIkeGatewayDhcpServer(d.Get("dhcp_server").([]interface{}))
	ret.Inst.DisableRekey = d.Get("disable_rekey").(int)
	ret.Inst.Dpd = getObjectVpnIkeGatewayDpd(d.Get("dpd").([]interface{}))
	ret.Inst.EncCfg = getSliceVpnIkeGatewayEncCfg(d.Get("enc_cfg").([]interface{}))
	ret.Inst.FragmentSize = d.Get("fragment_size").(int)
	ret.Inst.Hash = d.Get("hash").(string)
	ret.Inst.IkeVersion = d.Get("ike_version").(string)
	ret.Inst.InterfaceManagement = d.Get("interface_management").(int)
	ret.Inst.Key = d.Get("key").(string)
	ret.Inst.KeyPassphrase = d.Get("key_passphrase").(string)
	//omit key_passphrase_encrypted
	ret.Inst.Lifetime = d.Get("lifetime").(int)
	ret.Inst.LocalAddress = getObjectVpnIkeGatewayLocalAddress(d.Get("local_address").([]interface{}))
	ret.Inst.LocalCert = getObjectVpnIkeGatewayLocalCert(d.Get("local_cert").([]interface{}))
	ret.Inst.LocalId = d.Get("local_id").(string)
	ret.Inst.Mode = d.Get("mode").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NatTraversal = d.Get("nat_traversal").(int)
	//omit preshare_key_encrypted
	ret.Inst.PreshareKeyValue = d.Get("preshare_key_value").(string)
	ret.Inst.RadiusServer = getObjectVpnIkeGatewayRadiusServer(d.Get("radius_server").([]interface{}))
	ret.Inst.RemoteAddress = getObjectVpnIkeGatewayRemoteAddress(d.Get("remote_address").([]interface{}))
	ret.Inst.RemoteCaCert = getObjectVpnIkeGatewayRemoteCaCert(d.Get("remote_ca_cert").([]interface{}))
	ret.Inst.RemoteId = d.Get("remote_id").(string)
	ret.Inst.SamplingEnable = getSliceVpnIkeGatewaySamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Vrid = getObjectVpnIkeGatewayVrid(d.Get("vrid").([]interface{}))
	return ret
}
