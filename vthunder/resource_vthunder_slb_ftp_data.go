package vthunder

// vThunder resource Slb FTPData

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbFTPData() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbFTPDataCreate,
		Update: resourceSlbFTPDataUpdate,
		Read:   resourceSlbFTPDataRead,
		Delete: resourceSlbFTPDataDelete,

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

func resourceSlbFTPDataCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating FTPData (Inside resourceSlbFTPDataCreate)")

	if client.Host != "" {
		vc := dataToSlbFTPData(d)
		d.SetId("1")
		go_vthunder.PostSlbFTPData(client.Token, vc, client.Host)

		return resourceSlbFTPDataRead(d, meta)
	}
	return nil
}

func resourceSlbFTPDataRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading FTPData (Inside resourceSlbFTPDataRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbFTPData(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No FTPData found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbFTPDataUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbFTPDataRead(d, meta)
}

func resourceSlbFTPDataDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbFTPDataRead(d, meta)
}

//Utility method to instantiate FTPData Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbFTPData(d *schema.ResourceData) go_vthunder.FTPData {
	var vc go_vthunder.FTPData

	var c go_vthunder.FTPDataInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable21, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable21
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
