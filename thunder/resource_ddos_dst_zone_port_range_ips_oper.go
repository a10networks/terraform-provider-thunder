package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangeIpsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_range_ips_oper`: Operational Status for the object ips\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortRangeIpsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"signature_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"match_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
		},
	}
}

func resourceDdosDstZonePortRangeIpsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeIpsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeIpsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortRangeIpsOperOper := setObjectDdosDstZonePortRangeIpsOperOper(res)
		d.Set("oper", DdosDstZonePortRangeIpsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortRangeIpsOperOper(ret edpt.DataDdosDstZonePortRangeIpsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"signature_list": setSliceDdosDstZonePortRangeIpsOperOperSignatureList(ret.DtDdosDstZonePortRangeIpsOper.Oper.SignatureList),
		},
	}
}

func setSliceDdosDstZonePortRangeIpsOperOperSignatureList(d []edpt.DdosDstZonePortRangeIpsOperOperSignatureList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sid"] = item.Sid
		in["match_count"] = item.MatchCount
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZonePortRangeIpsOperOper(d []interface{}) edpt.DdosDstZonePortRangeIpsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeIpsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SignatureList = getSliceDdosDstZonePortRangeIpsOperOperSignatureList(in["signature_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortRangeIpsOperOperSignatureList(d []interface{}) []edpt.DdosDstZonePortRangeIpsOperOperSignatureList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeIpsOperOperSignatureList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeIpsOperOperSignatureList
		oi.Sid = in["sid"].(int)
		oi.MatchCount = in["match_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangeIpsOper(d *schema.ResourceData) edpt.DdosDstZonePortRangeIpsOper {
	var ret edpt.DdosDstZonePortRangeIpsOper

	ret.Oper = getObjectDdosDstZonePortRangeIpsOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortRangeEnd = d.Get("port_range_end").(string)

	ret.PortRangeStart = d.Get("port_range_start").(string)
	return ret
}
