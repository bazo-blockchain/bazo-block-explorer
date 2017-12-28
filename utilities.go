package main

import (
  "fmt"
  "bytes"
  "encoding/binary"
  "golang.org/x/crypto/sha3"
  "github.com/mchetelat/bazo_miner/protocol"
)

func ConvertBlock(unconvertedBlock *protocol.Block) block {
  var convertedBlock block
  var convertedTxHash string

  //convertedBlock.Header = fmt.Sprintf("%x", unconvertedBlock.Header)
  convertedBlock.Hash = fmt.Sprintf("%x", unconvertedBlock.Hash)
  convertedBlock.PrevHash = fmt.Sprintf("%x", unconvertedBlock.PrevHash)
  //convertedBlock.Nonce = fmt.Sprintf("%x", unconvertedBlock.Nonce)
  convertedBlock.Timestamp = unconvertedBlock.Timestamp
  convertedBlock.MerkleRoot = fmt.Sprintf("%x", unconvertedBlock.MerkleRoot)
  convertedBlock.Beneficiary = fmt.Sprintf("%x", unconvertedBlock.Beneficiary)
  convertedBlock.NrFundsTx = unconvertedBlock.NrFundsTx
  convertedBlock.NrAccTx = unconvertedBlock.NrAccTx
  convertedBlock.NrConfigTx = unconvertedBlock.NrConfigTx
  for _, hash := range unconvertedBlock.FundsTxData {
    convertedTxHash = fmt.Sprintf("%x", hash)
    convertedBlock.FundsTxData = append(convertedBlock.FundsTxData, convertedTxHash)
  }
  for _, hash := range unconvertedBlock.AccTxData {
    convertedTxHash = fmt.Sprintf("%x", hash)
    convertedBlock.AccTxData = append(convertedBlock.AccTxData, convertedTxHash)
  }
  for _, hash := range unconvertedBlock.ConfigTxData {
    convertedTxHash = fmt.Sprintf("%x", hash)
    convertedBlock.ConfigTxData = append(convertedBlock.ConfigTxData, convertedTxHash)
  }

  return convertedBlock
}

func ConvertFundsTransaction(unconvertedTx *protocol.FundsTx, unconvertedBlockHash [32]byte, unconvertedTxHash [32]byte) fundstx {
  var convertedTx fundstx

  //convertedTx.Header = fmt.Sprintf("%x", unconvertedTx.Header)
  convertedTx.Hash = fmt.Sprintf("%x", unconvertedTxHash)
  convertedTx.BlockHash = fmt.Sprintf("%x", unconvertedBlockHash)
  convertedTx.Amount = unconvertedTx.Amount
  convertedTx.Fee = unconvertedTx.Fee
  convertedTx.TxCount = unconvertedTx.TxCnt
  convertedTx.From = fmt.Sprintf("%x", unconvertedTx.From)
  convertedTx.To = fmt.Sprintf("%x", unconvertedTx.To)
  convertedTx.Signature = fmt.Sprintf("%x", unconvertedTx.Sig)

  return convertedTx
}

func ConvertAccTransaction(unconvertedTx *protocol.AccTx, unconvertedBlockHash [32]byte, unconvertedTxHash [32]byte) acctx {
  var convertedTx acctx

  //convertedTx.Header = fmt.Sprintf("%x", unconvertedTx.Header)
  convertedTx.Hash = fmt.Sprintf("%x", unconvertedTxHash)
  convertedTx.BlockHash = fmt.Sprintf("%x", unconvertedBlockHash)
  convertedTx.Fee = unconvertedTx.Fee
  convertedTx.Issuer = fmt.Sprintf("%x", unconvertedTx.Issuer)
  convertedTx.PubKey = fmt.Sprintf("%x", unconvertedTx.PubKey)
  convertedTx.Signature = fmt.Sprintf("%x", unconvertedTx.Sig)

  return convertedTx
}

func ConvertConfigTransaction(unconvertedTx *protocol.ConfigTx, unconvertedBlockHash [32]byte, unconvertedTxHash [32]byte) configtx {
  var convertedTx configtx

  //convertedTx.Header = fmt.Sprintf("%x", unconvertedTx.Header)
  convertedTx.Hash = fmt.Sprintf("%x", unconvertedTxHash)
  convertedTx.BlockHash = fmt.Sprintf("%x", unconvertedBlockHash)
  convertedTx.Id = unconvertedTx.Id
  convertedTx.Fee = unconvertedTx.Fee
  convertedTx.Payload = unconvertedTx.Payload
  convertedTx.TxCount = unconvertedTx.TxCnt
  convertedTx.Signature = fmt.Sprintf("%x", unconvertedTx.Sig)

  return convertedTx
}

func ConvertOpenFundsTransaction(unconvertedTx *protocol.FundsTx, unconvertedTxHash [32]byte) fundstx {
  var convertedTx fundstx

  //convertedTx.Header = fmt.Sprintf("%x", unconvertedTx.Header)
  convertedTx.Hash = fmt.Sprintf("%x", unconvertedTxHash)
  convertedTx.Amount = unconvertedTx.Amount
  convertedTx.Fee = unconvertedTx.Fee
  convertedTx.TxCount = unconvertedTx.TxCnt
  convertedTx.From = fmt.Sprintf("%x", unconvertedTx.From)
  convertedTx.To = fmt.Sprintf("%x", unconvertedTx.To)
  convertedTx.Signature = fmt.Sprintf("%x", unconvertedTx.Sig)

  return convertedTx
}

func invertBlockArray(array []*protocol.Block) []*protocol.Block {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}

	return array
}

func SerializeHashContent(data interface{}) (hash [32]byte) {
	var buf bytes.Buffer

	binary.Write(&buf, binary.BigEndian, data)

	return sha3.Sum256(buf.Bytes())
}
