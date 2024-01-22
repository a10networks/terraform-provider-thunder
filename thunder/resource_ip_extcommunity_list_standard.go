package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpExtcommunityListStandard() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_extcommunity_list_standard`: Configure Standard Community-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpExtcommunityListStandardCreate,
		UpdateContext: resourceIpExtcommunityListStandardUpdate,
		ReadContext:   resourceIpExtcommunityListStandardRead,
		DeleteContext: resourceIpExtcommunityListStandardDelete,

		Schema: map[string]*schema.Schema{
			"rules_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"standard_action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Specify community to reject; 'permit': Specify community to accept;",
						},
						"standard_value": {
							Type: schema.TypeString, Optional: true, Description: "rt Route Target extended community in aa:nn or IPaddr:nn format OR soo Site-of-Origin extended community in aa:nn or IPaddr:nn",
						},
					},
				},
			},
			"standard": {
				Type: schema.TypeString, Required: true, Description: "Add a standard extcommunity-list entry (Extended Community list name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpExtcommunityListStandardCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListStandardCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListStandard(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpExtcommunityListStandardRead(ctx, d, meta)
	}
	return diags
}

func resourceIpExtcommunityListStandardUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListStandardUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListStandard(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpExtcommunityListStandardRead(ctx, d, meta)
	}
	return diags
}
func resourceIpExtcommunityListStandardDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListStandardDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListStandard(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpExtcommunityListStandardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListStandardRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListStandard(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpExtcommunityListStandardRulesList(d []interface{}) []edpt.IpExtcommunityListStandardRulesList {

	count1 := len(d)
	ret := make([]edpt.IpExtcommunityListStandardRulesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpExtcommunityListStandardRulesList
		oi.StandardAction = in["standard_action"].(string)
		oi.StandardValue = in["standard_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpExtcommunityListStandard(d *schema.ResourceData) edpt.IpExtcommunityListStandard {
	var ret edpt.IpExtcommunityListStandard
	ret.Inst.RulesList = getSliceIpExtcommunityListStandardRulesList(d.Get("rules_list").([]interface{}))
	ret.Inst.Standard = d.Get("standard").(string)
	//omit uuid
	return ret
}
