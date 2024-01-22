package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsSrvRecordStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_service_dns_srv_record_stats`: Statistics for the object dns-srv-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneServiceDnsSrvRecordStatsRead,

		Schema: map[string]*schema.Schema{
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Specify Port (Port Number)",
			},
			"srv_name": {
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
			"service_port": {
				Type: schema.TypeString, Required: true, Description: "ServicePort",
			},
			"service_name": {
				Type: schema.TypeString, Required: true, Description: "ServiceName",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceGslbZoneServiceDnsSrvRecordStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsSrvRecordStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsSrvRecordStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneServiceDnsSrvRecordStatsStats := setObjectGslbZoneServiceDnsSrvRecordStatsStats(res)
		d.Set("stats", GslbZoneServiceDnsSrvRecordStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneServiceDnsSrvRecordStatsStats(ret edpt.DataGslbZoneServiceDnsSrvRecordStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtGslbZoneServiceDnsSrvRecordStats.Stats.Hits,
		},
	}
}

func getObjectGslbZoneServiceDnsSrvRecordStatsStats(d []interface{}) edpt.GslbZoneServiceDnsSrvRecordStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceDnsSrvRecordStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsSrvRecordStats(d *schema.ResourceData) edpt.GslbZoneServiceDnsSrvRecordStats {
	var ret edpt.GslbZoneServiceDnsSrvRecordStats

	ret.Port = d.Get("port").(int)

	ret.SrvName = d.Get("srv_name").(string)

	ret.Stats = getObjectGslbZoneServiceDnsSrvRecordStatsStats(d.Get("stats").([]interface{}))

	ret.ServicePort = d.Get("service_port").(string)

	ret.ServiceName = d.Get("service_name").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
