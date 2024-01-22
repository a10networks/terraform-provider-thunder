package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugQuic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_quic`: Debug QUIC\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugQuicCreate,
		UpdateContext: resourceDebugQuicUpdate,
		ReadContext:   resourceDebugQuicRead,
		DeleteContext: resourceDebugQuicDelete,

		Schema: map[string]*schema.Schema{
			"dummy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
			},
			"level": {
				Type: schema.TypeInt, Optional: true, Description: "Debug level (Level 1-5)",
			},
			"tls": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "QUIC TLS",
			},
			"tls_level": {
				Type: schema.TypeInt, Optional: true, Description: "Debug level (Level 1-5)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugQuicCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugQuicCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugQuic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugQuicRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugQuicUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugQuicUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugQuic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugQuicRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugQuicDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugQuicDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugQuic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugQuicRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugQuicRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugQuic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugQuic(d *schema.ResourceData) edpt.DebugQuic {
	var ret edpt.DebugQuic
	ret.Inst.Dummy = d.Get("dummy").(int)
	ret.Inst.Level = d.Get("level").(int)
	ret.Inst.Tls = d.Get("tls").(int)
	ret.Inst.TlsLevel = d.Get("tls_level").(int)
	//omit uuid
	return ret
}
