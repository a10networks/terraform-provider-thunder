package thunder

//Thunder resource FwSystemStatus

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwSystemStatus() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwSystemStatusCreate,
		Update: resourceFwSystemStatusUpdate,
		Read:   resourceFwSystemStatusRead,
		Delete: resourceFwSystemStatusDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwSystemStatusCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwSystemStatus (Inside resourceFwSystemStatusCreate) ")

		data := dataToFwSystemStatus(d)
		logger.Println("[INFO] received formatted data from method data to FwSystemStatus --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwSystemStatus(client.Token, data, client.Host)

		return resourceFwSystemStatusRead(d, meta)

	}
	return nil
}

func resourceFwSystemStatusRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwSystemStatus (Inside resourceFwSystemStatusRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwSystemStatus(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwSystemStatusUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwSystemStatusRead(d, meta)
}

func resourceFwSystemStatusDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwSystemStatusRead(d, meta)
}
func dataToFwSystemStatus(d *schema.ResourceData) go_thunder.FwSystemStatus {
	var vc go_thunder.FwSystemStatus
	var c go_thunder.FwSystemStatusInstance

	vc.UUID = c
	return vc
}
