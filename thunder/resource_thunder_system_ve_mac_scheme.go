package thunder

//Thunder resource SystemVeMacScheme

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceSystemVeMacScheme() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVeMacSchemeCreate,
		Update: resourceSystemVeMacSchemeUpdate,
		Read:   resourceSystemVeMacSchemeRead,
		Delete: resourceSystemVeMacSchemeDelete,
		Schema: map[string]*schema.Schema{
			"ve_mac_scheme_val": {
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

func resourceSystemVeMacSchemeCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SystemVeMacScheme (Inside resourceSystemVeMacSchemeCreate) ")

		data := dataToSystemVeMacScheme(d)
		logger.Println("[INFO] received formatted data from method data to SystemVeMacScheme --")
		d.SetId("1")
		go_thunder.PostSystemVeMacScheme(client.Token, data, client.Host)

		return resourceSystemVeMacSchemeRead(d, meta)

	}
	return nil
}

func resourceSystemVeMacSchemeRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SystemVeMacScheme (Inside resourceSystemVeMacSchemeRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSystemVeMacScheme(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSystemVeMacSchemeUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSystemVeMacSchemeRead(d, meta)
}

func resourceSystemVeMacSchemeDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSystemVeMacSchemeRead(d, meta)
}
func dataToSystemVeMacScheme(d *schema.ResourceData) go_thunder.SystemVeMacScheme {
	var vc go_thunder.SystemVeMacScheme
	var c go_thunder.SystemVeMacSchemeInstance
	c.VeMacSchemeVal = d.Get("ve_mac_scheme_val").(string)

	vc.VeMacSchemeVal = c
	return vc
}
