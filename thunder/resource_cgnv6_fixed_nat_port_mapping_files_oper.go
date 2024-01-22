package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatPortMappingFilesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_port_mapping_files_oper`: Operational Status for the object port-mapping-files\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatPortMappingFilesOperRead,

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
									"write_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"contain": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"contain_case_sensitive": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6FixedNatPortMappingFilesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatPortMappingFilesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatPortMappingFilesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatPortMappingFilesOperOper := setObjectCgnv6FixedNatPortMappingFilesOperOper(res)
		d.Set("oper", Cgnv6FixedNatPortMappingFilesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatPortMappingFilesOperOper(ret edpt.DataCgnv6FixedNatPortMappingFilesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"file_list":              setSliceCgnv6FixedNatPortMappingFilesOperOperFileList(ret.DtCgnv6FixedNatPortMappingFilesOper.Oper.FileList),
			"type":                   ret.DtCgnv6FixedNatPortMappingFilesOper.Oper.Type,
			"contain":                ret.DtCgnv6FixedNatPortMappingFilesOper.Oper.Contain,
			"contain_case_sensitive": ret.DtCgnv6FixedNatPortMappingFilesOper.Oper.ContainCaseSensitive,
		},
	}
}

func setSliceCgnv6FixedNatPortMappingFilesOperOperFileList(d []edpt.Cgnv6FixedNatPortMappingFilesOperOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["file_name"] = item.FileName
		in["write_status"] = item.WriteStatus
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6FixedNatPortMappingFilesOperOper(d []interface{}) edpt.Cgnv6FixedNatPortMappingFilesOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatPortMappingFilesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceCgnv6FixedNatPortMappingFilesOperOperFileList(in["file_list"].([]interface{}))
		ret.Type = in["type"].(string)
		ret.Contain = in["contain"].(string)
		ret.ContainCaseSensitive = in["contain_case_sensitive"].(string)
	}
	return ret
}

func getSliceCgnv6FixedNatPortMappingFilesOperOperFileList(d []interface{}) []edpt.Cgnv6FixedNatPortMappingFilesOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6FixedNatPortMappingFilesOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6FixedNatPortMappingFilesOperOperFileList
		oi.FileName = in["file_name"].(string)
		oi.WriteStatus = in["write_status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatPortMappingFilesOper(d *schema.ResourceData) edpt.Cgnv6FixedNatPortMappingFilesOper {
	var ret edpt.Cgnv6FixedNatPortMappingFilesOper

	ret.Oper = getObjectCgnv6FixedNatPortMappingFilesOperOper(d.Get("oper").([]interface{}))
	return ret
}
