package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneDnsMxRecordStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_dns_mx_record_stats`: Statistics for the object dns-mx-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneDnsMxRecordStatsRead,

		Schema: map[string]*schema.Schema{
			"mx_name": {
				Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
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

func resourceGslbZoneDnsMxRecordStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsMxRecordStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsMxRecordStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneDnsMxRecordStatsStats := setObjectGslbZoneDnsMxRecordStatsStats(res)
		d.Set("stats", GslbZoneDnsMxRecordStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneDnsMxRecordStatsStats(ret edpt.DataGslbZoneDnsMxRecordStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtGslbZoneDnsMxRecordStats.Stats.Hits,
		},
	}
}

func getObjectGslbZoneDnsMxRecordStatsStats(d []interface{}) edpt.GslbZoneDnsMxRecordStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneDnsMxRecordStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneDnsMxRecordStats(d *schema.ResourceData) edpt.GslbZoneDnsMxRecordStats {
	var ret edpt.GslbZoneDnsMxRecordStats

	ret.MxName = d.Get("mx_name").(string)

	ret.Stats = getObjectGslbZoneDnsMxRecordStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
