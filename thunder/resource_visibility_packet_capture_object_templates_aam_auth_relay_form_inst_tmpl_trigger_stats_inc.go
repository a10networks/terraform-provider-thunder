package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_aam_auth_relay_form_inst_tmpl_trigger_stats_inc`: Configure stats to triggers packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"bad_req": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Request",
			},
			"error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Server Error",
			},
			"invalid_cred": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid Credential",
			},
			"invalid_srv_rsp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid Server Response",
			},
			"not_fnd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Not Found",
			},
			"other_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Other Error",
			},
			"post_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for POST Failed",
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
func resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplTriggerStatsInc
	ret.Inst.Bad_req = d.Get("bad_req").(int)
	ret.Inst.Error = d.Get("error").(int)
	ret.Inst.Invalid_cred = d.Get("invalid_cred").(int)
	ret.Inst.Invalid_srv_rsp = d.Get("invalid_srv_rsp").(int)
	ret.Inst.Not_fnd = d.Get("not_fnd").(int)
	ret.Inst.Other_error = d.Get("other_error").(int)
	ret.Inst.Post_fail = d.Get("post_fail").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
