package thunder

//Thunder resource FwLogging

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwLogging() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwLoggingCreate,
		UpdateContext: resourceFwLoggingUpdate,
		ReadContext:   resourceFwLoggingRead,
		DeleteContext: resourceFwLoggingDelete,
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
			"name": {
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

func resourceFwLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwLogging (Inside resourceFwLoggingCreate) ")

		data := dataToFwLogging(d)
		logger.Println("[INFO] received formatted data from method data to FwLogging --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwLogging(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwLoggingRead(ctx, d, meta)

	}
	return diags
}

func resourceFwLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwLogging (Inside resourceFwLoggingRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwLogging(client.Token, client.Host)
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

func resourceFwLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwLoggingRead(ctx, d, meta)
}

func resourceFwLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwLoggingRead(ctx, d, meta)
}
func dataToFwLogging(d *schema.ResourceData) go_thunder.FwLogging {
	var vc go_thunder.FwLogging
	var c go_thunder.FwLoggingInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.FwLoggingSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.FwLoggingSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	c.Name = d.Get("name").(string)

	vc.SamplingEnable = c
	return vc
}
