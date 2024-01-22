package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6BindingTableFilesStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lw_4o6_binding_table_files_status_oper`: Operational Status for the object binding-table-files-status\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Lw4o6BindingTableFilesStatusOperRead,

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
									"active": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"modified": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6Lw4o6BindingTableFilesStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableFilesStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableFilesStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Lw4o6BindingTableFilesStatusOperOper := setObjectCgnv6Lw4o6BindingTableFilesStatusOperOper(res)
		d.Set("oper", Cgnv6Lw4o6BindingTableFilesStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Lw4o6BindingTableFilesStatusOperOper(ret edpt.DataCgnv6Lw4o6BindingTableFilesStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_list":  setSliceCgnv6Lw4o6BindingTableFilesStatusOperOperEntryList(ret.DtCgnv6Lw4o6BindingTableFilesStatusOper.Oper.EntryList),
			"entry_count": ret.DtCgnv6Lw4o6BindingTableFilesStatusOper.Oper.EntryCount,
		},
	}
}

func setSliceCgnv6Lw4o6BindingTableFilesStatusOperOperEntryList(d []edpt.Cgnv6Lw4o6BindingTableFilesStatusOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["active"] = item.Active
		in["modified"] = item.Modified
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6Lw4o6BindingTableFilesStatusOperOper(d []interface{}) edpt.Cgnv6Lw4o6BindingTableFilesStatusOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6Lw4o6BindingTableFilesStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceCgnv6Lw4o6BindingTableFilesStatusOperOperEntryList(in["entry_list"].([]interface{}))
		ret.EntryCount = in["entry_count"].(int)
	}
	return ret
}

func getSliceCgnv6Lw4o6BindingTableFilesStatusOperOperEntryList(d []interface{}) []edpt.Cgnv6Lw4o6BindingTableFilesStatusOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Lw4o6BindingTableFilesStatusOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Lw4o6BindingTableFilesStatusOperOperEntryList
		oi.Name = in["name"].(string)
		oi.Active = in["active"].(string)
		oi.Modified = in["modified"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Lw4o6BindingTableFilesStatusOper(d *schema.ResourceData) edpt.Cgnv6Lw4o6BindingTableFilesStatusOper {
	var ret edpt.Cgnv6Lw4o6BindingTableFilesStatusOper

	ret.Oper = getObjectCgnv6Lw4o6BindingTableFilesStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
