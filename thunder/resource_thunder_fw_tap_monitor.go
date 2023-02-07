package thunder

//Thunder resource FwTapMonitor

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTapMonitor() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwTapMonitorCreate,
		UpdateContext: resourceFwTapMonitorUpdate,
		ReadContext:   resourceFwTapMonitorRead,
		DeleteContext: resourceFwTapMonitorDelete,
		Schema: map[string]*schema.Schema{
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"tap_port_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tap_eth": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"tap_vlan": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwTapMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwTapMonitor (Inside resourceFwTapMonitorCreate) ")

		data := dataToFwTapMonitor(d)
		logger.Println("[INFO] received formatted data from method data to FwTapMonitor --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwTapMonitor(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwTapMonitorRead(ctx, d, meta)

	}
	return diags
}

func resourceFwTapMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwTapMonitor (Inside resourceFwTapMonitorRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwTapMonitor(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceFwTapMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwTapMonitorRead(ctx, d, meta)
}

func resourceFwTapMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwTapMonitorRead(ctx, d, meta)
}
func dataToFwTapMonitor(d *schema.ResourceData) go_thunder.FwTapMonitor {
	var vc go_thunder.FwTapMonitor
	var c go_thunder.FwTapMonitorInstance
	c.Status = d.Get("status").(string)

	TapPortCfgCount := d.Get("tap_port_cfg.#").(int)
	c.TapEth = make([]go_thunder.FwTapMonitorPortCfg, 0, TapPortCfgCount)

	for i := 0; i < TapPortCfgCount; i++ {
		var obj1 go_thunder.FwTapMonitorPortCfg
		prefix := fmt.Sprintf("tap_port_cfg.%d.", i)
		obj1.TapEth = d.Get(prefix + "tap_eth").(int)
		obj1.TapVlan = d.Get(prefix + "tap_vlan").(int)
		c.TapEth = append(c.TapEth, obj1)
	}

	vc.Status = c
	return vc
}
