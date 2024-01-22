package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteSlbDevVipServerVipServerV4Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_site_slb_dev_vip_server_vip_server_v4_stats`: Statistics for the object vip-server-v4\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbSiteSlbDevVipServerVipServerV4StatsRead,

		Schema: map[string]*schema.Schema{
			"ipv4": {
				Type: schema.TypeString, Required: true, Description: "Specify IP address",
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

func resourceGslbSiteSlbDevVipServerVipServerV4StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerV4StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerV4Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbSiteSlbDevVipServerVipServerV4StatsStats := setObjectGslbSiteSlbDevVipServerVipServerV4StatsStats(res)
		d.Set("stats", GslbSiteSlbDevVipServerVipServerV4StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbSiteSlbDevVipServerVipServerV4StatsStats(ret edpt.DataGslbSiteSlbDevVipServerVipServerV4Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dev_vip_hits":   ret.DtGslbSiteSlbDevVipServerVipServerV4Stats.Stats.Dev_vip_hits,
			"dev_vip_recent": ret.DtGslbSiteSlbDevVipServerVipServerV4Stats.Stats.Dev_vip_recent,
		},
	}
}

func getObjectGslbSiteSlbDevVipServerVipServerV4StatsStats(d []interface{}) edpt.GslbSiteSlbDevVipServerVipServerV4StatsStats {

	count1 := len(d)
	var ret edpt.GslbSiteSlbDevVipServerVipServerV4StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dev_vip_hits = in["dev_vip_hits"].(int)
		ret.Dev_vip_recent = in["dev_vip_recent"].(int)
	}
	return ret
}

func dataToEndpointGslbSiteSlbDevVipServerVipServerV4Stats(d *schema.ResourceData) edpt.GslbSiteSlbDevVipServerVipServerV4Stats {
	var ret edpt.GslbSiteSlbDevVipServerVipServerV4Stats

	ret.Ipv4 = d.Get("ipv4").(string)

	ret.Stats = getObjectGslbSiteSlbDevVipServerVipServerV4StatsStats(d.Get("stats").([]interface{}))

	ret.SiteName = d.Get("site_name").(string)

	ret.DeviceName = d.Get("device_name").(string)
	return ret
}
