package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func wrapFuncToMeasure(f func(execAtEveryIteration func(fa func()) )) {
	latencyArray := []int{}
	sum := 0
	noOfMeasurements := 0

	f(func(fa func()) {
		noOfMeasurements++
		startTime := time.Now()

		fa()

		latency := int(time.Now().Sub(startTime).Milliseconds())
		//fmt.Println(latency)
		latencyArray = append(latencyArray, latency)
		sum += latency
	})

	fmt.Println("============RESULTS=============")
	fmt.Printf("total time: %d ms\n", sum)
	fmt.Printf("avg: %d ms\n", sum/noOfMeasurements)
	sort.Ints(latencyArray)

	p95Index := int(0.95 * float32(noOfMeasurements))
	fmt.Printf("index for p99: %d\n", p95Index)
	fmt.Printf("p95: %d ms\n", latencyArray[p95Index])

	p99Index := int(0.99 * float32(noOfMeasurements))
	fmt.Printf("index for p99: %d\n", p99Index)
	fmt.Printf("p95: %d ms\n", latencyArray[p99Index])
	fmt.Println("===============================")
}

func RunBenchmark() {
	a := setup()

	runIndex := 100000

	fmt.Println("---------Adding users to group---------------")
	wrapFuncToMeasure(func(execAtEveryIteration func(fa func())) {
		for i := 0; i < runIndex; i++ {
			execAtEveryIteration(func() {
				strs := fmt.Sprintf("group:g%d#member@user:u%d", i/10, i)
				a.Permission.Add(strs)
			})
		}
	})

	fmt.Println("---------Adding group to firehose---------------")
	wrapFuncToMeasure(func(execAtEveryIteration func(fa func())) {
		for i := 0; i < runIndex; i++ {
			execAtEveryIteration(func() {
				stmt := fmt.Sprintf("resource/firehose:f%d#manager@group:g%d", i, i/10)
				a.Permission.Add(stmt)
			})
		}
	})

	a.Permission.Add("group:f_admins#member@user:u2")
	a.Permission.Add("project:p1#firehose_admins@group:f_admins#member")
	a.Permission.Add("resource/firehose:f1#manager@project:p1#firehose_admins")

	time.Sleep(5 * time.Second)

	fmt.Println("---------Checking Group Permission---------------")
	wrapFuncToMeasure(func(execAtEveryIteration func(fa func())) {
		for i := 0; i < runIndex; i++ {
			execAtEveryIteration(func() {
				a.Permission.Check(fmt.Sprintf("group:g%d#view@user:u%d", i/10, i))
			})
		}
	})

	fmt.Println("---------Checking Firehose Permission---------------")
	wrapFuncToMeasure(func(execAtEveryIteration func(fa func())) {
		for i := 0; i < runIndex; i++ {
			execAtEveryIteration(func() {
				a.Permission.Check(fmt.Sprintf("resource/firehose:f%d#manage@user:u%d#...", i, i))
			})
		}
	})

	fmt.Println("---------Checking incorrect Firehose Permission---------------")
	wrapFuncToMeasure(func(execAtEveryIteration func(fa func())) {
		for i := 0; i < runIndex; i++ {
			min := 1
			max := 10
			randomUser := rand.Intn(max - min) + min
			execAtEveryIteration(func() {
				a.Permission.Check(fmt.Sprintf("resource/firehose:f%d#manage@user:u%d#...", i, randomUser))
			})
		}
	})
}
