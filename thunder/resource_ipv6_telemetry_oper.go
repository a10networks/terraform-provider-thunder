package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6TelemetryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ipv6_telemetry_oper`: Operational Status for the object telemetry\n\n__PLACEHOLDER__",
		ReadContext: resourceIpv6TelemetryOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"db": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"net_mask_len": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"class_list_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"class_list_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"zone_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"origin_as": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"peer_as": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_nexthop": {
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

func resourceIpv6TelemetryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6TelemetryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6TelemetryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Ipv6TelemetryOperOper := setObjectIpv6TelemetryOperOper(res)
		d.Set("oper", Ipv6TelemetryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpv6TelemetryOperOper(ret edpt.DataIpv6TelemetryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total": ret.DtIpv6TelemetryOper.Oper.Total,
			"db":    setSliceIpv6TelemetryOperOperDb(ret.DtIpv6TelemetryOper.Oper.Db),
		},
	}
}

func setSliceIpv6TelemetryOperOperDb(d []edpt.Ipv6TelemetryOperOperDb) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipv6_addr"] = item.Ipv6_addr
		in["net_mask_len"] = item.NetMaskLen
		in["class_list_name"] = item.ClassListName
		in["class_list_id"] = item.ClassListId
		in["zone_name"] = item.ZoneName
		in["origin_as"] = item.Origin_as
		in["peer_as"] = item.Peer_as
		in["ipv6_nexthop"] = item.Ipv6Nexthop
		result = append(result, in)
	}
	return result
}

func getObjectIpv6TelemetryOperOper(d []interface{}) edpt.Ipv6TelemetryOperOper {

	count1 := len(d)
	var ret edpt.Ipv6TelemetryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total = in["total"].(int)
		ret.Db = getSliceIpv6TelemetryOperOperDb(in["db"].([]interface{}))
	}
	return ret
}

func getSliceIpv6TelemetryOperOperDb(d []interface{}) []edpt.Ipv6TelemetryOperOperDb {

	count1 := len(d)
	ret := make([]edpt.Ipv6TelemetryOperOperDb, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Ipv6TelemetryOperOperDb
		oi.Ipv6_addr = in["ipv6_addr"].(string)
		oi.NetMaskLen = in["net_mask_len"].(int)
		oi.ClassListName = in["class_list_name"].(string)
		oi.ClassListId = in["class_list_id"].(int)
		oi.ZoneName = in["zone_name"].(string)
		oi.Origin_as = in["origin_as"].(int)
		oi.Peer_as = in["peer_as"].(int)
		oi.Ipv6Nexthop = in["ipv6_nexthop"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpv6TelemetryOper(d *schema.ResourceData) edpt.Ipv6TelemetryOper {
	var ret edpt.Ipv6TelemetryOper

	ret.Oper = getObjectIpv6TelemetryOperOper(d.Get("oper").([]interface{}))
	return ret
}
