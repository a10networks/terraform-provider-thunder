package thunder

//Thunder resource FwClearSessionFilter

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwClearSessionFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwClearSessionFilterCreate,
		Update: resourceFwClearSessionFilterUpdate,
		Read:   resourceFwClearSessionFilterRead,
		Delete: resourceFwClearSessionFilterDelete,
		Schema: map[string]*schema.Schema{
			"status": {
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

func resourceFwClearSessionFilterCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwClearSessionFilter (Inside resourceFwClearSessionFilterCreate) ")

		data := dataToFwClearSessionFilter(d)
		logger.Println("[INFO] received formatted data from method data to FwClearSessionFilter --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwClearSessionFilter(client.Token, data, client.Host)

		return resourceFwClearSessionFilterRead(d, meta)

	}
	return nil
}

func resourceFwClearSessionFilterRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwClearSessionFilter (Inside resourceFwClearSessionFilterRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwClearSessionFilter(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwClearSessionFilterUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwClearSessionFilterRead(d, meta)
}

func resourceFwClearSessionFilterDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwClearSessionFilterRead(d, meta)
}
func dataToFwClearSessionFilter(d *schema.ResourceData) go_thunder.FwClearSessionFilter {
	var vc go_thunder.FwClearSessionFilter
	var c go_thunder.FwClearSessionFilterInstance
	c.Status = d.Get("status").(string)

	vc.Status = c
	return vc
}
