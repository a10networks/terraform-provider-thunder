package thunder

//Thunder resource SnmpServerManagementIndex

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceSnmpServerManagementIndex() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerManagementIndexCreate,
		UpdateContext: resourceSnmpServerManagementIndexUpdate,
		ReadContext:   resourceSnmpServerManagementIndexRead,
		DeleteContext: resourceSnmpServerManagementIndexDelete,
		Schema: map[string]*schema.Schema{
			"mgmt_index": {
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

func resourceSnmpServerManagementIndexCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerManagementIndex (Inside resourceSnmpServerManagementIndexCreate) ")

		data := dataToSnmpServerManagementIndex(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerManagementIndex --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerManagementIndex(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerManagementIndexRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerManagementIndexRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerManagementIndex (Inside resourceSnmpServerManagementIndexRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerManagementIndex(client.Token, client.Host)
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

func resourceSnmpServerManagementIndexUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerManagementIndexRead(ctx, d, meta)
}

func resourceSnmpServerManagementIndexDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerManagementIndexRead(ctx, d, meta)
}
func dataToSnmpServerManagementIndex(d *schema.ResourceData) go_thunder.SnmpServerManagementIndex {
	var vc go_thunder.SnmpServerManagementIndex
	var c go_thunder.SnmpServerManagementIndexInstance
	c.MgmtIndex = d.Get("mgmt_index").(int)

	vc.MgmtIndex = c
	return vc
}
