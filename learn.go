package main

type Learner struct {
	PieceLifetimes [12][64]int // Tracks how many turns a piece type is on a certain square
}

func (l *Learner) UpdatePieceLifetimes(b boardStruct) {
	for sq, pc := range b.sq {
		if pc != empty {
			l.PieceLifetimes[pc][sq] += 1
		}
	}
}
