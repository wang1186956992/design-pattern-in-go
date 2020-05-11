package main

import "fmt"

type IStream interface {
	Read(by []byte)
	Seek(n int)
	Write(by []byte)
}

type FileStream struct{}

func (fs *FileStream) Read(by []byte) {
	fmt.Printf("File Stream read: %v\n", string(by))
}

func (fs *FileStream) Seek(n int) {
	fmt.Printf("File Stream seek: %v\n", n)
}

func (fs *FileStream) Write(by []byte) {
	fmt.Printf("File Stream write: %v\n", string(by))
}

type NetworkStream struct{}

func (ns *NetworkStream) Read(by []byte) {
	fmt.Printf("Network Stream read: %v\n", string(by))
}

func (ns *NetworkStream) Seek(n int) {
	fmt.Printf("Network Stream seek: %v\n", n)
}

func (ns *NetworkStream) Write(by []byte) {
	fmt.Printf("Network Stream write: %v\n", by)
}

type MemoryStream struct{}

func (ms *MemoryStream) Read(by []byte) {
	fmt.Printf("Memory Stream read: %v\n", string(by))
}

func (ms *MemoryStream) Seek(n int) {
	fmt.Printf("Memory Stream seek: %v\n", n)
}

func (ms *MemoryStream) Write(by []byte) {
	fmt.Printf("Memory Stream write: %v\n", by)
}

type DecoratorStream struct {
	Stream IStream
}

func (ds *DecoratorStream) Read(by []byte) {
	ds.Stream.Read(by)
}

func (ds *DecoratorStream) Seek(n int) {
	ds.Stream.Seek(n)
}

func (ds *DecoratorStream) Write(by []byte) {
	ds.Stream.Write(by)
}

type CrptoStream struct {
	DecoratorStream
}

func (cs *CrptoStream) Read(by []byte) {
	fmt.Printf("Crpto Stream Read: %v\n", string(by))
	cs.DecoratorStream.Read(by)
}

func (cs *CrptoStream) Seek(n int) {
	fmt.Printf("Crpto Stream Seek: %v\n", n)
	cs.DecoratorStream.Seek(n)
}

func (cs *CrptoStream) Write(by []byte) {
	fmt.Printf("Crpto Stream Write: %v\n", string(by))
	cs.DecoratorStream.Write(by)
}

type BufferedStream struct {
	DecoratorStream
}

func (bs *BufferedStream) Read(by []byte) {
	fmt.Printf("Buffered Stream Read: %v\n", string(by))
	bs.DecoratorStream.Read(by)
}

func (bs *BufferedStream) Seek(n int) {
	fmt.Printf("Buffered Stream Seek: %v\n", n)
	bs.DecoratorStream.Seek(n)
}

func (bs *BufferedStream) Write(by []byte) {
	fmt.Printf("Buffered Stream Write: %v\n", string(by))
	bs.DecoratorStream.Write(by)
}
