package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowMonitorCustomRecord() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_monitor_custom_record`: Configure custom record types to be exported\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowMonitorCustomRecordCreate,
		UpdateContext: resourceNetflowMonitorCustomRecordUpdate,
		ReadContext:   resourceNetflowMonitorCustomRecordRead,
		DeleteContext: resourceNetflowMonitorCustomRecordDelete,

		Schema: map[string]*schema.Schema{
			"custom_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"event": {
							Type: schema.TypeString, Optional: true, Description: "'sesn-event-nat44-creation': Export NAT44 session creation events; 'sesn-event-nat44-deletion': Export NAT44 session deletion events; 'sesn-event-nat64-creation': Export NAT64 session creation events; 'sesn-event-nat64-deletion': Export NAT64 session deletion events; 'sesn-event-dslite-creation': Export Dslite session creation events; 'sesn-event-dslite-deletion': Export Dslite session deletion events; 'sesn-event-fw4-creation': Export FW4 session creation events; 'sesn-event-fw4-deletion': Export FW4 session deletion events; 'sesn-event-fw6-creation': Export FW6 session creation events; 'sesn-event-fw6-deletion': Export FW6 session deletion events; 'deny-reset-event-fw4': Export FW4 Deny Reset events; 'deny-reset-event-fw6': Export FW6 Deny Reset events; 'port-mapping-nat44-creation': Export NAT44 Port Mapping Creation Event; 'port-mapping-nat44-deletion': Export NAT44 Port Mapping Deletion Event; 'port-mapping-nat64-creation': Export NAT64 Port Mapping Creation Event; 'port-mapping-nat64-deletion': Export NAT64 Port Mapping Deletion Event; 'port-mapping-dslite-creation': Export Dslite Port Mapping Creation Event; 'port-mapping-dslite-deletion': Export Dslite Port Mapping Deletion Event; 'port-batch-nat44-creation': Export NAT44 Port Batch Creation Event; 'port-batch-nat44-deletion': Export NAT44 Port Batch Deletion Event; 'port-batch-nat64-creation': Export NAT64 Port Batch Creation Event; 'port-batch-nat64-deletion': Export NAT64 Port Batch Deletion Event; 'port-batch-dslite-creation': Export Dslite Port Batch Creation Event; 'port-batch-dslite-deletion': Export Dslite Port Batch Deletion Event; 'port-batch-v2-nat44-creation': Export NAT44 Port Batch v2 Creation Event; 'port-batch-v2-nat44-deletion': Export NAT44 Port Batch v2 Deletion Event; 'port-batch-v2-nat64-creation': Export NAT64 Port Batch v2 Creation Event; 'port-batch-v2-nat64-deletion': Export NAT64 Port Batch v2 Deletion Event; 'port-batch-v2-dslite-creation': Export Dslite Port Batch v2 Creation Event; 'port-batch-v2-dslite-deletion': Export Dslite Port Batch v2 Deletion Event; 'gtp-c-tunnel-event': Export GTP Control Tunnel Creation or Deletion Events; 'gtp-u-tunnel-event': Export GTP User Tunnel Creation or Deletion Events; 'gtp-deny-event': Export GTP Deny events on GTP C/U Tunnels; 'gtp-info-event': Export GTP Info events on GTP C/U Tunnels; 'fw-ddos-entry-creation': Export FW iDDoS Entry Created Record; 'fw-ddos-entry-deletion': Export FW iDDoS Entry Deleted Record; 'fw-session-limit-exceeded': Export FW Session Limit Exceeded Record; 'cgn-ddos-l3-entry-creation': Export CGN iDDoS L3 Entry Creation; 'cgn-ddos-l3-entry-deletion': Export CGN iDDoS L3 Entry Deletion; 'cgn-ddos-l4-entry-creation': Export CGN iDDoS L4 Entry Creation; 'cgn-ddos-l4-entry-deletion': Export CGN iDDoS L4 Entry Deletion; 'gtp-rate-limit-periodic': Export GTP Rate Limit Periodic;",
						},
						"ipfix_template": {
							Type: schema.TypeString, Optional: true, Description: "Custom IPFIX Template",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceNetflowMonitorCustomRecordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorCustomRecordCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorCustomRecord(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorCustomRecordRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowMonitorCustomRecordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorCustomRecordUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorCustomRecord(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorCustomRecordRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowMonitorCustomRecordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorCustomRecordDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorCustomRecord(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowMonitorCustomRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorCustomRecordRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorCustomRecord(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetflowMonitorCustomRecordCustomCfg(d []interface{}) []edpt.NetflowMonitorCustomRecordCustomCfg {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorCustomRecordCustomCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorCustomRecordCustomCfg
		oi.Event = in["event"].(string)
		oi.IpfixTemplate = in["ipfix_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetflowMonitorCustomRecord(d *schema.ResourceData) edpt.NetflowMonitorCustomRecord {
	var ret edpt.NetflowMonitorCustomRecord
	ret.Inst.CustomCfg = getSliceNetflowMonitorCustomRecordCustomCfg(d.Get("custom_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
