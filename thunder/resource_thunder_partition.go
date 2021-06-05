package thunder

//Thunder resource Partition

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourcePartition() *schema.Resource {
	return &schema.Resource{
		Create: resourcePartitionCreate,
		Update: resourcePartitionUpdate,
		Read:   resourcePartitionRead,
		Delete: resourcePartitionDelete,
		Schema: map[string]*schema.Schema{
			"partition_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"id1": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"application_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_vlan": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"mgmt_floating_ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"vrid": {
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
				},
			},
			"template": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_accounting": {
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
				},
			},
		},
	}
}

func resourcePartitionCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating Partition (Inside resourcePartitionCreate) ")
		name1 := d.Get("partition_name").(string)
		data := dataToPartition(d)
		logger.Println("[INFO] received formatted data from method data to Partition --")
		d.SetId(name1)
		go_thunder.PostPartition(client.Token, data, client.Host)

		return resourcePartitionRead(d, meta)

	}
	return nil
}

func resourcePartitionRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading Partition (Inside resourcePartitionRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetPartition(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourcePartitionUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourcePartitionRead(d, meta)
}

func resourcePartitionDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourcePartitionDelete) " + name1)
		err := go_thunder.DeletePartition(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToPartition(d *schema.ResourceData) go_thunder.Partition {
	var vc go_thunder.Partition
	var c go_thunder.PartitionInstance
	c.PartitionName = d.Get("partition_name").(string)
	c.ID = d.Get("id1").(int)
	c.ApplicationType = d.Get("application_type").(string)
	c.UserTag = d.Get("user_tag").(string)

	var obj1 go_thunder.SharedVlan
	prefix1 := "shared_vlan.0."
	obj1.Vlan = d.Get(prefix1 + "vlan").(int)
	obj1.MgmtFloatingIPAddress = d.Get(prefix1 + "mgmt_floating_ip_address").(string)
	obj1.Vrid = d.Get(prefix1 + "vrid").(int)

	c.Vlan = obj1

	var obj2 go_thunder.Template
	prefix2 := "template.0."
	obj2.ResourceAccounting = d.Get(prefix2 + "resource_accounting").(string)

	c.ResourceAccounting = obj2

	vc.PartitionName = c
	return vc
}
