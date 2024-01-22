package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneDnsNsRecordOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_dns_ns_record_oper`: Operational Status for the object dns-ns-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneDnsNsRecordOperRead,

		Schema: map[string]*schema.Schema{
			"ns_name": {
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
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceGslbZoneDnsNsRecordOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsNsRecordOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsNsRecordOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneDnsNsRecordOperOper := setObjectGslbZoneDnsNsRecordOperOper(res)
		d.Set("oper", GslbZoneDnsNsRecordOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneDnsNsRecordOperOper(ret edpt.DataGslbZoneDnsNsRecordOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"last_server": ret.DtGslbZoneDnsNsRecordOper.Oper.LastServer,
			"hits":        ret.DtGslbZoneDnsNsRecordOper.Oper.Hits,
		},
	}
}

func getObjectGslbZoneDnsNsRecordOperOper(d []interface{}) edpt.GslbZoneDnsNsRecordOperOper {

	count1 := len(d)
	var ret edpt.GslbZoneDnsNsRecordOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LastServer = in["last_server"].(string)
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneDnsNsRecordOper(d *schema.ResourceData) edpt.GslbZoneDnsNsRecordOper {
	var ret edpt.GslbZoneDnsNsRecordOper

	ret.NsName = d.Get("ns_name").(string)

	ret.Oper = getObjectGslbZoneDnsNsRecordOperOper(d.Get("oper").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
