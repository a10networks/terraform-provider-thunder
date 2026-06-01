package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionXflowInterfaceSelectionRegex() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_xflow_interface_selection_regex`: Configure regex rules to match interface name\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionXflowInterfaceSelectionRegexCreate,
		UpdateContext: resourceDdosDetectionXflowInterfaceSelectionRegexUpdate,
		ReadContext:   resourceDdosDetectionXflowInterfaceSelectionRegexRead,
		DeleteContext: resourceDdosDetectionXflowInterfaceSelectionRegexDelete,

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
			"type": {
				Type: schema.TypeString, Required: true, Description: "Type",
			},
		},
	}
}
func resourceDdosDetectionXflowInterfaceSelectionRegexCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionXflowInterfaceSelectionRegexCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionXflowInterfaceSelectionRegex(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionXflowInterfaceSelectionRegexRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionXflowInterfaceSelectionRegexUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionXflowInterfaceSelectionRegexUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionXflowInterfaceSelectionRegex(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionXflowInterfaceSelectionRegexRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionXflowInterfaceSelectionRegexDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionXflowInterfaceSelectionRegexDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionXflowInterfaceSelectionRegex(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionXflowInterfaceSelectionRegexRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionXflowInterfaceSelectionRegexRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionXflowInterfaceSelectionRegex(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDetectionXflowInterfaceSelectionRegexRuleList(d []interface{}) []edpt.DdosDetectionXflowInterfaceSelectionRegexRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionXflowInterfaceSelectionRegexRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionXflowInterfaceSelectionRegexRuleList
		oi.SingleRegex = in["single_regex"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDetectionXflowInterfaceSelectionRegex(d *schema.ResourceData) edpt.DdosDetectionXflowInterfaceSelectionRegex {
	var ret edpt.DdosDetectionXflowInterfaceSelectionRegex
	ret.Inst.RuleList = getSliceDdosDetectionXflowInterfaceSelectionRegexRuleList(d.Get("rule_list").([]interface{}))
	//omit uuid
	ret.Inst.Type = d.Get("type").(string)
	return ret
}
