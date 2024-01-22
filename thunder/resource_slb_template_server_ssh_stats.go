package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateServerSshStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_server_ssh_stats`: Statistics for the object server-ssh\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplateServerSshStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Server SSH Template Name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"successful_handshakes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"failed_handshakes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"forwarding_errors": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbTemplateServerSshStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerSshStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServerSshStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplateServerSshStatsStats := setObjectSlbTemplateServerSshStatsStats(res)
		d.Set("stats", SlbTemplateServerSshStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplateServerSshStatsStats(ret edpt.DataSlbTemplateServerSshStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"successful_handshakes": ret.DtSlbTemplateServerSshStats.Stats.Successful_handshakes,
			"failed_handshakes":     ret.DtSlbTemplateServerSshStats.Stats.Failed_handshakes,
			"forwarding_errors":     ret.DtSlbTemplateServerSshStats.Stats.Forwarding_errors,
		},
	}
}

func getObjectSlbTemplateServerSshStatsStats(d []interface{}) edpt.SlbTemplateServerSshStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplateServerSshStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Successful_handshakes = in["successful_handshakes"].(int)
		ret.Failed_handshakes = in["failed_handshakes"].(int)
		ret.Forwarding_errors = in["forwarding_errors"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplateServerSshStats(d *schema.ResourceData) edpt.SlbTemplateServerSshStats {
	var ret edpt.SlbTemplateServerSshStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectSlbTemplateServerSshStatsStats(d.Get("stats").([]interface{}))
	return ret
}
