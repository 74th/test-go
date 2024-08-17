package main

import (
	"log"
	"testing"

	"github.com/74th/testing-go/20240802-protobuf-nested-message/pb"
	"google.golang.org/protobuf/proto"
)

func TestNestedProtoBuf(t *testing.T) {
	input := &pb.Record{
		Id: 1,
		// ネストされたメッセージは、Record_Optionとアンスコで区切られて定義される
		Option: &pb.Record_Option{
			// enumはenumの型名は付かない
			// 型の定義はされている
			OptionCode: pb.Record_OPTION_1,
		},
		SubOption: &pb.Record_SubOption{
			// メッセージの中に含まれる分には、アンスコで区切られる
			OptionCode: pb.Record_SubOption_OPTION_1,
		},
	}

	b, err := proto.Marshal(input)
	if err != nil {
		t.Errorf("Failed to marshal record: %v", err)
		return
	}
	if len(b) == 0 {
		t.Errorf("Empty marshaled data")
		return
	}

	log.Printf("Marshaled data: %v", b)

	output1 := &pb.Record{}
	err = proto.Unmarshal(b, output1)
	if err != nil {
		t.Errorf("Failed to unmarshal record: %v", err)
		return
	}
	if !proto.Equal(input, output1) {
		t.Errorf("unmatch : %#v, %#v", input, output1)
	}
}
