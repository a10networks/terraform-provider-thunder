package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpCommunityListStandardNum() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_community_list_standard_num`: Configure Standard number Community-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpCommunityListStandardNumCreate,
		UpdateContext: resourceIpCommunityListStandardNumUpdate,
		ReadContext:   resourceIpCommunityListStandardNumRead,
		DeleteContext: resourceIpCommunityListStandardNumDelete,

		Schema: map[string]*schema.Schema{
			"rules_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"std_list_action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Specify community to reject; 'permit': Specify community to accept;",
						},
						"std_list_comm_value": {
							Type: schema.TypeString, Optional: true, Description: "community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export",
						},
					},
				},
			},
			"std_list_num": {
				Type: schema.TypeInt, Required: true, Description: "Community list number (standard)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpCommunityListStandardNumCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListStandardNumCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListStandardNum(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpCommunityListStandardNumRead(ctx, d, meta)
	}
	return diags
}

func resourceIpCommunityListStandardNumUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListStandardNumUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListStandardNum(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpCommunityListStandardNumRead(ctx, d, meta)
	}
	return diags
}
func resourceIpCommunityListStandardNumDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListStandardNumDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListStandardNum(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpCommunityListStandardNumRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpCommunityListStandardNumRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpCommunityListStandardNum(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpCommunityListStandardNumRulesList(d []interface{}) []edpt.IpCommunityListStandardNumRulesList {

	count1 := len(d)
	ret := make([]edpt.IpCommunityListStandardNumRulesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpCommunityListStandardNumRulesList
		oi.StdListAction = in["std_list_action"].(string)
		oi.StdListCommValue = in["std_list_comm_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpCommunityListStandardNum(d *schema.ResourceData) edpt.IpCommunityListStandardNum {
	var ret edpt.IpCommunityListStandardNum
	ret.Inst.RulesList = getSliceIpCommunityListStandardNumRulesList(d.Get("rules_list").([]interface{}))
	ret.Inst.StdListNum = d.Get("std_list_num").(int)
	//omit uuid
	return ret
}
