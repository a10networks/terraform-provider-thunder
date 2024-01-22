package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityDebugFilesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_debug_files_oper`: Operational Status for the object debug-files\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityDebugFilesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"debug_file_name_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entity_key": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "",
									},
									"flat_oid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv4_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipv6_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"protocol": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"debug_file_name": {
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

func resourceVisibilityDebugFilesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityDebugFilesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityDebugFilesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityDebugFilesOperOper := setObjectVisibilityDebugFilesOperOper(res)
		d.Set("oper", VisibilityDebugFilesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityDebugFilesOperOper(ret edpt.DataVisibilityDebugFilesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"debug_file_name_list": setSliceVisibilityDebugFilesOperOperDebugFileNameList(ret.DtVisibilityDebugFilesOper.Oper.DebugFileNameList),
		},
	}
}

func setSliceVisibilityDebugFilesOperOperDebugFileNameList(d []edpt.VisibilityDebugFilesOperOperDebugFileNameList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entity_key"] = item.EntityKey
		//omit uuid
		in["flat_oid"] = item.FlatOid
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv6_addr"] = item.Ipv6Addr
		in["protocol"] = item.Protocol
		in["port"] = item.Port
		in["debug_file_name"] = item.DebugFileName
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityDebugFilesOperOper(d []interface{}) edpt.VisibilityDebugFilesOperOper {

	count1 := len(d)
	var ret edpt.VisibilityDebugFilesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DebugFileNameList = getSliceVisibilityDebugFilesOperOperDebugFileNameList(in["debug_file_name_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityDebugFilesOperOperDebugFileNameList(d []interface{}) []edpt.VisibilityDebugFilesOperOperDebugFileNameList {

	count1 := len(d)
	ret := make([]edpt.VisibilityDebugFilesOperOperDebugFileNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityDebugFilesOperOperDebugFileNameList
		oi.EntityKey = in["entity_key"].(string)
		//omit uuid
		oi.FlatOid = in["flat_oid"].(int)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.Port = in["port"].(int)
		oi.DebugFileName = in["debug_file_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityDebugFilesOper(d *schema.ResourceData) edpt.VisibilityDebugFilesOper {
	var ret edpt.VisibilityDebugFilesOper

	ret.Oper = getObjectVisibilityDebugFilesOperOper(d.Get("oper").([]interface{}))
	return ret
}
