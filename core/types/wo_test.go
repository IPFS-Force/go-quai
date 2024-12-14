package types

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/dominant-strategies/go-quai/common"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

var locations = []common.Location{
	{0, 0},
	{0, 1},
	{0, 2},
	{0, 3},
	{1, 0},
	{1, 1},
	{1, 2},
	{1, 3},
	{2, 0},
	{2, 1},
	{2, 2},
	{2, 3},
	{3, 0},
	{3, 1},
	{3, 2},
	{3, 3},
}

func woTestData() (*WorkObject, common.Hash) {
	wo := &WorkObject{}
	wo.SetWorkObjectHeader(&WorkObjectHeader{})
	wo.woHeader.SetHeaderHash(EmptyHeader().Hash())
	wo.woHeader.SetParentHash(EmptyHeader().Hash())
	wo.woHeader.SetNumber(big.NewInt(1))
	wo.woHeader.SetDifficulty(big.NewInt(123456789))
	wo.woHeader.SetPrimeTerminusNumber(big.NewInt(42))
	wo.woHeader.SetTxHash(common.HexToHash("0x456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef3"))
	wo.woHeader.SetLocation(common.Location{0, 0})
	wo.woHeader.SetMixHash(common.HexToHash("0x56789abcdef0123456789abcdef0123456789abcdef0123456789abcdef4"))
	wo.woHeader.SetPrimaryCoinbase(common.HexToAddress("0x123456789abcdef0123456789abcdef0123456789", common.Location{0, 0}))
	wo.woHeader.SetTime(uint64(1))
	wo.woHeader.SetNonce(EncodeNonce(uint64(1)))
	wo.woHeader.SetLock(0)

	wo.woBody = EmptyWorkObjectBody()
	return wo, wo.Hash()
}

var (
	expectedWoHash         = common.HexToHash("0xd8e3ef0d1804c06495b219308535844169ccfdbd8565770077ea12f928fc8000")
	expectedUncleHash      = common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347")
	expectedPETXProtoBytes = []byte{
		0x0a, 0xc0, 0x01, 0x0a, 0x22, 0x0a, 0x20, 0x97, 0xb8, 0xd8, 0x2d, 0x3f, 0x97, 0x82, 0x7d, 0x2f,
		0x95, 0x8f, 0x53, 0xa5, 0x31, 0x4a, 0x3c, 0x36, 0xe5, 0x1c, 0x57, 0xb9, 0xbb, 0x77, 0x08, 0x80,
		0xb2, 0x48, 0x79, 0x5d, 0x40, 0xa0, 0x1e, 0x12, 0x22, 0x0a, 0x20, 0x97, 0xb8, 0xd8, 0x2d, 0x3f,
		0x97, 0x82, 0x7d, 0x2f, 0x95, 0x8f, 0x53, 0xa5, 0x31, 0x4a, 0x3c, 0x36, 0xe5, 0x1c, 0x57, 0xb9,
		0xbb, 0x77, 0x08, 0x80, 0xb2, 0x48, 0x79, 0x5d, 0x40, 0xa0, 0x1e, 0x1a, 0x01, 0x01, 0x22, 0x04,
		0x07, 0x5b, 0xcd, 0x15, 0x2a, 0x22, 0x0a, 0x20, 0x00, 0x04, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
		0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0, 0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
		0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf3, 0x30, 0x01, 0x3a, 0x04, 0x0a, 0x02, 0x00, 0x00,
		0x42, 0x22, 0x0a, 0x20, 0x00, 0x00, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0, 0x12, 0x34, 0x56, 0x78,
		0x9a, 0xbc, 0xde, 0xf0, 0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0, 0x12, 0x34, 0x56, 0x78,
		0x9a, 0xbc, 0xde, 0xf4, 0x48, 0x01, 0x52, 0x01, 0x2a, 0x58, 0x00, 0x62, 0x16, 0x0a, 0x14, 0x23,
		0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23,
		0x45, 0x67, 0x89, 0x12, 0x89, 0x05, 0x0a, 0x86, 0x05, 0x0a, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f,
		0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0,
		0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0x0a, 0x22, 0x0a,
		0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8,
		0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4,
		0x21, 0x12, 0x22, 0x0a, 0x20, 0x1d, 0xcc, 0x4d, 0xe8, 0xde, 0xc7, 0x5d, 0x7a, 0xab, 0x85, 0xb5,
		0x67, 0xb6, 0xcc, 0xd4, 0x1a, 0xd3, 0x12, 0x45, 0x1b, 0x94, 0x8a, 0x74, 0x13, 0xf0, 0xa1, 0x42,
		0xfd, 0x40, 0xd4, 0x93, 0x47, 0x1a, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55,
		0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad,
		0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0x22, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f,
		0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0,
		0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0x2a, 0x22, 0x0a,
		0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8,
		0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4,
		0x21, 0x32, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45,
		0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f,
		0xb5, 0xe3, 0x63, 0xb4, 0x21, 0x3a, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55,
		0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad,
		0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0x3a, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f,
		0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0,
		0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0x3a, 0x22, 0x0a,
		0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8,
		0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4,
		0x21, 0x42, 0x22, 0x0a, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x00, 0x52, 0x00, 0x52, 0x00, 0x5a, 0x00, 0x5a, 0x00, 0x5a,
		0x00, 0x62, 0x00, 0x62, 0x00, 0x62, 0x00, 0x6a, 0x00, 0x72, 0x00, 0x72, 0x00, 0x78, 0x00, 0x80,
		0x01, 0x00, 0x8a, 0x01, 0x00, 0x9a, 0x01, 0x00, 0xb2, 0x01, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f,
		0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0,
		0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0xba, 0x01, 0x22,
		0x0a, 0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0,
		0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63,
		0xb4, 0x21, 0xc0, 0x01, 0x00, 0xc8, 0x01, 0x00, 0xd0, 0x01, 0x00, 0xda, 0x01, 0x22, 0x0a, 0x20,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xe2, 0x01, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45,
		0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f,
		0xb5, 0xe3, 0x63, 0xb4, 0x21, 0xea, 0x01, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc,
		0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c,
		0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0xf0, 0x01, 0x00, 0xf8, 0x01, 0x00,
		0x82, 0x02, 0x00, 0x8a, 0x02, 0x00, 0x92, 0x02, 0x00, 0x9a, 0x02, 0x00, 0xa2, 0x02, 0x00,
	}
	expectedBlockProtoBytes = []byte{
		0x0a, 0xc0, 0x01, 0x0a, 0x22, 0x0a, 0x20, 0x97, 0xb8, 0xd8, 0x2d, 0x3f, 0x97, 0x82, 0x7d, 0x2f,
		0x95, 0x8f, 0x53, 0xa5, 0x31, 0x4a, 0x3c, 0x36, 0xe5, 0x1c, 0x57, 0xb9, 0xbb, 0x77, 0x08, 0x80,
		0xb2, 0x48, 0x79, 0x5d, 0x40, 0xa0, 0x1e, 0x12, 0x22, 0x0a, 0x20, 0x97, 0xb8, 0xd8, 0x2d, 0x3f,
		0x97, 0x82, 0x7d, 0x2f, 0x95, 0x8f, 0x53, 0xa5, 0x31, 0x4a, 0x3c, 0x36, 0xe5, 0x1c, 0x57, 0xb9,
		0xbb, 0x77, 0x08, 0x80, 0xb2, 0x48, 0x79, 0x5d, 0x40, 0xa0, 0x1e, 0x1a, 0x01, 0x01, 0x22, 0x04,
		0x07, 0x5b, 0xcd, 0x15, 0x2a, 0x22, 0x0a, 0x20, 0x00, 0x04, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
		0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0, 0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
		0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf3, 0x30, 0x01, 0x3a, 0x04, 0x0a, 0x02, 0x00, 0x00,
		0x42, 0x22, 0x0a, 0x20, 0x00, 0x00, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0, 0x12, 0x34, 0x56, 0x78,
		0x9a, 0xbc, 0xde, 0xf0, 0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0, 0x12, 0x34, 0x56, 0x78,
		0x9a, 0xbc, 0xde, 0xf4, 0x48, 0x01, 0x52, 0x01, 0x2a, 0x58, 0x00, 0x62, 0x16, 0x0a, 0x14, 0x23,
		0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23,
		0x45, 0x67, 0x89, 0x12, 0x93, 0x05, 0x0a, 0x86, 0x05, 0x0a, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f,
		0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0,
		0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0x0a, 0x22, 0x0a,
		0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8,
		0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4,
		0x21, 0x12, 0x22, 0x0a, 0x20, 0x1d, 0xcc, 0x4d, 0xe8, 0xde, 0xc7, 0x5d, 0x7a, 0xab, 0x85, 0xb5,
		0x67, 0xb6, 0xcc, 0xd4, 0x1a, 0xd3, 0x12, 0x45, 0x1b, 0x94, 0x8a, 0x74, 0x13, 0xf0, 0xa1, 0x42,
		0xfd, 0x40, 0xd4, 0x93, 0x47, 0x1a, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55,
		0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad,
		0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0x22, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f,
		0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0,
		0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0x2a, 0x22, 0x0a,
		0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8,
		0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4,
		0x21, 0x32, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45,
		0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f,
		0xb5, 0xe3, 0x63, 0xb4, 0x21, 0x3a, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55,
		0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad,
		0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0x3a, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f,
		0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0,
		0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0x3a, 0x22, 0x0a,
		0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8,
		0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4,
		0x21, 0x42, 0x22, 0x0a, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x00, 0x52, 0x00, 0x52, 0x00, 0x5a, 0x00, 0x5a, 0x00, 0x5a,
		0x00, 0x62, 0x00, 0x62, 0x00, 0x62, 0x00, 0x6a, 0x00, 0x72, 0x00, 0x72, 0x00, 0x78, 0x00, 0x80,
		0x01, 0x00, 0x8a, 0x01, 0x00, 0x9a, 0x01, 0x00, 0xb2, 0x01, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f,
		0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0,
		0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0xba, 0x01, 0x22,
		0x0a, 0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0,
		0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63,
		0xb4, 0x21, 0xc0, 0x01, 0x00, 0xc8, 0x01, 0x00, 0xd0, 0x01, 0x00, 0xda, 0x01, 0x22, 0x0a, 0x20,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xe2, 0x01, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45,
		0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f,
		0xb5, 0xe3, 0x63, 0xb4, 0x21, 0xea, 0x01, 0x22, 0x0a, 0x20, 0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc,
		0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c,
		0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, 0xf0, 0x01, 0x00, 0xf8, 0x01, 0x00,
		0x82, 0x02, 0x00, 0x8a, 0x02, 0x00, 0x92, 0x02, 0x00, 0x9a, 0x02, 0x00, 0xa2, 0x02, 0x00, 0x12,
		0x00, 0x1a, 0x00, 0x22, 0x00, 0x2a, 0x00, 0x32, 0x00,
	}
)

func TestWoHash(t *testing.T) {
	_, actualHash := woTestData()
	require.Equal(t, expectedWoHash, actualHash, "Hash not equal to expected hash")
}

func TestWoSealHash(t *testing.T) {
	testWo, _ := woTestData()
	actualHash := testWo.SealHash()
	expectedHash := common.HexToHash("0x83fd7f1dbd2f62d320c2dd974e2d1b3ce8108594a7c9e3aa99dfd993ca106090")
	require.Equal(t, expectedHash, actualHash, "Seal hash not equal to expected hash")
}

func FuzzHeaderHash(f *testing.F) {
	fuzzHash(f,
		func(woh *WorkObjectHeader) common.Hash { return woh.headerHash },
		func(woh *WorkObjectHeader, hash common.Hash) { woh.headerHash = hash })
}

func FuzzParentHash(f *testing.F) {
	fuzzHash(f,
		func(woh *WorkObjectHeader) common.Hash { return woh.parentHash },
		func(woh *WorkObjectHeader, hash common.Hash) { woh.parentHash = hash })
}

func FuzzDifficultyHash(f *testing.F) {
	fuzzBigInt(f,
		func(woh *WorkObjectHeader) *big.Int { return woh.difficulty },
		func(woh *WorkObjectHeader, val *big.Int) { woh.difficulty = val })
}

func FuzzNumberHash(f *testing.F) {
	fuzzBigInt(f,
		func(woh *WorkObjectHeader) *big.Int { return woh.number },
		func(woh *WorkObjectHeader, val *big.Int) { woh.number = val })
}

func FuzzTxHash(f *testing.F) {
	fuzzHash(f,
		func(woh *WorkObjectHeader) common.Hash { return woh.TxHash() },
		func(woh *WorkObjectHeader, hash common.Hash) { woh.SetTxHash(hash) })
}

func FuzzMixHash(f *testing.F) {
	fuzzHash(f,
		func(woh *WorkObjectHeader) common.Hash { return woh.MixHash() },
		func(woh *WorkObjectHeader, hash common.Hash) { woh.SetMixHash(hash) })
}

func TestLocationHash(t *testing.T) {
	wo, hash := woTestData()

	for _, loc := range locations[1:] {
		woCopy := *wo
		woCopy.woHeader.location = loc
		require.NotEqual(t, woCopy.Hash(), hash, "Hash equal for location \noriginal: %v, modified: %v", wo.woHeader.location, loc)
	}
}

func FuzzTimeHash(f *testing.F) {
	fuzzUint64Field(f,
		func(woh *WorkObjectHeader) uint64 { return woh.time },
		func(woh *WorkObjectHeader, time uint64) { woh.time = time })
}

func FuzzNonceHash(f *testing.F) {
	fuzzUint64Field(f,
		func(woh *WorkObjectHeader) uint64 { return woh.nonce.Uint64() },
		func(woh *WorkObjectHeader, nonce uint64) { woh.nonce = EncodeNonce(nonce) })
}

func TestCalcUncleHash(t *testing.T) {
	tests := []struct {
		uncleNum          int
		expectedUncleHash common.Hash
		expectedWoHash    common.Hash
		shouldPass        bool
	}{
		{
			uncleNum:          0,
			expectedUncleHash: expectedUncleHash,
			expectedWoHash:    expectedWoHash,
			shouldPass:        true,
		},
		{
			uncleNum:          1,
			expectedUncleHash: expectedUncleHash,
			expectedWoHash:    expectedWoHash,
			shouldPass:        false,
		},
		{
			uncleNum:          5,
			expectedUncleHash: common.HexToHash("0x3c9dd26495f9a6ddf36e1443bee2ff0a3bb59b0722a765b145d01ab1f78ccd44"),
			expectedWoHash:    common.HexToHash("0x67c4f50242b43f752a32574a28633ac08d316fdbd786ccdc675e90887643adc2"),
			shouldPass:        true,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(fmt.Sprintf("uncleNum=%d", tt.uncleNum), func(t *testing.T) {
			assertUncleHash(t, tt.uncleNum, tt.expectedUncleHash, tt.expectedWoHash, tt.shouldPass)
		})
	}
}

func TestProtoEncode(t *testing.T) {
	// Test data
	testWo, _ := woTestData()

	// Define test cases
	tests := []struct {
		name          string
		objectType    WorkObjectView
		expectedBytes []byte
	}{
		{
			name:          "PEtxObject",
			objectType:    PEtxObject,
			expectedBytes: expectedPETXProtoBytes,
		},
		{
			name:          "BlockObject",
			objectType:    BlockObject,
			expectedBytes: expectedBlockProtoBytes,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ProtoEncode the test WorkObject
			protoTestWo, err := testWo.ProtoEncode(tt.objectType)
			require.NoError(t, err)

			// Marshal to bytes
			protoTestWoBytes, err := proto.Marshal(protoTestWo)
			require.NoError(t, err)

			// Compare with expected bytes
			require.Equal(t, tt.expectedBytes, protoTestWoBytes)
		})
	}
}

func TestProtoDecode(t *testing.T) {
	_, testWoHash := woTestData()

	tests := []struct {
		name         string
		objectType   WorkObjectView
		testBytes    []byte
		expectedHash common.Hash
	}{
		{
			"PETX",
			PEtxObject,
			expectedPETXProtoBytes,
			testWoHash,
		},
		{
			"BlockObject",
			BlockObject,
			expectedBlockProtoBytes,
			testWoHash,
		},
		{
			"WorkShareObject",
			WorkShareObject,
			expectedBlockProtoBytes,
			testWoHash,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			protoWo := &ProtoWorkObject{}
			err := proto.Unmarshal(tt.testBytes, protoWo)
			require.NoError(t, err)

			decoded := &WorkObject{}
			decoded.ProtoDecode(protoWo, common.Location{0, 0}, tt.objectType)
			require.Equal(t, decoded.Hash(), tt.expectedHash)
		})
	}
}

func TestCopyWorkObject(t *testing.T) {
	originalWo, expectedHash := woTestData()

	newWo := CopyWorkObject(originalWo)

	require.Equal(t, expectedHash, newWo.Hash(), "Copied work object is different from new work object")

	// Test to make sure the copy doesn't modify original.
	newWo.WorkObjectHeader().SetLocation(common.Location{2, 2})
	require.NotEqual(t, expectedHash, newWo.Hash(), "WorkObject hash didn't change with a new location")
	require.Equal(t, originalWo.Hash(), expectedHash, "Copied WorkObject changed values from the original")
}

func TestNewWorkObject(t *testing.T) {
	// Verify that copy is same as original.
	originalWo, expectedHash := woTestData()
	newWo := NewWorkObject(originalWo.WorkObjectHeader(), originalWo.Body(), originalWo.Tx())

	require.Equal(t, expectedHash, newWo.Hash(), "NewWorkObject created a different WorkObject than the original")
}

func assertUncleHash(t *testing.T, uncleNum int, expectedUncleHash common.Hash, expectedWoHash common.Hash, shouldPass bool) {
	wo, _ := woTestData()
	wo.Body().uncles = make([]*WorkObjectHeader, uncleNum)
	for i := 0; i < uncleNum; i++ {
		uncle, _ := woTestData()
		wo.Body().uncles[i] = CopyWorkObjectHeader(uncle.WorkObjectHeader())
	}

	wo.Body().Header().SetUncleHash(CalcUncleHash(wo.Body().uncles))
	wo.woHeader.SetHeaderHash(wo.Body().header.Hash())

	if shouldPass {
		require.Equal(t, expectedUncleHash, wo.Header().UncleHash(), "Uncle hashes do not create the expected root hash")
		require.Equal(t, expectedWoHash, wo.Hash(), "Uncle hashes do not create the expected WorkObject hash")
	} else {
		require.NotEqual(t, expectedUncleHash, wo.Header().UncleHash(), "Uncle hashes do not create the expected root hash")
		require.NotEqual(t, expectedWoHash, wo.Hash(), "Uncle hashes do not create the expected WorkObject hash")
	}
}

func fuzzHash(f *testing.F, getField func(*WorkObjectHeader) common.Hash, setField func(*WorkObjectHeader, common.Hash)) {
	wo, _ := woTestData()
	f.Add(testByte)
	f.Add(getField(wo.woHeader).Bytes())
	f.Fuzz(func(t *testing.T, b []byte) {
		localWo, hash := woTestData()
		sc := common.BytesToHash(b)
		if getField(localWo.woHeader) != sc {
			setField(localWo.woHeader, sc)
			require.NotEqual(t, localWo.Hash(), hash, "Hash collision\noriginal: %v, modified: %v", getField(wo.woHeader).Bytes(), b)
		}
	})
}

func fuzzUint64Field(f *testing.F, getVal func(*WorkObjectHeader) uint64, setVal func(*WorkObjectHeader, uint64)) {
	wo, _ := woTestData()
	f.Add(testUInt64)
	f.Add(getVal(wo.woHeader))
	f.Fuzz(func(t *testing.T, i uint64) {
		localWo, hash := woTestData()
		if getVal(localWo.woHeader) != i {
			setVal(localWo.woHeader, i)
			require.NotEqual(t, localWo.Hash(), hash, "Hash collision\noriginal: %v, modified: %v", getVal(wo.woHeader), i)
		}
	})
}

func fuzzBigInt(f *testing.F, getVal func(*WorkObjectHeader) *big.Int, setVal func(*WorkObjectHeader, *big.Int)) {
	wo, _ := woTestData()
	f.Add(testInt64)
	f.Add(getVal(wo.woHeader).Int64())
	f.Fuzz(func(t *testing.T, i int64) {
		localWo, hash := woTestData()
		bi := big.NewInt(i)
		if getVal(localWo.woHeader).Cmp(bi) != 0 {
			setVal(localWo.woHeader, bi)
			require.NotEqual(t, localWo.Hash(), hash, "Hash collision\noriginal: %v, modified: %v", getVal(wo.woHeader), bi)
		}
	})
}
