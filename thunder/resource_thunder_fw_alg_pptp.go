package thunder

//Thunder resource FwAlgPptp

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwAlgPptp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwAlgPptpCreate,
		Update: resourceFwAlgPptpUpdate,
		Read:   resourceFwAlgPptpRead,
		Delete: resourceFwAlgPptpDelete,
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

func resourceFwAlgPptpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwAlgPptp (Inside resourceFwAlgPptpCreate) ")

		data := dataToFwAlgPptp(d)
		logger.Println("[INFO] received formatted data from method data to FwAlgPptp --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwAlgPptp(client.Token, data, client.Host)

		return resourceFwAlgPptpRead(d, meta)

	}
	return nil
}

func resourceFwAlgPptpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwAlgPptp (Inside resourceFwAlgPptpRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwAlgPptp(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwAlgPptpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgPptpRead(d, meta)
}

func resourceFwAlgPptpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgPptpRead(d, meta)
}
func dataToFwAlgPptp(d *schema.ResourceData) go_thunder.FwAlgPptp {
	var vc go_thunder.FwAlgPptp
	var c go_thunder.FwAlgPptpInstance
	c.DefaultPortDisable = d.Get("default_port_disable").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.FwAlgPptpSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.FwAlgPptpSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.DefaultPortDisable = c
	return vc
}
