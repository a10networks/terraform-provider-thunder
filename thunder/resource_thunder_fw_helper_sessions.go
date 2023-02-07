package thunder

//Thunder resource FwHelperSessions

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwHelperSessions() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwHelperSessionsCreate,
		UpdateContext: resourceFwHelperSessionsUpdate,
		ReadContext:   resourceFwHelperSessionsRead,
		DeleteContext: resourceFwHelperSessionsDelete,
		Schema: map[string]*schema.Schema{
			"idle_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mode": {
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

func resourceFwHelperSessionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwHelperSessions (Inside resourceFwHelperSessionsCreate) ")

		data := dataToFwHelperSessions(d)
		logger.Println("[INFO] received formatted data from method data to FwHelperSessions --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwHelperSessions(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwHelperSessionsRead(ctx, d, meta)

	}
	return diags
}

func resourceFwHelperSessionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwHelperSessions (Inside resourceFwHelperSessionsRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwHelperSessions(client.Token, client.Host)
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

func resourceFwHelperSessionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwHelperSessionsRead(ctx, d, meta)
}

func resourceFwHelperSessionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwHelperSessionsRead(ctx, d, meta)
}
func dataToFwHelperSessions(d *schema.ResourceData) go_thunder.FwHelperSessions {
	var vc go_thunder.FwHelperSessions
	var c go_thunder.FwHelperSessionsInstance
	c.IdleTimeout = d.Get("idle_timeout").(int)
	c.Limit = d.Get("limit").(int)
	c.Mode = d.Get("mode").(string)

	vc.IdleTimeout = c
	return vc
}
