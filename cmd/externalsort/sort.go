/*
  author='du'
  date='2020/1/19 20:19'
*/
package main

import (
	"bufio"
	"fmt"
	"golang20200117/pipeline"
	"os"
)

func main() {
	largeTest()
}

func largeTest() {
	//const inFileName="large.in"
	//const size=800000
	//const outFileName="large.out"
	//p:=createPipleline(inFileName,size,4)
	//write2File(p,outFileName)
	//println(outFileName)

	//p := createPipleline("du.in", 512000000, 4)
	//write2File(p, "du.out")
	//printFile("du.out")
	smallTest()
}

func smallTest() {
	//三个方法，第一个创建pipline,第二个写入文件，第三个打印文件。
	p := createPipleline("dusmall.in", 512, 4)
	write2File(p, "dusmall.out")
	printFile("dusmall.out")
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func write2File(p <-chan int, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriteSink(writer, p)

}

func createPipleline(fileName string, fileSize, chunkCount int) <-chan int {
	//假设这里chunkSize正好就是个整数
	chunkSize := fileSize / chunkCount
	sortResults := []<-chan int{}
	pipeline.InitStartTime()
	//先创建
	_, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	for i := 1; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		pipeline.InMemSort(source)
		sortResults = append(sortResults, pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResults...)
}

func createNetworkPipleline(fileName string, fileSize, chunkCount int) <-chan int {
	//假设这里chunkSize正好就是个整数
	chunkSize := fileSize / chunkCount
	sortResults := []<-chan int{}
	pipeline.InitStartTime()
	//先创建
	_, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	for i := 1; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		pipeline.InMemSort(source)
		sortResults = append(sortResults, pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResults...)
}
