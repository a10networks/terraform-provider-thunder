package thunder

//Thunder resource SlbTemplateServerSSL

import (
	"context"
	"fmt"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateServerSSL() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateServerSSLCreate,
		UpdateContext: resourceSlbTemplateServerSSLUpdate,
		ReadContext:   resourceSlbTemplateServerSSLRead,
		DeleteContext: resourceSlbTemplateServerSSLDelete,
		Schema: map[string]*schema.Schema{
			"session_cache_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"enable_tls_alert_logging": {
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
			"key_shared_str": {
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
					},
				},
			},
			"use_client_sni": {
				Type:        schema.TypeInt,
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
						"ca_cert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_ocsp_sg": {
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
					},
				},
			},
			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_cipher_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"key_shared_passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"close_notify": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ocsp_stapling": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cert_shared_str": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"session_ticket_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"session_cache_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"version": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"renegotiation_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dh_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"key_shared_encrypted": {
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
			"passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dgversion": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cipher_template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ssli_logging": {
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
		},
	}
}

func resourceSlbTemplateServerSSLCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateServerSSL (Inside resourceSlbTemplateServerSSLCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateServerSSL(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateServerSSL --")
		d.SetId(name)
		err := go_thunder.PostSlbTemplateServerSSL(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateServerSSLRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateServerSSLRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbTemplateServerSSL (Inside resourceSlbTemplateServerSSLRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTemplateServerSSL(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbTemplateServerSSLUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateServerSSL   (Inside resourceSlbTemplateServerSSLUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateServerSSL(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateServerSSL ")
		d.SetId(name)
		err := go_thunder.PutSlbTemplateServerSSL(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateServerSSLRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateServerSSLDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateServerSSLDelete) " + name)
		err := go_thunder.DeleteSlbTemplateServerSSL(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateServerSSL(d *schema.ResourceData) go_thunder.ServerSSL {
	var vc go_thunder.ServerSSL
	var c go_thunder.ServerSSLInstance

	c.SessionCacheTimeout = d.Get("session_cache_timeout").(int)
	c.CipherTemplate = d.Get("cipher_template").(string)
	c.Sslilogging = d.Get("sslilogging").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.Passphrase = d.Get("passphrase").(string)
	c.OcspStapling = d.Get("ocsp_stapling").(int)

	CrlCertsCount := d.Get("crl_certs.#").(int)
	c.Crl = make([]go_thunder.CrlCerts, 0, CrlCertsCount)

	for i := 0; i < CrlCertsCount; i++ {
		var obj1 go_thunder.CrlCerts
		prefix := fmt.Sprintf("crl_certs.%d.", i)
		obj1.Crl = d.Get(prefix + "crl").(string)
		c.Crl = append(c.Crl, obj1)
	}

	c.KeySharedStr = d.Get("key_shared_str").(string)
	c.TemplateCipherShared = d.Get("template_cipher_shared").(string)
	c.Dgversion = d.Get("dgversion").(int)
	c.CertSharedStr = d.Get("cert_shared_str").(string)
	c.Version = d.Get("version").(int)

	EcListCount := d.Get("ec_list.#").(int)
	c.Ec = make([]go_thunder.EcList, 0, EcListCount)

	for i := 0; i < EcListCount; i++ {
		var obj2 go_thunder.EcList
		prefix := fmt.Sprintf("ec_list.%d.", i)
		obj2.Ec = d.Get(prefix + "ec").(string)
		c.Ec = append(c.Ec, obj2)
	}

	c.Encrypted = d.Get("encrypted").(string)
	c.SsliLogging = d.Get("ssli_logging").(int)
	c.SessionCacheSize = d.Get("session_cache_size").(int)
	c.DhType = d.Get("dh_type").(string)
	c.UseClientSni = d.Get("use_client_sni").(int)
	c.ForwardProxyEnable = d.Get("forward_proxy_enable").(int)
	c.Key = d.Get("key").(string)
	c.KeySharedEncrypted = d.Get("key_shared_encrypted").(string)

	CipherWithoutPrioListCount := d.Get("cipher_without_prio_list.#").(int)
	c.CipherWoPrio = make([]go_thunder.CipherWithoutPrioList, 0, CipherWithoutPrioListCount)

	for i := 0; i < CipherWithoutPrioListCount; i++ {
		var obj3 go_thunder.CipherWithoutPrioList
		prefix := fmt.Sprintf("cipher_without_prio_list.%d.", i)
		obj3.CipherWoPrio = d.Get(prefix + "cipher_wo_prio").(string)
		c.CipherWoPrio = append(c.CipherWoPrio, obj3)
	}

	CaCertsCount := d.Get("ca_certs.#").(int)
	c.CaCert = make([]go_thunder.CaCerts, 0, CaCertsCount)

	for i := 0; i < CaCertsCount; i++ {
		var obj4 go_thunder.CaCerts
		prefix := fmt.Sprintf("ca_certs.%d.", i)
		obj4.CaCert = d.Get(prefix + "ca_cert").(string)
		obj4.CaCertPartitionShared = d.Get(prefix + "ca_cert_partition_shared").(int)
		obj4.ServerOcspSg = d.Get(prefix + "server_ocsp_sg").(string)
		obj4.ServerOcspSrvr = d.Get(prefix + "server_ocsp_srvr").(string)
		c.CaCert = append(c.CaCert, obj4)
	}

	c.Name = d.Get("name").(string)
	c.SharedPartitionCipherTemplate = d.Get("shared_partition_cipher_template").(int)
	c.EnableTLSAlertLogging = d.Get("enable_tls_alert_logging").(int)
	c.SessionTicketEnable = d.Get("session_ticket_enable").(int)
	c.AlertType = d.Get("alert_type").(string)
	c.Cert = d.Get("cert").(string)
	c.HandshakeLoggingEnable = d.Get("handshake_logging_enable").(int)
	c.RenegotiationDisable = d.Get("renegotiation_disable").(int)

	ServerCertificateErrorCount := d.Get("server_certificate_error.#").(int)
	c.ErrorType = make([]go_thunder.ServerCertificateError, 0, ServerCertificateErrorCount)

	for i := 0; i < ServerCertificateErrorCount; i++ {
		var obj5 go_thunder.ServerCertificateError
		prefix := fmt.Sprintf("server_certificate_error.%d.", i)
		obj5.ErrorType = d.Get(prefix + "error_type").(string)
		c.ErrorType = append(c.ErrorType, obj5)
	}

	c.CloseNotify = d.Get("close_notify").(int)
	c.KeySharedPassphrase = d.Get("key_shared_passphrase").(string)

	vc.UUID = c
	return vc
}
