package vthunder

//vThunder resource FwGtp

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwGtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwGtpCreate,
		Update: resourceFwGtpUpdate,
		Read:   resourceFwGtpRead,
		Delete: resourceFwGtpDelete,
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
			"gtp_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwGtpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwGtp (Inside resourceFwGtpCreate) ")

		data := dataToFwGtp(d)
		logger.Println("[INFO] received formatted data from method data to FwGtp --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwGtp(client.Token, data, client.Host)

		return resourceFwGtpRead(d, meta)

	}
	return nil
}

func resourceFwGtpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwGtp (Inside resourceFwGtpRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwGtp(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwGtpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwGtpRead(d, meta)
}

func resourceFwGtpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwGtpRead(d, meta)
}
func dataToFwGtp(d *schema.ResourceData) go_vthunder.FwGtp {
	var vc go_vthunder.FwGtp
	var c go_vthunder.FwGtpInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.FwGtpSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.FwGtpSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	c.GtpValue = d.Get("gtp_value").(string)

	vc.SamplingEnable = c
	return vc
}
