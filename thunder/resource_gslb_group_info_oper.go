package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbGroupInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_group_info_oper`: Operational Status for the object group-info\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbGroupInfoOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"member_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_master": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sys_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"learn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"passive": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipv6_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"connect_success": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connect_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"open_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"open_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"open_success": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sync_sequence_number": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"update_in": {
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

func resourceGslbGroupInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbGroupInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbGroupInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbGroupInfoOperOper := setObjectGslbGroupInfoOperOper(res)
		d.Set("oper", GslbGroupInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbGroupInfoOperOper(ret edpt.DataGslbGroupInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"member_list": setSliceGslbGroupInfoOperOperMemberList(ret.DtGslbGroupInfoOper.Oper.MemberList),
		},
	}
}

func setSliceGslbGroupInfoOperOperMemberList(d []edpt.GslbGroupInfoOperOperMemberList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["group_name"] = item.GroupName
		in["member_name"] = item.MemberName
		in["is_master"] = item.Is_master
		in["sys_id"] = item.SysId
		in["priority"] = item.Priority
		in["learn"] = item.Learn
		in["passive"] = item.Passive
		in["status"] = item.Status
		in["address"] = item.Address
		in["ipv6_address"] = item.Ipv6Address
		in["connect_success"] = item.ConnectSuccess
		in["connect_fail"] = item.ConnectFail
		in["open_in"] = item.OpenIn
		in["open_out"] = item.OpenOut
		in["open_success"] = item.OpenSuccess
		in["sync_sequence_number"] = item.SyncSequenceNumber
		in["update_in"] = item.UpdateIn
		result = append(result, in)
	}
	return result
}

func getObjectGslbGroupInfoOperOper(d []interface{}) edpt.GslbGroupInfoOperOper {

	count1 := len(d)
	var ret edpt.GslbGroupInfoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MemberList = getSliceGslbGroupInfoOperOperMemberList(in["member_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbGroupInfoOperOperMemberList(d []interface{}) []edpt.GslbGroupInfoOperOperMemberList {

	count1 := len(d)
	ret := make([]edpt.GslbGroupInfoOperOperMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbGroupInfoOperOperMemberList
		oi.GroupName = in["group_name"].(string)
		oi.MemberName = in["member_name"].(string)
		oi.Is_master = in["is_master"].(int)
		oi.SysId = in["sys_id"].(int)
		oi.Priority = in["priority"].(int)
		oi.Learn = in["learn"].(int)
		oi.Passive = in["passive"].(int)
		oi.Status = in["status"].(string)
		oi.Address = in["address"].(string)
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.ConnectSuccess = in["connect_success"].(int)
		oi.ConnectFail = in["connect_fail"].(int)
		oi.OpenIn = in["open_in"].(int)
		oi.OpenOut = in["open_out"].(int)
		oi.OpenSuccess = in["open_success"].(int)
		oi.SyncSequenceNumber = in["sync_sequence_number"].(int)
		oi.UpdateIn = in["update_in"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbGroupInfoOper(d *schema.ResourceData) edpt.GslbGroupInfoOper {
	var ret edpt.GslbGroupInfoOper

	ret.Oper = getObjectGslbGroupInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
