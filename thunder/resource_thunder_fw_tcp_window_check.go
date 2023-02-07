package thunder

//Thunder resource FwTcpWindowCheck

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTcpWindowCheck() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwTcpWindowCheckCreate,
		UpdateContext: resourceFwTcpWindowCheckUpdate,
		ReadContext:   resourceFwTcpWindowCheckRead,
		DeleteContext: resourceFwTcpWindowCheckDelete,
		Schema: map[string]*schema.Schema{
			"status": {
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwTcpWindowCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwTcpWindowCheck (Inside resourceFwTcpWindowCheckCreate) ")

		data := dataToFwTcpWindowCheck(d)
		logger.Println("[INFO] received formatted data from method data to FwTcpWindowCheck --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwTcpWindowCheck(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwTcpWindowCheckRead(ctx, d, meta)

	}
	return diags
}

func resourceFwTcpWindowCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwTcpWindowCheck (Inside resourceFwTcpWindowCheckRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwTcpWindowCheck(client.Token, client.Host)
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

func resourceFwTcpWindowCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwTcpWindowCheckRead(ctx, d, meta)
}

func resourceFwTcpWindowCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwTcpWindowCheckRead(ctx, d, meta)
}
func dataToFwTcpWindowCheck(d *schema.ResourceData) go_thunder.FwTcpWindowCheck {
	var vc go_thunder.FwTcpWindowCheck
	var c go_thunder.FwTCPWindowCheckInstance
	c.Status = d.Get("status").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.FwTCPWindowCheckSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.FwTCPWindowCheckSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Status = c
	return vc
}
