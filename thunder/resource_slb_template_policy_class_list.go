package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_policy_class_list`: Configure classification list\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePolicyClassListCreate,
		UpdateContext: resourceSlbTemplatePolicyClassListUpdate,
		ReadContext:   resourceSlbTemplatePolicyClassListRead,
		DeleteContext: resourceSlbTemplatePolicyClassListDelete,

		Schema: map[string]*schema.Schema{
			"client_ip_l3_dest": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use destination IP as client IP address",
			},
			"client_ip_l7_header": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use extract client IP address from L7 header",
			},
			"header_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify L7 header name",
			},
			"lid_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lidnum": {
							Type: schema.TypeInt, Required: true, Description: "Specify a limit ID",
						},
						"conn_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Connection limit",
						},
						"conn_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Specify connection rate limit",
						},
						"conn_per": {
							Type: schema.TypeInt, Optional: true, Description: "Per (Specify interval in number of 100ms)",
						},
						"request_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Request limit (Specify request limit)",
						},
						"request_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Request rate limit (Specify request rate limit)",
						},
						"request_per": {
							Type: schema.TypeInt, Optional: true, Description: "Per (Specify interval in number of 100ms)",
						},
						"bw_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Specify bandwidth rate limit (Bandwidth rate limit in bytes)",
						},
						"bw_per": {
							Type: schema.TypeInt, Optional: true, Description: "Per (Specify interval in number of 100ms)",
						},
						"over_limit_action": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set action when exceeds limit",
						},
						"action_value": {
							Type: schema.TypeString, Optional: true, Description: "'forward': Forward the traffic even it exceeds limit; 'reset': Reset the connection when it exceeds limit;",
						},
						"lockout": {
							Type: schema.TypeInt, Optional: true, Description: "Don't accept any new connection for certain time (Lockout duration in minutes)",
						},
						"log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log a message",
						},
						"interval": {
							Type: schema.TypeInt, Optional: true, Description: "Specify log interval in minutes, by default system will log every over limit instance",
						},
						"direct_action": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set action when match the lid",
						},
						"direct_service_group": {
							Type: schema.TypeString, Optional: true, Description: "Specify a service group (Specify the service group name)",
						},
						"direct_pbslb_logging": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure PBSLB logging",
						},
						"direct_pbslb_interval": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify logging interval in minutes(default is 3)",
						},
						"direct_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only log unsuccessful connections",
						},
						"direct_action_value": {
							Type: schema.TypeString, Optional: true, Description: "'drop': drop the packet; 'reset': Send reset back;",
						},
						"direct_logging_drp_rst": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure PBSLB logging",
						},
						"direct_action_interval": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify logging interval in minute (default is 3)",
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
				Type: schema.TypeString, Required: true, Description: "Class list name or geo-location-class-list name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplatePolicyClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePolicyClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePolicyClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePolicyClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplatePolicyClassListLidList(d []interface{}) []edpt.SlbTemplatePolicyClassListLidList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyClassListLidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyClassListLidList
		oi.Lidnum = in["lidnum"].(int)
		oi.ConnLimit = in["conn_limit"].(int)
		oi.ConnRateLimit = in["conn_rate_limit"].(int)
		oi.ConnPer = in["conn_per"].(int)
		oi.RequestLimit = in["request_limit"].(int)
		oi.RequestRateLimit = in["request_rate_limit"].(int)
		oi.RequestPer = in["request_per"].(int)
		oi.BwRateLimit = in["bw_rate_limit"].(int)
		oi.BwPer = in["bw_per"].(int)
		oi.OverLimitAction = in["over_limit_action"].(int)
		oi.ActionValue = in["action_value"].(string)
		oi.Lockout = in["lockout"].(int)
		oi.Log = in["log"].(int)
		oi.Interval = in["interval"].(int)
		oi.DirectAction = in["direct_action"].(int)
		oi.DirectServiceGroup = in["direct_service_group"].(string)
		oi.DirectPbslbLogging = in["direct_pbslb_logging"].(int)
		oi.DirectPbslbInterval = in["direct_pbslb_interval"].(int)
		oi.DirectFail = in["direct_fail"].(int)
		oi.DirectActionValue = in["direct_action_value"].(string)
		oi.DirectLoggingDrpRst = in["direct_logging_drp_rst"].(int)
		oi.DirectActionInterval = in["direct_action_interval"].(int)
		oi.ResponseCodeRateLimit = getSliceSlbTemplatePolicyClassListLidListResponseCodeRateLimit(in["response_code_rate_limit"].([]interface{}))
		oi.Dns64 = getObjectSlbTemplatePolicyClassListLidListDns64(in["dns64"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyClassListLidListResponseCodeRateLimit(d []interface{}) []edpt.SlbTemplatePolicyClassListLidListResponseCodeRateLimit {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyClassListLidListResponseCodeRateLimit, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyClassListLidListResponseCodeRateLimit
		oi.CodeRangeStart = in["code_range_start"].(int)
		oi.CodeRangeEnd = in["code_range_end"].(int)
		oi.Threshold = in["threshold"].(int)
		oi.Period = in["period"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplatePolicyClassListLidListDns64(d []interface{}) edpt.SlbTemplatePolicyClassListLidListDns64 {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyClassListLidListDns64
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Disable = in["disable"].(int)
		ret.ExclusiveAnswer = in["exclusive_answer"].(int)
		ret.Prefix = in["prefix"].(string)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyClassList(d *schema.ResourceData) edpt.SlbTemplatePolicyClassList {
	var ret edpt.SlbTemplatePolicyClassList
	ret.Inst.ClientIpL3Dest = d.Get("client_ip_l3_dest").(int)
	ret.Inst.ClientIpL7Header = d.Get("client_ip_l7_header").(int)
	ret.Inst.HeaderName = d.Get("header_name").(string)
	ret.Inst.LidList = getSliceSlbTemplatePolicyClassListLidList(d.Get("lid_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	//omit uuid
	return ret
}
