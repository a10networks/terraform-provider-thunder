package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsPtrRecordStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_service_dns_ptr_record_stats`: Statistics for the object dns-ptr-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneServiceDnsPtrRecordStatsRead,

		Schema: map[string]*schema.Schema{
			"ptr_name": {
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

func resourceGslbZoneServiceDnsPtrRecordStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsPtrRecordStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsPtrRecordStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneServiceDnsPtrRecordStatsStats := setObjectGslbZoneServiceDnsPtrRecordStatsStats(res)
		d.Set("stats", GslbZoneServiceDnsPtrRecordStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneServiceDnsPtrRecordStatsStats(ret edpt.DataGslbZoneServiceDnsPtrRecordStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtGslbZoneServiceDnsPtrRecordStats.Stats.Hits,
		},
	}
}

func getObjectGslbZoneServiceDnsPtrRecordStatsStats(d []interface{}) edpt.GslbZoneServiceDnsPtrRecordStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceDnsPtrRecordStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsPtrRecordStats(d *schema.ResourceData) edpt.GslbZoneServiceDnsPtrRecordStats {
	var ret edpt.GslbZoneServiceDnsPtrRecordStats

	ret.PtrName = d.Get("ptr_name").(string)

	ret.Stats = getObjectGslbZoneServiceDnsPtrRecordStatsStats(d.Get("stats").([]interface{}))

	ret.ServiceName = d.Get("service_name").(string)

	ret.ServicePort = d.Get("service_port").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
