package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6OneToOneSharedPoolGroupMembersOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_one_to_one_shared_pool_group_members_oper`: Operational Status for the object members\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6OneToOneSharedPoolGroupMembersOperRead,

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

func resourceCgnv6OneToOneSharedPoolGroupMembersOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOneSharedPoolGroupMembersOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOneSharedPoolGroupMembersOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6OneToOneSharedPoolGroupMembersOperOper := setObjectCgnv6OneToOneSharedPoolGroupMembersOperOper(res)
		d.Set("oper", Cgnv6OneToOneSharedPoolGroupMembersOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6OneToOneSharedPoolGroupMembersOperOper(ret edpt.DataCgnv6OneToOneSharedPoolGroupMembersOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"member_list":     setSliceCgnv6OneToOneSharedPoolGroupMembersOperOperMemberList(ret.DtCgnv6OneToOneSharedPoolGroupMembersOper.Oper.MemberList),
			"pool_group_name": ret.DtCgnv6OneToOneSharedPoolGroupMembersOper.Oper.PoolGroupName,
		},
	}
}

func setSliceCgnv6OneToOneSharedPoolGroupMembersOperOperMemberList(d []edpt.Cgnv6OneToOneSharedPoolGroupMembersOperOperMemberList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["pool_name"] = item.PoolName
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6OneToOneSharedPoolGroupMembersOperOper(d []interface{}) edpt.Cgnv6OneToOneSharedPoolGroupMembersOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6OneToOneSharedPoolGroupMembersOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MemberList = getSliceCgnv6OneToOneSharedPoolGroupMembersOperOperMemberList(in["member_list"].([]interface{}))
		ret.PoolGroupName = in["pool_group_name"].(string)
	}
	return ret
}

func getSliceCgnv6OneToOneSharedPoolGroupMembersOperOperMemberList(d []interface{}) []edpt.Cgnv6OneToOneSharedPoolGroupMembersOperOperMemberList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6OneToOneSharedPoolGroupMembersOperOperMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6OneToOneSharedPoolGroupMembersOperOperMemberList
		oi.PoolName = in["pool_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6OneToOneSharedPoolGroupMembersOper(d *schema.ResourceData) edpt.Cgnv6OneToOneSharedPoolGroupMembersOper {
	var ret edpt.Cgnv6OneToOneSharedPoolGroupMembersOper

	ret.Oper = getObjectCgnv6OneToOneSharedPoolGroupMembersOperOper(d.Get("oper").([]interface{}))
	return ret
}
