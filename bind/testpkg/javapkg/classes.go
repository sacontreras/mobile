// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package javapkg

import (
	"Java/java/beans"
	"Java/java/io"
	"Java/java/io/IOException"
	"Java/java/lang"
	"Java/java/lang/Character"
	"Java/java/lang/Integer"
	"Java/java/lang/Object"
	"Java/java/util"
	"Java/java/util/concurrent"
)

const (
	ToStringPrefix     = "Go toString: "
	IOExceptionMessage = "GoInputStream IOException"
)

type GoRunnable struct {
	lang.Runnable
	this lang.Runnable
}

func (r *GoRunnable) ToString(_ lang.Runnable) string {
	return ToStringPrefix
}

func (r *GoRunnable) Run(this lang.Runnable) {
	r.this = this // Careful: This creates a reference cycle
}

func (r *GoRunnable) GetThis() lang.Runnable {
	return r.this
}

type GoInputStream struct {
	io.InputStream
}

func (_ *GoInputStream) Read() (int32, error) {
	return 0, IOException.New_Ljava_lang_String_2(IOExceptionMessage)
}

func NewGoInputStream() *GoInputStream {
	return new(GoInputStream)
}

type GoFuture struct {
	concurrent.Future
}

func (_ *GoFuture) Cancel(_ bool) bool {
	return false
}

func (_ *GoFuture) Get() lang.Object {
	return nil
}

func (_ *GoFuture) Get2(_ int64, _ concurrent.TimeUnit) lang.Object {
	return nil
}

func (_ *GoFuture) IsCancelled() bool {
	return false
}

func (_ *GoFuture) IsDone() bool {
	return false
}

type GoObject struct {
	lang.Object
}

func (_ *GoObject) ToString(this lang.Object) string {
	return ToStringPrefix + this.Super().ToString()
}

func (_ *GoObject) HashCode() int32 {
	return 42
}

func RunRunnable(r lang.Runnable) {
	r.Run()
}

func RunnableRoundtrip(r lang.Runnable) lang.Runnable {
	return r
}

// Test constructing and returning Go instances of GoObject and GoRunnable
// outside a constructor
func ConstructGoRunnable() *GoRunnable {
	return new(GoRunnable)
}

func ConstructGoObject() *GoObject {
	return new(GoObject)
}

// java.beans.PropertyChangeEvent is a class a with no default constructors.
type GoPCE struct {
	beans.PropertyChangeEvent
}

func NewGoPCE(_ lang.Object, _ string, _ lang.Object, _ lang.Object) *GoPCE {
	return new(GoPCE)
}

// java.util.ArrayList is a class with multiple constructors
type GoArrayList struct {
	util.ArrayList
}

func NewGoArrayList() *GoArrayList {
	return new(GoArrayList)
}

func NewGoArrayListWithCap(_ int32) *GoArrayList {
	return new(GoArrayList)
}

func CallSubset(s Character.Subset) {
	s.ToString()
}

type GoSubset struct {
	Character.Subset
}

func NewGoSubset(_ string) *GoSubset {
	return new(GoSubset)
}

func NewJavaObject() lang.Object {
	return Object.New()
}

func NewJavaInteger() lang.Integer {
	return Integer.New_I(42)
}

type NoargConstructor struct {
	util.BitSet // An otherwise unused class with a no-arg constructor
}
