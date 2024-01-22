package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpCommunityListExpandedNum() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_community_list_expanded_num`: Configure Expanded number Community-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpCommunityListExpandedNumCreate,
		UpdateContext: resourceIpCommunityListExpandedNumUpdate,
		ReadContext:   resourceIpCommunityListExpandedNumRead,
		DeleteContext: resourceIpCommunityListExpandedNumDelete,

		Schema: map[string]*schema.Schema{
			"ext_list_num": {
				Type: schema.TypeInt, Required: true, Description: "Community list number (expanded)",
			},
			"rules_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ext_list_action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Specify community to reject; 'permit': Specify community to accept;",
						},
						"ext_list_value": {
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
func resourceIpCommunityListExpandedNumCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListExpandedNumCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListExpandedNum(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpCommunityListExpandedNumRead(ctx, d, meta)
	}
	return diags
}

func resourceIpCommunityListExpandedNumUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListExpandedNumUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListExpandedNum(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpCommunityListExpandedNumRead(ctx, d, meta)
	}
	return diags
}
func resourceIpCommunityListExpandedNumDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListExpandedNumDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListExpandedNum(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpCommunityListExpandedNumRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListExpandedNumRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListExpandedNum(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpCommunityListExpandedNumRulesList(d []interface{}) []edpt.IpCommunityListExpandedNumRulesList {

	count1 := len(d)
	ret := make([]edpt.IpCommunityListExpandedNumRulesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpCommunityListExpandedNumRulesList
		oi.ExtListAction = in["ext_list_action"].(string)
		oi.ExtListValue = in["ext_list_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpCommunityListExpandedNum(d *schema.ResourceData) edpt.IpCommunityListExpandedNum {
	var ret edpt.IpCommunityListExpandedNum
	ret.Inst.ExtListNum = d.Get("ext_list_num").(int)
	ret.Inst.RulesList = getSliceIpCommunityListExpandedNumRulesList(d.Get("rules_list").([]interface{}))
	//omit uuid
	return ret
}
