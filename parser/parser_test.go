package parser

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"testing"
)

func TestParser(t *testing.T) {
	b := bytes.Buffer{}
	d, _ := hex.DecodeString("42007801000001484200770100000038420069010000002042006A0200000004000000010000000042006B0200000004000000020000000042000D0200000004000000020000000042000F010000007842005C05000000040000000D000000004200930800000008606051f958d79b0f4200790100000050420094070000001424554E495155455F4944454E5449464945525F3000000000420008010000002842000A070000000F41637469766174696F6E20446174650042000B0900000008000000005B224CB142000F010000008042005C05000000040000000D0000000042009308000000087cb12802f6a52cf14200790100000058420094070000001424554E495155455F4944454E5449464945525F3000000000420008010000003042000A0700000011446561637469766174696F6E20446174650000000000000042000B0900000008000000005B224D29")
	b.Write(d)
	Parser(&b)



}



func TestParser_Marshal(t *testing.T) {

	k := &KmipStructResponse{}
	k.ResponseMessage = &ResponseMessage{}
	k.ResponseMessage.ResponseHeader = &ResponseHeader{}
	var kk kmipbin.KmipInt
	kk = kmipbin.KmipInt(7)
	k.ResponseMessage.ResponseHeader.BatchCount = &kk
	var gg = kmipbin.KmipByteString{0x40,0x41, 0x40,0x41, 0x40,0x41}
	k.ResponseMessage.BatchItem = append(k.ResponseMessage.BatchItem , &BatchItem{UniqueBatchItemID:&gg} )
	b := MarshalAllT(k)

	h := hex.EncodeToString(b)
	fmt.Println(h)



}
