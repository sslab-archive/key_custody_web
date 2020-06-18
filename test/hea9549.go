/*
 * Copyright 2019 hea9549
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/share"
	"go.dedis.ch/kyber/v3/share/pvss"
	"sync"
	"time"
)

func main()  {
	mu := sync.Mutex{}
	suite := edwards25519.NewBlakeSHA256Ed25519()
	G := suite.Point().Base()
	H := suite.Point().Pick(suite.XOF([]byte("H")))
	n := 512
	t := 2*n/3 + 1
	x := make([]kyber.Scalar, n) // trustee private keys
	X := make([]kyber.Point, n)  // trustee public keys
	secrets := make([]kyber.Scalar, n)
	for i := 0; i < n; i++ {
		x[i] = suite.Scalar().Pick(suite.RandomStream())
		X[i] = suite.Point().Mul(x[i], nil)
		secrets[i] = suite.Scalar().Pick(suite.RandomStream()) // Scalar of shared secret
	}

	encShareData := make([][]*pvss.PubVerShare, n)
	publicPolyData := make([]*share.PubPoly, n)
	decData := make([][]*pvss.PubVerShare, n)
	wgMake := new(sync.WaitGroup)
	makeShares := func(index int, secret kyber.Scalar, publicKeys []kyber.Point) (shares []*pvss.PubVerShare, commit *share.PubPoly) {
		encShares, pubPoly, err := pvss.EncShares(suite, H, X, secret, t)
		if err != nil {
			fmt.Println(err.Error())
		}
		mu.Lock()
		encShareData[index] = encShares
		publicPolyData[index] = pubPoly
		mu.Unlock()
		wgMake.Done()
		return encShares, pubPoly
	}
	wgMake.Add(n)

	fmt.Println("encrypt & make share start : " + time.Now().String())
	for i := 0; i < n; i++ {
		go makeShares(i, secrets[i], X)
	}
	wgMake.Wait()

	fmt.Println("encrypt & make share end : " + time.Now().String())

	fmt.Println("decrypt share start : " + time.Now().String())

	wgDecrypt := new(sync.WaitGroup)
	wgDecrypt.Add(n)
	decryptMyShares := func(myIdx int, pubPoly []*share.PubPoly, encShares []*pvss.PubVerShare, privateKey kyber.Scalar, publicKey kyber.Point) {
		var dsList []*pvss.PubVerShare
		for i := 0; i < len(pubPoly); i++ {
			sH := pubPoly[i].Eval(encShares[i].S.I).V
			ds, err := pvss.DecShare(suite, H, publicKey, sH, privateKey, encShares[i])
			if err != nil {
				fmt.Println(err.Error())
			}
			dsList = append(dsList, ds)
		}
		mu.Lock()
		decData[myIdx] = dsList
		mu.Unlock()
		wgDecrypt.Done()
	}
	for i := 0; i < n; i++ {
		var myShareData []*pvss.PubVerShare
		for j := 0; j < n; j++ {
			myShareData = append(myShareData, encShareData[j][i])
		}
		go decryptMyShares(i, publicPolyData, myShareData, x[i], X[i])

	}
	wgDecrypt.Wait()

	fmt.Println("decrypt share end : " + time.Now().String())

	fmt.Println("verify share start : " + time.Now().String())

	for i := 0; i < n; i++ {
		var D []*pvss.PubVerShare // good decrypted shares
		for j := 0; j < n; j++ {
			D = append(D, decData[j][i])
		}
		recovered, err := pvss.RecoverSecret(suite, G, X, encShareData[i], D, t, n)
		if err != nil {
			fmt.Println("verify Err")
			fmt.Println(err.Error())
		}
		if suite.Point().Mul(secrets[i], nil).Equal(recovered) == false{
			fmt.Println("result : ", suite.Point().Mul(secrets[i], nil).Equal(recovered))
		}

	}

	fmt.Println(time.Now().String())
}


