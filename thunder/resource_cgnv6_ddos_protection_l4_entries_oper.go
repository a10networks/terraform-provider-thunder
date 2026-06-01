package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DdosProtectionL4EntriesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_ddos_protection_l4_entries_oper`: Operational Status for the object l4-entries\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6DdosProtectionL4EntriesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ddos_l4_entries_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"v4_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"l4_protocol": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"in_hardware": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hardware_index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"expiration": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"is_deleted": {
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
						"not_in_hardware": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6DdosProtectionL4EntriesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DdosProtectionL4EntriesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DdosProtectionL4EntriesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6DdosProtectionL4EntriesOperOper := setObjectCgnv6DdosProtectionL4EntriesOperOper(res)
		d.Set("oper", Cgnv6DdosProtectionL4EntriesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6DdosProtectionL4EntriesOperOper(ret edpt.DataCgnv6DdosProtectionL4EntriesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_l4_entries_list": setSliceCgnv6DdosProtectionL4EntriesOperOperDdosL4EntriesList(ret.DtCgnv6DdosProtectionL4EntriesOper.Oper.DdosL4EntriesList),
			"total_entries":        ret.DtCgnv6DdosProtectionL4EntriesOper.Oper.TotalEntries,
			"all":                  ret.DtCgnv6DdosProtectionL4EntriesOper.Oper.All,
			"nat_pool":             ret.DtCgnv6DdosProtectionL4EntriesOper.Oper.NatPool,
			"v4_netmask":           ret.DtCgnv6DdosProtectionL4EntriesOper.Oper.V4Netmask,
			"not_in_hardware":      ret.DtCgnv6DdosProtectionL4EntriesOper.Oper.NotInHardware,
		},
	}
}

func setSliceCgnv6DdosProtectionL4EntriesOperOperDdosL4EntriesList(d []edpt.Cgnv6DdosProtectionL4EntriesOperOperDdosL4EntriesList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["v4_address"] = item.V4Address
		in["l4_protocol"] = item.L4Protocol
		in["port"] = item.Port
		in["in_hardware"] = item.InHardware
		in["hardware_index"] = item.HardwareIndex
		in["pps"] = item.Pps
		in["expiration"] = item.Expiration
		in["is_deleted"] = item.IsDeleted
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6DdosProtectionL4EntriesOperOper(d []interface{}) edpt.Cgnv6DdosProtectionL4EntriesOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6DdosProtectionL4EntriesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DdosL4EntriesList = getSliceCgnv6DdosProtectionL4EntriesOperOperDdosL4EntriesList(in["ddos_l4_entries_list"].([]interface{}))
		ret.TotalEntries = in["total_entries"].(int)
		ret.All = in["all"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.V4Netmask = in["v4_netmask"].(string)
		ret.NotInHardware = in["not_in_hardware"].(int)
	}
	return ret
}

func getSliceCgnv6DdosProtectionL4EntriesOperOperDdosL4EntriesList(d []interface{}) []edpt.Cgnv6DdosProtectionL4EntriesOperOperDdosL4EntriesList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6DdosProtectionL4EntriesOperOperDdosL4EntriesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6DdosProtectionL4EntriesOperOperDdosL4EntriesList
		oi.V4Address = in["v4_address"].(string)
		oi.L4Protocol = in["l4_protocol"].(string)
		oi.Port = in["port"].(int)
		oi.InHardware = in["in_hardware"].(int)
		oi.HardwareIndex = in["hardware_index"].(int)
		oi.Pps = in["pps"].(int)
		oi.Expiration = in["expiration"].(int)
		oi.IsDeleted = in["is_deleted"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6DdosProtectionL4EntriesOper(d *schema.ResourceData) edpt.Cgnv6DdosProtectionL4EntriesOper {
	var ret edpt.Cgnv6DdosProtectionL4EntriesOper

	ret.Oper = getObjectCgnv6DdosProtectionL4EntriesOperOper(d.Get("oper").([]interface{}))
	return ret
}
