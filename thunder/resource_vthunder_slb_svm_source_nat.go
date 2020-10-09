package thunder

//Thunder resource SlbSvmSourceNat

import (
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbSvmSourceNat() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbSvmSourceNatCreate,
		Update: resourceSlbSvmSourceNatUpdate,
		Read:   resourceSlbSvmSourceNatRead,
		Delete: resourceSlbSvmSourceNatDelete,
		Schema: map[string]*schema.Schema{
			"pool": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbSvmSourceNatCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSvmSourceNat (Inside resourceSlbSvmSourceNatCreate) ")

		data := dataToSlbSvmSourceNat(d)
		logger.Println("[INFO] received formatted data from method data to SlbSvmSourceNat --")
		d.SetId("1")
		go_thunder.PostSlbSvmSourceNat(client.Token, data, client.Host)

		return resourceSlbSvmSourceNatRead(d, meta)

	}
	return nil
}

func resourceSlbSvmSourceNatRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbSvmSourceNat (Inside resourceSlbSvmSourceNatRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbSvmSourceNat(client.Token, client.Host)
		if data == nil {

			return nil
		}
		return err
	}
	return nil
}

func resourceSlbSvmSourceNatUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSvmSourceNatRead(d, meta)
}

func resourceSlbSvmSourceNatDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSvmSourceNatRead(d, meta)
}

func dataToSlbSvmSourceNat(d *schema.ResourceData) go_thunder.SvmSourceNat {
	var vc go_thunder.SvmSourceNat
	var c go_thunder.SvmSourceNatInstance
	c.Pool = d.Get("pool").(string)

	vc.Pool = c
	return vc
}
