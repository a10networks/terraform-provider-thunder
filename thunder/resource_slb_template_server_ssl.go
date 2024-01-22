package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateServerSsl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_server_ssl`: Server Side SSL Template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateServerSslCreate,
		UpdateContext: resourceSlbTemplateServerSslUpdate,
		ReadContext:   resourceSlbTemplateServerSslRead,
		DeleteContext: resourceSlbTemplateServerSslDelete,

		Schema: map[string]*schema.Schema{
			"alert_type": {
				Type: schema.TypeString, Optional: true, Description: "'fatal': Log fatal alerts;",
			},
			"ca_certs": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_cert": {
							Type: schema.TypeString, Optional: true, Description: "Specify CA certificate",
						},
						"ca_cert_partition_shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "CA Certificate Partition Shared",
						},
						"server_ocsp_srvr": {
							Type: schema.TypeString, Optional: true, Description: "Specify authentication server",
						},
						"server_ocsp_sg": {
							Type: schema.TypeString, Optional: true, Description: "Specify service-group (Service group name)",
						},
					},
				},
			},
			"certificate": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert": {
							Type: schema.TypeString, Optional: true, Description: "Certificate Name",
						},
						"key": {
							Type: schema.TypeString, Optional: true, Description: "Client private-key (Key Name)",
						},
						"passphrase": {
							Type: schema.TypeString, Optional: true, Description: "Password Phrase",
						},
						"shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Client Certificate and Key Partition Shared",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"cipher_template": {
				Type: schema.TypeString, Optional: true, Description: "Cipher Template Name",
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
			"close_notify": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send close notification when terminate connection",
			},
			"crl_certs": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crl": {
							Type: schema.TypeString, Optional: true, Description: "Certificate Revocation Lists (Certificate Revocation Lists file name)",
						},
						"crl_partition_shared": {
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
			"forward_proxy_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL forward proxy",
			},
			"handshake_logging_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL handshake logging",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Server SSL Template Name",
			},
			"ocsp_stapling": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ocsp-stapling support",
			},
			"renegotiation_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable SSL renegotiation",
			},
			"server_certificate_error": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error_type": {
							Type: schema.TypeString, Optional: true, Description: "'email': Notify the error via email; 'ignore': Ignore the error, which mean the connection can continue; 'logging': Log the error; 'trap': Notify the error by SNMP trap;",
						},
					},
				},
			},
			"server_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify Server Name",
			},
			"session_cache_size": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Session Cache Size (Maximum cache size. Default value 0 (Session ID reuse disabled))",
			},
			"session_cache_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Session Cache Timeout (Timeout value, in seconds. Default no timeout.)",
			},
			"session_ticket_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable server side session ticket support",
			},
			"shared_partition_cipher_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a cipher template from shared partition",
			},
			"ssli_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SSLi logging level, default is error logging only",
			},
			"sslilogging": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable all logging; 'all': enable all logging(error, info);",
			},
			"template_cipher_shared": {
				Type: schema.TypeString, Optional: true, Description: "Cipher Template Name",
			},
			"use_client_sni": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "use client SNI",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"version": {
				Type: schema.TypeInt, Optional: true, Default: 33, Description: "TLS/SSL version, default is the highest number supported (TLS/SSL version: 30-SSLv3.0, 31-TLSv1.0, 32-TLSv1.1, 33-TLSv1.2 and 34-TLSv1.3)",
			},
		},
	}
}
func resourceSlbTemplateServerSslCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerSslCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServerSsl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateServerSslRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateServerSslUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerSslUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServerSsl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateServerSslRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateServerSslDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerSslDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServerSsl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateServerSslRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerSslRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServerSsl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateServerSslCaCerts(d []interface{}) []edpt.SlbTemplateServerSslCaCerts {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateServerSslCaCerts, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateServerSslCaCerts
		oi.CaCert = in["ca_cert"].(string)
		oi.CaCertPartitionShared = in["ca_cert_partition_shared"].(int)
		oi.ServerOcspSrvr = in["server_ocsp_srvr"].(string)
		oi.ServerOcspSg = in["server_ocsp_sg"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateServerSslCertificate1474(d []interface{}) edpt.SlbTemplateServerSslCertificate1474 {

	count1 := len(d)
	var ret edpt.SlbTemplateServerSslCertificate1474
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cert = in["cert"].(string)
		ret.Key = in["key"].(string)
		ret.Passphrase = in["passphrase"].(string)
		//omit encrypted
		ret.Shared = in["shared"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateServerSslCipherWithoutPrioList(d []interface{}) []edpt.SlbTemplateServerSslCipherWithoutPrioList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateServerSslCipherWithoutPrioList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateServerSslCipherWithoutPrioList
		oi.CipherWoPrio = in["cipher_wo_prio"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateServerSslCrlCerts(d []interface{}) []edpt.SlbTemplateServerSslCrlCerts {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateServerSslCrlCerts, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateServerSslCrlCerts
		oi.Crl = in["crl"].(string)
		oi.CrlPartitionShared = in["crl_partition_shared"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateServerSslEcList(d []interface{}) []edpt.SlbTemplateServerSslEcList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateServerSslEcList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateServerSslEcList
		oi.Ec = in["ec"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateServerSslServerCertificateError(d []interface{}) []edpt.SlbTemplateServerSslServerCertificateError {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateServerSslServerCertificateError, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateServerSslServerCertificateError
		oi.ErrorType = in["error_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateServerSsl(d *schema.ResourceData) edpt.SlbTemplateServerSsl {
	var ret edpt.SlbTemplateServerSsl
	ret.Inst.AlertType = d.Get("alert_type").(string)
	ret.Inst.CaCerts = getSliceSlbTemplateServerSslCaCerts(d.Get("ca_certs").([]interface{}))
	ret.Inst.Certificate = getObjectSlbTemplateServerSslCertificate1474(d.Get("certificate").([]interface{}))
	ret.Inst.CipherTemplate = d.Get("cipher_template").(string)
	ret.Inst.CipherWithoutPrioList = getSliceSlbTemplateServerSslCipherWithoutPrioList(d.Get("cipher_without_prio_list").([]interface{}))
	ret.Inst.CloseNotify = d.Get("close_notify").(int)
	ret.Inst.CrlCerts = getSliceSlbTemplateServerSslCrlCerts(d.Get("crl_certs").([]interface{}))
	ret.Inst.Dgversion = d.Get("dgversion").(int)
	ret.Inst.DhType = d.Get("dh_type").(string)
	ret.Inst.EarlyData = d.Get("early_data").(int)
	ret.Inst.EcList = getSliceSlbTemplateServerSslEcList(d.Get("ec_list").([]interface{}))
	ret.Inst.EnableSsliFtpAlg = d.Get("enable_ssli_ftp_alg").(int)
	ret.Inst.EnableTlsAlertLogging = d.Get("enable_tls_alert_logging").(int)
	ret.Inst.ForwardProxyEnable = d.Get("forward_proxy_enable").(int)
	ret.Inst.HandshakeLoggingEnable = d.Get("handshake_logging_enable").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OcspStapling = d.Get("ocsp_stapling").(int)
	ret.Inst.RenegotiationDisable = d.Get("renegotiation_disable").(int)
	ret.Inst.ServerCertificateError = getSliceSlbTemplateServerSslServerCertificateError(d.Get("server_certificate_error").([]interface{}))
	ret.Inst.ServerName = d.Get("server_name").(string)
	ret.Inst.SessionCacheSize = d.Get("session_cache_size").(int)
	ret.Inst.SessionCacheTimeout = d.Get("session_cache_timeout").(int)
	ret.Inst.SessionTicketEnable = d.Get("session_ticket_enable").(int)
	ret.Inst.SharedPartitionCipherTemplate = d.Get("shared_partition_cipher_template").(int)
	ret.Inst.SsliLogging = d.Get("ssli_logging").(int)
	ret.Inst.Sslilogging = d.Get("sslilogging").(string)
	ret.Inst.TemplateCipherShared = d.Get("template_cipher_shared").(string)
	ret.Inst.UseClientSni = d.Get("use_client_sni").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Version = d.Get("version").(int)
	return ret
}
