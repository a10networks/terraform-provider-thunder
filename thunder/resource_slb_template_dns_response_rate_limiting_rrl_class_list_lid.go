package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsResponseRateLimitingRrlClassListLid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_response_rate_limiting_rrl_class_list_lid`: DNS Response Rate Limiting Class-List Lid\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidCreate,
		UpdateContext: resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidUpdate,
		ReadContext:   resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidRead,
		DeleteContext: resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidDelete,

		Schema: map[string]*schema.Schema{
			"lid_action": {
				Type: schema.TypeString, Optional: true, Default: "rate-limit", Description: "'log-only': Only log rate-limiting, do not actually rate limit. Requires enable-log configuration; 'rate-limit': Rate-Limit based on configuration (Default); 'whitelist': Whitelist, disable rate-limiting;",
			},
			"lid_enable_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"lid_match_subnet": {
				Type: schema.TypeString, Optional: true, Default: "255.255.255.255", Description: "IP subnet mask (response rate by IP subnet mask)",
			},
			"lid_match_subnet_v6": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "IPV6 subnet mask (response rate by IPv6 subnet mask)",
			},
			"lid_response_rate": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Responses exceeding this rate within the window will be dropped (default 5 per second)",
			},
			"lid_slip_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Every n'th response that would be rate-limited will be let through instead",
			},
			"lid_src_ip_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"lid_tc_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Every n'th response that would be rate-limited will respond with TC bit",
			},
			"lid_window": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Rate-Limiting Interval in Seconds (default is one)",
			},
			"lidnum": {
				Type: schema.TypeInt, Required: true, Description: "Specify a limit ID",
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
func resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsResponseRateLimitingRrlClassListLid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsResponseRateLimitingRrlClassListLid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsResponseRateLimitingRrlClassListLid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsResponseRateLimitingRrlClassListLidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsResponseRateLimitingRrlClassListLid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDnsResponseRateLimitingRrlClassListLid(d *schema.ResourceData) edpt.SlbTemplateDnsResponseRateLimitingRrlClassListLid {
	var ret edpt.SlbTemplateDnsResponseRateLimitingRrlClassListLid
	ret.Inst.LidAction = d.Get("lid_action").(string)
	ret.Inst.LidEnableLog = d.Get("lid_enable_log").(int)
	ret.Inst.LidMatchSubnet = d.Get("lid_match_subnet").(string)
	ret.Inst.LidMatchSubnetV6 = d.Get("lid_match_subnet_v6").(int)
	ret.Inst.LidResponseRate = d.Get("lid_response_rate").(int)
	ret.Inst.LidSlipRate = d.Get("lid_slip_rate").(int)
	ret.Inst.LidSrcIpOnly = d.Get("lid_src_ip_only").(int)
	ret.Inst.LidTcRate = d.Get("lid_tc_rate").(int)
	ret.Inst.LidWindow = d.Get("lid_window").(int)
	ret.Inst.Lidnum = d.Get("lidnum").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
