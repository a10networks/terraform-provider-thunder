package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsMxRecordOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_service_dns_mx_record_oper`: Operational Status for the object dns-mx-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneServiceDnsMxRecordOperRead,

		Schema: map[string]*schema.Schema{
			"mx_name": {
				Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"last_server": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"service_name": {
				Type: schema.TypeString, Required: true, Description: "ServiceName",
			},
			"service_port": {
				Type: schema.TypeString, Required: true, Description: "ServicePort",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceGslbZoneServiceDnsMxRecordOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsMxRecordOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsMxRecordOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneServiceDnsMxRecordOperOper := setObjectGslbZoneServiceDnsMxRecordOperOper(res)
		d.Set("oper", GslbZoneServiceDnsMxRecordOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneServiceDnsMxRecordOperOper(ret edpt.DataGslbZoneServiceDnsMxRecordOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"last_server": ret.DtGslbZoneServiceDnsMxRecordOper.Oper.LastServer,
			"hits":        ret.DtGslbZoneServiceDnsMxRecordOper.Oper.Hits,
			"priority":    ret.DtGslbZoneServiceDnsMxRecordOper.Oper.Priority,
		},
	}
}

func getObjectGslbZoneServiceDnsMxRecordOperOper(d []interface{}) edpt.GslbZoneServiceDnsMxRecordOperOper {

	count1 := len(d)
	var ret edpt.GslbZoneServiceDnsMxRecordOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LastServer = in["last_server"].(string)
		ret.Hits = in["hits"].(int)
		ret.Priority = in["priority"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsMxRecordOper(d *schema.ResourceData) edpt.GslbZoneServiceDnsMxRecordOper {
	var ret edpt.GslbZoneServiceDnsMxRecordOper

	ret.MxName = d.Get("mx_name").(string)

	ret.Oper = getObjectGslbZoneServiceDnsMxRecordOperOper(d.Get("oper").([]interface{}))

	ret.ServiceName = d.Get("service_name").(string)

	ret.ServicePort = d.Get("service_port").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
