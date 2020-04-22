package vthunder

//vThunder resource SlbMysql

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbMysql() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbMysqlCreate,
		Update: resourceSlbMysqlUpdate,
		Read:   resourceSlbMysqlRead,
		Delete: resourceSlbMysqlDelete,
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

func resourceSlbMysqlCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbMysql (Inside resourceSlbMysqlCreate) ")

		data := dataToSlbMysql(d)
		logger.Println("[INFO] received V from method data to SlbMysql --")
		d.SetId("1")
		go_vthunder.PostSlbMysql(client.Token, data, client.Host)

		return resourceSlbMysqlRead(d, meta)

	}
	return nil
}

func resourceSlbMysqlRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbMysql (Inside resourceSlbMysqlRead)")

	if client.Host != "" {

		data, err := go_vthunder.GetSlbMysql(client.Token, client.Host)
		if data == nil {

			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}
func resourceSlbMysqlUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbMysqlRead(d, meta)
}

func resourceSlbMysqlDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbMysqlRead(d, meta)
}
func dataToSlbMysql(d *schema.ResourceData) go_vthunder.Mysql {
	var vc go_vthunder.Mysql
	var c go_vthunder.MysqlInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable46, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable46
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
