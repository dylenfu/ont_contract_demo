package main

import (
	"fmt"
	"github.com/ontio/ontology-go-sdk"
	cm "github.com/ontio/ontology/common"
	"time"
)

// 所有invoke接口,signer必须是admin，合约的管理者，合约中有CheckWitness

var (
	sdk                = ontology_go_sdk.NewOntologySdk()
	wallet             *ontology_go_sdk.Wallet
	admin              *ontology_go_sdk.Account
	pwd                = []byte("111111")
	contractAddress, _ = cm.AddressFromHexString("13d970986b2f8ff4a76c317cf6ba3965bb19f12c")
)

func Init() {
	sdk.ClientMgr.NewRestClient().SetAddress("http://localhost:20334")
	wallet, _ = ontology_go_sdk.OpenWallet("/Users/dylen/software/ont/testmode/wallet.dat")
	admin, _ = wallet.GetAccountByIndex(2, pwd)
}

const (
	bucket1 = "bucket_1"
	bucket2 = "bucket_2"
	bucket3 = "bucket_3"
	bucket4 = "bucket_4"

	did1 = "did_1"
	did2 = "did_2"
	did3 = "did_3"
	did4 = "did_4"

	avm = "5dc56b05322e302e306a00527ac41400000000000000000000000000000000000000006a51527ac41400000000000000000000000000000000000000016a52527ac41400000000000000000000000000000000000000026a53527ac42241476f43314c70744e3475524672796d57377037547533644652375971325242325068204f6e746f6c6f67792e52756e74696d652e426173653538546f416464726573736a54527ac4682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e67536372697074486173686a55527ac4681953797374656d2e53746f726167652e476574436f6e746578746a56527ac4026b316a57527ac4036b325f6a58527ac4036b335f6a59527ac4006a5a527ac401016a5b527ac46c0119c56b6a00527ac46a51527ac46a52527ac46a51c304696e69747d9c7c75641500006a00c306f802000000006e6c756662a1016a51c30462696e647d9c7c75642f006a52c300c36a54527ac46a52c351c36a55527ac46a55c36a54c3526a00c3067a03000000006e6c75666266016a51c306756e62696e647d9c7c75642f006a52c300c36a54527ac46a52c351c36a55527ac46a55c36a54c3526a00c306d605000000006e6c75666229016a51c30a6765745f6275636b65747d9c7c756422006a52c300c36a54527ac46a54c3516a00c3064b07000000006e6c756662f5006a51c3096765745f6f776e65727d9c7c75641500006a00c3068507000000006e6c756662cf006a51c30572657365747d9c7c75642f006a52c300c36a54527ac46a52c351c36a55527ac46a55c36a54c3526a00c306b607000000006e6c75666293006a51c307757067726164657d9c7c756422006a52c300c36a56527ac46a56c3516a00c3063b09000000006e6c75666262006a51c30b6765745f6f6e745f6964737d9c7c756422006a52c300c36a55527ac46a55c3516a00c306120c000000006e6c7566622d006a51c30d636865636b5f6d6967726174657d9c7c75641500006a00c306630c000000006e6c7566620300006c756655c56b6a00527ac46a51527ac46203006a00c354c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a00c354c36a00c357c36a00c356c3681253797374656d2e53746f726167652e50757404696e6974077375636365737352c176c9681553797374656d2e52756e74696d652e4e6f74696679516c75660115c56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a00c354c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a00c358c36a52c37e6a00c356c3681253797374656d2e53746f726167652e4765746a54527ac46a54c36a53c37d9c7c756426001e6f6e742069642062696e6420746f207468652073616d65206275636b6574f06203006a54c3916487006a53c3516a00c306960b000000006e6a56527ac4516a56c36a52c37bc46a53c36a00c358c36a52c37e6a00c356c3681253797374656d2e53746f726167652e5075746a56c3681853797374656d2e52756e74696d652e53657269616c697a656a00c359c36a53c37e6a00c356c3681253797374656d2e53746f726167652e5075746215016a00c359c36a54c37e6a00c356c3681253797374656d2e53746f726167652e4765746a57527ac46a57c3681a53797374656d2e52756e74696d652e446573657269616c697a656a58527ac46a58c36a52c3ca6a58c3681853797374656d2e52756e74696d652e53657269616c697a656a00c359c36a54c37e6a00c356c3681253797374656d2e53746f726167652e5075746a53c3516a00c306960b000000006e6a56527ac4516a56c36a52c37bc46a53c36a00c358c36a52c37e6a00c356c3681253797374656d2e53746f726167652e5075746a56c3681853797374656d2e52756e74696d652e53657269616c697a656a00c359c36a53c37e6a00c356c3681253797374656d2e53746f726167652e5075740462696e646a52c36a53c353c176c9681553797374656d2e52756e74696d652e4e6f74696679516c75665ec56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a00c354c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a00c358c36a52c37e6a00c356c3681253797374656d2e53746f726167652e4765746a54527ac46a54c391642000186f6e742069642062696e642077697468206e6f7468696e67f06203006a54c36a53c37d9c7c75f16a00c359c36a53c37e6a00c356c3681253797374656d2e53746f726167652e4765746a55527ac46a55c3681a53797374656d2e52756e74696d652e446573657269616c697a656a56527ac46a56c36a52c3ca6a56c3681853797374656d2e52756e74696d652e53657269616c697a656a00c359c36a53c37e6a00c356c3681253797374656d2e53746f726167652e5075746a00c358c36a52c37e6a00c356c3681553797374656d2e53746f726167652e44656c65746506756e62696e646a52c36a53c353c176c9681553797374656d2e52756e74696d652e4e6f74696679516c756657c56b6a00527ac46a51527ac46a52527ac46203006a00c358c36a52c37e6a00c356c3681253797374656d2e53746f726167652e4765746c756655c56b6a00527ac46a51527ac46203006a00c357c36a00c356c3681253797374656d2e53746f726167652e4765746c756660c56b6a00527ac46a51527ac46a52527ac46a53527ac4620300006a00c3068507000000006e681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a00c359c36a53c37e6a00c356c3681253797374656d2e53746f726167652e4765746a55527ac46a55c36a53c37d9e7c75f16a55c36460006a55c3516a00c306960b000000006e6a56527ac46a56c36a52c3ca6a56c3681853797374656d2e52756e74696d652e53657269616c697a656a00c359c36a55c37e6a00c356c3681253797374656d2e53746f726167652e5075746203006a53c3516a00c306960b000000006e6a57527ac4516a57c36a52c37bc46a57c3681853797374656d2e52756e74696d652e53657269616c697a656a00c359c36a53c37e6a00c356c3681253797374656d2e53746f726167652e5075746a53c36a00c358c36a52c37e6a00c356c3681253797374656d2e53746f726167652e5075740572657365746a52c36a53c353c176c9681553797374656d2e52756e74696d652e4e6f74696679516c75665fc56b6a00527ac46a51527ac46a52527ac46203006a00c354c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a00c355c351c66b6a00527ac46c0962616c616e63654f666a00c352c36a00c35ac368164f6e746f6c6f67792e4e61746976652e496e766f6b656a53527ac46a53c36a00c354c36a00c355c353c66b6a00527ac46a51527ac46a52527ac46c51c176c9087472616e736665726a00c352c36a00c35ac368164f6e746f6c6f67792e4e61746976652e496e766f6b656a54527ac46a54c36a00c35bc37d9e7c75642e0026636f6e74726163742075706772616465202d207472616e73666572206f6e74206661696c6564f06203006a00c355c351c66b6a00527ac46c0962616c616e63654f666a00c353c36a00c35ac368164f6e746f6c6f67792e4e61746976652e496e766f6b656a55527ac46a55c36a00c354c36a00c355c353c66b6a00527ac46a51527ac46a52527ac46c51c176c9087472616e736665726a00c353c36a00c35ac368164f6e746f6c6f67792e4e61746976652e496e766f6b656a56527ac46a56c36a00c35bc37d9e7c75642e0026636f6e74726163742075706772616465202d207472616e73666572206f6e67206661696c6564f06203000000000000006a52c368194f6e746f6c6f67792e436f6e74726163742e4d6967726174656a57527ac46a57c391642d0025636f6e74726163742075706772616465202d206d6967726174652061766d206661696c6564f06203000775706772616465077375636365737352c176c9681553797374656d2e52756e74696d652e4e6f74696679516c75665bc56b6a00527ac46a51527ac46a52527ac46203006a00c359c36a52c37e6a00c356c3681253797374656d2e53746f726167652e4765746a53527ac4006a54527ac46a53c391640c00c76a54527ac46227006a53c3681a53797374656d2e52756e74696d652e446573657269616c697a656a54527ac46a54c36c756658c56b6a00527ac46a51527ac46a52527ac46203006a52c3516a00c306960b000000006e6a54527ac46a54c3681853797374656d2e52756e74696d652e53657269616c697a656a55527ac46a55c36c756655c56b6a00527ac46a51527ac462030001756c7566"
)

func main() {
	Init()

	//SetInit()
	//
	//SetBind(did1, bucket1)
	//GetBucket(did1)
	//
	//SetBind(did2, bucket2)
	//GetBucket(did2)
	//
	//SetBind(did3, bucket3)
	//GetBucket(did3)
	//
	//SetBind(did4, bucket3)
	//GetBucket(did4)

	//SetUnbind(did4, bucket4)
	SetUnbind(did4, bucket3)
	//GetBucket(did4)

	//SetReset(did4, bucket4)
	//GetBucket(did4)

	//GetOwner()

	//Migrate(avm)
	ShowAllBuckets()
	CheckMigrate()
}

// contract init
func SetInit() {
	args := assembleCallArgs("init")
	res, _ := sdk.NeoVM.InvokeNeoVMContract(0, 20000, admin, contractAddress, args)
	_, _ = sdk.WaitForGenerateBlock(6*time.Second, 1)

	hash := res.ToHexString()
	fmt.Println(hash)

	event, _ := sdk.GetSmartContractEvent(hash)

	v := event.Notify[0]
	slice := v.States.([]interface{})

	bs, _ := cm.HexToBytes(slice[0].(string))
	method := string(bs)

	bs, _ = cm.HexToBytes(slice[1].(string))
	inf := string(bs)

	fmt.Println(method, inf)
}

func SetBind(did, bucket string) {
	args := assembleCallArgs("bind", did, bucket)
	res, _ := sdk.NeoVM.InvokeNeoVMContract(0, 20000, admin, contractAddress, args)
	_, _ = sdk.WaitForGenerateBlock(6*time.Second, 1)

	hash := res.ToHexString()
	fmt.Println(hash)
	event, _ := sdk.GetSmartContractEvent(hash)

	x := event.Notify[0]
	slice := x.States.([]interface{})

	bs, _ := cm.HexToBytes(slice[0].(string))
	method := string(bs)

	bs, _ = cm.HexToBytes(slice[1].(string))
	did = string(bs)

	bs, _ = cm.HexToBytes(slice[2].(string))
	bucket = string(bs)

	fmt.Println(method, did, bucket)
}

func SetUnbind(did, bucket string) {
	args := assembleCallArgs("unbind", did, bucket)
	res, err := sdk.NeoVM.InvokeNeoVMContract(0, 20000, admin, contractAddress, args)
	if err != nil {
		panic(err)
	}
	_, _ = sdk.WaitForGenerateBlock(6*time.Second, 1)

	hash := res.ToHexString()
	fmt.Println(hash)
	event, _ := sdk.GetSmartContractEvent(hash)

	x := event.Notify[0]
	slice := x.States.([]interface{})

	bs, _ := cm.HexToBytes(slice[0].(string))
	method := string(bs)

	bs, _ = cm.HexToBytes(slice[1].(string))
	did = string(bs)

	bs, _ = cm.HexToBytes(slice[2].(string))
	bucket = string(bs)

	fmt.Println(method, did, bucket)
}

func SetReset(did, bucket string) {
	args := assembleCallArgs("reset", did, bucket)
	res, _ := sdk.NeoVM.InvokeNeoVMContract(0, 20000, admin, contractAddress, args)
	_, _ = sdk.WaitForGenerateBlock(6*time.Second, 1)

	hash := res.ToHexString()
	fmt.Println(hash)
	event, _ := sdk.GetSmartContractEvent(hash)

	x := event.Notify[0]
	slice := x.States.([]interface{})

	bs, _ := cm.HexToBytes(slice[0].(string))
	method := string(bs)

	bs, _ = cm.HexToBytes(slice[1].(string))
	did = string(bs)

	bs, _ = cm.HexToBytes(slice[2].(string))
	bucket = string(bs)

	fmt.Println(method, did, bucket)
}

func GetBucket(ontId string) {
	args := assembleCallArgs("get_bucket", ontId)
	res, _ := sdk.NeoVM.PreExecInvokeNeoVMContract(contractAddress, args)
	bs, _ := res.Result.ToByteArray()
	bucket := string(bs)
	fmt.Println(bucket)
}

func GetOntIdList(bucket string) {
	args := assembleCallArgs("get_ont_ids", bucket)
	res, err := sdk.NeoVM.PreExecInvokeNeoVMContract(contractAddress, args)
	if err != nil {
		panic(err)
	}
	bs, _ := res.Result.ToByteArray()
	fmt.Println(string(bs))
}

func GetOwner() {
	args := assembleCallArgs("get_owner")
	res, _ := sdk.NeoVM.PreExecInvokeNeoVMContract(contractAddress, args)
	bs, _ := res.Result.ToByteArray()
	addr, _ := cm.AddressParseFromBytes(bs)
	fmt.Println(addr.ToBase58())
}

func Migrate(str string) {
	code, _ := cm.HexToBytes(str)
	args := assembleCallArgs("upgrade", code)
	res, _ := sdk.NeoVM.InvokeNeoVMContract(0, 30000000, admin, contractAddress, args)
	_, _ = sdk.WaitForGenerateBlock(6*time.Second, 1)

	hash := res.ToHexString()
	fmt.Println(hash)

	event, err := sdk.GetSmartContractEvent(hash)
	if err != nil {
		panic(err)
	}

	x := event.Notify[0]
	slice := x.States.([]interface{})

	bs, _ := cm.HexToBytes(slice[0].(string))
	method := string(bs)

	bs, _ = cm.HexToBytes(slice[1].(string))
	inf := string(bs)

	fmt.Println(method, inf)
}

func CheckMigrate() {
	args := assembleCallArgs("check_migrate")
	res, _ := sdk.NeoVM.PreExecInvokeNeoVMContract(contractAddress, args)
	num, _ := res.Result.ToInteger()
	fmt.Println(num.Uint64())
}

func ShowAllBuckets() {
	fmt.Println("check after migrate...")
	GetBucket(did1)
	GetBucket(did2)
	GetBucket(did3)
	GetBucket(did4)

	fmt.Println("----bucket1")
	GetOntIdList(bucket1)
	fmt.Println("----bucket2")
	GetOntIdList(bucket2)
	fmt.Println("----bucket3")
	GetOntIdList(bucket3)
	fmt.Println("----bucket4")
	GetOntIdList(bucket4)
}

func assembleCallArgs(method string, list ...interface{}) []interface{} {
	params := []interface{}{}
	params = append(params, list...)
	args := []interface{}{method, params}

	return args
}
