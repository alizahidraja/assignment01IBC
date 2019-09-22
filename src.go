type Block struct {
	//Hash
	//Data

	transaction string
	prevPointer *Block
	next *Block
	hash [32]byte
	prevHash [32]byte
}



func  DeriveHash(transaction string)[32]byte {
	return sha256.Sum256([]byte(transaction))
}


func InsertBlock(transaction string, chainHead *Block) *Block {
	if chainHead == nil {
		return &Block{transaction, nil,nil,DeriveHash(transaction),[32]byte{}}
	}
	for p := chainHead; p != nil; p = p.next {
		if p.next == nil {
			p.next = &Block{transaction, p,nil,DeriveHash(transaction),DeriveHash(p.transaction)}
			return chainHead
		}
	}
	return chainHead
}

func ListBlocks(chainHead *Block) {
	for p := chainHead; p != nil; p = p.next {
		fmt.Printf("Transaction: %s, Hash:%x, PrevHash:%x\n",p.transaction,p.hash,p.prevHash)

	}
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	for p := chainHead; p != nil; p = p.next {
		if p.transaction == oldTrans{
			p.transaction=newTrans
			p.hash=DeriveHash(newTrans)

		}
	}
}
func VerifyChain(chainHead *Block) {
	for p := chainHead; p != nil; p = p.next {
		if p.next !=nil{
			if p.hash != p.next.prevHash{
				fmt.Println("Chain is Invalid!")
				return
			}
		}
	}
	fmt.Println("Chain is Valid!")
}
