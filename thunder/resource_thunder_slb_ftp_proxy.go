package thunder

// Thunder resource Slb FTPProxy

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFTPProxy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbFTPProxyCreate,
		UpdateContext: resourceSlbFTPProxyUpdate,
		ReadContext:   resourceSlbFTPProxyRead,
		DeleteContext: resourceSlbFTPProxyDelete,

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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceSlbFTPProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating ftp-proxy (Inside resourceSlbFTPProxyCreate)")

	if client.Host != "" {
		vc := dataToSlbFTPProxy(d)
		d.SetId("1")
		err := go_thunder.PostSlbFTPProxy(client.Token, vc, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbFTPProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbFTPProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading ftp-proxy (Inside resourceSlbFTPProxyRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbFTPProxy(client.Token, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No ftp-proxy found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbFTPProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbFTPProxyRead(ctx, d, meta)
}

func resourceSlbFTPProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbFTPProxyRead(ctx, d, meta)
}

//Utility method to instantiate FTPProxy Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbFTPProxy(d *schema.ResourceData) go_thunder.FTPProxy {
	var vc go_thunder.FTPProxy

	var c go_thunder.FTPProxyInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable22, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable22
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
