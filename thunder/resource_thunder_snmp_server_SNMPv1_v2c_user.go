package thunder

//Thunder resource SnmpServerSNMPv1V2cUser

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"util"
)

func resourceSnmpServerSNMPv1V2cUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerSNMPv1V2cUserCreate,
		Update: resourceSnmpServerSNMPv1V2cUserUpdate,
		Read:   resourceSnmpServerSNMPv1V2cUserRead,
		Delete: resourceSnmpServerSNMPv1V2cUserDelete,
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"passwd": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"oid_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
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
					},
				},
			},
		},
	}
}

func resourceSnmpServerSNMPv1V2cUserCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerSNMPv1V2cUser (Inside resourceSnmpServerSNMPv1V2cUserCreate) ")
		name1 := d.Get("user").(string)
		data := dataToSnmpServerSNMPv1V2cUser(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerSNMPv1V2cUser --")
		d.SetId(name1)
		go_thunder.PostSnmpServerSNMPv1V2cUser(client.Token, data, client.Host)

		return resourceSnmpServerSNMPv1V2cUserRead(d, meta)

	}
	return nil
}

func resourceSnmpServerSNMPv1V2cUserRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerSNMPv1V2cUser (Inside resourceSnmpServerSNMPv1V2cUserRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerSNMPv1V2cUser(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerSNMPv1V2cUserUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SnmpServerSNMPv1V2cUser   (Inside resourceSnmpServerSNMPv1V2cUserUpdate) ")
		data := dataToSnmpServerSNMPv1V2cUser(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerSNMPv1V2cUser ")
		go_thunder.PutSnmpServerSNMPv1V2cUser(client.Token, name1, data, client.Host)

		return resourceSnmpServerSNMPv1V2cUserRead(d, meta)

	}
	return nil
}

func resourceSnmpServerSNMPv1V2cUserDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerSNMPv1V2cUserDelete) " + name1)
		err := go_thunder.DeleteSnmpServerSNMPv1V2cUser(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToSnmpServerSNMPv1V2cUser(d *schema.ResourceData) go_thunder.SnmpServerSNMPv1V2cUser {
	var vc go_thunder.SnmpServerSNMPv1V2cUser
	var c go_thunder.SnmpServerSNMPv1V2cUserInstance

	var obj1 go_thunder.SnmpServerSNMPv1V2cRemote
	prefix := "remote.0."

	HostListCount := d.Get(prefix + "host_list.#").(int)
	obj1.DNSHost = make([]go_thunder.SnmpServerSNMPv1V2cHostList, 0, HostListCount)

	for i := 0; i < HostListCount; i++ {
		var obj1_1 go_thunder.SnmpServerSNMPv1V2cHostList
		prefix1 := prefix + fmt.Sprintf("host_list.%d.", i)
		obj1_1.DNSHost = d.Get(prefix1 + "dns_host").(string)
		obj1_1.Ipv4Mask = d.Get(prefix1 + "ipv4_mask").(string)
		obj1.DNSHost = append(obj1.DNSHost, obj1_1)
	}

	Ipv4ListCount := d.Get(prefix + "ipv4_list.#").(int)
	obj1.Ipv4Host = make([]go_thunder.SnmpServerSNMPv1V2cIpv4List, 0, Ipv4ListCount)

	for i := 0; i < Ipv4ListCount; i++ {
		var obj1_2 go_thunder.SnmpServerSNMPv1V2cIpv4List
		prefix1_2 := prefix + fmt.Sprintf("ipv4_list.%d.", i)
		obj1_2.Ipv4Host = d.Get(prefix1_2 + "ipv4_host").(string)
		obj1_2.Ipv4Mask = d.Get(prefix1_2 + "ipv4_mask").(string)
		obj1.Ipv4Host = append(obj1.Ipv4Host, obj1_2)
	}

	Ipv6ListCount := d.Get(prefix + "ipv6_list.#").(int)
	obj1.Ipv6Host = make([]go_thunder.SnmpServerSNMPv1V2cIpv6List, 0, Ipv6ListCount)

	for i := 0; i < Ipv6ListCount; i++ {
		var obj1_3 go_thunder.SnmpServerSNMPv1V2cIpv6List
		prefix1_3 := prefix + fmt.Sprintf("ipv6_list.%d.", i)
		obj1_3.Ipv6Host = d.Get(prefix1_3 + "ipv6_host").(string)
		obj1_3.Ipv6Mask = d.Get(prefix1_3 + "ipv6_mask").(int)
		obj1.Ipv6Host = append(obj1.Ipv6Host, obj1_3)
	}

	c.HostList = obj1

	c.Passwd = d.Get("passwd").(string)
	c.Encrypted = d.Get("encrypted").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.User = d.Get("user").(string)

	OidListCount := d.Get("oid_list.#").(int)
	c.OidVal = make([]go_thunder.SnmpServerSNMPv1V2cOidList, 0, OidListCount)

	for i := 0; i < OidListCount; i++ {
		var obj2 go_thunder.SnmpServerSNMPv1V2cOidList
		prefix2 := fmt.Sprintf("oid_list.%d.", i)

		var obj2_1 go_thunder.SnmpServerSNMPv1V2cRemote
		prefix2_1 := prefix2 + "remote.0."

		HostListCount := d.Get(prefix2_1 + "host_list.#").(int)
		obj2_1.DNSHost = make([]go_thunder.SnmpServerSNMPv1V2cHostList, 0, HostListCount)

		for i := 0; i < HostListCount; i++ {
			var obj2_1_1 go_thunder.SnmpServerSNMPv1V2cHostList
			prefix2_1_1 := prefix2_1 + fmt.Sprintf("host_list.%d.", i)
			obj2_1_1.DNSHost = d.Get(prefix2_1_1 + "dns_host").(string)
			obj2_1_1.Ipv4Mask = d.Get(prefix2_1_1 + "ipv4_mask").(string)
			obj2_1.DNSHost = append(obj2_1.DNSHost, obj2_1_1)
		}

		Ipv4ListCount := d.Get(prefix2_1 + "ipv4_list.#").(int)
		obj2_1.Ipv4Host = make([]go_thunder.SnmpServerSNMPv1V2cIpv4List, 0, Ipv4ListCount)

		for i := 0; i < Ipv4ListCount; i++ {
			var obj2_1_2 go_thunder.SnmpServerSNMPv1V2cIpv4List
			prefix2_1_2 := prefix2_1 + fmt.Sprintf("ipv4_list.%d.", i)
			obj2_1_2.Ipv4Host = d.Get(prefix2_1_2 + "ipv4_host").(string)
			obj2_1_2.Ipv4Mask = d.Get(prefix2_1_2 + "ipv4_mask").(string)
			obj2_1.Ipv4Host = append(obj2_1.Ipv4Host, obj2_1_2)
		}

		Ipv6ListCount := d.Get(prefix2_1 + "ipv6_list.#").(int)
		obj2_1.Ipv6Host = make([]go_thunder.SnmpServerSNMPv1V2cIpv6List, 0, Ipv6ListCount)

		for i := 0; i < Ipv6ListCount; i++ {
			var obj2_1_3 go_thunder.SnmpServerSNMPv1V2cIpv6List
			prefix2_1_3 := prefix2_1 + fmt.Sprintf("ipv6_list.%d.", i)
			obj2_1_3.Ipv6Host = d.Get(prefix2_1_3 + "ipv6_host").(string)
			obj2_1_3.Ipv6Mask = d.Get(prefix2_1_3 + "ipv6_mask").(int)
			obj2_1.Ipv6Host = append(obj2_1.Ipv6Host, obj2_1_3)
		}

		obj2.HostList = obj2_1

		obj2.OidVal = d.Get(prefix2 + "oid_val").(string)
		obj2.UserTag = d.Get(prefix2 + "user_tag").(string)
		c.OidVal = append(c.OidVal, obj2)
	}

	vc.User = c
	return vc
}
