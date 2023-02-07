package thunder

//Thunder resource SnmpServerSlbDataCacheTimeout

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerSlbDataCacheTimeout() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerSlbDataCacheTimeoutCreate,
		UpdateContext: resourceSnmpServerSlbDataCacheTimeoutUpdate,
		ReadContext:   resourceSnmpServerSlbDataCacheTimeoutRead,
		DeleteContext: resourceSnmpServerSlbDataCacheTimeoutDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"slblimit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerSlbDataCacheTimeoutCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerSlbDataCacheTimeout (Inside resourceSnmpServerSlbDataCacheTimeoutCreate) ")

		data := dataToSnmpServerSlbDataCacheTimeout(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerSlbDataCacheTimeout --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerSlbDataCacheTimeout(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerSlbDataCacheTimeoutRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerSlbDataCacheTimeoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerSlbDataCacheTimeout (Inside resourceSnmpServerSlbDataCacheTimeoutRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerSlbDataCacheTimeout(client.Token, client.Host)
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

func resourceSnmpServerSlbDataCacheTimeoutUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerSlbDataCacheTimeoutRead(ctx, d, meta)
}

func resourceSnmpServerSlbDataCacheTimeoutDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerSlbDataCacheTimeoutRead(ctx, d, meta)
}
func dataToSnmpServerSlbDataCacheTimeout(d *schema.ResourceData) go_thunder.SnmpServerSlbDataCacheTimeout {
	var vc go_thunder.SnmpServerSlbDataCacheTimeout
	var c go_thunder.SnmpServerSlbDataCacheTimeoutInstance
	c.Slblimit = d.Get("slblimit").(int)

	vc.UUID = c
	return vc
}
