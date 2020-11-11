package thunder

//Thunder resource SnmpServerCommunityReadOid

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"strings"
	"util"
)

func resourceSnmpServerCommunityReadOid() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerCommunityReadOidCreate,
		Update: resourceSnmpServerCommunityReadOidUpdate,
		Read:   resourceSnmpServerCommunityReadOidRead,
		Delete: resourceSnmpServerCommunityReadOidDelete,
		Schema: map[string]*schema.Schema{
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
			"read_user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerCommunityReadOidCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerCommunityReadOid (Inside resourceSnmpServerCommunityReadOidCreate) ")
		name2 := d.Get("oid-val").(string)
		name1 := d.Get("user").(string)
		data := dataToSnmpServerCommunityReadOid(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerCommunityReadOid --")
		d.SetId(name1 + "," + name2)
		go_thunder.PostSnmpServerCommunityReadOid(client.Token, name1, data, client.Host)

		return resourceSnmpServerCommunityReadOidRead(d, meta)

	}
	return nil
}

func resourceSnmpServerCommunityReadOidRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerCommunityReadOid (Inside resourceSnmpServerCommunityReadOidRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerCommunityReadOid(client.Token, name1, name2, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerCommunityReadOidUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying SnmpServerCommunityReadOid   (Inside resourceSnmpServerCommunityReadOidUpdate) ")
		data := dataToSnmpServerCommunityReadOid(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerCommunityReadOid ")
		go_thunder.PutSnmpServerCommunityReadOid(client.Token, name1, name2, data, client.Host)

		return resourceSnmpServerCommunityReadOidRead(d, meta)

	}
	return nil
}

func resourceSnmpServerCommunityReadOidDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerCommunityReadOidDelete) " + name1)
		err := go_thunder.DeleteSnmpServerCommunityReadOid(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToSnmpServerCommunityReadOid(d *schema.ResourceData) go_thunder.SnmpServerCommunityReadOid {
	var vc go_thunder.SnmpServerCommunityReadOid
	var c go_thunder.SnmpServerCommunityReadOidInstance

	var obj1 go_thunder.SnmpServerCommunityReadRemote
	prefix1 := "remote.0."

	HostListCount := d.Get(prefix1 + "host_list.#").(int)
	obj1.DNSHost = make([]go_thunder.SnmpServerCommunityReadHostList, 0, HostListCount)

	for i := 0; i < HostListCount; i++ {
		var obj1_1 go_thunder.SnmpServerCommunityReadHostList
		prefix1_1 := prefix1 + fmt.Sprintf("host_list.%d.", i)
		obj1_1.DNSHost = d.Get(prefix1_1 + "dns_host").(string)
		obj1_1.Ipv4Mask = d.Get(prefix1_1 + "ipv4_mask").(string)
		obj1.DNSHost = append(obj1.DNSHost, obj1_1)
	}

	Ipv4ListCount := d.Get(prefix1 + "ipv4_list.#").(int)
	obj1.Ipv4Host = make([]go_thunder.SnmpServerCommunityReadIpv4List, 0, Ipv4ListCount)

	for i := 0; i < Ipv4ListCount; i++ {
		var obj1_2 go_thunder.SnmpServerCommunityReadIpv4List
		prefix1_2 := prefix1 + fmt.Sprintf("ipv4_list.%d.", i)
		obj1_2.Ipv4Host = d.Get(prefix1_2 + "ipv4_host").(string)
		obj1_2.Ipv4Mask = d.Get(prefix1_2 + "ipv4_mask").(string)
		obj1.Ipv4Host = append(obj1.Ipv4Host, obj1_2)
	}

	Ipv6ListCount := d.Get(prefix1 + "ipv6_list.#").(int)
	obj1.Ipv6Host = make([]go_thunder.SnmpServerCommunityReadIpv6List, 0, Ipv6ListCount)

	for i := 0; i < Ipv6ListCount; i++ {
		var obj1_3 go_thunder.SnmpServerCommunityReadIpv6List
		prefix1_3 := prefix1 + fmt.Sprintf("ipv6_list.%d.", i)
		obj1_3.Ipv6Host = d.Get(prefix1_3 + "ipv6_host").(string)
		obj1_3.Ipv6Mask = d.Get(prefix1_3 + "ipv6_mask").(int)
		obj1.Ipv6Host = append(obj1.Ipv6Host, obj1_3)
	}

	c.HostList = obj1

	c.OidVal = d.Get("oid_val").(string)
	c.UserTag = d.Get("user_tag").(string)

	vc.Remote = c
	return vc
}
