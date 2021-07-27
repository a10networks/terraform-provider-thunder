package thunder

//Thunder resource ActivePartition

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceActivePartition() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceActivePartitionCreate,
		UpdateContext: resourceActivePartitionUpdate,
		ReadContext:   resourceActivePartitionRead,
		DeleteContext: resourceActivePartitionDelete,
		Schema: map[string]*schema.Schema{
			"shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"curr_part_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "",
				ValidateFunc: validation.StringLenBetween(0, 14),
			},
		},
	}
}

func resourceActivePartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating ActivePartition (Inside resourceActivePartitionCreate) ")

		data := dataToActivePartition(d)
		logger.Println("[INFO] received formatted data from method data to ActivePartition --")
		d.SetId("1")
		err := go_thunder.PostActivePartition(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceActivePartitionRead(ctx, d, meta)

	}
	return diags
}

func resourceActivePartitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading ActivePartition (Inside resourceActivePartitionRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetActivePartition(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return diags
	}
	return nil
}

func resourceActivePartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceActivePartitionRead(ctx, d, meta)
}

func resourceActivePartitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceActivePartitionRead(ctx, d, meta)
}
func dataToActivePartition(d *schema.ResourceData) go_thunder.ActivePartition {
	var vc go_thunder.ActivePartition
	var c go_thunder.ActivePartitionInstance
	c.Shared = d.Get("shared").(int)
	c.CurrPartName = d.Get("curr_part_name").(string)

	vc.Shared = c
	return vc
}
