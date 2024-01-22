package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVlanGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_vlan_global_stats`: Statistics for the object vlan-global\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkVlanGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"xparent_vlan_list_err": {
							Type: schema.TypeInt, Optional: true, Description: "Transparent Mode VLAN List Errors",
						},
					},
				},
			},
		},
	}
}

func resourceNetworkVlanGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlanGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkVlanGlobalStatsStats := setObjectNetworkVlanGlobalStatsStats(res)
		d.Set("stats", NetworkVlanGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkVlanGlobalStatsStats(ret edpt.DataNetworkVlanGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"xparent_vlan_list_err": ret.DtNetworkVlanGlobalStats.Stats.Xparent_vlan_list_err,
		},
	}
}

func getObjectNetworkVlanGlobalStatsStats(d []interface{}) edpt.NetworkVlanGlobalStatsStats {

	count1 := len(d)
	var ret edpt.NetworkVlanGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Xparent_vlan_list_err = in["xparent_vlan_list_err"].(int)
	}
	return ret
}

func dataToEndpointNetworkVlanGlobalStats(d *schema.ResourceData) edpt.NetworkVlanGlobalStats {
	var ret edpt.NetworkVlanGlobalStats

	ret.Stats = getObjectNetworkVlanGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
