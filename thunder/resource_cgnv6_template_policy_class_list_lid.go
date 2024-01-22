package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TemplatePolicyClassListLid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_template_policy_class_list_lid`: Limit ID\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TemplatePolicyClassListLidCreate,
		UpdateContext: resourceCgnv6TemplatePolicyClassListLidUpdate,
		ReadContext:   resourceCgnv6TemplatePolicyClassListLidRead,
		DeleteContext: resourceCgnv6TemplatePolicyClassListLidDelete,

		Schema: map[string]*schema.Schema{
			"action_value": {
				Type: schema.TypeString, Optional: true, Description: "'forward': Forward the traffic even it exceeds limit; 'reset': Reset the connection when it exceeds limit;",
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
func resourceCgnv6TemplatePolicyClassListLidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplatePolicyClassListLidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplatePolicyClassListLid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplatePolicyClassListLidRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TemplatePolicyClassListLidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplatePolicyClassListLidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplatePolicyClassListLid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplatePolicyClassListLidRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TemplatePolicyClassListLidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplatePolicyClassListLidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplatePolicyClassListLid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TemplatePolicyClassListLidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplatePolicyClassListLidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplatePolicyClassListLid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6TemplatePolicyClassListLidDns64(d []interface{}) edpt.Cgnv6TemplatePolicyClassListLidDns64 {

	count1 := len(d)
	var ret edpt.Cgnv6TemplatePolicyClassListLidDns64
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Disable = in["disable"].(int)
		ret.ExclusiveAnswer = in["exclusive_answer"].(int)
		ret.Prefix = in["prefix"].(string)
	}
	return ret
}

func dataToEndpointCgnv6TemplatePolicyClassListLid(d *schema.ResourceData) edpt.Cgnv6TemplatePolicyClassListLid {
	var ret edpt.Cgnv6TemplatePolicyClassListLid
	ret.Inst.ActionValue = d.Get("action_value").(string)
	ret.Inst.ConnLimit = d.Get("conn_limit").(int)
	ret.Inst.ConnPer = d.Get("conn_per").(int)
	ret.Inst.ConnRateLimit = d.Get("conn_rate_limit").(int)
	ret.Inst.Dns64 = getObjectCgnv6TemplatePolicyClassListLidDns64(d.Get("dns64").([]interface{}))
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.Lidnum = d.Get("lidnum").(int)
	ret.Inst.Lockout = d.Get("lockout").(int)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.OverLimitAction = d.Get("over_limit_action").(int)
	ret.Inst.RequestLimit = d.Get("request_limit").(int)
	ret.Inst.RequestPer = d.Get("request_per").(int)
	ret.Inst.RequestRateLimit = d.Get("request_rate_limit").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
