package thunder

//Thunder resource FwAlgIcmp

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgIcmp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwAlgIcmpCreate,
		UpdateContext: resourceFwAlgIcmpUpdate,
		ReadContext:   resourceFwAlgIcmpRead,
		DeleteContext: resourceFwAlgIcmpDelete,
		Schema: map[string]*schema.Schema{
			"disable": {
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

func resourceFwAlgIcmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwAlgIcmp (Inside resourceFwAlgIcmpCreate) ")

		data := dataToFwAlgIcmp(d)
		logger.Println("[INFO] received formatted data from method data to FwAlgIcmp --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwAlgIcmp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwAlgIcmpRead(ctx, d, meta)

	}
	return diags
}

func resourceFwAlgIcmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwAlgIcmp (Inside resourceFwAlgIcmpRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwAlgIcmp(client.Token, client.Host)
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

func resourceFwAlgIcmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwAlgIcmpRead(ctx, d, meta)
}

func resourceFwAlgIcmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwAlgIcmpRead(ctx, d, meta)
}
func dataToFwAlgIcmp(d *schema.ResourceData) go_thunder.FwAlgIcmp {
	var vc go_thunder.FwAlgIcmp
	var c go_thunder.FwAlgIcmpInstance
	c.Disable = d.Get("disable").(string)

	vc.Disable = c
	return vc
}
