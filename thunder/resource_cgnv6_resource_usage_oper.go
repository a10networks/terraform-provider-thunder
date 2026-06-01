package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6ResourceUsageOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_resource_usage_oper`: Operational Status for the object resource-usage\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6ResourceUsageOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lsn_nat_addr_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"lsn_nat_addr_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"lsn_nat_addr_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fixed_nat_ip_addr_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fixed_nat_ip_addr_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fixed_nat_ip_addr_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fixed_nat_inside_user_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fixed_nat_inside_user_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fixed_nat_inside_user_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_table_size_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_table_size_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_table_size_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6ResourceUsageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ResourceUsageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ResourceUsageOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6ResourceUsageOperOper := setObjectCgnv6ResourceUsageOperOper(res)
		d.Set("oper", Cgnv6ResourceUsageOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6ResourceUsageOperOper(ret edpt.DataCgnv6ResourceUsageOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"lsn_nat_addr_count_min":              ret.DtCgnv6ResourceUsageOper.Oper.LsnNatAddrCountMin,
			"lsn_nat_addr_count_max":              ret.DtCgnv6ResourceUsageOper.Oper.LsnNatAddrCountMax,
			"lsn_nat_addr_count_default":          ret.DtCgnv6ResourceUsageOper.Oper.LsnNatAddrCountDefault,
			"fixed_nat_ip_addr_count_min":         ret.DtCgnv6ResourceUsageOper.Oper.FixedNatIpAddrCountMin,
			"fixed_nat_ip_addr_count_max":         ret.DtCgnv6ResourceUsageOper.Oper.FixedNatIpAddrCountMax,
			"fixed_nat_ip_addr_count_default":     ret.DtCgnv6ResourceUsageOper.Oper.FixedNatIpAddrCountDefault,
			"fixed_nat_inside_user_count_min":     ret.DtCgnv6ResourceUsageOper.Oper.FixedNatInsideUserCountMin,
			"fixed_nat_inside_user_count_max":     ret.DtCgnv6ResourceUsageOper.Oper.FixedNatInsideUserCountMax,
			"fixed_nat_inside_user_count_default": ret.DtCgnv6ResourceUsageOper.Oper.FixedNatInsideUserCountDefault,
			"radius_table_size_min":               ret.DtCgnv6ResourceUsageOper.Oper.RadiusTableSizeMin,
			"radius_table_size_max":               ret.DtCgnv6ResourceUsageOper.Oper.RadiusTableSizeMax,
			"radius_table_size_default":           ret.DtCgnv6ResourceUsageOper.Oper.RadiusTableSizeDefault,
		},
	}
}

func getObjectCgnv6ResourceUsageOperOper(d []interface{}) edpt.Cgnv6ResourceUsageOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6ResourceUsageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LsnNatAddrCountMin = in["lsn_nat_addr_count_min"].(int)
		ret.LsnNatAddrCountMax = in["lsn_nat_addr_count_max"].(int)
		ret.LsnNatAddrCountDefault = in["lsn_nat_addr_count_default"].(int)
		ret.FixedNatIpAddrCountMin = in["fixed_nat_ip_addr_count_min"].(int)
		ret.FixedNatIpAddrCountMax = in["fixed_nat_ip_addr_count_max"].(int)
		ret.FixedNatIpAddrCountDefault = in["fixed_nat_ip_addr_count_default"].(int)
		ret.FixedNatInsideUserCountMin = in["fixed_nat_inside_user_count_min"].(int)
		ret.FixedNatInsideUserCountMax = in["fixed_nat_inside_user_count_max"].(int)
		ret.FixedNatInsideUserCountDefault = in["fixed_nat_inside_user_count_default"].(int)
		ret.RadiusTableSizeMin = in["radius_table_size_min"].(int)
		ret.RadiusTableSizeMax = in["radius_table_size_max"].(int)
		ret.RadiusTableSizeDefault = in["radius_table_size_default"].(int)
	}
	return ret
}

func dataToEndpointCgnv6ResourceUsageOper(d *schema.ResourceData) edpt.Cgnv6ResourceUsageOper {
	var ret edpt.Cgnv6ResourceUsageOper

	ret.Oper = getObjectCgnv6ResourceUsageOperOper(d.Get("oper").([]interface{}))
	return ret
}
