package main

import (
	"fmt"
	"strconv"
	"testing"
)

// --- 動かないコード ---

// comparable と自己言及関数を持つインターフェイスを定義するとコンパイルエラーになる
// comparable を含む interface は自己言及関数を作成できない
type WrongGenericsType interface {
	fmt.Stringer
	comparable
	Less(WrongGenericsType) bool
}

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

type Instance1[L CollectGenericsType1] struct {
	// mapのキーに使える
	source map[L]string
}

// --- 気になること ---

// comparableを使わない自己言及インターフェイス
type CollectGenericsType2 interface {
	fmt.Stringer
	Less(CollectGenericsType2) bool
}

// これは普通に Generics の制約のinterfaceとして使える
type Instance2[L CollectGenericsType2] struct {
	// しかし、mapのキーには使えない
	source map[L]string
}

// --- 気になること ---

// comparableを使わない自己言及インターフェイス
type CollectGenericsType3[L any] interface {
	fmt.Stringer
	comparable
	Less(L) bool
}

type Instance3[T CollectGenericsType3[T]] struct {
	source map[T]string
}

type ID int

func (id ID) Less(other ID) bool {
	return id < other
}

func (id ID) String() string {
	return strconv.Itoa(int(id))
}

func TestHoge(t *testing.T) {

	instance3 := Instance3[ID]{source: map[ID]string{}}

	fmt.Printf("%#v\n", instance3)

}
