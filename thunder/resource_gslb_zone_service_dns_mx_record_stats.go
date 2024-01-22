package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsMxRecordStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_service_dns_mx_record_stats`: Statistics for the object dns-mx-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneServiceDnsMxRecordStatsRead,

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

func resourceGslbZoneServiceDnsMxRecordStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsMxRecordStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsMxRecordStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneServiceDnsMxRecordStatsStats := setObjectGslbZoneServiceDnsMxRecordStatsStats(res)
		d.Set("stats", GslbZoneServiceDnsMxRecordStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneServiceDnsMxRecordStatsStats(ret edpt.DataGslbZoneServiceDnsMxRecordStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtGslbZoneServiceDnsMxRecordStats.Stats.Hits,
		},
	}
}

func getObjectGslbZoneServiceDnsMxRecordStatsStats(d []interface{}) edpt.GslbZoneServiceDnsMxRecordStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceDnsMxRecordStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsMxRecordStats(d *schema.ResourceData) edpt.GslbZoneServiceDnsMxRecordStats {
	var ret edpt.GslbZoneServiceDnsMxRecordStats

	ret.MxName = d.Get("mx_name").(string)

	ret.Stats = getObjectGslbZoneServiceDnsMxRecordStatsStats(d.Get("stats").([]interface{}))

	ret.ServiceName = d.Get("service_name").(string)

	ret.ServicePort = d.Get("service_port").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
