package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkTrunkOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_trunk_oper`: Operational Status for the object trunk\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkTrunkOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trunk": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trunk_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"member_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"trunk_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"trunk_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"trunk_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"admin_key": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"trunk_member_status": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"members": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"cfg_status": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"oper_status": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"ports_threshold": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"timer": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"timer_running": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ports_threshold_block": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"working_lead": {
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

func resourceNetworkTrunkOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTrunkOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTrunkOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkTrunkOperOper := setObjectNetworkTrunkOperOper(res)
		d.Set("oper", NetworkTrunkOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkTrunkOperOper(ret edpt.DataNetworkTrunkOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"trunk": setSliceNetworkTrunkOperOperTrunk(ret.DtNetworkTrunkOper.Oper.Trunk),
		},
	}
}

func setSliceNetworkTrunkOperOperTrunk(d []edpt.NetworkTrunkOperOperTrunk) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["trunk_id"] = item.Trunk_id
		in["member_count"] = item.Member_count
		in["trunk_name"] = item.Trunk_name
		in["trunk_status"] = item.Trunk_status
		in["trunk_type"] = item.Trunk_type
		in["admin_key"] = item.Admin_key
		in["trunk_member_status"] = setSliceNetworkTrunkOperOperTrunkTrunkMemberStatus(item.TrunkMemberStatus)
		in["ports_threshold"] = item.Ports_threshold
		in["timer"] = item.Timer
		in["timer_running"] = item.Timer_running
		in["ports_threshold_block"] = item.Ports_threshold_block
		in["working_lead"] = item.Working_lead
		result = append(result, in)
	}
	return result
}

func setSliceNetworkTrunkOperOperTrunkTrunkMemberStatus(d []edpt.NetworkTrunkOperOperTrunkTrunkMemberStatus) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["members"] = item.Members
		in["cfg_status"] = item.Cfg_status
		in["oper_status"] = item.Oper_status
		result = append(result, in)
	}
	return result
}

func getObjectNetworkTrunkOperOper(d []interface{}) edpt.NetworkTrunkOperOper {

	count1 := len(d)
	var ret edpt.NetworkTrunkOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Trunk = getSliceNetworkTrunkOperOperTrunk(in["trunk"].([]interface{}))
	}
	return ret
}

func getSliceNetworkTrunkOperOperTrunk(d []interface{}) []edpt.NetworkTrunkOperOperTrunk {

	count1 := len(d)
	ret := make([]edpt.NetworkTrunkOperOperTrunk, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkTrunkOperOperTrunk
		oi.Trunk_id = in["trunk_id"].(int)
		oi.Member_count = in["member_count"].(int)
		oi.Trunk_name = in["trunk_name"].(string)
		oi.Trunk_status = in["trunk_status"].(string)
		oi.Trunk_type = in["trunk_type"].(string)
		oi.Admin_key = in["admin_key"].(int)
		oi.TrunkMemberStatus = getSliceNetworkTrunkOperOperTrunkTrunkMemberStatus(in["trunk_member_status"].([]interface{}))
		oi.Ports_threshold = in["ports_threshold"].(int)
		oi.Timer = in["timer"].(int)
		oi.Timer_running = in["timer_running"].(string)
		oi.Ports_threshold_block = in["ports_threshold_block"].(string)
		oi.Working_lead = in["working_lead"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetworkTrunkOperOperTrunkTrunkMemberStatus(d []interface{}) []edpt.NetworkTrunkOperOperTrunkTrunkMemberStatus {

	count1 := len(d)
	ret := make([]edpt.NetworkTrunkOperOperTrunkTrunkMemberStatus, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkTrunkOperOperTrunkTrunkMemberStatus
		oi.Members = in["members"].(int)
		oi.Cfg_status = in["cfg_status"].(string)
		oi.Oper_status = in["oper_status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkTrunkOper(d *schema.ResourceData) edpt.NetworkTrunkOper {
	var ret edpt.NetworkTrunkOper

	ret.Oper = getObjectNetworkTrunkOperOper(d.Get("oper").([]interface{}))
	return ret
}
