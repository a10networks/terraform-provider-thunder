package thunder

//Thunder resource SnmpServerLocation

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerLocation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerLocationCreate,
		UpdateContext: resourceSnmpServerLocationUpdate,
		ReadContext:   resourceSnmpServerLocationRead,
		DeleteContext: resourceSnmpServerLocationDelete,
		Schema: map[string]*schema.Schema{
			"loc": {
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

func resourceSnmpServerLocationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerLocation (Inside resourceSnmpServerLocationCreate) ")

		data := dataToSnmpServerLocation(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerLocation --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerLocation(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerLocationRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerLocationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerLocation (Inside resourceSnmpServerLocationRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerLocation(client.Token, client.Host)
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

func resourceSnmpServerLocationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerLocationRead(ctx, d, meta)
}

func resourceSnmpServerLocationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerLocationRead(ctx, d, meta)
}
func dataToSnmpServerLocation(d *schema.ResourceData) go_thunder.SnmpServerLocation {
	var vc go_thunder.SnmpServerLocation
	var c go_thunder.SnmpServerLocationInstance
	c.Loc = d.Get("loc").(string)

	vc.Loc = c
	return vc
}
