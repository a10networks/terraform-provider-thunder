package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_packet_capture_oper`: Operational Status for the object packet-capture\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityPacketCaptureOperRead,

		Schema: map[string]*schema.Schema{
			"automated_captures": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"capture_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"automated_capture_config": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"automated_capture_config_line": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packet_capture_files": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"memory_usage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_number_of_files": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"packet_capture_file_name_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"packet_capture_file_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"file_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"last_modified": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"total_memory": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"used_memory": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityPacketCaptureOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityPacketCaptureOperAutomatedCaptures := setObjectVisibilityPacketCaptureOperAutomatedCaptures(res)
		d.Set("automated_captures", VisibilityPacketCaptureOperAutomatedCaptures)
		VisibilityPacketCaptureOperOper := setObjectVisibilityPacketCaptureOperOper(res)
		d.Set("oper", VisibilityPacketCaptureOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityPacketCaptureOperAutomatedCaptures(ret edpt.DataVisibilityPacketCaptureOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVisibilityPacketCaptureOperAutomatedCapturesOper(ret.DtVisibilityPacketCaptureOper.AutomatedCaptures.Oper),
		},
	}
}

func setObjectVisibilityPacketCaptureOperAutomatedCapturesOper(d edpt.VisibilityPacketCaptureOperAutomatedCapturesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["capture_name"] = d.CaptureName
	in["automated_capture_config"] = setSliceVisibilityPacketCaptureOperAutomatedCapturesOperAutomatedCaptureConfig(d.AutomatedCaptureConfig)
	result = append(result, in)
	return result
}

func setSliceVisibilityPacketCaptureOperAutomatedCapturesOperAutomatedCaptureConfig(d []edpt.VisibilityPacketCaptureOperAutomatedCapturesOperAutomatedCaptureConfig) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["automated_capture_config_line"] = item.AutomatedCaptureConfigLine
		result = append(result, in)
	}
	return result
}

func setObjectVisibilityPacketCaptureOperOper(ret edpt.DataVisibilityPacketCaptureOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packet_capture_files":          ret.DtVisibilityPacketCaptureOper.Oper.PacketCaptureFiles,
			"memory_usage":                  ret.DtVisibilityPacketCaptureOper.Oper.MemoryUsage,
			"total_number_of_files":         ret.DtVisibilityPacketCaptureOper.Oper.TotalNumberOfFiles,
			"packet_capture_file_name_list": setSliceVisibilityPacketCaptureOperOperPacketCaptureFileNameList(ret.DtVisibilityPacketCaptureOper.Oper.PacketCaptureFileNameList),
			"total_memory":                  ret.DtVisibilityPacketCaptureOper.Oper.TotalMemory,
			"used_memory":                   ret.DtVisibilityPacketCaptureOper.Oper.UsedMemory,
		},
	}
}

func setSliceVisibilityPacketCaptureOperOperPacketCaptureFileNameList(d []edpt.VisibilityPacketCaptureOperOperPacketCaptureFileNameList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["packet_capture_file_name"] = item.PacketCaptureFileName
		in["file_size"] = item.FileSize
		in["last_modified"] = item.LastModified
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityPacketCaptureOperAutomatedCaptures(d []interface{}) edpt.VisibilityPacketCaptureOperAutomatedCaptures {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureOperAutomatedCaptures
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityPacketCaptureOperAutomatedCapturesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureOperAutomatedCapturesOper(d []interface{}) edpt.VisibilityPacketCaptureOperAutomatedCapturesOper {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureOperAutomatedCapturesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CaptureName = in["capture_name"].(string)
		ret.AutomatedCaptureConfig = getSliceVisibilityPacketCaptureOperAutomatedCapturesOperAutomatedCaptureConfig(in["automated_capture_config"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityPacketCaptureOperAutomatedCapturesOperAutomatedCaptureConfig(d []interface{}) []edpt.VisibilityPacketCaptureOperAutomatedCapturesOperAutomatedCaptureConfig {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureOperAutomatedCapturesOperAutomatedCaptureConfig, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureOperAutomatedCapturesOperAutomatedCaptureConfig
		oi.AutomatedCaptureConfigLine = in["automated_capture_config_line"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureOperOper(d []interface{}) edpt.VisibilityPacketCaptureOperOper {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PacketCaptureFiles = in["packet_capture_files"].(int)
		ret.MemoryUsage = in["memory_usage"].(int)
		ret.TotalNumberOfFiles = in["total_number_of_files"].(int)
		ret.PacketCaptureFileNameList = getSliceVisibilityPacketCaptureOperOperPacketCaptureFileNameList(in["packet_capture_file_name_list"].([]interface{}))
		ret.TotalMemory = in["total_memory"].(string)
		ret.UsedMemory = in["used_memory"].(string)
	}
	return ret
}

func getSliceVisibilityPacketCaptureOperOperPacketCaptureFileNameList(d []interface{}) []edpt.VisibilityPacketCaptureOperOperPacketCaptureFileNameList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureOperOperPacketCaptureFileNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureOperOperPacketCaptureFileNameList
		oi.PacketCaptureFileName = in["packet_capture_file_name"].(string)
		oi.FileSize = in["file_size"].(int)
		oi.LastModified = in["last_modified"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureOper(d *schema.ResourceData) edpt.VisibilityPacketCaptureOper {
	var ret edpt.VisibilityPacketCaptureOper

	ret.AutomatedCaptures = getObjectVisibilityPacketCaptureOperAutomatedCaptures(d.Get("automated_captures").([]interface{}))

	ret.Oper = getObjectVisibilityPacketCaptureOperOper(d.Get("oper").([]interface{}))
	return ret
}
