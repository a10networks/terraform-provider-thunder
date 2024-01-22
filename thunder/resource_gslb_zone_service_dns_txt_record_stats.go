package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsTxtRecordStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_service_dns_txt_record_stats`: Statistics for the object dns-txt-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneServiceDnsTxtRecordStatsRead,

		Schema: map[string]*schema.Schema{
			"record_name": {
				Type: schema.TypeString, Required: true, Description: "Specify the Object Name for TXT Data",
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

func resourceGslbZoneServiceDnsTxtRecordStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsTxtRecordStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsTxtRecordStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneServiceDnsTxtRecordStatsStats := setObjectGslbZoneServiceDnsTxtRecordStatsStats(res)
		d.Set("stats", GslbZoneServiceDnsTxtRecordStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneServiceDnsTxtRecordStatsStats(ret edpt.DataGslbZoneServiceDnsTxtRecordStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtGslbZoneServiceDnsTxtRecordStats.Stats.Hits,
		},
	}
}

func getObjectGslbZoneServiceDnsTxtRecordStatsStats(d []interface{}) edpt.GslbZoneServiceDnsTxtRecordStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceDnsTxtRecordStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsTxtRecordStats(d *schema.ResourceData) edpt.GslbZoneServiceDnsTxtRecordStats {
	var ret edpt.GslbZoneServiceDnsTxtRecordStats

	ret.RecordName = d.Get("record_name").(string)

	ret.Stats = getObjectGslbZoneServiceDnsTxtRecordStatsStats(d.Get("stats").([]interface{}))

	ret.ServiceName = d.Get("service_name").(string)

	ret.ServicePort = d.Get("service_port").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
