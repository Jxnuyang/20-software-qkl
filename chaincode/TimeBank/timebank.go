package main

import (
	"fmt"
	"time"

	"github.com/Jxnuyang/20-software-qkl/chaincode/TimeBank/lib"
	"github.com/Jxnuyang/20-software-qkl/chaincode/TimeBank/routers"
	"github.com/Jxnuyang/20-software-qkl/chaincode/TimeBank/utils"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	//"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	//"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/routers"
	//"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
)

type BlockChainTimeBank struct {
}

// Init 链码初始化
func (t *BlockChainTimeBank) Init(stub shim.ChaincodeStubInterface) peer.Response {
	//fmt.Sprintf("Init chaincode start ...")
	timeLocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return shim.Error(fmt.Sprintf("时区设置失败%s", err))
	}
	time.Local = timeLocal
	//初始化默认数据
	var managerID = "666666"
	var managerAsset float64 = 100
	manager := &lib.Manager{
		ManagerID:    managerID,
		ManagerAsset: managerAsset,
	}
	// 写入账本
	utils.WriteLedger(manager, stub, lib.ManagerKey, []string{managerID})
	var orgids = [3]string{
		"330027", //青山湖校区邮政编码
		"330022", //瑶湖校区邮政编码
		"332020", //共青城校区邮政编码
	}
	var orgNames = [3]string{"青山湖校区", "瑶湖校区", "共青城校区"}
	var userNums = [3]int{0, 0, 0}
	//var empty []string
	//初始化账号数据
	for i, val := range orgids {
		oragnization := &lib.Organization{
			OrgID:      val,
			OrgName:    orgNames[i],
			UserSum:    userNums[i],
			HaveUserID: nil,
		}
		// 写入账本
		utils.WriteLedger(oragnization, stub, lib.OrganizationKey, []string{val})
	}
	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainTimeBank) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	funcName, args := stub.GetFunctionAndParameters()
	//fmt.Sprintf("Invoke %s start ...", funcName)
	if funcName == "CreateUser" {
		return routers.CreateUser(stub, args)
	} else if funcName == "CreateService" {
		return routers.CreateService(stub, args)
	} else if funcName == "CreateOrg" {
		return routers.CreateOrg(stub, args)
	} else if funcName == "UserList" {
		return routers.UserList(stub, args)
	} else if funcName == "OrgList" {
		return routers.OrgList(stub, args)
	} else if funcName == "ManagerList" {
		return routers.ManagerList(stub, args)
	} else if funcName == "CreateServicing" {
		return routers.CreateServicing(stub, args)
	} else if funcName == "AcceptServicing" {
		return routers.AcceptServicing(stub, args)
	} else if funcName == "DoneServicing" {
		return routers.DoneServicing(stub, args)
	} else if funcName == "CloseServicing" {
		return routers.CloseServicing(stub, args)
	} else if funcName == "QueryServicingStatus" {
		return routers.QueryServicingStatus(stub, args)
	} else if funcName == "QueryServiceTrade" {
		return routers.QueryServiceTrade(stub, args)
	} else if funcName == "TransferAsset" {
		return routers.TransferAsset(stub, args)
	} else if funcName == "InheritAsset" {
		return routers.InheritAsset(stub, args)
	} else if funcName == "RechargeAsset" {
		return routers.RechargeAsset(stub, args)
	} else if funcName == "" {
		return routers.SpecailTradeList(stub, args)
	} else {
		return shim.Error("Invoke funcName error !!!")
	}
	//return shim.Success([]byte(""))
}

//启动并进入链码
func main() {
	err := shim.Start(new(BlockChainTimeBank))
	if err != nil {
		fmt.Printf("Chaincode start error %s", err)
	}
}
