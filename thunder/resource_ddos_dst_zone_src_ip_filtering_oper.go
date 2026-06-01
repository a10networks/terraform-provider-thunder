package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneSrcIpFilteringOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_src_ip_filtering_oper`: Operational Status for the object src-ip-filtering\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneSrcIpFilteringOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}

func resourceDdosDstZoneSrcIpFilteringOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcIpFilteringOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcIpFilteringOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneSrcIpFilteringOperOper := setObjectDdosDstZoneSrcIpFilteringOperOper(res)
		d.Set("oper", DdosDstZoneSrcIpFilteringOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneSrcIpFilteringOperOper(ret edpt.DataDdosDstZoneSrcIpFilteringOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"class_list": setSliceDdosDstZoneSrcIpFilteringOperOperClassList(ret.DtDdosDstZoneSrcIpFilteringOper.Oper.ClassList),
		},
	}
}

func setSliceDdosDstZoneSrcIpFilteringOperOperClassList(d []edpt.DdosDstZoneSrcIpFilteringOperOperClassList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["hit"] = item.Hit
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZoneSrcIpFilteringOperOper(d []interface{}) edpt.DdosDstZoneSrcIpFilteringOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcIpFilteringOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassList = getSliceDdosDstZoneSrcIpFilteringOperOperClassList(in["class_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneSrcIpFilteringOperOperClassList(d []interface{}) []edpt.DdosDstZoneSrcIpFilteringOperOperClassList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcIpFilteringOperOperClassList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcIpFilteringOperOperClassList
		oi.Name = in["name"].(string)
		oi.Hit = in["hit"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneSrcIpFilteringOper(d *schema.ResourceData) edpt.DdosDstZoneSrcIpFilteringOper {
	var ret edpt.DdosDstZoneSrcIpFilteringOper

	ret.Oper = getObjectDdosDstZoneSrcIpFilteringOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
