package vthunder

//vThunder resource FwAlgTftp

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwAlgTftp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwAlgTftpCreate,
		Update: resourceFwAlgTftpUpdate,
		Read:   resourceFwAlgTftpRead,
		Delete: resourceFwAlgTftpDelete,
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

func resourceFwAlgTftpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwAlgTftp (Inside resourceFwAlgTftpCreate) ")

		data := dataToFwAlgTftp(d)
		logger.Println("[INFO] received formatted data from method data to FwAlgTftp --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwAlgTftp(client.Token, data, client.Host)

		return resourceFwAlgTftpRead(d, meta)

	}
	return nil
}

func resourceFwAlgTftpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwAlgTftp (Inside resourceFwAlgTftpRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwAlgTftp(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwAlgTftpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgTftpRead(d, meta)
}

func resourceFwAlgTftpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgTftpRead(d, meta)
}
func dataToFwAlgTftp(d *schema.ResourceData) go_vthunder.FwAlgTftp {
	var vc go_vthunder.FwAlgTftp
	var c go_vthunder.FwAlgTftpInstance
	c.DefaultPortDisable = d.Get("default_port_disable").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.FwAlgTftpSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.FwAlgTftpSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.DefaultPortDisable = c
	return vc
}
