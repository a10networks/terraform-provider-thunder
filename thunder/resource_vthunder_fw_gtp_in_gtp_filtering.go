package thunder

//Thunder resource FwGtpInGtpFiltering

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwGtpInGtpFiltering() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwGtpInGtpFilteringCreate,
		Update: resourceFwGtpInGtpFilteringUpdate,
		Read:   resourceFwGtpInGtpFilteringRead,
		Delete: resourceFwGtpInGtpFilteringDelete,
		Schema: map[string]*schema.Schema{
			"gtp_in_gtp_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwGtpInGtpFilteringCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwGtpInGtpFiltering (Inside resourceFwGtpInGtpFilteringCreate) ")

		data := dataToFwGtpInGtpFiltering(d)
		logger.Println("[INFO] received formatted data from method data to FwGtpInGtpFiltering --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwGtpInGtpFiltering(client.Token, data, client.Host)

		return resourceFwGtpInGtpFilteringRead(d, meta)

	}
	return nil
}

func resourceFwGtpInGtpFilteringRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwGtpInGtpFiltering (Inside resourceFwGtpInGtpFilteringRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwGtpInGtpFiltering(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwGtpInGtpFilteringUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwGtpInGtpFilteringRead(d, meta)
}

func resourceFwGtpInGtpFilteringDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwGtpInGtpFilteringRead(d, meta)
}
func dataToFwGtpInGtpFiltering(d *schema.ResourceData) go_thunder.FwGtpInGtpFiltering {
	var vc go_thunder.FwGtpInGtpFiltering
	var c go_thunder.FwGtpInGtpFilteringInstance
	c.GtpInGtpValue = d.Get("gtp_in_gtp_value").(string)

	vc.GtpInGtpValue = c
	return vc
}
