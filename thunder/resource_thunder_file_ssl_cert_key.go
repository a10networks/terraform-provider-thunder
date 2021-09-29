package thunder

//Thunder resource FileSslCertKey

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFileSslCertKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFileSslCertKeyCreate,
		UpdateContext: resourceFileSslCertKeyUpdate,
		ReadContext:   resourceFileSslCertKeyRead,
		DeleteContext: resourceFileSslCertKeyDelete,
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
			"secured": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"file_local_path": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "local path for ssl-cert file",
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
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
		},
	}
}

func resourceFileSslCertKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating FileSslCertKey (Inside resourceFileSslCertKeyCreate) ")

		data := dataToFileSslCertKey(d)
		FileHandleLocal := d.Get("file_local_path").(string)
		logger.Println("[INFO] received formatted data from method data to FileSslCertKey --")
		d.SetId("1")
		err := go_thunder.PostFileSslCertKey(client.Token, data, client.Host, FileHandleLocal)
		if err != nil {
			return diag.FromErr(err)
		}
		return diags
		//return resourceFileSslCertKeyRead(ctx, d, meta)

	}
	return diags
}

func resourceFileSslCertKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FileSslCertKey (Inside resourceFileSslCertKeyRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetFileSslCertKey(client.Token, client.Host)
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

func resourceFileSslCertKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFileSslCertKeyRead(ctx, d, meta)
}

func resourceFileSslCertKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFileSslCertKeyRead(ctx, d, meta)
}
func dataToFileSslCertKey(d *schema.ResourceData) go_thunder.FileSslCertKey {
	var vc go_thunder.FileSslCertKey
	var c go_thunder.FileSslCertKeyInstance
	c.FileSslCertKeyInstanceFile = d.Get("file").(string)
	c.FileSslCertKeyInstanceSize = d.Get("size").(int)
	c.FileSslCertKeyInstanceFileHandle = d.Get("file_handle").(string)
	c.FileSslCertKeyInstanceSecured = d.Get("secured").(int)
	c.FileSslCertKeyInstanceAction = d.Get("action").(string)
	c.FileSslCertKeyInstanceDstFile = d.Get("dst_file").(string)

	vc.FileSslCertKeyInstanceFile = c
	return vc
}
