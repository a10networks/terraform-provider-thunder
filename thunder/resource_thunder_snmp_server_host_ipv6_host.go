package thunder

//Thunder resource SnmpServerHostIpv6Host

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
	"util"
)

func resourceSnmpServerHostIpv6Host() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerHostIpv6HostCreate,
		Update: resourceSnmpServerHostIpv6HostUpdate,
		Read:   resourceSnmpServerHostIpv6HostRead,
		Delete: resourceSnmpServerHostIpv6HostDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6_addr": {
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

func resourceSnmpServerHostIpv6HostCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerHostIpv6Host (Inside resourceSnmpServerHostIpv6HostCreate) ")
		name1 := d.Get("ipv6_addr").(string)
		name2 := d.Get("version").(string)
		data := dataToSnmpServerHostIpv6Host(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerHostIpv6Host --")
		d.SetId(name1 + "," + name2)
		go_thunder.PostSnmpServerHostIpv6Host(client.Token, data, client.Host)

		return resourceSnmpServerHostIpv6HostRead(d, meta)

	}
	return nil
}

func resourceSnmpServerHostIpv6HostRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerHostIpv6Host (Inside resourceSnmpServerHostIpv6HostRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerHostIpv6Host(client.Token, name1, name2, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerHostIpv6HostUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying SnmpServerHostIpv6Host   (Inside resourceSnmpServerHostIpv6HostUpdate) ")
		data := dataToSnmpServerHostIpv6Host(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerHostIpv6Host ")
		go_thunder.PutSnmpServerHostIpv6Host(client.Token, name1, name2, data, client.Host)

		return resourceSnmpServerHostIpv6HostRead(d, meta)

	}
	return nil
}

func resourceSnmpServerHostIpv6HostDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerHostIpv6HostDelete) " + name1)
		err := go_thunder.DeleteSnmpServerHostIpv6Host(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToSnmpServerHostIpv6Host(d *schema.ResourceData) go_thunder.SnmpServerHostIpv6Host {
	var vc go_thunder.SnmpServerHostIpv6Host
	var c go_thunder.SnmpServerHostIpv6HostInstance
	c.Ipv6Addr = d.Get("ipv6_addr").(string)
	c.UDPPort = d.Get("udp_port").(int)
	c.V1V2CComm = d.Get("v1_v2c_comm").(string)
	c.User = d.Get("user").(string)
	c.Version = d.Get("version").(string)

	vc.UUID = c
	return vc
}
