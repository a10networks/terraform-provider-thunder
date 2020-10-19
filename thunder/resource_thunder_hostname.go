package thunder

//Thunder resource Hostname

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceHostname() *schema.Resource {
	return &schema.Resource{
		Create: resourceHostnameCreate,
		Update: resourceHostnameUpdate,
		Read:   resourceHostnameRead,
		Delete: resourceHostnameDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceHostnameCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating Hostname (Inside resourceHostnameCreate) ")

		data := dataToHostname(d)
		logger.Println("[INFO] received formatted data from method data to Hostname --")
		d.SetId("1")
		go_thunder.PostHostname(client.Token, data, client.Host)

		return resourceHostnameRead(d, meta)

	}
	return nil
}

func resourceHostnameRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading Hostname (Inside resourceHostnameRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read hostname")
		data, err := go_thunder.GetHostname(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found hostname ")
			return nil
		}
		return err
	}
	return nil
}

func resourceHostnameUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceHostnameRead(d, meta)
}

func resourceHostnameDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceHostnameRead(d, meta)
}
func dataToHostname(d *schema.ResourceData) go_thunder.Hostname {
	var vc go_thunder.Hostname
	var c go_thunder.HostnameInstance
	c.Value = d.Get("value").(string)

	vc.UUID = c
	return vc
}
