package thunder

//Thunder resource SlbSmtp

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbSmtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbSmtpCreate,
		Update: resourceSlbSmtpUpdate,
		Read:   resourceSlbSmtpRead,
		Delete: resourceSlbSmtpDelete,
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

func resourceSlbSmtpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSmtp (Inside resourceSlbSmtpCreate) ")

		data := dataToSlbSmtp(d)
		logger.Println("[INFO] received formatted data from method data to SlbSmtp --")
		d.SetId("1")
		go_thunder.PostSlbSmtp(client.Token, data, client.Host)

		return resourceSlbSmtpRead(d, meta)

	}
	return nil
}

func resourceSlbSmtpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbSmtp (Inside resourceSlbSmtpRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbSmtp(client.Token, client.Host)
		if data == nil {

			return nil
		}
		return err
	}
	return nil
}

func resourceSlbSmtpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSmtpRead(d, meta)
}

func resourceSlbSmtpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSmtpRead(d, meta)
}
func dataToSlbSmtp(d *schema.ResourceData) go_thunder.SlbSmtp {
	var vc go_thunder.SlbSmtp
	var c go_thunder.SlbSmtpInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable12, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable12
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		obj1.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
