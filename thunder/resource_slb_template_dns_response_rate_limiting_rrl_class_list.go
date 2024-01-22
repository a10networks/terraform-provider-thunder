package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsResponseRateLimitingRrlClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_response_rate_limiting_rrl_class_list`: DNS Response Rate Limiting Class-List\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsResponseRateLimitingRrlClassListCreate,
		UpdateContext: resourceSlbTemplateDnsResponseRateLimitingRrlClassListUpdate,
		ReadContext:   resourceSlbTemplateDnsResponseRateLimitingRrlClassListRead,
		DeleteContext: resourceSlbTemplateDnsResponseRateLimitingRrlClassListDelete,

		Schema: map[string]*schema.Schema{
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Class-list name",
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
func resourceSlbTemplateDnsResponseRateLimitingRrlClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsResponseRateLimitingRrlClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsResponseRateLimitingRrlClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsResponseRateLimitingRrlClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsResponseRateLimitingRrlClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsResponseRateLimitingRrlClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsResponseRateLimitingRrlClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsResponseRateLimitingRrlClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsResponseRateLimitingRrlClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsResponseRateLimitingRrlClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsResponseRateLimitingRrlClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsResponseRateLimitingRrlClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsResponseRateLimitingRrlClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsResponseRateLimitingRrlClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDnsResponseRateLimitingRrlClassListLidList(d []interface{}) []edpt.SlbTemplateDnsResponseRateLimitingRrlClassListLidList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsResponseRateLimitingRrlClassListLidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsResponseRateLimitingRrlClassListLidList
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

func dataToEndpointSlbTemplateDnsResponseRateLimitingRrlClassList(d *schema.ResourceData) edpt.SlbTemplateDnsResponseRateLimitingRrlClassList {
	var ret edpt.SlbTemplateDnsResponseRateLimitingRrlClassList
	ret.Inst.LidList = getSliceSlbTemplateDnsResponseRateLimitingRrlClassListLidList(d.Get("lid_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
