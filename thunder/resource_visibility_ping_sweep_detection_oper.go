package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPingSweepDetectionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_ping_sweep_detection_oper`: Operational Status for the object ping-sweep-detection\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityPingSweepDetectionOperRead,

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
									"scanned_ips": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ip_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip": {
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

func resourceVisibilityPingSweepDetectionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPingSweepDetectionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPingSweepDetectionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityPingSweepDetectionOperOper := setObjectVisibilityPingSweepDetectionOperOper(res)
		d.Set("oper", VisibilityPingSweepDetectionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityPingSweepDetectionOperOper(ret edpt.DataVisibilityPingSweepDetectionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"src_ip_list": setSliceVisibilityPingSweepDetectionOperOperSrcIpList(ret.DtVisibilityPingSweepDetectionOper.Oper.SrcIpList),
		},
	}
}

func setSliceVisibilityPingSweepDetectionOperOperSrcIpList(d []edpt.VisibilityPingSweepDetectionOperOperSrcIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_ip_addr"] = item.SrcIpAddr
		in["scanned_ips"] = item.ScannedIps
		in["ip_list"] = setSliceVisibilityPingSweepDetectionOperOperSrcIpListIpList(item.IpList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityPingSweepDetectionOperOperSrcIpListIpList(d []edpt.VisibilityPingSweepDetectionOperOperSrcIpListIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip"] = item.Ip
		in["scanned_time"] = item.Scanned_time
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityPingSweepDetectionOperOper(d []interface{}) edpt.VisibilityPingSweepDetectionOperOper {

	count1 := len(d)
	var ret edpt.VisibilityPingSweepDetectionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcIpList = getSliceVisibilityPingSweepDetectionOperOperSrcIpList(in["src_ip_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityPingSweepDetectionOperOperSrcIpList(d []interface{}) []edpt.VisibilityPingSweepDetectionOperOperSrcIpList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPingSweepDetectionOperOperSrcIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPingSweepDetectionOperOperSrcIpList
		oi.SrcIpAddr = in["src_ip_addr"].(string)
		oi.ScannedIps = in["scanned_ips"].(int)
		oi.IpList = getSliceVisibilityPingSweepDetectionOperOperSrcIpListIpList(in["ip_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityPingSweepDetectionOperOperSrcIpListIpList(d []interface{}) []edpt.VisibilityPingSweepDetectionOperOperSrcIpListIpList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPingSweepDetectionOperOperSrcIpListIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPingSweepDetectionOperOperSrcIpListIpList
		oi.Ip = in["ip"].(string)
		oi.Scanned_time = in["scanned_time"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityPingSweepDetectionOper(d *schema.ResourceData) edpt.VisibilityPingSweepDetectionOper {
	var ret edpt.VisibilityPingSweepDetectionOper

	ret.Oper = getObjectVisibilityPingSweepDetectionOperOper(d.Get("oper").([]interface{}))
	return ret
}
