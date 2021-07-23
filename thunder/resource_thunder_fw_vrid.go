package thunder

//Thunder resource FwVrid

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwVrid() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwVridCreate,
		UpdateContext: resourceFwVridUpdate,
		ReadContext:   resourceFwVridRead,
		DeleteContext: resourceFwVridDelete,
		Schema: map[string]*schema.Schema{
			"vrid": {
				Type:        schema.TypeInt,
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

func resourceFwVridCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwVrid (Inside resourceFwVridCreate) ")

		data := dataToFwVrid(d)
		logger.Println("[INFO] received formatted data from method data to FwVrid --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwVrid(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwVridRead(ctx, d, meta)

	}
	return diags
}

func resourceFwVridRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwVrid (Inside resourceFwVridRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwVrid(client.Token, client.Host)
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

func resourceFwVridUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwVridRead(ctx, d, meta)
}

func resourceFwVridDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwVridRead(ctx, d, meta)
}
func dataToFwVrid(d *schema.ResourceData) go_thunder.FwVrid {
	var vc go_thunder.FwVrid
	var c go_thunder.FwVridInstance
	c.Vrid = d.Get("vrid").(int)

	vc.Vrid = c
	return vc
}
