package block

type Block struct {
	Index        int64
	Hash         string
	PreviousHash string
	Timestamp    int64
	Data         string
}

func CalculateHash()
