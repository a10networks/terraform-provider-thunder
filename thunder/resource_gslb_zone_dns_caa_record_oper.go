package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneDnsCaaRecordOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_dns_caa_record_oper`: Operational Status for the object dns-caa-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneDnsCaaRecordOperRead,

		Schema: map[string]*schema.Schema{
			"critical_flag": {
				Type: schema.TypeInt, Required: true, Description: "Issuer Critical Flag",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"last_server": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"property_tag": {
				Type: schema.TypeString, Required: true, Description: "Specify other property tags, only allowed lowercase alphanumeric",
			},
			"rdata": {
				Type: schema.TypeString, Required: true, Description: "Specify the Issuer Domain Name or a URL",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceGslbZoneDnsCaaRecordOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsCaaRecordOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsCaaRecordOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneDnsCaaRecordOperOper := setObjectGslbZoneDnsCaaRecordOperOper(res)
		d.Set("oper", GslbZoneDnsCaaRecordOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneDnsCaaRecordOperOper(ret edpt.DataGslbZoneDnsCaaRecordOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"last_server": ret.DtGslbZoneDnsCaaRecordOper.Oper.LastServer,
		},
	}
}

func getObjectGslbZoneDnsCaaRecordOperOper(d []interface{}) edpt.GslbZoneDnsCaaRecordOperOper {

	count1 := len(d)
	var ret edpt.GslbZoneDnsCaaRecordOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LastServer = in["last_server"].(string)
	}
	return ret
}

func dataToEndpointGslbZoneDnsCaaRecordOper(d *schema.ResourceData) edpt.GslbZoneDnsCaaRecordOper {
	var ret edpt.GslbZoneDnsCaaRecordOper

	ret.CriticalFlag = d.Get("critical_flag").(int)

	ret.Oper = getObjectGslbZoneDnsCaaRecordOperOper(d.Get("oper").([]interface{}))

	ret.PropertyTag = d.Get("property_tag").(string)

	ret.Rdata = d.Get("rdata").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
