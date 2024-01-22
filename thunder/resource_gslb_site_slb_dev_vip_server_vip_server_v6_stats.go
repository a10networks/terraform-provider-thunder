package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteSlbDevVipServerVipServerV6Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_site_slb_dev_vip_server_vip_server_v6_stats`: Statistics for the object vip-server-v6\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbSiteSlbDevVipServerVipServerV6StatsRead,

		Schema: map[string]*schema.Schema{
			"ipv6": {
				Type: schema.TypeString, Required: true, Description: "Specify IP address (IPv6 address)",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dev_vip_hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of times the service-ip was selected",
						},
						"dev_vip_recent": {
							Type: schema.TypeInt, Optional: true, Description: "Recent hits",
						},
					},
				},
			},
			"site_name": {
				Type: schema.TypeString, Required: true, Description: "SiteName",
			},
			"device_name": {
				Type: schema.TypeString, Required: true, Description: "DeviceName",
			},
		},
	}
}

func resourceGslbSiteSlbDevVipServerVipServerV6StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerV6StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerV6Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbSiteSlbDevVipServerVipServerV6StatsStats := setObjectGslbSiteSlbDevVipServerVipServerV6StatsStats(res)
		d.Set("stats", GslbSiteSlbDevVipServerVipServerV6StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbSiteSlbDevVipServerVipServerV6StatsStats(ret edpt.DataGslbSiteSlbDevVipServerVipServerV6Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dev_vip_hits":   ret.DtGslbSiteSlbDevVipServerVipServerV6Stats.Stats.Dev_vip_hits,
			"dev_vip_recent": ret.DtGslbSiteSlbDevVipServerVipServerV6Stats.Stats.Dev_vip_recent,
		},
	}
}

func getObjectGslbSiteSlbDevVipServerVipServerV6StatsStats(d []interface{}) edpt.GslbSiteSlbDevVipServerVipServerV6StatsStats {

	count1 := len(d)
	var ret edpt.GslbSiteSlbDevVipServerVipServerV6StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dev_vip_hits = in["dev_vip_hits"].(int)
		ret.Dev_vip_recent = in["dev_vip_recent"].(int)
	}
	return ret
}

func dataToEndpointGslbSiteSlbDevVipServerVipServerV6Stats(d *schema.ResourceData) edpt.GslbSiteSlbDevVipServerVipServerV6Stats {
	var ret edpt.GslbSiteSlbDevVipServerVipServerV6Stats

	ret.Ipv6 = d.Get("ipv6").(string)

	ret.Stats = getObjectGslbSiteSlbDevVipServerVipServerV6StatsStats(d.Get("stats").([]interface{}))

	ret.SiteName = d.Get("site_name").(string)

	ret.DeviceName = d.Get("device_name").(string)
	return ret
}
