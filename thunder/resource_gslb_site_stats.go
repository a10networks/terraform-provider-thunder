package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_site_stats`: Statistics for the object site\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbSiteStatsRead,

		Schema: map[string]*schema.Schema{
			"ip_server_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
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
					},
				},
			},
			"site_name": {
				Type: schema.TypeString, Required: true, Description: "Specify GSLB site name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of times the site was selected",
						},
					},
				},
			},
		},
	}
}

func resourceGslbSiteStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbSiteStatsIpServerList := setSliceGslbSiteStatsIpServerList(res)
		d.Set("ip_server_list", GslbSiteStatsIpServerList)
		GslbSiteStatsStats := setObjectGslbSiteStatsStats(res)
		d.Set("stats", GslbSiteStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceGslbSiteStatsIpServerList(d edpt.DataGslbSiteStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbSiteStats.IpServerList {
		in := make(map[string]interface{})
		in["ip_server_name"] = item.IpServerName
		in["stats"] = setObjectGslbSiteStatsIpServerListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbSiteStatsIpServerListStats(d edpt.GslbSiteStatsIpServerListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits

	in["recent"] = d.Recent
	result = append(result, in)
	return result
}

func setObjectGslbSiteStatsStats(ret edpt.DataGslbSiteStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtGslbSiteStats.Stats.Hits,
		},
	}
}

func getSliceGslbSiteStatsIpServerList(d []interface{}) []edpt.GslbSiteStatsIpServerList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteStatsIpServerList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteStatsIpServerList
		oi.IpServerName = in["ip_server_name"].(string)
		oi.Stats = getObjectGslbSiteStatsIpServerListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbSiteStatsIpServerListStats(d []interface{}) edpt.GslbSiteStatsIpServerListStats {

	count1 := len(d)
	var ret edpt.GslbSiteStatsIpServerListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
		ret.Recent = in["recent"].(int)
	}
	return ret
}

func getObjectGslbSiteStatsStats(d []interface{}) edpt.GslbSiteStatsStats {

	count1 := len(d)
	var ret edpt.GslbSiteStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbSiteStats(d *schema.ResourceData) edpt.GslbSiteStats {
	var ret edpt.GslbSiteStats

	ret.IpServerList = getSliceGslbSiteStatsIpServerList(d.Get("ip_server_list").([]interface{}))

	ret.SiteName = d.Get("site_name").(string)

	ret.Stats = getObjectGslbSiteStatsStats(d.Get("stats").([]interface{}))
	return ret
}
