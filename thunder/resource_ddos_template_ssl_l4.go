package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateSslL4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_ssl_l4`: SSL-L4 template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateSslL4Create,
		UpdateContext: resourceDdosTemplateSslL4Update,
		ReadContext:   resourceDdosTemplateSslL4Read,
		DeleteContext: resourceDdosTemplateSslL4Delete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': drop; 'reset': reset;",
			},
			"allow_non_tls": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow Non-TLS (SSLv3 and lower) traffic (Warning: security may be compromised)",
			},
			"auth_config_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timeout": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Connection timeout",
						},
						"trials": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Number of failed handshakes",
						},
						"auth_handshake_fail_action": {
							Type: schema.TypeString, Optional: true, Description: "'blacklist-src': Blacklist-src when auth handshake fails;",
						},
					},
				},
			},
			"cert_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert": {
							Type: schema.TypeString, Optional: true, Description: "SSL certificate",
						},
						"key": {
							Type: schema.TypeString, Optional: true, Description: "SSL key",
						},
						"key_passphrase": {
							Type: schema.TypeString, Optional: true, Description: "Password Phrase",
						},
					},
				},
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable this template",
			},
			"multi_pu_threshold_distribution": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"multi_pu_threshold_distribution_value": {
							Type: schema.TypeInt, Optional: true, Description: "Destination side rate limit only. Default: 0",
						},
						"multi_pu_threshold_distribution_disable": {
							Type: schema.TypeString, Optional: true, Description: "'disable': Destination side rate limit only. Default: Enable;",
						},
					},
				},
			},
			"renegotiation": {
				Type: schema.TypeInt, Optional: true, Description: "Configure renegotiation limiting for SSL (Number of renegotiation allowed)",
			},
			"request_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Configure rate limiting for SSL",
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
						"server_key": {
							Type: schema.TypeString, Optional: true, Description: "Server Private Key associated to SNI (Server Private Key Name)",
						},
						"server_passphrase": {
							Type: schema.TypeString, Optional: true, Description: "Password Phrase",
						},
					},
				},
			},
			"ssl_l4_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"ssl_traffic_check": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header_inspection": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inspect ssl header",
						},
						"header_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets with bad ssl header; 'ignore': Forward packets with bad ssl header;",
						},
						"check_resumed_connection": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply checks to SSL connections initialized by ACK packets",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
		},
	}
}
func resourceDdosTemplateSslL4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSslL4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSslL4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateSslL4Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateSslL4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSslL4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSslL4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateSslL4Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateSslL4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSslL4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSslL4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateSslL4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSslL4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSslL4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosTemplateSslL4AuthConfigCfg(d []interface{}) edpt.DdosTemplateSslL4AuthConfigCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSslL4AuthConfigCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Timeout = in["timeout"].(int)
		ret.Trials = in["trials"].(int)
		ret.AuthHandshakeFailAction = in["auth_handshake_fail_action"].(string)
	}
	return ret
}

func getObjectDdosTemplateSslL4CertCfg(d []interface{}) edpt.DdosTemplateSslL4CertCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSslL4CertCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cert = in["cert"].(string)
		ret.Key = in["key"].(string)
		ret.KeyPassphrase = in["key_passphrase"].(string)
		//omit key_encrypted
	}
	return ret
}

func getObjectDdosTemplateSslL4MultiPuThresholdDistribution(d []interface{}) edpt.DdosTemplateSslL4MultiPuThresholdDistribution {

	count1 := len(d)
	var ret edpt.DdosTemplateSslL4MultiPuThresholdDistribution
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MultiPuThresholdDistributionValue = in["multi_pu_threshold_distribution_value"].(int)
		ret.MultiPuThresholdDistributionDisable = in["multi_pu_threshold_distribution_disable"].(string)
	}
	return ret
}

func getSliceDdosTemplateSslL4ServerNameList(d []interface{}) []edpt.DdosTemplateSslL4ServerNameList {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateSslL4ServerNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateSslL4ServerNameList
		oi.ServerName = in["server_name"].(string)
		oi.ServerCert = in["server_cert"].(string)
		oi.ServerKey = in["server_key"].(string)
		oi.ServerPassphrase = in["server_passphrase"].(string)
		//omit server_encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateSslL4SslTrafficCheck299(d []interface{}) edpt.DdosTemplateSslL4SslTrafficCheck299 {

	count1 := len(d)
	var ret edpt.DdosTemplateSslL4SslTrafficCheck299
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HeaderInspection = in["header_inspection"].(int)
		ret.HeaderAction = in["header_action"].(string)
		ret.CheckResumedConnection = in["check_resumed_connection"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosTemplateSslL4(d *schema.ResourceData) edpt.DdosTemplateSslL4 {
	var ret edpt.DdosTemplateSslL4
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.AllowNonTls = d.Get("allow_non_tls").(int)
	ret.Inst.AuthConfigCfg = getObjectDdosTemplateSslL4AuthConfigCfg(d.Get("auth_config_cfg").([]interface{}))
	ret.Inst.CertCfg = getObjectDdosTemplateSslL4CertCfg(d.Get("cert_cfg").([]interface{}))
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.MultiPuThresholdDistribution = getObjectDdosTemplateSslL4MultiPuThresholdDistribution(d.Get("multi_pu_threshold_distribution").([]interface{}))
	ret.Inst.Renegotiation = d.Get("renegotiation").(int)
	ret.Inst.RequestRateLimit = d.Get("request_rate_limit").(int)
	ret.Inst.ServerNameList = getSliceDdosTemplateSslL4ServerNameList(d.Get("server_name_list").([]interface{}))
	ret.Inst.SslL4TmplName = d.Get("ssl_l4_tmpl_name").(string)
	ret.Inst.SslTrafficCheck = getObjectDdosTemplateSslL4SslTrafficCheck299(d.Get("ssl_traffic_check").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
