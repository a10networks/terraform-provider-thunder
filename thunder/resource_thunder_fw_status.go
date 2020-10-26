package thunder

//Thunder resource FwStatus

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwStatus() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwStatusCreate,
		Update: resourceFwStatusUpdate,
		Read:   resourceFwStatusRead,
		Delete: resourceFwStatusDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwStatusCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwStatus (Inside resourceFwStatusCreate) ")

		data := dataToFwStatus(d)
		logger.Println("[INFO] received formatted data from method data to FwStatus --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwStatus(client.Token, data, client.Host)

		return resourceFwStatusRead(d, meta)

	}
	return nil
}

func resourceFwStatusRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwStatus (Inside resourceFwStatusRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwStatus(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwStatusUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwStatusRead(d, meta)
}

func resourceFwStatusDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwStatusRead(d, meta)
}
func dataToFwStatus(d *schema.ResourceData) go_thunder.FwStatus {
	var vc go_thunder.FwStatus
	var c go_thunder.FwStatusInstance

	vc.UUID = c
	return vc
}
