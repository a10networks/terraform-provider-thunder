package thunder

//Thunder resource WriteMemory

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWriteMemory() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWriteMemoryCreate,
		UpdateContext: resourceWriteMemoryUpdate,
		ReadContext:   resourceWriteMemoryRead,
		DeleteContext: resourceWriteMemoryDelete,
		Schema: map[string]*schema.Schema{
			"profile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"specified_partition": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"destination": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"partition": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceWriteMemoryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating WriteMemory (Inside resourceWriteMemoryCreate) ")

		data := dataToWriteMemory(d)
		logger.Println("[INFO] received formatted data from method data to WriteMemory --")
		d.SetId("1")
		err := go_thunder.PostWriteMemory(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceWriteMemoryRead(ctx, d, meta)

	}
	return diags
}

func resourceWriteMemoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading WriteMemory (Inside resourceWriteMemoryRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetWriteMemory(client.Token, client.Host)
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

func resourceWriteMemoryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceWriteMemoryRead(ctx, d, meta)
}

func resourceWriteMemoryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceWriteMemoryRead(ctx, d, meta)
}
func dataToWriteMemory(d *schema.ResourceData) go_thunder.WriteMemory {
	var vc go_thunder.WriteMemory
	var c go_thunder.WriteMemoryInstance
	c.Profile = d.Get("profile").(string)
	c.SpecifiedPartition = d.Get("specified_partition").(string)
	c.Destination = d.Get("destination").(string)
	c.Partition = d.Get("partition").(string)

	vc.Profile = c
	return vc
}
