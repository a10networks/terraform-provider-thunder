package thunder

//Thunder resource SlbSSLCertRevoke

import (
	"context"
	"fmt"
	"util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSSLCertRevoke() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbSSLCertRevokeCreate,
		UpdateContext: resourceSlbSSLCertRevokeUpdate,
		ReadContext:   resourceSlbSSLCertRevokeRead,
		DeleteContext: resourceSlbSSLCertRevokeDelete,
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

func resourceSlbSSLCertRevokeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSSLCertRevoke (Inside resourceSlbSSLCertRevokeCreate) ")

		data := dataToSlbSSLCertRevoke(d)
		logger.Println("[INFO] received V from method data to SlbSSLCertRevoke --")
		d.SetId("1")
		err := go_thunder.PostSlbSSLCertRevoke(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbSSLCertRevokeRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbSSLCertRevokeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbSSLCertRevoke (Inside resourceSlbSSLCertRevokeRead)")

	if client.Host != "" {
		data, err := go_thunder.GetSlbSSLCertRevoke(client.Token, client.Host)
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

func resourceSlbSSLCertRevokeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSSLCertRevokeRead(ctx, d, meta)
}

func resourceSlbSSLCertRevokeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSSLCertRevokeRead(ctx, d, meta)
}

func dataToSlbSSLCertRevoke(d *schema.ResourceData) go_thunder.SSLCertRevoke {
	var vc go_thunder.SSLCertRevoke
	var c go_thunder.SSLCertRevokeInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable37, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable37
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
