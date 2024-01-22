package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpExtcommunityListExpanded() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_extcommunity_list_expanded`: Configure Extended Community-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpExtcommunityListExpandedCreate,
		UpdateContext: resourceIpExtcommunityListExpandedUpdate,
		ReadContext:   resourceIpExtcommunityListExpandedRead,
		DeleteContext: resourceIpExtcommunityListExpandedDelete,

		Schema: map[string]*schema.Schema{
			"expanded": {
				Type: schema.TypeString, Required: true, Description: "Add an expanded extcommunity-list entry (Extended Community list name)",
			},
			"rules_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expanded_action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Specify community to reject; 'permit': Specify community to accept;",
						},
						"expanded_value": {
							Type: schema.TypeString, Optional: true, Description: "An ordered list as a regular-expression",
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
func resourceIpExtcommunityListExpandedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListExpandedCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListExpanded(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpExtcommunityListExpandedRead(ctx, d, meta)
	}
	return diags
}

func resourceIpExtcommunityListExpandedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListExpandedUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListExpanded(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpExtcommunityListExpandedRead(ctx, d, meta)
	}
	return diags
}
func resourceIpExtcommunityListExpandedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListExpandedDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListExpanded(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpExtcommunityListExpandedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListExpandedRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListExpanded(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpExtcommunityListExpandedRulesList(d []interface{}) []edpt.IpExtcommunityListExpandedRulesList {

	count1 := len(d)
	ret := make([]edpt.IpExtcommunityListExpandedRulesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpExtcommunityListExpandedRulesList
		oi.ExpandedAction = in["expanded_action"].(string)
		oi.ExpandedValue = in["expanded_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpExtcommunityListExpanded(d *schema.ResourceData) edpt.IpExtcommunityListExpanded {
	var ret edpt.IpExtcommunityListExpanded
	ret.Inst.Expanded = d.Get("expanded").(string)
	ret.Inst.RulesList = getSliceIpExtcommunityListExpandedRulesList(d.Get("rules_list").([]interface{}))
	//omit uuid
	return ret
}
