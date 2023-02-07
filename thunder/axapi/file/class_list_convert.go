package file

//pseudo endpoint for upload common class-list files

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi/oper"
)

type ClassListConvert struct {
	Name          string
	ClassListType string
	Protocol      string
	Host          string
	Path          string
	Username      string
	Password      string
	UseMgmtPort   int
	Overwrite     int
}

func (p *ClassListConvert) GetId() string {
	return p.Name
}

func (p *ClassListConvert) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("file.ClassListConvert::Post")

	realInst := edpt.Import{}
	realInst.Inst.ClassListConvert = p.Name
	realInst.Inst.ClassListType = p.ClassListType
	realInst.Inst.UseMgmtPort = p.UseMgmtPort
	realInst.Inst.Overwrite = p.Overwrite
	if len(p.Password) > 0 {
		realInst.Inst.Password = p.Password
	}
	realInst.Inst.RemoteFile = axapi.GenUrlForDownload(p.Protocol, p.Host, p.Path, p.Username)

	return realInst.Post(authToken, host, logger)
}

func (p *ClassListConvert) Get(authToken string, host, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.ClassListConvert::Get")
	err := oper.FileClassListGet(authToken, host, instId, logger)
	return err
}

func (p *ClassListConvert) Delete(authToken string, host, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.ClassListConvert::Delete")
	obj := edpt.ClassList{}
	err := obj.Delete(authToken, host, instId, logger)
	return err
}
