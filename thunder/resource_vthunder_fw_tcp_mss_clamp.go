package thunder

//Thunder resource FwTcpMssClamp

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwTcpMssClamp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwTcpMssClampCreate,
		Update: resourceFwTcpMssClampUpdate,
		Read:   resourceFwTcpMssClampRead,
		Delete: resourceFwTcpMssClampDelete,
		Schema: map[string]*schema.Schema{
			"mss_subtract": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mss_value": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mss_clamp_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"min": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwTcpMssClampCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwTcpMssClamp (Inside resourceFwTcpMssClampCreate) ")

		data := dataToFwTcpMssClamp(d)
		logger.Println("[INFO] received formatted data from method data to FwTcpMssClamp --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwTcpMssClamp(client.Token, data, client.Host)

		return resourceFwTcpMssClampRead(d, meta)

	}
	return nil
}

func resourceFwTcpMssClampRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwTcpMssClamp (Inside resourceFwTcpMssClampRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwTcpMssClamp(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwTcpMssClampUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwTcpMssClampRead(d, meta)
}

func resourceFwTcpMssClampDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwTcpMssClampRead(d, meta)
}
func dataToFwTcpMssClamp(d *schema.ResourceData) go_thunder.FwTcpMssClamp {
	var vc go_thunder.FwTcpMssClamp
	var c go_thunder.FwTcpMssClampInstance
	c.MssSubtract = d.Get("mss_subtract").(int)
	c.MssValue = d.Get("mss_value").(int)
	c.MssClampType = d.Get("mss_clamp_type").(string)
	c.Min = d.Get("min").(int)

	vc.MssSubtract = c
	return vc
}
