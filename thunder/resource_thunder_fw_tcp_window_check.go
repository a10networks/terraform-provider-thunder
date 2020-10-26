package thunder

//Thunder resource FwTcpWindowCheck

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwTcpWindowCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwTcpWindowCheckCreate,
		Update: resourceFwTcpWindowCheckUpdate,
		Read:   resourceFwTcpWindowCheckRead,
		Delete: resourceFwTcpWindowCheckDelete,
		Schema: map[string]*schema.Schema{
			"status": {
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

func resourceFwTcpWindowCheckCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwTcpWindowCheck (Inside resourceFwTcpWindowCheckCreate) ")

		data := dataToFwTcpWindowCheck(d)
		logger.Println("[INFO] received formatted data from method data to FwTcpWindowCheck --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwTcpWindowCheck(client.Token, data, client.Host)

		return resourceFwTcpWindowCheckRead(d, meta)

	}
	return nil
}

func resourceFwTcpWindowCheckRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwTcpWindowCheck (Inside resourceFwTcpWindowCheckRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwTcpWindowCheck(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwTcpWindowCheckUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwTcpWindowCheckRead(d, meta)
}

func resourceFwTcpWindowCheckDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwTcpWindowCheckRead(d, meta)
}
func dataToFwTcpWindowCheck(d *schema.ResourceData) go_thunder.FwTcpWindowCheck {
	var vc go_thunder.FwTcpWindowCheck
	var c go_thunder.FwTCPWindowCheckInstance
	c.Status = d.Get("status").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.FwTCPWindowCheckSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.FwTCPWindowCheckSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Status = c
	return vc
}
