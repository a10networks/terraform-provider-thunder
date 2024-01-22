package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpLargeCommunityListStandardNum() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_large_community_list_standard_num`: Configure Standard number Large Community-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpLargeCommunityListStandardNumCreate,
		UpdateContext: resourceIpLargeCommunityListStandardNumUpdate,
		ReadContext:   resourceIpLargeCommunityListStandardNumRead,
		DeleteContext: resourceIpLargeCommunityListStandardNumDelete,

		Schema: map[string]*schema.Schema{
			"rules_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"std_list_lcom_action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Specify large community to reject; 'permit': Specify large community to accept;",
						},
						"std_list_lcomm_value": {
							Type: schema.TypeString, Optional: true, Description: "Large community value in the format XX:YY:ZZ",
						},
					},
				},
			},
			"std_list_num": {
				Type: schema.TypeInt, Required: true, Description: "Large Community list number (standard)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpLargeCommunityListStandardNumCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListStandardNumCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListStandardNum(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpLargeCommunityListStandardNumRead(ctx, d, meta)
	}
	return diags
}

func resourceIpLargeCommunityListStandardNumUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListStandardNumUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListStandardNum(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpLargeCommunityListStandardNumRead(ctx, d, meta)
	}
	return diags
}
func resourceIpLargeCommunityListStandardNumDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListStandardNumDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListStandardNum(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpLargeCommunityListStandardNumRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpLargeCommunityListStandardNumRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpLargeCommunityListStandardNum(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpLargeCommunityListStandardNumRulesList(d []interface{}) []edpt.IpLargeCommunityListStandardNumRulesList {

	count1 := len(d)
	ret := make([]edpt.IpLargeCommunityListStandardNumRulesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpLargeCommunityListStandardNumRulesList
		oi.StdListLcomAction = in["std_list_lcom_action"].(string)
		oi.StdListLcommValue = in["std_list_lcomm_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpLargeCommunityListStandardNum(d *schema.ResourceData) edpt.IpLargeCommunityListStandardNum {
	var ret edpt.IpLargeCommunityListStandardNum
	ret.Inst.RulesList = getSliceIpLargeCommunityListStandardNumRulesList(d.Get("rules_list").([]interface{}))
	ret.Inst.StdListNum = d.Get("std_list_num").(int)
	//omit uuid
	return ret
}
