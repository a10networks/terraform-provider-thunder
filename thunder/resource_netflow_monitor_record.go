package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowMonitorRecord() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_monitor_record`: Configure record types to be exported\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowMonitorRecordCreate,
		UpdateContext: resourceNetflowMonitorRecordUpdate,
		ReadContext:   resourceNetflowMonitorRecordRead,
		DeleteContext: resourceNetflowMonitorRecordDelete,

		Schema: map[string]*schema.Schema{
			"ddos_general_stat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "General DDOS statistics",
			},
			"ddos_http_stat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "HTTP DDOS statistics",
			},
			"dslite": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "DS-Lite Flow Record Template",
			},
			"nat44": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "NAT44 Flow Record Template",
			},
			"nat64": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "NAT64 Flow Record Template",
			},
			"netflow_v5": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "NetFlow V5 Flow Record Template",
			},
			"netflow_v5_ext": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Extended NetFlow V5 Flow Record Template, supports ipv6",
			},
			"port_batch_dslite": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
			},
			"port_batch_nat44": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
			},
			"port_batch_nat64": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
			},
			"port_batch_v2_dslite": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
			},
			"port_batch_v2_nat44": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
			},
			"port_batch_v2_nat64": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
			},
			"port_mapping_dslite": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
			},
			"port_mapping_nat44": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
			},
			"port_mapping_nat64": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
			},
			"sesn_event_dslite": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
			},
			"sesn_event_fw4": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
			},
			"sesn_event_fw6": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
			},
			"sesn_event_nat44": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
			},
			"sesn_event_nat64": {
				Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
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
func resourceNetflowMonitorRecordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorRecordCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorRecord(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorRecordRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowMonitorRecordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorRecordUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorRecord(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorRecordRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowMonitorRecordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorRecordDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorRecord(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowMonitorRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorRecordRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorRecord(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetflowMonitorRecord(d *schema.ResourceData) edpt.NetflowMonitorRecord {
	var ret edpt.NetflowMonitorRecord
	ret.Inst.DdosGeneralStat = d.Get("ddos_general_stat").(int)
	ret.Inst.DdosHttpStat = d.Get("ddos_http_stat").(int)
	ret.Inst.Dslite = d.Get("dslite").(int)
	ret.Inst.Nat44 = d.Get("nat44").(int)
	ret.Inst.Nat64 = d.Get("nat64").(int)
	ret.Inst.NetflowV5 = d.Get("netflow_v5").(int)
	ret.Inst.NetflowV5Ext = d.Get("netflow_v5_ext").(int)
	ret.Inst.PortBatchDslite = d.Get("port_batch_dslite").(string)
	ret.Inst.PortBatchNat44 = d.Get("port_batch_nat44").(string)
	ret.Inst.PortBatchNat64 = d.Get("port_batch_nat64").(string)
	ret.Inst.PortBatchV2Dslite = d.Get("port_batch_v2_dslite").(string)
	ret.Inst.PortBatchV2Nat44 = d.Get("port_batch_v2_nat44").(string)
	ret.Inst.PortBatchV2Nat64 = d.Get("port_batch_v2_nat64").(string)
	ret.Inst.PortMappingDslite = d.Get("port_mapping_dslite").(string)
	ret.Inst.PortMappingNat44 = d.Get("port_mapping_nat44").(string)
	ret.Inst.PortMappingNat64 = d.Get("port_mapping_nat64").(string)
	ret.Inst.SesnEventDslite = d.Get("sesn_event_dslite").(string)
	ret.Inst.SesnEventFw4 = d.Get("sesn_event_fw4").(string)
	ret.Inst.SesnEventFw6 = d.Get("sesn_event_fw6").(string)
	ret.Inst.SesnEventNat44 = d.Get("sesn_event_nat44").(string)
	ret.Inst.SesnEventNat64 = d.Get("sesn_event_nat64").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
