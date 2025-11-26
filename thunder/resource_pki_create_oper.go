package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePkiCreateOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_pki_create_oper`: \n\n__PLACEHOLDER__",
		ReadContext: resourcePkiCreateOperRead,

		Schema: map[string]*schema.Schema{
			"bits": {
				Type: schema.TypeString, Optional: true, Default: "1024", Description: "'256': 256; '384': 384; '1024': 1024; '2048': 2048; '4096': 4096;",
			},
			"cert_type": {
				Type: schema.TypeString, Optional: true, Default: "rsa", Description: "'rsa': rsa; 'ecdsa': ecdsa;",
			},
			"common_name": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"country": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"csr_generate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"digest": {
				Type: schema.TypeString, Optional: true, Default: "sha1", Description: "'sha1': sha1; 'sha256': sha256; 'sha384': sha384; 'sha512': sha512;",
			},
			"division": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"email": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"filename": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"locality": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"organization": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"rootca": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"secured": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"state_province": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"v3_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"valid_days": {
				Type: schema.TypeInt, Optional: true, Default: 730, Description: "",
			},
		},
	}
}

func resourcePkiCreateOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCreateOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCreateOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointPkiCreateOper(d *schema.ResourceData) edpt.PkiCreateOper {
	var ret edpt.PkiCreateOper

	ret.Bits = d.Get("bits").(string)

	ret.CertType = d.Get("cert_type").(string)

	ret.CommonName = d.Get("common_name").(string)

	ret.Country = d.Get("country").(string)

	ret.CsrGenerate = d.Get("csr_generate").(int)

	ret.Digest = d.Get("digest").(string)

	ret.Division = d.Get("division").(string)

	ret.Email = d.Get("email").(string)

	ret.Filename = d.Get("filename").(string)

	ret.Locality = d.Get("locality").(string)

	ret.Organization = d.Get("organization").(string)

	ret.Rootca = d.Get("rootca").(int)

	ret.Secured = d.Get("secured").(int)

	ret.StateProvince = d.Get("state_province").(string)

	ret.V3Request = d.Get("v3_request").(int)

	ret.ValidDays = d.Get("valid_days").(int)
	return ret
}
