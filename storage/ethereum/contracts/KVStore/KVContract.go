// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"data\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
const StoreBin = `608060405234801561001057600080fd5b5061072a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063a18c751e14610046578063d3d29df114610076578063d6d7d525146100a6575b600080fd5b610060600480360381019061005b9190610478565b6100d6565b60405161006d919061050b565b60405180910390f35b610090600480360381019061008b9190610526565b610115565b60405161009d91906105f7565b60405180910390f35b6100c060048036038101906100bb9190610526565b6101cb565b6040516100cd91906105f7565b60405180910390f35b6000816000846040516100e99190610655565b9081526020016040518091039020908051906020019061010a92919061027b565b506001905092915050565b600081805160208101820180518482526020830160208501208183528095505050505050600091509050805461014a9061069b565b80601f01602080910402602001604051908101604052809291908181526020018280546101769061069b565b80156101c35780601f10610198576101008083540402835291602001916101c3565b820191906000526020600020905b8154815290600101906020018083116101a657829003601f168201915b505050505081565b60606000826040516101dd9190610655565b908152602001604051809103902080546101f69061069b565b80601f01602080910402602001604051908101604052809291908181526020018280546102229061069b565b801561026f5780601f106102445761010080835404028352916020019161026f565b820191906000526020600020905b81548152906001019060200180831161025257829003601f168201915b50505050509050919050565b8280546102879061069b565b90600052602060002090601f0160209004810192826102a957600085556102f0565b82601f106102c257805160ff19168380011785556102f0565b828001600101855582156102f0579182015b828111156102ef5782518255916020019190600101906102d4565b5b5090506102fd9190610301565b5090565b5b8082111561031a576000816000905550600101610302565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103858261033c565b810181811067ffffffffffffffff821117156103a4576103a361034d565b5b80604052505050565b60006103b761031e565b90506103c3828261037c565b919050565b600067ffffffffffffffff8211156103e3576103e261034d565b5b6103ec8261033c565b9050602081019050919050565b82818337600083830152505050565b600061041b610416846103c8565b6103ad565b90508281526020810184848401111561043757610436610337565b5b6104428482856103f9565b509392505050565b600082601f83011261045f5761045e610332565b5b813561046f848260208601610408565b91505092915050565b6000806040838503121561048f5761048e610328565b5b600083013567ffffffffffffffff8111156104ad576104ac61032d565b5b6104b98582860161044a565b925050602083013567ffffffffffffffff8111156104da576104d961032d565b5b6104e68582860161044a565b9150509250929050565b60008115159050919050565b610505816104f0565b82525050565b600060208201905061052060008301846104fc565b92915050565b60006020828403121561053c5761053b610328565b5b600082013567ffffffffffffffff81111561055a5761055961032d565b5b6105668482850161044a565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156105a957808201518184015260208101905061058e565b838111156105b8576000848401525b50505050565b60006105c98261056f565b6105d3818561057a565b93506105e381856020860161058b565b6105ec8161033c565b840191505092915050565b6000602082019050818103600083015261061181846105be565b905092915050565b600081905092915050565b600061062f8261056f565b6106398185610619565b935061064981856020860161058b565b80840191505092915050565b60006106618284610624565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806106b357607f821691505b602082108114156106c7576106c661066c565b5b5091905056fea2646970667358221220a892def3a09d2be3d88a94af6e7da1dbe48144e3756338e2a5f008c979c25b4d64736f6c637829302e382e31302d646576656c6f702e323032312e31302e32342b636f6d6d69742e3337373439353564005a`

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Data is a paid mutator transaction binding the contract method 0xd3d29df1.
//
// Solidity: function data(bytes ) returns(bytes)
func (_Store *StoreTransactor) Data(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "data", arg0)
}

// Data is a paid mutator transaction binding the contract method 0xd3d29df1.
//
// Solidity: function data(bytes ) returns(bytes)
func (_Store *StoreSession) Data(arg0 []byte) (*types.Transaction, error) {
	return _Store.Contract.Data(&_Store.TransactOpts, arg0)
}

// Data is a paid mutator transaction binding the contract method 0xd3d29df1.
//
// Solidity: function data(bytes ) returns(bytes)
func (_Store *StoreTransactorSession) Data(arg0 []byte) (*types.Transaction, error) {
	return _Store.Contract.Data(&_Store.TransactOpts, arg0)
}

// Get is a paid mutator transaction binding the contract method 0xd6d7d525.
//
// Solidity: function get(bytes key) returns(bytes value)
func (_Store *StoreTransactor) Get(opts *bind.TransactOpts, key []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "get", key)
}

// Get is a paid mutator transaction binding the contract method 0xd6d7d525.
//
// Solidity: function get(bytes key) returns(bytes value)
func (_Store *StoreSession) Get(key []byte) (*types.Transaction, error) {
	return _Store.Contract.Get(&_Store.TransactOpts, key)
}

// Get is a paid mutator transaction binding the contract method 0xd6d7d525.
//
// Solidity: function get(bytes key) returns(bytes value)
func (_Store *StoreTransactorSession) Get(key []byte) (*types.Transaction, error) {
	return _Store.Contract.Get(&_Store.TransactOpts, key)
}

// Set is a paid mutator transaction binding the contract method 0xa18c751e.
//
// Solidity: function set(bytes key, bytes val) returns(bool success)
func (_Store *StoreTransactor) Set(opts *bind.TransactOpts, key []byte, val []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "set", key, val)
}

// Set is a paid mutator transaction binding the contract method 0xa18c751e.
//
// Solidity: function set(bytes key, bytes val) returns(bool success)
func (_Store *StoreSession) Set(key []byte, val []byte) (*types.Transaction, error) {
	return _Store.Contract.Set(&_Store.TransactOpts, key, val)
}

// Set is a paid mutator transaction binding the contract method 0xa18c751e.
//
// Solidity: function set(bytes key, bytes val) returns(bool success)
func (_Store *StoreTransactorSession) Set(key []byte, val []byte) (*types.Transaction, error) {
	return _Store.Contract.Set(&_Store.TransactOpts, key, val)
}
