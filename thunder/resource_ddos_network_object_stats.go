package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_network_object_stats`: Statistics for the object network-object\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosNetworkObjectStatsRead,

		Schema: map[string]*schema.Schema{
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_learned": {
							Type: schema.TypeInt, Optional: true, Description: "Subnet Entry Learned",
						},
						"subnet_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Subnet Entry Aged",
						},
						"subnet_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Subnet Entry Create Failures",
						},
						"ip_learned": {
							Type: schema.TypeInt, Optional: true, Description: "IP Entry Learned",
						},
						"ip_aged": {
							Type: schema.TypeInt, Optional: true, Description: "IP Entry Aged",
						},
						"ip_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "IP Entry Create Failures",
						},
						"service_learned": {
							Type: schema.TypeInt, Optional: true, Description: "Service Entry Learned",
						},
						"service_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Service Entry Aged",
						},
						"service_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Service Entry Create Failures",
						},
					},
				},
			},
		},
	}
}

func resourceDdosNetworkObjectStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosNetworkObjectStatsStats := setObjectDdosNetworkObjectStatsStats(res)
		d.Set("stats", DdosNetworkObjectStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosNetworkObjectStatsStats(ret edpt.DataDdosNetworkObjectStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"subnet_learned":      ret.DtDdosNetworkObjectStats.Stats.Subnet_learned,
			"subnet_aged":         ret.DtDdosNetworkObjectStats.Stats.Subnet_aged,
			"subnet_create_fail":  ret.DtDdosNetworkObjectStats.Stats.Subnet_create_fail,
			"ip_learned":          ret.DtDdosNetworkObjectStats.Stats.Ip_learned,
			"ip_aged":             ret.DtDdosNetworkObjectStats.Stats.Ip_aged,
			"ip_create_fail":      ret.DtDdosNetworkObjectStats.Stats.Ip_create_fail,
			"service_learned":     ret.DtDdosNetworkObjectStats.Stats.Service_learned,
			"service_aged":        ret.DtDdosNetworkObjectStats.Stats.Service_aged,
			"service_create_fail": ret.DtDdosNetworkObjectStats.Stats.Service_create_fail,
		},
	}
}

func getObjectDdosNetworkObjectStatsStats(d []interface{}) edpt.DdosNetworkObjectStatsStats {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Subnet_learned = in["subnet_learned"].(int)
		ret.Subnet_aged = in["subnet_aged"].(int)
		ret.Subnet_create_fail = in["subnet_create_fail"].(int)
		ret.Ip_learned = in["ip_learned"].(int)
		ret.Ip_aged = in["ip_aged"].(int)
		ret.Ip_create_fail = in["ip_create_fail"].(int)
		ret.Service_learned = in["service_learned"].(int)
		ret.Service_aged = in["service_aged"].(int)
		ret.Service_create_fail = in["service_create_fail"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectStats(d *schema.ResourceData) edpt.DdosNetworkObjectStats {
	var ret edpt.DdosNetworkObjectStats

	ret.ObjectName = d.Get("object_name").(string)

	ret.Stats = getObjectDdosNetworkObjectStatsStats(d.Get("stats").([]interface{}))
	return ret
}
