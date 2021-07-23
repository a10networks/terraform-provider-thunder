package thunder

//Thunder resource Hostname

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHostname() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHostnameCreate,
		UpdateContext: resourceHostnameUpdate,
		ReadContext:   resourceHostnameRead,
		DeleteContext: resourceHostnameDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceHostnameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating Hostname (Inside resourceHostnameCreate) ")

		data := dataToHostname(d)
		logger.Println("[INFO] received formatted data from method data to Hostname --")
		d.SetId("1")
		err := go_thunder.PostHostname(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceHostnameRead(ctx, d, meta)

	}
	return diags
}

func resourceHostnameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading Hostname (Inside resourceHostnameRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read hostname")
		data, err := go_thunder.GetHostname(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found hostname ")
			return nil
		}
		return diags
	}
	return nil
}

func resourceHostnameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceHostnameRead(ctx, d, meta)
}

func resourceHostnameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceHostnameRead(ctx, d, meta)
}
func dataToHostname(d *schema.ResourceData) go_thunder.Hostname {
	var vc go_thunder.Hostname
	var c go_thunder.HostnameInstance
	c.Value = d.Get("value").(string)

	vc.UUID = c
	return vc
}
