package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_cgnv6_dns64_vs_port_tmpl_trigger_stats_inc`: Configure stats to triggers packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"es_total_failure_actions": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total failure actions",
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
func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsInc
	ret.Inst.Es_total_failure_actions = d.Get("es_total_failure_actions").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
