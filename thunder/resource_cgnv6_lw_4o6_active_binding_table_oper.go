package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6ActiveBindingTableOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lw_4o6_active_binding_table_oper`: Operational Status for the object active-binding-table\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Lw4o6ActiveBindingTableOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tunnel_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"nat_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_start": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"port_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fwd_match_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rev_match_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aftr_tunnel_address": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceCgnv6Lw4o6ActiveBindingTableOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6ActiveBindingTableOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6ActiveBindingTableOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Lw4o6ActiveBindingTableOperOper := setObjectCgnv6Lw4o6ActiveBindingTableOperOper(res)
		d.Set("oper", Cgnv6Lw4o6ActiveBindingTableOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Lw4o6ActiveBindingTableOperOper(ret edpt.DataCgnv6Lw4o6ActiveBindingTableOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_list": setSliceCgnv6Lw4o6ActiveBindingTableOperOperEntryList(ret.DtCgnv6Lw4o6ActiveBindingTableOper.Oper.EntryList),
		},
	}
}

func setSliceCgnv6Lw4o6ActiveBindingTableOperOperEntryList(d []edpt.Cgnv6Lw4o6ActiveBindingTableOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["tunnel_address"] = item.TunnelAddress
		in["nat_address"] = item.NatAddress
		in["port_start"] = item.PortStart
		in["port_end"] = item.PortEnd
		in["fwd_match_count"] = item.FwdMatchCount
		in["rev_match_count"] = item.RevMatchCount
		in["aftr_tunnel_address"] = item.AftrTunnelAddress
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6Lw4o6ActiveBindingTableOperOper(d []interface{}) edpt.Cgnv6Lw4o6ActiveBindingTableOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6Lw4o6ActiveBindingTableOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceCgnv6Lw4o6ActiveBindingTableOperOperEntryList(in["entry_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6Lw4o6ActiveBindingTableOperOperEntryList(d []interface{}) []edpt.Cgnv6Lw4o6ActiveBindingTableOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Lw4o6ActiveBindingTableOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Lw4o6ActiveBindingTableOperOperEntryList
		oi.TunnelAddress = in["tunnel_address"].(string)
		oi.NatAddress = in["nat_address"].(string)
		oi.PortStart = in["port_start"].(int)
		oi.PortEnd = in["port_end"].(int)
		oi.FwdMatchCount = in["fwd_match_count"].(int)
		oi.RevMatchCount = in["rev_match_count"].(int)
		oi.AftrTunnelAddress = in["aftr_tunnel_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Lw4o6ActiveBindingTableOper(d *schema.ResourceData) edpt.Cgnv6Lw4o6ActiveBindingTableOper {
	var ret edpt.Cgnv6Lw4o6ActiveBindingTableOper

	ret.Oper = getObjectCgnv6Lw4o6ActiveBindingTableOperOper(d.Get("oper").([]interface{}))
	return ret
}
