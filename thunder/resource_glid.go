package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_glid`: Configure global limit ID\n\n__PLACEHOLDER__",
		CreateContext: resourceGlidCreate,
		UpdateContext: resourceGlidUpdate,
		ReadContext:   resourceGlidRead,
		DeleteContext: resourceGlidDelete,

		Schema: map[string]*schema.Schema{
			"bit_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Kibit (kibibit / 1024-bit) rate limit per rate-interval",
			},
			"conn_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Connection Limit for the GLID (PBSLB range 1-1048575)",
			},
			"conn_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Connection rate limit per rate-interval (TPS range 1-16000000)",
			},
			"conn_rate_limit_interval": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Description for glid",
			},
			"dns": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Default: "cache-disable", Description: "'cache-disable': Disable dns cache; 'cache-enable': Enable dns cache;",
						},
						"weight": {
							Type: schema.TypeInt, Optional: true, Description: "Weight for cache entry",
						},
						"ttl": {
							Type: schema.TypeInt, Optional: true, Description: "TTL for cache entry (TTL in seconds)",
						},
					},
				},
			},
			"dns64": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable",
						},
						"exclusive_answer": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Exclusive Answer in DNS Response",
						},
						"prefix": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 prefix",
						},
					},
				},
			},
			"frag_pkt_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Fragmented packet rate limit per rate-interval",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Global Limit ID Name (PBSLB allows number only)",
			},
			"over_limit_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"over_limit_action": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Action to take when limit(s) exceeds",
						},
						"action_type": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Silently Drop the new connection / new packet when it exceeds limit; 'blacklist-src': Black List source entry for X minutes (only applied to src and src-dst entries);",
						},
						"blacklist_src_min": {
							Type: schema.TypeInt, Optional: true, Description: "Black List source entry for X minutes",
						},
						"action_value": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Silently Drop the new connection / new packet when it exceeds limit; 'dns-cache-disable': Disable dns cache when it exceeds limit; 'dns-cache-enable': Enable dns cache when it exceeds limit; 'forward': Forward the traffic even it exceeds limit; 'reset': Reset the connection when it exceeds limit;",
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
					},
				},
			},
			"pkt_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Packet rate limit per rate-interval",
			},
			"rate_unit": {
				Type: schema.TypeString, Optional: true, Default: "system-global-setting", Description: "'1sec': 1sec for internal glid rate unit; 'system-global-setting': use global rate interval;",
			},
			"request_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Request limit",
			},
			"request_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Request rate limit",
			},
			"request_rate_limit_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Number of 100ms",
			},
			"syn_cookie_thr": {
				Type: schema.TypeInt, Optional: true, Description: "Syn Cookie threshold for the GLID",
			},
			"use_nat_pool": {
				Type: schema.TypeString, Optional: true, Description: "Use NAT pool specified to do reverse NAT for class list members bound to the lid",
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
func resourceGlidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlidRead(ctx, d, meta)
	}
	return diags
}

func resourceGlidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlidRead(ctx, d, meta)
	}
	return diags
}
func resourceGlidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGlidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectGlidDns(d []interface{}) edpt.GlidDns {

	count1 := len(d)
	var ret edpt.GlidDns
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.Weight = in["weight"].(int)
		ret.Ttl = in["ttl"].(int)
	}
	return ret
}

func getObjectGlidDns64(d []interface{}) edpt.GlidDns64 {

	count1 := len(d)
	var ret edpt.GlidDns64
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Disable = in["disable"].(int)
		ret.ExclusiveAnswer = in["exclusive_answer"].(int)
		ret.Prefix = in["prefix"].(string)
	}
	return ret
}

func getObjectGlidOverLimitCfg(d []interface{}) edpt.GlidOverLimitCfg {

	count1 := len(d)
	var ret edpt.GlidOverLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OverLimitAction = in["over_limit_action"].(int)
		ret.ActionType = in["action_type"].(string)
		ret.BlacklistSrcMin = in["blacklist_src_min"].(int)
		ret.ActionValue = in["action_value"].(string)
		ret.Lockout = in["lockout"].(int)
		ret.Log = in["log"].(int)
		ret.LogInterval = in["log_interval"].(int)
	}
	return ret
}

func dataToEndpointGlid(d *schema.ResourceData) edpt.Glid {
	var ret edpt.Glid
	ret.Inst.BitRateLimit = d.Get("bit_rate_limit").(int)
	ret.Inst.ConnLimit = d.Get("conn_limit").(int)
	ret.Inst.ConnRateLimit = d.Get("conn_rate_limit").(int)
	ret.Inst.ConnRateLimitInterval = d.Get("conn_rate_limit_interval").(int)
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.Dns = getObjectGlidDns(d.Get("dns").([]interface{}))
	ret.Inst.Dns64 = getObjectGlidDns64(d.Get("dns64").([]interface{}))
	ret.Inst.FragPktRateLimit = d.Get("frag_pkt_rate_limit").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OverLimitCfg = getObjectGlidOverLimitCfg(d.Get("over_limit_cfg").([]interface{}))
	ret.Inst.PktRateLimit = d.Get("pkt_rate_limit").(int)
	ret.Inst.RateUnit = d.Get("rate_unit").(string)
	ret.Inst.RequestLimit = d.Get("request_limit").(int)
	ret.Inst.RequestRateLimit = d.Get("request_rate_limit").(int)
	ret.Inst.RequestRateLimitInterval = d.Get("request_rate_limit_interval").(int)
	ret.Inst.SynCookieThr = d.Get("syn_cookie_thr").(int)
	ret.Inst.UseNatPool = d.Get("use_nat_pool").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
