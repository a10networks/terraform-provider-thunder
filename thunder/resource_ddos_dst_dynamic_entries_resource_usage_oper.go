package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstDynamicEntriesResourceUsageOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_dynamic_entries_resource_usage_oper`: Operational Status for the object dynamic-entries-resource-usage\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstDynamicEntriesResourceUsageOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_entry_ip_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dst_entry_ip_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dst_entry_ip_remaining": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"conn_total_ip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dst_entry_ipv6_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dst_entry_ipv6_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dst_entry_ipv6_remaining": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"conn_total_ipv6": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDdosDstDynamicEntriesResourceUsageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntriesResourceUsageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntriesResourceUsageOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstDynamicEntriesResourceUsageOperOper := setObjectDdosDstDynamicEntriesResourceUsageOperOper(res)
		d.Set("oper", DdosDstDynamicEntriesResourceUsageOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstDynamicEntriesResourceUsageOperOper(ret edpt.DataDdosDstDynamicEntriesResourceUsageOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dst_entry_ip_limit":       ret.DtDdosDstDynamicEntriesResourceUsageOper.Oper.Dst_entry_ip_limit,
			"dst_entry_ip_allocated":   ret.DtDdosDstDynamicEntriesResourceUsageOper.Oper.Dst_entry_ip_allocated,
			"dst_entry_ip_remaining":   ret.DtDdosDstDynamicEntriesResourceUsageOper.Oper.Dst_entry_ip_remaining,
			"conn_total_ip":            ret.DtDdosDstDynamicEntriesResourceUsageOper.Oper.Conn_total_ip,
			"dst_entry_ipv6_limit":     ret.DtDdosDstDynamicEntriesResourceUsageOper.Oper.Dst_entry_ipv6_limit,
			"dst_entry_ipv6_allocated": ret.DtDdosDstDynamicEntriesResourceUsageOper.Oper.Dst_entry_ipv6_allocated,
			"dst_entry_ipv6_remaining": ret.DtDdosDstDynamicEntriesResourceUsageOper.Oper.Dst_entry_ipv6_remaining,
			"conn_total_ipv6":          ret.DtDdosDstDynamicEntriesResourceUsageOper.Oper.Conn_total_ipv6,
		},
	}
}

func getObjectDdosDstDynamicEntriesResourceUsageOperOper(d []interface{}) edpt.DdosDstDynamicEntriesResourceUsageOperOper {

	count1 := len(d)
	var ret edpt.DdosDstDynamicEntriesResourceUsageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dst_entry_ip_limit = in["dst_entry_ip_limit"].(int)
		ret.Dst_entry_ip_allocated = in["dst_entry_ip_allocated"].(int)
		ret.Dst_entry_ip_remaining = in["dst_entry_ip_remaining"].(string)
		ret.Conn_total_ip = in["conn_total_ip"].(int)
		ret.Dst_entry_ipv6_limit = in["dst_entry_ipv6_limit"].(int)
		ret.Dst_entry_ipv6_allocated = in["dst_entry_ipv6_allocated"].(int)
		ret.Dst_entry_ipv6_remaining = in["dst_entry_ipv6_remaining"].(string)
		ret.Conn_total_ipv6 = in["conn_total_ipv6"].(int)
	}
	return ret
}

func dataToEndpointDdosDstDynamicEntriesResourceUsageOper(d *schema.ResourceData) edpt.DdosDstDynamicEntriesResourceUsageOper {
	var ret edpt.DdosDstDynamicEntriesResourceUsageOper

	ret.Oper = getObjectDdosDstDynamicEntriesResourceUsageOperOper(d.Get("oper").([]interface{}))
	return ret
}
