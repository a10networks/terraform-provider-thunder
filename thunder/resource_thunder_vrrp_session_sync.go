package thunder

//Thunder resource Vrrp common

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpSessionSync() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVrrpSessionSyncCreate,
		UpdateContext: resourceVrrpSessionSyncUpdate,
		ReadContext:   resourceVrrpSessionSyncRead,
		DeleteContext: resourceVrrpSessionSyncDelete,

		Schema: map[string]*schema.Schema{
			"action": {
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

func resourceVrrpSessionSyncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating VrrpSessionSync (Inside resourceVrrpSessionSyncCreate)")

	if client.Host != "" {
		vc := dataToVrrpSessionSync(d)
		err := go_thunder.PostVrrpSessionSync(client.Token, vc, client.Host)
if err != nil {
			return diag.FromErr(err)
		}

		d.SetId("1")

		return resourceVrrpSessionSyncRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpSessionSyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading VrrpSessionSync (Inside resourceVrrpSessionSyncRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		vc, err := go_thunder.GetVrrpSessionSync(client.Token, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No VrrpSessionSync found")
			//d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceVrrpSessionSyncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceVrrpSessionSyncRead(ctx, d, meta)
}

func resourceVrrpSessionSyncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceVrrpSessionSyncRead(ctx, d, meta)
}

//Utility method to instantiate DnsPrimary Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToVrrpSessionSync(d *schema.ResourceData) go_thunder.SessionSync {
	var vc go_thunder.SessionSync

	var c go_thunder.SessionSyncInstance

	c.Action = d.Get("action").(string)

	vc.Action = c

	return vc
}
