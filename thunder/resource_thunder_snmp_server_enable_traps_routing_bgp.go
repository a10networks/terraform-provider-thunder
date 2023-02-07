package thunder

//Thunder resource SnmpServerEnableTrapsRoutingBgp

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsRoutingBgp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerEnableTrapsRoutingBgpCreate,
		UpdateContext: resourceSnmpServerEnableTrapsRoutingBgpUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsRoutingBgpRead,
		DeleteContext: resourceSnmpServerEnableTrapsRoutingBgpDelete,
		Schema: map[string]*schema.Schema{
			"bgp_established_notification": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"bgp_backward_trans_notification": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEnableTrapsRoutingBgpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsRoutingBgp (Inside resourceSnmpServerEnableTrapsRoutingBgpCreate) ")

		data := dataToSnmpServerEnableTrapsRoutingBgp(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsRoutingBgp --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerEnableTrapsRoutingBgp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerEnableTrapsRoutingBgpRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerEnableTrapsRoutingBgpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerEnableTrapsRoutingBgp (Inside resourceSnmpServerEnableTrapsRoutingBgpRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsRoutingBgp(client.Token, client.Host)
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

func resourceSnmpServerEnableTrapsRoutingBgpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsRoutingBgpRead(ctx, d, meta)
}

func resourceSnmpServerEnableTrapsRoutingBgpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsRoutingBgpRead(ctx, d, meta)
}
func dataToSnmpServerEnableTrapsRoutingBgp(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsRoutingBgp {
	var vc go_thunder.SnmpServerEnableTrapsRoutingBgp
	var c go_thunder.SnmpServerEnableTrapsRoutingBgpInstance
	c.BgpEstablishedNotification = d.Get("bgp_established_notification").(int)
	c.BgpBackwardTransNotification = d.Get("bgp_backward_trans_notification").(int)

	vc.BgpEstablishedNotification = c
	return vc
}
