package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateSslL4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_ssl_l4`: SSL-L4 template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateSslL4Create,
		UpdateContext: resourceDdosZoneTemplateSslL4Update,
		ReadContext:   resourceDdosZoneTemplateSslL4Read,
		DeleteContext: resourceDdosZoneTemplateSslL4Delete,

		Schema: map[string]*schema.Schema{
			"allow_non_tls": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow Non-TLS (SSLv3 and lower) traffic (Warning: security may be compromised)",
			},
			"auth_handshake": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_handshake_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Connection timeout (default 5 seconds) and trials (default 5 times) (DST support only)",
						},
						"auth_handshake_trials": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Number of failed handshakes before entry marked black",
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
						"auth_handshake_pass_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for passing the authentication",
						},
						"auth_handshake_pass_action": {
							Type: schema.TypeString, Optional: true, Description: "'authenticate-src': authenticate-src (Default);",
						},
						"auth_handshake_fail_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for failing the authentication",
						},
						"auth_handshake_fail_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
						},
					},
				},
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable this template",
			},
			"dst": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate_limit": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"request": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dst_request_rate_limit": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dst_request_rate_limit_action_list_name": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"dst_request_rate_limit_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'reset': Reset client connection;",
												},
											},
										},
									},
								},
							},
						},
					},
				},
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
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num_renegotiation": {
							Type: schema.TypeInt, Optional: true, Description: "Number of renegotiation allowed",
						},
						"ssl_l4_reneg_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"ssl_l4_reneg_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
						},
					},
				},
			},
			"src": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate_limit": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"request": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_request_rate_limit": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"src_request_rate_limit_action_list_name": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"src_request_rate_limit_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'reset': Reset client connection;",
												},
											},
										},
									},
								},
							},
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
func resourceDdosZoneTemplateSslL4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSslL4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSslL4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateSslL4Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateSslL4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSslL4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSslL4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateSslL4Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateSslL4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSslL4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSslL4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateSslL4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSslL4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSslL4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneTemplateSslL4AuthHandshake(d []interface{}) edpt.DdosZoneTemplateSslL4AuthHandshake {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSslL4AuthHandshake
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthHandshakeTimeout = in["auth_handshake_timeout"].(int)
		ret.AuthHandshakeTrials = in["auth_handshake_trials"].(int)
		ret.CertCfg = getObjectDdosZoneTemplateSslL4AuthHandshakeCertCfg(in["cert_cfg"].([]interface{}))
		ret.ServerNameList = getSliceDdosZoneTemplateSslL4AuthHandshakeServerNameList(in["server_name_list"].([]interface{}))
		ret.AuthHandshakePassActionListName = in["auth_handshake_pass_action_list_name"].(string)
		ret.AuthHandshakePassAction = in["auth_handshake_pass_action"].(string)
		ret.AuthHandshakeFailActionListName = in["auth_handshake_fail_action_list_name"].(string)
		ret.AuthHandshakeFailAction = in["auth_handshake_fail_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateSslL4AuthHandshakeCertCfg(d []interface{}) edpt.DdosZoneTemplateSslL4AuthHandshakeCertCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSslL4AuthHandshakeCertCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cert = in["cert"].(string)
		ret.Key = in["key"].(string)
		ret.KeyPassphrase = in["key_passphrase"].(string)
		//omit key_encrypted
	}
	return ret
}

func getSliceDdosZoneTemplateSslL4AuthHandshakeServerNameList(d []interface{}) []edpt.DdosZoneTemplateSslL4AuthHandshakeServerNameList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateSslL4AuthHandshakeServerNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateSslL4AuthHandshakeServerNameList
		oi.ServerName = in["server_name"].(string)
		oi.ServerCert = in["server_cert"].(string)
		oi.ServerKey = in["server_key"].(string)
		oi.ServerPassphrase = in["server_passphrase"].(string)
		//omit server_encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateSslL4Dst(d []interface{}) edpt.DdosZoneTemplateSslL4Dst {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSslL4Dst
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimit = getObjectDdosZoneTemplateSslL4DstRateLimit(in["rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateSslL4DstRateLimit(d []interface{}) edpt.DdosZoneTemplateSslL4DstRateLimit {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSslL4DstRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Request = getObjectDdosZoneTemplateSslL4DstRateLimitRequest(in["request"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateSslL4DstRateLimitRequest(d []interface{}) edpt.DdosZoneTemplateSslL4DstRateLimitRequest {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSslL4DstRateLimitRequest
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstRequestRateLimit = in["dst_request_rate_limit"].(int)
		ret.DstRequestRateLimitActionListName = in["dst_request_rate_limit_action_list_name"].(string)
		ret.DstRequestRateLimitAction = in["dst_request_rate_limit_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateSslL4MultiPuThresholdDistribution(d []interface{}) edpt.DdosZoneTemplateSslL4MultiPuThresholdDistribution {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSslL4MultiPuThresholdDistribution
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MultiPuThresholdDistributionValue = in["multi_pu_threshold_distribution_value"].(int)
		ret.MultiPuThresholdDistributionDisable = in["multi_pu_threshold_distribution_disable"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateSslL4Renegotiation(d []interface{}) edpt.DdosZoneTemplateSslL4Renegotiation {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSslL4Renegotiation
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NumRenegotiation = in["num_renegotiation"].(int)
		ret.SslL4RenegActionListName = in["ssl_l4_reneg_action_list_name"].(string)
		ret.SslL4RenegAction = in["ssl_l4_reneg_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateSslL4Src(d []interface{}) edpt.DdosZoneTemplateSslL4Src {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSslL4Src
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimit = getObjectDdosZoneTemplateSslL4SrcRateLimit(in["rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateSslL4SrcRateLimit(d []interface{}) edpt.DdosZoneTemplateSslL4SrcRateLimit {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSslL4SrcRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Request = getObjectDdosZoneTemplateSslL4SrcRateLimitRequest(in["request"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateSslL4SrcRateLimitRequest(d []interface{}) edpt.DdosZoneTemplateSslL4SrcRateLimitRequest {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSslL4SrcRateLimitRequest
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcRequestRateLimit = in["src_request_rate_limit"].(int)
		ret.SrcRequestRateLimitActionListName = in["src_request_rate_limit_action_list_name"].(string)
		ret.SrcRequestRateLimitAction = in["src_request_rate_limit_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateSslL4SslTrafficCheck317(d []interface{}) edpt.DdosZoneTemplateSslL4SslTrafficCheck317 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSslL4SslTrafficCheck317
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HeaderInspection = in["header_inspection"].(int)
		ret.HeaderAction = in["header_action"].(string)
		ret.CheckResumedConnection = in["check_resumed_connection"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosZoneTemplateSslL4(d *schema.ResourceData) edpt.DdosZoneTemplateSslL4 {
	var ret edpt.DdosZoneTemplateSslL4
	ret.Inst.AllowNonTls = d.Get("allow_non_tls").(int)
	ret.Inst.AuthHandshake = getObjectDdosZoneTemplateSslL4AuthHandshake(d.Get("auth_handshake").([]interface{}))
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.Dst = getObjectDdosZoneTemplateSslL4Dst(d.Get("dst").([]interface{}))
	ret.Inst.MultiPuThresholdDistribution = getObjectDdosZoneTemplateSslL4MultiPuThresholdDistribution(d.Get("multi_pu_threshold_distribution").([]interface{}))
	ret.Inst.Renegotiation = getObjectDdosZoneTemplateSslL4Renegotiation(d.Get("renegotiation").([]interface{}))
	ret.Inst.Src = getObjectDdosZoneTemplateSslL4Src(d.Get("src").([]interface{}))
	ret.Inst.SslL4TmplName = d.Get("ssl_l4_tmpl_name").(string)
	ret.Inst.SslTrafficCheck = getObjectDdosZoneTemplateSslL4SslTrafficCheck317(d.Get("ssl_traffic_check").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
