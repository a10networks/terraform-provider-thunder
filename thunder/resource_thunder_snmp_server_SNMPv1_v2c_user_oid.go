package thunder

//Thunder resource SnmpServerSNMPv1V2cUserOid

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
	"util"
)

func resourceSnmpServerSNMPv1V2cUserOid() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerSNMPv1V2cUserOidCreate,
		Update: resourceSnmpServerSNMPv1V2cUserOidUpdate,
		Read:   resourceSnmpServerSNMPv1V2cUserOidRead,
		Delete: resourceSnmpServerSNMPv1V2cUserOidDelete,
		Schema: map[string]*schema.Schema{
			"oid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"remote": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns_host": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ipv4_mask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ipv4_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_host": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ipv4_mask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ipv6_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_host": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ipv6_mask": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"oid_val": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
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

func resourceSnmpServerSNMPv1V2cUserOidCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerSNMPv1V2cUserOid (Inside resourceSnmpServerSNMPv1V2cUserOidCreate) ")
		name2 := d.Get("oid_val").(string)
		name1 := d.Get("user_tag").(string)
		data := dataToSnmpServerSNMPv1V2cUserOid(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerSNMPv1V2cUserOid --")
		d.SetId(name1 + "," + name2)
		go_thunder.PostSnmpServerSNMPv1V2cUserOid(client.Token, name1, data, client.Host)

		return resourceSnmpServerSNMPv1V2cUserOidRead(d, meta)

	}
	return nil
}

func resourceSnmpServerSNMPv1V2cUserOidRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerSNMPv1V2cUserOid (Inside resourceSnmpServerSNMPv1V2cUserOidRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerSNMPv1V2cUserOid(client.Token, name1, name2, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerSNMPv1V2cUserOidUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying SnmpServerSNMPv1V2cUserOid   (Inside resourceSnmpServerSNMPv1V2cUserOidUpdate) ")
		data := dataToSnmpServerSNMPv1V2cUserOid(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerSNMPv1V2cUserOid ")
		go_thunder.PutSnmpServerSNMPv1V2cUserOid(client.Token, name1, name2, data, client.Host)

		return resourceSnmpServerSNMPv1V2cUserOidRead(d, meta)

	}
	return nil
}

func resourceSnmpServerSNMPv1V2cUserOidDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerSNMPv1V2cUserOidDelete) " + name1)
		err := go_thunder.DeleteSnmpServerSNMPv1V2cUserOid(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToSnmpServerSNMPv1V2cUserOid(d *schema.ResourceData) go_thunder.SnmpServerSNMPv1V2cUserOid {
	var vc go_thunder.SnmpServerSNMPv1V2cUserOid
	var c go_thunder.SnmpServerSNMPv1V2cUserOidInstance

	var obj1 go_thunder.SnmpServerSNMPv1V2cUserRemote
	prefix := "remote.0."

	HostListCount := d.Get(prefix + "host_list.#").(int)
	obj1.DNSHost = make([]go_thunder.SnmpServerSNMPv1V2cUserHostList, 0, HostListCount)

	for i := 0; i < HostListCount; i++ {
		var obj1_1 go_thunder.SnmpServerSNMPv1V2cUserHostList
		prefix1_1 := prefix + fmt.Sprintf("host_list.%d.", i)
		obj1_1.DNSHost = d.Get(prefix1_1 + "dns_host").(string)
		obj1_1.Ipv4Mask = d.Get(prefix1_1 + "ipv4_mask").(string)
		obj1.DNSHost = append(obj1.DNSHost, obj1_1)
	}

	Ipv4ListCount := d.Get(prefix + "ipv4_list.#").(int)
	obj1.Ipv4Host = make([]go_thunder.SnmpServerSNMPv1V2cUserIpv4List, 0, Ipv4ListCount)

	for i := 0; i < Ipv4ListCount; i++ {
		var obj1_2 go_thunder.SnmpServerSNMPv1V2cUserIpv4List
		prefix1_2 := prefix + fmt.Sprintf("ipv4_list.%d.", i)
		obj1_2.Ipv4Host = d.Get(prefix1_2 + "ipv4_host").(string)
		obj1_2.Ipv4Mask = d.Get(prefix1_2 + "ipv4_mask").(string)
		obj1.Ipv4Host = append(obj1.Ipv4Host, obj1_2)
	}

	Ipv6ListCount := d.Get(prefix + "ipv6_list.#").(int)
	obj1.Ipv6Host = make([]go_thunder.SnmpServerSNMPv1V2cUserIpv6List, 0, Ipv6ListCount)

	for i := 0; i < Ipv6ListCount; i++ {
		var obj1_3 go_thunder.SnmpServerSNMPv1V2cUserIpv6List
		prefix1_3 := prefix + fmt.Sprintf("ipv6_list.%d.", i)
		obj1_3.Ipv6Host = d.Get(prefix1_3 + "ipv6_host").(string)
		obj1_3.Ipv6Mask = d.Get(prefix1_3 + "ipv6_mask").(int)
		obj1.Ipv6Host = append(obj1.Ipv6Host, obj1_3)
	}

	c.DNSHost = obj1

	c.OidVal = d.Get("oid_val").(string)
	c.UserTag = d.Get("user_tag").(string)

	vc.OidVal = c
	return vc
}
