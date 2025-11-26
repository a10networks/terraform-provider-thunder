package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonEntityTelemetryData() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_mon_entity_telemetry_data`: dummy schema for sflow exports\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonEntityTelemetryDataCreate,
		UpdateContext: resourceVisibilityMonEntityTelemetryDataUpdate,
		ReadContext:   resourceVisibilityMonEntityTelemetryDataRead,
		DeleteContext: resourceVisibilityMonEntityTelemetryDataDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'in_pkts': Monitored Entity telemetry Metric IN pkts; 'out_pkts': Monitored Entity telemetry Metric OUT pkts; 'in_bytes': Monitored Entity telemetry Metric IN bytes; 'out_bytes': Monitored Entity telemetry Metric OUT bytes; 'errors': Monitored Entity telemetry Metric ERRORS; 'in_small_pkt': Monitored Entity telemetry Metric IN SMALL pkt; 'in_frag': Monitored Entity telemetry Metric IN frag; 'out_small_pkt': Monitored Entity telemetry Metric OUT SMALL pkt; 'out_frag': Monitored Entity telemetry Metric OUT frag; 'new-conn': Monitored Entity telemetry Metric New Sessions; 'avg_data_cpu_util': Monitored Entity telemetry Metric Average data CPU utilization; 'outside_intf_util': Monitored Entity telemetry Metric Outside interface utilization; 'concurrent-conn': concurrent-conn; 'in_bytes_per_out_bytes': Monitored Entity telemetry Metric IN bytes per OUT bytes; 'drop_pkts_per_pkts': Monitored Entity telemetry Metric Drop pkts per pkts; 'tcp_in_syn': Monitored Entity telemetry Metric TCP IN syn; 'tcp_out_syn': Monitored Entity telemetry Metric TCP OUT syn; 'tcp_in_fin': Monitored Entity telemetry Metric TCP IN fin; 'tcp_out_fin': Monitored Entity telemetry Metric TCP OUT fin; 'tcp_in_payload': Monitored Entity telemetry Metric TCP IN payload; 'tcp_out_payload': Monitored Entity telemetry Metric TCP OUT payload; 'tcp_in_rexmit': Monitored Entity telemetry Metric TCP IN rexmit; 'tcp_out_rexmit': Monitored Entity telemetry Metric TCP OUT rexmit; 'tcp_in_rst': Monitored Entity telemetry Metric TCP IN rst; 'tcp_out_rst': Monitored Entity telemetry Metric TCP OUT rst; 'tcp_in_empty_ack': Monitored Entity telemetry Metric TCP_IN EMPTY ack; 'tcp_out_empty_ack': Monitored Entity telemetry Metric TCP OUT EMPTY ack; 'tcp_in_zero_wnd': Monitored Entity telemetry Metric TCP IN ZERO wnd; 'tcp_out_zero_wnd': Monitored Entity telemetry Metric TCP OUT ZERO wnd; 'tcp_conn_miss': Monitored Entity telemetry Metric TCP connection miss; 'tcp_fwd_syn_per_fin': Monitored Entity telemetry Metric TCP FWD SYN per FIN;",
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
func resourceVisibilityMonEntityTelemetryDataCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonEntityTelemetryDataCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonEntityTelemetryData(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonEntityTelemetryDataRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonEntityTelemetryDataUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonEntityTelemetryDataUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonEntityTelemetryData(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonEntityTelemetryDataRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonEntityTelemetryDataDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonEntityTelemetryDataDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonEntityTelemetryData(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonEntityTelemetryDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonEntityTelemetryDataRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonEntityTelemetryData(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVisibilityMonEntityTelemetryDataSamplingEnable(d []interface{}) []edpt.VisibilityMonEntityTelemetryDataSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonEntityTelemetryDataSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonEntityTelemetryDataSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityMonEntityTelemetryData(d *schema.ResourceData) edpt.VisibilityMonEntityTelemetryData {
	var ret edpt.VisibilityMonEntityTelemetryData
	ret.Inst.SamplingEnable = getSliceVisibilityMonEntityTelemetryDataSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
