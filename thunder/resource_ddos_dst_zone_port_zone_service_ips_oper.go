package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceIpsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_ips_oper`: Operational Status for the object ips\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceIpsOperRead,

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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}

func resourceDdosDstZonePortZoneServiceIpsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceIpsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceIpsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceIpsOperOper := setObjectDdosDstZonePortZoneServiceIpsOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceIpsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceIpsOperOper(ret edpt.DataDdosDstZonePortZoneServiceIpsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"signature_list": setSliceDdosDstZonePortZoneServiceIpsOperOperSignatureList(ret.DtDdosDstZonePortZoneServiceIpsOper.Oper.SignatureList),
		},
	}
}

func setSliceDdosDstZonePortZoneServiceIpsOperOperSignatureList(d []edpt.DdosDstZonePortZoneServiceIpsOperOperSignatureList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sid"] = item.Sid
		in["match_count"] = item.MatchCount
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZonePortZoneServiceIpsOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceIpsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceIpsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SignatureList = getSliceDdosDstZonePortZoneServiceIpsOperOperSignatureList(in["signature_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceIpsOperOperSignatureList(d []interface{}) []edpt.DdosDstZonePortZoneServiceIpsOperOperSignatureList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceIpsOperOperSignatureList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceIpsOperOperSignatureList
		oi.Sid = in["sid"].(int)
		oi.MatchCount = in["match_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceIpsOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceIpsOper {
	var ret edpt.DdosDstZonePortZoneServiceIpsOper

	ret.Oper = getObjectDdosDstZonePortZoneServiceIpsOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortNum = d.Get("port_num").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
