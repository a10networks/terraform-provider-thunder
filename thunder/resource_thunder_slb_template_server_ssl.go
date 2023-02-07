package thunder

//Thunder resource SlbTemplateServerSsl

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateServerSsl() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateServerSslCreate,
		UpdateContext: resourceSlbTemplateServerSslUpdate,
		ReadContext:   resourceSlbTemplateServerSslRead,
		DeleteContext: resourceSlbTemplateServerSslDelete,
		Schema: map[string]*schema.Schema{
			"name": {
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
						"ca_cert_partition_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_ocsp_srvr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_ocsp_sg": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"server_name": {
				Type:        schema.TypeString,
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
						"crl_partition_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
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
			"handshake_logging_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"close_notify": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"session_ticket_enable": {
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
			"server_certificate_error": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
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
			"ocsp_stapling": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"use_client_sni": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"renegotiation_disable": {
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
			"cipher_template": {
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
			"certificate": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
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
			"ifnum": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateServerSslCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateServerSsl (Inside resourceSlbTemplateServerSslCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateServerSsl(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateServerSsl --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateServerSsl(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateServerSslRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateServerSslRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateServerSsl (Inside resourceSlbTemplateServerSslRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateServerSsl(client.Token, name1, client.Host)
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

func resourceSlbTemplateServerSslUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplateServerSsl   (Inside resourceSlbTemplateServerSslUpdate) ")
		data := dataToSlbTemplateServerSsl(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateServerSsl ")
		err := go_thunder.PutSlbTemplateServerSsl(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateServerSslRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateServerSslDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateServerSslDelete) " + name1)
		err := go_thunder.DeleteSlbTemplateServerSsl(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplateServerSsl(d *schema.ResourceData) go_thunder.SlbTemplateServerSsl {
	var vc go_thunder.SlbTemplateServerSsl
	var c go_thunder.SlbTemplateServerSslInstance
	c.SlbTemplateServerSslInstanceName = d.Get("name").(string)

	SlbTemplateServerSslInstanceCaCertsCount := d.Get("ca_certs.#").(int)
	c.SlbTemplateServerSslInstanceCaCertsCaCert = make([]go_thunder.SlbTemplateServerSslInstanceCaCerts, 0, SlbTemplateServerSslInstanceCaCertsCount)

	for i := 0; i < SlbTemplateServerSslInstanceCaCertsCount; i++ {
		var obj1 go_thunder.SlbTemplateServerSslInstanceCaCerts
		prefix1 := fmt.Sprintf("ca_certs.%d.", i)
		obj1.SlbTemplateServerSslInstanceCaCertsCaCert = d.Get(prefix1 + "ca_cert").(string)
		obj1.SlbTemplateServerSslInstanceCaCertsCaCertPartitionShared = d.Get(prefix1 + "ca_cert_partition_shared").(int)
		obj1.SlbTemplateServerSslInstanceCaCertsServerOcspSrvr = d.Get(prefix1 + "server_ocsp_srvr").(string)
		obj1.SlbTemplateServerSslInstanceCaCertsServerOcspSg = d.Get(prefix1 + "server_ocsp_sg").(string)
		c.SlbTemplateServerSslInstanceCaCertsCaCert = append(c.SlbTemplateServerSslInstanceCaCertsCaCert, obj1)
	}

	c.SlbTemplateServerSslInstanceServerName = d.Get("server_name").(string)

	SlbTemplateServerSslInstanceCrlCertsCount := d.Get("crl_certs.#").(int)
	c.SlbTemplateServerSslInstanceCrlCertsCrl = make([]go_thunder.SlbTemplateServerSslInstanceCrlCerts, 0, SlbTemplateServerSslInstanceCrlCertsCount)

	for i := 0; i < SlbTemplateServerSslInstanceCrlCertsCount; i++ {
		var obj2 go_thunder.SlbTemplateServerSslInstanceCrlCerts
		prefix2 := fmt.Sprintf("crl_certs.%d.", i)
		obj2.SlbTemplateServerSslInstanceCrlCertsCrl = d.Get(prefix2 + "crl").(string)
		obj2.SlbTemplateServerSslInstanceCrlCertsCrlPartitionShared = d.Get(prefix2 + "crl_partition_shared").(int)
		c.SlbTemplateServerSslInstanceCrlCertsCrl = append(c.SlbTemplateServerSslInstanceCrlCertsCrl, obj2)
	}

	SlbTemplateServerSslInstanceCipherWithoutPrioListCount := d.Get("cipher_without_prio_list.#").(int)
	c.SlbTemplateServerSslInstanceCipherWithoutPrioListCipherWoPrio = make([]go_thunder.SlbTemplateServerSslInstanceCipherWithoutPrioList, 0, SlbTemplateServerSslInstanceCipherWithoutPrioListCount)

	for i := 0; i < SlbTemplateServerSslInstanceCipherWithoutPrioListCount; i++ {
		var obj3 go_thunder.SlbTemplateServerSslInstanceCipherWithoutPrioList
		prefix3 := fmt.Sprintf("cipher_without_prio_list.%d.", i)
		obj3.SlbTemplateServerSslInstanceCipherWithoutPrioListCipherWoPrio = d.Get(prefix3 + "cipher_wo_prio").(string)
		c.SlbTemplateServerSslInstanceCipherWithoutPrioListCipherWoPrio = append(c.SlbTemplateServerSslInstanceCipherWithoutPrioListCipherWoPrio, obj3)
	}

	c.SlbTemplateServerSslInstanceDhType = d.Get("dh_type").(string)

	SlbTemplateServerSslInstanceEcListCount := d.Get("ec_list.#").(int)
	c.SlbTemplateServerSslInstanceEcListEc = make([]go_thunder.SlbTemplateServerSslInstanceEcList, 0, SlbTemplateServerSslInstanceEcListCount)

	for i := 0; i < SlbTemplateServerSslInstanceEcListCount; i++ {
		var obj4 go_thunder.SlbTemplateServerSslInstanceEcList
		prefix4 := fmt.Sprintf("ec_list.%d.", i)
		obj4.SlbTemplateServerSslInstanceEcListEc = d.Get(prefix4 + "ec").(string)
		c.SlbTemplateServerSslInstanceEcListEc = append(c.SlbTemplateServerSslInstanceEcListEc, obj4)
	}

	c.SlbTemplateServerSslInstanceEnableTLSAlertLogging = d.Get("enable_tls_alert_logging").(int)
	c.SlbTemplateServerSslInstanceAlertType = d.Get("alert_type").(string)
	c.SlbTemplateServerSslInstanceHandshakeLoggingEnable = d.Get("handshake_logging_enable").(int)
	c.SlbTemplateServerSslInstanceCloseNotify = d.Get("close_notify").(int)
	c.SlbTemplateServerSslInstanceForwardProxyEnable = d.Get("forward_proxy_enable").(int)
	c.SlbTemplateServerSslInstanceSessionTicketEnable = d.Get("session_ticket_enable").(int)
	c.SlbTemplateServerSslInstanceVersion = d.Get("version").(int)
	c.SlbTemplateServerSslInstanceDgversion = d.Get("dgversion").(int)

	SlbTemplateServerSslInstanceServerCertificateErrorCount := d.Get("server_certificate_error.#").(int)
	c.SlbTemplateServerSslInstanceServerCertificateErrorErrorType = make([]go_thunder.SlbTemplateServerSslInstanceServerCertificateError, 0, SlbTemplateServerSslInstanceServerCertificateErrorCount)

	for i := 0; i < SlbTemplateServerSslInstanceServerCertificateErrorCount; i++ {
		var obj5 go_thunder.SlbTemplateServerSslInstanceServerCertificateError
		prefix5 := fmt.Sprintf("server_certificate_error.%d.", i)
		obj5.SlbTemplateServerSslInstanceServerCertificateErrorErrorType = d.Get(prefix5 + "error_type").(string)
		c.SlbTemplateServerSslInstanceServerCertificateErrorErrorType = append(c.SlbTemplateServerSslInstanceServerCertificateErrorErrorType, obj5)
	}

	c.SlbTemplateServerSslInstanceSsliLogging = d.Get("ssli_logging").(int)
	c.SlbTemplateServerSslInstanceSslilogging = d.Get("sslilogging").(string)
	c.SlbTemplateServerSslInstanceOcspStapling = d.Get("ocsp_stapling").(int)
	c.SlbTemplateServerSslInstanceUseClientSni = d.Get("use_client_sni").(int)
	c.SlbTemplateServerSslInstanceRenegotiationDisable = d.Get("renegotiation_disable").(int)
	c.SlbTemplateServerSslInstanceSessionCacheSize = d.Get("session_cache_size").(int)
	c.SlbTemplateServerSslInstanceSessionCacheTimeout = d.Get("session_cache_timeout").(int)
	c.SlbTemplateServerSslInstanceCipherTemplate = d.Get("cipher_template").(string)
	c.SlbTemplateServerSslInstanceSharedPartitionCipherTemplate = d.Get("shared_partition_cipher_template").(int)
	c.SlbTemplateServerSslInstanceTemplateCipherShared = d.Get("template_cipher_shared").(string)
	c.SlbTemplateServerSslInstanceEnableSsliFtpAlg = d.Get("enable_ssli_ftp_alg").(int)
	c.SlbTemplateServerSslInstanceEarlyData = d.Get("early_data").(int)
	c.SlbTemplateServerSslInstanceUserTag = d.Get("user_tag").(string)

	var obj6 go_thunder.SlbTemplateServerSslInstanceCertificate
	prefix6 := "certificate.0."
	obj6.SlbTemplateServerSslInstanceCertificateCert = d.Get(prefix6 + "cert").(string)
	obj6.SlbTemplateServerSslInstanceCertificateKey = d.Get(prefix6 + "key").(string)
	obj6.SlbTemplateServerSslInstanceCertificatePassphrase = d.Get(prefix6 + "passphrase").(string)
	obj6.SlbTemplateServerSslInstanceCertificateShared = d.Get(prefix6 + "shared").(int)

	c.SlbTemplateServerSslInstanceCertificateCert = obj6

	vc.SlbTemplateServerSslInstanceName = c
	return vc
}
