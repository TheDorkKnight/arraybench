package main

import (
	"math/rand"
	"testing"
)

const (
	consistentRandomSeed = 42
	valueLength = 16
)

type sourceData struct {
	sliceOfArrays [][valueLength]byte
}

var (
	compilerThinksThisCouldBeTrue bool = false
)

func makeSourceData(b *testing.B) sourceData {
	// Set up values
	sd := sourceData{sliceOfArrays: make([][valueLength]byte, b.N, b.N)}
	rng := rand.New(rand.NewSource(consistentRandomSeed))
	for i := 0; i < b.N; i++ {
		var data []byte = sd.sliceOfArrays[i][:]
		n, err := rng.Read(data)
		if err != nil {
			b.Fatal(err)
		} else if n != valueLength {
			b.Fatal("failed to read bytes from RNG. Expected to read ", valueLength, " bytes, but got ", n)
		}
	}
	return sd
}

func (sd *sourceData) slice(i int) []byte {
	s := make([]byte, valueLength, valueLength)
	for j := 0; j < valueLength; j++ {
		s[j] = sd.sliceOfArrays[i][j]
	}
	return s
}

func (sd *sourceData) sliceAroundArray(i int) []byte {
	return sd.sliceOfArrays[i][:]
}

func (sd *sourceData) array(i int) [valueLength]byte {
	return sd.sliceOfArrays[i]
}

func (sd *sourceData) arrayCopy(i int) [valueLength]byte {
	var a [valueLength]byte
	for j := 0; j < valueLength; j++ {
		a[j] = sd.sliceOfArrays[i][j]
	}
	return a
}

func BenchmarkSlice(b *testing.B) {
	srcData := makeSourceData(b)

	var cantOptimizeMeOut byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceI := srcData.slice(i)
		for j := 0; j < valueLength; j++ {
			cantOptimizeMeOut ^= sliceI[j]
		}
	}
	b.StopTimer()
	if compilerThinksThisCouldBeTrue {
		b.Log(cantOptimizeMeOut)
	}
	b.StartTimer()
}

func BenchmarkSliceWrappingArray(b *testing.B) {
	srcData := makeSourceData(b)

	var cantOptimizeMeOut byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceI := srcData.sliceAroundArray(i)
		for j := 0; j < valueLength; j++ {
			cantOptimizeMeOut ^= sliceI[j]
		}
	}
	b.StopTimer()
	if compilerThinksThisCouldBeTrue {
		b.Log(cantOptimizeMeOut)
	}
	b.StartTimer()
}

func BenchmarkArray(b *testing.B) {
	srcData := makeSourceData(b)

	var cantOptimizeMeOut byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrayI := srcData.array(i)
		for j := 0; j < valueLength; j++ {
			cantOptimizeMeOut ^= arrayI[j]
		}
	}
	b.StopTimer()
	if compilerThinksThisCouldBeTrue {
		b.Log(cantOptimizeMeOut)
	}
	b.StartTimer()
}

func BenchmarkArrayCopy(b *testing.B) {
	srcData := makeSourceData(b)

	var cantOptimizeMeOut byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrayI := srcData.arrayCopy(i)
		for j := 0; j < valueLength; j++ {
			cantOptimizeMeOut ^= arrayI[j]
		}
	}
	b.StopTimer()
	if compilerThinksThisCouldBeTrue {
		b.Log(cantOptimizeMeOut)
	}
	b.StartTimer()
}
