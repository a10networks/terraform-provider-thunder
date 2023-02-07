package thunder

// Thunder resource Slb FTPData

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFTPData() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbFTPDataCreate,
		UpdateContext: resourceSlbFTPDataUpdate,
		ReadContext:   resourceSlbFTPDataRead,
		DeleteContext: resourceSlbFTPDataDelete,

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

func resourceSlbFTPDataCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating FTPData (Inside resourceSlbFTPDataCreate)")

	if client.Host != "" {
		vc := dataToSlbFTPData(d)
		d.SetId("1")
		err := go_thunder.PostSlbFTPData(client.Token, vc, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbFTPDataRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbFTPDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading FTPData (Inside resourceSlbFTPDataRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbFTPData(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No FTPData found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbFTPDataUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbFTPDataRead(ctx, d, meta)
}

func resourceSlbFTPDataDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbFTPDataRead(ctx, d, meta)
}

//Utility method to instantiate FTPData Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbFTPData(d *schema.ResourceData) go_thunder.FTPData {
	var vc go_thunder.FTPData

	var c go_thunder.FTPDataInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable21, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable21
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
