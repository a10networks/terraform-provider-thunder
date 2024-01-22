package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePkiCopyCert() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_pki_copy_cert`: Copy SSL certificate to another file\n\n__PLACEHOLDER__",
		CreateContext: resourcePkiCopyCertCreate,
		UpdateContext: resourcePkiCopyCertUpdate,
		ReadContext:   resourcePkiCopyCertRead,
		DeleteContext: resourcePkiCopyCertDelete,

		Schema: map[string]*schema.Schema{
			"dest_cert": {
				Type: schema.TypeString, Optional: true, Description: "Destination certificate file",
			},
			"overwrite": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite the destination file if already present",
			},
			"rotation": {
				Type: schema.TypeInt, Optional: true, Description: "Specify rotation number of SCEP/CMP generated certificate file",
			},
			"src_cert": {
				Type: schema.TypeString, Optional: true, Description: "Source certificate file",
			},
		},
	}
}
func resourcePkiCopyCertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCopyCertCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCopyCert(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiCopyCertRead(ctx, d, meta)
	}
	return diags
}

func resourcePkiCopyCertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCopyCertUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCopyCert(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiCopyCertRead(ctx, d, meta)
	}
	return diags
}
func resourcePkiCopyCertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCopyCertDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCopyCert(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePkiCopyCertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCopyCertRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCopyCert(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointPkiCopyCert(d *schema.ResourceData) edpt.PkiCopyCert {
	var ret edpt.PkiCopyCert
	ret.Inst.DestCert = d.Get("dest_cert").(string)
	ret.Inst.Overwrite = d.Get("overwrite").(int)
	ret.Inst.Rotation = d.Get("rotation").(int)
	ret.Inst.SrcCert = d.Get("src_cert").(string)
	return ret
}
