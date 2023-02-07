package thunder

// Thunder resource Slb FTPCtl

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFTPCtl() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbFTPCtlCreate,
		UpdateContext: resourceSlbFTPCtlUpdate,
		ReadContext:   resourceSlbFTPCtlRead,
		DeleteContext: resourceSlbFTPCtlDelete,

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

func resourceSlbFTPCtlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating FTPCtl (Inside resourceSlbFTPCtlCreate)")

	if client.Host != "" {
		vc := dataToSlbFTPCtl(d)
		d.SetId("1")
		err := go_thunder.PostSlbFTPCtl(client.Token, vc, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbFTPCtlRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbFTPCtlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading FTPCtl (Inside resourceSlbFTPCtlRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbFTPCtl(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No FTPCtl found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbFTPCtlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbFTPCtlRead(ctx, d, meta)
}

func resourceSlbFTPCtlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbFTPCtlRead(ctx, d, meta)
}

//Utility method to instantiate FTPCtl Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbFTPCtl(d *schema.ResourceData) go_thunder.FTPCtl {
	var vc go_thunder.FTPCtl

	var c go_thunder.FTPCtlInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable20, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable20
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
