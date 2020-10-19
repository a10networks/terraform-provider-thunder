package thunder

//Thunder resource GlmSend

import (
	go_thunder "github.com/go_thunder/thunder"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceGlmSend() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlmSendCreate,
		Update: resourceGlmSendUpdate,
		Read:   resourceGlmSendRead,
		Delete: resourceGlmSendDelete,
		Schema: map[string]*schema.Schema{
			"license_request": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceGlmSendCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GlmSend (Inside resourceGlmSendCreate) ")

		data := dataToGlmSend(d)
		logger.Println("[INFO] received formatted data from method data to GlmSend --")
		d.SetId("1")
		go_thunder.PostGlmSend(client.Token, data, client.Host)

		return resourceGlmSendRead(d, meta)

	}
	return nil
}

func resourceGlmSendRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading GlmSend (Inside resourceGlmSendRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetGlmSend(client.Token, client.Host)
		if data == nil {

			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceGlmSendUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceGlmSendRead(d, meta)
}

func resourceGlmSendDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceGlmSendRead(d, meta)
}
func dataToGlmSend(d *schema.ResourceData) go_thunder.GlmSend {
	var vc go_thunder.GlmSend
	var c go_thunder.GlmSendInstance
	c.LicenseRequest = d.Get("license_request").(int)

	vc.LicenseRequest = c
	return vc
}
