// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BuyFGE

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BuyFGEPayment is an auto generated low-level Go binding around an user-defined struct.
type BuyFGEPayment struct {
	PaymentId     []byte
	Paid          bool
	ActuallyPaid  *big.Int
	PayAmount     *big.Int
	Price         *big.Int
	OrderId       common.Address
	PaymentStatus []byte
	Fge           *big.Int
	Data          []byte
}

// BuyFGEABI is the input ABI used to generate the binding from.
const BuyFGEABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"payment_id\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"paid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"actually_paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pay_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"order_id\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payment_status\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structBuyFGE.Payment\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"SendFGE\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"Transaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"payment_id\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"paid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"actually_paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pay_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"order_id\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payment_status\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payment_id\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actually_paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pay_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"order_id\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payment_status\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fge\",\"type\":\"uint256\"}],\"name\":\"buyFGE\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payment_id\",\"type\":\"bytes\"}],\"name\":\"getTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"payment_id\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"paid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"actually_paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pay_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"order_id\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payment_status\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structBuyFGE.Payment\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"payment_id\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"paid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"actually_paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pay_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"order_id\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payment_status\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structBuyFGE.Payment\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"sendFGE\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// BuyFGEBin is the compiled bytecode used for deploying new contracts.
var BuyFGEBin = "0x60806040527fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4706001557f7b33a07f1644fab5bc0af6f7c06c8dceb370139ee7f9ae35b70dfee8eda3c95a6002557f8fbb8dad9e3865f36bbe2a27c9cb2cb4c4e3db9188c84707107220db2318a41860035534801561007c57600080fd5b506111358061008c6000396000f3fe60806040526004361061003f5760003560e01c8063032d64751461004457806378025a7714610059578063e54f992a14610079578063faa9e569146100b7575b600080fd5b610057610052366004610c63565b6100e4565b005b34801561006557600080fd5b50610057610074366004610df7565b610233565b34801561008557600080fd5b50610099610094366004610dba565b610638565b6040516100ae99989796959493929190610eee565b60405180910390f35b3480156100c357600080fd5b506100d76100d2366004610d78565b610836565b6040516100ae9190610f6d565b60008160e00151670de0b6b3a76400006100fe9190611055565b905060008260c001518051906020012090508260a001516001600160a01b0316846001600160a01b031614801561016f5750602083015115801561015357506002548160001a60f81b6001600160f81b031916145b8061016f57506003548160001a60f81b6001600160f81b031916145b1561022d57600080856001600160a01b03168460405160006040518083038185875af1925050503d80600081146101c2576040519150601f19603f3d011682016040523d82523d6000602084013e6101c7565b606091505b5091509150816102145760405162461bcd60e51b81526020600482015260146024820152732330b4b632b2103a379039b2b7321022ba3432b960611b604482015260640160405180910390fd5b811561022a576001602086015261010085018190525b50505b50505050565b60006040518061012001604052808a8152602001600015158152602001898152602001888152602001878152602001866001600160a01b0316815260200185858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509385525050506020808301869052604080518083018252928352909201528a51908b01209091507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470141561062d578060008a6040516103029190610ed2565b9081526020016040518091039020600082015181600001908051906020019061032c929190610ac8565b5060208281015160018301805460ff191691151591909117905560408301516002830155606083015160038301556080830151600483015560a08301516005830180546001600160a01b0319166001600160a01b0390921691909117905560c083015180516103a19260068501920190610ac8565b5060e0820151600782015561010082015180516103c8916008840191602090910190610ac8565b5090505061062d60008a6040516103df9190610ed2565b908152604051908190036020018120600501546001600160a01b03169060009061040a908d90610ed2565b908152602001604051809103902060405180610120016040529081600082018054610434906110ae565b80601f0160208091040260200160405190810160405280929190818152602001828054610460906110ae565b80156104ad5780601f10610482576101008083540402835291602001916104ad565b820191906000526020600020905b81548152906001019060200180831161049057829003601f168201915b5050509183525050600182015460ff161515602082015260028201546040820152600382015460608201526004820154608082015260058201546001600160a01b031660a082015260068201805460c09092019161050a906110ae565b80601f0160208091040260200160405190810160405280929190818152602001828054610536906110ae565b80156105835780601f1061055857610100808354040283529160200191610583565b820191906000526020600020905b81548152906001019060200180831161056657829003601f168201915b50505050508152602001600782015481526020016008820180546105a6906110ae565b80601f01602080910402602001604051908101604052809291908181526020018280546105d2906110ae565b801561061f5780601f106105f45761010080835404028352916020019161061f565b820191906000526020600020905b81548152906001019060200180831161060257829003601f168201915b5050505050815250506100e4565b505050505050505050565b805160208183018101805160008252928201919093012091528054819061065e906110ae565b80601f016020809104026020016040519081016040528092919081815260200182805461068a906110ae565b80156106d75780601f106106ac576101008083540402835291602001916106d7565b820191906000526020600020905b8154815290600101906020018083116106ba57829003601f168201915b5050505060018301546002840154600385015460048601546005870154600688018054979860ff90961697949650929491936001600160a01b03909116929061071f906110ae565b80601f016020809104026020016040519081016040528092919081815260200182805461074b906110ae565b80156107985780601f1061076d57610100808354040283529160200191610798565b820191906000526020600020905b81548152906001019060200180831161077b57829003601f168201915b5050505050908060070154908060080180546107b3906110ae565b80601f01602080910402602001604051908101604052809291908181526020018280546107df906110ae565b801561082c5780601f106108015761010080835404028352916020019161082c565b820191906000526020600020905b81548152906001019060200180831161080f57829003601f168201915b5050505050905089565b6108906040518061012001604052806060815260200160001515815260200160008152602001600081526020016000815260200160006001600160a01b031681526020016060815260200160008152602001606081525090565b600083836040516108a2929190610ec2565b9081526020016040518091039020604051806101200160405290816000820180546108cc906110ae565b80601f01602080910402602001604051908101604052809291908181526020018280546108f8906110ae565b80156109455780601f1061091a57610100808354040283529160200191610945565b820191906000526020600020905b81548152906001019060200180831161092857829003601f168201915b5050509183525050600182015460ff161515602082015260028201546040820152600382015460608201526004820154608082015260058201546001600160a01b031660a082015260068201805460c0909201916109a2906110ae565b80601f01602080910402602001604051908101604052809291908181526020018280546109ce906110ae565b8015610a1b5780601f106109f057610100808354040283529160200191610a1b565b820191906000526020600020905b8154815290600101906020018083116109fe57829003601f168201915b5050505050815260200160078201548152602001600882018054610a3e906110ae565b80601f0160208091040260200160405190810160405280929190818152602001828054610a6a906110ae565b8015610ab75780601f10610a8c57610100808354040283529160200191610ab7565b820191906000526020600020905b815481529060010190602001808311610a9a57829003601f168201915b505050505081525050905092915050565b828054610ad4906110ae565b90600052602060002090601f016020900481019282610af65760008555610b3c565b82601f10610b0f57805160ff1916838001178555610b3c565b82800160010185558215610b3c579182015b82811115610b3c578251825591602001919060010190610b21565b50610b48929150610b4c565b5090565b5b80821115610b485760008155600101610b4d565b80356001600160a01b0381168114610b7857600080fd5b919050565b80358015158114610b7857600080fd5b60008083601f840112610b9f57600080fd5b50813567ffffffffffffffff811115610bb757600080fd5b602083019150836020828501011115610bcf57600080fd5b9250929050565b600082601f830112610be757600080fd5b813567ffffffffffffffff80821115610c0257610c026110e9565b604051601f8301601f19908116603f01168101908282118183101715610c2a57610c2a6110e9565b81604052838152866020858801011115610c4357600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215610c7657600080fd5b610c7f83610b61565b9150602083013567ffffffffffffffff80821115610c9c57600080fd5b908401906101208287031215610cb157600080fd5b610cb961102b565b823582811115610cc857600080fd5b610cd488828601610bd6565b825250610ce360208401610b7d565b6020820152604083013560408201526060830135606082015260808301356080820152610d1260a08401610b61565b60a082015260c083013582811115610d2957600080fd5b610d3588828601610bd6565b60c08301525060e083013560e08201526101008084013583811115610d5957600080fd5b610d6589828701610bd6565b8284015250508093505050509250929050565b60008060208385031215610d8b57600080fd5b823567ffffffffffffffff811115610da257600080fd5b610dae85828601610b8d565b90969095509350505050565b600060208284031215610dcc57600080fd5b813567ffffffffffffffff811115610de357600080fd5b610def84828501610bd6565b949350505050565b60008060008060008060008060e0898b031215610e1357600080fd5b883567ffffffffffffffff80821115610e2b57600080fd5b610e378c838d01610bd6565b995060208b0135985060408b0135975060608b01359650610e5a60808c01610b61565b955060a08b0135915080821115610e7057600080fd5b50610e7d8b828c01610b8d565b999c989b50969995989497949560c00135949350505050565b60008151808452610eae816020860160208601611082565b601f01601f19169290920160200192915050565b8183823760009101908152919050565b60008251610ee4818460208701611082565b9190910192915050565b6000610120808352610f028184018d610e96565b8b15156020850152604084018b9052606084018a9052608084018990526001600160a01b03881660a085015283810360c08501529050610f428187610e96565b90508460e0840152828103610100840152610f5d8185610e96565b9c9b505050505050505050505050565b6020815260008251610120806020850152610f8c610140850183610e96565b91506020850151610fa1604086018215159052565b506040850151606085015260608501516080850152608085015160a085015260a0850151610fda60c08601826001600160a01b03169052565b5060c0850151601f19808685030160e0870152610ff78483610e96565b60e08801516101008881019190915288015187820390920184880152935090506110218382610e96565b9695505050505050565b604051610120810167ffffffffffffffff8111828210171561104f5761104f6110e9565b60405290565b600081600019048311821515161561107d57634e487b7160e01b600052601160045260246000fd5b500290565b60005b8381101561109d578181015183820152602001611085565b8381111561022d5750506000910152565b600181811c908216806110c257607f821691505b602082108114156110e357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea264697066735822122053ccca0710093a57d702e3b44b8f5b2854c89593e6524ba74a4be816e6e7d42d64736f6c63430008060033"

// DeployBuyFGE deploys a new Ethereum contract, binding an instance of BuyFGE to it.
func DeployBuyFGE(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BuyFGE, error) {
	parsed, err := abi.JSON(strings.NewReader(BuyFGEABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BuyFGEBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BuyFGE{BuyFGECaller: BuyFGECaller{contract: contract}, BuyFGETransactor: BuyFGETransactor{contract: contract}, BuyFGEFilterer: BuyFGEFilterer{contract: contract}}, nil
}

// BuyFGE is an auto generated Go binding around an Ethereum contract.
type BuyFGE struct {
	BuyFGECaller     // Read-only binding to the contract
	BuyFGETransactor // Write-only binding to the contract
	BuyFGEFilterer   // Log filterer for contract events
}

// BuyFGECaller is an auto generated read-only Go binding around an Ethereum contract.
type BuyFGECaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyFGETransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuyFGETransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyFGEFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuyFGEFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyFGESession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuyFGESession struct {
	Contract     *BuyFGE           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyFGECallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuyFGECallerSession struct {
	Contract *BuyFGECaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BuyFGETransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuyFGETransactorSession struct {
	Contract     *BuyFGETransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyFGERaw is an auto generated low-level Go binding around an Ethereum contract.
type BuyFGERaw struct {
	Contract *BuyFGE // Generic contract binding to access the raw methods on
}

// BuyFGECallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuyFGECallerRaw struct {
	Contract *BuyFGECaller // Generic read-only contract binding to access the raw methods on
}

// BuyFGETransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuyFGETransactorRaw struct {
	Contract *BuyFGETransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuyFGE creates a new instance of BuyFGE, bound to a specific deployed contract.
func NewBuyFGE(address common.Address, backend bind.ContractBackend) (*BuyFGE, error) {
	contract, err := bindBuyFGE(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BuyFGE{BuyFGECaller: BuyFGECaller{contract: contract}, BuyFGETransactor: BuyFGETransactor{contract: contract}, BuyFGEFilterer: BuyFGEFilterer{contract: contract}}, nil
}

// NewBuyFGECaller creates a new read-only instance of BuyFGE, bound to a specific deployed contract.
func NewBuyFGECaller(address common.Address, caller bind.ContractCaller) (*BuyFGECaller, error) {
	contract, err := bindBuyFGE(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuyFGECaller{contract: contract}, nil
}

// NewBuyFGETransactor creates a new write-only instance of BuyFGE, bound to a specific deployed contract.
func NewBuyFGETransactor(address common.Address, transactor bind.ContractTransactor) (*BuyFGETransactor, error) {
	contract, err := bindBuyFGE(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuyFGETransactor{contract: contract}, nil
}

// NewBuyFGEFilterer creates a new log filterer instance of BuyFGE, bound to a specific deployed contract.
func NewBuyFGEFilterer(address common.Address, filterer bind.ContractFilterer) (*BuyFGEFilterer, error) {
	contract, err := bindBuyFGE(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuyFGEFilterer{contract: contract}, nil
}

// bindBuyFGE binds a generic wrapper to an already deployed contract.
func bindBuyFGE(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BuyFGEABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuyFGE *BuyFGERaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuyFGE.Contract.BuyFGECaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuyFGE *BuyFGERaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyFGE.Contract.BuyFGETransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuyFGE *BuyFGERaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuyFGE.Contract.BuyFGETransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuyFGE *BuyFGECallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuyFGE.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuyFGE *BuyFGETransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyFGE.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuyFGE *BuyFGETransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuyFGE.Contract.contract.Transact(opts, method, params...)
}

// Transaction is a free data retrieval call binding the contract method 0xe54f992a.
//
// Solidity: function Transaction(bytes ) view returns(bytes payment_id, bool paid, uint256 actually_paid, uint256 pay_amount, uint256 price, address order_id, bytes payment_status, uint256 fge, bytes data)
func (_BuyFGE *BuyFGECaller) Transaction(opts *bind.CallOpts, arg0 []byte) (struct {
	PaymentId     []byte
	Paid          bool
	ActuallyPaid  *big.Int
	PayAmount     *big.Int
	Price         *big.Int
	OrderId       common.Address
	PaymentStatus []byte
	Fge           *big.Int
	Data          []byte
}, error) {
	var out []interface{}
	err := _BuyFGE.contract.Call(opts, &out, "Transaction", arg0)

	outstruct := new(struct {
		PaymentId     []byte
		Paid          bool
		ActuallyPaid  *big.Int
		PayAmount     *big.Int
		Price         *big.Int
		OrderId       common.Address
		PaymentStatus []byte
		Fge           *big.Int
		Data          []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PaymentId = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Paid = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.ActuallyPaid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PayAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OrderId = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.PaymentStatus = *abi.ConvertType(out[6], new([]byte)).(*[]byte)
	outstruct.Fge = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[8], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Transaction is a free data retrieval call binding the contract method 0xe54f992a.
//
// Solidity: function Transaction(bytes ) view returns(bytes payment_id, bool paid, uint256 actually_paid, uint256 pay_amount, uint256 price, address order_id, bytes payment_status, uint256 fge, bytes data)
func (_BuyFGE *BuyFGESession) Transaction(arg0 []byte) (struct {
	PaymentId     []byte
	Paid          bool
	ActuallyPaid  *big.Int
	PayAmount     *big.Int
	Price         *big.Int
	OrderId       common.Address
	PaymentStatus []byte
	Fge           *big.Int
	Data          []byte
}, error) {
	return _BuyFGE.Contract.Transaction(&_BuyFGE.CallOpts, arg0)
}

// Transaction is a free data retrieval call binding the contract method 0xe54f992a.
//
// Solidity: function Transaction(bytes ) view returns(bytes payment_id, bool paid, uint256 actually_paid, uint256 pay_amount, uint256 price, address order_id, bytes payment_status, uint256 fge, bytes data)
func (_BuyFGE *BuyFGECallerSession) Transaction(arg0 []byte) (struct {
	PaymentId     []byte
	Paid          bool
	ActuallyPaid  *big.Int
	PayAmount     *big.Int
	Price         *big.Int
	OrderId       common.Address
	PaymentStatus []byte
	Fge           *big.Int
	Data          []byte
}, error) {
	return _BuyFGE.Contract.Transaction(&_BuyFGE.CallOpts, arg0)
}

// GetTransaction is a free data retrieval call binding the contract method 0xfaa9e569.
//
// Solidity: function getTransaction(bytes payment_id) view returns((bytes,bool,uint256,uint256,uint256,address,bytes,uint256,bytes))
func (_BuyFGE *BuyFGECaller) GetTransaction(opts *bind.CallOpts, payment_id []byte) (BuyFGEPayment, error) {
	var out []interface{}
	err := _BuyFGE.contract.Call(opts, &out, "getTransaction", payment_id)

	if err != nil {
		return *new(BuyFGEPayment), err
	}

	out0 := *abi.ConvertType(out[0], new(BuyFGEPayment)).(*BuyFGEPayment)

	return out0, err

}

// GetTransaction is a free data retrieval call binding the contract method 0xfaa9e569.
//
// Solidity: function getTransaction(bytes payment_id) view returns((bytes,bool,uint256,uint256,uint256,address,bytes,uint256,bytes))
func (_BuyFGE *BuyFGESession) GetTransaction(payment_id []byte) (BuyFGEPayment, error) {
	return _BuyFGE.Contract.GetTransaction(&_BuyFGE.CallOpts, payment_id)
}

// GetTransaction is a free data retrieval call binding the contract method 0xfaa9e569.
//
// Solidity: function getTransaction(bytes payment_id) view returns((bytes,bool,uint256,uint256,uint256,address,bytes,uint256,bytes))
func (_BuyFGE *BuyFGECallerSession) GetTransaction(payment_id []byte) (BuyFGEPayment, error) {
	return _BuyFGE.Contract.GetTransaction(&_BuyFGE.CallOpts, payment_id)
}

// BuyFGE is a paid mutator transaction binding the contract method 0x78025a77.
//
// Solidity: function buyFGE(bytes payment_id, uint256 actually_paid, uint256 pay_amount, uint256 price, address order_id, bytes payment_status, uint256 fge) returns()
func (_BuyFGE *BuyFGETransactor) BuyFGE(opts *bind.TransactOpts, payment_id []byte, actually_paid *big.Int, pay_amount *big.Int, price *big.Int, order_id common.Address, payment_status []byte, fge *big.Int) (*types.Transaction, error) {
	return _BuyFGE.contract.Transact(opts, "buyFGE", payment_id, actually_paid, pay_amount, price, order_id, payment_status, fge)
}

// BuyFGE is a paid mutator transaction binding the contract method 0x78025a77.
//
// Solidity: function buyFGE(bytes payment_id, uint256 actually_paid, uint256 pay_amount, uint256 price, address order_id, bytes payment_status, uint256 fge) returns()
func (_BuyFGE *BuyFGESession) BuyFGE(payment_id []byte, actually_paid *big.Int, pay_amount *big.Int, price *big.Int, order_id common.Address, payment_status []byte, fge *big.Int) (*types.Transaction, error) {
	return _BuyFGE.Contract.BuyFGE(&_BuyFGE.TransactOpts, payment_id, actually_paid, pay_amount, price, order_id, payment_status, fge)
}

// BuyFGE is a paid mutator transaction binding the contract method 0x78025a77.
//
// Solidity: function buyFGE(bytes payment_id, uint256 actually_paid, uint256 pay_amount, uint256 price, address order_id, bytes payment_status, uint256 fge) returns()
func (_BuyFGE *BuyFGETransactorSession) BuyFGE(payment_id []byte, actually_paid *big.Int, pay_amount *big.Int, price *big.Int, order_id common.Address, payment_status []byte, fge *big.Int) (*types.Transaction, error) {
	return _BuyFGE.Contract.BuyFGE(&_BuyFGE.TransactOpts, payment_id, actually_paid, pay_amount, price, order_id, payment_status, fge)
}

// SendFGE is a paid mutator transaction binding the contract method 0x032d6475.
//
// Solidity: function sendFGE(address _to, (bytes,bool,uint256,uint256,uint256,address,bytes,uint256,bytes) transaction) payable returns()
func (_BuyFGE *BuyFGETransactor) SendFGE(opts *bind.TransactOpts, _to common.Address, transaction BuyFGEPayment) (*types.Transaction, error) {
	return _BuyFGE.contract.Transact(opts, "sendFGE", _to, transaction)
}

// SendFGE is a paid mutator transaction binding the contract method 0x032d6475.
//
// Solidity: function sendFGE(address _to, (bytes,bool,uint256,uint256,uint256,address,bytes,uint256,bytes) transaction) payable returns()
func (_BuyFGE *BuyFGESession) SendFGE(_to common.Address, transaction BuyFGEPayment) (*types.Transaction, error) {
	return _BuyFGE.Contract.SendFGE(&_BuyFGE.TransactOpts, _to, transaction)
}

// SendFGE is a paid mutator transaction binding the contract method 0x032d6475.
//
// Solidity: function sendFGE(address _to, (bytes,bool,uint256,uint256,uint256,address,bytes,uint256,bytes) transaction) payable returns()
func (_BuyFGE *BuyFGETransactorSession) SendFGE(_to common.Address, transaction BuyFGEPayment) (*types.Transaction, error) {
	return _BuyFGE.Contract.SendFGE(&_BuyFGE.TransactOpts, _to, transaction)
}

// BuyFGESendFGEIterator is returned from FilterSendFGE and is used to iterate over the raw logs and unpacked data for SendFGE events raised by the BuyFGE contract.
type BuyFGESendFGEIterator struct {
	Event *BuyFGESendFGE // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BuyFGESendFGEIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyFGESendFGE)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BuyFGESendFGE)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BuyFGESendFGEIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyFGESendFGEIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyFGESendFGE represents a SendFGE event raised by the BuyFGE contract.
type BuyFGESendFGE struct {
	To          common.Address
	Transaction BuyFGEPayment
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSendFGE is a free log retrieval operation binding the contract event 0xe7e9ebbbcd30fbde2b2533292813701723c078316675bec53e74067fbfd5e8f7.
//
// Solidity: event SendFGE(address _to, (bytes,bool,uint256,uint256,uint256,address,bytes,uint256,bytes) transaction)
func (_BuyFGE *BuyFGEFilterer) FilterSendFGE(opts *bind.FilterOpts) (*BuyFGESendFGEIterator, error) {

	logs, sub, err := _BuyFGE.contract.FilterLogs(opts, "SendFGE")
	if err != nil {
		return nil, err
	}
	return &BuyFGESendFGEIterator{contract: _BuyFGE.contract, event: "SendFGE", logs: logs, sub: sub}, nil
}

// WatchSendFGE is a free log subscription operation binding the contract event 0xe7e9ebbbcd30fbde2b2533292813701723c078316675bec53e74067fbfd5e8f7.
//
// Solidity: event SendFGE(address _to, (bytes,bool,uint256,uint256,uint256,address,bytes,uint256,bytes) transaction)
func (_BuyFGE *BuyFGEFilterer) WatchSendFGE(opts *bind.WatchOpts, sink chan<- *BuyFGESendFGE) (event.Subscription, error) {

	logs, sub, err := _BuyFGE.contract.WatchLogs(opts, "SendFGE")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyFGESendFGE)
				if err := _BuyFGE.contract.UnpackLog(event, "SendFGE", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSendFGE is a log parse operation binding the contract event 0xe7e9ebbbcd30fbde2b2533292813701723c078316675bec53e74067fbfd5e8f7.
//
// Solidity: event SendFGE(address _to, (bytes,bool,uint256,uint256,uint256,address,bytes,uint256,bytes) transaction)
func (_BuyFGE *BuyFGEFilterer) ParseSendFGE(log types.Log) (*BuyFGESendFGE, error) {
	event := new(BuyFGESendFGE)
	if err := _BuyFGE.contract.UnpackLog(event, "SendFGE", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
