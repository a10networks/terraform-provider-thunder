package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DdosProtectionIpEntriesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_ddos_protection_ip_entries_oper`: Operational Status for the object ip-entries\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6DdosProtectionIpEntriesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ddos_ip_entries_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"v4_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"expiration": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hints": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"l4_entries_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sw_receive_pps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sw_l4_drop_pps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hw_l4_drop_pps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sw_l3_drop_pps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hw_l3_drop_pps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_pps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"in_blacklist": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"in_hardware": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hardware_index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"total_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"v4_netmask": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6DdosProtectionIpEntriesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DdosProtectionIpEntriesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DdosProtectionIpEntriesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6DdosProtectionIpEntriesOperOper := setObjectCgnv6DdosProtectionIpEntriesOperOper(res)
		d.Set("oper", Cgnv6DdosProtectionIpEntriesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6DdosProtectionIpEntriesOperOper(ret edpt.DataCgnv6DdosProtectionIpEntriesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_ip_entries_list": setSliceCgnv6DdosProtectionIpEntriesOperOperDdosIpEntriesList(ret.DtCgnv6DdosProtectionIpEntriesOper.Oper.DdosIpEntriesList),
			"total_entries":        ret.DtCgnv6DdosProtectionIpEntriesOper.Oper.TotalEntries,
			"all":                  ret.DtCgnv6DdosProtectionIpEntriesOper.Oper.All,
			"nat_pool":             ret.DtCgnv6DdosProtectionIpEntriesOper.Oper.NatPool,
			"v4_netmask":           ret.DtCgnv6DdosProtectionIpEntriesOper.Oper.V4Netmask,
		},
	}
}

func setSliceCgnv6DdosProtectionIpEntriesOperOperDdosIpEntriesList(d []edpt.Cgnv6DdosProtectionIpEntriesOperOperDdosIpEntriesList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["v4_address"] = item.V4Address
		in["expiration"] = item.Expiration
		in["hints"] = item.Hints
		in["l4_entries_count"] = item.L4EntriesCount
		in["sw_receive_pps"] = item.SwReceivePps
		in["sw_l4_drop_pps"] = item.SwL4DropPps
		in["hw_l4_drop_pps"] = item.HwL4DropPps
		in["sw_l3_drop_pps"] = item.SwL3DropPps
		in["hw_l3_drop_pps"] = item.HwL3DropPps
		in["total_pps"] = item.TotalPps
		in["in_blacklist"] = item.InBlacklist
		in["in_hardware"] = item.InHardware
		in["hardware_index"] = item.HardwareIndex
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6DdosProtectionIpEntriesOperOper(d []interface{}) edpt.Cgnv6DdosProtectionIpEntriesOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6DdosProtectionIpEntriesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DdosIpEntriesList = getSliceCgnv6DdosProtectionIpEntriesOperOperDdosIpEntriesList(in["ddos_ip_entries_list"].([]interface{}))
		ret.TotalEntries = in["total_entries"].(int)
		ret.All = in["all"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.V4Netmask = in["v4_netmask"].(string)
	}
	return ret
}

func getSliceCgnv6DdosProtectionIpEntriesOperOperDdosIpEntriesList(d []interface{}) []edpt.Cgnv6DdosProtectionIpEntriesOperOperDdosIpEntriesList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6DdosProtectionIpEntriesOperOperDdosIpEntriesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6DdosProtectionIpEntriesOperOperDdosIpEntriesList
		oi.V4Address = in["v4_address"].(string)
		oi.Expiration = in["expiration"].(int)
		oi.Hints = in["hints"].(string)
		oi.L4EntriesCount = in["l4_entries_count"].(int)
		oi.SwReceivePps = in["sw_receive_pps"].(int)
		oi.SwL4DropPps = in["sw_l4_drop_pps"].(int)
		oi.HwL4DropPps = in["hw_l4_drop_pps"].(int)
		oi.SwL3DropPps = in["sw_l3_drop_pps"].(int)
		oi.HwL3DropPps = in["hw_l3_drop_pps"].(int)
		oi.TotalPps = in["total_pps"].(int)
		oi.InBlacklist = in["in_blacklist"].(int)
		oi.InHardware = in["in_hardware"].(int)
		oi.HardwareIndex = in["hardware_index"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6DdosProtectionIpEntriesOper(d *schema.ResourceData) edpt.Cgnv6DdosProtectionIpEntriesOper {
	var ret edpt.Cgnv6DdosProtectionIpEntriesOper

	ret.Oper = getObjectCgnv6DdosProtectionIpEntriesOperOper(d.Get("oper").([]interface{}))
	return ret
}
