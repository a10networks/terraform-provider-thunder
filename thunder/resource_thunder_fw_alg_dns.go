package thunder

//Thunder resource FwAlgDns

import (
	"context"
	"strconv"
	"util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgDns() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwAlgDnsCreate,
		UpdateContext: resourceFwAlgDnsUpdate,
		ReadContext:   resourceFwAlgDnsRead,
		DeleteContext: resourceFwAlgDnsDelete,
		Schema: map[string]*schema.Schema{
			"default_port_disable": {
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

func resourceFwAlgDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwAlgDns (Inside resourceFwAlgDnsCreate) ")

		data := dataToFwAlgDns(d)
		logger.Println("[INFO] received formatted data from method data to FwAlgDns --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwAlgDns(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwAlgDnsRead(ctx, d, meta)

	}
	return diags
}

func resourceFwAlgDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwAlgDns (Inside resourceFwAlgDnsRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwAlgDns(client.Token, client.Host)
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

func resourceFwAlgDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwAlgDnsRead(ctx, d, meta)
}

func resourceFwAlgDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwAlgDnsRead(ctx, d, meta)
}
func dataToFwAlgDns(d *schema.ResourceData) go_thunder.FwAlgDns {
	var vc go_thunder.FwAlgDns
	var c go_thunder.FwAlgDNSInstance
	c.DefaultPortDisable = d.Get("default_port_disable").(string)

	vc.DefaultPortDisable = c
	return vc
}
