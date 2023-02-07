package thunder

//Thunder resource SnmpServerContact

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerContact() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerContactCreate,
		UpdateContext: resourceSnmpServerContactUpdate,
		ReadContext:   resourceSnmpServerContactRead,
		DeleteContext: resourceSnmpServerContactDelete,
		Schema: map[string]*schema.Schema{
			"contact_name": {
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

func resourceSnmpServerContactCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerContact (Inside resourceSnmpServerContactCreate) ")

		data := dataToSnmpServerContact(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerContact --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerContact(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerContactRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerContactRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerContact (Inside resourceSnmpServerContactRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerContact(client.Token, client.Host)
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

func resourceSnmpServerContactUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerContactRead(ctx, d, meta)
}

func resourceSnmpServerContactDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerContactRead(ctx, d, meta)
}
func dataToSnmpServerContact(d *schema.ResourceData) go_thunder.SnmpServerContact {
	var vc go_thunder.SnmpServerContact
	var c go_thunder.SnmpServerContactInstance
	c.ContactName = d.Get("contact_name").(string)

	vc.ContactName = c
	return vc
}
