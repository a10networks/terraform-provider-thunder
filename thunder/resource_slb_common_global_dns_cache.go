package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCommonGlobalDnsCache() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_common_global_dns_cache`: Configure global DNS cache related settings\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbCommonGlobalDnsCacheCreate,
		UpdateContext: resourceSlbCommonGlobalDnsCacheUpdate,
		ReadContext:   resourceSlbCommonGlobalDnsCacheRead,
		DeleteContext: resourceSlbCommonGlobalDnsCacheDelete,

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
										Type: schema.TypeString, Optional: true, Default: "drop", Description: "'ignore': Ignore the limit and proceed; 'drop': Drop the query when it exceeds limit;",
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
													Type: schema.TypeString, Optional: true, Default: "cache-enable", Description: "'cache-disable': Disable dns cache; 'cache-enable': Enable dns cache;",
												},
												"ttl": {
													Type: schema.TypeInt, Optional: true, Default: 300, Description: "TTL for cache entry (TTL in seconds)",
												},
												"weight": {
													Type: schema.TypeInt, Optional: true, Default: 1, Description: "Weight for cache entry",
												},
												"honor_server_response_ttl": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Honor the server response TTL",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbCommonGlobalDnsCacheCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonGlobalDnsCacheCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonGlobalDnsCache(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonGlobalDnsCacheRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbCommonGlobalDnsCacheUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonGlobalDnsCacheUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonGlobalDnsCache(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonGlobalDnsCacheRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbCommonGlobalDnsCacheDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonGlobalDnsCacheDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonGlobalDnsCache(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbCommonGlobalDnsCacheRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonGlobalDnsCacheRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonGlobalDnsCache(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbCommonGlobalDnsCacheClassList1499(d []interface{}) edpt.SlbCommonGlobalDnsCacheClassList1499 {

	count1 := len(d)
	var ret edpt.SlbCommonGlobalDnsCacheClassList1499
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		//omit uuid
		ret.LidList = getSliceSlbCommonGlobalDnsCacheClassListLidList1500(in["lid_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbCommonGlobalDnsCacheClassListLidList1500(d []interface{}) []edpt.SlbCommonGlobalDnsCacheClassListLidList1500 {

	count1 := len(d)
	ret := make([]edpt.SlbCommonGlobalDnsCacheClassListLidList1500, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbCommonGlobalDnsCacheClassListLidList1500
		oi.Lidnum = in["lidnum"].(int)
		oi.ConnRateLimit = in["conn_rate_limit"].(int)
		oi.Per = in["per"].(int)
		oi.OverLimitAction = in["over_limit_action"].(string)
		oi.Lockout = in["lockout"].(int)
		oi.Log = in["log"].(int)
		oi.LogInterval = in["log_interval"].(int)
		oi.Dns = getObjectSlbCommonGlobalDnsCacheClassListLidListDns1501(in["dns"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbCommonGlobalDnsCacheClassListLidListDns1501(d []interface{}) edpt.SlbCommonGlobalDnsCacheClassListLidListDns1501 {

	count1 := len(d)
	var ret edpt.SlbCommonGlobalDnsCacheClassListLidListDns1501
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheAction = in["cache_action"].(string)
		ret.Ttl = in["ttl"].(int)
		ret.Weight = in["weight"].(int)
		ret.HonorServerResponseTtl = in["honor_server_response_ttl"].(int)
	}
	return ret
}

func dataToEndpointSlbCommonGlobalDnsCache(d *schema.ResourceData) edpt.SlbCommonGlobalDnsCache {
	var ret edpt.SlbCommonGlobalDnsCache
	ret.Inst.ClassList = getObjectSlbCommonGlobalDnsCacheClassList1499(d.Get("class_list").([]interface{}))
	//omit uuid
	return ret
}
