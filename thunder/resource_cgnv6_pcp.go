package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Pcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_pcp`: Set Port Control Protocol parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6PcpCreate,
		UpdateContext: resourceCgnv6PcpUpdate,
		ReadContext:   resourceCgnv6PcpRead,
		DeleteContext: resourceCgnv6PcpDelete,

		Schema: map[string]*schema.Schema{
			"default_template": {
				Type: schema.TypeString, Optional: true, Description: "Bind the default template for PCP (Bind a PCP template)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'packets-rcv': Packets Received; 'lsn-map-process-success': PCP MAP Request Processing Success (NAT44); 'dslite-map-process-success': PCP MAP Request Processing Success (DS-Lite); 'nat64-map-process-success': PCP MAP Request Processing Success (NAT64); 'lsn-peer-process-success': PCP PEER Request Processing Success (NAT44); 'dslite-peer-process-success': PCP PEER Request Processing Success (DS-Lite); 'nat64-peer-process-success': PCP PEER Request Processing Success (NAT64); 'lsn-announce-process-success': PCP ANNOUNCE Request Processing Success (NAT44); 'dslite-announce-process-success': PCP ANNOUNCE Request Processing Success (DS-Lite); 'nat64-announce-process-success': PCP ANNOUNCE Request Processing Success (NAT64); 'pkt-not-request-drop': Packet Not a PCP Request; 'pkt-too-short-drop': Packet Too Short; 'noroute-drop': Response No Route; 'unsupported-version': Unsupported PCP version; 'not-authorized': PCP Request Not Authorized; 'malform-request': PCP Request Malformed; 'unsupp-opcode': Unsupported PCP Opcode; 'unsupp-option': Unsupported PCP Option; 'malform-option': PCP Option Malformed; 'no-resources': No System or NAT Resources; 'unsupp-protocol': Unsupported Mapping Protocol; 'user-quota-exceeded': User Quota Exceeded; 'cannot-provide-suggest': Cannot Provide Suggested Port When PREFER_FAILURE; 'address-mismatch': PCP Client Address Mismatch; 'excessive-remote-peers': Excessive Remote Peers; 'pkt-not-from-nat-inside': Packet Dropped For Not Coming From NAT Inside; 'l4-process-error': L3/L4 Process Error; 'internal-error-drop': Internal Error; 'unsol_ance_sent_succ': Unsolicited Announce Sent; 'unsol_ance_sent_fail': Unsolicited Announce Send Failure; 'ha_sync_epoch_sent': HA Sync PCP Epoch Sent; 'ha_sync_epoch_rcv': HA Sync PCP Epoch Recv; 'fullcone-ext-alloc': PCP Fullcone Extension Alloc; 'fullcone-ext-free': PCP Fullcone Extension Free; 'fullcone-ext-alloc-failure': PCP Fullcone Extension Alloc Failure; 'fullcone-ext-notfound': PCP Fullcone Extension Not Found; 'fullcone-ext-reuse': PCP Fullcone Extension Reuse; 'client-nonce-mismatch': PCP Client Nonce Mismatch; 'map-filter-set': PCP MAP Filter Set; 'map-filter-deny': PCP MAP Filter Deny Inbound; 'inter-board-pkts': PCP Inter board packets;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6PcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6PcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Pcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6PcpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6PcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6PcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Pcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6PcpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6PcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6PcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Pcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6PcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6PcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Pcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6PcpSamplingEnable(d []interface{}) []edpt.Cgnv6PcpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6PcpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6PcpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Pcp(d *schema.ResourceData) edpt.Cgnv6Pcp {
	var ret edpt.Cgnv6Pcp
	ret.Inst.DefaultTemplate = d.Get("default_template").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6PcpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
