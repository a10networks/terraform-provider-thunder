package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpCommunityListStandard() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_community_list_standard`: Configure Standard Community-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpCommunityListStandardCreate,
		UpdateContext: resourceIpCommunityListStandardUpdate,
		ReadContext:   resourceIpCommunityListStandardRead,
		DeleteContext: resourceIpCommunityListStandardDelete,

		Schema: map[string]*schema.Schema{
			"rules_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"standard_action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Specify community to reject; 'permit': Specify community to accept;",
						},
						"standard_comm_value": {
							Type: schema.TypeString, Optional: true, Description: "community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export",
						},
					},
				},
			},
			"standard": {
				Type: schema.TypeString, Required: true, Description: "Add a standard community-list entry (Community list name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpCommunityListStandardCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListStandardCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListStandard(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpCommunityListStandardRead(ctx, d, meta)
	}
	return diags
}

func resourceIpCommunityListStandardUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListStandardUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListStandard(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpCommunityListStandardRead(ctx, d, meta)
	}
	return diags
}
func resourceIpCommunityListStandardDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListStandardDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListStandard(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpCommunityListStandardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListStandardRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListStandard(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpCommunityListStandardRulesList(d []interface{}) []edpt.IpCommunityListStandardRulesList {

	count1 := len(d)
	ret := make([]edpt.IpCommunityListStandardRulesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpCommunityListStandardRulesList
		oi.StandardAction = in["standard_action"].(string)
		oi.StandardCommValue = in["standard_comm_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpCommunityListStandard(d *schema.ResourceData) edpt.IpCommunityListStandard {
	var ret edpt.IpCommunityListStandard
	ret.Inst.RulesList = getSliceIpCommunityListStandardRulesList(d.Get("rules_list").([]interface{}))
	ret.Inst.Standard = d.Get("standard").(string)
	//omit uuid
	return ret
}
