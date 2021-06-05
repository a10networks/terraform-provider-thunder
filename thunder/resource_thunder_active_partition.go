package thunder

//Thunder resource ActivePartition

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceActivePartition() *schema.Resource {
	return &schema.Resource{
		Create: resourceActivePartitionCreate,
		Update: resourceActivePartitionUpdate,
		Read:   resourceActivePartitionRead,
		Delete: resourceActivePartitionDelete,
		Schema: map[string]*schema.Schema{
			"shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"curr_part_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceActivePartitionCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating ActivePartition (Inside resourceActivePartitionCreate) ")

		data := dataToActivePartition(d)
		logger.Println("[INFO] received formatted data from method data to ActivePartition --")
		d.SetId("1")
		go_thunder.PostActivePartition(client.Token, data, client.Host)

		return resourceActivePartitionRead(d, meta)

	}
	return nil
}

func resourceActivePartitionRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading ActivePartition (Inside resourceActivePartitionRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetActivePartition(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceActivePartitionUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceActivePartitionRead(d, meta)
}

func resourceActivePartitionDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceActivePartitionRead(d, meta)
}
func dataToActivePartition(d *schema.ResourceData) go_thunder.ActivePartition {
	var vc go_thunder.ActivePartition
	var c go_thunder.ActivePartitionInstance
	c.Shared = d.Get("shared").(int)
	c.CurrPartName = d.Get("curr_part_name").(string)

	vc.Shared = c
	return vc
}
