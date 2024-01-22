package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcDynamicEntriesResourceUsageOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_src_dynamic_entries_resource_usage_oper`: Operational Status for the object dynamic-entries-resource-usage\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosSrcDynamicEntriesResourceUsageOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_entry_ip_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_entry_ip_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_entry_ip_remaining": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"conn_total_ip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_entry_ipv6_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_entry_ipv6_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_entry_ipv6_remaining": {
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

func resourceDdosSrcDynamicEntriesResourceUsageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcDynamicEntriesResourceUsageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcDynamicEntriesResourceUsageOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosSrcDynamicEntriesResourceUsageOperOper := setObjectDdosSrcDynamicEntriesResourceUsageOperOper(res)
		d.Set("oper", DdosSrcDynamicEntriesResourceUsageOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosSrcDynamicEntriesResourceUsageOperOper(ret edpt.DataDdosSrcDynamicEntriesResourceUsageOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"src_entry_ip_limit":       ret.DtDdosSrcDynamicEntriesResourceUsageOper.Oper.Src_entry_ip_limit,
			"src_entry_ip_allocated":   ret.DtDdosSrcDynamicEntriesResourceUsageOper.Oper.Src_entry_ip_allocated,
			"src_entry_ip_remaining":   ret.DtDdosSrcDynamicEntriesResourceUsageOper.Oper.Src_entry_ip_remaining,
			"conn_total_ip":            ret.DtDdosSrcDynamicEntriesResourceUsageOper.Oper.Conn_total_ip,
			"src_entry_ipv6_limit":     ret.DtDdosSrcDynamicEntriesResourceUsageOper.Oper.Src_entry_ipv6_limit,
			"src_entry_ipv6_allocated": ret.DtDdosSrcDynamicEntriesResourceUsageOper.Oper.Src_entry_ipv6_allocated,
			"src_entry_ipv6_remaining": ret.DtDdosSrcDynamicEntriesResourceUsageOper.Oper.Src_entry_ipv6_remaining,
			"conn_total_ipv6":          ret.DtDdosSrcDynamicEntriesResourceUsageOper.Oper.Conn_total_ipv6,
		},
	}
}

func getObjectDdosSrcDynamicEntriesResourceUsageOperOper(d []interface{}) edpt.DdosSrcDynamicEntriesResourceUsageOperOper {

	count1 := len(d)
	var ret edpt.DdosSrcDynamicEntriesResourceUsageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Src_entry_ip_limit = in["src_entry_ip_limit"].(int)
		ret.Src_entry_ip_allocated = in["src_entry_ip_allocated"].(int)
		ret.Src_entry_ip_remaining = in["src_entry_ip_remaining"].(string)
		ret.Conn_total_ip = in["conn_total_ip"].(int)
		ret.Src_entry_ipv6_limit = in["src_entry_ipv6_limit"].(int)
		ret.Src_entry_ipv6_allocated = in["src_entry_ipv6_allocated"].(int)
		ret.Src_entry_ipv6_remaining = in["src_entry_ipv6_remaining"].(string)
		ret.Conn_total_ipv6 = in["conn_total_ipv6"].(int)
	}
	return ret
}

func dataToEndpointDdosSrcDynamicEntriesResourceUsageOper(d *schema.ResourceData) edpt.DdosSrcDynamicEntriesResourceUsageOper {
	var ret edpt.DdosSrcDynamicEntriesResourceUsageOper

	ret.Oper = getObjectDdosSrcDynamicEntriesResourceUsageOperOper(d.Get("oper").([]interface{}))
	return ret
}
