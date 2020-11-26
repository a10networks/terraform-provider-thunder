package thunder

//Thunder resource SnmpServerHostIpv4Host

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
	"util"
)

func resourceSnmpServerHostIpv4Host() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerHostIpv4HostCreate,
		Update: resourceSnmpServerHostIpv4HostUpdate,
		Read:   resourceSnmpServerHostIpv4HostRead,
		Delete: resourceSnmpServerHostIpv4HostDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv4_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"udp_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"v1_v2c_comm": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerHostIpv4HostCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerHostIpv4Host (Inside resourceSnmpServerHostIpv4HostCreate) ")
		name2 := d.Get("version").(string)
		name1 := d.Get("ipv4_addr").(string)
		data := dataToSnmpServerHostIpv4Host(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerHostIpv4Host --")
		d.SetId(name1 + "," + name2)
		go_thunder.PostSnmpServerHostIpv4Host(client.Token, data, client.Host)

		return resourceSnmpServerHostIpv4HostRead(d, meta)

	}
	return nil
}

func resourceSnmpServerHostIpv4HostRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerHostIpv4Host (Inside resourceSnmpServerHostIpv4HostRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerHostIpv4Host(client.Token, name1, name2, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerHostIpv4HostUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying SnmpServerHostIpv4Host   (Inside resourceSnmpServerHostIpv4HostUpdate) ")
		data := dataToSnmpServerHostIpv4Host(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerHostIpv4Host ")
		go_thunder.PutSnmpServerHostIpv4Host(client.Token, name1, name2, data, client.Host)

		return resourceSnmpServerHostIpv4HostRead(d, meta)

	}
	return nil
}

func resourceSnmpServerHostIpv4HostDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerHostIpv4HostDelete) " + name1)
		err := go_thunder.DeleteSnmpServerHostIpv4Host(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToSnmpServerHostIpv4Host(d *schema.ResourceData) go_thunder.SnmpServerHostIpv4Host {
	var vc go_thunder.SnmpServerHostIpv4Host
	var c go_thunder.SnmpServerHostIpv4HostInstance
	c.Ipv4Addr = d.Get("ipv4_addr").(string)
	c.UDPPort = d.Get("udp_port").(int)
	c.V1V2CComm = d.Get("v1_v2c_comm").(string)
	c.User = d.Get("user").(string)
	c.Version = d.Get("version").(string)

	vc.UUID = c
	return vc
}
