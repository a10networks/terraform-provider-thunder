package thunder

//Thunder resource SnmpServerEnableTrapsNetwork

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsNetwork() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerEnableTrapsNetworkCreate,
		UpdateContext: resourceSnmpServerEnableTrapsNetworkUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsNetworkRead,
		DeleteContext: resourceSnmpServerEnableTrapsNetworkDelete,
		Schema: map[string]*schema.Schema{
			"trunk_port_threshold": {
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

func resourceSnmpServerEnableTrapsNetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsNetwork (Inside resourceSnmpServerEnableTrapsNetworkCreate) ")

		data := dataToSnmpServerEnableTrapsNetwork(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsNetwork --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerEnableTrapsNetwork(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerEnableTrapsNetworkRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerEnableTrapsNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerEnableTrapsNetwork (Inside resourceSnmpServerEnableTrapsNetworkRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsNetwork(client.Token, client.Host)
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

func resourceSnmpServerEnableTrapsNetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsNetworkRead(ctx, d, meta)
}

func resourceSnmpServerEnableTrapsNetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsNetworkRead(ctx, d, meta)
}
func dataToSnmpServerEnableTrapsNetwork(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsNetwork {
	var vc go_thunder.SnmpServerEnableTrapsNetwork
	var c go_thunder.SnmpServerEnableTrapsNetworkInstance
	c.TrunkPortThreshold = d.Get("trunk_port_threshold").(int)

	vc.TrunkPortThreshold = c
	return vc
}
