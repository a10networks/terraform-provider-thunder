package file

//pseudo endpoint for upload and delete black-white list files

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi/oper"
)

type BwList struct {
	Name        string
	Protocol    string
	Host        string
	Path        string
	Username    string
	Password    string
	UseMgmtPort int
	Overwrite   int
}

func (p *BwList) GetId() string {
	return p.Name
}

func (p *BwList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("file.BwList::Post")

	realInst := edpt.Import{}
	realInst.Inst.BwList = p.Name
	realInst.Inst.UseMgmtPort = p.UseMgmtPort
	realInst.Inst.Overwrite = p.Overwrite
	if len(p.Password) > 0 {
		realInst.Inst.Password = p.Password
	}
	realInst.Inst.RemoteFile = axapi.GenUrlForDownload(p.Protocol, p.Host, p.Path, p.Username)

	return realInst.Post(authToken, host, logger)
}

func (p *BwList) Get(authToken string, host, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.BwList::Get")
	err := oper.FileBwListGet(authToken, host, instId, logger)
	return err
}

func (p *BwList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.BwList::Delet")
	obj := edpt.DeleteBwList{}
	obj.Inst.FileName = instId
	return obj.Post(authToken, host, logger)
}
