package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6OneToOneSharedPoolGroupOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_one_to_one_shared_pool_group_oper`: Operational Status for the object shared-pool-group\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6OneToOneSharedPoolGroupOperRead,

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

func resourceCgnv6OneToOneSharedPoolGroupOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOneSharedPoolGroupOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOneSharedPoolGroupOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6OneToOneSharedPoolGroupOperMembers := setObjectCgnv6OneToOneSharedPoolGroupOperMembers(res)
		d.Set("members", Cgnv6OneToOneSharedPoolGroupOperMembers)
		Cgnv6OneToOneSharedPoolGroupOperOper := setObjectCgnv6OneToOneSharedPoolGroupOperOper(res)
		d.Set("oper", Cgnv6OneToOneSharedPoolGroupOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6OneToOneSharedPoolGroupOperMembers(ret edpt.DataCgnv6OneToOneSharedPoolGroupOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectCgnv6OneToOneSharedPoolGroupOperMembersOper(ret.DtCgnv6OneToOneSharedPoolGroupOper.Members.Oper),
		},
	}
}

func setObjectCgnv6OneToOneSharedPoolGroupOperMembersOper(d edpt.Cgnv6OneToOneSharedPoolGroupOperMembersOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["member_list"] = setSliceCgnv6OneToOneSharedPoolGroupOperMembersOperMemberList(d.MemberList)

	in["pool_group_name"] = d.PoolGroupName
	result = append(result, in)
	return result
}

func setSliceCgnv6OneToOneSharedPoolGroupOperMembersOperMemberList(d []edpt.Cgnv6OneToOneSharedPoolGroupOperMembersOperMemberList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["pool_name"] = item.PoolName
		result = append(result, in)
	}
	return result
}

func setObjectCgnv6OneToOneSharedPoolGroupOperOper(ret edpt.DataCgnv6OneToOneSharedPoolGroupOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"shared_pool_group_list": setSliceCgnv6OneToOneSharedPoolGroupOperOperSharedPoolGroupList(ret.DtCgnv6OneToOneSharedPoolGroupOper.Oper.SharedPoolGroupList),
		},
	}
}

func setSliceCgnv6OneToOneSharedPoolGroupOperOperSharedPoolGroupList(d []edpt.Cgnv6OneToOneSharedPoolGroupOperOperSharedPoolGroupList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["pool_group_name"] = item.PoolGroupName
		in["vird"] = item.Vird
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6OneToOneSharedPoolGroupOperMembers(d []interface{}) edpt.Cgnv6OneToOneSharedPoolGroupOperMembers {

	count1 := len(d)
	var ret edpt.Cgnv6OneToOneSharedPoolGroupOperMembers
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectCgnv6OneToOneSharedPoolGroupOperMembersOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectCgnv6OneToOneSharedPoolGroupOperMembersOper(d []interface{}) edpt.Cgnv6OneToOneSharedPoolGroupOperMembersOper {

	count1 := len(d)
	var ret edpt.Cgnv6OneToOneSharedPoolGroupOperMembersOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MemberList = getSliceCgnv6OneToOneSharedPoolGroupOperMembersOperMemberList(in["member_list"].([]interface{}))
		ret.PoolGroupName = in["pool_group_name"].(string)
	}
	return ret
}

func getSliceCgnv6OneToOneSharedPoolGroupOperMembersOperMemberList(d []interface{}) []edpt.Cgnv6OneToOneSharedPoolGroupOperMembersOperMemberList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6OneToOneSharedPoolGroupOperMembersOperMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6OneToOneSharedPoolGroupOperMembersOperMemberList
		oi.PoolName = in["pool_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6OneToOneSharedPoolGroupOperOper(d []interface{}) edpt.Cgnv6OneToOneSharedPoolGroupOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6OneToOneSharedPoolGroupOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SharedPoolGroupList = getSliceCgnv6OneToOneSharedPoolGroupOperOperSharedPoolGroupList(in["shared_pool_group_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6OneToOneSharedPoolGroupOperOperSharedPoolGroupList(d []interface{}) []edpt.Cgnv6OneToOneSharedPoolGroupOperOperSharedPoolGroupList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6OneToOneSharedPoolGroupOperOperSharedPoolGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6OneToOneSharedPoolGroupOperOperSharedPoolGroupList
		oi.PoolGroupName = in["pool_group_name"].(string)
		oi.Vird = in["vird"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6OneToOneSharedPoolGroupOper(d *schema.ResourceData) edpt.Cgnv6OneToOneSharedPoolGroupOper {
	var ret edpt.Cgnv6OneToOneSharedPoolGroupOper

	ret.Members = getObjectCgnv6OneToOneSharedPoolGroupOperMembers(d.Get("members").([]interface{}))

	ret.Oper = getObjectCgnv6OneToOneSharedPoolGroupOperOper(d.Get("oper").([]interface{}))
	return ret
}
