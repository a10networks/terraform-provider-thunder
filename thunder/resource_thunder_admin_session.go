package thunder

//Thunder resource AdminSession

import (
	"context"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceAdminSession() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAdminSessionCreate,
		UpdateContext: resourceAdminSessionUpdate,
		ReadContext:   resourceAdminSessionRead,
		DeleteContext: resourceAdminSessionDelete,
		Schema: map[string]*schema.Schema{
			"clear": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sid": {
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

func resourceAdminSessionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating AdminSession (Inside resourceAdminSessionCreate) ")

		data := dataToAdminSession(d)
		logger.Println("[INFO] received formatted data from method data to AdminSession --")
		d.SetId("1")
		err := go_thunder.PostAdminSession(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceAdminSessionRead(ctx, d, meta)

	}
	return diags
}

func resourceAdminSessionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading AdminSession (Inside resourceAdminSessionRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetAdminSession(client.Token, client.Host)
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

func resourceAdminSessionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceAdminSessionRead(ctx, d, meta)
}

func resourceAdminSessionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceAdminSessionRead(ctx, d, meta)
}
func dataToAdminSession(d *schema.ResourceData) go_thunder.AdminSession {
	var vc go_thunder.AdminSession
	var c go_thunder.AdminSessionInstance
	c.AdminSessionInstanceClear = d.Get("clear").(int)
	c.AdminSessionInstanceAll = d.Get("all").(int)
	c.AdminSessionInstanceSid = d.Get("sid").(int)

	vc.AdminSessionInstanceClear = c
	return vc
}
