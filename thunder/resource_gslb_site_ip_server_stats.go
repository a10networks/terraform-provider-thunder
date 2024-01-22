package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteIpServerStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_site_ip_server_stats`: Statistics for the object ip-server\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbSiteIpServerStatsRead,

		Schema: map[string]*schema.Schema{
			"ip_server_name": {
				Type: schema.TypeString, Required: true, Description: "Specify the real server name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of times the IP was selected",
						},
						"recent": {
							Type: schema.TypeInt, Optional: true, Description: "Recent hits",
						},
					},
				},
			},
			"site_name": {
				Type: schema.TypeString, Required: true, Description: "SiteName",
			},
		},
	}
}

func resourceGslbSiteIpServerStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteIpServerStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteIpServerStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbSiteIpServerStatsStats := setObjectGslbSiteIpServerStatsStats(res)
		d.Set("stats", GslbSiteIpServerStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbSiteIpServerStatsStats(ret edpt.DataGslbSiteIpServerStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits":   ret.DtGslbSiteIpServerStats.Stats.Hits,
			"recent": ret.DtGslbSiteIpServerStats.Stats.Recent,
		},
	}
}

func getObjectGslbSiteIpServerStatsStats(d []interface{}) edpt.GslbSiteIpServerStatsStats {

	count1 := len(d)
	var ret edpt.GslbSiteIpServerStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
		ret.Recent = in["recent"].(int)
	}
	return ret
}

func dataToEndpointGslbSiteIpServerStats(d *schema.ResourceData) edpt.GslbSiteIpServerStats {
	var ret edpt.GslbSiteIpServerStats

	ret.IpServerName = d.Get("ip_server_name").(string)

	ret.Stats = getObjectGslbSiteIpServerStatsStats(d.Get("stats").([]interface{}))

	ret.SiteName = d.Get("site_name").(string)
	return ret
}
