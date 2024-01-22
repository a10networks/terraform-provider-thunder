package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TemplatePolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_template_policy`: Policy config\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TemplatePolicyCreate,
		UpdateContext: resourceCgnv6TemplatePolicyUpdate,
		ReadContext:   resourceCgnv6TemplatePolicyRead,
		DeleteContext: resourceCgnv6TemplatePolicyDelete,

		Schema: map[string]*schema.Schema{
			"class_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Class list name",
						},
						"client_ip_l3_dest": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use destination IP as client IP address",
						},
						"client_ip_l7_header": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use extract client IP address from L7 header",
						},
						"header_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify L7 header name",
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
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Policy template name",
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
func resourceCgnv6TemplatePolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplatePolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplatePolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplatePolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TemplatePolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplatePolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplatePolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplatePolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TemplatePolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplatePolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplatePolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TemplatePolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplatePolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplatePolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6TemplatePolicyClassList125(d []interface{}) edpt.Cgnv6TemplatePolicyClassList125 {

	count1 := len(d)
	var ret edpt.Cgnv6TemplatePolicyClassList125
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.ClientIpL3Dest = in["client_ip_l3_dest"].(int)
		ret.ClientIpL7Header = in["client_ip_l7_header"].(int)
		ret.HeaderName = in["header_name"].(string)
		//omit uuid
		ret.LidList = getSliceCgnv6TemplatePolicyClassListLidList126(in["lid_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6TemplatePolicyClassListLidList126(d []interface{}) []edpt.Cgnv6TemplatePolicyClassListLidList126 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplatePolicyClassListLidList126, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplatePolicyClassListLidList126
		oi.Lidnum = in["lidnum"].(int)
		oi.ConnLimit = in["conn_limit"].(int)
		oi.ConnRateLimit = in["conn_rate_limit"].(int)
		oi.ConnPer = in["conn_per"].(int)
		oi.RequestLimit = in["request_limit"].(int)
		oi.RequestRateLimit = in["request_rate_limit"].(int)
		oi.RequestPer = in["request_per"].(int)
		oi.OverLimitAction = in["over_limit_action"].(int)
		oi.ActionValue = in["action_value"].(string)
		oi.Lockout = in["lockout"].(int)
		oi.Log = in["log"].(int)
		oi.Interval = in["interval"].(int)
		oi.Dns64 = getObjectCgnv6TemplatePolicyClassListLidListDns64127(in["dns64"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6TemplatePolicyClassListLidListDns64127(d []interface{}) edpt.Cgnv6TemplatePolicyClassListLidListDns64127 {

	count1 := len(d)
	var ret edpt.Cgnv6TemplatePolicyClassListLidListDns64127
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Disable = in["disable"].(int)
		ret.ExclusiveAnswer = in["exclusive_answer"].(int)
		ret.Prefix = in["prefix"].(string)
	}
	return ret
}

func dataToEndpointCgnv6TemplatePolicy(d *schema.ResourceData) edpt.Cgnv6TemplatePolicy {
	var ret edpt.Cgnv6TemplatePolicy
	ret.Inst.ClassList = getObjectCgnv6TemplatePolicyClassList125(d.Get("class_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
