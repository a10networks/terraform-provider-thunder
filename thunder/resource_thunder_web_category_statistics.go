package thunder

//Thunder resource WebCategoryStatistics

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceWebCategoryStatistics() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebCategoryStatisticsCreate,
		Update: resourceWebCategoryStatisticsUpdate,
		Read:   resourceWebCategoryStatisticsRead,
		Delete: resourceWebCategoryStatisticsDelete,
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

func resourceWebCategoryStatisticsCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating WebCategoryStatistics (Inside resourceWebCategoryStatisticsCreate) ")

		data := dataToWebCategoryStatistics(d)
		logger.Println("[INFO] received formatted data from method data to WebCategoryStatistics --")
		d.SetId("1")
		go_thunder.PostWebCategoryStatistics(client.Token, data, client.Host)

		return resourceWebCategoryStatisticsRead(d, meta)

	}
	return nil
}

func resourceWebCategoryStatisticsRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading WebCategoryStatistics (Inside resourceWebCategoryStatisticsRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetWebCategoryStatistics(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceWebCategoryStatisticsUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceWebCategoryStatisticsRead(d, meta)
}

func resourceWebCategoryStatisticsDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceWebCategoryStatisticsRead(d, meta)
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
