package thunder

//Thunder resource FwTcpRstCloseImmediate

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTcpRstCloseImmediate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwTcpRstCloseImmediateCreate,
		UpdateContext: resourceFwTcpRstCloseImmediateUpdate,
		ReadContext:   resourceFwTcpRstCloseImmediateRead,
		DeleteContext: resourceFwTcpRstCloseImmediateDelete,
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

func resourceFwTcpRstCloseImmediateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwTcpRstCloseImmediate (Inside resourceFwTcpRstCloseImmediateCreate) ")

		data := dataToFwTcpRstCloseImmediate(d)
		logger.Println("[INFO] received formatted data from method data to FwTcpRstCloseImmediate --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwTcpRstCloseImmediate(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwTcpRstCloseImmediateRead(ctx, d, meta)

	}
	return diags
}

func resourceFwTcpRstCloseImmediateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwTcpRstCloseImmediate (Inside resourceFwTcpRstCloseImmediateRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwTcpRstCloseImmediate(client.Token, client.Host)
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

func resourceFwTcpRstCloseImmediateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwTcpRstCloseImmediateRead(ctx, d, meta)
}

func resourceFwTcpRstCloseImmediateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwTcpRstCloseImmediateRead(ctx, d, meta)
}
func dataToFwTcpRstCloseImmediate(d *schema.ResourceData) go_thunder.FwTcpRstCloseImmediate {
	var vc go_thunder.FwTcpRstCloseImmediate
	var c go_thunder.FwTCPRstCloseImmediateInstance
	c.Status = d.Get("status").(string)

	vc.Status = c
	return vc
}
