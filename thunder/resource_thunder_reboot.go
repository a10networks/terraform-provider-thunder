package thunder

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceReboot() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRebootCreate,
		UpdateContext: resourceRebootUpdate,
		ReadContext:   resourceRebootRead,
		DeleteContext: resourceRebootDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reason_3": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cancel": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reason": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"reason_2": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"in": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"month_2": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"at": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"month": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"day_of_month_2": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"day_of_month": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"device": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceRebootCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating Reboot (Inside resourceRebootCreate)")

	if client.Host != "" {
		vc := dataToReboot(d)
		err := go_thunder.PostReboot(client.Token, vc, client.Host)
if err != nil {
			return diag.FromErr(err)
		}

		d.SetId("1")

		return resourceRebootRead(ctx, d, meta)
	}
	return diags
}

func resourceRebootRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading Reboot (Inside resourceRebootRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		vc, err := go_thunder.GetReboot(client.Token, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No Reboot found")
			//d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceRebootUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRebootRead(ctx, d, meta)
}

func resourceRebootDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRebootRead(ctx, d, meta)
}

//Utility method to instantiate Reboot Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToReboot(d *schema.ResourceData) go_thunder.Reboot {
	var vc go_thunder.Reboot

	var c go_thunder.RebootInstance

	c.All = d.Get("all").(int)

	vc.All = c

	return vc
}
