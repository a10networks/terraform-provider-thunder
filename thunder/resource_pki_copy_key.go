package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePkiCopyKey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_pki_copy_key`: Copy SSL key to another file\n\n__PLACEHOLDER__",
		CreateContext: resourcePkiCopyKeyCreate,
		UpdateContext: resourcePkiCopyKeyUpdate,
		ReadContext:   resourcePkiCopyKeyRead,
		DeleteContext: resourcePkiCopyKeyDelete,

		Schema: map[string]*schema.Schema{
			"dest_key": {
				Type: schema.TypeString, Optional: true, Description: "Destination key file",
			},
			"overwrite": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite the destination file if already present",
			},
			"rotation": {
				Type: schema.TypeInt, Optional: true, Description: "Specify rotation number of SCEP/CMP generated key file",
			},
			"src_key": {
				Type: schema.TypeString, Optional: true, Description: "Source key file",
			},
		},
	}
}
func resourcePkiCopyKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCopyKeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCopyKey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiCopyKeyRead(ctx, d, meta)
	}
	return diags
}

func resourcePkiCopyKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCopyKeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCopyKey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiCopyKeyRead(ctx, d, meta)
	}
	return diags
}
func resourcePkiCopyKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCopyKeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCopyKey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePkiCopyKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCopyKeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCopyKey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointPkiCopyKey(d *schema.ResourceData) edpt.PkiCopyKey {
	var ret edpt.PkiCopyKey
	ret.Inst.DestKey = d.Get("dest_key").(string)
	ret.Inst.Overwrite = d.Get("overwrite").(int)
	ret.Inst.Rotation = d.Get("rotation").(int)
	ret.Inst.SrcKey = d.Get("src_key").(string)
	return ret
}
