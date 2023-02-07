package thunder

//Thunder resource WebCategoryStatistics

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryStatistics() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWebCategoryStatisticsCreate,
		UpdateContext: resourceWebCategoryStatisticsUpdate,
		ReadContext:   resourceWebCategoryStatisticsRead,
		DeleteContext: resourceWebCategoryStatisticsDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
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

func resourceWebCategoryStatisticsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating WebCategoryStatistics (Inside resourceWebCategoryStatisticsCreate) ")

		data := dataToWebCategoryStatistics(d)
		logger.Println("[INFO] received formatted data from method data to WebCategoryStatistics --")
		d.SetId("1")
		err := go_thunder.PostWebCategoryStatistics(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceWebCategoryStatisticsRead(ctx, d, meta)

	}
	return diags
}

func resourceWebCategoryStatisticsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading WebCategoryStatistics (Inside resourceWebCategoryStatisticsRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetWebCategoryStatistics(client.Token, client.Host)
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

func resourceWebCategoryStatisticsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceWebCategoryStatisticsRead(ctx, d, meta)
}

func resourceWebCategoryStatisticsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceWebCategoryStatisticsRead(ctx, d, meta)
}
func dataToWebCategoryStatistics(d *schema.ResourceData) go_thunder.WebCategoryStatistics {
	var vc go_thunder.WebCategoryStatistics
	var c go_thunder.WebCategoryStatisticsInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.WebCategoryStatisticsSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.WebCategoryStatisticsSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.UUID = c
	return vc
}
