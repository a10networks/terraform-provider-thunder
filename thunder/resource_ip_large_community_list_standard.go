package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpLargeCommunityListStandard() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_large_community_list_standard`: Configure Standard Large Community-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpLargeCommunityListStandardCreate,
		UpdateContext: resourceIpLargeCommunityListStandardUpdate,
		ReadContext:   resourceIpLargeCommunityListStandardRead,
		DeleteContext: resourceIpLargeCommunityListStandardDelete,

		Schema: map[string]*schema.Schema{
			"rules_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"standard_lcom_action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Specify large community to reject; 'permit': Specify large community to accept;",
						},
						"standard_lcomm_value": {
							Type: schema.TypeString, Optional: true, Description: "Large community value in the format XX:YY:ZZ",
						},
					},
				},
			},
			"standard": {
				Type: schema.TypeString, Required: true, Description: "Add a standard large community-list entry (Large community list name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpLargeCommunityListStandardCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListStandardCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListStandard(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpLargeCommunityListStandardRead(ctx, d, meta)
	}
	return diags
}

func resourceIpLargeCommunityListStandardUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListStandardUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListStandard(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpLargeCommunityListStandardRead(ctx, d, meta)
	}
	return diags
}
func resourceIpLargeCommunityListStandardDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListStandardDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListStandard(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpLargeCommunityListStandardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListStandardRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListStandard(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpLargeCommunityListStandardRulesList(d []interface{}) []edpt.IpLargeCommunityListStandardRulesList {

	count1 := len(d)
	ret := make([]edpt.IpLargeCommunityListStandardRulesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpLargeCommunityListStandardRulesList
		oi.StandardLcomAction = in["standard_lcom_action"].(string)
		oi.StandardLcommValue = in["standard_lcomm_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpLargeCommunityListStandard(d *schema.ResourceData) edpt.IpLargeCommunityListStandard {
	var ret edpt.IpLargeCommunityListStandard
	ret.Inst.RulesList = getSliceIpLargeCommunityListStandardRulesList(d.Get("rules_list").([]interface{}))
	ret.Inst.Standard = d.Get("standard").(string)
	//omit uuid
	return ret
}
