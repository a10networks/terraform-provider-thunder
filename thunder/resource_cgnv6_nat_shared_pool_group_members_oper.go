package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatSharedPoolGroupMembersOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_nat_shared_pool_group_members_oper`: Operational Status for the object members\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6NatSharedPoolGroupMembersOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pool_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"pool_group_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6NatSharedPoolGroupMembersOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatSharedPoolGroupMembersOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatSharedPoolGroupMembersOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6NatSharedPoolGroupMembersOperOper := setObjectCgnv6NatSharedPoolGroupMembersOperOper(res)
		d.Set("oper", Cgnv6NatSharedPoolGroupMembersOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6NatSharedPoolGroupMembersOperOper(ret edpt.DataCgnv6NatSharedPoolGroupMembersOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"member_list":     setSliceCgnv6NatSharedPoolGroupMembersOperOperMemberList(ret.DtCgnv6NatSharedPoolGroupMembersOper.Oper.MemberList),
			"pool_group_name": ret.DtCgnv6NatSharedPoolGroupMembersOper.Oper.PoolGroupName,
		},
	}
}

func setSliceCgnv6NatSharedPoolGroupMembersOperOperMemberList(d []edpt.Cgnv6NatSharedPoolGroupMembersOperOperMemberList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["pool_name"] = item.PoolName
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6NatSharedPoolGroupMembersOperOper(d []interface{}) edpt.Cgnv6NatSharedPoolGroupMembersOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6NatSharedPoolGroupMembersOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MemberList = getSliceCgnv6NatSharedPoolGroupMembersOperOperMemberList(in["member_list"].([]interface{}))
		ret.PoolGroupName = in["pool_group_name"].(string)
	}
	return ret
}

func getSliceCgnv6NatSharedPoolGroupMembersOperOperMemberList(d []interface{}) []edpt.Cgnv6NatSharedPoolGroupMembersOperOperMemberList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6NatSharedPoolGroupMembersOperOperMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6NatSharedPoolGroupMembersOperOperMemberList
		oi.PoolName = in["pool_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6NatSharedPoolGroupMembersOper(d *schema.ResourceData) edpt.Cgnv6NatSharedPoolGroupMembersOper {
	var ret edpt.Cgnv6NatSharedPoolGroupMembersOper

	ret.Oper = getObjectCgnv6NatSharedPoolGroupMembersOperOper(d.Get("oper").([]interface{}))
	return ret
}
