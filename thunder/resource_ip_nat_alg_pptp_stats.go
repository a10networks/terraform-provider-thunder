package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatAlgPptpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ip_nat_alg_pptp_stats`: Statistics for the object pptp\n\n__PLACEHOLDER__",
		ReadContext: resourceIpNatAlgPptpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"current_smp_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"current_gre_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"smp_session_creation_failure": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"truncated_pns_message": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"truncated_pac_message": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mismatched_pns_call_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mismatched_pac_call_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"retransmitted_pns_message": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"retransmitted_pac_message": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"truncated_gre_packet": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"unknown_gre_version": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"no_matching_gre_session": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceIpNatAlgPptpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatAlgPptpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatAlgPptpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpNatAlgPptpStatsStats := setObjectIpNatAlgPptpStatsStats(res)
		d.Set("stats", IpNatAlgPptpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpNatAlgPptpStatsStats(ret edpt.DataIpNatAlgPptpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"current_smp_sessions":         ret.DtIpNatAlgPptpStats.Stats.CurrentSmpSessions,
			"current_gre_sessions":         ret.DtIpNatAlgPptpStats.Stats.CurrentGreSessions,
			"smp_session_creation_failure": ret.DtIpNatAlgPptpStats.Stats.SmpSessionCreationFailure,
			"truncated_pns_message":        ret.DtIpNatAlgPptpStats.Stats.TruncatedPnsMessage,
			"truncated_pac_message":        ret.DtIpNatAlgPptpStats.Stats.TruncatedPacMessage,
			"mismatched_pns_call_id":       ret.DtIpNatAlgPptpStats.Stats.MismatchedPnsCallId,
			"mismatched_pac_call_id":       ret.DtIpNatAlgPptpStats.Stats.MismatchedPacCallId,
			"retransmitted_pns_message":    ret.DtIpNatAlgPptpStats.Stats.RetransmittedPnsMessage,
			"retransmitted_pac_message":    ret.DtIpNatAlgPptpStats.Stats.RetransmittedPacMessage,
			"truncated_gre_packet":         ret.DtIpNatAlgPptpStats.Stats.TruncatedGrePacket,
			"unknown_gre_version":          ret.DtIpNatAlgPptpStats.Stats.UnknownGreVersion,
			"no_matching_gre_session":      ret.DtIpNatAlgPptpStats.Stats.NoMatchingGreSession,
		},
	}
}

func getObjectIpNatAlgPptpStatsStats(d []interface{}) edpt.IpNatAlgPptpStatsStats {

	count1 := len(d)
	var ret edpt.IpNatAlgPptpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CurrentSmpSessions = in["current_smp_sessions"].(int)
		ret.CurrentGreSessions = in["current_gre_sessions"].(int)
		ret.SmpSessionCreationFailure = in["smp_session_creation_failure"].(int)
		ret.TruncatedPnsMessage = in["truncated_pns_message"].(int)
		ret.TruncatedPacMessage = in["truncated_pac_message"].(int)
		ret.MismatchedPnsCallId = in["mismatched_pns_call_id"].(int)
		ret.MismatchedPacCallId = in["mismatched_pac_call_id"].(int)
		ret.RetransmittedPnsMessage = in["retransmitted_pns_message"].(int)
		ret.RetransmittedPacMessage = in["retransmitted_pac_message"].(int)
		ret.TruncatedGrePacket = in["truncated_gre_packet"].(int)
		ret.UnknownGreVersion = in["unknown_gre_version"].(int)
		ret.NoMatchingGreSession = in["no_matching_gre_session"].(int)
	}
	return ret
}

func dataToEndpointIpNatAlgPptpStats(d *schema.ResourceData) edpt.IpNatAlgPptpStats {
	var ret edpt.IpNatAlgPptpStats

	ret.Stats = getObjectIpNatAlgPptpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
