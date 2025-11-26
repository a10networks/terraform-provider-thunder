package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkTrunkStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_trunk_stats`: Enables/Disables trunk interface stats generation\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkTrunkStatsRead,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable trunk interface stats generation",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}

func resourceNetworkTrunkStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTrunkStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTrunkStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkTrunkStats(d *schema.ResourceData) edpt.NetworkTrunkStats {
	var ret edpt.NetworkTrunkStats

	ret.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
