package thunder

//Thunder resource SlbMssql

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbMssql() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbMssqlCreate,
		UpdateContext: resourceSlbMssqlUpdate,
		ReadContext:   resourceSlbMssqlRead,
		DeleteContext: resourceSlbMssqlDelete,
		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbMssqlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbMssql (Inside resourceSlbMssqlCreate) ")

		data := dataToSlbMssql(d)
		logger.Println("[INFO] received V from method data to SlbMssql --")
		d.SetId("1")
		err := go_thunder.PostSlbMssql(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbMssqlRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbMssqlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbMssql (Inside resourceSlbMssqlRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbMssql(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {

			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbMssqlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbMssqlRead(ctx, d, meta)
}

func resourceSlbMssqlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbMssqlRead(ctx, d, meta)
}

func dataToSlbMssql(d *schema.ResourceData) go_thunder.Mssql {
	var vc go_thunder.Mssql
	var c go_thunder.MssqlInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable45, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable45
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
