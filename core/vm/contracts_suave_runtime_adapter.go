// Code generated by suave/gen. DO NOT EDIT.
// Hash: 3560bf1169d2bb413a9a0303c4043027fe6111a2f3cbfa25bf264612b2694ee4
package vm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/suave/artifacts"
	"github.com/mitchellh/mapstructure"
)

var (
	errFailedToUnpackInput = fmt.Errorf("failed to decode input")
	errFailedToDecodeField = fmt.Errorf("failed to decode field")
	errFailedToPackOutput  = fmt.Errorf("failed to encode output")
)

type SuaveRuntime interface {
	buildEthBlock(blockArgs types.BuildBlockArgs, bidId types.BidId, namespace string) ([]byte, []byte, error)
	confidentialInputs() ([]byte, error)
	confidentialStoreRetrieve(bidId types.BidId, key string) ([]byte, error)
	confidentialStoreStore(bidId types.BidId, key string, data1 []byte) error
	extractHint(bundleData []byte) ([]byte, error)
	fetchBids(cond uint64, namespace string) ([]types.Bid, error)
	newBid(decryptionCondition uint64, allowedPeekers []common.Address, bidType string) (types.Bid, error)
	simulateBundle(bundleData []byte) (uint64, error)
	submitEthBlockBidToRelay(relayUrl string, builderBid []byte) ([]byte, error)
}

type SuaveRuntimeAdapter struct {
	impl SuaveRuntime
}

func (b *SuaveRuntimeAdapter) buildEthBlock(input []byte) (res []byte, err error) {
	var (
		unpacked []interface{}
		result   []byte
	)

	_ = unpacked
	_ = result

	unpacked, err = artifacts.SuaveAbi.Methods["buildEthBlock"].Inputs.Unpack(input)
	if err != nil {
		err = errFailedToUnpackInput
		return
	}

	var (
		blockArgs types.BuildBlockArgs
		bidId     types.BidId
		namespace string
	)

	if err = mapstructure.Decode(unpacked[0], &blockArgs); err != nil {
		err = errFailedToDecodeField
		return
	}

	if err = mapstructure.Decode(unpacked[1], &bidId); err != nil {
		err = errFailedToDecodeField
		return
	}

	namespace = unpacked[2].(string)

	var (
		output1 []byte
		output2 []byte
	)

	if output1, output2, err = b.impl.buildEthBlock(blockArgs, bidId, namespace); err != nil {
		return
	}

	result, err = artifacts.SuaveAbi.Methods["buildEthBlock"].Outputs.Pack(output1, output2)
	if err != nil {
		err = errFailedToPackOutput
		return
	}
	return result, nil

}

func (b *SuaveRuntimeAdapter) confidentialInputs(input []byte) (res []byte, err error) {
	var (
		unpacked []interface{}
		result   []byte
	)

	_ = unpacked
	_ = result

	unpacked, err = artifacts.SuaveAbi.Methods["confidentialInputs"].Inputs.Unpack(input)
	if err != nil {
		err = errFailedToUnpackInput
		return
	}

	var ()

	var (
		output1 []byte
	)

	if output1, err = b.impl.confidentialInputs(); err != nil {
		return
	}

	result = output1
	return result, nil

}

func (b *SuaveRuntimeAdapter) confidentialStoreRetrieve(input []byte) (res []byte, err error) {
	var (
		unpacked []interface{}
		result   []byte
	)

	_ = unpacked
	_ = result

	unpacked, err = artifacts.SuaveAbi.Methods["confidentialStoreRetrieve"].Inputs.Unpack(input)
	if err != nil {
		err = errFailedToUnpackInput
		return
	}

	var (
		bidId types.BidId
		key   string
	)

	if err = mapstructure.Decode(unpacked[0], &bidId); err != nil {
		err = errFailedToDecodeField
		return
	}

	key = unpacked[1].(string)

	var (
		output1 []byte
	)

	if output1, err = b.impl.confidentialStoreRetrieve(bidId, key); err != nil {
		return
	}

	result = output1
	return result, nil

}

func (b *SuaveRuntimeAdapter) confidentialStoreStore(input []byte) (res []byte, err error) {
	var (
		unpacked []interface{}
		result   []byte
	)

	_ = unpacked
	_ = result

	unpacked, err = artifacts.SuaveAbi.Methods["confidentialStoreStore"].Inputs.Unpack(input)
	if err != nil {
		err = errFailedToUnpackInput
		return
	}

	var (
		bidId types.BidId
		key   string
		data1 []byte
	)

	if err = mapstructure.Decode(unpacked[0], &bidId); err != nil {
		err = errFailedToDecodeField
		return
	}

	key = unpacked[1].(string)
	data1 = unpacked[2].([]byte)

	var ()

	if err = b.impl.confidentialStoreStore(bidId, key, data1); err != nil {
		return
	}

	return nil, nil

}

func (b *SuaveRuntimeAdapter) extractHint(input []byte) (res []byte, err error) {
	var (
		unpacked []interface{}
		result   []byte
	)

	_ = unpacked
	_ = result

	unpacked, err = artifacts.SuaveAbi.Methods["extractHint"].Inputs.Unpack(input)
	if err != nil {
		err = errFailedToUnpackInput
		return
	}

	var (
		bundleData []byte
	)

	bundleData = unpacked[0].([]byte)

	var (
		output1 []byte
	)

	if output1, err = b.impl.extractHint(bundleData); err != nil {
		return
	}

	result = output1
	return result, nil

}

func (b *SuaveRuntimeAdapter) fetchBids(input []byte) (res []byte, err error) {
	var (
		unpacked []interface{}
		result   []byte
	)

	_ = unpacked
	_ = result

	unpacked, err = artifacts.SuaveAbi.Methods["fetchBids"].Inputs.Unpack(input)
	if err != nil {
		err = errFailedToUnpackInput
		return
	}

	var (
		cond      uint64
		namespace string
	)

	cond = unpacked[0].(uint64)
	namespace = unpacked[1].(string)

	var (
		bid []types.Bid
	)

	if bid, err = b.impl.fetchBids(cond, namespace); err != nil {
		return
	}

	result, err = artifacts.SuaveAbi.Methods["fetchBids"].Outputs.Pack(bid)
	if err != nil {
		err = errFailedToPackOutput
		return
	}
	return result, nil

}

func (b *SuaveRuntimeAdapter) newBid(input []byte) (res []byte, err error) {
	var (
		unpacked []interface{}
		result   []byte
	)

	_ = unpacked
	_ = result

	unpacked, err = artifacts.SuaveAbi.Methods["newBid"].Inputs.Unpack(input)
	if err != nil {
		err = errFailedToUnpackInput
		return
	}

	var (
		decryptionCondition uint64
		allowedPeekers      []common.Address
		bidType             string
	)

	decryptionCondition = unpacked[0].(uint64)
	allowedPeekers = unpacked[1].([]common.Address)
	bidType = unpacked[2].(string)

	var (
		bid types.Bid
	)

	if bid, err = b.impl.newBid(decryptionCondition, allowedPeekers, bidType); err != nil {
		return
	}

	result, err = artifacts.SuaveAbi.Methods["newBid"].Outputs.Pack(bid)
	if err != nil {
		err = errFailedToPackOutput
		return
	}
	return result, nil

}

func (b *SuaveRuntimeAdapter) simulateBundle(input []byte) (res []byte, err error) {
	var (
		unpacked []interface{}
		result   []byte
	)

	_ = unpacked
	_ = result

	unpacked, err = artifacts.SuaveAbi.Methods["simulateBundle"].Inputs.Unpack(input)
	if err != nil {
		err = errFailedToUnpackInput
		return
	}

	var (
		bundleData []byte
	)

	bundleData = unpacked[0].([]byte)

	var (
		output1 uint64
	)

	if output1, err = b.impl.simulateBundle(bundleData); err != nil {
		return
	}

	result, err = artifacts.SuaveAbi.Methods["simulateBundle"].Outputs.Pack(output1)
	if err != nil {
		err = errFailedToPackOutput
		return
	}
	return result, nil

}

func (b *SuaveRuntimeAdapter) submitEthBlockBidToRelay(input []byte) (res []byte, err error) {
	var (
		unpacked []interface{}
		result   []byte
	)

	_ = unpacked
	_ = result

	unpacked, err = artifacts.SuaveAbi.Methods["submitEthBlockBidToRelay"].Inputs.Unpack(input)
	if err != nil {
		err = errFailedToUnpackInput
		return
	}

	var (
		relayUrl   string
		builderBid []byte
	)

	relayUrl = unpacked[0].(string)
	builderBid = unpacked[1].([]byte)

	var (
		output1 []byte
	)

	if output1, err = b.impl.submitEthBlockBidToRelay(relayUrl, builderBid); err != nil {
		return
	}

	result = output1
	return result, nil

}