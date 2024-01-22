package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyClassListLid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_policy_class_list_lid`: Limit ID\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePolicyClassListLidCreate,
		UpdateContext: resourceSlbTemplatePolicyClassListLidUpdate,
		ReadContext:   resourceSlbTemplatePolicyClassListLidRead,
		DeleteContext: resourceSlbTemplatePolicyClassListLidDelete,

		Schema: map[string]*schema.Schema{
			"action_value": {
				Type: schema.TypeString, Optional: true, Description: "'forward': Forward the traffic even it exceeds limit; 'reset': Reset the connection when it exceeds limit;",
			},
			"bw_per": {
				Type: schema.TypeInt, Optional: true, Description: "Per (Specify interval in number of 100ms)",
			},
			"bw_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Specify bandwidth rate limit (Bandwidth rate limit in bytes)",
			},
			"conn_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Connection limit",
			},
			"conn_per": {
				Type: schema.TypeInt, Optional: true, Description: "Per (Specify interval in number of 100ms)",
			},
			"conn_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Specify connection rate limit",
			},
			"direct_action": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set action when match the lid",
			},
			"direct_action_interval": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify logging interval in minute (default is 3)",
			},
			"direct_action_value": {
				Type: schema.TypeString, Optional: true, Description: "'drop': drop the packet; 'reset': Send reset back;",
			},
			"direct_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only log unsuccessful connections",
			},
			"direct_logging_drp_rst": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure PBSLB logging",
			},
			"direct_pbslb_interval": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify logging interval in minutes(default is 3)",
			},
			"direct_pbslb_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure PBSLB logging",
			},
			"direct_service_group": {
				Type: schema.TypeString, Optional: true, Description: "Specify a service group (Specify the service group name)",
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
			"interval": {
				Type: schema.TypeInt, Optional: true, Description: "Specify log interval in minutes, by default system will log every over limit instance",
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
			"over_limit_action": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set action when exceeds limit",
			},
			"request_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Request limit (Specify request limit)",
			},
			"request_per": {
				Type: schema.TypeInt, Optional: true, Description: "Per (Specify interval in number of 100ms)",
			},
			"request_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Request rate limit (Specify request rate limit)",
			},
			"response_code_rate_limit": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code_range_start": {
							Type: schema.TypeInt, Optional: true, Description: "server response code range start",
						},
						"code_range_end": {
							Type: schema.TypeInt, Optional: true, Description: "server response code range end",
						},
						"threshold": {
							Type: schema.TypeInt, Optional: true, Description: "the times of getting the response code",
						},
						"period": {
							Type: schema.TypeInt, Optional: true, Description: "seconds",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplatePolicyClassListLidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyClassListLidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyClassListLid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyClassListLidRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePolicyClassListLidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyClassListLidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyClassListLid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyClassListLidRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePolicyClassListLidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyClassListLidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyClassListLid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePolicyClassListLidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyClassListLidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyClassListLid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbTemplatePolicyClassListLidDns64(d []interface{}) edpt.SlbTemplatePolicyClassListLidDns64 {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyClassListLidDns64
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Disable = in["disable"].(int)
		ret.ExclusiveAnswer = in["exclusive_answer"].(int)
		ret.Prefix = in["prefix"].(string)
	}
	return ret
}

func getSliceSlbTemplatePolicyClassListLidResponseCodeRateLimit(d []interface{}) []edpt.SlbTemplatePolicyClassListLidResponseCodeRateLimit {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyClassListLidResponseCodeRateLimit, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyClassListLidResponseCodeRateLimit
		oi.CodeRangeStart = in["code_range_start"].(int)
		oi.CodeRangeEnd = in["code_range_end"].(int)
		oi.Threshold = in["threshold"].(int)
		oi.Period = in["period"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyClassListLid(d *schema.ResourceData) edpt.SlbTemplatePolicyClassListLid {
	var ret edpt.SlbTemplatePolicyClassListLid
	ret.Inst.ActionValue = d.Get("action_value").(string)
	ret.Inst.BwPer = d.Get("bw_per").(int)
	ret.Inst.BwRateLimit = d.Get("bw_rate_limit").(int)
	ret.Inst.ConnLimit = d.Get("conn_limit").(int)
	ret.Inst.ConnPer = d.Get("conn_per").(int)
	ret.Inst.ConnRateLimit = d.Get("conn_rate_limit").(int)
	ret.Inst.DirectAction = d.Get("direct_action").(int)
	ret.Inst.DirectActionInterval = d.Get("direct_action_interval").(int)
	ret.Inst.DirectActionValue = d.Get("direct_action_value").(string)
	ret.Inst.DirectFail = d.Get("direct_fail").(int)
	ret.Inst.DirectLoggingDrpRst = d.Get("direct_logging_drp_rst").(int)
	ret.Inst.DirectPbslbInterval = d.Get("direct_pbslb_interval").(int)
	ret.Inst.DirectPbslbLogging = d.Get("direct_pbslb_logging").(int)
	ret.Inst.DirectServiceGroup = d.Get("direct_service_group").(string)
	ret.Inst.Dns64 = getObjectSlbTemplatePolicyClassListLidDns64(d.Get("dns64").([]interface{}))
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.Lidnum = d.Get("lidnum").(int)
	ret.Inst.Lockout = d.Get("lockout").(int)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.OverLimitAction = d.Get("over_limit_action").(int)
	ret.Inst.RequestLimit = d.Get("request_limit").(int)
	ret.Inst.RequestPer = d.Get("request_per").(int)
	ret.Inst.RequestRateLimit = d.Get("request_rate_limit").(int)
	ret.Inst.ResponseCodeRateLimit = getSliceSlbTemplatePolicyClassListLidResponseCodeRateLimit(d.Get("response_code_rate_limit").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
