package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneDnsMxRecordOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_dns_mx_record_oper`: Operational Status for the object dns-mx-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneDnsMxRecordOperRead,

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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceGslbZoneDnsMxRecordOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsMxRecordOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsMxRecordOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneDnsMxRecordOperOper := setObjectGslbZoneDnsMxRecordOperOper(res)
		d.Set("oper", GslbZoneDnsMxRecordOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneDnsMxRecordOperOper(ret edpt.DataGslbZoneDnsMxRecordOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"last_server": ret.DtGslbZoneDnsMxRecordOper.Oper.LastServer,
			"hits":        ret.DtGslbZoneDnsMxRecordOper.Oper.Hits,
			"priority":    ret.DtGslbZoneDnsMxRecordOper.Oper.Priority,
		},
	}
}

func getObjectGslbZoneDnsMxRecordOperOper(d []interface{}) edpt.GslbZoneDnsMxRecordOperOper {

	count1 := len(d)
	var ret edpt.GslbZoneDnsMxRecordOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LastServer = in["last_server"].(string)
		ret.Hits = in["hits"].(int)
		ret.Priority = in["priority"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneDnsMxRecordOper(d *schema.ResourceData) edpt.GslbZoneDnsMxRecordOper {
	var ret edpt.GslbZoneDnsMxRecordOper

	ret.MxName = d.Get("mx_name").(string)

	ret.Oper = getObjectGslbZoneDnsMxRecordOperOper(d.Get("oper").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
