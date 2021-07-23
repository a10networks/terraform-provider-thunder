package thunder

//Thunder resource SnmpServerEnableTrapsGslb

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsGslb() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerEnableTrapsGslbCreate,
		UpdateContext: resourceSnmpServerEnableTrapsGslbUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsGslbRead,
		DeleteContext: resourceSnmpServerEnableTrapsGslbDelete,
		Schema: map[string]*schema.Schema{
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"group": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"zone": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"site": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEnableTrapsGslbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsGslb (Inside resourceSnmpServerEnableTrapsGslbCreate) ")

		data := dataToSnmpServerEnableTrapsGslb(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsGslb --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerEnableTrapsGslb(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerEnableTrapsGslbRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerEnableTrapsGslbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerEnableTrapsGslb (Inside resourceSnmpServerEnableTrapsGslbRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsGslb(client.Token, client.Host)
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

func resourceSnmpServerEnableTrapsGslbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsGslbRead(ctx, d, meta)
}

func resourceSnmpServerEnableTrapsGslbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsGslbRead(ctx, d, meta)
}
func dataToSnmpServerEnableTrapsGslb(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsGslb {
	var vc go_thunder.SnmpServerEnableTrapsGslb
	var c go_thunder.SnmpServerEnableTrapsGslbInstance
	c.All = d.Get("all").(int)
	c.Group = d.Get("group").(int)
	c.Zone = d.Get("zone").(int)
	c.Site = d.Get("site").(int)
	c.ServiceIP = d.Get("service_ip").(int)

	vc.All = c
	return vc
}
