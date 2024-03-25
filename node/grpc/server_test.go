package grpc_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"

	commonpb "github.com/Layr-Labs/eigenda/api/grpc/common"
	pb "github.com/Layr-Labs/eigenda/api/grpc/node"
	"github.com/Layr-Labs/eigenda/common"
	commonmock "github.com/Layr-Labs/eigenda/common/mock"
	"github.com/Layr-Labs/eigenda/core"
	core_mock "github.com/Layr-Labs/eigenda/core/mock"
	"github.com/Layr-Labs/eigenda/encoding"
	"github.com/Layr-Labs/eigenda/encoding/kzg"
	"github.com/Layr-Labs/eigenda/encoding/kzg/prover"
	"github.com/Layr-Labs/eigenda/encoding/kzg/verifier"
	"github.com/Layr-Labs/eigenda/node"
	"github.com/Layr-Labs/eigenda/node/grpc"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wealdtech/go-merkletree"
	"github.com/wealdtech/go-merkletree/keccak256"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"
)

var (
	encodedChunk = []byte{42, 255, 129, 3, 1, 1, 5, 67, 104, 117, 110, 107, 1, 255, 130, 0, 1, 2, 1, 6, 67, 111, 101, 102, 102, 115, 1, 255, 134, 0, 1, 5, 80, 114, 111, 111, 102, 1, 255, 136, 0, 0, 0, 25, 255, 133, 2, 1, 1, 10, 91, 93, 98, 110, 50, 53, 52, 46, 70, 114, 1, 255, 134, 0, 1, 255, 132, 0, 0, 18, 255, 131, 1, 1, 1, 2, 70, 114, 1, 255, 132, 0, 1, 6, 1, 8, 0, 0, 35, 255, 135, 3, 1, 1, 7, 71, 49, 80, 111, 105, 110, 116, 1, 255, 136, 0, 1, 2, 1, 1, 88, 1, 255, 138, 0, 1, 1, 89, 1, 255, 138, 0, 0, 0, 23, 255, 137, 1, 1, 1, 7, 69, 108, 101, 109, 101, 110, 116, 1, 255, 138, 0, 1, 6, 1, 8, 0, 0, 254, 4, 243, 255, 130, 1, 32, 4, 248, 186, 196, 96, 34, 212, 35, 97, 83, 248, 121, 9, 252, 220, 181, 118, 97, 134, 248, 186, 26, 225, 204, 191, 144, 133, 234, 248, 7, 223, 191, 156, 83, 115, 21, 36, 4, 248, 43, 196, 225, 43, 61, 88, 43, 49, 248, 28, 200, 121, 122, 178, 119, 200, 17, 248, 29, 172, 61, 194, 130, 114, 50, 171, 248, 33, 141, 185, 47, 11, 129, 128, 116, 4, 248, 246, 236, 255, 207, 43, 92, 176, 63, 248, 103, 179, 139, 80, 75, 57, 128, 89, 248, 107, 170, 70, 254, 95, 17, 101, 158, 248, 8, 106, 82, 82, 25, 78, 95, 104, 4, 248, 28, 125, 21, 116, 243, 255, 206, 10, 248, 153, 249, 156, 88, 61, 254, 171, 171, 248, 103, 66, 131, 8, 12, 165, 173, 173, 248, 36, 227, 189, 242, 180, 18, 171, 208, 4, 248, 19, 159, 205, 146, 86, 81, 57, 28, 248, 161, 130, 249, 92, 236, 82, 103, 4, 248, 84, 44, 63, 43, 249, 88, 187, 12, 248, 42, 121, 83, 118, 55, 127, 180, 134, 4, 248, 193, 39, 155, 110, 195, 113, 118, 46, 248, 47, 92, 162, 69, 188, 120, 94, 161, 248, 101, 214, 253, 103, 243, 8, 246, 176, 248, 41, 1, 238, 37, 43, 132, 228, 244, 4, 248, 70, 34, 194, 33, 68, 87, 108, 180, 248, 203, 230, 97, 137, 162, 177, 142, 23, 248, 101, 25, 216, 255, 137, 96, 240, 73, 248, 40, 50, 167, 154, 63, 108, 55, 240, 4, 248, 78, 40, 51, 224, 193, 131, 8, 90, 248, 162, 203, 245, 119, 83, 125, 219, 33, 248, 85, 109, 106, 231, 162, 152, 229, 110, 248, 38, 189, 66, 40, 176, 177, 114, 84, 4, 248, 193, 67, 43, 158, 218, 245, 83, 116, 248, 100, 165, 217, 161, 166, 209, 98, 172, 248, 231, 23, 45, 28, 225, 102, 143, 157, 248, 20, 12, 146, 122, 104, 126, 51, 235, 4, 248, 19, 118, 59, 144, 83, 246, 144, 229, 248, 203, 168, 161, 194, 137, 34, 191, 157, 248, 252, 196, 212, 78, 99, 166, 6, 225, 248, 29, 41, 54, 112, 125, 128, 240, 209, 4, 248, 24, 175, 53, 2, 113, 155, 113, 233, 248, 162, 189, 238, 198, 233, 31, 199, 239, 248, 205, 162, 128, 190, 163, 250, 181, 226, 248, 40, 205, 5, 117, 16, 49, 205, 45, 4, 248, 78, 49, 135, 21, 90, 93, 196, 50, 248, 115, 105, 77, 122, 222, 27, 224, 166, 248, 44, 0, 255, 63, 67, 184, 234, 235, 248, 45, 88, 39, 211, 138, 80, 43, 243, 4, 248, 244, 239, 154, 119, 68, 204, 215, 5, 248, 53, 82, 219, 150, 72, 243, 20, 147, 248, 141, 131, 101, 73, 11, 218, 234, 89, 248, 25, 246, 203, 17, 86, 91, 107, 199, 4, 248, 111, 106, 155, 101, 22, 163, 231, 214, 248, 86, 123, 235, 222, 87, 192, 80, 167, 248, 107, 38, 156, 175, 73, 123, 184, 189, 248, 23, 12, 154, 39, 153, 2, 158, 213, 4, 248, 40, 166, 62, 99, 6, 145, 128, 237, 248, 77, 160, 235, 64, 123, 181, 120, 66, 248, 116, 0, 126, 221, 26, 18, 100, 74, 248, 46, 92, 161, 252, 177, 177, 191, 127, 4, 248, 227, 144, 223, 154, 232, 249, 22, 233, 248, 53, 82, 148, 149, 84, 76, 107, 93, 248, 71, 251, 7, 58, 156, 200, 102, 4, 248, 3, 147, 75, 172, 199, 222, 109, 87, 4, 248, 169, 207, 109, 252, 37, 85, 158, 78, 248, 237, 12, 207, 255, 117, 62, 171, 3, 248, 43, 93, 155, 238, 136, 102, 150, 139, 248, 40, 174, 6, 46, 62, 50, 174, 104, 4, 248, 156, 217, 228, 156, 76, 202, 37, 121, 248, 80, 44, 200, 177, 237, 112, 103, 44, 248, 211, 172, 202, 164, 34, 242, 190, 204, 248, 15, 241, 94, 33, 88, 13, 34, 66, 4, 248, 198, 229, 9, 111, 155, 117, 84, 125, 248, 69, 115, 47, 6, 35, 132, 39, 86, 248, 243, 113, 79, 216, 240, 35, 72, 75, 248, 7, 29, 38, 85, 134, 106, 213, 236, 4, 248, 8, 8, 251, 11, 97, 66, 8, 55, 248, 159, 67, 100, 214, 31, 167, 88, 221, 248, 151, 110, 49, 190, 136, 249, 55, 217, 248, 47, 94, 78, 30, 0, 220, 176, 125, 4, 248, 246, 81, 132, 144, 151, 161, 113, 102, 248, 229, 8, 10, 180, 28, 223, 222, 8, 248, 158, 88, 212, 24, 77, 31, 96, 232, 248, 41, 65, 45, 216, 25, 224, 221, 4, 4, 248, 11, 189, 86, 122, 64, 254, 107, 253, 248, 242, 174, 32, 144, 43, 116, 187, 77, 248, 16, 163, 127, 128, 4, 233, 82, 168, 248, 4, 90, 126, 233, 232, 220, 81, 74, 4, 248, 54, 17, 20, 36, 220, 10, 168, 78, 248, 77, 61, 41, 4, 95, 154, 130, 70, 248, 37, 180, 163, 188, 242, 88, 81, 28, 248, 37, 195, 179, 103, 195, 0, 252, 30, 4, 248, 148, 154, 198, 22, 110, 201, 164, 240, 248, 242, 100, 163, 103, 30, 185, 139, 205, 248, 198, 168, 87, 116, 135, 219, 11, 230, 248, 43, 163, 196, 37, 51, 32, 130, 241, 4, 248, 160, 22, 80, 69, 111, 126, 3, 23, 248, 76, 89, 182, 79, 244, 245, 155, 42, 248, 144, 203, 89, 203, 85, 216, 109, 139, 248, 36, 125, 246, 94, 210, 7, 236, 50, 4, 248, 244, 42, 154, 219, 137, 78, 64, 167, 248, 73, 57, 191, 50, 122, 120, 124, 249, 248, 192, 102, 139, 159, 135, 150, 18, 35, 248, 40, 167, 252, 247, 112, 215, 52, 61, 4, 248, 151, 181, 121, 81, 121, 147, 227, 13, 248, 236, 181, 178, 176, 243, 4, 136, 195, 248, 62, 97, 145, 239, 166, 114, 175, 107, 248, 23, 91, 75, 217, 198, 192, 155, 92, 4, 248, 182, 191, 150, 70, 229, 96, 122, 14, 248, 134, 0, 111, 72, 36, 162, 244, 220, 248, 168, 72, 14, 253, 239, 166, 139, 197, 248, 44, 139, 158, 151, 191, 127, 27, 222, 4, 248, 74, 171, 39, 27, 36, 31, 102, 30, 248, 41, 77, 140, 191, 229, 182, 30, 16, 248, 219, 194, 193, 143, 239, 141, 47, 73, 248, 23, 1, 236, 49, 51, 57, 155, 228, 4, 248, 128, 145, 254, 105, 104, 55, 224, 206, 248, 195, 70, 112, 120, 42, 171, 202, 23, 248, 242, 232, 247, 249, 215, 77, 208, 121, 248, 29, 0, 45, 26, 151, 224, 199, 214, 4, 248, 235, 253, 108, 246, 112, 139, 56, 187, 248, 214, 211, 157, 43, 210, 247, 57, 203, 248, 150, 28, 35, 231, 169, 220, 146, 139, 248, 48, 54, 207, 130, 116, 140, 125, 197, 4, 248, 23, 120, 154, 57, 66, 85, 149, 5, 248, 170, 172, 192, 127, 230, 130, 224, 17, 248, 117, 98, 19, 140, 134, 78, 47, 98, 248, 40, 206, 62, 254, 165, 238, 160, 130, 1, 1, 4, 248, 164, 40, 240, 180, 149, 114, 87, 82, 248, 195, 115, 109, 187, 95, 132, 65, 10, 248, 176, 59, 100, 197, 207, 37, 161, 253, 248, 10, 19, 137, 98, 39, 77, 128, 20, 1, 4, 248, 213, 212, 69, 58, 138, 39, 69, 249, 248, 99, 187, 162, 108, 114, 239, 78, 157, 248, 62, 166, 165, 148, 83, 202, 37, 169, 248, 47, 253, 18, 76, 216, 168, 22, 21, 0, 0}
	chainState   *core_mock.ChainDataMock
	opID         [32]byte
)

func TestMain(m *testing.M) {
	chainState, _ = core_mock.MakeChainDataMock(core.OperatorIndex(4))
	os.Exit(m.Run())
}

// makeTestVerifier makes a verifier currently using the only supported backend.
func makeTestComponents() (encoding.Prover, encoding.Verifier, error) {

	config := &kzg.KzgConfig{
		G1Path:          "../../inabox/resources/kzg/g1.point.300000",
		G2Path:          "../../inabox/resources/kzg/g2.point.300000",
		CacheDir:        "../../inabox/resources/kzg/SRSTables",
		SRSOrder:        300000,
		SRSNumberToLoad: 300000,
		NumWorker:       uint64(runtime.GOMAXPROCS(0)),
	}

	p, err := prover.NewProver(config, true)
	if err != nil {
		return nil, nil, err
	}

	v, err := verifier.NewVerifier(config, true)
	if err != nil {
		return nil, nil, err
	}

	return p, v, nil
}

func newTestServer(t *testing.T, mockValidator bool) *grpc.Server {
	dbPath := t.TempDir()
	keyPair, err := core.GenRandomBlsKeys()
	if err != nil {
		panic("failed to create a BLS Key")
	}
	opID = [32]byte{}
	copy(opID[:], []byte(fmt.Sprintf("%d", 3)))
	config := &node.Config{
		Timeout:                   10 * time.Second,
		ExpirationPollIntervalSec: 1,
		QuorumIDList:              []core.QuorumID{0},
		DbPath:                    dbPath,
		ID:                        opID,
		NumBatchValidators:        runtime.GOMAXPROCS(0),
	}
	loggerConfig := common.DefaultLoggerConfig()
	logger, err := common.NewLogger(loggerConfig)
	if err != nil {
		panic("failed to create a logger")
	}

	err = os.MkdirAll(config.DbPath, os.ModePerm)
	if err != nil {
		panic("failed to create a directory for db")
	}
	noopMetrics := metrics.NewNoopMetrics()
	reg := prometheus.NewRegistry()
	metrics := node.NewMetrics(noopMetrics, reg, logger, ":9090")
	store, err := node.NewLevelDBStore(dbPath, logger, metrics, 1e9, 1e9)
	if err != nil {
		panic("failed to create a new levelDB store")
	}
	defer os.Remove(dbPath)

	ratelimiter := &commonmock.NoopRatelimiter{}

	var val core.ShardValidator

	if mockValidator {
		mockVal := core_mock.NewMockShardValidator()
		mockVal.On("ValidateBlob", mock.Anything, mock.Anything).Return(nil)
		mockVal.On("ValidateBatch", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		val = mockVal
	} else {

		_, v, err := makeTestComponents()
		if err != nil {
			panic("failed to create test encoder")
		}

		asn := &core.StdAssignmentCoordinator{}

		cst, err := core_mock.MakeChainDataMock(core.OperatorIndex(10))
		if err != nil {
			panic("failed to create test encoder")
		}

		val = core.NewShardValidator(v, asn, cst, opID)
	}

	node := &node.Node{
		Config:     config,
		Logger:     logger,
		KeyPair:    keyPair,
		Metrics:    metrics,
		Store:      store,
		ChainState: chainState,
		Validator:  val,
	}
	return grpc.NewServer(config, node, logger, ratelimiter)
}

func makeStoreChunksRequest(t *testing.T, quorumThreshold, adversaryThreshold uint8) (*pb.StoreChunksRequest, [32]byte, [32]byte, []*core.BlobHeader, []*pb.BlobHeader) {
	var commitX, commitY fp.Element
	_, err := commitX.SetString("21661178944771197726808973281966770251114553549453983978976194544185382599016")
	assert.NoError(t, err)
	_, err = commitY.SetString("9207254729396071334325696286939045899948985698134704137261649190717970615186")
	assert.NoError(t, err)

	commitment := &encoding.G1Commitment{
		X: commitX,
		Y: commitY,
	}
	var lengthXA0, lengthXA1, lengthYA0, lengthYA1 fp.Element
	_, err = lengthXA0.SetString("10857046999023057135944570762232829481370756359578518086990519993285655852781")
	assert.NoError(t, err)
	_, err = lengthXA1.SetString("11559732032986387107991004021392285783925812861821192530917403151452391805634")
	assert.NoError(t, err)
	_, err = lengthYA0.SetString("8495653923123431417604973247489272438418190587263600148770280649306958101930")
	assert.NoError(t, err)
	_, err = lengthYA1.SetString("4082367875863433681332203403145435568316851327593401208105741076214120093531")
	assert.NoError(t, err)

	var lengthProof, lengthCommitment encoding.G2Commitment
	lengthProof.X.A0 = lengthXA0
	lengthProof.X.A1 = lengthXA1
	lengthProof.Y.A0 = lengthYA0
	lengthProof.Y.A1 = lengthYA1

	lengthCommitment = lengthProof

	quorumHeader := &core.BlobQuorumInfo{
		SecurityParam: core.SecurityParam{
			QuorumID:              0,
			ConfirmationThreshold: quorumThreshold,
			AdversaryThreshold:    adversaryThreshold,
		},
		ChunkLength: 10,
	}

	quorumHeader1 := &core.BlobQuorumInfo{
		SecurityParam: core.SecurityParam{
			QuorumID:              1,
			ConfirmationThreshold: 65,
			AdversaryThreshold:    15,
		},
		ChunkLength: 10,
	}

	blobHeaders := []*core.BlobHeader{
		{
			BlobCommitments: encoding.BlobCommitments{
				Commitment:       commitment,
				LengthCommitment: &lengthCommitment,
				LengthProof:      &lengthProof,
				Length:           48,
			},
			QuorumInfos: []*core.BlobQuorumInfo{quorumHeader, quorumHeader1},
		},
		{
			BlobCommitments: encoding.BlobCommitments{
				Commitment:       commitment,
				LengthCommitment: &lengthCommitment,
				LengthProof:      &lengthProof,
				Length:           50,
			},
			QuorumInfos: []*core.BlobQuorumInfo{quorumHeader},
		},
	}
	batchHeader := core.BatchHeader{
		BatchRoot:            [32]byte{0},
		ReferenceBlockNumber: 1,
	}

	_, err = batchHeader.SetBatchRoot(blobHeaders)
	assert.NoError(t, err)

	batchHeaderHash, err := batchHeader.GetBatchHeaderHash()
	assert.NoError(t, err)

	blobHeaderProto0 := blobHeaderToProto(blobHeaders[0])
	blobHeaderProto1 := blobHeaderToProto(blobHeaders[1])

	req := &pb.StoreChunksRequest{
		BatchHeader: &pb.BatchHeader{
			BatchRoot:            batchHeader.BatchRoot[:],
			ReferenceBlockNumber: uint32(batchHeader.ReferenceBlockNumber),
		},
		Blobs: []*pb.Blob{
			{
				Header: blobHeaderProto0,
				Bundles: []*pb.Bundle{
					{
						Chunks: [][]byte{encodedChunk},
					},
					{
						// Empty bundle for the second quorum
						Chunks: [][]byte{},
					},
				},
			},
			{
				Header: blobHeaderProto1,
				Bundles: []*pb.Bundle{
					{
						Chunks: [][]byte{encodedChunk},
					},
				},
			},
		},
	}
	blobHeadersProto := []*pb.BlobHeader{blobHeaderProto0, blobHeaderProto1}

	return req, batchHeaderHash, batchHeader.BatchRoot, blobHeaders, blobHeadersProto
}

func storeChunks(t *testing.T, server *grpc.Server) ([32]byte, [32]byte, []*core.BlobHeader, []*pb.BlobHeader) {
	adversaryThreshold := uint8(90)
	quorumThreshold := uint8(100)
	req, batchHeaderHash, batchRoot, blobHeaders, blobHeadersProto := makeStoreChunksRequest(t, quorumThreshold, adversaryThreshold)

	reply, err := server.StoreChunks(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, reply.GetSignature())

	return batchHeaderHash, batchRoot, blobHeaders, blobHeadersProto
}

func TestStoreChunksRequestValidation(t *testing.T) {
	server := newTestServer(t, true)

	req, _, _, _, _ := makeStoreChunksRequest(t, 66, 33)
	req.BatchHeader = nil
	_, err := server.StoreChunks(context.Background(), req)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "missing batch_header"))

	req, _, _, _, _ = makeStoreChunksRequest(t, 66, 33)
	req.BatchHeader.BatchRoot = nil
	_, err = server.StoreChunks(context.Background(), req)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "missing batch_root in request"))

	req, _, _, _, _ = makeStoreChunksRequest(t, 66, 33)
	req.BatchHeader.ReferenceBlockNumber = 0
	_, err = server.StoreChunks(context.Background(), req)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "missing reference_block_number in request"))

	req, _, _, _, _ = makeStoreChunksRequest(t, 66, 33)
	req.Blobs = nil
	_, err = server.StoreChunks(context.Background(), req)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "missing blobs in request"))

	req, _, _, _, _ = makeStoreChunksRequest(t, 66, 33)
	req.Blobs[0].Header = nil
	_, err = server.StoreChunks(context.Background(), req)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "missing blob header in request"))

	req, _, _, _, _ = makeStoreChunksRequest(t, 66, 33)
	req.Blobs[0].Header.QuorumHeaders = nil
	_, err = server.StoreChunks(context.Background(), req)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "missing quorum headers in request"))

	req, _, _, _, _ = makeStoreChunksRequest(t, 66, 66)
	_, err = server.StoreChunks(context.Background(), req)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "confirmation threshold must be >= 10 + adversary threshold"))

	req, _, _, _, _ = makeStoreChunksRequest(t, 101, 66)
	_, err = server.StoreChunks(context.Background(), req)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "confimration threshold exceeds 100"))

	req, _, _, _, _ = makeStoreChunksRequest(t, 66, 0)
	_, err = server.StoreChunks(context.Background(), req)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "adversary threshold equals 0"))
}

func TestRetrieveChunks(t *testing.T) {
	server := newTestServer(t, true)
	batchHeaderHash, _, _, _ := storeChunks(t, server)

	p := &peer.Peer{
		Addr: &net.TCPAddr{
			IP:   net.ParseIP("0.0.0.0"),
			Port: 3000,
		},
	}
	ctx := peer.NewContext(context.Background(), p)

	retrievalReply, err := server.RetrieveChunks(ctx, &pb.RetrieveChunksRequest{
		BatchHeaderHash: batchHeaderHash[:],
		BlobIndex:       0,
		QuorumId:        0,
	})
	assert.NoError(t, err)
	recovered, err := new(encoding.Frame).Deserialize(retrievalReply.GetChunks()[0])
	assert.NoError(t, err)
	chunk, err := new(encoding.Frame).Deserialize(encodedChunk)
	assert.NoError(t, err)
	assert.Equal(t, recovered, chunk)

	retrievalReply, err = server.RetrieveChunks(ctx, &pb.RetrieveChunksRequest{
		BatchHeaderHash: batchHeaderHash[:],
		BlobIndex:       0,
		QuorumId:        1,
	})
	assert.NoError(t, err)
	assert.Empty(t, retrievalReply.GetChunks())
}

// If a batch fails to validate, it should not be stored in the store.
func TestRevertInvalidBatch(t *testing.T) {
	// This will fail the validation because the quorum threshold cannot be greater than 100.
	quorumThreshold := uint8(100)
	adversaryThreshold := uint8(100)
	req, batchHeaderHash, _, _, _ := makeStoreChunksRequest(t, quorumThreshold, adversaryThreshold)

	// Fail to store chunks, because invalid adversaryThreshold.
	server := newTestServer(t, false)
	_, err := server.StoreChunks(context.Background(), req)
	assert.Error(t, err)

	// Fail to get the blob header, because invalid batch will not be stored.
	_, err = server.GetBlobHeader(context.Background(), &pb.GetBlobHeaderRequest{
		BatchHeaderHash: batchHeaderHash[:],
		BlobIndex:       0,
		QuorumId:        0,
	})
	assert.Error(t, err)
}

func TestGetBlobHeader(t *testing.T) {
	server := newTestServer(t, true)
	batchHeaderHash, batchRoot, blobHeaders, protoBlobHeaders := storeChunks(t, server)
	reply, err := server.GetBlobHeader(context.Background(), &pb.GetBlobHeaderRequest{
		BatchHeaderHash: batchHeaderHash[:],
		BlobIndex:       0,
		QuorumId:        0,
	})
	assert.NoError(t, err)

	actual := reply.GetBlobHeader()
	expected := protoBlobHeaders[0]
	assert.True(t, proto.Equal(expected, actual))

	blobHeaderHash, err := blobHeaders[0].GetBlobHeaderHash()
	assert.NoError(t, err)

	proof := &merkletree.Proof{
		Hashes: reply.GetProof().GetHashes(),
		Index:  uint64(reply.GetProof().GetIndex()),
	}

	ok, err := merkletree.VerifyProofUsing(blobHeaderHash[:], false, proof, [][]byte{batchRoot[:]}, keccak256.New())
	assert.NoError(t, err)
	assert.True(t, ok)

	// Get blob header for the second quorum
	reply, err = server.GetBlobHeader(context.Background(), &pb.GetBlobHeaderRequest{
		BatchHeaderHash: batchHeaderHash[:],
		BlobIndex:       0,
		QuorumId:        1,
	})
	assert.NoError(t, err)

	actual = reply.GetBlobHeader()
	expected = protoBlobHeaders[0]
	assert.True(t, proto.Equal(expected, actual))

	blobHeaderHash, err = blobHeaders[0].GetBlobHeaderHash()
	assert.NoError(t, err)

	proof = &merkletree.Proof{
		Hashes: reply.GetProof().GetHashes(),
		Index:  uint64(reply.GetProof().GetIndex()),
	}

	ok, err = merkletree.VerifyProofUsing(blobHeaderHash[:], false, proof, [][]byte{batchRoot[:]}, keccak256.New())
	assert.NoError(t, err)
	assert.True(t, ok)
}

func blobHeaderToProto(blobHeader *core.BlobHeader) *pb.BlobHeader {
	var lengthCommitment, lengthProof pb.G2Commitment
	if blobHeader.LengthCommitment != nil {
		lengthCommitment.XA0 = blobHeader.LengthCommitment.X.A0.Marshal()
		lengthCommitment.XA1 = blobHeader.LengthCommitment.X.A1.Marshal()
		lengthCommitment.YA0 = blobHeader.LengthCommitment.Y.A0.Marshal()
		lengthCommitment.YA1 = blobHeader.LengthCommitment.Y.A1.Marshal()
	}
	if blobHeader.LengthProof != nil {
		lengthProof.XA0 = blobHeader.LengthProof.X.A0.Marshal()
		lengthProof.XA1 = blobHeader.LengthProof.X.A1.Marshal()
		lengthProof.YA0 = blobHeader.LengthProof.Y.A0.Marshal()
		lengthProof.YA1 = blobHeader.LengthProof.Y.A1.Marshal()
	}
	quorumHeaders := make([]*pb.BlobQuorumInfo, len(blobHeader.QuorumInfos))
	for i, quorumInfo := range blobHeader.QuorumInfos {
		quorumHeaders[i] = &pb.BlobQuorumInfo{
			QuorumId:              uint32(quorumInfo.QuorumID),
			ConfirmationThreshold: uint32(quorumInfo.ConfirmationThreshold),
			AdversaryThreshold:    uint32(quorumInfo.AdversaryThreshold),
			ChunkLength:           uint32(quorumInfo.ChunkLength),
		}
	}

	return &pb.BlobHeader{
		Commitment: &commonpb.G1Commitment{
			X: blobHeader.Commitment.X.Marshal(),
			Y: blobHeader.Commitment.Y.Marshal(),
		},
		LengthCommitment: &lengthCommitment,
		LengthProof:      &lengthProof,
		Length:           uint32(blobHeader.Length),
		QuorumHeaders:    quorumHeaders,
	}
}
