package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbLinkProbeEntry() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_link_probe_entry`: Show SLB Link Probe entries\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbLinkProbeEntryCreate,
		UpdateContext: resourceSlbLinkProbeEntryUpdate,
		ReadContext:   resourceSlbLinkProbeEntryRead,
		DeleteContext: resourceSlbLinkProbeEntryDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_entries': Current Entry Count; 'total_created': Total Entry Created; 'total_inserted': Total Entry Inserted; 'total_ready_to_free': Total Entry Ready To Free; 'total_freed': Total Entry Freed; 'err_entry_create_failed': Entry Creation Failure; 'err_entry_create_oom': Entry Creation Out of Memory; 'err_entry_insert_failed': Entry Insert Failed; 'err_tmpl_probe_create_failed': Probe Template Creation Failure; 'err_tmpl_probe_create_oom': Probe Template Creation Out of Memory; 'total_http_probes_sent': Total HTTP Probes Sent out; 'total_http_response_received': Total HTTP responses received; 'total_http_response_good': Total HTTP responses matching probe template config; 'total_http_response_bad': Total HTTP responses not matching probe template config; 'total_tcp_err': Total TCP errors in probes sent out; 'err_smart_nat_alloc': Error creating Smart NAT Instance; 'err_smart_nat_port_alloc': Error obtaining Smart NAT source port; 'err_l4_sess_alloc': Error allocating L4 session for probe; 'err_probe_tcp_conn_send': Error in initiating TCP connection for probe; 'probe_tcp_conn_sent': TCP connection sent for probe;",
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
func resourceSlbLinkProbeEntryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbLinkProbeEntryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbLinkProbeEntry(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbLinkProbeEntryRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbLinkProbeEntryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbLinkProbeEntryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbLinkProbeEntry(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbLinkProbeEntryRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbLinkProbeEntryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbLinkProbeEntryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbLinkProbeEntry(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbLinkProbeEntryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbLinkProbeEntryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbLinkProbeEntry(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbLinkProbeEntrySamplingEnable(d []interface{}) []edpt.SlbLinkProbeEntrySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbLinkProbeEntrySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbLinkProbeEntrySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbLinkProbeEntry(d *schema.ResourceData) edpt.SlbLinkProbeEntry {
	var ret edpt.SlbLinkProbeEntry
	ret.Inst.SamplingEnable = getSliceSlbLinkProbeEntrySamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
