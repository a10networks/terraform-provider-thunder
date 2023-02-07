package thunder

//Thunder resource SnmpServerEngineID

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceSnmpServerEngineID() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerEngineIDCreate,
		UpdateContext: resourceSnmpServerEngineIDUpdate,
		ReadContext:   resourceSnmpServerEngineIDRead,
		DeleteContext: resourceSnmpServerEngineIDDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"eng_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEngineIDCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEngineID (Inside resourceSnmpServerEngineIDCreate) ")

		data := dataToSnmpServerEngineID(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEngineID --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerEngineID(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerEngineIDRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerEngineIDRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerEngineID (Inside resourceSnmpServerEngineIDRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEngineID(client.Token, client.Host)
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

func resourceSnmpServerEngineIDUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEngineIDRead(ctx, d, meta)
}

func resourceSnmpServerEngineIDDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEngineIDRead(ctx, d, meta)
}
func dataToSnmpServerEngineID(d *schema.ResourceData) go_thunder.SnmpServerEngineID {
	var vc go_thunder.SnmpServerEngineID
	var c go_thunder.SnmpServerEngineIDInstance
	c.EngID = d.Get("eng_id").(string)

	vc.UUID = c
	return vc
}
