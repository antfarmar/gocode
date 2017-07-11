package main

import (
	"fmt"
	//"time"
	"math/big"
	//"math/rand"
	"crypto/rand"
)

// RSA-100
// 1522605027922533360535618378132637429718068114961380688657908494580122963258952897654000350692006139
// p = 37975227936943673922808872755445627854565536638199
// q = 40094690950920881030683735292761468389214899724061

// RSA-220
// 2260138526203405784941654048610197513508038915719776718321197768109445641817966676608593121306582577250631562886676970448070001811149711863002112487928199487482066070131066586646083327982803560379205391980139946496955261

func main() {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n, p, m := new(big.Int), new(big.Int), new(big.Int)
	n.SetString("2260138526203405784941654048610197513508038915719776718321197768109445641817966676608593121306582577250631562886676970448070001811149711863002112487928199487482066070131066586646083327982803560379205391980139946496955261", 10)
	s := SqrtBig(n)
	fmt.Println(s, "<--SQRT")

	for {
		//p.Rand(r, s)
		p, _ = rand.Prime(rand.Reader, 364)
		//if !p.ProbablyPrime(1) {continue}
		m.Mod(n, p)
		if m.Cmp(big.NewInt(0)) == 0 {
			break
		}
		//fmt.Println(p)
	}
	fmt.Println("FACTOR: ", p)
}

// SqrtBig returns floor(sqrt(n)). It panics on n < 0.
func SqrtBig(n *big.Int) (x *big.Int) {
	switch n.Sign() {
	case -1:
		panic(-1)
	case 0:
		return big.NewInt(0)
	}

	var px, nx big.Int
	x = big.NewInt(0)
	x.SetBit(x, n.BitLen()/2+1, 1)
	for {
		nx.Rsh(nx.Add(x, nx.Div(n, x)), 1)
		if nx.Cmp(x) == 0 || nx.Cmp(&px) == 0 {
			break
		}
		px.Set(x)
		x.Set(&nx)
	}
	return
}
