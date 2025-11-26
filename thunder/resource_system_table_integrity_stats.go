package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTableIntegrityStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_table_integrity_stats`: Statistics for the object table-integrity\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemTableIntegrityStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arp_tbl_sync_start_ts_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync start time stamp master",
						},
						"nd6_tbl_sync_start_ts_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync start time stamp master",
						},
						"ipv4_fib_tbl_sync_start_ts_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync start time stamp master",
						},
						"ipv6_fib_tbl_sync_start_ts_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync start time stamp master",
						},
						"mac_tbl_sync_start_ts_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync start time stamp master",
						},
						"arp_tbl_sync_start_ts_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync start time stamp blade",
						},
						"nd6_tbl_sync_start_ts_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync start time stamp blade",
						},
						"ipv4_fib_tbl_sync_start_ts_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync start time stamp blade",
						},
						"ipv6_fib_tbl_sync_start_ts_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync start time stamp blade",
						},
						"mac_tbl_sync_start_ts_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync start time stamp blade",
						},
						"arp_tbl_sync_entries_sent_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries sent from master for T0 synchronization",
						},
						"nd6_tbl_sync_entries_sent_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries sent from master for T0 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_sent_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries sent from master for T0 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_sent_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries sent from master for T0 synchronization",
						},
						"mac_tbl_sync_entries_sent_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries sent from master for T0 synchronization",
						},
						"arp_tbl_sync_entries_rcvd_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries received on blade for T0 synchronization",
						},
						"nd6_tbl_sync_entries_rcvd_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries received on blade for T0 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_rcvd_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries received on blade for T0 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_rcvd_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries received on blade for T0 synchronization",
						},
						"mac_tbl_sync_entries_rcvd_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries received on blade for T0 synchronization",
						},
						"arp_tbl_sync_entries_added_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries added on blade for T0 synchronization",
						},
						"nd6_tbl_sync_entries_added_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries added on blade for T0 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_added_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries added on blade for T0 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_added_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries added on blade for T0 synchronization",
						},
						"mac_tbl_sync_entries_added_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries added on blade for T0 synchronization",
						},
						"arp_tbl_sync_entries_removed_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries removed on blade for T0 synchronization",
						},
						"nd6_tbl_sync_entries_removed_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries removed on blade for T0 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_removed_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries removed on blade for T0 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_removed_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries removed on blade for T0 synchronization",
						},
						"mac_tbl_sync_entries_removed_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries removed on blade for T0 synchronization",
						},
						"arp_tbl_sync_end_ts_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync end time stamp master for T0 synchronization",
						},
						"nd6_tbl_sync_end_ts_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync end time stamp master for T0 synchronization",
						},
						"ipv4_fib_tbl_sync_end_ts_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync end time stamp master for T0 synchronization",
						},
						"ipv6_fib_tbl_sync_end_ts_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync end time stamp master for T0 synchronization",
						},
						"mac_tbl_sync_end_ts_m_1st": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync end time stamp master for T0 synchronization",
						},
						"arp_tbl_sync_end_ts_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync end time stamp blade for T0 synchronization",
						},
						"nd6_tbl_sync_end_ts_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync end time stamp blade for T0 synchronization",
						},
						"ipv4_fib_tbl_sync_end_ts_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync end time stamp blade for T0 synchronization",
						},
						"ipv6_fib_tbl_sync_end_ts_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync end time stamp blade for T0 synchronization",
						},
						"mac_tbl_sync_end_ts_b_1st": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync end time stamp blade for T0 synchronization",
						},
						"arp_tbl_sync_start_ts_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync start time stamp master for T-1 synchronization",
						},
						"nd6_tbl_sync_start_ts_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync start time stamp master for T-1 synchronization",
						},
						"ipv4_fib_tbl_sync_start_ts_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync start time stamp master for T-1 synchronization",
						},
						"ipv6_fib_tbl_sync_start_ts_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync start time stamp master for T-1 synchronization",
						},
						"mac_tbl_sync_start_ts_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync start time stamp master for T-1 synchronization",
						},
						"arp_tbl_sync_start_ts_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync start time stamp blade for T-1 synchronization",
						},
						"nd6_tbl_sync_start_ts_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync start time stamp blade for T-1 synchronization",
						},
						"ipv4_fib_tbl_sync_start_ts_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync start time stamp blade for T-1 synchronization",
						},
						"ipv6_fib_tbl_sync_start_ts_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync start time stamp blade for T-1 synchronization",
						},
						"mac_tbl_sync_start_ts_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync start time stamp blade for T-1 synchronization",
						},
						"arp_tbl_sync_entries_sent_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries sent from master for T-1 synchronization",
						},
						"nd6_tbl_sync_entries_sent_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries sent from master for T-1 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_sent_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries sent from master for T-1 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_sent_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries sent from master for T-1 synchronization",
						},
						"mac_tbl_sync_entries_sent_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries sent from master for T-1 synchronization",
						},
						"arp_tbl_sync_entries_rcvd_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries received in blade for T-1 synchronization",
						},
						"nd6_tbl_sync_entries_rcvd_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries received in blade for T-1 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_rcvd_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries received in blade for T-1 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_rcvd_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries received in blade for T-1 synchronization",
						},
						"mac_tbl_sync_entries_rcvd_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries received in blade for T-1 synchronization",
						},
						"arp_tbl_sync_entries_added_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries added in blade for T-1 synchronization",
						},
						"nd6_tbl_sync_entries_added_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries added in blade for T-1 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_added_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries added in blade for T-1 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_added_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries added in blade for T-1 synchronization",
						},
						"mac_tbl_sync_entries_added_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries added in blade for T-1 synchronization",
						},
						"arp_tbl_sync_entries_removed_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries removed in blade for T-1 synchronization",
						},
						"nd6_tbl_sync_entries_removed_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries removed in blade for T-1 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_removed_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries removed in blade for T-1 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_removed_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries removed in blade for T-1 synchronization",
						},
						"mac_tbl_sync_entries_removed_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries removed in blade for T-1 synchronization",
						},
						"arp_tbl_sync_end_ts_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync end time stamp master for T-1 synchronization",
						},
						"nd6_tbl_sync_end_ts_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync end time stamp master for T-1 synchronization",
						},
						"ipv4_fib_tbl_sync_end_ts_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync end time stamp master for T-1 synchronization",
						},
						"ipv6_fib_tbl_sync_end_ts_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync end time stamp master for T-1 synchronization",
						},
						"mac_tbl_sync_end_ts_m_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync end time stamp master for T-1 synchronization",
						},
						"arp_tbl_sync_end_ts_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync end time stamp blade for T-1 synchronization",
						},
						"nd6_tbl_sync_end_ts_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync end time stamp blade for T-1 synchronization",
						},
						"ipv4_fib_tbl_sync_end_ts_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync end time stamp blade for T-1 synchronization",
						},
						"ipv6_fib_tbl_sync_end_ts_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync end time stamp blade for T-1 synchronization",
						},
						"mac_tbl_sync_end_ts_b_2nd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync end time stamp blade for T-1 synchronization",
						},
						"arp_tbl_sync_start_ts_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync start time stamp master for T-2 synchronization",
						},
						"nd6_tbl_sync_start_ts_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync start time stamp master for T-2 synchronization",
						},
						"ipv4_fib_tbl_sync_start_ts_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync start time stamp master for T-2 synchronization",
						},
						"ipv6_fib_tbl_sync_start_ts_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync start time stamp master for T-2 synchronization",
						},
						"mac_tbl_sync_start_ts_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync start time stamp master for T-2 synchronization",
						},
						"arp_tbl_sync_start_ts_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync start time stamp blade for T-2 synchronization",
						},
						"nd6_tbl_sync_start_ts_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync start time stamp blade for T-2 synchronization",
						},
						"ipv4_fib_tbl_sync_start_ts_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync start time stamp blade for T-2 synchronization",
						},
						"ipv6_fib_tbl_sync_start_ts_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync start time stamp blade for T-2 synchronization",
						},
						"mac_tbl_sync_start_ts_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync start time stamp blade for T-2 synchronization",
						},
						"arp_tbl_sync_entries_sent_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries sent from master for T-2 synchronization",
						},
						"nd6_tbl_sync_entries_sent_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries sent from master for T-2 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_sent_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries sent from master for T-2 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_sent_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries sent from master for T-2 synchronization",
						},
						"mac_tbl_sync_entries_sent_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries sent from master for T-2 synchronization",
						},
						"arp_tbl_sync_entries_rcvd_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries received in blade for T-2 synchronization",
						},
						"nd6_tbl_sync_entries_rcvd_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries received in blade for T-2 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_rcvd_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries received in blade for T-2 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_rcvd_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries received in blade for T-2 synchronization",
						},
						"mac_tbl_sync_entries_rcvd_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries received in blade for T-2 synchronization",
						},
						"arp_tbl_sync_entries_added_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries added in blade for T-2 synchronization",
						},
						"nd6_tbl_sync_entries_added_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries added in blade for T-2 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_added_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries added in blade for T-2 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_added_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries added in blade for T-2 synchronization",
						},
						"mac_tbl_sync_entries_added_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries added in blade for T-2 synchronization",
						},
						"arp_tbl_sync_entries_removed_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries removed in blade for T-2 synchronization",
						},
						"nd6_tbl_sync_entries_removed_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries removed in blade for T-2 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_removed_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries removed in blade for T-2 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_removed_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries removed in blade for T-2 synchronization",
						},
						"mac_tbl_sync_entries_removed_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries removed in blade for T-2 synchronization",
						},
						"arp_tbl_sync_end_ts_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync end time stamp master for T-2 synchronization",
						},
						"nd6_tbl_sync_end_ts_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync end time stamp master for T-2 synchronization",
						},
						"ipv4_fib_tbl_sync_end_ts_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync end time stamp master for T-2 synchronization",
						},
						"ipv6_fib_tbl_sync_end_ts_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync end time stamp master for T-2 synchronization",
						},
						"mac_tbl_sync_end_ts_m_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync end time stamp master for T-2 synchronization",
						},
						"arp_tbl_sync_end_ts_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync end time stamp blade for T-2 synchronization",
						},
						"nd6_tbl_sync_end_ts_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync end time stamp blade for T-2 synchronization",
						},
						"ipv4_fib_tbl_sync_end_ts_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync end time stamp blade for T-2 synchronization",
						},
						"ipv6_fib_tbl_sync_end_ts_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync end time stamp blade for T-2 synchronization",
						},
						"mac_tbl_sync_end_ts_b_3rd": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync end time stamp blade for T-2 synchronization",
						},
						"arp_tbl_sync_start_ts_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync start time stamp master for T-3 synchronization",
						},
						"nd6_tbl_sync_start_ts_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync start time stamp master for T-3 synchronization",
						},
						"ipv4_fib_tbl_sync_start_ts_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync start time stamp master for T-3 synchronization",
						},
						"ipv6_fib_tbl_sync_start_ts_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync start time stamp master for T-3 synchronization",
						},
						"mac_tbl_sync_start_ts_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync start time stamp master for T-3 synchronization",
						},
						"arp_tbl_sync_start_ts_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync start time stamp blade for T-3 synchronization",
						},
						"nd6_tbl_sync_start_ts_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync start time stamp blade for T-3 synchronization",
						},
						"ipv4_fib_tbl_sync_start_ts_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync start time stamp blade for T-3 synchronization",
						},
						"ipv6_fib_tbl_sync_start_ts_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync start time stamp blade for T-3 synchronization",
						},
						"mac_tbl_sync_start_ts_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync start time stamp blade for T-3 synchronization",
						},
						"arp_tbl_sync_entries_sent_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries sent from master for T-3 synchronization",
						},
						"nd6_tbl_sync_entries_sent_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries sent from master for T-3 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_sent_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries sent from master for T-3 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_sent_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries sent from master for T-3 synchronization",
						},
						"mac_tbl_sync_entries_sent_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries sent from master for T-3 synchronization",
						},
						"arp_tbl_sync_entries_rcvd_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries received in blade for T-3 synchronization",
						},
						"nd6_tbl_sync_entries_rcvd_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries received in blade for T-3 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_rcvd_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries received in blade for T-3 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_rcvd_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries received in blade for T-3 synchronization",
						},
						"mac_tbl_sync_entries_rcvd_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries received in blade for T-3 synchronization",
						},
						"arp_tbl_sync_entries_added_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries added in blade for T-3 synchronization",
						},
						"nd6_tbl_sync_entries_added_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries added in blade for T-3 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_added_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries added in blade for T-3 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_added_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries added in blade for T-3 synchronization",
						},
						"mac_tbl_sync_entries_added_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries added in blade for T-3 synchronization",
						},
						"arp_tbl_sync_entries_removed_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries removed in blade for T-3 synchronization",
						},
						"nd6_tbl_sync_entries_removed_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries removed in blade for T-3 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_removed_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries removed in blade for T-3 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_removed_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries removed in blade for T-3 synchronization",
						},
						"mac_tbl_sync_entries_removed_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries removed in blade for T-3 synchronization",
						},
						"arp_tbl_sync_end_ts_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync end time stamp master for T-3 synchronization",
						},
						"nd6_tbl_sync_end_ts_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync end time stamp master for T-3 synchronization",
						},
						"ipv4_fib_tbl_sync_end_ts_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync end time stamp master for T-3 synchronization",
						},
						"ipv6_fib_tbl_sync_end_ts_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync end time stamp master for T-3 synchronization",
						},
						"mac_tbl_sync_end_ts_m_4th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync end time stamp master for T-3 synchronization",
						},
						"arp_tbl_sync_end_ts_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync end time stamp blade for T-3 synchronization",
						},
						"nd6_tbl_sync_end_ts_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync end time stamp blade for T-3 synchronization",
						},
						"ipv4_fib_tbl_sync_end_ts_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync end time stamp blade for T-3 synchronization",
						},
						"ipv6_fib_tbl_sync_end_ts_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync end time stamp blade for T-3 synchronization",
						},
						"mac_tbl_sync_end_ts_b_4th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync end time stamp blade for T-3 synchronization",
						},
						"arp_tbl_sync_start_ts_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync start time stamp master for T-4 synchronization",
						},
						"nd6_tbl_sync_start_ts_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync start time stamp master for T-4 synchronization",
						},
						"ipv4_fib_tbl_sync_start_ts_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync start time stamp master for T-4 synchronization",
						},
						"ipv6_fib_tbl_sync_start_ts_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync start time stamp master for T-4 synchronization",
						},
						"mac_tbl_sync_start_ts_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync start time stamp master for T-4 synchronization",
						},
						"arp_tbl_sync_start_ts_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync start time stamp blade for T-4 synchronization",
						},
						"nd6_tbl_sync_start_ts_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync start time stamp blade for T-4 synchronization",
						},
						"ipv4_fib_tbl_sync_start_ts_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync start time stamp blade for T-4 synchronization",
						},
						"ipv6_fib_tbl_sync_start_ts_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync start time stamp blade for T-4 synchronization",
						},
						"mac_tbl_sync_start_ts_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync start time stamp blade for T-4 synchronization",
						},
						"arp_tbl_sync_entries_sent_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync start time stamp blade for T-4 synchronization",
						},
						"nd6_tbl_sync_entries_sent_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync start time stamp blade for T-4 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_sent_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync start time stamp blade for T-4 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_sent_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync start time stamp blade for T-4 synchronization",
						},
						"mac_tbl_sync_entries_sent_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync start time stamp blade for T-4 synchronization",
						},
						"arp_tbl_sync_entries_rcvd_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries received in blade for T-4 synchronization",
						},
						"nd6_tbl_sync_entries_rcvd_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries received in blade for T-4 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_rcvd_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries received in blade for T-4 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_rcvd_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries received in blade for T-4 synchronization",
						},
						"mac_tbl_sync_entries_rcvd_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries received in blade for T-4 synchronization",
						},
						"arp_tbl_sync_entries_added_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries added in blade for T-4 synchronization",
						},
						"nd6_tbl_sync_entries_added_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries added in blade for T-4 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_added_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries added in blade for T-4 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_added_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries added in blade for T-4 synchronization",
						},
						"mac_tbl_sync_entries_added_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries added in blade for T-4 synchronization",
						},
						"arp_tbl_sync_entries_removed_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table entries removed in blade for T-4 synchronization",
						},
						"nd6_tbl_sync_entries_removed_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table entries removed in blade for T-4 synchronization",
						},
						"ipv4_fib_tbl_sync_entries_removed_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table entries removed in blade for T-4 synchronization",
						},
						"ipv6_fib_tbl_sync_entries_removed_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table entries removed in blade for T-4 synchronization",
						},
						"mac_tbl_sync_entries_removed_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table entries removed in blade for T-4 synchronization",
						},
						"arp_tbl_sync_end_ts_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync end time stamp master for T-4 synchronization",
						},
						"nd6_tbl_sync_end_ts_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync end time stamp master for T-4 synchronization",
						},
						"ipv4_fib_tbl_sync_end_ts_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync end time stamp master for T-4 synchronization",
						},
						"ipv6_fib_tbl_sync_end_ts_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync end time stamp master for T-4 synchronization",
						},
						"mac_tbl_sync_end_ts_m_5th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync end time stamp master for T-4 synchronization",
						},
						"arp_tbl_sync_end_ts_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync end time stamp blade for T-4 synchronization",
						},
						"nd6_tbl_sync_end_ts_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync end time stamp blade for T-4 synchronization",
						},
						"ipv4_fib_tbl_sync_end_ts_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync end time stamp blade for T-4 synchronization",
						},
						"ipv6_fib_tbl_sync_end_ts_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync end time stamp blade for T-4 synchronization",
						},
						"mac_tbl_sync_end_ts_b_5th": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync end time stamp blade for T-4 synchronization",
						},
						"arp_tbl_sync_m": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync count in master",
						},
						"nd6_tbl_sync_m": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync count in master",
						},
						"ipv4_fib_tbl_sync_m": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync count in master",
						},
						"ipv6_fib_tbl_sync_m": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync count in master",
						},
						"mac_tbl_sync_m": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync count in master",
						},
						"arp_tbl_sync_b": {
							Type: schema.TypeInt, Optional: true, Description: "arp table sync count in blade",
						},
						"nd6_tbl_sync_b": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table sync count in blade",
						},
						"ipv4_fib_tbl_sync_b": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table sync count in blade",
						},
						"ipv6_fib_tbl_sync_b": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table sync count in blade",
						},
						"mac_tbl_sync_b": {
							Type: schema.TypeInt, Optional: true, Description: "mac table sync count in blade",
						},
						"arp_tbl_cksum_m": {
							Type: schema.TypeInt, Optional: true, Description: "arp table checksum count in master",
						},
						"nd6_tbl_cksum_m": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table checksum count in master",
						},
						"ipv4_fib_tbl_cksum_m": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table checksum count in master",
						},
						"ipv6_fib_tbl_cksum_m": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table checksum count in master",
						},
						"mac_tbl_cksum_m": {
							Type: schema.TypeInt, Optional: true, Description: "mac table checksum count in master",
						},
						"arp_tbl_cksum_b": {
							Type: schema.TypeInt, Optional: true, Description: "arp table checksum count in blade",
						},
						"nd6_tbl_cksum_b": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table checksum count in blade",
						},
						"ipv4_fib_tbl_cksum_b": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table checksum count in blade",
						},
						"ipv6_fib_tbl_cksum_b": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table checksum count in blade",
						},
						"mac_tbl_cksum_b": {
							Type: schema.TypeInt, Optional: true, Description: "mac table checksum count in blade",
						},
						"arp_tbl_cksum_mismatch_b": {
							Type: schema.TypeInt, Optional: true, Description: "arp table checksum mismatch count in blade",
						},
						"nd6_tbl_cksum_mismatch_b": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table checksum mismatch count in blade",
						},
						"ipv4_fib_tbl_cksum_mismatch_b": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table checksum mismatch count in blade",
						},
						"ipv6_fib_tbl_cksum_mismatch_b": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table checksum mismatch count in blade",
						},
						"mac_tbl_cksum_mismatch_b": {
							Type: schema.TypeInt, Optional: true, Description: "mac table checksum mismatch count in blade",
						},
						"arp_tbl_cksum_cancel_m": {
							Type: schema.TypeInt, Optional: true, Description: "arp table checksum cancelled count in master",
						},
						"nd6_tbl_cksum_cancel_m": {
							Type: schema.TypeInt, Optional: true, Description: "nd6 table checksum cancelled count in master",
						},
						"ipv4_fib_tbl_cksum_cancel_m": {
							Type: schema.TypeInt, Optional: true, Description: "ipv4-fib table checksum cancelled count in master",
						},
						"ipv6_fib_tbl_cksum_cancel_m": {
							Type: schema.TypeInt, Optional: true, Description: "ipv6-fib table checksum cancelled count in master",
						},
						"mac_tbl_cksum_cancel_m": {
							Type: schema.TypeInt, Optional: true, Description: "mac table checksum cancelled count in master",
						},
					},
				},
			},
		},
	}
}

func resourceSystemTableIntegrityStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTableIntegrityStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTableIntegrityStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemTableIntegrityStatsStats := setObjectSystemTableIntegrityStatsStats(res)
		d.Set("stats", SystemTableIntegrityStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemTableIntegrityStatsStats(ret edpt.DataSystemTableIntegrityStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"arp_tbl_sync_start_ts_m_1st":             ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncStartTsM1st,
			"nd6_tbl_sync_start_ts_m_1st":             ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncStartTsM1st,
			"ipv4_fib_tbl_sync_start_ts_m_1st":        ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncStartTsM1st,
			"ipv6_fib_tbl_sync_start_ts_m_1st":        ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncStartTsM1st,
			"mac_tbl_sync_start_ts_m_1st":             ret.DtSystemTableIntegrityStats.Stats.MacTblSyncStartTsM1st,
			"arp_tbl_sync_start_ts_b_1st":             ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncStartTsB1st,
			"nd6_tbl_sync_start_ts_b_1st":             ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncStartTsB1st,
			"ipv4_fib_tbl_sync_start_ts_b_1st":        ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncStartTsB1st,
			"ipv6_fib_tbl_sync_start_ts_b_1st":        ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncStartTsB1st,
			"mac_tbl_sync_start_ts_b_1st":             ret.DtSystemTableIntegrityStats.Stats.MacTblSyncStartTsB1st,
			"arp_tbl_sync_entries_sent_m_1st":         ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesSentM1st,
			"nd6_tbl_sync_entries_sent_m_1st":         ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesSentM1st,
			"ipv4_fib_tbl_sync_entries_sent_m_1st":    ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesSentM1st,
			"ipv6_fib_tbl_sync_entries_sent_m_1st":    ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesSentM1st,
			"mac_tbl_sync_entries_sent_m_1st":         ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesSentM1st,
			"arp_tbl_sync_entries_rcvd_b_1st":         ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesRcvdB1st,
			"nd6_tbl_sync_entries_rcvd_b_1st":         ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesRcvdB1st,
			"ipv4_fib_tbl_sync_entries_rcvd_b_1st":    ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesRcvdB1st,
			"ipv6_fib_tbl_sync_entries_rcvd_b_1st":    ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesRcvdB1st,
			"mac_tbl_sync_entries_rcvd_b_1st":         ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesRcvdB1st,
			"arp_tbl_sync_entries_added_b_1st":        ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesAddedB1st,
			"nd6_tbl_sync_entries_added_b_1st":        ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesAddedB1st,
			"ipv4_fib_tbl_sync_entries_added_b_1st":   ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesAddedB1st,
			"ipv6_fib_tbl_sync_entries_added_b_1st":   ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesAddedB1st,
			"mac_tbl_sync_entries_added_b_1st":        ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesAddedB1st,
			"arp_tbl_sync_entries_removed_b_1st":      ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesRemovedB1st,
			"nd6_tbl_sync_entries_removed_b_1st":      ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesRemovedB1st,
			"ipv4_fib_tbl_sync_entries_removed_b_1st": ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesRemovedB1st,
			"ipv6_fib_tbl_sync_entries_removed_b_1st": ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesRemovedB1st,
			"mac_tbl_sync_entries_removed_b_1st":      ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesRemovedB1st,
			"arp_tbl_sync_end_ts_m_1st":               ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEndTsM1st,
			"nd6_tbl_sync_end_ts_m_1st":               ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEndTsM1st,
			"ipv4_fib_tbl_sync_end_ts_m_1st":          ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEndTsM1st,
			"ipv6_fib_tbl_sync_end_ts_m_1st":          ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEndTsM1st,
			"mac_tbl_sync_end_ts_m_1st":               ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEndTsM1st,
			"arp_tbl_sync_end_ts_b_1st":               ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEndTsB1st,
			"nd6_tbl_sync_end_ts_b_1st":               ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEndTsB1st,
			"ipv4_fib_tbl_sync_end_ts_b_1st":          ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEndTsB1st,
			"ipv6_fib_tbl_sync_end_ts_b_1st":          ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEndTsB1st,
			"mac_tbl_sync_end_ts_b_1st":               ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEndTsB1st,
			"arp_tbl_sync_start_ts_m_2nd":             ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncStartTsM2nd,
			"nd6_tbl_sync_start_ts_m_2nd":             ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncStartTsM2nd,
			"ipv4_fib_tbl_sync_start_ts_m_2nd":        ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncStartTsM2nd,
			"ipv6_fib_tbl_sync_start_ts_m_2nd":        ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncStartTsM2nd,
			"mac_tbl_sync_start_ts_m_2nd":             ret.DtSystemTableIntegrityStats.Stats.MacTblSyncStartTsM2nd,
			"arp_tbl_sync_start_ts_b_2nd":             ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncStartTsB2nd,
			"nd6_tbl_sync_start_ts_b_2nd":             ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncStartTsB2nd,
			"ipv4_fib_tbl_sync_start_ts_b_2nd":        ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncStartTsB2nd,
			"ipv6_fib_tbl_sync_start_ts_b_2nd":        ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncStartTsB2nd,
			"mac_tbl_sync_start_ts_b_2nd":             ret.DtSystemTableIntegrityStats.Stats.MacTblSyncStartTsB2nd,
			"arp_tbl_sync_entries_sent_m_2nd":         ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesSentM2nd,
			"nd6_tbl_sync_entries_sent_m_2nd":         ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesSentM2nd,
			"ipv4_fib_tbl_sync_entries_sent_m_2nd":    ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesSentM2nd,
			"ipv6_fib_tbl_sync_entries_sent_m_2nd":    ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesSentM2nd,
			"mac_tbl_sync_entries_sent_m_2nd":         ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesSentM2nd,
			"arp_tbl_sync_entries_rcvd_b_2nd":         ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesRcvdB2nd,
			"nd6_tbl_sync_entries_rcvd_b_2nd":         ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesRcvdB2nd,
			"ipv4_fib_tbl_sync_entries_rcvd_b_2nd":    ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesRcvdB2nd,
			"ipv6_fib_tbl_sync_entries_rcvd_b_2nd":    ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesRcvdB2nd,
			"mac_tbl_sync_entries_rcvd_b_2nd":         ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesRcvdB2nd,
			"arp_tbl_sync_entries_added_b_2nd":        ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesAddedB2nd,
			"nd6_tbl_sync_entries_added_b_2nd":        ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesAddedB2nd,
			"ipv4_fib_tbl_sync_entries_added_b_2nd":   ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesAddedB2nd,
			"ipv6_fib_tbl_sync_entries_added_b_2nd":   ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesAddedB2nd,
			"mac_tbl_sync_entries_added_b_2nd":        ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesAddedB2nd,
			"arp_tbl_sync_entries_removed_b_2nd":      ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesRemovedB2nd,
			"nd6_tbl_sync_entries_removed_b_2nd":      ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesRemovedB2nd,
			"ipv4_fib_tbl_sync_entries_removed_b_2nd": ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesRemovedB2nd,
			"ipv6_fib_tbl_sync_entries_removed_b_2nd": ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesRemovedB2nd,
			"mac_tbl_sync_entries_removed_b_2nd":      ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesRemovedB2nd,
			"arp_tbl_sync_end_ts_m_2nd":               ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEndTsM2nd,
			"nd6_tbl_sync_end_ts_m_2nd":               ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEndTsM2nd,
			"ipv4_fib_tbl_sync_end_ts_m_2nd":          ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEndTsM2nd,
			"ipv6_fib_tbl_sync_end_ts_m_2nd":          ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEndTsM2nd,
			"mac_tbl_sync_end_ts_m_2nd":               ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEndTsM2nd,
			"arp_tbl_sync_end_ts_b_2nd":               ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEndTsB2nd,
			"nd6_tbl_sync_end_ts_b_2nd":               ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEndTsB2nd,
			"ipv4_fib_tbl_sync_end_ts_b_2nd":          ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEndTsB2nd,
			"ipv6_fib_tbl_sync_end_ts_b_2nd":          ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEndTsB2nd,
			"mac_tbl_sync_end_ts_b_2nd":               ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEndTsB2nd,
			"arp_tbl_sync_start_ts_m_3rd":             ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncStartTsM3rd,
			"nd6_tbl_sync_start_ts_m_3rd":             ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncStartTsM3rd,
			"ipv4_fib_tbl_sync_start_ts_m_3rd":        ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncStartTsM3rd,
			"ipv6_fib_tbl_sync_start_ts_m_3rd":        ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncStartTsM3rd,
			"mac_tbl_sync_start_ts_m_3rd":             ret.DtSystemTableIntegrityStats.Stats.MacTblSyncStartTsM3rd,
			"arp_tbl_sync_start_ts_b_3rd":             ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncStartTsB3rd,
			"nd6_tbl_sync_start_ts_b_3rd":             ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncStartTsB3rd,
			"ipv4_fib_tbl_sync_start_ts_b_3rd":        ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncStartTsB3rd,
			"ipv6_fib_tbl_sync_start_ts_b_3rd":        ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncStartTsB3rd,
			"mac_tbl_sync_start_ts_b_3rd":             ret.DtSystemTableIntegrityStats.Stats.MacTblSyncStartTsB3rd,
			"arp_tbl_sync_entries_sent_m_3rd":         ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesSentM3rd,
			"nd6_tbl_sync_entries_sent_m_3rd":         ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesSentM3rd,
			"ipv4_fib_tbl_sync_entries_sent_m_3rd":    ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesSentM3rd,
			"ipv6_fib_tbl_sync_entries_sent_m_3rd":    ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesSentM3rd,
			"mac_tbl_sync_entries_sent_m_3rd":         ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesSentM3rd,
			"arp_tbl_sync_entries_rcvd_b_3rd":         ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesRcvdB3rd,
			"nd6_tbl_sync_entries_rcvd_b_3rd":         ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesRcvdB3rd,
			"ipv4_fib_tbl_sync_entries_rcvd_b_3rd":    ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesRcvdB3rd,
			"ipv6_fib_tbl_sync_entries_rcvd_b_3rd":    ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesRcvdB3rd,
			"mac_tbl_sync_entries_rcvd_b_3rd":         ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesRcvdB3rd,
			"arp_tbl_sync_entries_added_b_3rd":        ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesAddedB3rd,
			"nd6_tbl_sync_entries_added_b_3rd":        ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesAddedB3rd,
			"ipv4_fib_tbl_sync_entries_added_b_3rd":   ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesAddedB3rd,
			"ipv6_fib_tbl_sync_entries_added_b_3rd":   ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesAddedB3rd,
			"mac_tbl_sync_entries_added_b_3rd":        ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesAddedB3rd,
			"arp_tbl_sync_entries_removed_b_3rd":      ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesRemovedB3rd,
			"nd6_tbl_sync_entries_removed_b_3rd":      ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesRemovedB3rd,
			"ipv4_fib_tbl_sync_entries_removed_b_3rd": ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesRemovedB3rd,
			"ipv6_fib_tbl_sync_entries_removed_b_3rd": ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesRemovedB3rd,
			"mac_tbl_sync_entries_removed_b_3rd":      ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesRemovedB3rd,
			"arp_tbl_sync_end_ts_m_3rd":               ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEndTsM3rd,
			"nd6_tbl_sync_end_ts_m_3rd":               ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEndTsM3rd,
			"ipv4_fib_tbl_sync_end_ts_m_3rd":          ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEndTsM3rd,
			"ipv6_fib_tbl_sync_end_ts_m_3rd":          ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEndTsM3rd,
			"mac_tbl_sync_end_ts_m_3rd":               ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEndTsM3rd,
			"arp_tbl_sync_end_ts_b_3rd":               ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEndTsB3rd,
			"nd6_tbl_sync_end_ts_b_3rd":               ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEndTsB3rd,
			"ipv4_fib_tbl_sync_end_ts_b_3rd":          ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEndTsB3rd,
			"ipv6_fib_tbl_sync_end_ts_b_3rd":          ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEndTsB3rd,
			"mac_tbl_sync_end_ts_b_3rd":               ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEndTsB3rd,
			"arp_tbl_sync_start_ts_m_4th":             ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncStartTsM4th,
			"nd6_tbl_sync_start_ts_m_4th":             ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncStartTsM4th,
			"ipv4_fib_tbl_sync_start_ts_m_4th":        ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncStartTsM4th,
			"ipv6_fib_tbl_sync_start_ts_m_4th":        ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncStartTsM4th,
			"mac_tbl_sync_start_ts_m_4th":             ret.DtSystemTableIntegrityStats.Stats.MacTblSyncStartTsM4th,
			"arp_tbl_sync_start_ts_b_4th":             ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncStartTsB4th,
			"nd6_tbl_sync_start_ts_b_4th":             ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncStartTsB4th,
			"ipv4_fib_tbl_sync_start_ts_b_4th":        ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncStartTsB4th,
			"ipv6_fib_tbl_sync_start_ts_b_4th":        ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncStartTsB4th,
			"mac_tbl_sync_start_ts_b_4th":             ret.DtSystemTableIntegrityStats.Stats.MacTblSyncStartTsB4th,
			"arp_tbl_sync_entries_sent_m_4th":         ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesSentM4th,
			"nd6_tbl_sync_entries_sent_m_4th":         ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesSentM4th,
			"ipv4_fib_tbl_sync_entries_sent_m_4th":    ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesSentM4th,
			"ipv6_fib_tbl_sync_entries_sent_m_4th":    ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesSentM4th,
			"mac_tbl_sync_entries_sent_m_4th":         ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesSentM4th,
			"arp_tbl_sync_entries_rcvd_b_4th":         ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesRcvdB4th,
			"nd6_tbl_sync_entries_rcvd_b_4th":         ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesRcvdB4th,
			"ipv4_fib_tbl_sync_entries_rcvd_b_4th":    ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesRcvdB4th,
			"ipv6_fib_tbl_sync_entries_rcvd_b_4th":    ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesRcvdB4th,
			"mac_tbl_sync_entries_rcvd_b_4th":         ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesRcvdB4th,
			"arp_tbl_sync_entries_added_b_4th":        ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesAddedB4th,
			"nd6_tbl_sync_entries_added_b_4th":        ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesAddedB4th,
			"ipv4_fib_tbl_sync_entries_added_b_4th":   ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesAddedB4th,
			"ipv6_fib_tbl_sync_entries_added_b_4th":   ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesAddedB4th,
			"mac_tbl_sync_entries_added_b_4th":        ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesAddedB4th,
			"arp_tbl_sync_entries_removed_b_4th":      ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesRemovedB4th,
			"nd6_tbl_sync_entries_removed_b_4th":      ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesRemovedB4th,
			"ipv4_fib_tbl_sync_entries_removed_b_4th": ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesRemovedB4th,
			"ipv6_fib_tbl_sync_entries_removed_b_4th": ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesRemovedB4th,
			"mac_tbl_sync_entries_removed_b_4th":      ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesRemovedB4th,
			"arp_tbl_sync_end_ts_m_4th":               ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEndTsM4th,
			"nd6_tbl_sync_end_ts_m_4th":               ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEndTsM4th,
			"ipv4_fib_tbl_sync_end_ts_m_4th":          ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEndTsM4th,
			"ipv6_fib_tbl_sync_end_ts_m_4th":          ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEndTsM4th,
			"mac_tbl_sync_end_ts_m_4th":               ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEndTsM4th,
			"arp_tbl_sync_end_ts_b_4th":               ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEndTsB4th,
			"nd6_tbl_sync_end_ts_b_4th":               ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEndTsB4th,
			"ipv4_fib_tbl_sync_end_ts_b_4th":          ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEndTsB4th,
			"ipv6_fib_tbl_sync_end_ts_b_4th":          ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEndTsB4th,
			"mac_tbl_sync_end_ts_b_4th":               ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEndTsB4th,
			"arp_tbl_sync_start_ts_m_5th":             ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncStartTsM5th,
			"nd6_tbl_sync_start_ts_m_5th":             ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncStartTsM5th,
			"ipv4_fib_tbl_sync_start_ts_m_5th":        ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncStartTsM5th,
			"ipv6_fib_tbl_sync_start_ts_m_5th":        ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncStartTsM5th,
			"mac_tbl_sync_start_ts_m_5th":             ret.DtSystemTableIntegrityStats.Stats.MacTblSyncStartTsM5th,
			"arp_tbl_sync_start_ts_b_5th":             ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncStartTsB5th,
			"nd6_tbl_sync_start_ts_b_5th":             ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncStartTsB5th,
			"ipv4_fib_tbl_sync_start_ts_b_5th":        ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncStartTsB5th,
			"ipv6_fib_tbl_sync_start_ts_b_5th":        ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncStartTsB5th,
			"mac_tbl_sync_start_ts_b_5th":             ret.DtSystemTableIntegrityStats.Stats.MacTblSyncStartTsB5th,
			"arp_tbl_sync_entries_sent_m_5th":         ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesSentM5th,
			"nd6_tbl_sync_entries_sent_m_5th":         ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesSentM5th,
			"ipv4_fib_tbl_sync_entries_sent_m_5th":    ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesSentM5th,
			"ipv6_fib_tbl_sync_entries_sent_m_5th":    ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesSentM5th,
			"mac_tbl_sync_entries_sent_m_5th":         ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesSentM5th,
			"arp_tbl_sync_entries_rcvd_b_5th":         ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesRcvdB5th,
			"nd6_tbl_sync_entries_rcvd_b_5th":         ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesRcvdB5th,
			"ipv4_fib_tbl_sync_entries_rcvd_b_5th":    ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesRcvdB5th,
			"ipv6_fib_tbl_sync_entries_rcvd_b_5th":    ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesRcvdB5th,
			"mac_tbl_sync_entries_rcvd_b_5th":         ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesRcvdB5th,
			"arp_tbl_sync_entries_added_b_5th":        ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesAddedB5th,
			"nd6_tbl_sync_entries_added_b_5th":        ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesAddedB5th,
			"ipv4_fib_tbl_sync_entries_added_b_5th":   ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesAddedB5th,
			"ipv6_fib_tbl_sync_entries_added_b_5th":   ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesAddedB5th,
			"mac_tbl_sync_entries_added_b_5th":        ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesAddedB5th,
			"arp_tbl_sync_entries_removed_b_5th":      ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEntriesRemovedB5th,
			"nd6_tbl_sync_entries_removed_b_5th":      ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEntriesRemovedB5th,
			"ipv4_fib_tbl_sync_entries_removed_b_5th": ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEntriesRemovedB5th,
			"ipv6_fib_tbl_sync_entries_removed_b_5th": ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEntriesRemovedB5th,
			"mac_tbl_sync_entries_removed_b_5th":      ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEntriesRemovedB5th,
			"arp_tbl_sync_end_ts_m_5th":               ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEndTsM5th,
			"nd6_tbl_sync_end_ts_m_5th":               ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEndTsM5th,
			"ipv4_fib_tbl_sync_end_ts_m_5th":          ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEndTsM5th,
			"ipv6_fib_tbl_sync_end_ts_m_5th":          ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEndTsM5th,
			"mac_tbl_sync_end_ts_m_5th":               ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEndTsM5th,
			"arp_tbl_sync_end_ts_b_5th":               ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncEndTsB5th,
			"nd6_tbl_sync_end_ts_b_5th":               ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncEndTsB5th,
			"ipv4_fib_tbl_sync_end_ts_b_5th":          ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncEndTsB5th,
			"ipv6_fib_tbl_sync_end_ts_b_5th":          ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncEndTsB5th,
			"mac_tbl_sync_end_ts_b_5th":               ret.DtSystemTableIntegrityStats.Stats.MacTblSyncEndTsB5th,
			"arp_tbl_sync_m":                          ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncM,
			"nd6_tbl_sync_m":                          ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncM,
			"ipv4_fib_tbl_sync_m":                     ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncM,
			"ipv6_fib_tbl_sync_m":                     ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncM,
			"mac_tbl_sync_m":                          ret.DtSystemTableIntegrityStats.Stats.MacTblSyncM,
			"arp_tbl_sync_b":                          ret.DtSystemTableIntegrityStats.Stats.ArpTblSyncB,
			"nd6_tbl_sync_b":                          ret.DtSystemTableIntegrityStats.Stats.Nd6TblSyncB,
			"ipv4_fib_tbl_sync_b":                     ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblSyncB,
			"ipv6_fib_tbl_sync_b":                     ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblSyncB,
			"mac_tbl_sync_b":                          ret.DtSystemTableIntegrityStats.Stats.MacTblSyncB,
			"arp_tbl_cksum_m":                         ret.DtSystemTableIntegrityStats.Stats.ArpTblCksumM,
			"nd6_tbl_cksum_m":                         ret.DtSystemTableIntegrityStats.Stats.Nd6TblCksumM,
			"ipv4_fib_tbl_cksum_m":                    ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblCksumM,
			"ipv6_fib_tbl_cksum_m":                    ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblCksumM,
			"mac_tbl_cksum_m":                         ret.DtSystemTableIntegrityStats.Stats.MacTblCksumM,
			"arp_tbl_cksum_b":                         ret.DtSystemTableIntegrityStats.Stats.ArpTblCksumB,
			"nd6_tbl_cksum_b":                         ret.DtSystemTableIntegrityStats.Stats.Nd6TblCksumB,
			"ipv4_fib_tbl_cksum_b":                    ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblCksumB,
			"ipv6_fib_tbl_cksum_b":                    ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblCksumB,
			"mac_tbl_cksum_b":                         ret.DtSystemTableIntegrityStats.Stats.MacTblCksumB,
			"arp_tbl_cksum_mismatch_b":                ret.DtSystemTableIntegrityStats.Stats.ArpTblCksumMismatchB,
			"nd6_tbl_cksum_mismatch_b":                ret.DtSystemTableIntegrityStats.Stats.Nd6TblCksumMismatchB,
			"ipv4_fib_tbl_cksum_mismatch_b":           ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblCksumMismatchB,
			"ipv6_fib_tbl_cksum_mismatch_b":           ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblCksumMismatchB,
			"mac_tbl_cksum_mismatch_b":                ret.DtSystemTableIntegrityStats.Stats.MacTblCksumMismatchB,
			"arp_tbl_cksum_cancel_m":                  ret.DtSystemTableIntegrityStats.Stats.ArpTblCksumCancelM,
			"nd6_tbl_cksum_cancel_m":                  ret.DtSystemTableIntegrityStats.Stats.Nd6TblCksumCancelM,
			"ipv4_fib_tbl_cksum_cancel_m":             ret.DtSystemTableIntegrityStats.Stats.Ipv4FibTblCksumCancelM,
			"ipv6_fib_tbl_cksum_cancel_m":             ret.DtSystemTableIntegrityStats.Stats.Ipv6FibTblCksumCancelM,
			"mac_tbl_cksum_cancel_m":                  ret.DtSystemTableIntegrityStats.Stats.MacTblCksumCancelM,
		},
	}
}

func getObjectSystemTableIntegrityStatsStats(d []interface{}) edpt.SystemTableIntegrityStatsStats {

	count1 := len(d)
	var ret edpt.SystemTableIntegrityStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ArpTblSyncStartTsM1st = in["arp_tbl_sync_start_ts_m_1st"].(int)
		ret.Nd6TblSyncStartTsM1st = in["nd6_tbl_sync_start_ts_m_1st"].(int)
		ret.Ipv4FibTblSyncStartTsM1st = in["ipv4_fib_tbl_sync_start_ts_m_1st"].(int)
		ret.Ipv6FibTblSyncStartTsM1st = in["ipv6_fib_tbl_sync_start_ts_m_1st"].(int)
		ret.MacTblSyncStartTsM1st = in["mac_tbl_sync_start_ts_m_1st"].(int)
		ret.ArpTblSyncStartTsB1st = in["arp_tbl_sync_start_ts_b_1st"].(int)
		ret.Nd6TblSyncStartTsB1st = in["nd6_tbl_sync_start_ts_b_1st"].(int)
		ret.Ipv4FibTblSyncStartTsB1st = in["ipv4_fib_tbl_sync_start_ts_b_1st"].(int)
		ret.Ipv6FibTblSyncStartTsB1st = in["ipv6_fib_tbl_sync_start_ts_b_1st"].(int)
		ret.MacTblSyncStartTsB1st = in["mac_tbl_sync_start_ts_b_1st"].(int)
		ret.ArpTblSyncEntriesSentM1st = in["arp_tbl_sync_entries_sent_m_1st"].(int)
		ret.Nd6TblSyncEntriesSentM1st = in["nd6_tbl_sync_entries_sent_m_1st"].(int)
		ret.Ipv4FibTblSyncEntriesSentM1st = in["ipv4_fib_tbl_sync_entries_sent_m_1st"].(int)
		ret.Ipv6FibTblSyncEntriesSentM1st = in["ipv6_fib_tbl_sync_entries_sent_m_1st"].(int)
		ret.MacTblSyncEntriesSentM1st = in["mac_tbl_sync_entries_sent_m_1st"].(int)
		ret.ArpTblSyncEntriesRcvdB1st = in["arp_tbl_sync_entries_rcvd_b_1st"].(int)
		ret.Nd6TblSyncEntriesRcvdB1st = in["nd6_tbl_sync_entries_rcvd_b_1st"].(int)
		ret.Ipv4FibTblSyncEntriesRcvdB1st = in["ipv4_fib_tbl_sync_entries_rcvd_b_1st"].(int)
		ret.Ipv6FibTblSyncEntriesRcvdB1st = in["ipv6_fib_tbl_sync_entries_rcvd_b_1st"].(int)
		ret.MacTblSyncEntriesRcvdB1st = in["mac_tbl_sync_entries_rcvd_b_1st"].(int)
		ret.ArpTblSyncEntriesAddedB1st = in["arp_tbl_sync_entries_added_b_1st"].(int)
		ret.Nd6TblSyncEntriesAddedB1st = in["nd6_tbl_sync_entries_added_b_1st"].(int)
		ret.Ipv4FibTblSyncEntriesAddedB1st = in["ipv4_fib_tbl_sync_entries_added_b_1st"].(int)
		ret.Ipv6FibTblSyncEntriesAddedB1st = in["ipv6_fib_tbl_sync_entries_added_b_1st"].(int)
		ret.MacTblSyncEntriesAddedB1st = in["mac_tbl_sync_entries_added_b_1st"].(int)
		ret.ArpTblSyncEntriesRemovedB1st = in["arp_tbl_sync_entries_removed_b_1st"].(int)
		ret.Nd6TblSyncEntriesRemovedB1st = in["nd6_tbl_sync_entries_removed_b_1st"].(int)
		ret.Ipv4FibTblSyncEntriesRemovedB1st = in["ipv4_fib_tbl_sync_entries_removed_b_1st"].(int)
		ret.Ipv6FibTblSyncEntriesRemovedB1st = in["ipv6_fib_tbl_sync_entries_removed_b_1st"].(int)
		ret.MacTblSyncEntriesRemovedB1st = in["mac_tbl_sync_entries_removed_b_1st"].(int)
		ret.ArpTblSyncEndTsM1st = in["arp_tbl_sync_end_ts_m_1st"].(int)
		ret.Nd6TblSyncEndTsM1st = in["nd6_tbl_sync_end_ts_m_1st"].(int)
		ret.Ipv4FibTblSyncEndTsM1st = in["ipv4_fib_tbl_sync_end_ts_m_1st"].(int)
		ret.Ipv6FibTblSyncEndTsM1st = in["ipv6_fib_tbl_sync_end_ts_m_1st"].(int)
		ret.MacTblSyncEndTsM1st = in["mac_tbl_sync_end_ts_m_1st"].(int)
		ret.ArpTblSyncEndTsB1st = in["arp_tbl_sync_end_ts_b_1st"].(int)
		ret.Nd6TblSyncEndTsB1st = in["nd6_tbl_sync_end_ts_b_1st"].(int)
		ret.Ipv4FibTblSyncEndTsB1st = in["ipv4_fib_tbl_sync_end_ts_b_1st"].(int)
		ret.Ipv6FibTblSyncEndTsB1st = in["ipv6_fib_tbl_sync_end_ts_b_1st"].(int)
		ret.MacTblSyncEndTsB1st = in["mac_tbl_sync_end_ts_b_1st"].(int)
		ret.ArpTblSyncStartTsM2nd = in["arp_tbl_sync_start_ts_m_2nd"].(int)
		ret.Nd6TblSyncStartTsM2nd = in["nd6_tbl_sync_start_ts_m_2nd"].(int)
		ret.Ipv4FibTblSyncStartTsM2nd = in["ipv4_fib_tbl_sync_start_ts_m_2nd"].(int)
		ret.Ipv6FibTblSyncStartTsM2nd = in["ipv6_fib_tbl_sync_start_ts_m_2nd"].(int)
		ret.MacTblSyncStartTsM2nd = in["mac_tbl_sync_start_ts_m_2nd"].(int)
		ret.ArpTblSyncStartTsB2nd = in["arp_tbl_sync_start_ts_b_2nd"].(int)
		ret.Nd6TblSyncStartTsB2nd = in["nd6_tbl_sync_start_ts_b_2nd"].(int)
		ret.Ipv4FibTblSyncStartTsB2nd = in["ipv4_fib_tbl_sync_start_ts_b_2nd"].(int)
		ret.Ipv6FibTblSyncStartTsB2nd = in["ipv6_fib_tbl_sync_start_ts_b_2nd"].(int)
		ret.MacTblSyncStartTsB2nd = in["mac_tbl_sync_start_ts_b_2nd"].(int)
		ret.ArpTblSyncEntriesSentM2nd = in["arp_tbl_sync_entries_sent_m_2nd"].(int)
		ret.Nd6TblSyncEntriesSentM2nd = in["nd6_tbl_sync_entries_sent_m_2nd"].(int)
		ret.Ipv4FibTblSyncEntriesSentM2nd = in["ipv4_fib_tbl_sync_entries_sent_m_2nd"].(int)
		ret.Ipv6FibTblSyncEntriesSentM2nd = in["ipv6_fib_tbl_sync_entries_sent_m_2nd"].(int)
		ret.MacTblSyncEntriesSentM2nd = in["mac_tbl_sync_entries_sent_m_2nd"].(int)
		ret.ArpTblSyncEntriesRcvdB2nd = in["arp_tbl_sync_entries_rcvd_b_2nd"].(int)
		ret.Nd6TblSyncEntriesRcvdB2nd = in["nd6_tbl_sync_entries_rcvd_b_2nd"].(int)
		ret.Ipv4FibTblSyncEntriesRcvdB2nd = in["ipv4_fib_tbl_sync_entries_rcvd_b_2nd"].(int)
		ret.Ipv6FibTblSyncEntriesRcvdB2nd = in["ipv6_fib_tbl_sync_entries_rcvd_b_2nd"].(int)
		ret.MacTblSyncEntriesRcvdB2nd = in["mac_tbl_sync_entries_rcvd_b_2nd"].(int)
		ret.ArpTblSyncEntriesAddedB2nd = in["arp_tbl_sync_entries_added_b_2nd"].(int)
		ret.Nd6TblSyncEntriesAddedB2nd = in["nd6_tbl_sync_entries_added_b_2nd"].(int)
		ret.Ipv4FibTblSyncEntriesAddedB2nd = in["ipv4_fib_tbl_sync_entries_added_b_2nd"].(int)
		ret.Ipv6FibTblSyncEntriesAddedB2nd = in["ipv6_fib_tbl_sync_entries_added_b_2nd"].(int)
		ret.MacTblSyncEntriesAddedB2nd = in["mac_tbl_sync_entries_added_b_2nd"].(int)
		ret.ArpTblSyncEntriesRemovedB2nd = in["arp_tbl_sync_entries_removed_b_2nd"].(int)
		ret.Nd6TblSyncEntriesRemovedB2nd = in["nd6_tbl_sync_entries_removed_b_2nd"].(int)
		ret.Ipv4FibTblSyncEntriesRemovedB2nd = in["ipv4_fib_tbl_sync_entries_removed_b_2nd"].(int)
		ret.Ipv6FibTblSyncEntriesRemovedB2nd = in["ipv6_fib_tbl_sync_entries_removed_b_2nd"].(int)
		ret.MacTblSyncEntriesRemovedB2nd = in["mac_tbl_sync_entries_removed_b_2nd"].(int)
		ret.ArpTblSyncEndTsM2nd = in["arp_tbl_sync_end_ts_m_2nd"].(int)
		ret.Nd6TblSyncEndTsM2nd = in["nd6_tbl_sync_end_ts_m_2nd"].(int)
		ret.Ipv4FibTblSyncEndTsM2nd = in["ipv4_fib_tbl_sync_end_ts_m_2nd"].(int)
		ret.Ipv6FibTblSyncEndTsM2nd = in["ipv6_fib_tbl_sync_end_ts_m_2nd"].(int)
		ret.MacTblSyncEndTsM2nd = in["mac_tbl_sync_end_ts_m_2nd"].(int)
		ret.ArpTblSyncEndTsB2nd = in["arp_tbl_sync_end_ts_b_2nd"].(int)
		ret.Nd6TblSyncEndTsB2nd = in["nd6_tbl_sync_end_ts_b_2nd"].(int)
		ret.Ipv4FibTblSyncEndTsB2nd = in["ipv4_fib_tbl_sync_end_ts_b_2nd"].(int)
		ret.Ipv6FibTblSyncEndTsB2nd = in["ipv6_fib_tbl_sync_end_ts_b_2nd"].(int)
		ret.MacTblSyncEndTsB2nd = in["mac_tbl_sync_end_ts_b_2nd"].(int)
		ret.ArpTblSyncStartTsM3rd = in["arp_tbl_sync_start_ts_m_3rd"].(int)
		ret.Nd6TblSyncStartTsM3rd = in["nd6_tbl_sync_start_ts_m_3rd"].(int)
		ret.Ipv4FibTblSyncStartTsM3rd = in["ipv4_fib_tbl_sync_start_ts_m_3rd"].(int)
		ret.Ipv6FibTblSyncStartTsM3rd = in["ipv6_fib_tbl_sync_start_ts_m_3rd"].(int)
		ret.MacTblSyncStartTsM3rd = in["mac_tbl_sync_start_ts_m_3rd"].(int)
		ret.ArpTblSyncStartTsB3rd = in["arp_tbl_sync_start_ts_b_3rd"].(int)
		ret.Nd6TblSyncStartTsB3rd = in["nd6_tbl_sync_start_ts_b_3rd"].(int)
		ret.Ipv4FibTblSyncStartTsB3rd = in["ipv4_fib_tbl_sync_start_ts_b_3rd"].(int)
		ret.Ipv6FibTblSyncStartTsB3rd = in["ipv6_fib_tbl_sync_start_ts_b_3rd"].(int)
		ret.MacTblSyncStartTsB3rd = in["mac_tbl_sync_start_ts_b_3rd"].(int)
		ret.ArpTblSyncEntriesSentM3rd = in["arp_tbl_sync_entries_sent_m_3rd"].(int)
		ret.Nd6TblSyncEntriesSentM3rd = in["nd6_tbl_sync_entries_sent_m_3rd"].(int)
		ret.Ipv4FibTblSyncEntriesSentM3rd = in["ipv4_fib_tbl_sync_entries_sent_m_3rd"].(int)
		ret.Ipv6FibTblSyncEntriesSentM3rd = in["ipv6_fib_tbl_sync_entries_sent_m_3rd"].(int)
		ret.MacTblSyncEntriesSentM3rd = in["mac_tbl_sync_entries_sent_m_3rd"].(int)
		ret.ArpTblSyncEntriesRcvdB3rd = in["arp_tbl_sync_entries_rcvd_b_3rd"].(int)
		ret.Nd6TblSyncEntriesRcvdB3rd = in["nd6_tbl_sync_entries_rcvd_b_3rd"].(int)
		ret.Ipv4FibTblSyncEntriesRcvdB3rd = in["ipv4_fib_tbl_sync_entries_rcvd_b_3rd"].(int)
		ret.Ipv6FibTblSyncEntriesRcvdB3rd = in["ipv6_fib_tbl_sync_entries_rcvd_b_3rd"].(int)
		ret.MacTblSyncEntriesRcvdB3rd = in["mac_tbl_sync_entries_rcvd_b_3rd"].(int)
		ret.ArpTblSyncEntriesAddedB3rd = in["arp_tbl_sync_entries_added_b_3rd"].(int)
		ret.Nd6TblSyncEntriesAddedB3rd = in["nd6_tbl_sync_entries_added_b_3rd"].(int)
		ret.Ipv4FibTblSyncEntriesAddedB3rd = in["ipv4_fib_tbl_sync_entries_added_b_3rd"].(int)
		ret.Ipv6FibTblSyncEntriesAddedB3rd = in["ipv6_fib_tbl_sync_entries_added_b_3rd"].(int)
		ret.MacTblSyncEntriesAddedB3rd = in["mac_tbl_sync_entries_added_b_3rd"].(int)
		ret.ArpTblSyncEntriesRemovedB3rd = in["arp_tbl_sync_entries_removed_b_3rd"].(int)
		ret.Nd6TblSyncEntriesRemovedB3rd = in["nd6_tbl_sync_entries_removed_b_3rd"].(int)
		ret.Ipv4FibTblSyncEntriesRemovedB3rd = in["ipv4_fib_tbl_sync_entries_removed_b_3rd"].(int)
		ret.Ipv6FibTblSyncEntriesRemovedB3rd = in["ipv6_fib_tbl_sync_entries_removed_b_3rd"].(int)
		ret.MacTblSyncEntriesRemovedB3rd = in["mac_tbl_sync_entries_removed_b_3rd"].(int)
		ret.ArpTblSyncEndTsM3rd = in["arp_tbl_sync_end_ts_m_3rd"].(int)
		ret.Nd6TblSyncEndTsM3rd = in["nd6_tbl_sync_end_ts_m_3rd"].(int)
		ret.Ipv4FibTblSyncEndTsM3rd = in["ipv4_fib_tbl_sync_end_ts_m_3rd"].(int)
		ret.Ipv6FibTblSyncEndTsM3rd = in["ipv6_fib_tbl_sync_end_ts_m_3rd"].(int)
		ret.MacTblSyncEndTsM3rd = in["mac_tbl_sync_end_ts_m_3rd"].(int)
		ret.ArpTblSyncEndTsB3rd = in["arp_tbl_sync_end_ts_b_3rd"].(int)
		ret.Nd6TblSyncEndTsB3rd = in["nd6_tbl_sync_end_ts_b_3rd"].(int)
		ret.Ipv4FibTblSyncEndTsB3rd = in["ipv4_fib_tbl_sync_end_ts_b_3rd"].(int)
		ret.Ipv6FibTblSyncEndTsB3rd = in["ipv6_fib_tbl_sync_end_ts_b_3rd"].(int)
		ret.MacTblSyncEndTsB3rd = in["mac_tbl_sync_end_ts_b_3rd"].(int)
		ret.ArpTblSyncStartTsM4th = in["arp_tbl_sync_start_ts_m_4th"].(int)
		ret.Nd6TblSyncStartTsM4th = in["nd6_tbl_sync_start_ts_m_4th"].(int)
		ret.Ipv4FibTblSyncStartTsM4th = in["ipv4_fib_tbl_sync_start_ts_m_4th"].(int)
		ret.Ipv6FibTblSyncStartTsM4th = in["ipv6_fib_tbl_sync_start_ts_m_4th"].(int)
		ret.MacTblSyncStartTsM4th = in["mac_tbl_sync_start_ts_m_4th"].(int)
		ret.ArpTblSyncStartTsB4th = in["arp_tbl_sync_start_ts_b_4th"].(int)
		ret.Nd6TblSyncStartTsB4th = in["nd6_tbl_sync_start_ts_b_4th"].(int)
		ret.Ipv4FibTblSyncStartTsB4th = in["ipv4_fib_tbl_sync_start_ts_b_4th"].(int)
		ret.Ipv6FibTblSyncStartTsB4th = in["ipv6_fib_tbl_sync_start_ts_b_4th"].(int)
		ret.MacTblSyncStartTsB4th = in["mac_tbl_sync_start_ts_b_4th"].(int)
		ret.ArpTblSyncEntriesSentM4th = in["arp_tbl_sync_entries_sent_m_4th"].(int)
		ret.Nd6TblSyncEntriesSentM4th = in["nd6_tbl_sync_entries_sent_m_4th"].(int)
		ret.Ipv4FibTblSyncEntriesSentM4th = in["ipv4_fib_tbl_sync_entries_sent_m_4th"].(int)
		ret.Ipv6FibTblSyncEntriesSentM4th = in["ipv6_fib_tbl_sync_entries_sent_m_4th"].(int)
		ret.MacTblSyncEntriesSentM4th = in["mac_tbl_sync_entries_sent_m_4th"].(int)
		ret.ArpTblSyncEntriesRcvdB4th = in["arp_tbl_sync_entries_rcvd_b_4th"].(int)
		ret.Nd6TblSyncEntriesRcvdB4th = in["nd6_tbl_sync_entries_rcvd_b_4th"].(int)
		ret.Ipv4FibTblSyncEntriesRcvdB4th = in["ipv4_fib_tbl_sync_entries_rcvd_b_4th"].(int)
		ret.Ipv6FibTblSyncEntriesRcvdB4th = in["ipv6_fib_tbl_sync_entries_rcvd_b_4th"].(int)
		ret.MacTblSyncEntriesRcvdB4th = in["mac_tbl_sync_entries_rcvd_b_4th"].(int)
		ret.ArpTblSyncEntriesAddedB4th = in["arp_tbl_sync_entries_added_b_4th"].(int)
		ret.Nd6TblSyncEntriesAddedB4th = in["nd6_tbl_sync_entries_added_b_4th"].(int)
		ret.Ipv4FibTblSyncEntriesAddedB4th = in["ipv4_fib_tbl_sync_entries_added_b_4th"].(int)
		ret.Ipv6FibTblSyncEntriesAddedB4th = in["ipv6_fib_tbl_sync_entries_added_b_4th"].(int)
		ret.MacTblSyncEntriesAddedB4th = in["mac_tbl_sync_entries_added_b_4th"].(int)
		ret.ArpTblSyncEntriesRemovedB4th = in["arp_tbl_sync_entries_removed_b_4th"].(int)
		ret.Nd6TblSyncEntriesRemovedB4th = in["nd6_tbl_sync_entries_removed_b_4th"].(int)
		ret.Ipv4FibTblSyncEntriesRemovedB4th = in["ipv4_fib_tbl_sync_entries_removed_b_4th"].(int)
		ret.Ipv6FibTblSyncEntriesRemovedB4th = in["ipv6_fib_tbl_sync_entries_removed_b_4th"].(int)
		ret.MacTblSyncEntriesRemovedB4th = in["mac_tbl_sync_entries_removed_b_4th"].(int)
		ret.ArpTblSyncEndTsM4th = in["arp_tbl_sync_end_ts_m_4th"].(int)
		ret.Nd6TblSyncEndTsM4th = in["nd6_tbl_sync_end_ts_m_4th"].(int)
		ret.Ipv4FibTblSyncEndTsM4th = in["ipv4_fib_tbl_sync_end_ts_m_4th"].(int)
		ret.Ipv6FibTblSyncEndTsM4th = in["ipv6_fib_tbl_sync_end_ts_m_4th"].(int)
		ret.MacTblSyncEndTsM4th = in["mac_tbl_sync_end_ts_m_4th"].(int)
		ret.ArpTblSyncEndTsB4th = in["arp_tbl_sync_end_ts_b_4th"].(int)
		ret.Nd6TblSyncEndTsB4th = in["nd6_tbl_sync_end_ts_b_4th"].(int)
		ret.Ipv4FibTblSyncEndTsB4th = in["ipv4_fib_tbl_sync_end_ts_b_4th"].(int)
		ret.Ipv6FibTblSyncEndTsB4th = in["ipv6_fib_tbl_sync_end_ts_b_4th"].(int)
		ret.MacTblSyncEndTsB4th = in["mac_tbl_sync_end_ts_b_4th"].(int)
		ret.ArpTblSyncStartTsM5th = in["arp_tbl_sync_start_ts_m_5th"].(int)
		ret.Nd6TblSyncStartTsM5th = in["nd6_tbl_sync_start_ts_m_5th"].(int)
		ret.Ipv4FibTblSyncStartTsM5th = in["ipv4_fib_tbl_sync_start_ts_m_5th"].(int)
		ret.Ipv6FibTblSyncStartTsM5th = in["ipv6_fib_tbl_sync_start_ts_m_5th"].(int)
		ret.MacTblSyncStartTsM5th = in["mac_tbl_sync_start_ts_m_5th"].(int)
		ret.ArpTblSyncStartTsB5th = in["arp_tbl_sync_start_ts_b_5th"].(int)
		ret.Nd6TblSyncStartTsB5th = in["nd6_tbl_sync_start_ts_b_5th"].(int)
		ret.Ipv4FibTblSyncStartTsB5th = in["ipv4_fib_tbl_sync_start_ts_b_5th"].(int)
		ret.Ipv6FibTblSyncStartTsB5th = in["ipv6_fib_tbl_sync_start_ts_b_5th"].(int)
		ret.MacTblSyncStartTsB5th = in["mac_tbl_sync_start_ts_b_5th"].(int)
		ret.ArpTblSyncEntriesSentM5th = in["arp_tbl_sync_entries_sent_m_5th"].(int)
		ret.Nd6TblSyncEntriesSentM5th = in["nd6_tbl_sync_entries_sent_m_5th"].(int)
		ret.Ipv4FibTblSyncEntriesSentM5th = in["ipv4_fib_tbl_sync_entries_sent_m_5th"].(int)
		ret.Ipv6FibTblSyncEntriesSentM5th = in["ipv6_fib_tbl_sync_entries_sent_m_5th"].(int)
		ret.MacTblSyncEntriesSentM5th = in["mac_tbl_sync_entries_sent_m_5th"].(int)
		ret.ArpTblSyncEntriesRcvdB5th = in["arp_tbl_sync_entries_rcvd_b_5th"].(int)
		ret.Nd6TblSyncEntriesRcvdB5th = in["nd6_tbl_sync_entries_rcvd_b_5th"].(int)
		ret.Ipv4FibTblSyncEntriesRcvdB5th = in["ipv4_fib_tbl_sync_entries_rcvd_b_5th"].(int)
		ret.Ipv6FibTblSyncEntriesRcvdB5th = in["ipv6_fib_tbl_sync_entries_rcvd_b_5th"].(int)
		ret.MacTblSyncEntriesRcvdB5th = in["mac_tbl_sync_entries_rcvd_b_5th"].(int)
		ret.ArpTblSyncEntriesAddedB5th = in["arp_tbl_sync_entries_added_b_5th"].(int)
		ret.Nd6TblSyncEntriesAddedB5th = in["nd6_tbl_sync_entries_added_b_5th"].(int)
		ret.Ipv4FibTblSyncEntriesAddedB5th = in["ipv4_fib_tbl_sync_entries_added_b_5th"].(int)
		ret.Ipv6FibTblSyncEntriesAddedB5th = in["ipv6_fib_tbl_sync_entries_added_b_5th"].(int)
		ret.MacTblSyncEntriesAddedB5th = in["mac_tbl_sync_entries_added_b_5th"].(int)
		ret.ArpTblSyncEntriesRemovedB5th = in["arp_tbl_sync_entries_removed_b_5th"].(int)
		ret.Nd6TblSyncEntriesRemovedB5th = in["nd6_tbl_sync_entries_removed_b_5th"].(int)
		ret.Ipv4FibTblSyncEntriesRemovedB5th = in["ipv4_fib_tbl_sync_entries_removed_b_5th"].(int)
		ret.Ipv6FibTblSyncEntriesRemovedB5th = in["ipv6_fib_tbl_sync_entries_removed_b_5th"].(int)
		ret.MacTblSyncEntriesRemovedB5th = in["mac_tbl_sync_entries_removed_b_5th"].(int)
		ret.ArpTblSyncEndTsM5th = in["arp_tbl_sync_end_ts_m_5th"].(int)
		ret.Nd6TblSyncEndTsM5th = in["nd6_tbl_sync_end_ts_m_5th"].(int)
		ret.Ipv4FibTblSyncEndTsM5th = in["ipv4_fib_tbl_sync_end_ts_m_5th"].(int)
		ret.Ipv6FibTblSyncEndTsM5th = in["ipv6_fib_tbl_sync_end_ts_m_5th"].(int)
		ret.MacTblSyncEndTsM5th = in["mac_tbl_sync_end_ts_m_5th"].(int)
		ret.ArpTblSyncEndTsB5th = in["arp_tbl_sync_end_ts_b_5th"].(int)
		ret.Nd6TblSyncEndTsB5th = in["nd6_tbl_sync_end_ts_b_5th"].(int)
		ret.Ipv4FibTblSyncEndTsB5th = in["ipv4_fib_tbl_sync_end_ts_b_5th"].(int)
		ret.Ipv6FibTblSyncEndTsB5th = in["ipv6_fib_tbl_sync_end_ts_b_5th"].(int)
		ret.MacTblSyncEndTsB5th = in["mac_tbl_sync_end_ts_b_5th"].(int)
		ret.ArpTblSyncM = in["arp_tbl_sync_m"].(int)
		ret.Nd6TblSyncM = in["nd6_tbl_sync_m"].(int)
		ret.Ipv4FibTblSyncM = in["ipv4_fib_tbl_sync_m"].(int)
		ret.Ipv6FibTblSyncM = in["ipv6_fib_tbl_sync_m"].(int)
		ret.MacTblSyncM = in["mac_tbl_sync_m"].(int)
		ret.ArpTblSyncB = in["arp_tbl_sync_b"].(int)
		ret.Nd6TblSyncB = in["nd6_tbl_sync_b"].(int)
		ret.Ipv4FibTblSyncB = in["ipv4_fib_tbl_sync_b"].(int)
		ret.Ipv6FibTblSyncB = in["ipv6_fib_tbl_sync_b"].(int)
		ret.MacTblSyncB = in["mac_tbl_sync_b"].(int)
		ret.ArpTblCksumM = in["arp_tbl_cksum_m"].(int)
		ret.Nd6TblCksumM = in["nd6_tbl_cksum_m"].(int)
		ret.Ipv4FibTblCksumM = in["ipv4_fib_tbl_cksum_m"].(int)
		ret.Ipv6FibTblCksumM = in["ipv6_fib_tbl_cksum_m"].(int)
		ret.MacTblCksumM = in["mac_tbl_cksum_m"].(int)
		ret.ArpTblCksumB = in["arp_tbl_cksum_b"].(int)
		ret.Nd6TblCksumB = in["nd6_tbl_cksum_b"].(int)
		ret.Ipv4FibTblCksumB = in["ipv4_fib_tbl_cksum_b"].(int)
		ret.Ipv6FibTblCksumB = in["ipv6_fib_tbl_cksum_b"].(int)
		ret.MacTblCksumB = in["mac_tbl_cksum_b"].(int)
		ret.ArpTblCksumMismatchB = in["arp_tbl_cksum_mismatch_b"].(int)
		ret.Nd6TblCksumMismatchB = in["nd6_tbl_cksum_mismatch_b"].(int)
		ret.Ipv4FibTblCksumMismatchB = in["ipv4_fib_tbl_cksum_mismatch_b"].(int)
		ret.Ipv6FibTblCksumMismatchB = in["ipv6_fib_tbl_cksum_mismatch_b"].(int)
		ret.MacTblCksumMismatchB = in["mac_tbl_cksum_mismatch_b"].(int)
		ret.ArpTblCksumCancelM = in["arp_tbl_cksum_cancel_m"].(int)
		ret.Nd6TblCksumCancelM = in["nd6_tbl_cksum_cancel_m"].(int)
		ret.Ipv4FibTblCksumCancelM = in["ipv4_fib_tbl_cksum_cancel_m"].(int)
		ret.Ipv6FibTblCksumCancelM = in["ipv6_fib_tbl_cksum_cancel_m"].(int)
		ret.MacTblCksumCancelM = in["mac_tbl_cksum_cancel_m"].(int)
	}
	return ret
}

func dataToEndpointSystemTableIntegrityStats(d *schema.ResourceData) edpt.SystemTableIntegrityStats {
	var ret edpt.SystemTableIntegrityStats

	ret.Stats = getObjectSystemTableIntegrityStatsStats(d.Get("stats").([]interface{}))
	return ret
}
