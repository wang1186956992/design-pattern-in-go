package chain

import (
	"os"
	"testing"
)

func TestChain(t *testing.T) {
	f, err := os.OpenFile("D:/log/log.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatalf("打开文件失败, %v\n", err)
	}

	chain := LogChain{}
	fLog := Log{
		Name: "文件",
		W:    f,
	}

	tLog := Log{
		Name: "终端",
		W:    os.Stdout,
	}

	chain.AddLog(&tLog)
	chain.AddLog(&fLog)
	chain.Handle()

}
