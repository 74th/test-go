package main

import (
	"log"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestRun(t *testing.T) {
	tm := time.Time{}
	log.Printf("time.Time(): %v", tm)

	ptm := &timestamppb.Timestamp{}
	log.Printf("time.Time(): %v", ptm.AsTime())
	log.Printf("isValid(): %v", ptm.IsValid())
	log.Print(len(ptm.String()))
}

func TestRun2(t *testing.T) {
	tm, _ := time.Parse(time.RFC3339, "1970-01-01T00:00:00Z")
	log.Printf("time.Time(): %v", tm)
	ptm := timestamppb.New(tm)
	log.Printf("time.Time(): %v", ptm.AsTime())
	log.Printf("isValid(): %v", ptm.IsValid())
	log.Print(len(ptm.String()))

	ptm = timestamppb.Now()
	log.Printf("time.Time(): %v", ptm.AsTime())
	log.Printf("isValid(): %v", ptm.IsValid())
	log.Print(len(ptm.String()))
}
