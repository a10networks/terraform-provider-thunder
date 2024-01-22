package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpExtcommunityListExpandedNum() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_extcommunity_list_expanded_num`: Configure Expanded number Community-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpExtcommunityListExpandedNumCreate,
		UpdateContext: resourceIpExtcommunityListExpandedNumUpdate,
		ReadContext:   resourceIpExtcommunityListExpandedNumRead,
		DeleteContext: resourceIpExtcommunityListExpandedNumDelete,

		Schema: map[string]*schema.Schema{
			"ext_list_num": {
				Type: schema.TypeInt, Required: true, Description: "Extended Community list number (expanded)",
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
func resourceIpExtcommunityListExpandedNumCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListExpandedNumCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListExpandedNum(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpExtcommunityListExpandedNumRead(ctx, d, meta)
	}
	return diags
}

func resourceIpExtcommunityListExpandedNumUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListExpandedNumUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListExpandedNum(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpExtcommunityListExpandedNumRead(ctx, d, meta)
	}
	return diags
}
func resourceIpExtcommunityListExpandedNumDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListExpandedNumDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListExpandedNum(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpExtcommunityListExpandedNumRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListExpandedNumRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListExpandedNum(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpExtcommunityListExpandedNumRulesList(d []interface{}) []edpt.IpExtcommunityListExpandedNumRulesList {

	count1 := len(d)
	ret := make([]edpt.IpExtcommunityListExpandedNumRulesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpExtcommunityListExpandedNumRulesList
		oi.ExtListAction = in["ext_list_action"].(string)
		oi.ExtListValue = in["ext_list_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpExtcommunityListExpandedNum(d *schema.ResourceData) edpt.IpExtcommunityListExpandedNum {
	var ret edpt.IpExtcommunityListExpandedNum
	ret.Inst.ExtListNum = d.Get("ext_list_num").(int)
	ret.Inst.RulesList = getSliceIpExtcommunityListExpandedNumRulesList(d.Get("rules_list").([]interface{}))
	//omit uuid
	return ret
}
