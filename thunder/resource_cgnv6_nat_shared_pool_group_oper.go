package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatSharedPoolGroupOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_nat_shared_pool_group_oper`: Operational Status for the object shared-pool-group\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6NatSharedPoolGroupOperRead,

		Schema: map[string]*schema.Schema{
			"members": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"shared_pool_group_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pool_group_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vird": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6NatSharedPoolGroupOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatSharedPoolGroupOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatSharedPoolGroupOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6NatSharedPoolGroupOperMembers := setObjectCgnv6NatSharedPoolGroupOperMembers(res)
		d.Set("members", Cgnv6NatSharedPoolGroupOperMembers)
		Cgnv6NatSharedPoolGroupOperOper := setObjectCgnv6NatSharedPoolGroupOperOper(res)
		d.Set("oper", Cgnv6NatSharedPoolGroupOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6NatSharedPoolGroupOperMembers(ret edpt.DataCgnv6NatSharedPoolGroupOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectCgnv6NatSharedPoolGroupOperMembersOper(ret.DtCgnv6NatSharedPoolGroupOper.Members.Oper),
		},
	}
}

func setObjectCgnv6NatSharedPoolGroupOperMembersOper(d edpt.Cgnv6NatSharedPoolGroupOperMembersOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["member_list"] = setSliceCgnv6NatSharedPoolGroupOperMembersOperMemberList(d.MemberList)

	in["pool_group_name"] = d.PoolGroupName
	result = append(result, in)
	return result
}

func setSliceCgnv6NatSharedPoolGroupOperMembersOperMemberList(d []edpt.Cgnv6NatSharedPoolGroupOperMembersOperMemberList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["pool_name"] = item.PoolName
		result = append(result, in)
	}
	return result
}

func setObjectCgnv6NatSharedPoolGroupOperOper(ret edpt.DataCgnv6NatSharedPoolGroupOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"shared_pool_group_list": setSliceCgnv6NatSharedPoolGroupOperOperSharedPoolGroupList(ret.DtCgnv6NatSharedPoolGroupOper.Oper.SharedPoolGroupList),
		},
	}
}

func setSliceCgnv6NatSharedPoolGroupOperOperSharedPoolGroupList(d []edpt.Cgnv6NatSharedPoolGroupOperOperSharedPoolGroupList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["pool_group_name"] = item.PoolGroupName
		in["vird"] = item.Vird
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6NatSharedPoolGroupOperMembers(d []interface{}) edpt.Cgnv6NatSharedPoolGroupOperMembers {

	count1 := len(d)
	var ret edpt.Cgnv6NatSharedPoolGroupOperMembers
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectCgnv6NatSharedPoolGroupOperMembersOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectCgnv6NatSharedPoolGroupOperMembersOper(d []interface{}) edpt.Cgnv6NatSharedPoolGroupOperMembersOper {

	count1 := len(d)
	var ret edpt.Cgnv6NatSharedPoolGroupOperMembersOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MemberList = getSliceCgnv6NatSharedPoolGroupOperMembersOperMemberList(in["member_list"].([]interface{}))
		ret.PoolGroupName = in["pool_group_name"].(string)
	}
	return ret
}

func getSliceCgnv6NatSharedPoolGroupOperMembersOperMemberList(d []interface{}) []edpt.Cgnv6NatSharedPoolGroupOperMembersOperMemberList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6NatSharedPoolGroupOperMembersOperMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6NatSharedPoolGroupOperMembersOperMemberList
		oi.PoolName = in["pool_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6NatSharedPoolGroupOperOper(d []interface{}) edpt.Cgnv6NatSharedPoolGroupOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6NatSharedPoolGroupOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SharedPoolGroupList = getSliceCgnv6NatSharedPoolGroupOperOperSharedPoolGroupList(in["shared_pool_group_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6NatSharedPoolGroupOperOperSharedPoolGroupList(d []interface{}) []edpt.Cgnv6NatSharedPoolGroupOperOperSharedPoolGroupList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6NatSharedPoolGroupOperOperSharedPoolGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6NatSharedPoolGroupOperOperSharedPoolGroupList
		oi.PoolGroupName = in["pool_group_name"].(string)
		oi.Vird = in["vird"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6NatSharedPoolGroupOper(d *schema.ResourceData) edpt.Cgnv6NatSharedPoolGroupOper {
	var ret edpt.Cgnv6NatSharedPoolGroupOper

	ret.Members = getObjectCgnv6NatSharedPoolGroupOperMembers(d.Get("members").([]interface{}))

	ret.Oper = getObjectCgnv6NatSharedPoolGroupOperOper(d.Get("oper").([]interface{}))
	return ret
}
