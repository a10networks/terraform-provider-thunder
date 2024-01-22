package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpCommunityListExpanded() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_community_list_expanded`: Configure Expanded Community-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpCommunityListExpandedCreate,
		UpdateContext: resourceIpCommunityListExpandedUpdate,
		ReadContext:   resourceIpCommunityListExpandedRead,
		DeleteContext: resourceIpCommunityListExpandedDelete,

		Schema: map[string]*schema.Schema{
			"expanded": {
				Type: schema.TypeString, Required: true, Description: "Add an expanded community-list entry (Community list name)",
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
func resourceIpCommunityListExpandedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListExpandedCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListExpanded(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpCommunityListExpandedRead(ctx, d, meta)
	}
	return diags
}

func resourceIpCommunityListExpandedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListExpandedUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListExpanded(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpCommunityListExpandedRead(ctx, d, meta)
	}
	return diags
}
func resourceIpCommunityListExpandedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListExpandedDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListExpanded(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpCommunityListExpandedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListExpandedRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListExpanded(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpCommunityListExpandedRulesList(d []interface{}) []edpt.IpCommunityListExpandedRulesList {

	count1 := len(d)
	ret := make([]edpt.IpCommunityListExpandedRulesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpCommunityListExpandedRulesList
		oi.ExpandedAction = in["expanded_action"].(string)
		oi.ExpandedValue = in["expanded_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpCommunityListExpanded(d *schema.ResourceData) edpt.IpCommunityListExpanded {
	var ret edpt.IpCommunityListExpanded
	ret.Inst.Expanded = d.Get("expanded").(string)
	ret.Inst.RulesList = getSliceIpCommunityListExpandedRulesList(d.Get("rules_list").([]interface{}))
	//omit uuid
	return ret
}
