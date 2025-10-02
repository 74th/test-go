package main

import (
	"fmt"
	"testing"
)

// --- 回避策 ---

// 自己言及インターフェイスを別に切り出して
type subType interface {
	fmt.Stringer
	Less(subType) bool
}

// comparable と組み合わせると使える
type CollectGenericsType1 interface {
	subType
	comparable
}

type Insance1[L CollectGenericsType1] struct {
	sources map[L]string
}

// --- 気になること ---

// comparableを使わない自己言及インターフェイス
type CollectGenericsType2 interface {
	Less(CollectGenericsType2) bool
}

// これは普通に Generics の制約のinterfaceとして使える
type Insance2[L CollectGenericsType2] struct {
	sources map[L]string
}

func TestHoge(t *testing.T) {

}
