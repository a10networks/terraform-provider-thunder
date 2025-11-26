package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePkiDelete() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_pki_delete`: Delete SSL cert\n\n__PLACEHOLDER__",
		CreateContext: resourcePkiDeleteCreate,
		UpdateContext: resourcePkiDeleteUpdate,
		ReadContext:   resourcePkiDeleteRead,
		DeleteContext: resourcePkiDeleteDelete,

		Schema: map[string]*schema.Schema{
			"ca": {
				Type: schema.TypeString, Optional: true, Description: "CA certificate file name",
			},
			"cert_name": {
				Type: schema.TypeString, Optional: true, Description: "Certificate file name",
			},
			"crl": {
				Type: schema.TypeString, Optional: true, Description: "CRL file name",
			},
			"csr": {
				Type: schema.TypeString, Optional: true, Description: "CSR file name",
			},
			"private_key": {
				Type: schema.TypeString, Optional: true, Description: "Private key file name",
			},
		},
	}
}
func resourcePkiDeleteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiDeleteCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiDelete(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiDeleteRead(ctx, d, meta)
	}
	return diags
}

func resourcePkiDeleteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiDeleteUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiDelete(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiDeleteRead(ctx, d, meta)
	}
	return diags
}
func resourcePkiDeleteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiDeleteDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiDelete(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePkiDeleteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiDeleteRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiDelete(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointPkiDelete(d *schema.ResourceData) edpt.PkiDelete {
	var ret edpt.PkiDelete
	ret.Inst.Ca = d.Get("ca").(string)
	ret.Inst.CertName = d.Get("cert_name").(string)
	ret.Inst.Crl = d.Get("crl").(string)
	ret.Inst.Csr = d.Get("csr").(string)
	ret.Inst.PrivateKey = d.Get("private_key").(string)
	return ret
}
