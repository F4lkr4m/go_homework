package main

import (
	"sort"
	"strconv"
	"sync"
)

const MultiHashTh = 6

func elementaryHash(data string, out chan string, process func(data string) string) {
	result := process(data)
	out <- result
}

func elementaryHashMutex(mu *sync.Mutex, data string, out chan string, process func(data string) string) {
	mu.Lock()
	result := process(data)
	mu.Unlock()
	out <- result
}

func singleHashWorker(outWg *sync.WaitGroup, outMu *sync.Mutex, input interface{}, out chan interface{}) {
	defer outWg.Done()

	crcOneChan := make(chan string)
	go elementaryHash(input.(string), crcOneChan, DataSignerCrc32)

	mdOneChan := make(chan string)
	go elementaryHashMutex(outMu, input.(string), mdOneChan, DataSignerMd5)

	md := <-mdOneChan
	crcTwoChan := make(chan string)
	go elementaryHash(md, crcTwoChan, DataSignerCrc32)

	crcOne := <-crcOneChan
	crcTwo := <-crcTwoChan

	result := crcOne + "~" + crcTwo
	out <- result

}

func SingleHash(in, out chan interface{}) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for input := range in {
		input = strconv.Itoa(input.(int))
		wg.Add(1)
		go singleHashWorker(&wg, &mu, input, out)
	}
	wg.Wait()
}

type orderedData struct {
	data        string
	orderNumber int
}

func elementaryHashWg(data orderedData, wg *sync.WaitGroup, out chan orderedData, process func(data string) string) {
	defer wg.Done()

	result := process(data.data)
	out <- orderedData{data: result, orderNumber: data.orderNumber}
}

func multiHashWorker(outWg *sync.WaitGroup, input interface{}, out chan interface{}) {
	defer outWg.Done()

	wg := &sync.WaitGroup{}
	elems := make(chan orderedData, MultiHashTh)
	for i := 0; i < MultiHashTh; i++ {
		wg.Add(1)
		go elementaryHashWg(orderedData{strconv.Itoa(i) + input.(string), i}, wg, elems, DataSignerCrc32)
	}
	wg.Wait()
	close(elems)

	var crcArray []orderedData
	for i := range elems {
		crcArray = append(crcArray, i)
	}

	// sort data in order
	sort.Slice(crcArray, func(i, j int) bool {
		return crcArray[i].orderNumber < crcArray[j].orderNumber
	})

	// concat data
	var result string
	for i := 0; i < len(crcArray); i++ {
		result += crcArray[i].data
	}
	out <- result
}

func MultiHash(in, out chan interface{}) {
	localWg := sync.WaitGroup{}
	for input := range in {
		localWg.Add(1)
		go multiHashWorker(&localWg, input, out)
	}
	localWg.Wait()
}

func sortAndConcatElementsInSlice(slice []string) string {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	result := slice[0]
	for i := 1; i < len(slice); i++ {
		result += "_" + slice[i]
	}
	return result
}

func CombineResults(in, out chan interface{}) {
	var resultSlice []string
	for i := range in {
		resultSlice = append(resultSlice, i.(string))
	}

	out <- sortAndConcatElementsInSlice(resultSlice)
}

func executeJob(in, out chan interface{}, wg *sync.WaitGroup, job func(in, out chan interface{})) {
	defer wg.Done()
	defer close(out)

	job(in, out)
}

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{})
	for _, currJob := range jobs {
		out := make(chan interface{})

		wg.Add(1)
		go executeJob(in, out, wg, currJob)
		in = out
	}
	wg.Wait()
}
