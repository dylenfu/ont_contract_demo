package main

import (
	"fmt"
	"github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/common"
	cm "github.com/ontio/ontology/common"
)

var (
	sdk                = ontology_go_sdk.NewOntologySdk()
	wallet             *ontology_go_sdk.Wallet
	signer             *ontology_go_sdk.Account
	contractAddress, _ = cm.AddressFromHexString("4b1968e68924dd273022349717be7dcc4c2151d1")
)

func Init() {
	sdk.ClientMgr.NewRestClient().SetAddress("http://localhost:20334")
	wallet, _ = ontology_go_sdk.OpenWallet("/Users/dylen/software/ont/testmode/wallet.dat")
	signer, _ = wallet.GetAccountByIndex(1, []byte("111111"))
}

func main() {
	Init()

	did()

	// exec contract function
	// exec("put", "cat", 50)

	// query pre exec tx
	//res := query("get", "tom")
	//r, err := res.ToInteger()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(r.Uint64())
}

func query(method string, list ...interface{}) *common.ResultItem {
	var (
		res    *common.PreExecResult
		err    error
		params = []interface{}{}
	)

	params = append(params, list...)
	q := []interface{}{method, params}
	res, err = sdk.NeoVM.PreExecInvokeNeoVMContract(contractAddress, q)
	if err != nil {
		panic(err)
	}

	return res.Result
}

func exec(method string, list ...interface{}) {
	var (
		err    error
		params = []interface{}{}
		res    cm.Uint256
	)

	params = append(params, list...)
	query := []interface{}{method, params}

	res, err = sdk.NeoVM.InvokeNeoVMContract(0, 20000, signer, contractAddress, query)
	if err != nil {
		panic(err)
	}
	hash := res.ToHexString()
	fmt.Println(hash)
}

func notify(hash string) {
	var (
		event *common.SmartContactEvent
		err   error
	)

	if event, err = sdk.GetSmartContractEvent(hash); err != nil {
		panic(err)
	}

	for _, v := range event.Notify {
		fmt.Println(v.States)
	}
}

func did() {
	identity, _ := wallet.NewDefaultSettingIdentity([]byte("111111"))
	fmt.Println(identity.ID)                               // generated from [32]bytes
	fmt.Println(len("TUQnNzzGv1nfCpsNxda9aKSqwZYGDchUkf")) //  total length 42
}

func queryBlock() {
	height := uint32(7869)

	block, err := sdk.GetBlockByHeight(height)
	if err != nil {
		panic(err)
	}

	for _, tx := range block.Transactions {
		fmt.Println("version", uint8(tx.Version))
		fmt.Println("payer", tx.Payer.ToBase58())
		// payload 只能是transfer或者transferFrom
		//depolyCode,ok := tx.Payload.(*payload.DeployCode)
		//if !ok {
		//	panic("convert tx deploy code failed")
		//}
		//res, err := ontology_go_sdk.ParsePayload(depolyCode.Code)
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println(res)
	}
}

/*
	if tx, err = sdk.GetTransaction(hash); err != nil {
		panic(err)
	}

	invokeCode, ok := tx.Payload.(*payload.InvokeCode)
	if !ok {
		panic("convert invoke failed")
	}

	if res, err = ontology_go_sdk.ParsePayload(invokeCode.Code); err != nil {
		panic(err)
	}
	fmt.Println(res)

	if event, err = sdk.GetSmartContractEvent(hash); err != nil {
		panic(err)
	}
	fmt.Println(event)
*/
