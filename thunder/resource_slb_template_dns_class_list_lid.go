package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsClassListLid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_class_list_lid`: Limit ID\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsClassListLidCreate,
		UpdateContext: resourceSlbTemplateDnsClassListLidUpdate,
		ReadContext:   resourceSlbTemplateDnsClassListLidRead,
		DeleteContext: resourceSlbTemplateDnsClassListLidDelete,

		Schema: map[string]*schema.Schema{
			"action_value": {
				Type: schema.TypeString, Optional: true, Description: "'dns-cache-disable': Disable DNS cache when it exceeds limit; 'dns-cache-enable': Enable DNS cache when it exceeds limit; 'forward': Forward the traffic even it exceeds limit;",
			},
			"conn_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Connection rate limit",
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
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Action when exceeds limit",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplateDnsClassListLidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsClassListLidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsClassListLid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsClassListLidRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsClassListLidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsClassListLidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsClassListLid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsClassListLidRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsClassListLidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsClassListLidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsClassListLid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsClassListLidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsClassListLidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsClassListLid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbTemplateDnsClassListLidDns(d []interface{}) edpt.SlbTemplateDnsClassListLidDns {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsClassListLidDns
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheAction = in["cache_action"].(string)
		ret.Ttl = in["ttl"].(int)
		ret.Weight = in["weight"].(int)
		ret.HonorServerResponseTtl = in["honor_server_response_ttl"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsClassListLid(d *schema.ResourceData) edpt.SlbTemplateDnsClassListLid {
	var ret edpt.SlbTemplateDnsClassListLid
	ret.Inst.ActionValue = d.Get("action_value").(string)
	ret.Inst.ConnRateLimit = d.Get("conn_rate_limit").(int)
	ret.Inst.Dns = getObjectSlbTemplateDnsClassListLidDns(d.Get("dns").([]interface{}))
	ret.Inst.Lidnum = d.Get("lidnum").(int)
	ret.Inst.Lockout = d.Get("lockout").(int)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.LogInterval = d.Get("log_interval").(int)
	ret.Inst.OverLimitAction = d.Get("over_limit_action").(int)
	ret.Inst.Per = d.Get("per").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
