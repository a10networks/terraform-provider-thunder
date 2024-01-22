package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpLargeCommunityListExpandedNum() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_large_community_list_expanded_num`: Configure Expanded number Large Community-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpLargeCommunityListExpandedNumCreate,
		UpdateContext: resourceIpLargeCommunityListExpandedNumUpdate,
		ReadContext:   resourceIpLargeCommunityListExpandedNumRead,
		DeleteContext: resourceIpLargeCommunityListExpandedNumDelete,

		Schema: map[string]*schema.Schema{
			"ext_list_num": {
				Type: schema.TypeInt, Required: true, Description: "Large community list number (expanded)",
			},
			"rules_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ext_list_lcom_action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Specify large community to reject; 'permit': Specify large community to accept;",
						},
						"ext_list_lcom_value": {
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
func resourceIpLargeCommunityListExpandedNumCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListExpandedNumCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListExpandedNum(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpLargeCommunityListExpandedNumRead(ctx, d, meta)
	}
	return diags
}

func resourceIpLargeCommunityListExpandedNumUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListExpandedNumUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListExpandedNum(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpLargeCommunityListExpandedNumRead(ctx, d, meta)
	}
	return diags
}
func resourceIpLargeCommunityListExpandedNumDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListExpandedNumDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListExpandedNum(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpLargeCommunityListExpandedNumRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListExpandedNumRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListExpandedNum(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpLargeCommunityListExpandedNumRulesList(d []interface{}) []edpt.IpLargeCommunityListExpandedNumRulesList {

	count1 := len(d)
	ret := make([]edpt.IpLargeCommunityListExpandedNumRulesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpLargeCommunityListExpandedNumRulesList
		oi.ExtListLcomAction = in["ext_list_lcom_action"].(string)
		oi.ExtListLcomValue = in["ext_list_lcom_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpLargeCommunityListExpandedNum(d *schema.ResourceData) edpt.IpLargeCommunityListExpandedNum {
	var ret edpt.IpLargeCommunityListExpandedNum
	ret.Inst.ExtListNum = d.Get("ext_list_num").(int)
	ret.Inst.RulesList = getSliceIpLargeCommunityListExpandedNumRulesList(d.Get("rules_list").([]interface{}))
	//omit uuid
	return ret
}
