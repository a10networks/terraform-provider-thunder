package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpTelemetryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ip_telemetry_oper`: Operational Status for the object telemetry\n\n__PLACEHOLDER__",
		ReadContext: resourceIpTelemetryOperRead,

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
									"ipv4_addr": {
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
									"ipv4_nexthop": {
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

func resourceIpTelemetryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpTelemetryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpTelemetryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpTelemetryOperOper := setObjectIpTelemetryOperOper(res)
		d.Set("oper", IpTelemetryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpTelemetryOperOper(ret edpt.DataIpTelemetryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total": ret.DtIpTelemetryOper.Oper.Total,
			"db":    setSliceIpTelemetryOperOperDb(ret.DtIpTelemetryOper.Oper.Db),
		},
	}
}

func setSliceIpTelemetryOperOperDb(d []edpt.IpTelemetryOperOperDb) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipv4_addr"] = item.Ipv4_addr
		in["net_mask_len"] = item.NetMaskLen
		in["class_list_name"] = item.ClassListName
		in["class_list_id"] = item.ClassListId
		in["zone_name"] = item.ZoneName
		in["origin_as"] = item.Origin_as
		in["peer_as"] = item.Peer_as
		in["ipv4_nexthop"] = item.Ipv4Nexthop
		result = append(result, in)
	}
	return result
}

func getObjectIpTelemetryOperOper(d []interface{}) edpt.IpTelemetryOperOper {

	count1 := len(d)
	var ret edpt.IpTelemetryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total = in["total"].(int)
		ret.Db = getSliceIpTelemetryOperOperDb(in["db"].([]interface{}))
	}
	return ret
}

func getSliceIpTelemetryOperOperDb(d []interface{}) []edpt.IpTelemetryOperOperDb {

	count1 := len(d)
	ret := make([]edpt.IpTelemetryOperOperDb, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpTelemetryOperOperDb
		oi.Ipv4_addr = in["ipv4_addr"].(string)
		oi.NetMaskLen = in["net_mask_len"].(int)
		oi.ClassListName = in["class_list_name"].(string)
		oi.ClassListId = in["class_list_id"].(int)
		oi.ZoneName = in["zone_name"].(string)
		oi.Origin_as = in["origin_as"].(int)
		oi.Peer_as = in["peer_as"].(int)
		oi.Ipv4Nexthop = in["ipv4_nexthop"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpTelemetryOper(d *schema.ResourceData) edpt.IpTelemetryOper {
	var ret edpt.IpTelemetryOper

	ret.Oper = getObjectIpTelemetryOperOper(d.Get("oper").([]interface{}))
	return ret
}
