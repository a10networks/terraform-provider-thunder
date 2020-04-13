package vthunder

// vThunder resource Slb FTPProxy

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbFTPProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbFTPProxyCreate,
		Update: resourceSlbFTPProxyUpdate,
		Read:   resourceSlbFTPProxyRead,
		Delete: resourceSlbFTPProxyDelete,

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

func resourceSlbFTPProxyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating ftp-proxy (Inside resourceSlbFTPProxyCreate)")

	if client.Host != "" {
		vc := dataToSlbFTPProxy(d)
		d.SetId("1")
		go_vthunder.PostSlbFTPProxy(client.Token, vc, client.Host)

		return resourceSlbFTPProxyRead(d, meta)
	}
	return nil
}

func resourceSlbFTPProxyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading ftp-proxy (Inside resourceSlbFTPProxyRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbFTPProxy(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No ftp-proxy found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbFTPProxyUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbFTPProxyRead(d, meta)
}

func resourceSlbFTPProxyDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbFTPProxyRead(d, meta)
}

//Utility method to instantiate FTPProxy Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbFTPProxy(d *schema.ResourceData) go_vthunder.FTPProxy {
	var vc go_vthunder.FTPProxy

	var c go_vthunder.FTPProxyInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable22, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable22
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
