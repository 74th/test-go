package main

import (
	"log"
	"testing"
	"time"

	"github.com/74th/testing-go/20240730-protobuf-optional/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestHoge(t *testing.T) {

	ts1, _ := time.Parse(time.RFC3339, "2024-07-30T13:00:00Z")
	ts2, _ := time.Parse(time.RFC3339, "2024-07-30T14:00:00Z")
	id1 := int32(100)
	id2 := int32(200)

	input1 := &pb.Record{
		Id:    id1,
		OptId: &id2,
		Ts:    timestamppb.New(ts1),
		OptTs: timestamppb.New(ts2),
	}

	b, err := proto.Marshal(input1)
	if err != nil {
		t.Errorf("Failed to marshal record: %v", err)
		return
	}
	if len(b) == 0 {
		t.Errorf("Empty marshaled data")
		return
	}

	log.Printf("Marshaled data: %v", b)

	// Optionalを含むもの
	output1 := &pb.Record{}
	err = proto.Unmarshal(b, output1)
	if err != nil {
		t.Errorf("Failed to unmarshal record: %v", err)
		return
	}

	if output1.GetId() != id1 {
		t.Errorf("Invalid id: %v", output1.GetId())
	}
	if output1.GetOptId() != id2 {
		t.Errorf("Invalid opt_id: %v", output1.GetOptId())
	}
	if output1.OptId == nil {
		t.Errorf("Invalid opt_id: %v", output1.GetOptId())
	}
	if output1.GetTs().AsTime() != ts1 {
		t.Errorf("Invalid ts: %v", output1.GetTs())
	}
	if output1.GetOptTs().AsTime() != ts2 {
		t.Errorf("Invalid opt_ts: %v", output1.GetOptTs())
	}

	input2 := &pb.Record{
		Id:    id1,
		OptId: nil,
		Ts:    timestamppb.New(ts1),
		OptTs: nil,
	}

	b, err = proto.Marshal(input2)
	if err != nil {
		t.Errorf("Failed to marshal record: %v", err)
		return
	}
	if len(b) == 0 {
		t.Errorf("Empty marshaled data")
		return
	}

	log.Printf("Marshaled data: %v", b)

	output2 := &pb.Record{}
	err = proto.Unmarshal(b, output2)
	if err != nil {
		t.Errorf("Failed to unmarshal record: %v", err)
		return
	}

	if output2.GetId() != id1 {
		t.Errorf("Invalid id: %v", output2.GetId())
	}
	if output2.GetOptId() != 0 {
		// GetXX() ではゼロ値と見分けが付かない
		t.Errorf("Invalid opt_id: %v", output2.GetOptId())
	}
	if output2.OptId != nil {
		t.Errorf("Invalid opt_id: %v", output2.OptId)
	}
	if output2.GetTs().AsTime() != ts1 {
		t.Errorf("Invalid ts: %v", output2.GetTs())
	}
	if output2.GetOptTs() != nil {
		t.Errorf("Invalid opt_ts: %v", output1.GetOptTs())
	}
	log.Printf("IsZeroではnilは評価できない AsTime():%v AsTime().IsZero():%t", output2.GetOptTs().AsTime(), output2.GetOptTs().AsTime().IsZero())
}
