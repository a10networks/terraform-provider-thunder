package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsLogParameterization() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_log_parameterization`: Harmony controller log parameterization configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsLogParameterizationCreate,
		UpdateContext: resourceAcosEventsLogParameterizationUpdate,
		ReadContext:   resourceAcosEventsLogParameterizationRead,
		DeleteContext: resourceAcosEventsLogParameterizationDelete,

		Schema: map[string]*schema.Schema{
			"log_rate": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Max number of parameterized logs sent per second",
			},
			"message_selector": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"rule_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"index": {
										Type: schema.TypeInt, Required: true, Description: "Specify rule index - rules are applied in numeric order",
									},
									"action": {
										Type: schema.TypeString, Optional: true, Default: "send", Description: "'send': log messages selected by this rule will be sent (Default); 'drop': log messages selected by this rule will be dropped;",
									},
									"message_id": {
										Type: schema.TypeString, Optional: true, Description: "Select a specific message by message-id and optionally severity",
									},
									"message_id_scope": {
										Type: schema.TypeString, Optional: true, Description: "'all': Log messages at this level and all sub-trees; 'node-only': Log messages at this node only; 'children-only': Log messages at all sub-trees only; 'log-field-only': Log message for this Log Field only;",
									},
									"severity_oper": {
										Type: schema.TypeString, Optional: true, Description: "'equal-and-higher': emergency is highest, debugging lowest; 'equal': single severity;",
									},
									"severity_val": {
										Type: schema.TypeString, Optional: true, Description: "'emergency': System unusable log messages (Most Important); 'alert': Action must be taken immediately; 'critical': Critical conditions; 'error': Error conditions; 'warning': Warning conditions; 'notification': Normal but significant conditions; 'information': Informational messages; 'debugging': Debug level messages (Least Important);",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAcosEventsLogParameterizationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogParameterizationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogParameterization(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLogParameterizationRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsLogParameterizationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogParameterizationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogParameterization(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLogParameterizationRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsLogParameterizationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogParameterizationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogParameterization(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsLogParameterizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogParameterizationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogParameterization(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAcosEventsLogParameterizationMessageSelector54(d []interface{}) edpt.AcosEventsLogParameterizationMessageSelector54 {

	count1 := len(d)
	var ret edpt.AcosEventsLogParameterizationMessageSelector54
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.RuleList = getSliceAcosEventsLogParameterizationMessageSelectorRuleList55(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceAcosEventsLogParameterizationMessageSelectorRuleList55(d []interface{}) []edpt.AcosEventsLogParameterizationMessageSelectorRuleList55 {

	count1 := len(d)
	ret := make([]edpt.AcosEventsLogParameterizationMessageSelectorRuleList55, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsLogParameterizationMessageSelectorRuleList55
		oi.Index = in["index"].(int)
		oi.Action = in["action"].(string)
		oi.MessageId = in["message_id"].(string)
		oi.MessageIdScope = in["message_id_scope"].(string)
		oi.SeverityOper = in["severity_oper"].(string)
		oi.SeverityVal = in["severity_val"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAcosEventsLogParameterization(d *schema.ResourceData) edpt.AcosEventsLogParameterization {
	var ret edpt.AcosEventsLogParameterization
	ret.Inst.LogRate = d.Get("log_rate").(int)
	ret.Inst.MessageSelector = getObjectAcosEventsLogParameterizationMessageSelector54(d.Get("message_selector").([]interface{}))
	//omit uuid
	return ret
}
