package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPortScanDetectionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_port_scan_detection_oper`: Operational Status for the object port-scan-detection\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityPortScanDetectionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_ip_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"scanned_ports": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"port_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"protocol": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"scanned_time": {
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
		},
	}
}

func resourceVisibilityPortScanDetectionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPortScanDetectionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPortScanDetectionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityPortScanDetectionOperOper := setObjectVisibilityPortScanDetectionOperOper(res)
		d.Set("oper", VisibilityPortScanDetectionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityPortScanDetectionOperOper(ret edpt.DataVisibilityPortScanDetectionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"src_ip_list": setSliceVisibilityPortScanDetectionOperOperSrcIpList(ret.DtVisibilityPortScanDetectionOper.Oper.SrcIpList),
		},
	}
}

func setSliceVisibilityPortScanDetectionOperOperSrcIpList(d []edpt.VisibilityPortScanDetectionOperOperSrcIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_ip_addr"] = item.SrcIpAddr
		in["scanned_ports"] = item.ScannedPorts
		in["port_list"] = setSliceVisibilityPortScanDetectionOperOperSrcIpListPortList(item.PortList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityPortScanDetectionOperOperSrcIpListPortList(d []edpt.VisibilityPortScanDetectionOperOperSrcIpListPortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port"] = item.Port
		in["protocol"] = item.Protocol
		in["scanned_time"] = item.Scanned_time
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityPortScanDetectionOperOper(d []interface{}) edpt.VisibilityPortScanDetectionOperOper {

	count1 := len(d)
	var ret edpt.VisibilityPortScanDetectionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcIpList = getSliceVisibilityPortScanDetectionOperOperSrcIpList(in["src_ip_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityPortScanDetectionOperOperSrcIpList(d []interface{}) []edpt.VisibilityPortScanDetectionOperOperSrcIpList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPortScanDetectionOperOperSrcIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPortScanDetectionOperOperSrcIpList
		oi.SrcIpAddr = in["src_ip_addr"].(string)
		oi.ScannedPorts = in["scanned_ports"].(int)
		oi.PortList = getSliceVisibilityPortScanDetectionOperOperSrcIpListPortList(in["port_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityPortScanDetectionOperOperSrcIpListPortList(d []interface{}) []edpt.VisibilityPortScanDetectionOperOperSrcIpListPortList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPortScanDetectionOperOperSrcIpListPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPortScanDetectionOperOperSrcIpListPortList
		oi.Port = in["port"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Scanned_time = in["scanned_time"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityPortScanDetectionOper(d *schema.ResourceData) edpt.VisibilityPortScanDetectionOper {
	var ret edpt.VisibilityPortScanDetectionOper

	ret.Oper = getObjectVisibilityPortScanDetectionOperOper(d.Get("oper").([]interface{}))
	return ret
}
