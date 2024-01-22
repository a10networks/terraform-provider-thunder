package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneDnsCaaRecordStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_dns_caa_record_stats`: Statistics for the object dns-caa-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneDnsCaaRecordStatsRead,

		Schema: map[string]*schema.Schema{
			"critical_flag": {
				Type: schema.TypeInt, Required: true, Description: "Issuer Critical Flag",
			},
			"property_tag": {
				Type: schema.TypeString, Required: true, Description: "Specify other property tags, only allowed lowercase alphanumeric",
			},
			"rdata": {
				Type: schema.TypeString, Required: true, Description: "Specify the Issuer Domain Name or a URL",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of times the record has been used",
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

func resourceGslbZoneDnsCaaRecordStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsCaaRecordStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsCaaRecordStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneDnsCaaRecordStatsStats := setObjectGslbZoneDnsCaaRecordStatsStats(res)
		d.Set("stats", GslbZoneDnsCaaRecordStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneDnsCaaRecordStatsStats(ret edpt.DataGslbZoneDnsCaaRecordStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtGslbZoneDnsCaaRecordStats.Stats.Hits,
		},
	}
}

func getObjectGslbZoneDnsCaaRecordStatsStats(d []interface{}) edpt.GslbZoneDnsCaaRecordStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneDnsCaaRecordStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneDnsCaaRecordStats(d *schema.ResourceData) edpt.GslbZoneDnsCaaRecordStats {
	var ret edpt.GslbZoneDnsCaaRecordStats

	ret.CriticalFlag = d.Get("critical_flag").(int)

	ret.PropertyTag = d.Get("property_tag").(string)

	ret.Rdata = d.Get("rdata").(string)

	ret.Stats = getObjectGslbZoneDnsCaaRecordStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
