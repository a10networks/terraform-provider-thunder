package file

//pseudo endpoint for upload and delete class-list files

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi/oper"
)

type ClassList struct {
	Name        string
	Protocol    string
	Host        string
	Path        string
	Username    string
	Password    string
	UseMgmtPort int
	Overwrite   int
}

func (p *ClassList) GetId() string {
	return p.Name
}

func (p *ClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("file.ClassList::Post")

	realInst := edpt.Import{}
	realInst.Inst.ClassList = p.Name
	realInst.Inst.UseMgmtPort = p.UseMgmtPort
	realInst.Inst.Overwrite = p.Overwrite
	if len(p.Password) > 0 {
		realInst.Inst.Password = p.Password
	}
	realInst.Inst.RemoteFile = axapi.GenUrlForDownload(p.Protocol, p.Host, p.Path, p.Username)

	return realInst.Post(authToken, host, logger)
}

func (p *ClassList) Get(authToken string, host, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.ClassList::Get")
	err := oper.FileClassListGet(authToken, host, instId, logger)
	return err
}

func (p *ClassList) Delete(authToken string, host, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.ClassList::Delete")
	obj := edpt.ClassList{}
	err := obj.Delete(authToken, host, instId, logger)
	return err
}
