package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateClientSshStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_client_ssh_stats`: Statistics for the object client-ssh\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplateClientSshStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Client SSH Template Name",
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

func resourceSlbTemplateClientSshStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSshStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSshStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplateClientSshStatsStats := setObjectSlbTemplateClientSshStatsStats(res)
		d.Set("stats", SlbTemplateClientSshStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplateClientSshStatsStats(ret edpt.DataSlbTemplateClientSshStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"successful_handshakes": ret.DtSlbTemplateClientSshStats.Stats.Successful_handshakes,
			"failed_handshakes":     ret.DtSlbTemplateClientSshStats.Stats.Failed_handshakes,
			"forwarding_errors":     ret.DtSlbTemplateClientSshStats.Stats.Forwarding_errors,
		},
	}
}

func getObjectSlbTemplateClientSshStatsStats(d []interface{}) edpt.SlbTemplateClientSshStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplateClientSshStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Successful_handshakes = in["successful_handshakes"].(int)
		ret.Failed_handshakes = in["failed_handshakes"].(int)
		ret.Forwarding_errors = in["forwarding_errors"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplateClientSshStats(d *schema.ResourceData) edpt.SlbTemplateClientSshStats {
	var ret edpt.SlbTemplateClientSshStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectSlbTemplateClientSshStatsStats(d.Get("stats").([]interface{}))
	return ret
}
