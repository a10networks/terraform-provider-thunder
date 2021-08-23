package thunder

//Thunder resource FileSslCert

import (
	"context"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"util"
)

func resourceFileSslCert() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFileSslCertCreate,
		UpdateContext: resourceFileSslCertUpdate,
		ReadContext:   resourceFileSslCertRead,
		DeleteContext: resourceFileSslCertDelete,
		Schema: map[string]*schema.Schema{
			"file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"file_handle": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"file_local_path": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "local path for ssl-cert file",
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
			},
			"certificate_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pfx_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pfx_password_export": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dst_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFileSslCertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating FileSslCert (Inside resourceFileSslCertCreate) ")

		data := dataToFileSslCert(d)
		FileHandleLocal := d.Get("file_local_path").(string)
		logger.Println("[INFO] received formatted data from method data to FileSslCert --")
		d.SetId("1")
		err := go_thunder.PostFileSslCert(client.Token, data, client.Host, FileHandleLocal)
		if err != nil {
			return diag.FromErr(err)
		}

		//return resourceFileSslCertRead(ctx, d, meta)
		return diags

	}
	return diags
}

func resourceFileSslCertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FileSslCert (Inside resourceFileSslCertRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetFileSslCert(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return diags
	}
	return diags
}

func resourceFileSslCertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFileSslCertCreate(ctx, d, meta)
	//return nil
}

func resourceFileSslCertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFileSslCertRead(ctx, d, meta)
	//return nil
}
func dataToFileSslCert(d *schema.ResourceData) go_thunder.FileSslCert {
	var vc go_thunder.FileSslCert
	var c go_thunder.FileSslCertInstance
	c.File = d.Get("file").(string)
	c.Size = d.Get("size").(int)
	c.FileHandle = d.Get("file_handle").(string)
	//c.FileHandleLocal = d.Get("file_local_path").(string)
	c.CertificateType = d.Get("certificate_type").(string)
	c.PfxPassword = d.Get("pfx_password").(string)
	c.PfxPasswordExport = d.Get("pfx_password_export").(string)
	c.Action = d.Get("action").(string)
	c.DstFile = d.Get("dst_file").(string)

	vc.File = c
	return vc
}
