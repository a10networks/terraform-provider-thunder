package thunder

import (
	"context"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi/file"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

//Not direct mapping to any axapi endpoint
func resourceFileCsr() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFileCsrCreate,
		ReadContext:   resourceFileCsrRead,
		DeleteContext: resourceFileCsrDelete,
		Schema: map[string]*schema.Schema{
			"cert_type": {
				Type: schema.TypeString, Optional: true, ForceNew: true, Default: "rsa", Description: "Certificate type",
				ValidateFunc: validation.StringInSlice([]string{"rsa", "ecdsa"}, false),
			},
			"common_name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Common name",
				ValidateFunc: validation.StringLenBetween(1, 64),
			},
			"country": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Country",
				ValidateFunc: validation.StringLenBetween(2, 3),
			},
			"digest": {
				Type: schema.TypeString, Optional: true, ForceNew: true, Default: "sha1", Description: "Digest type",
				ValidateFunc: validation.StringInSlice([]string{"sha1", "sha256", "sha384", "sha512"}, false),
			},
			"division": {
				Type: schema.TypeString, Optional: true, ForceNew: true, Description: "Division",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"email": {
				Type: schema.TypeString, Optional: true, ForceNew: true, Description: "Email address",
				ValidateFunc: validation.StringLenBetween(1, 64),
			},
			"key_size": {
				Type: schema.TypeString, Optional: true, ForceNew: true, Default: "1024", Description: "Key size",
				ValidateFunc: validation.StringInSlice([]string{"256", "384", "1024", "2048", "4096"}, false),
			},
			"locality": {
				Type: schema.TypeString, Optional: true, ForceNew: true, Description: "Locality",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Local file name",
				ValidateFunc: validation.StringLenBetween(1, 245),
			},
			"organization": {
				Type: schema.TypeString, Optional: true, ForceNew: true, Description: "Organization",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"secured": {
				Type: schema.TypeInt, Optional: true, ForceNew: true, Default: 0, Description: "Mark key as non-exportable",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"state_province": {
				Type: schema.TypeString, Optional: true, ForceNew: true, Description: "State province",
				ValidateFunc: validation.StringLenBetween(1, 31),
			},
			"valid_days": {
				Type: schema.TypeInt, Optional: true, ForceNew: true, Default: 730, Description: "Valid days",
				ValidateFunc: validation.IntBetween(30, 3650),
			},
		},
	}
}

func resourceFileCsrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCsrCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToFileCsr(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileCsrRead(ctx, d, meta)
	}
	return diags
}

func resourceFileCsrRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCsrRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := file.SslCsr{}
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileCsrDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCsrDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := file.SslCsr{}
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToFileCsr(d *schema.ResourceData) file.SslCsr {
	var ret file.SslCsr
	ret.CertType = d.Get("cert_type").(string)
	ret.CommonName = d.Get("common_name").(string)
	ret.Country = d.Get("country").(string)
	ret.Digest = d.Get("digest").(string)
	ret.Division = d.Get("division").(string)
	ret.Email = d.Get("email").(string)
	ret.KeySize = d.Get("key_size").(string)
	ret.Locality = d.Get("locality").(string)
	ret.Name = d.Get("name").(string)
	ret.Organization = d.Get("organization").(string)
	ret.Secured = d.Get("secured").(int)
	ret.StateProvince = d.Get("state_province").(string)
	ret.ValidDays = d.Get("valid_days").(int)
	return ret
}
