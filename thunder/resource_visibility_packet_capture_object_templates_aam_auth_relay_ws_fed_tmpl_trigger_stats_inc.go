package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_aam_auth_relay_ws_fed_tmpl_trigger_stats_inc`: Configure stats to triggers packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failure",
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
func resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplTriggerStatsInc
	ret.Inst.Failure = d.Get("failure").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
