package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6ValidationLogFilesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lw_4o6_validation_log_files_oper`: Operational Status for the object validation-log-files\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Lw4o6ValidationLogFilesOperRead,

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
						"entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6Lw4o6ValidationLogFilesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6ValidationLogFilesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6ValidationLogFilesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Lw4o6ValidationLogFilesOperOper := setObjectCgnv6Lw4o6ValidationLogFilesOperOper(res)
		d.Set("oper", Cgnv6Lw4o6ValidationLogFilesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Lw4o6ValidationLogFilesOperOper(ret edpt.DataCgnv6Lw4o6ValidationLogFilesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_list":  setSliceCgnv6Lw4o6ValidationLogFilesOperOperEntryList(ret.DtCgnv6Lw4o6ValidationLogFilesOper.Oper.EntryList),
			"entry_count": ret.DtCgnv6Lw4o6ValidationLogFilesOper.Oper.EntryCount,
		},
	}
}

func setSliceCgnv6Lw4o6ValidationLogFilesOperOperEntryList(d []edpt.Cgnv6Lw4o6ValidationLogFilesOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6Lw4o6ValidationLogFilesOperOper(d []interface{}) edpt.Cgnv6Lw4o6ValidationLogFilesOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6Lw4o6ValidationLogFilesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceCgnv6Lw4o6ValidationLogFilesOperOperEntryList(in["entry_list"].([]interface{}))
		ret.EntryCount = in["entry_count"].(int)
	}
	return ret
}

func getSliceCgnv6Lw4o6ValidationLogFilesOperOperEntryList(d []interface{}) []edpt.Cgnv6Lw4o6ValidationLogFilesOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Lw4o6ValidationLogFilesOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Lw4o6ValidationLogFilesOperOperEntryList
		oi.Name = in["name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Lw4o6ValidationLogFilesOper(d *schema.ResourceData) edpt.Cgnv6Lw4o6ValidationLogFilesOper {
	var ret edpt.Cgnv6Lw4o6ValidationLogFilesOper

	ret.Oper = getObjectCgnv6Lw4o6ValidationLogFilesOperOper(d.Get("oper").([]interface{}))
	return ret
}
