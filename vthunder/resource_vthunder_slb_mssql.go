package vthunder

//vThunder resource SlbMssql

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbMssql() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbMssqlCreate,
		Update: resourceSlbMssqlUpdate,
		Read:   resourceSlbMssqlRead,
		Delete: resourceSlbMssqlDelete,
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

func resourceSlbMssqlCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbMssql (Inside resourceSlbMssqlCreate) ")

		data := dataToSlbMssql(d)
		logger.Println("[INFO] received V from method data to SlbMssql --")
		d.SetId("1")
		go_vthunder.PostSlbMssql(client.Token, data, client.Host)

		return resourceSlbMssqlRead(d, meta)

	}
	return nil
}

func resourceSlbMssqlRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbMssql (Inside resourceSlbMssqlRead)")

	if client.Host != "" {

		data, err := go_vthunder.GetSlbMssql(client.Token, client.Host)
		if data == nil {

			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbMssqlUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbMssqlRead(d, meta)
}

func resourceSlbMssqlDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbMssqlRead(d, meta)
}

func dataToSlbMssql(d *schema.ResourceData) go_vthunder.Mssql {
	var vc go_vthunder.Mssql
	var c go_vthunder.MssqlInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable45, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable45
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
