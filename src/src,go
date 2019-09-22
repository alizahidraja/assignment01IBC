type Block struct {
	//Hash
	//Data
	transaction string
	prevPointer *Block
	next *Block
}



//func (b Block) DeriveHash() {
//	hash := sha256.Sum256([]byte(b.transaction))
//	b.Hash=hash[:]
//}


func InsertBlock(transaction string, chainHead *Block) *Block {
	if chainHead == nil {
		return &Block{transaction, nil,nil}
	}
	for p := chainHead; p != nil; p = p.next {
		if p.next == nil {
			p.next = &Block{transaction, p,nil}
			return chainHead
		}
	}
	return chainHead
}

func ListBlocks(chainHead *Block) {
	for p := chainHead; p != nil; p = p.next {
		fmt.Println(p.transaction)
	}
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	for p := chainHead; p != nil; p = p.next {
		if p.transaction == oldTrans{
			p.transaction=newTrans
		}
		if p.next == nil {
			if p.transaction == oldTrans{
				p.transaction=newTrans
			}
			return
		}
	}
}
func VerifyChain(chainHead *Block) {
	for p := chainHead; p != nil; p = p.next {
	}
	fmt.Println("Chain is Verified!")
}
