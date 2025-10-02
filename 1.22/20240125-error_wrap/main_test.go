package main

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

type SpecialError struct {
	when string
}

func (e SpecialError) Error() string {
	if len(e.when) == 0 {
		return "special error"
	}
	return fmt.Sprintf("special error when %s", e.when)
}

func NewSpecialError(when string) SpecialError {
	return SpecialError{
		when: when,
	}
}

// Unwrap() 可能なエラー
type WrappingError struct {
	wrappedErr error
}

func (e WrappingError) Error() string { return fmt.Sprintf("wrapping error: %s", e.wrappedErr.Error()) }

func (e WrappingError) Unwrap() error { return e.wrappedErr }

func NewWrappingError(err error) WrappingError {
	return WrappingError{
		wrappedErr: err,
	}
}

func TestErrorIs(t *testing.T) {
	// ラップエラー
	deadlineWrapErr := fmt.Errorf("wrap error: %w", context.DeadlineExceeded)

	// ラップ
	fmt.Print("Unwrap: ", errors.Unwrap(deadlineWrapErr) == context.DeadlineExceeded) // true

	// 同じエラー
	fmt.Println("deadlineWrapErr Is deadlineWrapErr:", errors.Is(deadlineWrapErr, deadlineWrapErr)) // true

	// 1つめが子（wrapped)、2つめが親
	fmt.Println("deadlineWrapErr Is context.DeadlineExceeded:", errors.Is(deadlineWrapErr, context.DeadlineExceeded)) // true
	fmt.Println("context.DeadlineExceeded Is deadlinewrapErr:", errors.Is(context.DeadlineExceeded, deadlineWrapErr)) // false

	// 自分で作ったラップエラー
	myWrapErr := NewWrappingError(context.DeadlineExceeded)
	fmt.Print("Unwrap(myWrapErr): ", errors.Unwrap(myWrapErr) == context.DeadlineExceeded) // true
}

func TestErrorAs(t *testing.T) {
	// ラップエラー
	deadlineWrapErr := fmt.Errorf("wrap error: %w", context.DeadlineExceeded)
	// コンテキストを含んだエラー
	spErr := NewSpecialError("bad case")

	// errors.As() は、別の構造体 SomeError でエラーを実装した場合、
	// 第一引数のエラーが、第二引数のエラーの構造体のエラーであるかを見る
	// 第二に構造体の型の代わりにポインタを渡す
	var someErr SpecialError

	// 違うエラーの場合は代入されない
	fmt.Println("deadlineWrapErr As SomeError:", errors.As(deadlineWrapErr, &someErr)) // false
	fmt.Println("someErr:", someErr.Error())                                           // someErr: my error error:

	// specialError が someErr に代入されている
	fmt.Println("anError As SomeErr:", errors.As(spErr, &someErr)) // true
	fmt.Println("someErr:", someErr.Error())                       // someErr: my error error: an

	// wrap されていても errors.As() は代入できる
	someErr = SpecialError{}
	wrappedAnError := fmt.Errorf("wrap error: %w", spErr)
	fmt.Println("wrappedAnError As SomeErr:", errors.As(wrappedAnError, &someErr)) // true
	fmt.Println("someErr:", someErr.Error())                                       // someErr: my error error: an

	// 自分で作ったラップエラー
	myWrapErr := NewWrappingError(spErr)
	someErr = SpecialError{}
	fmt.Println("myWrapErr As SomeErr:", errors.As(myWrapErr, &someErr)) // true
	fmt.Println("someErr:", someErr.Error())                             // someErr: my error error: an
}
