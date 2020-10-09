package thunder

//Thunder resource Partition

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
	"util"
)

func resourceOverlayTunnelPartition() *schema.Resource {
	return &schema.Resource{
		Create: resourcePartitionCreate,
		Update: resourcePartitionUpdate,
		Read:   resourcePartitionRead,
		Delete: resourcePartitionDelete,

		Schema: map[string]*schema.Schema{
			"template": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"resource_accounting": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"partition_name": {
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
						"vrid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vlan": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						//						"allowable_ipv6_range": {
						//							Type:     schema.TypeList,
						//							Optional: true,
						//						},
						"mgmt_floating_ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						//						"allowable_ip_range": {
						//							Type:     schema.TypeList,
						//							Optional: true,
						//						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"id2": {
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
		},
	}
}

func resourcePartitionCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := strconv.Itoa(d.Get("id2").(int))
		logger.Println("[INFO] Creating partition (Inside resourcePartitionCreate    " + name)
		v := dataToPartition(name, d)
		d.SetId(name)
		go_thunder.PostPartition(client.Token, v, client.Host)

		return resourcePartitionRead(d, meta)
	}
	return nil
}

func resourcePartitionRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading partition (Inside resourcePartitionRead)")

	if client.Host != "" {
		client := meta.(Thunder)

		name := d.Id()

		logger.Println("[INFO] Fetching partition Read" + name)

		partition, err := go_thunder.GetPartition(client.Token, name, client.Host)

		if partition == nil {
			logger.Println("[INFO] No partition found " + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourcePartitionUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := strconv.Itoa(d.Get("id2").(int))
		logger.Println("[INFO] Modifying partition (Inside resourcePartitionUpdate    " + name)
		v := dataToPartition(name, d)
		d.SetId(name)
		go_thunder.PutPartition(client.Token, name, v, client.Host)

		return resourcePartitionRead(d, meta)
	}
	return nil
}

func resourcePartitionDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(Thunder)
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting partition (Inside resourcePartitionDelete) " + name)

		err := go_thunder.DeletePartition(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete partition  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate partition structure
func dataToPartition(name string, d *schema.ResourceData) go_thunder.Partition {
	var s go_thunder.Partition

	var sInstance go_thunder.PartitionInstance

	sInstance.UserTag = d.Get("user_tag").(string)
	sInstance.PartitionName = d.Get("partition_name").(string)
	sInstance.ApplicationType = d.Get("application_type").(string)
	sInstance.ID = d.Get("id2").(int)

	var vlans go_thunder.SharedVlan

	prefix := fmt.Sprintf("shared_vlan.0.")
	vlans.MgmtFloatingIPAddress = d.Get(prefix + "mgmt_floating_ip_address").(string)
	vlans.Vrid = d.Get(prefix + "vrid").(int)
	vlans.Vlan = d.Get(prefix + "vlan").(int)

	sInstance.Vrid = vlans

	var template go_thunder.Template

	prefix2 := fmt.Sprintf("template.0.")
	template.ResourceAccounting = d.Get(prefix2 + "resource_accounting").(string)

	sInstance.ResourceAccounting = template

	s.ID = sInstance

	return s
}
