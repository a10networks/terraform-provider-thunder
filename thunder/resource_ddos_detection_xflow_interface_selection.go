package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionXflowInterfaceSelection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_xflow_interface_selection`: Configure rules to select interface\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionXflowInterfaceSelectionCreate,
		UpdateContext: resourceDdosDetectionXflowInterfaceSelectionUpdate,
		ReadContext:   resourceDdosDetectionXflowInterfaceSelectionRead,
		DeleteContext: resourceDdosDetectionXflowInterfaceSelectionDelete,

		Schema: map[string]*schema.Schema{
			"regex": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"single_regex": {
										Type: schema.TypeString, Optional: true, Description: "Specify the regular expression rules",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"type": {
				Type: schema.TypeString, Required: true, Description: "'internet-side': internet-side;",
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
func resourceDdosDetectionXflowInterfaceSelectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionXflowInterfaceSelectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionXflowInterfaceSelection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionXflowInterfaceSelectionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionXflowInterfaceSelectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionXflowInterfaceSelectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionXflowInterfaceSelection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionXflowInterfaceSelectionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionXflowInterfaceSelectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionXflowInterfaceSelectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionXflowInterfaceSelection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionXflowInterfaceSelectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionXflowInterfaceSelectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionXflowInterfaceSelection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDetectionXflowInterfaceSelectionRegex155(d []interface{}) edpt.DdosDetectionXflowInterfaceSelectionRegex155 {

	count1 := len(d)
	var ret edpt.DdosDetectionXflowInterfaceSelectionRegex155
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDetectionXflowInterfaceSelectionRegexRuleList156(in["rule_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceDdosDetectionXflowInterfaceSelectionRegexRuleList156(d []interface{}) []edpt.DdosDetectionXflowInterfaceSelectionRegexRuleList156 {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionXflowInterfaceSelectionRegexRuleList156, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionXflowInterfaceSelectionRegexRuleList156
		oi.SingleRegex = in["single_regex"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDetectionXflowInterfaceSelection(d *schema.ResourceData) edpt.DdosDetectionXflowInterfaceSelection {
	var ret edpt.DdosDetectionXflowInterfaceSelection
	ret.Inst.Regex = getObjectDdosDetectionXflowInterfaceSelectionRegex155(d.Get("regex").([]interface{}))
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
