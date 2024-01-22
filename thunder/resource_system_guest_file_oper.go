package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGuestFileOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_guest_file_oper`: Operational Status for the object guest-file\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemGuestFileOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"update_time": {
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

func resourceSystemGuestFileOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGuestFileOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGuestFileOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemGuestFileOperOper := setObjectSystemGuestFileOperOper(res)
		d.Set("oper", SystemGuestFileOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemGuestFileOperOper(ret edpt.DataSystemGuestFileOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"file_list": setSliceSystemGuestFileOperOperFileList(ret.DtSystemGuestFileOper.Oper.FileList),
		},
	}
}

func setSliceSystemGuestFileOperOperFileList(d []edpt.SystemGuestFileOperOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["file_name"] = item.FileName
		in["size"] = item.Size
		in["update_time"] = item.UpdateTime
		result = append(result, in)
	}
	return result
}

func getObjectSystemGuestFileOperOper(d []interface{}) edpt.SystemGuestFileOperOper {

	count1 := len(d)
	var ret edpt.SystemGuestFileOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceSystemGuestFileOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemGuestFileOperOperFileList(d []interface{}) []edpt.SystemGuestFileOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.SystemGuestFileOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGuestFileOperOperFileList
		oi.FileName = in["file_name"].(string)
		oi.Size = in["size"].(int)
		oi.UpdateTime = in["update_time"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemGuestFileOper(d *schema.ResourceData) edpt.SystemGuestFileOper {
	var ret edpt.SystemGuestFileOper

	ret.Oper = getObjectSystemGuestFileOperOper(d.Get("oper").([]interface{}))
	return ret
}
