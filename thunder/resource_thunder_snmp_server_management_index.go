package thunder

//Thunder resource SnmpServerManagementIndex

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceSnmpServerManagementIndex() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerManagementIndexCreate,
		Update: resourceSnmpServerManagementIndexUpdate,
		Read:   resourceSnmpServerManagementIndexRead,
		Delete: resourceSnmpServerManagementIndexDelete,
		Schema: map[string]*schema.Schema{
			"mgmt_index": {
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

func resourceSnmpServerManagementIndexCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerManagementIndex (Inside resourceSnmpServerManagementIndexCreate) ")

		data := dataToSnmpServerManagementIndex(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerManagementIndex --")
		d.SetId("1")
		go_thunder.PostSnmpServerManagementIndex(client.Token, data, client.Host)

		return resourceSnmpServerManagementIndexRead(d, meta)

	}
	return nil
}

func resourceSnmpServerManagementIndexRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerManagementIndex (Inside resourceSnmpServerManagementIndexRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerManagementIndex(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerManagementIndexUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerManagementIndexRead(d, meta)
}

func resourceSnmpServerManagementIndexDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerManagementIndexRead(d, meta)
}
func dataToSnmpServerManagementIndex(d *schema.ResourceData) go_thunder.SnmpServerManagementIndex {
	var vc go_thunder.SnmpServerManagementIndex
	var c go_thunder.SnmpServerManagementIndexInstance
	c.MgmtIndex = d.Get("mgmt_index").(int)

	vc.MgmtIndex = c
	return vc
}
