package thunder

//Thunder resource FwUrpf

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwUrpf() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwUrpfCreate,
		UpdateContext: resourceFwUrpfUpdate,
		ReadContext:   resourceFwUrpfRead,
		DeleteContext: resourceFwUrpfDelete,
		Schema: map[string]*schema.Schema{
			"status": {
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

func resourceFwUrpfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwUrpf (Inside resourceFwUrpfCreate) ")

		data := dataToFwUrpf(d)
		logger.Println("[INFO] received formatted data from method data to FwUrpf --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwUrpf(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwUrpfRead(ctx, d, meta)

	}
	return diags
}

func resourceFwUrpfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwUrpf (Inside resourceFwUrpfRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwUrpf(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceFwUrpfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwUrpfRead(ctx, d, meta)
}

func resourceFwUrpfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwUrpfRead(ctx, d, meta)
}
func dataToFwUrpf(d *schema.ResourceData) go_thunder.FwUrpf {
	var vc go_thunder.FwUrpf
	var c go_thunder.FwUrpfInstance
	c.Status = d.Get("status").(string)

	vc.Status = c
	return vc
}
