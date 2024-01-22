

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTableIntegrityStats struct {
    
    Stats SystemTableIntegrityStatsStats `json:"stats"`

}
type DataSystemTableIntegrityStats struct {
    DtSystemTableIntegrityStats SystemTableIntegrityStats `json:"table-integrity"`
}


type SystemTableIntegrityStatsStats struct {
    ArpTblSyncStartTsM1st int `json:"arp-tbl-sync-start-ts-m-1st"`
    Nd6TblSyncStartTsM1st int `json:"nd6-tbl-sync-start-ts-m-1st"`
    Ipv4FibTblSyncStartTsM1st int `json:"ipv4-fib-tbl-sync-start-ts-m-1st"`
    Ipv6FibTblSyncStartTsM1st int `json:"ipv6-fib-tbl-sync-start-ts-m-1st"`
    MacTblSyncStartTsM1st int `json:"mac-tbl-sync-start-ts-m-1st"`
    ArpTblSyncStartTsB1st int `json:"arp-tbl-sync-start-ts-b-1st"`
    Nd6TblSyncStartTsB1st int `json:"nd6-tbl-sync-start-ts-b-1st"`
    Ipv4FibTblSyncStartTsB1st int `json:"ipv4-fib-tbl-sync-start-ts-b-1st"`
    Ipv6FibTblSyncStartTsB1st int `json:"ipv6-fib-tbl-sync-start-ts-b-1st"`
    MacTblSyncStartTsB1st int `json:"mac-tbl-sync-start-ts-b-1st"`
    ArpTblSyncEntriesSentM1st int `json:"arp-tbl-sync-entries-sent-m-1st"`
    Nd6TblSyncEntriesSentM1st int `json:"nd6-tbl-sync-entries-sent-m-1st"`
    Ipv4FibTblSyncEntriesSentM1st int `json:"ipv4-fib-tbl-sync-entries-sent-m-1st"`
    Ipv6FibTblSyncEntriesSentM1st int `json:"ipv6-fib-tbl-sync-entries-sent-m-1st"`
    MacTblSyncEntriesSentM1st int `json:"mac-tbl-sync-entries-sent-m-1st"`
    ArpTblSyncEntriesRcvdB1st int `json:"arp-tbl-sync-entries-rcvd-b-1st"`
    Nd6TblSyncEntriesRcvdB1st int `json:"nd6-tbl-sync-entries-rcvd-b-1st"`
    Ipv4FibTblSyncEntriesRcvdB1st int `json:"ipv4-fib-tbl-sync-entries-rcvd-b-1st"`
    Ipv6FibTblSyncEntriesRcvdB1st int `json:"ipv6-fib-tbl-sync-entries-rcvd-b-1st"`
    MacTblSyncEntriesRcvdB1st int `json:"mac-tbl-sync-entries-rcvd-b-1st"`
    ArpTblSyncEntriesAddedB1st int `json:"arp-tbl-sync-entries-added-b-1st"`
    Nd6TblSyncEntriesAddedB1st int `json:"nd6-tbl-sync-entries-added-b-1st"`
    Ipv4FibTblSyncEntriesAddedB1st int `json:"ipv4-fib-tbl-sync-entries-added-b-1st"`
    Ipv6FibTblSyncEntriesAddedB1st int `json:"ipv6-fib-tbl-sync-entries-added-b-1st"`
    MacTblSyncEntriesAddedB1st int `json:"mac-tbl-sync-entries-added-b-1st"`
    ArpTblSyncEntriesRemovedB1st int `json:"arp-tbl-sync-entries-removed-b-1st"`
    Nd6TblSyncEntriesRemovedB1st int `json:"nd6-tbl-sync-entries-removed-b-1st"`
    Ipv4FibTblSyncEntriesRemovedB1st int `json:"ipv4-fib-tbl-sync-entries-removed-b-1st"`
    Ipv6FibTblSyncEntriesRemovedB1st int `json:"ipv6-fib-tbl-sync-entries-removed-b-1st"`
    MacTblSyncEntriesRemovedB1st int `json:"mac-tbl-sync-entries-removed-b-1st"`
    ArpTblSyncEndTsM1st int `json:"arp-tbl-sync-end-ts-m-1st"`
    Nd6TblSyncEndTsM1st int `json:"nd6-tbl-sync-end-ts-m-1st"`
    Ipv4FibTblSyncEndTsM1st int `json:"ipv4-fib-tbl-sync-end-ts-m-1st"`
    Ipv6FibTblSyncEndTsM1st int `json:"ipv6-fib-tbl-sync-end-ts-m-1st"`
    MacTblSyncEndTsM1st int `json:"mac-tbl-sync-end-ts-m-1st"`
    ArpTblSyncEndTsB1st int `json:"arp-tbl-sync-end-ts-b-1st"`
    Nd6TblSyncEndTsB1st int `json:"nd6-tbl-sync-end-ts-b-1st"`
    Ipv4FibTblSyncEndTsB1st int `json:"ipv4-fib-tbl-sync-end-ts-b-1st"`
    Ipv6FibTblSyncEndTsB1st int `json:"ipv6-fib-tbl-sync-end-ts-b-1st"`
    MacTblSyncEndTsB1st int `json:"mac-tbl-sync-end-ts-b-1st"`
    ArpTblSyncStartTsM2nd int `json:"arp-tbl-sync-start-ts-m-2nd"`
    Nd6TblSyncStartTsM2nd int `json:"nd6-tbl-sync-start-ts-m-2nd"`
    Ipv4FibTblSyncStartTsM2nd int `json:"ipv4-fib-tbl-sync-start-ts-m-2nd"`
    Ipv6FibTblSyncStartTsM2nd int `json:"ipv6-fib-tbl-sync-start-ts-m-2nd"`
    MacTblSyncStartTsM2nd int `json:"mac-tbl-sync-start-ts-m-2nd"`
    ArpTblSyncStartTsB2nd int `json:"arp-tbl-sync-start-ts-b-2nd"`
    Nd6TblSyncStartTsB2nd int `json:"nd6-tbl-sync-start-ts-b-2nd"`
    Ipv4FibTblSyncStartTsB2nd int `json:"ipv4-fib-tbl-sync-start-ts-b-2nd"`
    Ipv6FibTblSyncStartTsB2nd int `json:"ipv6-fib-tbl-sync-start-ts-b-2nd"`
    MacTblSyncStartTsB2nd int `json:"mac-tbl-sync-start-ts-b-2nd"`
    ArpTblSyncEntriesSentM2nd int `json:"arp-tbl-sync-entries-sent-m-2nd"`
    Nd6TblSyncEntriesSentM2nd int `json:"nd6-tbl-sync-entries-sent-m-2nd"`
    Ipv4FibTblSyncEntriesSentM2nd int `json:"ipv4-fib-tbl-sync-entries-sent-m-2nd"`
    Ipv6FibTblSyncEntriesSentM2nd int `json:"ipv6-fib-tbl-sync-entries-sent-m-2nd"`
    MacTblSyncEntriesSentM2nd int `json:"mac-tbl-sync-entries-sent-m-2nd"`
    ArpTblSyncEntriesRcvdB2nd int `json:"arp-tbl-sync-entries-rcvd-b-2nd"`
    Nd6TblSyncEntriesRcvdB2nd int `json:"nd6-tbl-sync-entries-rcvd-b-2nd"`
    Ipv4FibTblSyncEntriesRcvdB2nd int `json:"ipv4-fib-tbl-sync-entries-rcvd-b-2nd"`
    Ipv6FibTblSyncEntriesRcvdB2nd int `json:"ipv6-fib-tbl-sync-entries-rcvd-b-2nd"`
    MacTblSyncEntriesRcvdB2nd int `json:"mac-tbl-sync-entries-rcvd-b-2nd"`
    ArpTblSyncEntriesAddedB2nd int `json:"arp-tbl-sync-entries-added-b-2nd"`
    Nd6TblSyncEntriesAddedB2nd int `json:"nd6-tbl-sync-entries-added-b-2nd"`
    Ipv4FibTblSyncEntriesAddedB2nd int `json:"ipv4-fib-tbl-sync-entries-added-b-2nd"`
    Ipv6FibTblSyncEntriesAddedB2nd int `json:"ipv6-fib-tbl-sync-entries-added-b-2nd"`
    MacTblSyncEntriesAddedB2nd int `json:"mac-tbl-sync-entries-added-b-2nd"`
    ArpTblSyncEntriesRemovedB2nd int `json:"arp-tbl-sync-entries-removed-b-2nd"`
    Nd6TblSyncEntriesRemovedB2nd int `json:"nd6-tbl-sync-entries-removed-b-2nd"`
    Ipv4FibTblSyncEntriesRemovedB2nd int `json:"ipv4-fib-tbl-sync-entries-removed-b-2nd"`
    Ipv6FibTblSyncEntriesRemovedB2nd int `json:"ipv6-fib-tbl-sync-entries-removed-b-2nd"`
    MacTblSyncEntriesRemovedB2nd int `json:"mac-tbl-sync-entries-removed-b-2nd"`
    ArpTblSyncEndTsM2nd int `json:"arp-tbl-sync-end-ts-m-2nd"`
    Nd6TblSyncEndTsM2nd int `json:"nd6-tbl-sync-end-ts-m-2nd"`
    Ipv4FibTblSyncEndTsM2nd int `json:"ipv4-fib-tbl-sync-end-ts-m-2nd"`
    Ipv6FibTblSyncEndTsM2nd int `json:"ipv6-fib-tbl-sync-end-ts-m-2nd"`
    MacTblSyncEndTsM2nd int `json:"mac-tbl-sync-end-ts-m-2nd"`
    ArpTblSyncEndTsB2nd int `json:"arp-tbl-sync-end-ts-b-2nd"`
    Nd6TblSyncEndTsB2nd int `json:"nd6-tbl-sync-end-ts-b-2nd"`
    Ipv4FibTblSyncEndTsB2nd int `json:"ipv4-fib-tbl-sync-end-ts-b-2nd"`
    Ipv6FibTblSyncEndTsB2nd int `json:"ipv6-fib-tbl-sync-end-ts-b-2nd"`
    MacTblSyncEndTsB2nd int `json:"mac-tbl-sync-end-ts-b-2nd"`
    ArpTblSyncStartTsM3rd int `json:"arp-tbl-sync-start-ts-m-3rd"`
    Nd6TblSyncStartTsM3rd int `json:"nd6-tbl-sync-start-ts-m-3rd"`
    Ipv4FibTblSyncStartTsM3rd int `json:"ipv4-fib-tbl-sync-start-ts-m-3rd"`
    Ipv6FibTblSyncStartTsM3rd int `json:"ipv6-fib-tbl-sync-start-ts-m-3rd"`
    MacTblSyncStartTsM3rd int `json:"mac-tbl-sync-start-ts-m-3rd"`
    ArpTblSyncStartTsB3rd int `json:"arp-tbl-sync-start-ts-b-3rd"`
    Nd6TblSyncStartTsB3rd int `json:"nd6-tbl-sync-start-ts-b-3rd"`
    Ipv4FibTblSyncStartTsB3rd int `json:"ipv4-fib-tbl-sync-start-ts-b-3rd"`
    Ipv6FibTblSyncStartTsB3rd int `json:"ipv6-fib-tbl-sync-start-ts-b-3rd"`
    MacTblSyncStartTsB3rd int `json:"mac-tbl-sync-start-ts-b-3rd"`
    ArpTblSyncEntriesSentM3rd int `json:"arp-tbl-sync-entries-sent-m-3rd"`
    Nd6TblSyncEntriesSentM3rd int `json:"nd6-tbl-sync-entries-sent-m-3rd"`
    Ipv4FibTblSyncEntriesSentM3rd int `json:"ipv4-fib-tbl-sync-entries-sent-m-3rd"`
    Ipv6FibTblSyncEntriesSentM3rd int `json:"ipv6-fib-tbl-sync-entries-sent-m-3rd"`
    MacTblSyncEntriesSentM3rd int `json:"mac-tbl-sync-entries-sent-m-3rd"`
    ArpTblSyncEntriesRcvdB3rd int `json:"arp-tbl-sync-entries-rcvd-b-3rd"`
    Nd6TblSyncEntriesRcvdB3rd int `json:"nd6-tbl-sync-entries-rcvd-b-3rd"`
    Ipv4FibTblSyncEntriesRcvdB3rd int `json:"ipv4-fib-tbl-sync-entries-rcvd-b-3rd"`
    Ipv6FibTblSyncEntriesRcvdB3rd int `json:"ipv6-fib-tbl-sync-entries-rcvd-b-3rd"`
    MacTblSyncEntriesRcvdB3rd int `json:"mac-tbl-sync-entries-rcvd-b-3rd"`
    ArpTblSyncEntriesAddedB3rd int `json:"arp-tbl-sync-entries-added-b-3rd"`
    Nd6TblSyncEntriesAddedB3rd int `json:"nd6-tbl-sync-entries-added-b-3rd"`
    Ipv4FibTblSyncEntriesAddedB3rd int `json:"ipv4-fib-tbl-sync-entries-added-b-3rd"`
    Ipv6FibTblSyncEntriesAddedB3rd int `json:"ipv6-fib-tbl-sync-entries-added-b-3rd"`
    MacTblSyncEntriesAddedB3rd int `json:"mac-tbl-sync-entries-added-b-3rd"`
    ArpTblSyncEntriesRemovedB3rd int `json:"arp-tbl-sync-entries-removed-b-3rd"`
    Nd6TblSyncEntriesRemovedB3rd int `json:"nd6-tbl-sync-entries-removed-b-3rd"`
    Ipv4FibTblSyncEntriesRemovedB3rd int `json:"ipv4-fib-tbl-sync-entries-removed-b-3rd"`
    Ipv6FibTblSyncEntriesRemovedB3rd int `json:"ipv6-fib-tbl-sync-entries-removed-b-3rd"`
    MacTblSyncEntriesRemovedB3rd int `json:"mac-tbl-sync-entries-removed-b-3rd"`
    ArpTblSyncEndTsM3rd int `json:"arp-tbl-sync-end-ts-m-3rd"`
    Nd6TblSyncEndTsM3rd int `json:"nd6-tbl-sync-end-ts-m-3rd"`
    Ipv4FibTblSyncEndTsM3rd int `json:"ipv4-fib-tbl-sync-end-ts-m-3rd"`
    Ipv6FibTblSyncEndTsM3rd int `json:"ipv6-fib-tbl-sync-end-ts-m-3rd"`
    MacTblSyncEndTsM3rd int `json:"mac-tbl-sync-end-ts-m-3rd"`
    ArpTblSyncEndTsB3rd int `json:"arp-tbl-sync-end-ts-b-3rd"`
    Nd6TblSyncEndTsB3rd int `json:"nd6-tbl-sync-end-ts-b-3rd"`
    Ipv4FibTblSyncEndTsB3rd int `json:"ipv4-fib-tbl-sync-end-ts-b-3rd"`
    Ipv6FibTblSyncEndTsB3rd int `json:"ipv6-fib-tbl-sync-end-ts-b-3rd"`
    MacTblSyncEndTsB3rd int `json:"mac-tbl-sync-end-ts-b-3rd"`
    ArpTblSyncStartTsM4th int `json:"arp-tbl-sync-start-ts-m-4th"`
    Nd6TblSyncStartTsM4th int `json:"nd6-tbl-sync-start-ts-m-4th"`
    Ipv4FibTblSyncStartTsM4th int `json:"ipv4-fib-tbl-sync-start-ts-m-4th"`
    Ipv6FibTblSyncStartTsM4th int `json:"ipv6-fib-tbl-sync-start-ts-m-4th"`
    MacTblSyncStartTsM4th int `json:"mac-tbl-sync-start-ts-m-4th"`
    ArpTblSyncStartTsB4th int `json:"arp-tbl-sync-start-ts-b-4th"`
    Nd6TblSyncStartTsB4th int `json:"nd6-tbl-sync-start-ts-b-4th"`
    Ipv4FibTblSyncStartTsB4th int `json:"ipv4-fib-tbl-sync-start-ts-b-4th"`
    Ipv6FibTblSyncStartTsB4th int `json:"ipv6-fib-tbl-sync-start-ts-b-4th"`
    MacTblSyncStartTsB4th int `json:"mac-tbl-sync-start-ts-b-4th"`
    ArpTblSyncEntriesSentM4th int `json:"arp-tbl-sync-entries-sent-m-4th"`
    Nd6TblSyncEntriesSentM4th int `json:"nd6-tbl-sync-entries-sent-m-4th"`
    Ipv4FibTblSyncEntriesSentM4th int `json:"ipv4-fib-tbl-sync-entries-sent-m-4th"`
    Ipv6FibTblSyncEntriesSentM4th int `json:"ipv6-fib-tbl-sync-entries-sent-m-4th"`
    MacTblSyncEntriesSentM4th int `json:"mac-tbl-sync-entries-sent-m-4th"`
    ArpTblSyncEntriesRcvdB4th int `json:"arp-tbl-sync-entries-rcvd-b-4th"`
    Nd6TblSyncEntriesRcvdB4th int `json:"nd6-tbl-sync-entries-rcvd-b-4th"`
    Ipv4FibTblSyncEntriesRcvdB4th int `json:"ipv4-fib-tbl-sync-entries-rcvd-b-4th"`
    Ipv6FibTblSyncEntriesRcvdB4th int `json:"ipv6-fib-tbl-sync-entries-rcvd-b-4th"`
    MacTblSyncEntriesRcvdB4th int `json:"mac-tbl-sync-entries-rcvd-b-4th"`
    ArpTblSyncEntriesAddedB4th int `json:"arp-tbl-sync-entries-added-b-4th"`
    Nd6TblSyncEntriesAddedB4th int `json:"nd6-tbl-sync-entries-added-b-4th"`
    Ipv4FibTblSyncEntriesAddedB4th int `json:"ipv4-fib-tbl-sync-entries-added-b-4th"`
    Ipv6FibTblSyncEntriesAddedB4th int `json:"ipv6-fib-tbl-sync-entries-added-b-4th"`
    MacTblSyncEntriesAddedB4th int `json:"mac-tbl-sync-entries-added-b-4th"`
    ArpTblSyncEntriesRemovedB4th int `json:"arp-tbl-sync-entries-removed-b-4th"`
    Nd6TblSyncEntriesRemovedB4th int `json:"nd6-tbl-sync-entries-removed-b-4th"`
    Ipv4FibTblSyncEntriesRemovedB4th int `json:"ipv4-fib-tbl-sync-entries-removed-b-4th"`
    Ipv6FibTblSyncEntriesRemovedB4th int `json:"ipv6-fib-tbl-sync-entries-removed-b-4th"`
    MacTblSyncEntriesRemovedB4th int `json:"mac-tbl-sync-entries-removed-b-4th"`
    ArpTblSyncEndTsM4th int `json:"arp-tbl-sync-end-ts-m-4th"`
    Nd6TblSyncEndTsM4th int `json:"nd6-tbl-sync-end-ts-m-4th"`
    Ipv4FibTblSyncEndTsM4th int `json:"ipv4-fib-tbl-sync-end-ts-m-4th"`
    Ipv6FibTblSyncEndTsM4th int `json:"ipv6-fib-tbl-sync-end-ts-m-4th"`
    MacTblSyncEndTsM4th int `json:"mac-tbl-sync-end-ts-m-4th"`
    ArpTblSyncEndTsB4th int `json:"arp-tbl-sync-end-ts-b-4th"`
    Nd6TblSyncEndTsB4th int `json:"nd6-tbl-sync-end-ts-b-4th"`
    Ipv4FibTblSyncEndTsB4th int `json:"ipv4-fib-tbl-sync-end-ts-b-4th"`
    Ipv6FibTblSyncEndTsB4th int `json:"ipv6-fib-tbl-sync-end-ts-b-4th"`
    MacTblSyncEndTsB4th int `json:"mac-tbl-sync-end-ts-b-4th"`
    ArpTblSyncStartTsM5th int `json:"arp-tbl-sync-start-ts-m-5th"`
    Nd6TblSyncStartTsM5th int `json:"nd6-tbl-sync-start-ts-m-5th"`
    Ipv4FibTblSyncStartTsM5th int `json:"ipv4-fib-tbl-sync-start-ts-m-5th"`
    Ipv6FibTblSyncStartTsM5th int `json:"ipv6-fib-tbl-sync-start-ts-m-5th"`
    MacTblSyncStartTsM5th int `json:"mac-tbl-sync-start-ts-m-5th"`
    ArpTblSyncStartTsB5th int `json:"arp-tbl-sync-start-ts-b-5th"`
    Nd6TblSyncStartTsB5th int `json:"nd6-tbl-sync-start-ts-b-5th"`
    Ipv4FibTblSyncStartTsB5th int `json:"ipv4-fib-tbl-sync-start-ts-b-5th"`
    Ipv6FibTblSyncStartTsB5th int `json:"ipv6-fib-tbl-sync-start-ts-b-5th"`
    MacTblSyncStartTsB5th int `json:"mac-tbl-sync-start-ts-b-5th"`
    ArpTblSyncEntriesSentM5th int `json:"arp-tbl-sync-entries-sent-m-5th"`
    Nd6TblSyncEntriesSentM5th int `json:"nd6-tbl-sync-entries-sent-m-5th"`
    Ipv4FibTblSyncEntriesSentM5th int `json:"ipv4-fib-tbl-sync-entries-sent-m-5th"`
    Ipv6FibTblSyncEntriesSentM5th int `json:"ipv6-fib-tbl-sync-entries-sent-m-5th"`
    MacTblSyncEntriesSentM5th int `json:"mac-tbl-sync-entries-sent-m-5th"`
    ArpTblSyncEntriesRcvdB5th int `json:"arp-tbl-sync-entries-rcvd-b-5th"`
    Nd6TblSyncEntriesRcvdB5th int `json:"nd6-tbl-sync-entries-rcvd-b-5th"`
    Ipv4FibTblSyncEntriesRcvdB5th int `json:"ipv4-fib-tbl-sync-entries-rcvd-b-5th"`
    Ipv6FibTblSyncEntriesRcvdB5th int `json:"ipv6-fib-tbl-sync-entries-rcvd-b-5th"`
    MacTblSyncEntriesRcvdB5th int `json:"mac-tbl-sync-entries-rcvd-b-5th"`
    ArpTblSyncEntriesAddedB5th int `json:"arp-tbl-sync-entries-added-b-5th"`
    Nd6TblSyncEntriesAddedB5th int `json:"nd6-tbl-sync-entries-added-b-5th"`
    Ipv4FibTblSyncEntriesAddedB5th int `json:"ipv4-fib-tbl-sync-entries-added-b-5th"`
    Ipv6FibTblSyncEntriesAddedB5th int `json:"ipv6-fib-tbl-sync-entries-added-b-5th"`
    MacTblSyncEntriesAddedB5th int `json:"mac-tbl-sync-entries-added-b-5th"`
    ArpTblSyncEntriesRemovedB5th int `json:"arp-tbl-sync-entries-removed-b-5th"`
    Nd6TblSyncEntriesRemovedB5th int `json:"nd6-tbl-sync-entries-removed-b-5th"`
    Ipv4FibTblSyncEntriesRemovedB5th int `json:"ipv4-fib-tbl-sync-entries-removed-b-5th"`
    Ipv6FibTblSyncEntriesRemovedB5th int `json:"ipv6-fib-tbl-sync-entries-removed-b-5th"`
    MacTblSyncEntriesRemovedB5th int `json:"mac-tbl-sync-entries-removed-b-5th"`
    ArpTblSyncEndTsM5th int `json:"arp-tbl-sync-end-ts-m-5th"`
    Nd6TblSyncEndTsM5th int `json:"nd6-tbl-sync-end-ts-m-5th"`
    Ipv4FibTblSyncEndTsM5th int `json:"ipv4-fib-tbl-sync-end-ts-m-5th"`
    Ipv6FibTblSyncEndTsM5th int `json:"ipv6-fib-tbl-sync-end-ts-m-5th"`
    MacTblSyncEndTsM5th int `json:"mac-tbl-sync-end-ts-m-5th"`
    ArpTblSyncEndTsB5th int `json:"arp-tbl-sync-end-ts-b-5th"`
    Nd6TblSyncEndTsB5th int `json:"nd6-tbl-sync-end-ts-b-5th"`
    Ipv4FibTblSyncEndTsB5th int `json:"ipv4-fib-tbl-sync-end-ts-b-5th"`
    Ipv6FibTblSyncEndTsB5th int `json:"ipv6-fib-tbl-sync-end-ts-b-5th"`
    MacTblSyncEndTsB5th int `json:"mac-tbl-sync-end-ts-b-5th"`
    ArpTblSyncM int `json:"arp-tbl-sync-m"`
    Nd6TblSyncM int `json:"nd6-tbl-sync-m"`
    Ipv4FibTblSyncM int `json:"ipv4-fib-tbl-sync-m"`
    Ipv6FibTblSyncM int `json:"ipv6-fib-tbl-sync-m"`
    MacTblSyncM int `json:"mac-tbl-sync-m"`
    ArpTblSyncB int `json:"arp-tbl-sync-b"`
    Nd6TblSyncB int `json:"nd6-tbl-sync-b"`
    Ipv4FibTblSyncB int `json:"ipv4-fib-tbl-sync-b"`
    Ipv6FibTblSyncB int `json:"ipv6-fib-tbl-sync-b"`
    MacTblSyncB int `json:"mac-tbl-sync-b"`
    ArpTblCksumM int `json:"arp-tbl-cksum-m"`
    Nd6TblCksumM int `json:"nd6-tbl-cksum-m"`
    Ipv4FibTblCksumM int `json:"ipv4-fib-tbl-cksum-m"`
    Ipv6FibTblCksumM int `json:"ipv6-fib-tbl-cksum-m"`
    MacTblCksumM int `json:"mac-tbl-cksum-m"`
    ArpTblCksumB int `json:"arp-tbl-cksum-b"`
    Nd6TblCksumB int `json:"nd6-tbl-cksum-b"`
    Ipv4FibTblCksumB int `json:"ipv4-fib-tbl-cksum-b"`
    Ipv6FibTblCksumB int `json:"ipv6-fib-tbl-cksum-b"`
    MacTblCksumB int `json:"mac-tbl-cksum-b"`
    ArpTblCksumMismatchB int `json:"arp-tbl-cksum-mismatch-b"`
    Nd6TblCksumMismatchB int `json:"nd6-tbl-cksum-mismatch-b"`
    Ipv4FibTblCksumMismatchB int `json:"ipv4-fib-tbl-cksum-mismatch-b"`
    Ipv6FibTblCksumMismatchB int `json:"ipv6-fib-tbl-cksum-mismatch-b"`
    MacTblCksumMismatchB int `json:"mac-tbl-cksum-mismatch-b"`
    ArpTblCksumCancelM int `json:"arp-tbl-cksum-cancel-m"`
    Nd6TblCksumCancelM int `json:"nd6-tbl-cksum-cancel-m"`
    Ipv4FibTblCksumCancelM int `json:"ipv4-fib-tbl-cksum-cancel-m"`
    Ipv6FibTblCksumCancelM int `json:"ipv6-fib-tbl-cksum-cancel-m"`
    MacTblCksumCancelM int `json:"mac-tbl-cksum-cancel-m"`
}

func (p *SystemTableIntegrityStats) GetId() string{
    return "1"
}

func (p *SystemTableIntegrityStats) getPath() string{
    return "system/table-integrity/stats"
}

func (p *SystemTableIntegrityStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemTableIntegrityStats,error) {
logger.Println("SystemTableIntegrityStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemTableIntegrityStats
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
