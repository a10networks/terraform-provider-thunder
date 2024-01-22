package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteSlbDevVipServerVipServerNameStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_site_slb_dev_vip_server_vip_server_name_stats`: Statistics for the object vip-server-name\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbSiteSlbDevVipServerVipServerNameStatsRead,

		Schema: map[string]*schema.Schema{
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
			"vip_name": {
				Type: schema.TypeString, Required: true, Description: "Specify a VIP name for the SLB device",
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

func resourceGslbSiteSlbDevVipServerVipServerNameStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerNameStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerNameStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbSiteSlbDevVipServerVipServerNameStatsStats := setObjectGslbSiteSlbDevVipServerVipServerNameStatsStats(res)
		d.Set("stats", GslbSiteSlbDevVipServerVipServerNameStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbSiteSlbDevVipServerVipServerNameStatsStats(ret edpt.DataGslbSiteSlbDevVipServerVipServerNameStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dev_vip_hits":   ret.DtGslbSiteSlbDevVipServerVipServerNameStats.Stats.Dev_vip_hits,
			"dev_vip_recent": ret.DtGslbSiteSlbDevVipServerVipServerNameStats.Stats.Dev_vip_recent,
		},
	}
}

func getObjectGslbSiteSlbDevVipServerVipServerNameStatsStats(d []interface{}) edpt.GslbSiteSlbDevVipServerVipServerNameStatsStats {

	count1 := len(d)
	var ret edpt.GslbSiteSlbDevVipServerVipServerNameStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dev_vip_hits = in["dev_vip_hits"].(int)
		ret.Dev_vip_recent = in["dev_vip_recent"].(int)
	}
	return ret
}

func dataToEndpointGslbSiteSlbDevVipServerVipServerNameStats(d *schema.ResourceData) edpt.GslbSiteSlbDevVipServerVipServerNameStats {
	var ret edpt.GslbSiteSlbDevVipServerVipServerNameStats

	ret.Stats = getObjectGslbSiteSlbDevVipServerVipServerNameStatsStats(d.Get("stats").([]interface{}))

	ret.VipName = d.Get("vip_name").(string)

	ret.SiteName = d.Get("site_name").(string)

	ret.DeviceName = d.Get("device_name").(string)
	return ret
}
