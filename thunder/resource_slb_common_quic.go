package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCommonQuic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_common_quic`: Configure QUIC related settings\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbCommonQuicCreate,
		UpdateContext: resourceSlbCommonQuicUpdate,
		ReadContext:   resourceSlbCommonQuicRead,
		DeleteContext: resourceSlbCommonQuicDelete,

		Schema: map[string]*schema.Schema{
			"cid_len": {
				Type: schema.TypeInt, Optional: true, Default: 4, Description: "Length of CID",
			},
			"cpu_offset": {
				Type: schema.TypeInt, Optional: true, Default: 12, Description: "Offset for Encoded CPU",
			},
			"enable_hash": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CID Hashing",
			},
			"enable_signature": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CID Signature Validation",
			},
			"quic_lb_offset": {
				Type: schema.TypeInt, Optional: true, Default: 8, Description: "Offset for QUIC-LB",
			},
			"signature": {
				Type: schema.TypeString, Optional: true, Description: "Set CID Signature",
			},
			"signature_len": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Offset for CID Signature",
			},
			"signature_offset": {
				Type: schema.TypeInt, Optional: true, Default: 4, Description: "Offset for CID Signature",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbCommonQuicCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonQuicCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonQuic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonQuicRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbCommonQuicUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonQuicUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonQuic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonQuicRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbCommonQuicDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonQuicDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonQuic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbCommonQuicRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonQuicRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonQuic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbCommonQuic(d *schema.ResourceData) edpt.SlbCommonQuic {
	var ret edpt.SlbCommonQuic
	ret.Inst.CidLen = d.Get("cid_len").(int)
	ret.Inst.CpuOffset = d.Get("cpu_offset").(int)
	ret.Inst.EnableHash = d.Get("enable_hash").(int)
	ret.Inst.EnableSignature = d.Get("enable_signature").(int)
	ret.Inst.QuicLbOffset = d.Get("quic_lb_offset").(int)
	ret.Inst.Signature = d.Get("signature").(string)
	ret.Inst.SignatureLen = d.Get("signature_len").(int)
	ret.Inst.SignatureOffset = d.Get("signature_offset").(int)
	//omit uuid
	return ret
}
