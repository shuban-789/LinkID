package mechanisms

func getBlockByHash(TargetChain chain, hash string) (block, bool) {
	for _, b := range TargetChain.Chain {
		if b.CurrentHash == hash {
			return b, true
		}
	}
	return block{}, false
}