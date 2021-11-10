package thunder

//Thunder resource FileSslKey

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFileSslKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFileSslKeyCreate,
		UpdateContext: resourceFileSslKeyUpdate,
		ReadContext:   resourceFileSslKeyRead,
		DeleteContext: resourceFileSslKeyDelete,
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFileSslKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating FileSslKey (Inside resourceFileSslKeyCreate) ")

		data := dataToFileSslKey(d)
		FileHandleLocal := d.Get("file_local_path").(string)
		action_import := d.Get("action").(string)
		logger.Println("action_import = " + action_import)
		logger.Println("[INFO] received formatted data from method data to FileSslKey --")
		d.SetId("1")

		if action_import == "import" {
			logger.Println(action_import + " for ssl_cert")
			err := go_thunder.PostFileSslKeyImport(client.Token, data, client.Host, FileHandleLocal)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			logger.Println(action_import + " for ssl_cert")
			err := go_thunder.PostFileSslKey(client.Token, data, client.Host)
			if err != nil {
				return diag.FromErr(err)
			}
		}

		return resourceFileSslKeyRead(ctx, d, meta)

	}
	return diags
}

func resourceFileSslKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	//client := meta.(Thunder)
	logger.Println("[INFO] Reading FileSslKey (Inside resourceFileSslKeyRead)")

	var diags diag.Diagnostics
	// if client.Host != "" {
	// 	logger.Println("[INFO] Fetching service Read")
	// 	data, err := go_thunder.GetFileSslKey(client.Token, client.Host)
	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// 	if data == nil {
	// 		logger.Println("[INFO] No data found ")
	// 		return nil
	// 	}
	// 	return diags
	// }
	return diags
}

func resourceFileSslKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFileSslKeyRead(ctx, d, meta)
}

func resourceFileSslKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFileSslKeyRead(ctx, d, meta)
}
func dataToFileSslKey(d *schema.ResourceData) go_thunder.FileSslKey {
	var vc go_thunder.FileSslKey
	var c go_thunder.FileSslKeyInstance
	c.FileSslKeyInstanceFile = d.Get("file").(string)
	c.FileSslKeyInstanceSize = d.Get("size").(int)
	c.FileSslKeyInstanceFileHandle = d.Get("file_handle").(string)
	c.FileSslKeyInstanceSecured = d.Get("secured").(int)
	c.FileSslKeyInstanceAction = d.Get("action").(string)
	c.FileSslKeyInstanceDstFile = d.Get("dst_file").(string)

	vc.FileSslKeyInstanceFile = c
	return vc
}
