package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatDisabledConfigOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_disabled_config_oper`: Operational Status for the object disabled-config\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatDisabledConfigOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disabled_config_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inside_start_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"inside_end_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"inside_netmask": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inside_ip_list": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"partition": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"active_users": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"clear_session": {
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

func resourceCgnv6FixedNatDisabledConfigOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatDisabledConfigOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatDisabledConfigOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatDisabledConfigOperOper := setObjectCgnv6FixedNatDisabledConfigOperOper(res)
		d.Set("oper", Cgnv6FixedNatDisabledConfigOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatDisabledConfigOperOper(ret edpt.DataCgnv6FixedNatDisabledConfigOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"disabled_config_list": setSliceCgnv6FixedNatDisabledConfigOperOperDisabledConfigList(ret.DtCgnv6FixedNatDisabledConfigOper.Oper.DisabledConfigList),
		},
	}
}

func setSliceCgnv6FixedNatDisabledConfigOperOperDisabledConfigList(d []edpt.Cgnv6FixedNatDisabledConfigOperOperDisabledConfigList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["inside_start_address"] = item.InsideStartAddress
		in["inside_end_address"] = item.InsideEndAddress
		in["inside_netmask"] = item.InsideNetmask
		in["inside_ip_list"] = item.InsideIpList
		in["partition"] = item.Partition
		in["active_users"] = item.ActiveUsers
		in["clear_session"] = item.ClearSession
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6FixedNatDisabledConfigOperOper(d []interface{}) edpt.Cgnv6FixedNatDisabledConfigOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatDisabledConfigOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DisabledConfigList = getSliceCgnv6FixedNatDisabledConfigOperOperDisabledConfigList(in["disabled_config_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6FixedNatDisabledConfigOperOperDisabledConfigList(d []interface{}) []edpt.Cgnv6FixedNatDisabledConfigOperOperDisabledConfigList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6FixedNatDisabledConfigOperOperDisabledConfigList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6FixedNatDisabledConfigOperOperDisabledConfigList
		oi.InsideStartAddress = in["inside_start_address"].(string)
		oi.InsideEndAddress = in["inside_end_address"].(string)
		oi.InsideNetmask = in["inside_netmask"].(int)
		oi.InsideIpList = in["inside_ip_list"].(string)
		oi.Partition = in["partition"].(string)
		oi.ActiveUsers = in["active_users"].(int)
		oi.ClearSession = in["clear_session"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatDisabledConfigOper(d *schema.ResourceData) edpt.Cgnv6FixedNatDisabledConfigOper {
	var ret edpt.Cgnv6FixedNatDisabledConfigOper

	ret.Oper = getObjectCgnv6FixedNatDisabledConfigOperOper(d.Get("oper").([]interface{}))
	return ret
}
