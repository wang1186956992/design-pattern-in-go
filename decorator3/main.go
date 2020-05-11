package main

func main() {
	//构造一个缓存加密文件流的方法
	var stream IStream = &BufferedStream{
		DecoratorStream: DecoratorStream{
			Stream: &CrptoStream{
				DecoratorStream: DecoratorStream{
					Stream: &FileStream{},
				},
			},
		},
	}

	strRead := "Read Str"
	stream.Read([]byte(strRead))
	stream.Seek(2)
	strWrite := "Write Str"
	stream.Write([]byte(strWrite))
}
