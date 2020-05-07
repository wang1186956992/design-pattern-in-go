package chain

import (
	"os"
	"testing"
)

func TestChain(t *testing.T) {
	f, err := os.OpenFile("D:/log/log.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatalf("打开文件失败%s\n", err)
	}
	fLog := Log{
		Name: "File",
		W:    f,
	}

	tLog := Log{
		Name: "终端",
		W:    os.Stdout,
	}

	ch2 := LogChain{
		L: &tLog,
	}
	ch1 := LogChain{
		Next: &ch2,
		L:    &fLog,
	}

	ch1.Handle()

}
