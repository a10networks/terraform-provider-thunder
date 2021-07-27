package thunder

//Thunder resource FwAlgFtp

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgFtp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwAlgFtpCreate,
		UpdateContext: resourceFwAlgFtpUpdate,
		ReadContext:   resourceFwAlgFtpRead,
		DeleteContext: resourceFwAlgFtpDelete,
		Schema: map[string]*schema.Schema{
			"default_port_disable": {
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

func resourceFwAlgFtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwAlgFtp (Inside resourceFwAlgFtpCreate) ")

		data := dataToFwAlgFtp(d)
		logger.Println("[INFO] received formatted data from method data to FwAlgFtp --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwAlgFtp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwAlgFtpRead(ctx, d, meta)

	}
	return diags
}

func resourceFwAlgFtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwAlgFtp (Inside resourceFwAlgFtpRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwAlgFtp(client.Token, client.Host)
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

func resourceFwAlgFtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwAlgFtpRead(ctx, d, meta)
}

func resourceFwAlgFtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwAlgFtpRead(ctx, d, meta)
}
func dataToFwAlgFtp(d *schema.ResourceData) go_thunder.FwAlgFtp {
	var vc go_thunder.FwAlgFtp
	var c go_thunder.FwAlgFtpInstance
	c.DefaultPortDisable = d.Get("default_port_disable").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.FwAlgFtpSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.FwAlgFtpSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.DefaultPortDisable = c
	return vc
}
