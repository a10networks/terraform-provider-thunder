package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpLargeCommunityListExpanded() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_large_community_list_expanded`: Configure Expanded Large Community-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpLargeCommunityListExpandedCreate,
		UpdateContext: resourceIpLargeCommunityListExpandedUpdate,
		ReadContext:   resourceIpLargeCommunityListExpandedRead,
		DeleteContext: resourceIpLargeCommunityListExpandedDelete,

		Schema: map[string]*schema.Schema{
			"expanded": {
				Type: schema.TypeString, Required: true, Description: "Add an expanded large community-list entry (Large community list name)",
			},
			"rules_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expanded_lcom_action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Specify large community to reject; 'permit': Specify large community to accept;",
						},
						"expanded_lcom_value": {
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
func resourceIpLargeCommunityListExpandedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListExpandedCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListExpanded(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpLargeCommunityListExpandedRead(ctx, d, meta)
	}
	return diags
}

func resourceIpLargeCommunityListExpandedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListExpandedUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListExpanded(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpLargeCommunityListExpandedRead(ctx, d, meta)
	}
	return diags
}
func resourceIpLargeCommunityListExpandedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListExpandedDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListExpanded(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpLargeCommunityListExpandedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListExpandedRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListExpanded(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpLargeCommunityListExpandedRulesList(d []interface{}) []edpt.IpLargeCommunityListExpandedRulesList {

	count1 := len(d)
	ret := make([]edpt.IpLargeCommunityListExpandedRulesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpLargeCommunityListExpandedRulesList
		oi.ExpandedLcomAction = in["expanded_lcom_action"].(string)
		oi.ExpandedLcomValue = in["expanded_lcom_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpLargeCommunityListExpanded(d *schema.ResourceData) edpt.IpLargeCommunityListExpanded {
	var ret edpt.IpLargeCommunityListExpanded
	ret.Inst.Expanded = d.Get("expanded").(string)
	ret.Inst.RulesList = getSliceIpLargeCommunityListExpandedRulesList(d.Get("rules_list").([]interface{}))
	//omit uuid
	return ret
}
