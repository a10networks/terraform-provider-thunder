package thunder

// Thunder resource Slb HwCompress

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbHwCompress() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbHwCompressCreate,
		Update: resourceSlbHwCompressUpdate,
		Read:   resourceSlbHwCompressRead,
		Delete: resourceSlbHwCompressDelete,

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

func resourceSlbHwCompressCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	logger.Println("[INFO] Creating hw-compress (Inside resourceSlbHwCompressCreate)")

	if client.Host != "" {
		vc := dataToSlbHwCompress(d)
		d.SetId("1")
		go_thunder.PostSlbHwCompress(client.Token, vc, client.Host)

		return resourceSlbHwCompressRead(d, meta)
	}
	return nil
}

func resourceSlbHwCompressRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading hw-compress (Inside resourceSlbHwCompressRead)")

	client := meta.(Thunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbHwCompress(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No hw-compress found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbHwCompressUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbHwCompressRead(d, meta)
}

func resourceSlbHwCompressDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbHwCompressRead(d, meta)
}

//Utility method to instantiate HwCompress Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbHwCompress(d *schema.ResourceData) go_thunder.HwCompress {
	var vc go_thunder.HwCompress

	var c go_thunder.HwCompressInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable36, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable36
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
