package types

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"testing"
)

func TestMarshalHashBilling(t *testing.T) {
	v := Billing{}
	binary.Read(rand.Reader, binary.BigEndian, &v)
	bts1, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	bts2, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bts1, bts2) {
		t.Fatal("hash not stable")
	}
}

func BenchmarkMarshalHashBilling(b *testing.B) {
	v := Billing{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalHash()
	}
}

func BenchmarkAppendMsgBilling(b *testing.B) {
	v := Billing{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalHash()
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalHash()
	}
}

func TestMarshalHashBillingHeader(t *testing.T) {
	v := BillingHeader{}
	binary.Read(rand.Reader, binary.BigEndian, &v)
	bts1, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	bts2, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bts1, bts2) {
		t.Fatal("hash not stable")
	}
}

func BenchmarkMarshalHashBillingHeader(b *testing.B) {
	v := BillingHeader{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalHash()
	}
}

func BenchmarkAppendMsgBillingHeader(b *testing.B) {
	v := BillingHeader{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalHash()
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalHash()
	}
}
