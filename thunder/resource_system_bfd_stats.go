package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemBfdStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_bfd_stats`: Statistics for the object bfd\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemBfdStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_checksum_error": {
							Type: schema.TypeInt, Optional: true, Description: "IP packet checksum errors",
						},
						"udp_checksum_error": {
							Type: schema.TypeInt, Optional: true, Description: "UDP packet checksum errors",
						},
						"session_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Session not found",
						},
						"multihop_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Multihop session or packet mismatch",
						},
						"version_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "BFD version mismatch",
						},
						"length_too_small": {
							Type: schema.TypeInt, Optional: true, Description: "Packets too small",
						},
						"data_is_short": {
							Type: schema.TypeInt, Optional: true, Description: "Packet data length too short",
						},
						"invalid_detect_mult": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid detect multiplier",
						},
						"invalid_multipoint": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid multipoint setting",
						},
						"invalid_my_disc": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid my descriptor",
						},
						"invalid_ttl": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid TTL",
						},
						"auth_length_invalid": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid authentication length",
						},
						"auth_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Authentication mismatch",
						},
						"auth_type_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Authentication type mismatch",
						},
						"auth_key_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Authentication key-id mismatch",
						},
						"auth_key_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Authentication key mismatch",
						},
						"auth_seqnum_invalid": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid authentication sequence number",
						},
						"auth_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Authentication failures",
						},
						"local_state_admin_down": {
							Type: schema.TypeInt, Optional: true, Description: "Local admin down session state",
						},
						"dest_unreachable": {
							Type: schema.TypeInt, Optional: true, Description: "Destination unreachable",
						},
						"no_ipv6_enable": {
							Type: schema.TypeInt, Optional: true, Description: "No IPv6 enable",
						},
						"other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Other errors",
						},
					},
				},
			},
		},
	}
}

func resourceSystemBfdStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBfdStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBfdStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemBfdStatsStats := setObjectSystemBfdStatsStats(res)
		d.Set("stats", SystemBfdStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemBfdStatsStats(ret edpt.DataSystemBfdStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ip_checksum_error":      ret.DtSystemBfdStats.Stats.Ip_checksum_error,
			"udp_checksum_error":     ret.DtSystemBfdStats.Stats.Udp_checksum_error,
			"session_not_found":      ret.DtSystemBfdStats.Stats.Session_not_found,
			"multihop_mismatch":      ret.DtSystemBfdStats.Stats.Multihop_mismatch,
			"version_mismatch":       ret.DtSystemBfdStats.Stats.Version_mismatch,
			"length_too_small":       ret.DtSystemBfdStats.Stats.Length_too_small,
			"data_is_short":          ret.DtSystemBfdStats.Stats.Data_is_short,
			"invalid_detect_mult":    ret.DtSystemBfdStats.Stats.Invalid_detect_mult,
			"invalid_multipoint":     ret.DtSystemBfdStats.Stats.Invalid_multipoint,
			"invalid_my_disc":        ret.DtSystemBfdStats.Stats.Invalid_my_disc,
			"invalid_ttl":            ret.DtSystemBfdStats.Stats.Invalid_ttl,
			"auth_length_invalid":    ret.DtSystemBfdStats.Stats.Auth_length_invalid,
			"auth_mismatch":          ret.DtSystemBfdStats.Stats.Auth_mismatch,
			"auth_type_mismatch":     ret.DtSystemBfdStats.Stats.Auth_type_mismatch,
			"auth_key_id_mismatch":   ret.DtSystemBfdStats.Stats.Auth_key_id_mismatch,
			"auth_key_mismatch":      ret.DtSystemBfdStats.Stats.Auth_key_mismatch,
			"auth_seqnum_invalid":    ret.DtSystemBfdStats.Stats.Auth_seqnum_invalid,
			"auth_failed":            ret.DtSystemBfdStats.Stats.Auth_failed,
			"local_state_admin_down": ret.DtSystemBfdStats.Stats.Local_state_admin_down,
			"dest_unreachable":       ret.DtSystemBfdStats.Stats.Dest_unreachable,
			"no_ipv6_enable":         ret.DtSystemBfdStats.Stats.No_ipv6_enable,
			"other_error":            ret.DtSystemBfdStats.Stats.Other_error,
		},
	}
}

func getObjectSystemBfdStatsStats(d []interface{}) edpt.SystemBfdStatsStats {

	count1 := len(d)
	var ret edpt.SystemBfdStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ip_checksum_error = in["ip_checksum_error"].(int)
		ret.Udp_checksum_error = in["udp_checksum_error"].(int)
		ret.Session_not_found = in["session_not_found"].(int)
		ret.Multihop_mismatch = in["multihop_mismatch"].(int)
		ret.Version_mismatch = in["version_mismatch"].(int)
		ret.Length_too_small = in["length_too_small"].(int)
		ret.Data_is_short = in["data_is_short"].(int)
		ret.Invalid_detect_mult = in["invalid_detect_mult"].(int)
		ret.Invalid_multipoint = in["invalid_multipoint"].(int)
		ret.Invalid_my_disc = in["invalid_my_disc"].(int)
		ret.Invalid_ttl = in["invalid_ttl"].(int)
		ret.Auth_length_invalid = in["auth_length_invalid"].(int)
		ret.Auth_mismatch = in["auth_mismatch"].(int)
		ret.Auth_type_mismatch = in["auth_type_mismatch"].(int)
		ret.Auth_key_id_mismatch = in["auth_key_id_mismatch"].(int)
		ret.Auth_key_mismatch = in["auth_key_mismatch"].(int)
		ret.Auth_seqnum_invalid = in["auth_seqnum_invalid"].(int)
		ret.Auth_failed = in["auth_failed"].(int)
		ret.Local_state_admin_down = in["local_state_admin_down"].(int)
		ret.Dest_unreachable = in["dest_unreachable"].(int)
		ret.No_ipv6_enable = in["no_ipv6_enable"].(int)
		ret.Other_error = in["other_error"].(int)
	}
	return ret
}

func dataToEndpointSystemBfdStats(d *schema.ResourceData) edpt.SystemBfdStats {
	var ret edpt.SystemBfdStats

	ret.Stats = getObjectSystemBfdStatsStats(d.Get("stats").([]interface{}))
	return ret
}
