package thunder

//Thunder resource SlbSmtp

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSmtp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbSmtpCreate,
		UpdateContext: resourceSlbSmtpUpdate,
		ReadContext:   resourceSlbSmtpRead,
		DeleteContext: resourceSlbSmtpDelete,
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

func resourceSlbSmtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSmtp (Inside resourceSlbSmtpCreate) ")

		data := dataToSlbSmtp(d)
		logger.Println("[INFO] received formatted data from method data to SlbSmtp --")
		d.SetId("1")
		err := go_thunder.PostSlbSmtp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbSmtpRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbSmtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbSmtp (Inside resourceSlbSmtpRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbSmtp(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {

			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbSmtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSmtpRead(ctx, d, meta)
}

func resourceSlbSmtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSmtpRead(ctx, d, meta)
}
func dataToSlbSmtp(d *schema.ResourceData) go_thunder.SlbSmtp {
	var vc go_thunder.SlbSmtp
	var c go_thunder.SlbSmtpInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable12, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable12
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		obj1.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
