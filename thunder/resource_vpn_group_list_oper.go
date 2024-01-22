package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnGroupListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_group_list_oper`: Operational Status for the object group-list\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnGroupListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"group_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipsec_sa_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_gateway_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"role": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_new_group": {
										Type: schema.TypeInt, Optional: true, Description: "a value of 1 indicates new group",
									},
									"grp_member_count": {
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

func resourceVpnGroupListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnGroupListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnGroupListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnGroupListOperOper := setObjectVpnGroupListOperOper(res)
		d.Set("oper", VpnGroupListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnGroupListOperOper(ret edpt.DataVpnGroupListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"group_name": ret.DtVpnGroupListOper.Oper.GroupName,
			"group_list": setSliceVpnGroupListOperOperGroupList(ret.DtVpnGroupListOper.Oper.GroupList),
		},
	}
}

func setSliceVpnGroupListOperOperGroupList(d []edpt.VpnGroupListOperOperGroupList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["ipsec_sa_name"] = item.IpsecSaName
		in["ike_gateway_name"] = item.IkeGatewayName
		in["priority"] = item.Priority
		in["status"] = item.Status
		in["role"] = item.Role
		in["is_new_group"] = item.IsNewGroup
		in["grp_member_count"] = item.GrpMemberCount
		result = append(result, in)
	}
	return result
}

func getObjectVpnGroupListOperOper(d []interface{}) edpt.VpnGroupListOperOper {

	count1 := len(d)
	var ret edpt.VpnGroupListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GroupName = in["group_name"].(string)
		ret.GroupList = getSliceVpnGroupListOperOperGroupList(in["group_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnGroupListOperOperGroupList(d []interface{}) []edpt.VpnGroupListOperOperGroupList {

	count1 := len(d)
	ret := make([]edpt.VpnGroupListOperOperGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnGroupListOperOperGroupList
		oi.Name = in["name"].(string)
		oi.IpsecSaName = in["ipsec_sa_name"].(string)
		oi.IkeGatewayName = in["ike_gateway_name"].(string)
		oi.Priority = in["priority"].(int)
		oi.Status = in["status"].(string)
		oi.Role = in["role"].(string)
		oi.IsNewGroup = in["is_new_group"].(int)
		oi.GrpMemberCount = in["grp_member_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnGroupListOper(d *schema.ResourceData) edpt.VpnGroupListOper {
	var ret edpt.VpnGroupListOper

	ret.Oper = getObjectVpnGroupListOperOper(d.Get("oper").([]interface{}))
	return ret
}
