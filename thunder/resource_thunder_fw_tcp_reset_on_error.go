package thunder

//Thunder resource FwTcpResetOnError

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTcpResetOnError() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwTcpResetOnErrorCreate,
		UpdateContext: resourceFwTcpResetOnErrorUpdate,
		ReadContext:   resourceFwTcpResetOnErrorRead,
		DeleteContext: resourceFwTcpResetOnErrorDelete,
		Schema: map[string]*schema.Schema{
			"enable": {
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

func resourceFwTcpResetOnErrorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwTcpResetOnError (Inside resourceFwTcpResetOnErrorCreate) ")

		data := dataToFwTcpResetOnError(d)
		logger.Println("[INFO] received formatted data from method data to FwTcpResetOnError --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwTcpResetOnError(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwTcpResetOnErrorRead(ctx, d, meta)

	}
	return diags
}

func resourceFwTcpResetOnErrorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwTcpResetOnError (Inside resourceFwTcpResetOnErrorRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwTcpResetOnError(client.Token, client.Host)
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

func resourceFwTcpResetOnErrorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwTcpResetOnErrorRead(ctx, d, meta)
}

func resourceFwTcpResetOnErrorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwTcpResetOnErrorRead(ctx, d, meta)
}
func dataToFwTcpResetOnError(d *schema.ResourceData) go_thunder.FwTcpResetOnError {
	var vc go_thunder.FwTcpResetOnError
	var c go_thunder.FwTcpResetOnErrorInstance
	c.Enable = d.Get("enable").(int)

	vc.Enable = c
	return vc
}
