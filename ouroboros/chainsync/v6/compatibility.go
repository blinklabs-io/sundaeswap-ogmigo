// Package v6 provides the Ogmios v6 chainsync API from the v7 module.
//
// The v6 wire representation is compatible with the v7 representation for
// the types exposed by this package. Type aliases keep the v6 API available
// without maintaining a second copy of the implementation.
package v6

import chainsync "github.com/SundaeSwap-finance/ogmigo/v7/ouroboros/chainsync"

type Block = chainsync.Block
type Nonce = chainsync.Nonce
type BlockSize = chainsync.BlockSize
type Protocol = chainsync.Protocol
type BlockIssuer = chainsync.BlockIssuer
type OpCert = chainsync.OpCert
type Kes = chainsync.Kes
type LeaderValue = chainsync.LeaderValue
type PointType = chainsync.PointType
type PointString = chainsync.PointString
type PointStruct = chainsync.PointStruct
type Point = chainsync.Point
type Points = chainsync.Points
type ProtocolVersion = chainsync.ProtocolVersion
type RollBackward = chainsync.RollBackward
type RollBackwardPoint = chainsync.RollBackwardPoint
type RollForward = chainsync.RollForward
type ResultFindIntersectionPraos = chainsync.ResultFindIntersectionPraos
type ResultError = chainsync.ResultError
type ResultNextBlockPraos = chainsync.ResultNextBlockPraos
type ResponsePraos = chainsync.ResponsePraos
type Signature = chainsync.Signature
type Tx = chainsync.Tx
type TxID = chainsync.TxID
type TxIn = chainsync.TxIn
type TxIns = chainsync.TxIns
type TxInID = chainsync.TxInID
type TxOut = chainsync.TxOut
type TxOuts = chainsync.TxOuts
type Datums = chainsync.Datums
type TxInQuery = chainsync.TxInQuery
type Witness = chainsync.Witness
type ValidityInterval = chainsync.ValidityInterval
type OgmiosAuxiliaryDataV6 = chainsync.OgmiosAuxiliaryDataV6
type OgmiosAuxiliaryDataLabelsV6 = chainsync.OgmiosAuxiliaryDataLabelsV6
type OgmiosMetadatumRecordV6 = chainsync.OgmiosMetadatumRecordV6
type OgmiosMetadatumKind = chainsync.OgmiosMetadatumKind
type OgmiosMetadatum = chainsync.OgmiosMetadatum
type OgmiosMetadatumMap = chainsync.OgmiosMetadatumMap
type ByronBlockBFT = chainsync.ByronBlockBFT
type ByronBlockEBB = chainsync.ByronBlockEBB
type ByronBlockDelegate = chainsync.ByronBlockDelegate
type ByronBlockIssuer = chainsync.ByronBlockIssuer
type ByronProtocol = chainsync.ByronProtocol
type ResultByronEBB = chainsync.ResultByronEBB
type ResultByronBFT = chainsync.ResultByronBFT
type ResponseByronEBB = chainsync.ResponseByronEBB
type ResponseByronBFT = chainsync.ResponseByronBFT

const (
	PointTypeString           = chainsync.PointTypeString
	PointTypeStruct           = chainsync.PointTypeStruct
	FindIntersectionMethod    = chainsync.FindIntersectionMethod
	NextBlockMethod           = chainsync.NextBlockMethod
	FindIntersectMethod       = chainsync.FindIntersectMethod
	RequestNextMethod         = chainsync.RequestNextMethod
	RollForwardString         = chainsync.RollForwardString
	RollBackwardString        = chainsync.RollBackwardString
	OgmiosMetadatumTagUnknown = chainsync.OgmiosMetadatumTagUnknown
	OgmiosMetadatumTagInt     = chainsync.OgmiosMetadatumTagInt
	OgmiosMetadatumTagString  = chainsync.OgmiosMetadatumTagString
	OgmiosMetadatumTagBytes   = chainsync.OgmiosMetadatumTagBytes
	OgmiosMetadatumTagList    = chainsync.OgmiosMetadatumTagList
	OgmiosMetadatumTagMap     = chainsync.OgmiosMetadatumTagMap
)

var Origin = chainsync.Origin

var NewTxID = chainsync.NewTxID
var GetMetadataDatums = chainsync.GetMetadataDatums
var GetMetadataDatumsV6 = chainsync.GetMetadataDatumsV6
var GetMetadataDatumMapV6 = chainsync.GetMetadataDatumMapV6
var ReconstructDatums = chainsync.ReconstructDatums
