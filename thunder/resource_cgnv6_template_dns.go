package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TemplateDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_template_dns`: DNS template\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TemplateDnsCreate,
		UpdateContext: resourceCgnv6TemplateDnsUpdate,
		ReadContext:   resourceCgnv6TemplateDnsRead,
		DeleteContext: resourceCgnv6TemplateDnsDelete,

		Schema: map[string]*schema.Schema{
			"class_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Specify a class list name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"lid_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lidnum": {
										Type: schema.TypeInt, Required: true, Description: "Specify a limit ID",
									},
									"conn_rate_limit": {
										Type: schema.TypeInt, Optional: true, Description: "Connection rate limit",
									},
									"per": {
										Type: schema.TypeInt, Optional: true, Description: "Per (Number of 100ms)",
									},
									"over_limit_action": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Action when exceeds limit",
									},
									"action_value": {
										Type: schema.TypeString, Optional: true, Description: "'dns-cache-disable': Disable DNS cache when it exceeds limit; 'dns-cache-enable': Enable DNS cache when it exceeds limit; 'forward': Forward the traffic even it exceeds limit;",
									},
									"lockout": {
										Type: schema.TypeInt, Optional: true, Description: "Don't accept any new connection for certain time (Lockout duration in minutes)",
									},
									"log": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log a message",
									},
									"log_interval": {
										Type: schema.TypeInt, Optional: true, Description: "Log interval (minute, by default system will log every over limit instance)",
									},
									"dns": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cache_action": {
													Type: schema.TypeString, Optional: true, Default: "cache-disable", Description: "'cache-disable': Disable dns cache; 'cache-enable': Enable dns cache;",
												},
												"ttl": {
													Type: schema.TypeInt, Optional: true, Description: "TTL for cache entry (TTL in seconds)",
												},
												"weight": {
													Type: schema.TypeInt, Optional: true, Description: "Weight for cache entry",
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
					},
				},
			},
			"default_policy": {
				Type: schema.TypeString, Optional: true, Default: "nocache", Description: "'nocache': Cache disable; 'cache': Cache enable;",
			},
			"disable_dns_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable DNS template",
			},
			"dns64": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS64 (Need to config this option before config any other dns64 options)",
						},
						"answer_only_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Only translate the Answer Section",
						},
						"auth_data": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set AA flag in DNS Response",
						},
						"cache": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use a cached A-query response to provide AAAA query responses for the same hostname",
						},
						"change_query": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always change incoming AAAA DNS Query to A",
						},
						"compress_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Always try DNS Compression",
						},
						"deep_check_qr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Check DNS Question Record",
						},
						"deep_check_rr_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Check DNS Response Records",
						},
						"drop_cname_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Drop DNS CNAME Response",
						},
						"edns_append": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Append EDNS Record when send A Query to server",
						},
						"fast_append": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Append translated Records when original Response only has Answer Section",
						},
						"ignore_rcode3_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Ignore DNS error Response with rcode 3",
						},
						"max_qr_length": {
							Type: schema.TypeInt, Optional: true, Default: 128, Description: "Max Question Record Length, default is 128",
						},
						"parallel_query": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward AAAA Query & generate A Query in parallel",
						},
						"passive_query_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Generate A query upon empty or error Response",
						},
						"retry": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Retry count, default is 3 (Retry Number)",
						},
						"single_response_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Single Response which is used to avoid ambiguity",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Timeout to send additional Queries, unit: second, default is 1",
						},
						"trans_ptr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Translate DNS PTR Records",
						},
						"trans_ptr_query": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Translate DNS PTR Query",
						},
						"ttl": {
							Type: schema.TypeInt, Optional: true, Description: "Specify Max TTL in DNS Response, unit: second",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop the malformed query",
			},
			"forward": {
				Type: schema.TypeString, Optional: true, Description: "Forward to service group (Service group name)",
			},
			"max_cache_size": {
				Type: schema.TypeInt, Optional: true, Description: "Define maximum cache size (Maximum cache entry per VIP)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DNS Template Name",
			},
			"period": {
				Type: schema.TypeInt, Optional: true, Description: "Period in minutes",
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
func resourceCgnv6TemplateDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateDnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateDns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateDnsRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TemplateDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateDnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateDns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateDnsRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TemplateDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateDnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateDns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TemplateDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateDnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateDns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6TemplateDnsClassList111(d []interface{}) edpt.Cgnv6TemplateDnsClassList111 {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateDnsClassList111
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		//omit uuid
		ret.LidList = getSliceCgnv6TemplateDnsClassListLidList112(in["lid_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6TemplateDnsClassListLidList112(d []interface{}) []edpt.Cgnv6TemplateDnsClassListLidList112 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateDnsClassListLidList112, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateDnsClassListLidList112
		oi.Lidnum = in["lidnum"].(int)
		oi.ConnRateLimit = in["conn_rate_limit"].(int)
		oi.Per = in["per"].(int)
		oi.OverLimitAction = in["over_limit_action"].(int)
		oi.ActionValue = in["action_value"].(string)
		oi.Lockout = in["lockout"].(int)
		oi.Log = in["log"].(int)
		oi.LogInterval = in["log_interval"].(int)
		oi.Dns = getObjectCgnv6TemplateDnsClassListLidListDns113(in["dns"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6TemplateDnsClassListLidListDns113(d []interface{}) edpt.Cgnv6TemplateDnsClassListLidListDns113 {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateDnsClassListLidListDns113
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheAction = in["cache_action"].(string)
		ret.Ttl = in["ttl"].(int)
		ret.Weight = in["weight"].(int)
	}
	return ret
}

func getObjectCgnv6TemplateDnsDns64114(d []interface{}) edpt.Cgnv6TemplateDnsDns64114 {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateDnsDns64114
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.AnswerOnlyDisable = in["answer_only_disable"].(int)
		ret.AuthData = in["auth_data"].(int)
		ret.Cache = in["cache"].(int)
		ret.ChangeQuery = in["change_query"].(int)
		ret.CompressDisable = in["compress_disable"].(int)
		ret.DeepCheckQr = in["deep_check_qr"].(int)
		ret.DeepCheckRrDisable = in["deep_check_rr_disable"].(int)
		ret.DropCnameDisable = in["drop_cname_disable"].(int)
		ret.EdnsAppend = in["edns_append"].(int)
		ret.FastAppend = in["fast_append"].(int)
		ret.IgnoreRcode3Disable = in["ignore_rcode3_disable"].(int)
		ret.MaxQrLength = in["max_qr_length"].(int)
		ret.ParallelQuery = in["parallel_query"].(int)
		ret.PassiveQueryDisable = in["passive_query_disable"].(int)
		ret.Retry = in["retry"].(int)
		ret.SingleResponseDisable = in["single_response_disable"].(int)
		ret.Timeout = in["timeout"].(int)
		ret.TransPtr = in["trans_ptr"].(int)
		ret.TransPtrQuery = in["trans_ptr_query"].(int)
		ret.Ttl = in["ttl"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointCgnv6TemplateDns(d *schema.ResourceData) edpt.Cgnv6TemplateDns {
	var ret edpt.Cgnv6TemplateDns
	ret.Inst.ClassList = getObjectCgnv6TemplateDnsClassList111(d.Get("class_list").([]interface{}))
	ret.Inst.DefaultPolicy = d.Get("default_policy").(string)
	ret.Inst.DisableDnsTemplate = d.Get("disable_dns_template").(int)
	ret.Inst.Dns64 = getObjectCgnv6TemplateDnsDns64114(d.Get("dns64").([]interface{}))
	ret.Inst.Drop = d.Get("drop").(int)
	ret.Inst.Forward = d.Get("forward").(string)
	ret.Inst.MaxCacheSize = d.Get("max_cache_size").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
