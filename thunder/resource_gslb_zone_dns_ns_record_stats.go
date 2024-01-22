package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneDnsNsRecordStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_dns_ns_record_stats`: Statistics for the object dns-ns-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneDnsNsRecordStatsRead,

		Schema: map[string]*schema.Schema{
			"ns_name": {
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

func resourceGslbZoneDnsNsRecordStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneDnsNsRecordStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneDnsNsRecordStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneDnsNsRecordStatsStats := setObjectGslbZoneDnsNsRecordStatsStats(res)
		d.Set("stats", GslbZoneDnsNsRecordStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneDnsNsRecordStatsStats(ret edpt.DataGslbZoneDnsNsRecordStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtGslbZoneDnsNsRecordStats.Stats.Hits,
		},
	}
}

func getObjectGslbZoneDnsNsRecordStatsStats(d []interface{}) edpt.GslbZoneDnsNsRecordStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneDnsNsRecordStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneDnsNsRecordStats(d *schema.ResourceData) edpt.GslbZoneDnsNsRecordStats {
	var ret edpt.GslbZoneDnsNsRecordStats

	ret.NsName = d.Get("ns_name").(string)

	ret.Stats = getObjectGslbZoneDnsNsRecordStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
