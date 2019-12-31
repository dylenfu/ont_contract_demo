package main

import (
	"fmt"
	"github.com/ontio/ontology-go-sdk"
	cm "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"time"
)

var (
	sdk                                           = ontology_go_sdk.NewOntologySdk()
	wallet                                        *ontology_go_sdk.Wallet
	signer                                        *ontology_go_sdk.Account
	admin, account3, account4, account5, account6 *ontology_go_sdk.Account
	defaultIdentity                               *ontology_go_sdk.Identity
	pwd                                           = []byte("111111")
	contractAddress, _                            = cm.AddressFromHexString("c827b5738e7d50d55d8cab88b1aba4f949c7bc7c") //"8d6e81e43cf97fcb1ec7ed8e32f61cac9c4ed27a")
)

func Init() {
	sdk.ClientMgr.NewRestClient().SetAddress("http://localhost:20334")
	wallet, _ = ontology_go_sdk.OpenWallet("/Users/dylen/software/ont/testmode/wallet.dat")
	signer, _ = wallet.GetAccountByIndex(1, pwd)
	admin, _ = wallet.GetAccountByIndex(2, pwd)
	account3, _ = wallet.GetAccountByIndex(3, pwd)
	account4, _ = wallet.GetAccountByIndex(4, pwd)
	account5, _ = wallet.GetAccountByIndex(5, pwd)
	account6, _ = wallet.GetAccountByIndex(6, pwd)
	defaultIdentity, _ = wallet.NewDefaultSettingIdentity(pwd)
}

const (
	did1 = "did:ont:TDBhNSD6LNRhBcQV9WW9X26AHGk9jAzz7w"
	did2 = "did:ont:TDBhNSD6LNRhBcQV9WW9X26AHGk9jAyy7w"
	did3 = "did:ont:TDBhNSD6LNRhBcQV9WW9X26AHGk9jAxx7w"
	did4 = "did:ont:TDBhNSD6LNRhBcQV9WW9X26AHGk9jAmm7w"

	avm = "5ec56b05322e302e306a00527ac41400000000000000000000000000000000000000006a51527ac41400000000000000000000000000000000000000016a52527ac41400000000000000000000000000000000000000026a53527ac42241476f43314c70744e3475524672796d57377037547533644652375971325242325068204f6e746f6c6f67792e52756e74696d652e426173653538546f416464726573736a54527ac4682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e67536372697074486173686a55527ac4681953797374656d2e53746f726167652e476574436f6e746578746a56527ac4026b316a57527ac4036b325f6a58527ac4036b335f6a59527ac4026b346a5a527ac4006a5b527ac401016a5c527ac46c011ac56b6a00527ac46a51527ac46a52527ac46a51c304696e69747d9c7c75641500006a00c3069a03000000006e6c756662c0016a51c30462696e647d9c7c75642f006a52c300c36a54527ac46a52c351c36a55527ac46a55c36a54c3526a00c3061c04000000006e6c75666285016a51c306756e62696e647d9c7c75642f006a52c300c36a54527ac46a52c351c36a55527ac46a55c36a54c3526a00c3063105000000006e6c75666248016a51c30a6765745f6f6e745f69647d9c7c756422006a52c300c36a55527ac46a55c3516a00c3064e06000000006e6c75666214016a51c30b6765745f6163636f756e747d9c7c756422006a52c300c36a54527ac46a54c3516a00c306bc06000000006e6c756662df006a51c3096765745f6f776e65727d9c7c75641500006a00c3063307000000006e6c756662b9006a51c3096765745f636f756e747d9c7c75641500006a00c3066407000000006e6c75666293006a51c30572657365747d9c7c75642f006a52c300c36a54527ac46a52c351c36a56527ac46a56c36a54c3526a00c306ad07000000006e6c75666257006a51c307757067726164657d9c7c756422006a52c300c36a57527ac46a57c3516a00c306a808000000006e6c75666226006a51c30673696d706c657d9c7c75641500006a00c3068503000000006e6c7566620300006c756658c56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a52c36a53c3936a54527ac46a54c36a52c37da27c75f16a54c36c756657c56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a52c36a53c37da27c75f16a52c36a53c3946c756655c56b6a00527ac46a51527ac462030001586c756655c56b6a00527ac46a51527ac46203006a00c354c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a00c354c36a00c357c36a00c356c3681253797374656d2e53746f726167652e50757404696e6974077375636365737352c176c9681553797374656d2e52756e74696d652e4e6f74696679516c75665bc56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a53c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a00c359c36a53c37e6a00c356c3681253797374656d2e53746f726167652e4765746a54527ac46a54c36a52c37d9c7c756427001f6163636f756e742062696e6420746f207468652073616d65206f6e74206964f06203006a52c36a00c359c36a53c37e6a00c356c3681253797374656d2e53746f726167652e5075746a53c36a00c358c36a52c37e6a00c356c3681253797374656d2e53746f726167652e50757451516a00c306030b000000006e750462696e646a52c36a53c353c176c9681553797374656d2e52756e74696d652e4e6f74696679516c75665bc56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a53c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a00c359c36a53c37e6a00c356c3681253797374656d2e53746f726167652e4765746a54527ac46a54c391642800206163636f756e74206e6f742062696e64207769746820616e79206f6e74206964f06203006a54c36a52c37d9c7c75f16a00c359c36a53c37e6a00c356c3681553797374656d2e53746f726167652e44656c6574656a00c358c36a52c37e6a00c356c3681553797374656d2e53746f726167652e44656c65746500516a00c306030b000000006e7506756e62696e646a52c36a53c353c176c9681553797374656d2e52756e74696d652e4e6f74696679516c756658c56b6a00527ac46a51527ac46a52527ac46203006a00c359c36a52c37e6a00c356c3681253797374656d2e53746f726167652e4765746a53527ac46a53c391642800206163636f756e74206e6f742062696e64207769746820616e79206f6e74206964f06203006a53c36c756658c56b6a00527ac46a51527ac46a52527ac46203006a00c358c36a52c37e6a00c356c3681253797374656d2e53746f726167652e4765746a53527ac46a53c3916431002963616e206e6f742066696e6420616e79206163636f756e742062696e642077697468206f6e74206964f06203006a53c36c756655c56b6a00527ac46a51527ac46203006a00c357c36a00c356c3681253797374656d2e53746f726167652e4765746c756657c56b6a00527ac46a51527ac46203006a00c35ac36a00c356c3681253797374656d2e53746f726167652e4765746a52527ac46a52c391640c00006a52527ac46203006a52c36c75665cc56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a53c3516a00c306bf0b000000006ef1006a00c3063307000000006e681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a00c359c36a53c37e6a00c356c3681253797374656d2e53746f726167652e4765746a55527ac46a55c36a52c37d9e7c75f16a52c36a00c359c36a53c37e6a00c356c3681253797374656d2e53746f726167652e5075746a53c36a00c358c36a52c37e6a00c356c3681253797374656d2e53746f726167652e5075740572657365746a52c36a53c353c176c9681553797374656d2e52756e74696d652e4e6f74696679516c75665fc56b6a00527ac46a51527ac46a52527ac46203006a00c354c3681b53797374656d2e52756e74696d652e436865636b5769746e657373f16a00c355c351c66b6a00527ac46c0962616c616e63654f666a00c352c36a00c35bc368164f6e746f6c6f67792e4e61746976652e496e766f6b656a53527ac46a53c36a00c354c36a00c355c353c66b6a00527ac46a51527ac46a52527ac46c51c176c9087472616e736665726a00c352c36a00c35bc368164f6e746f6c6f67792e4e61746976652e496e766f6b656a54527ac46a54c36a00c35cc37d9e7c75642e0026636f6e74726163742075706772616465202d207472616e73666572206f6e74206661696c6564f06203006a00c355c351c66b6a00527ac46c0962616c616e63654f666a00c353c36a00c35bc368164f6e746f6c6f67792e4e61746976652e496e766f6b656a55527ac46a55c36a00c354c36a00c355c353c66b6a00527ac46a51527ac46a52527ac46c51c176c9087472616e736665726a00c353c36a00c35bc368164f6e746f6c6f67792e4e61746976652e496e766f6b656a56527ac46a56c36a00c35cc37d9e7c75642e0026636f6e74726163742075706772616465202d207472616e73666572206f6e67206661696c6564f06203000000000000006a52c368194f6e746f6c6f67792e436f6e74726163742e4d6967726174656a57527ac46a57c391642d0025636f6e74726163742075706772616465202d206d6967726174652061766d206661696c6564f06203000775706772616465077375636365737352c176c9681553797374656d2e52756e74696d652e4e6f74696679516c75665bc56b6a00527ac46a51527ac46a52527ac46203006a00c35ac36a00c356c3681253797374656d2e53746f726167652e4765746a53527ac46a53c391640c00006a53527ac46203006a52c3517d877c75641b00516a53c3526a00c3061f03000000006e6a53527ac4621800516a53c3526a00c3065603000000006e6a53527ac46a52c36a53c3526a00c3061f03000000006e6a53527ac46a53c36a00c35ac36a00c356c3681253797374656d2e53746f726167652e507574516c756656c56b6a00527ac46a51527ac46a52527ac46203006a52c3c001147d9c7c7576641000756a52c36a00c351c37d9e7c75f1516c7566"
)

func main() {
	Init()

	//SetInit()

	//SetBind(account3, did1)
	//GetOntID(account3.Address)
	//GetAccount(did1)
	//
	//SetBind(account4, did2)
	//GetOntID(account4.Address)
	//GetAccount(did2)
	//
	//SetBind(account5, did3)
	//GetOntID(account5.Address)
	//GetAccount(did3)
	//
	//SetBind(account6, did4)
	//GetOntID(account6.Address)
	//GetAccount(did4)
	//
	//SetUnbind(account6, did4)
	//
	//SetReset(account5, did3)
	//GetOntID(account5.Address)
	//
	//GetOwner()

	//fmt.Println("before migrate...")
	//GetCount()
	//CheckMigrate()
	//
	//Migrate(avm)

	fmt.Println("after migrate...")
	CheckMigrate()
	GetCount()

	GetOntID(account3.Address)
	GetOntID(account4.Address)
}

func Deposit() { deposit() }

// contract init
func SetInit() {
	args := assembleCallArgs("init")
	// 这里参数signer必须是admin，合约的管理者，合约中有CheckWitness
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

	fmt.Println(hash, method, inf)
}

func SetBind(acc *ontology_go_sdk.Account, did string) {
	addr := acc.Address
	args := assembleCallArgs("bind", did, addr)
	res, _ := sdk.NeoVM.InvokeNeoVMContract(0, 20000, acc, contractAddress, args)
	_, _ = sdk.WaitForGenerateBlock(6*time.Second, 1)

	hash := res.ToHexString()
	fmt.Println(hash)
	event, _ := sdk.GetSmartContractEvent(hash)

	x := event.Notify[0]
	slice := x.States.([]interface{})

	bs, _ := cm.HexToBytes(slice[0].(string))
	method := string(bs)

	bs, _ = cm.HexToBytes(slice[1].(string))
	ontid := string(bs)

	bs, _ = cm.HexToBytes(slice[2].(string))
	_addr, _ := cm.AddressParseFromBytes(bs)

	fmt.Println(method, ontid, _addr.ToBase58())
}

func SetUnbind(acc *ontology_go_sdk.Account, did string) {
	addr := acc.Address
	args := assembleCallArgs("unbind", did, addr)
	res, err := sdk.NeoVM.InvokeNeoVMContract(0, 20000, acc, contractAddress, args)
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
	ontid := string(bs)

	bs, _ = cm.HexToBytes(slice[2].(string))
	_addr, _ := cm.AddressParseFromBytes(bs)

	fmt.Println(method, ontid, _addr.ToBase58())
}

func SetReset(acc *ontology_go_sdk.Account, did string) {
	addr := acc.Address
	args := assembleCallArgs("reset", did, addr)
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
	ontid := string(bs)

	bs, _ = cm.HexToBytes(slice[2].(string))
	_addr, _ := cm.AddressParseFromBytes(bs)

	fmt.Println(method, ontid, _addr.ToBase58())
}

// contract get_ont_id
func GetOntID(add cm.Address) {
	args := assembleCallArgs("get_ont_id", add)
	res, _ := sdk.NeoVM.PreExecInvokeNeoVMContract(contractAddress, args)
	ontId, _ := res.Result.ToString()
	fmt.Println(ontId)
}

func GetAccount(ontid string) {
	args := assembleCallArgs("get_account", ontid)
	res, _ := sdk.NeoVM.PreExecInvokeNeoVMContract(contractAddress, args)
	bs, _ := res.Result.ToByteArray()
	addr, _ := cm.AddressParseFromBytes(bs)
	fmt.Println(addr.ToBase58())
}

func GetOwner() {
	args := assembleCallArgs("get_owner")
	res, _ := sdk.NeoVM.PreExecInvokeNeoVMContract(contractAddress, args)
	bs, _ := res.Result.ToByteArray()
	addr, _ := cm.AddressParseFromBytes(bs)
	fmt.Println(addr.ToBase58())
}

func GetCount() {
	args := assembleCallArgs("get_count")
	res, _ := sdk.NeoVM.PreExecInvokeNeoVMContract(contractAddress, args)
	num, _ := res.Result.ToInteger()
	fmt.Println(num.Uint64())
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
	args := assembleCallArgs("simple")
	res, err := sdk.NeoVM.PreExecInvokeNeoVMContract(contractAddress, args)
	if err != nil {
		panic(err)
	}
	num, err := res.Result.ToInteger()
	if err != nil {
		panic(err)
	}
	fmt.Println("num ", num.Uint64())
}

func assembleCallArgs(method string, list ...interface{}) []interface{} {
	params := []interface{}{}
	params = append(params, list...)
	args := []interface{}{method, params}

	return args
}

// signer deposit some ont coin to admin
func deposit() {
	var (
		res cm.Uint256
		err error
	)
	if res, err = sdk.Native.Ont.Transfer(0, 20000, signer, admin.Address, 90); err != nil {
		panic(err)
	}
	hash := res.ToHexString()
	fmt.Println(hash)
}

// 注册默认地址ontid
func regOntId(passphrase string) {
	pin := []byte(passphrase)

	//scrypt := keypair.GetScryptParameters()
	//identity, err := ontology_go_sdk.NewIdentity(scrypt)
	identity, err := wallet.NewDefaultSettingIdentity(pin)
	if err != nil {
		panic(err)
	}

	controller, err := identity.GetControllerByIndex(1, pin)
	if err != nil {
		panic(err)
	}
	sender := ontology_go_sdk.NewAccount(controller.SigScheme)

	pubkey := controller.GetPublicKey()
	address := types.AddressFromPubKey(pubkey)

	pubkey1 := sender.GetPublicKey()
	address1 := types.AddressFromPubKey(pubkey1)
	fmt.Println("address", address.ToHexString(), address1.ToHexString())

	res, err := sdk.Native.OntId.RegIDWithPublicKey(0, 20000, sender, controller.ID, controller)
	if err != nil {
		panic(err)
	}

	hash := res.ToHexString()
	_, _ = sdk.WaitForGenerateBlock(3*time.Second, 1)

	getRegOntId(hash)
}

// 查询注册ontid时发送的事件
func getRegOntId(hash string) string {
	evt, err := sdk.GetSmartContractEvent(hash)
	if err != nil {
		panic(err)
	}

	notify := evt.Notify[0]
	slice := notify.States.([]interface{})
	method := slice[0].(string)
	ontid := slice[1].(string)
	fmt.Println(method, ontid)

	return ontid
}
