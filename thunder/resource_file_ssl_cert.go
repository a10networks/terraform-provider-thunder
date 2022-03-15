package thunder

import (
	"context"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi/file"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

//Not direct mapping to any axapi endpoint
func resourceFileSslCert() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFileSslCertCreate,
		ReadContext:   resourceFileSslCertRead,
		UpdateContext: resourceFileSslCertUpdate,
		DeleteContext: resourceFileSslCertDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Local file name",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Transfer protocol",
				ValidateFunc: validation.StringInSlice([]string{"ftp", "http", "https", "scp", "sftp", "tftp"}, false),
			},
			"host": {
				Type: schema.TypeString, Required: true, Description: "Remote site (IP or domain name)",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"path": {
				Type: schema.TypeString, Required: true, Description: "Remote path",
				ValidateFunc: axapi.IsAbsolutePath,
			},
			"username": {
				Type: schema.TypeString, Optional: true, Description: "Username for the remote site",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "Password for the remote site",
				ValidateFunc: validation.StringLenBetween(1, 128),
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"overwrite": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"certificate_type": {
				Type: schema.TypeString, Optional: true, Default: "pem", Description: "Certificate format",
				ValidateFunc: validation.StringInSlice([]string{"pem", "der", "p7b", "pfx"}, false),
			},
			"pfx_password": {
				Type: schema.TypeString, Optional: true, Description: "Password for pfx format",
				ValidateFunc: validation.StringLenBetween(1, 128),
			},
			"secured": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mark key as non-exportable (for pfx format)",
				ValidateFunc: validation.IntBetween(0, 1),
			},
		},
	}
}

func resourceFileSslCertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslCertCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToFileSslCert(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSslCertRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSslCertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslCertRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := file.SslCert{}
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSslCertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslCertUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToFileSslCert(d)
		//file.SslCert is immutable. Overwrite it.
		obj.Overwrite = 1
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSslCertRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSslCertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslCertDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := file.SslCert{}
		obj.CertificateType = d.Get("certificate_type").(string)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToFileSslCert(d *schema.ResourceData) file.SslCert {
	var ret file.SslCert
	ret.Name = d.Get("name").(string)
	ret.Protocol = d.Get("protocol").(string)
	ret.Host = d.Get("host").(string)
	ret.Path = d.Get("path").(string)
	ret.Username = d.Get("username").(string)
	ret.Password = d.Get("password").(string)
	ret.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Overwrite = d.Get("overwrite").(int)
	ret.CertificateType = d.Get("certificate_type").(string)
	ret.PfxPassword = d.Get("pfx_password").(string)
	ret.Secured = d.Get("secured").(int)
	return ret
}
