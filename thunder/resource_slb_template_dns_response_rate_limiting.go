package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsResponseRateLimiting() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_response_rate_limiting`: DNS Response Rate Limiting\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsResponseRateLimitingCreate,
		UpdateContext: resourceSlbTemplateDnsResponseRateLimitingUpdate,
		ReadContext:   resourceSlbTemplateDnsResponseRateLimitingRead,
		DeleteContext: resourceSlbTemplateDnsResponseRateLimitingDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "rate-limit", Description: "'log-only': Only log rate-limiting, do not actually rate limit. Requires enable-log configuration; 'rate-limit': Rate-Limit based on configuration (Default); 'whitelist': Whitelist, disable rate-limiting;",
			},
			"enable_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"filter_response_rate": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Maximum allowed request rate for the filter. This should match average traffic. (default 10 per seconds)",
			},
			"match_subnet": {
				Type: schema.TypeString, Optional: true, Default: "255.255.255.255", Description: "IP subnet mask (response rate by IP subnet mask)",
			},
			"match_subnet_v6": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "IPV6 subnet mask (response rate by IPv6 subnet mask)",
			},
			"response_rate": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Responses exceeding this rate within the window will be dropped (default 5 per second)",
			},
			"rrl_class_list_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Class-list name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"lid_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lidnum": {
										Type: schema.TypeInt, Required: true, Description: "Specify a limit ID",
									},
									"lid_response_rate": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Responses exceeding this rate within the window will be dropped (default 5 per second)",
									},
									"lid_slip_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Every n'th response that would be rate-limited will be let through instead",
									},
									"lid_tc_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Every n'th response that would be rate-limited will respond with TC bit",
									},
									"lid_match_subnet": {
										Type: schema.TypeString, Optional: true, Default: "255.255.255.255", Description: "IP subnet mask (response rate by IP subnet mask)",
									},
									"lid_match_subnet_v6": {
										Type: schema.TypeInt, Optional: true, Default: 128, Description: "IPV6 subnet mask (response rate by IPv6 subnet mask)",
									},
									"lid_window": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Rate-Limiting Interval in Seconds (default is one)",
									},
									"lid_src_ip_only": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
									},
									"lid_enable_log": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
									},
									"lid_action": {
										Type: schema.TypeString, Optional: true, Default: "rate-limit", Description: "'log-only': Only log rate-limiting, do not actually rate limit. Requires enable-log configuration; 'rate-limit': Rate-Limit based on configuration (Default); 'whitelist': Whitelist, disable rate-limiting;",
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
			"slip_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Every n'th response that would be rate-limited will be let through instead",
			},
			"src_ip_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"tc_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Every n'th response that would be rate-limited will respond with TC bit",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"window": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Rate-Limiting Interval in Seconds (default is one)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplateDnsResponseRateLimitingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsResponseRateLimitingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsResponseRateLimiting(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsResponseRateLimitingRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsResponseRateLimitingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsResponseRateLimitingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsResponseRateLimiting(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsResponseRateLimitingRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsResponseRateLimitingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsResponseRateLimitingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsResponseRateLimiting(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsResponseRateLimitingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsResponseRateLimitingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsResponseRateLimiting(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDnsResponseRateLimitingRrlClassListList(d []interface{}) []edpt.SlbTemplateDnsResponseRateLimitingRrlClassListList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsResponseRateLimitingRrlClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsResponseRateLimitingRrlClassListList
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.LidList = getSliceSlbTemplateDnsResponseRateLimitingRrlClassListListLidList(in["lid_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateDnsResponseRateLimitingRrlClassListListLidList(d []interface{}) []edpt.SlbTemplateDnsResponseRateLimitingRrlClassListListLidList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsResponseRateLimitingRrlClassListListLidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsResponseRateLimitingRrlClassListListLidList
		oi.Lidnum = in["lidnum"].(int)
		oi.LidResponseRate = in["lid_response_rate"].(int)
		oi.LidSlipRate = in["lid_slip_rate"].(int)
		oi.LidTcRate = in["lid_tc_rate"].(int)
		oi.LidMatchSubnet = in["lid_match_subnet"].(string)
		oi.LidMatchSubnetV6 = in["lid_match_subnet_v6"].(int)
		oi.LidWindow = in["lid_window"].(int)
		oi.LidSrcIpOnly = in["lid_src_ip_only"].(int)
		oi.LidEnableLog = in["lid_enable_log"].(int)
		oi.LidAction = in["lid_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsResponseRateLimiting(d *schema.ResourceData) edpt.SlbTemplateDnsResponseRateLimiting {
	var ret edpt.SlbTemplateDnsResponseRateLimiting
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.EnableLog = d.Get("enable_log").(int)
	ret.Inst.FilterResponseRate = d.Get("filter_response_rate").(int)
	ret.Inst.MatchSubnet = d.Get("match_subnet").(string)
	ret.Inst.MatchSubnetV6 = d.Get("match_subnet_v6").(int)
	ret.Inst.ResponseRate = d.Get("response_rate").(int)
	ret.Inst.RrlClassListList = getSliceSlbTemplateDnsResponseRateLimitingRrlClassListList(d.Get("rrl_class_list_list").([]interface{}))
	ret.Inst.SlipRate = d.Get("slip_rate").(int)
	ret.Inst.SrcIpOnly = d.Get("src_ip_only").(int)
	ret.Inst.TcRate = d.Get("tc_rate").(int)
	//omit uuid
	ret.Inst.Window = d.Get("window").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
