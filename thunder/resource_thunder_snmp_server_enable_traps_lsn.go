package thunder

//Thunder resource SnmpServerEnableTrapsLsn

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsLsn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerEnableTrapsLsnCreate,
		UpdateContext: resourceSnmpServerEnableTrapsLsnUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsLsnRead,
		DeleteContext: resourceSnmpServerEnableTrapsLsnDelete,
		Schema: map[string]*schema.Schema{
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fixed_nat_port_mapping_file_change": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"per_ip_port_usage_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"total_port_usage_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_port_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_ipport_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"traffic_exceeded": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEnableTrapsLsnCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsLsn (Inside resourceSnmpServerEnableTrapsLsnCreate) ")

		data := dataToSnmpServerEnableTrapsLsn(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsLsn --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerEnableTrapsLsn(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerEnableTrapsLsnRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerEnableTrapsLsnRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerEnableTrapsLsn (Inside resourceSnmpServerEnableTrapsLsnRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsLsn(client.Token, client.Host)
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

func resourceSnmpServerEnableTrapsLsnUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsLsnRead(ctx, d, meta)
}

func resourceSnmpServerEnableTrapsLsnDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsLsnRead(ctx, d, meta)
}
func dataToSnmpServerEnableTrapsLsn(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsLsn {
	var vc go_thunder.SnmpServerEnableTrapsLsn
	var c go_thunder.SnmpServerEnableTrapsLsnInstance
	c.All = d.Get("all").(int)
	c.FixedNatPortMappingFileChange = d.Get("fixed_nat_port_mapping_file_change").(int)
	c.PerIPPortUsageThreshold = d.Get("per_ip_port_usage_threshold").(int)
	c.TotalPortUsageThreshold = d.Get("total_port_usage_threshold").(int)
	c.MaxPortThreshold = d.Get("max_port_threshold").(int)
	c.MaxIpportThreshold = d.Get("max_ipport_threshold").(int)
	c.TrafficExceeded = d.Get("traffic_exceeded").(int)

	vc.All = c
	return vc
}
