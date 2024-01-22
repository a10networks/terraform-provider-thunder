package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6AllBindingTablesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lw_4o6_all_binding_tables_oper`: Operational Status for the object all-binding-tables\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Lw4o6AllBindingTablesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
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

func resourceCgnv6Lw4o6AllBindingTablesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6AllBindingTablesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6AllBindingTablesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Lw4o6AllBindingTablesOperOper := setObjectCgnv6Lw4o6AllBindingTablesOperOper(res)
		d.Set("oper", Cgnv6Lw4o6AllBindingTablesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Lw4o6AllBindingTablesOperOper(ret edpt.DataCgnv6Lw4o6AllBindingTablesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_list": setSliceCgnv6Lw4o6AllBindingTablesOperOperEntryList(ret.DtCgnv6Lw4o6AllBindingTablesOper.Oper.EntryList),
		},
	}
}

func setSliceCgnv6Lw4o6AllBindingTablesOperOperEntryList(d []edpt.Cgnv6Lw4o6AllBindingTablesOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6Lw4o6AllBindingTablesOperOper(d []interface{}) edpt.Cgnv6Lw4o6AllBindingTablesOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6Lw4o6AllBindingTablesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceCgnv6Lw4o6AllBindingTablesOperOperEntryList(in["entry_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6Lw4o6AllBindingTablesOperOperEntryList(d []interface{}) []edpt.Cgnv6Lw4o6AllBindingTablesOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Lw4o6AllBindingTablesOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Lw4o6AllBindingTablesOperOperEntryList
		oi.Name = in["name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Lw4o6AllBindingTablesOper(d *schema.ResourceData) edpt.Cgnv6Lw4o6AllBindingTablesOper {
	var ret edpt.Cgnv6Lw4o6AllBindingTablesOper

	ret.Oper = getObjectCgnv6Lw4o6AllBindingTablesOperOper(d.Get("oper").([]interface{}))
	return ret
}
