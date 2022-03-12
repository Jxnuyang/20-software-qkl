package main

import (
	"encoding/json"
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

	//杨大爷
	userwang := &lib.User{
		UserID:             "01",
		UserName:           "杨大爷",
		UserIdentification: "110101199003076675",
		Sex:                "男",
		Birthday:           "1990-3-7",
		Address:            "江西师范大学青山湖校区",
		Postcode:           "330027",
		Ability:            []string{"做饭", "教学高数"},
		StarSign:           3,
		UserAsset:          float64(1000),
		Comment:            []string{"很细心", "Good!!!", "王老师教的高数就是好"},
		RecommenderID:      "",
	}
	var orgList lib.Organization
	org, err := utils.QueryLedger(stub, lib.OrganizationKey, []string{userwang.Postcode})
	if err != nil {
		return shim.Error("Failed to find the organization !!!")
	}
	_ = json.Unmarshal(org[0], &orgList)

	orgList.UserSum++
	orgList.HaveUserID = append(orgList.HaveUserID, userwang.UserID)

	_ = utils.WriteLedger(orgList, stub, lib.OrganizationKey, []string{userwang.Postcode})
	//吕婆婆
	userlv := &lib.User{
		UserID:             "02",
		UserName:           "吕婆婆",
		UserIdentification: "110101199006068985",
		Sex:                "女",
		Birthday:           "1990-6-6",
		Address:            "江西师范大学瑶湖校区",
		Postcode:           "330022",
		Ability:            []string{"做饭", "洗衣服", "教学化学"},
		StarSign:           5,
		UserAsset:          float64(6000),
		Comment:            []string{"很细心", "Very Good!!!", "太可爱了"},
		RecommenderID:      "",
	}
	org, err = utils.QueryLedger(stub, lib.OrganizationKey, []string{userlv.Postcode})
	if err != nil {
		return shim.Error("Failed to find the organization !!!")
	}
	//var orgList lib.Organization
	_ = json.Unmarshal(org[0], &orgList)

	orgList.UserSum++
	orgList.HaveUserID = append(orgList.HaveUserID, userlv.UserID)

	_ = utils.WriteLedger(orgList, stub, lib.OrganizationKey, []string{userwang.Postcode})
	_ = utils.WriteLedger(userwang, stub, lib.UserKey, []string{userwang.UserID})
	_ = utils.WriteLedger(userlv, stub, lib.UserKey, []string{userlv.UserID})
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
