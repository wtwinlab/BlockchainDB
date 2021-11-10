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
const StoreABI = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
const StoreBin = `608060405234801561001057600080fd5b50610612806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063a18c751e1461003b578063d6d7d5251461006b575b600080fd5b61005560048036038101906100509190610387565b61009b565b604051610062919061041a565b60405180910390f35b61008560048036038101906100809190610435565b6100da565b6040516100929190610506565b60405180910390f35b6000816000846040516100ae9190610564565b908152602001604051809103902090805190602001906100cf92919061018a565b506001905092915050565b60606000826040516100ec9190610564565b90815260200160405180910390208054610105906105aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610131906105aa565b801561017e5780601f106101535761010080835404028352916020019161017e565b820191906000526020600020905b81548152906001019060200180831161016157829003601f168201915b50505050509050919050565b828054610196906105aa565b90600052602060002090601f0160209004810192826101b857600085556101ff565b82601f106101d157805160ff19168380011785556101ff565b828001600101855582156101ff579182015b828111156101fe5782518255916020019190600101906101e3565b5b50905061020c9190610210565b5090565b5b80821115610229576000816000905550600101610211565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6102948261024b565b810181811067ffffffffffffffff821117156102b3576102b261025c565b5b80604052505050565b60006102c661022d565b90506102d2828261028b565b919050565b600067ffffffffffffffff8211156102f2576102f161025c565b5b6102fb8261024b565b9050602081019050919050565b82818337600083830152505050565b600061032a610325846102d7565b6102bc565b90508281526020810184848401111561034657610345610246565b5b610351848285610308565b509392505050565b600082601f83011261036e5761036d610241565b5b813561037e848260208601610317565b91505092915050565b6000806040838503121561039e5761039d610237565b5b600083013567ffffffffffffffff8111156103bc576103bb61023c565b5b6103c885828601610359565b925050602083013567ffffffffffffffff8111156103e9576103e861023c565b5b6103f585828601610359565b9150509250929050565b60008115159050919050565b610414816103ff565b82525050565b600060208201905061042f600083018461040b565b92915050565b60006020828403121561044b5761044a610237565b5b600082013567ffffffffffffffff8111156104695761046861023c565b5b61047584828501610359565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156104b857808201518184015260208101905061049d565b838111156104c7576000848401525b50505050565b60006104d88261047e565b6104e28185610489565b93506104f281856020860161049a565b6104fb8161024b565b840191505092915050565b6000602082019050818103600083015261052081846104cd565b905092915050565b600081905092915050565b600061053e8261047e565b6105488185610528565b935061055881856020860161049a565b80840191505092915050565b60006105708284610533565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806105c257607f821691505b602082108114156105d6576105d561057b565b5b5091905056fea264697066735822122037e006b40b3f7c319b73b62edd654ca486b0334fa8a62e44241cb09b5c4d6e8d64736f6c63430008090033`

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
