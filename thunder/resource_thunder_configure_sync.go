package thunder

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigureSync() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConfigureSyncCreate,
		UpdateContext: resourceConfigureSyncUpdate,
		ReadContext:   resourceConfigureSyncRead,
		DeleteContext: resourceConfigureSyncDelete,

		Schema: map[string]*schema.Schema{
			"shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"partition_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"usr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"private_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auto_authentication": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"all_partitions": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pwd_enc": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pwd": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceConfigureSyncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating ConfigureSync (Inside resourceConfigureSyncCreate)")

	if client.Host != "" {
		vc := dataToConfigureSync(d)
		err := go_thunder.PostConfigureSync(client.Token, vc, client.Host)
if err != nil {
			return diag.FromErr(err)
		}

		d.SetId("1")

		return resourceConfigureSyncRead(ctx, d, meta)
	}
	return diags
}

func resourceConfigureSyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading ConfigureSync (Inside resourceConfigureSyncRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		vc, err := go_thunder.GetConfigureSync(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No ConfigureSync found")

			return nil
		}

		return diags
	}
	return nil
}

func resourceConfigureSyncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceConfigureSyncRead(ctx, d, meta)
}

func resourceConfigureSyncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceConfigureSyncRead(ctx, d, meta)
}

//Utility method to instantiate Glm Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToConfigureSync(d *schema.ResourceData) go_thunder.Sync {
	var vc go_thunder.Sync

	var c go_thunder.SyncInstance

	c.AllPartitions = d.Get("all_partitions").(int)
	c.PrivateKey = d.Get("private_key").(string)
	c.PartitionName = d.Get("partition_name").(string)
	c.Pwd = d.Get("pwd").(string)
	c.AutoAuthentication = d.Get("auto_authentication").(int)
	c.Address = d.Get("address").(string)
	c.Shared = d.Get("shared").(int)
	c.Type = d.Get("type").(string)
	c.PwdEnc = d.Get("pwd_enc").(string)
	c.Usr = d.Get("usr").(string)

	vc.AllPartitions = c

	return vc
}
