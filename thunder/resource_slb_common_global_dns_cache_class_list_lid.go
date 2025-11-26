package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCommonGlobalDnsCacheClassListLid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_common_global_dns_cache_class_list_lid`: Limit ID\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbCommonGlobalDnsCacheClassListLidCreate,
		UpdateContext: resourceSlbCommonGlobalDnsCacheClassListLidUpdate,
		ReadContext:   resourceSlbCommonGlobalDnsCacheClassListLidRead,
		DeleteContext: resourceSlbCommonGlobalDnsCacheClassListLidDelete,

		Schema: map[string]*schema.Schema{
			"conn_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Connection rate limit",
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
			"lidnum": {
				Type: schema.TypeInt, Required: true, Description: "Specify a limit ID",
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
			"over_limit_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'ignore': Ignore the limit and proceed; 'drop': Drop the query when it exceeds limit;",
			},
			"per": {
				Type: schema.TypeInt, Optional: true, Description: "Per (Number of 100ms)",
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
func resourceSlbCommonGlobalDnsCacheClassListLidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonGlobalDnsCacheClassListLidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonGlobalDnsCacheClassListLid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonGlobalDnsCacheClassListLidRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbCommonGlobalDnsCacheClassListLidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonGlobalDnsCacheClassListLidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonGlobalDnsCacheClassListLid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonGlobalDnsCacheClassListLidRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbCommonGlobalDnsCacheClassListLidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonGlobalDnsCacheClassListLidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonGlobalDnsCacheClassListLid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbCommonGlobalDnsCacheClassListLidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonGlobalDnsCacheClassListLidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonGlobalDnsCacheClassListLid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbCommonGlobalDnsCacheClassListLidDns(d []interface{}) edpt.SlbCommonGlobalDnsCacheClassListLidDns {

	count1 := len(d)
	var ret edpt.SlbCommonGlobalDnsCacheClassListLidDns
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheAction = in["cache_action"].(string)
		ret.Ttl = in["ttl"].(int)
		ret.Weight = in["weight"].(int)
		ret.HonorServerResponseTtl = in["honor_server_response_ttl"].(int)
	}
	return ret
}

func dataToEndpointSlbCommonGlobalDnsCacheClassListLid(d *schema.ResourceData) edpt.SlbCommonGlobalDnsCacheClassListLid {
	var ret edpt.SlbCommonGlobalDnsCacheClassListLid
	ret.Inst.ConnRateLimit = d.Get("conn_rate_limit").(int)
	ret.Inst.Dns = getObjectSlbCommonGlobalDnsCacheClassListLidDns(d.Get("dns").([]interface{}))
	ret.Inst.Lidnum = d.Get("lidnum").(int)
	ret.Inst.Lockout = d.Get("lockout").(int)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.LogInterval = d.Get("log_interval").(int)
	ret.Inst.OverLimitAction = d.Get("over_limit_action").(string)
	ret.Inst.Per = d.Get("per").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
