package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateQuicActionOnInitial() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_quic_action_on_initial`: Configure action on initial\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateQuicActionOnInitialCreate,
		UpdateContext: resourceDdosZoneTemplateQuicActionOnInitialUpdate,
		ReadContext:   resourceDdosZoneTemplateQuicActionOnInitialRead,
		DeleteContext: resourceDdosZoneTemplateQuicActionOnInitialDelete,

		Schema: map[string]*schema.Schema{
			"retry_token_authentication_dynamic": {
				Type: schema.TypeString, Optional: true, Description: "'cid-hash': Dynamic retry token based on Destination and Source CID;",
			},
			"retry_token_authentication_static": {
				Type: schema.TypeString, Optional: true, Description: "Static retry token (Maximum of 4 bytes)",
			},
			"retry_unauthenticated_initial_decrypt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Decrypt incoming initial packets for unauthenticated clients",
			},
			"scid_length": {
				Type: schema.TypeInt, Optional: true, Default: 20, Description: "Set Max SCID length for SCID generation",
			},
			"unauth_short_hdr_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop short header packet;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"quic_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "QuicTmplName",
			},
		},
	}
}
func resourceDdosZoneTemplateQuicActionOnInitialCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicActionOnInitialCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuicActionOnInitial(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateQuicActionOnInitialRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateQuicActionOnInitialUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicActionOnInitialUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuicActionOnInitial(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateQuicActionOnInitialRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateQuicActionOnInitialDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicActionOnInitialDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuicActionOnInitial(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateQuicActionOnInitialRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicActionOnInitialRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuicActionOnInitial(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneTemplateQuicActionOnInitial(d *schema.ResourceData) edpt.DdosZoneTemplateQuicActionOnInitial {
	var ret edpt.DdosZoneTemplateQuicActionOnInitial
	ret.Inst.RetryTokenAuthenticationDynamic = d.Get("retry_token_authentication_dynamic").(string)
	ret.Inst.RetryTokenAuthenticationStatic = d.Get("retry_token_authentication_static").(string)
	ret.Inst.RetryUnauthenticatedInitialDecrypt = d.Get("retry_unauthenticated_initial_decrypt").(int)
	ret.Inst.ScidLength = d.Get("scid_length").(int)
	ret.Inst.UnauthShortHdrAction = d.Get("unauth_short_hdr_action").(string)
	//omit uuid
	ret.Inst.QuicTmplName = d.Get("quic_tmpl_name").(string)
	return ret
}
