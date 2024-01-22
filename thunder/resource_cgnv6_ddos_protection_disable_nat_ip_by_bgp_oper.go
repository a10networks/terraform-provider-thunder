package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DdosProtectionDisableNatIpByBgpOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_ddos_protection_disable_nat_ip_by_bgp_oper`: Operational Status for the object disable-nat-ip-by-bgp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6DdosProtectionDisableNatIpByBgpOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ddos_disabled_by_bgp_entries_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"v4_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"nat_pool_name": {
										Type: schema.TypeString, Optional: true, Description: "",
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
						"v4_address": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6DdosProtectionDisableNatIpByBgpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DdosProtectionDisableNatIpByBgpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DdosProtectionDisableNatIpByBgpOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6DdosProtectionDisableNatIpByBgpOperOper := setObjectCgnv6DdosProtectionDisableNatIpByBgpOperOper(res)
		d.Set("oper", Cgnv6DdosProtectionDisableNatIpByBgpOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6DdosProtectionDisableNatIpByBgpOperOper(ret edpt.DataCgnv6DdosProtectionDisableNatIpByBgpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_disabled_by_bgp_entries_list": setSliceCgnv6DdosProtectionDisableNatIpByBgpOperOperDdosDisabledByBgpEntriesList(ret.DtCgnv6DdosProtectionDisableNatIpByBgpOper.Oper.DdosDisabledByBgpEntriesList),
			"total_entries":                     ret.DtCgnv6DdosProtectionDisableNatIpByBgpOper.Oper.TotalEntries,
			"all":                               ret.DtCgnv6DdosProtectionDisableNatIpByBgpOper.Oper.All,
			"v4_address":                        ret.DtCgnv6DdosProtectionDisableNatIpByBgpOper.Oper.V4Address,
		},
	}
}

func setSliceCgnv6DdosProtectionDisableNatIpByBgpOperOperDdosDisabledByBgpEntriesList(d []edpt.Cgnv6DdosProtectionDisableNatIpByBgpOperOperDdosDisabledByBgpEntriesList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["v4_address"] = item.V4Address
		in["nat_pool_name"] = item.NatPoolName
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6DdosProtectionDisableNatIpByBgpOperOper(d []interface{}) edpt.Cgnv6DdosProtectionDisableNatIpByBgpOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6DdosProtectionDisableNatIpByBgpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DdosDisabledByBgpEntriesList = getSliceCgnv6DdosProtectionDisableNatIpByBgpOperOperDdosDisabledByBgpEntriesList(in["ddos_disabled_by_bgp_entries_list"].([]interface{}))
		ret.TotalEntries = in["total_entries"].(int)
		ret.All = in["all"].(int)
		ret.V4Address = in["v4_address"].(string)
	}
	return ret
}

func getSliceCgnv6DdosProtectionDisableNatIpByBgpOperOperDdosDisabledByBgpEntriesList(d []interface{}) []edpt.Cgnv6DdosProtectionDisableNatIpByBgpOperOperDdosDisabledByBgpEntriesList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6DdosProtectionDisableNatIpByBgpOperOperDdosDisabledByBgpEntriesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6DdosProtectionDisableNatIpByBgpOperOperDdosDisabledByBgpEntriesList
		oi.V4Address = in["v4_address"].(string)
		oi.NatPoolName = in["nat_pool_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6DdosProtectionDisableNatIpByBgpOper(d *schema.ResourceData) edpt.Cgnv6DdosProtectionDisableNatIpByBgpOper {
	var ret edpt.Cgnv6DdosProtectionDisableNatIpByBgpOper

	ret.Oper = getObjectCgnv6DdosProtectionDisableNatIpByBgpOperOper(d.Get("oper").([]interface{}))
	return ret
}
