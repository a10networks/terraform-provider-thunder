package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_class_list`: Classification list\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsClassListCreate,
		UpdateContext: resourceSlbTemplateDnsClassListUpdate,
		ReadContext:   resourceSlbTemplateDnsClassListRead,
		DeleteContext: resourceSlbTemplateDnsClassListDelete,

		Schema: map[string]*schema.Schema{
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
									"honor_server_response_ttl": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Honor the server reponse TTL",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify a class list name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplateDnsClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDnsClassListLidList(d []interface{}) []edpt.SlbTemplateDnsClassListLidList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsClassListLidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsClassListLidList
		oi.Lidnum = in["lidnum"].(int)
		oi.ConnRateLimit = in["conn_rate_limit"].(int)
		oi.Per = in["per"].(int)
		oi.OverLimitAction = in["over_limit_action"].(int)
		oi.ActionValue = in["action_value"].(string)
		oi.Lockout = in["lockout"].(int)
		oi.Log = in["log"].(int)
		oi.LogInterval = in["log_interval"].(int)
		oi.Dns = getObjectSlbTemplateDnsClassListLidListDns(in["dns"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsClassListLidListDns(d []interface{}) edpt.SlbTemplateDnsClassListLidListDns {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsClassListLidListDns
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheAction = in["cache_action"].(string)
		ret.Ttl = in["ttl"].(int)
		ret.Weight = in["weight"].(int)
		ret.HonorServerResponseTtl = in["honor_server_response_ttl"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsClassList(d *schema.ResourceData) edpt.SlbTemplateDnsClassList {
	var ret edpt.SlbTemplateDnsClassList
	ret.Inst.LidList = getSliceSlbTemplateDnsClassListLidList(d.Get("lid_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	//omit uuid
	return ret
}
