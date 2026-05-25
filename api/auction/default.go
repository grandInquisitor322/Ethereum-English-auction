// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auction

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AuctionMetaData contains all meta data concerning the Auction contract.
var AuctionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_nft\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nftId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allBids\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bid\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"end\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ended\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"highestBid\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"highestBidder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nft\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC721\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nftId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"start\",\"inputs\":[{\"name\":\"_openingBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"started\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Bid\",\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"End\",\"inputs\":[{\"name\":\"highestBidder\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Start\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161143538038061143583398181016040528101906100319190610139565b3373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508060c081815250505050610177565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100d5826100ac565b9050919050565b6100e5816100cb565b81146100ef575f5ffd5b50565b5f81519050610100816100dc565b92915050565b5f819050919050565b61011881610106565b8114610122575f5ffd5b50565b5f815190506101338161010f565b92915050565b5f5f6040838503121561014f5761014e6100a8565b5b5f61015c858286016100f2565b925050602061016d85828601610125565b9150509250929050565b60805160a05160c05161125c6101d95f395f8181610778015281816108760152610a8801525f81816105dd0152818161071a0152610a2901525f81816106010152818161062501528181610756015281816108a00152610af4015261125c5ff3fe6080604052600436106100c1575f3560e01c80638da5cb5b1161007e578063bd20160711610058578063bd20160714610209578063c6bc518214610245578063d57bde791461026f578063efbe1c1c14610299576100c1565b80638da5cb5b1461018d5780638fb4b573146101b757806391f90157146101df576100c1565b806312fa6feb146100c55780631998aeef146100ef5780631f2698ab146100f95780633197cbb6146101235780633ccfd60b1461014d57806347ccca0214610163575b5f5ffd5b3480156100d0575f5ffd5b506100d96102af565b6040516100e69190610bd0565b60405180910390f35b6100f76102c1565b005b348015610104575f5ffd5b5061010d6104a3565b60405161011a9190610bd0565b60405180910390f35b34801561012e575f5ffd5b506101376104b4565b6040516101449190610c01565b60405180910390f35b348015610158575f5ffd5b506101616104ba565b005b34801561016e575f5ffd5b506101776105db565b6040516101849190610c94565b60405180910390f35b348015610198575f5ffd5b506101a16105ff565b6040516101ae9190610ccd565b60405180910390f35b3480156101c2575f5ffd5b506101dd60048036038101906101d89190610d14565b610623565b005b3480156101ea575f5ffd5b506101f361083a565b6040516102009190610d72565b60405180910390f35b348015610214575f5ffd5b5061022f600480360381019061022a9190610db5565b61085f565b60405161023c9190610c01565b60405180910390f35b348015610250575f5ffd5b50610259610874565b6040516102669190610c01565b60405180910390f35b34801561027a575f5ffd5b50610283610898565b6040516102909190610c01565b60405180910390f35b3480156102a4575f5ffd5b506102ad61089e565b005b5f60019054906101000a900460ff1681565b5f5f9054906101000a900460ff1661030e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030590610e3a565b60405180910390fd5b6002543411610352576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034990610ec8565b60405180910390fd5b6001544210610396576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038d90610f30565b60405180910390fd5b60025460045f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546104059190610f7b565b92505081905550346002819055503360035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167fe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2346040516104999190610c01565b60405180910390a2565b5f5f9054906101000a900460ff1681565b60015481565b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f81111561058a573373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015610588573d5f5f3e3d5ffd5b505b3373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364826040516105d09190610c01565b60405180910390a250565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a890610ff8565b60405180910390fd5b5f5f9054906101000a900460ff16156106ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f690611060565b60405180910390fd5b8160028190555080426107129190610f7b565b6001819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd7f0000000000000000000000000000000000000000000000000000000000000000307f00000000000000000000000000000000000000000000000000000000000000006040518463ffffffff1660e01b81526004016107b59392919061109e565b5f604051808303815f87803b1580156107cc575f5ffd5b505af11580156107de573d5f5f3e3d5ffd5b5050505060015f5f6101000a81548160ff0219169083151502179055507f5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee4260015460405161082e9291906110d3565b60405180910390a15050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6004602052805f5260405f205f915090505481565b7f000000000000000000000000000000000000000000000000000000000000000081565b60025481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461092c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092390610ff8565b60405180910390fd5b5f5f9054906101000a900460ff16610979576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097090610e3a565b60405180910390fd5b6001544210156109be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b590611144565b60405180910390fd5b5f60019054906101000a900460ff1615610a0d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a04906111ac565b60405180910390fd5b60015f60016101000a81548160ff0219169083151502179055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff167f00000000000000000000000000000000000000000000000000000000000000006040518463ffffffff1660e01b8152600401610ac5939291906111ca565b5f604051808303815f87803b158015610adc575f5ffd5b505af1158015610aee573d5f5f3e3d5ffd5b505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc60025490811502906040515f60405180830381858888f19350505050158015610b57573d5f5f3e3d5ffd5b507f7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc560035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600254604051610bac9291906111ff565b60405180910390a1565b5f8115159050919050565b610bca81610bb6565b82525050565b5f602082019050610be35f830184610bc1565b92915050565b5f819050919050565b610bfb81610be9565b82525050565b5f602082019050610c145f830184610bf2565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610c5c610c57610c5284610c1a565b610c39565b610c1a565b9050919050565b5f610c6d82610c42565b9050919050565b5f610c7e82610c63565b9050919050565b610c8e81610c74565b82525050565b5f602082019050610ca75f830184610c85565b92915050565b5f610cb782610c1a565b9050919050565b610cc781610cad565b82525050565b5f602082019050610ce05f830184610cbe565b92915050565b5f5ffd5b610cf381610be9565b8114610cfd575f5ffd5b50565b5f81359050610d0e81610cea565b92915050565b5f5f60408385031215610d2a57610d29610ce6565b5b5f610d3785828601610d00565b9250506020610d4885828601610d00565b9150509250929050565b5f610d5c82610c1a565b9050919050565b610d6c81610d52565b82525050565b5f602082019050610d855f830184610d63565b92915050565b610d9481610d52565b8114610d9e575f5ffd5b50565b5f81359050610daf81610d8b565b92915050565b5f60208284031215610dca57610dc9610ce6565b5b5f610dd784828501610da1565b91505092915050565b5f82825260208201905092915050565b7f41756374696f6e20686173206e6f7420737461727465640000000000000000005f82015250565b5f610e24601783610de0565b9150610e2f82610df0565b602082019050919050565b5f6020820190508181035f830152610e5181610e18565b9050919050565b7f426964207072696365206973206c6f776572207468616e2074686520637572725f8201527f656e742068696768657374206269640000000000000000000000000000000000602082015250565b5f610eb2602f83610de0565b9150610ebd82610e58565b604082019050919050565b5f6020820190508181035f830152610edf81610ea6565b9050919050565b7f41756374696f6e2068617320656e6465640000000000000000000000000000005f82015250565b5f610f1a601183610de0565b9150610f2582610ee6565b602082019050919050565b5f6020820190508181035f830152610f4781610f0e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610f8582610be9565b9150610f9083610be9565b9250828201905080821115610fa857610fa7610f4e565b5b92915050565b7f4f6e6c79206f776e65722063616e2063616c6c207468652066756e6374696f6e5f82015250565b5f610fe2602083610de0565b9150610fed82610fae565b602082019050919050565b5f6020820190508181035f83015261100f81610fd6565b9050919050565b7f41756374696f6e2068617320616c7265616479207374617274656400000000005f82015250565b5f61104a601b83610de0565b915061105582611016565b602082019050919050565b5f6020820190508181035f8301526110778161103e565b9050919050565b5f61108882610c63565b9050919050565b6110988161107e565b82525050565b5f6060820190506110b15f83018661108f565b6110be6020830185610d63565b6110cb6040830184610bf2565b949350505050565b5f6040820190506110e65f830185610bf2565b6110f36020830184610bf2565b9392505050565b7f41756374696f6e20686173206e6f7420656e64656400000000000000000000005f82015250565b5f61112e601583610de0565b9150611139826110fa565b602082019050919050565b5f6020820190508181035f83015261115b81611122565b9050919050565b7f41756374696f6e2068617320616c726561647920656e646564000000000000005f82015250565b5f611196601983610de0565b91506111a182611162565b602082019050919050565b5f6020820190508181035f8301526111c38161118a565b9050919050565b5f6060820190506111dd5f830186610d63565b6111ea6020830185610d63565b6111f76040830184610bf2565b949350505050565b5f6040820190506112125f830185610d63565b61121f6020830184610bf2565b939250505056fea26469706673582212206ecbaa45cbc0c3bd68b05a2a74c315b03d30f3f10665b7c947d3b0bcd5a9c16664736f6c63430008230033",
}

// AuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionMetaData.ABI instead.
var AuctionABI = AuctionMetaData.ABI

// AuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuctionMetaData.Bin instead.
var AuctionBin = AuctionMetaData.Bin

// DeployAuction deploys a new Ethereum contract, binding an instance of Auction to it.
func DeployAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _nft common.Address, _nftId *big.Int) (common.Address, *types.Transaction, *Auction, error) {
	parsed, err := AuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuctionBin), backend, _nft, _nftId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// Auction is an auto generated Go binding around an Ethereum contract.
type Auction struct {
	AuctionCaller     // Read-only binding to the contract
	AuctionTransactor // Write-only binding to the contract
	AuctionFilterer   // Log filterer for contract events
}

// AuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionSession struct {
	Contract     *Auction          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionCallerSession struct {
	Contract *AuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionTransactorSession struct {
	Contract     *AuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionRaw struct {
	Contract *Auction // Generic contract binding to access the raw methods on
}

// AuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionCallerRaw struct {
	Contract *AuctionCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionTransactorRaw struct {
	Contract *AuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuction creates a new instance of Auction, bound to a specific deployed contract.
func NewAuction(address common.Address, backend bind.ContractBackend) (*Auction, error) {
	contract, err := bindAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// NewAuctionCaller creates a new read-only instance of Auction, bound to a specific deployed contract.
func NewAuctionCaller(address common.Address, caller bind.ContractCaller) (*AuctionCaller, error) {
	contract, err := bindAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionCaller{contract: contract}, nil
}

// NewAuctionTransactor creates a new write-only instance of Auction, bound to a specific deployed contract.
func NewAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionTransactor, error) {
	contract, err := bindAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionTransactor{contract: contract}, nil
}

// NewAuctionFilterer creates a new log filterer instance of Auction, bound to a specific deployed contract.
func NewAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionFilterer, error) {
	contract, err := bindAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionFilterer{contract: contract}, nil
}

// bindAuction binds a generic wrapper to an already deployed contract.
func bindAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.AuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transact(opts, method, params...)
}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Auction *AuctionCaller) AllBids(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "allBids", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Auction *AuctionSession) AllBids(arg0 common.Address) (*big.Int, error) {
	return _Auction.Contract.AllBids(&_Auction.CallOpts, arg0)
}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Auction *AuctionCallerSession) AllBids(arg0 common.Address) (*big.Int, error) {
	return _Auction.Contract.AllBids(&_Auction.CallOpts, arg0)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Auction *AuctionCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Auction *AuctionSession) EndTime() (*big.Int, error) {
	return _Auction.Contract.EndTime(&_Auction.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Auction *AuctionCallerSession) EndTime() (*big.Int, error) {
	return _Auction.Contract.EndTime(&_Auction.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Auction *AuctionCaller) Ended(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "ended")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Auction *AuctionSession) Ended() (bool, error) {
	return _Auction.Contract.Ended(&_Auction.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Auction *AuctionCallerSession) Ended() (bool, error) {
	return _Auction.Contract.Ended(&_Auction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Auction *AuctionCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Auction *AuctionSession) HighestBid() (*big.Int, error) {
	return _Auction.Contract.HighestBid(&_Auction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Auction *AuctionCallerSession) HighestBid() (*big.Int, error) {
	return _Auction.Contract.HighestBid(&_Auction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Auction *AuctionCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "highestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Auction *AuctionSession) HighestBidder() (common.Address, error) {
	return _Auction.Contract.HighestBidder(&_Auction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Auction *AuctionCallerSession) HighestBidder() (common.Address, error) {
	return _Auction.Contract.HighestBidder(&_Auction.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Auction *AuctionCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Auction *AuctionSession) Nft() (common.Address, error) {
	return _Auction.Contract.Nft(&_Auction.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Auction *AuctionCallerSession) Nft() (common.Address, error) {
	return _Auction.Contract.Nft(&_Auction.CallOpts)
}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_Auction *AuctionCaller) NftId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "nftId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_Auction *AuctionSession) NftId() (*big.Int, error) {
	return _Auction.Contract.NftId(&_Auction.CallOpts)
}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_Auction *AuctionCallerSession) NftId() (*big.Int, error) {
	return _Auction.Contract.NftId(&_Auction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionSession) Owner() (common.Address, error) {
	return _Auction.Contract.Owner(&_Auction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionCallerSession) Owner() (common.Address, error) {
	return _Auction.Contract.Owner(&_Auction.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Auction *AuctionCaller) Started(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "started")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Auction *AuctionSession) Started() (bool, error) {
	return _Auction.Contract.Started(&_Auction.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Auction *AuctionCallerSession) Started() (bool, error) {
	return _Auction.Contract.Started(&_Auction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction *AuctionTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction *AuctionSession) Bid() (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction *AuctionTransactorSession) Bid() (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts)
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Auction *AuctionTransactor) End(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "end")
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Auction *AuctionSession) End() (*types.Transaction, error) {
	return _Auction.Contract.End(&_Auction.TransactOpts)
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Auction *AuctionTransactorSession) End() (*types.Transaction, error) {
	return _Auction.Contract.End(&_Auction.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Auction *AuctionTransactor) Start(opts *bind.TransactOpts, _openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "start", _openingBid, _duration)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Auction *AuctionSession) Start(_openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Start(&_Auction.TransactOpts, _openingBid, _duration)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Auction *AuctionTransactorSession) Start(_openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Start(&_Auction.TransactOpts, _openingBid, _duration)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionSession) Withdraw() (*types.Transaction, error) {
	return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}

// AuctionBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the Auction contract.
type AuctionBidIterator struct {
	Event *AuctionBid // Event containing the contract specifics and raw log

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
func (it *AuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBid)
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
		it.Event = new(AuctionBid)
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
func (it *AuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBid represents a Bid event raised by the Auction contract.
type AuctionBid struct {
	Bidder common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) FilterBid(opts *bind.FilterOpts, bidder []common.Address) (*AuctionBidIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionBidIterator{contract: _Auction.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *AuctionBid, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBid)
				if err := _Auction.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) ParseBid(log types.Log) (*AuctionBid, error) {
	event := new(AuctionBid)
	if err := _Auction.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionEndIterator is returned from FilterEnd and is used to iterate over the raw logs and unpacked data for End events raised by the Auction contract.
type AuctionEndIterator struct {
	Event *AuctionEnd // Event containing the contract specifics and raw log

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
func (it *AuctionEndIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionEnd)
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
		it.Event = new(AuctionEnd)
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
func (it *AuctionEndIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionEndIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionEnd represents a End event raised by the Auction contract.
type AuctionEnd struct {
	HighestBidder common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEnd is a free log retrieval operation binding the contract event 0x7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5.
//
// Solidity: event End(address highestBidder, uint256 value)
func (_Auction *AuctionFilterer) FilterEnd(opts *bind.FilterOpts) (*AuctionEndIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "End")
	if err != nil {
		return nil, err
	}
	return &AuctionEndIterator{contract: _Auction.contract, event: "End", logs: logs, sub: sub}, nil
}

// WatchEnd is a free log subscription operation binding the contract event 0x7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5.
//
// Solidity: event End(address highestBidder, uint256 value)
func (_Auction *AuctionFilterer) WatchEnd(opts *bind.WatchOpts, sink chan<- *AuctionEnd) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "End")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionEnd)
				if err := _Auction.contract.UnpackLog(event, "End", log); err != nil {
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

// ParseEnd is a log parse operation binding the contract event 0x7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5.
//
// Solidity: event End(address highestBidder, uint256 value)
func (_Auction *AuctionFilterer) ParseEnd(log types.Log) (*AuctionEnd, error) {
	event := new(AuctionEnd)
	if err := _Auction.contract.UnpackLog(event, "End", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionStartIterator is returned from FilterStart and is used to iterate over the raw logs and unpacked data for Start events raised by the Auction contract.
type AuctionStartIterator struct {
	Event *AuctionStart // Event containing the contract specifics and raw log

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
func (it *AuctionStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionStart)
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
		it.Event = new(AuctionStart)
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
func (it *AuctionStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionStart represents a Start event raised by the Auction contract.
type AuctionStart struct {
	StartTime *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStart is a free log retrieval operation binding the contract event 0x5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee.
//
// Solidity: event Start(uint256 startTime, uint256 endTime)
func (_Auction *AuctionFilterer) FilterStart(opts *bind.FilterOpts) (*AuctionStartIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return &AuctionStartIterator{contract: _Auction.contract, event: "Start", logs: logs, sub: sub}, nil
}

// WatchStart is a free log subscription operation binding the contract event 0x5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee.
//
// Solidity: event Start(uint256 startTime, uint256 endTime)
func (_Auction *AuctionFilterer) WatchStart(opts *bind.WatchOpts, sink chan<- *AuctionStart) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionStart)
				if err := _Auction.contract.UnpackLog(event, "Start", log); err != nil {
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

// ParseStart is a log parse operation binding the contract event 0x5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee.
//
// Solidity: event Start(uint256 startTime, uint256 endTime)
func (_Auction *AuctionFilterer) ParseStart(log types.Log) (*AuctionStart, error) {
	event := new(AuctionStart)
	if err := _Auction.contract.UnpackLog(event, "Start", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Auction contract.
type AuctionWithdrawIterator struct {
	Event *AuctionWithdraw // Event containing the contract specifics and raw log

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
func (it *AuctionWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionWithdraw)
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
		it.Event = new(AuctionWithdraw)
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
func (it *AuctionWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionWithdraw represents a Withdraw event raised by the Auction contract.
type AuctionWithdraw struct {
	Bidder common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) FilterWithdraw(opts *bind.FilterOpts, bidder []common.Address) (*AuctionWithdrawIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Withdraw", bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionWithdrawIterator{contract: _Auction.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AuctionWithdraw, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Withdraw", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionWithdraw)
				if err := _Auction.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) ParseWithdraw(log types.Log) (*AuctionWithdraw, error) {
	event := new(AuctionWithdraw)
	if err := _Auction.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
