package thunder

//Thunder resource FwTcpResetOnError

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwTcpResetOnError() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwTcpResetOnErrorCreate,
		Update: resourceFwTcpResetOnErrorUpdate,
		Read:   resourceFwTcpResetOnErrorRead,
		Delete: resourceFwTcpResetOnErrorDelete,
		Schema: map[string]*schema.Schema{
			"enable": {
				Type:        schema.TypeInt,
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

func resourceFwTcpResetOnErrorCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwTcpResetOnError (Inside resourceFwTcpResetOnErrorCreate) ")

		data := dataToFwTcpResetOnError(d)
		logger.Println("[INFO] received formatted data from method data to FwTcpResetOnError --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwTcpResetOnError(client.Token, data, client.Host)

		return resourceFwTcpResetOnErrorRead(d, meta)

	}
	return nil
}

func resourceFwTcpResetOnErrorRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwTcpResetOnError (Inside resourceFwTcpResetOnErrorRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwTcpResetOnError(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwTcpResetOnErrorUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwTcpResetOnErrorRead(d, meta)
}

func resourceFwTcpResetOnErrorDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwTcpResetOnErrorRead(d, meta)
}
func dataToFwTcpResetOnError(d *schema.ResourceData) go_thunder.FwTcpResetOnError {
	var vc go_thunder.FwTcpResetOnError
	var c go_thunder.FwTcpResetOnErrorInstance
	c.Enable = d.Get("enable").(int)

	vc.Enable = c
	return vc
}
