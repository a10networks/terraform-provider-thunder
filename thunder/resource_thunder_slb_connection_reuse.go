package thunder

// Thunder resource Slb ConnectionReuse

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbConnectionReuse() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbConnectionReuseCreate,
		Update: resourceSlbConnectionReuseUpdate,
		Read:   resourceSlbConnectionReuseRead,
		Delete: resourceSlbConnectionReuseDelete,

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

func resourceSlbConnectionReuseCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	logger.Println("[INFO] Creating connection-reuse (Inside resourceSlbConnectionReuseCreate)")

	if client.Host != "" {
		vc := dataToSlbConnectionReuse(d)
		d.SetId("1")
		go_thunder.PostSlbConnectionReuse(client.Token, vc, client.Host)

		return resourceSlbConnectionReuseRead(d, meta)
	}
	return nil
}

func resourceSlbConnectionReuseRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading connection-reuse (Inside resourceSlbConnectionReuseRead)")

	client := meta.(Thunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbConnectionReuse(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No connection-reuse found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbConnectionReuseUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbConnectionReuseRead(d, meta)
}

func resourceSlbConnectionReuseDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbConnectionReuseRead(d, meta)
}

//Utility method to instantiate ConnectionReuse Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbConnectionReuse(d *schema.ResourceData) go_thunder.ConnectionReuse {
	var vc go_thunder.ConnectionReuse

	var c go_thunder.SlbConnectionReuseInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable7, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable7
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
