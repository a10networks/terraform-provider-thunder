package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneDetectionServiceDiscoveryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_detection_service_discovery_oper`: Operational Status for the object service-discovery\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneDetectionServiceDiscoveryOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"discovered_service_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"protocol": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rate": {
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

func resourceDdosDstZoneDetectionServiceDiscoveryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionServiceDiscoveryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionServiceDiscoveryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneDetectionServiceDiscoveryOperOper := setObjectDdosDstZoneDetectionServiceDiscoveryOperOper(res)
		d.Set("oper", DdosDstZoneDetectionServiceDiscoveryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneDetectionServiceDiscoveryOperOper(ret edpt.DataDdosDstZoneDetectionServiceDiscoveryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"discovered_service_list": setSliceDdosDstZoneDetectionServiceDiscoveryOperOperDiscoveredServiceList(ret.DtDdosDstZoneDetectionServiceDiscoveryOper.Oper.DiscoveredServiceList),
		},
	}
}

func setSliceDdosDstZoneDetectionServiceDiscoveryOperOperDiscoveredServiceList(d []edpt.DdosDstZoneDetectionServiceDiscoveryOperOperDiscoveredServiceList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port"] = item.Port
		in["protocol"] = item.Protocol
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZoneDetectionServiceDiscoveryOperOper(d []interface{}) edpt.DdosDstZoneDetectionServiceDiscoveryOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionServiceDiscoveryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DiscoveredServiceList = getSliceDdosDstZoneDetectionServiceDiscoveryOperOperDiscoveredServiceList(in["discovered_service_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneDetectionServiceDiscoveryOperOperDiscoveredServiceList(d []interface{}) []edpt.DdosDstZoneDetectionServiceDiscoveryOperOperDiscoveredServiceList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionServiceDiscoveryOperOperDiscoveredServiceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionServiceDiscoveryOperOperDiscoveredServiceList
		oi.Port = in["port"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Rate = in["rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneDetectionServiceDiscoveryOper(d *schema.ResourceData) edpt.DdosDstZoneDetectionServiceDiscoveryOper {
	var ret edpt.DdosDstZoneDetectionServiceDiscoveryOper

	ret.Oper = getObjectDdosDstZoneDetectionServiceDiscoveryOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
