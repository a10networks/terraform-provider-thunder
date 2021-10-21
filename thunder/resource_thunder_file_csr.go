package thunder

//Thunder resource FileCsr

import (
	"context"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceFileCsr() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFileCsrCreate,
		UpdateContext: resourceFileCsrUpdate,
		ReadContext:   resourceFileCsrRead,
		DeleteContext: resourceFileCsrDelete,
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

func resourceFileCsrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating FileCsr (Inside resourceFileCsrCreate) ")
		data := dataToFileCsr(d)
		logger.Println("[INFO] received formatted data from method data to FileCsr --")
		d.SetId("1")
		err := go_thunder.PostFileCsr(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFileCsrRead(ctx, d, meta)

	}
	return diags
}

func resourceFileCsrRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FileCsr (Inside resourceFileCsrRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetFileCsr(client.Token, client.Host)
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

func resourceFileCsrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFileCsrRead(ctx, d, meta)
}

func resourceFileCsrDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFileCsrRead(ctx, d, meta)
}
func dataToFileCsr(d *schema.ResourceData) go_thunder.FileCsr {
	var vc go_thunder.FileCsr
	var c go_thunder.FileCsrInstance
	c.FileCsrInstanceFile = d.Get("file").(string)
	c.FileCsrInstanceSize = d.Get("size").(int)
	c.FileCsrInstanceFileHandle = d.Get("file_handle").(string)
	c.FileCsrInstanceAction = d.Get("action").(string)
	c.FileCsrInstanceDstFile = d.Get("dst_file").(string)

	vc.FileCsrInstanceFile = c
	return vc
}
