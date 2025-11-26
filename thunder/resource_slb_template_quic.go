package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateQuic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_quic`: QUIC template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateQuicCreate,
		UpdateContext: resourceSlbTemplateQuicUpdate,
		ReadContext:   resourceSlbTemplateQuicRead,
		DeleteContext: resourceSlbTemplateQuicDelete,

		Schema: map[string]*schema.Schema{
			"burst_len": {
				Type: schema.TypeInt, Optional: true, Description: "Number of burst packet, default 16",
			},
			"connection_id_length": {
				Type: schema.TypeInt, Optional: true, Description: "Connection id length in byte, default 8 bytes",
			},
			"idle_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Idle Timeout (interval of 60 seconds), default 120 seconds (idle timeout in second, default 120)",
			},
			"initial_wnd": {
				Type: schema.TypeInt, Optional: true, Description: "Initial window size in byte, default 10000 (Initial window size, default 10000)",
			},
			"key_update_to_client": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Initiate key update on the client-side",
			},
			"key_update_to_server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Initiate key update on the server-side",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "QUIC Template Name",
			},
			"receive_buffer": {
				Type: schema.TypeInt, Optional: true, Description: "Receive buffer size in byte, default 200000 (Receive buffer size, default 200000)",
			},
			"server_retry": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable server retry",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplateQuicCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateQuicCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateQuic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateQuicRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateQuicUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateQuicUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateQuic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateQuicRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateQuicDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateQuicDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateQuic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateQuicRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateQuicRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateQuic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateQuic(d *schema.ResourceData) edpt.SlbTemplateQuic {
	var ret edpt.SlbTemplateQuic
	ret.Inst.BurstLen = d.Get("burst_len").(int)
	ret.Inst.ConnectionIdLength = d.Get("connection_id_length").(int)
	ret.Inst.IdleTimeout = d.Get("idle_timeout").(int)
	ret.Inst.InitialWnd = d.Get("initial_wnd").(int)
	ret.Inst.KeyUpdateToClient = d.Get("key_update_to_client").(int)
	ret.Inst.KeyUpdateToServer = d.Get("key_update_to_server").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.ReceiveBuffer = d.Get("receive_buffer").(int)
	ret.Inst.ServerRetry = d.Get("server_retry").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
