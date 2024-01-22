package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpExtcommunityListStandardNum() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_extcommunity_list_standard_num`: Configure Standard number Community-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpExtcommunityListStandardNumCreate,
		UpdateContext: resourceIpExtcommunityListStandardNumUpdate,
		ReadContext:   resourceIpExtcommunityListStandardNumRead,
		DeleteContext: resourceIpExtcommunityListStandardNumDelete,

		Schema: map[string]*schema.Schema{
			"rules_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"std_list_action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Specify community to reject; 'permit': Specify community to accept;",
						},
						"std_list_value": {
							Type: schema.TypeString, Optional: true, Description: "rt Route Target extended community in aa:nn or IPaddr:nn format OR soo Site-of-Origin extended community in aa:nn or IPaddr:nn",
						},
					},
				},
			},
			"std_list_num": {
				Type: schema.TypeInt, Required: true, Description: "Extended Community list number (standard)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpExtcommunityListStandardNumCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListStandardNumCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListStandardNum(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpExtcommunityListStandardNumRead(ctx, d, meta)
	}
	return diags
}

func resourceIpExtcommunityListStandardNumUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListStandardNumUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListStandardNum(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpExtcommunityListStandardNumRead(ctx, d, meta)
	}
	return diags
}
func resourceIpExtcommunityListStandardNumDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListStandardNumDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListStandardNum(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpExtcommunityListStandardNumRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpExtcommunityListStandardNumRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpExtcommunityListStandardNum(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpExtcommunityListStandardNumRulesList(d []interface{}) []edpt.IpExtcommunityListStandardNumRulesList {

	count1 := len(d)
	ret := make([]edpt.IpExtcommunityListStandardNumRulesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpExtcommunityListStandardNumRulesList
		oi.StdListAction = in["std_list_action"].(string)
		oi.StdListValue = in["std_list_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpExtcommunityListStandardNum(d *schema.ResourceData) edpt.IpExtcommunityListStandardNum {
	var ret edpt.IpExtcommunityListStandardNum
	ret.Inst.RulesList = getSliceIpExtcommunityListStandardNumRulesList(d.Get("rules_list").([]interface{}))
	ret.Inst.StdListNum = d.Get("std_list_num").(int)
	//omit uuid
	return ret
}
