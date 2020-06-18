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
			"go.dedis.ch/kyber/v3/share"
	"go.dedis.ch/kyber/v3/group/edwards25519"
			)

func main() {

	GeneratePartialKey("15e82363cd90d31dd16fb677927a335326dbe99472244668bc31451b192c7d90", 5)
}

func GeneratePartialKey(privateKey string, providerNum int) []*share.PriShare {

	// 암호 알고리즘인듯
	fmt.Println("시작이다 Partial Key 생성~~ ", privateKey)
	suite := edwards25519.NewBlakeSHA256Ed25519()
	threshold := providerNum - 1
	// 다항식 세우기 랜덤으로...
	secret := suite.Scalar().SetBytes([]byte(privateKey))

	priPoly := share.NewPriPoly(suite, threshold, secret, suite.RandomStream())

	// Create secret set of shares
	priShares := priPoly.Shares(providerNum)

	// Partial keys
	return priShares
}

func Test(){

	//G := suite.Point().Base()
	//H := suite.Point().Pick(suite.XOF([]byte("H")))
	// 이놈 역할은 뭐지?
	//x := make([]kyber.Scalar, providerNum)
	// 다항식 세우기 랜덤으로...
	//X := make([]kyber.Point, providerNum) // trustee public keys
	// Create public polynomial commitments with respect to basis H
	//pubPoly := priPoly.Commit(H)
	//
	//// Encrypt Share
	//encShares, pubPoly, err := pvss.EncShares(suite, H, X, secret, threshold)
	//if err != nil {
	//	log.Println(err)
	//}
	//// share 된 결과
	//fmt.Println(encShares)
	//fmt.Println(pubPoly)
	//
	//// Decrypt 해보는 시도
	//decData := make([]*pvss.PubVerShare, providerNum)
	//// 각 비밀을 decrypt 하는 함수
	//decryptMyShares := func(myIdx int, pubPoly *share.PubPoly, encShares []*pvss.PubVerShare, privateKey kyber.Scalar, publicKey kyber.Point) {
	//	for i := 0; i < len(encShares); i++ {
	//		sH := pubPoly.Eval(encShares[i].S.I).V
	//		ds, err := pvss.DecShare(suite, H, publicKey, sH, privateKey, encShares[i])
	//		if err != nil {
	//			fmt.Println(err.Error())
	//		}
	//		decData[myIdx] = ds
	//	}
	//}
	//
	//// 참가자들 마다 decrypt 한 후 결과 값을 decData에 추가.
	//for i := 0; i < n; i++ {
	//	decryptMyShares(i, pubPoly, encShares, x[i], X[i])
	//}
	//
	//// 데이터들을 가지고 직접 Recover를 해보자.
	//var D []*pvss.PubVerShare // good decrypted shares
	//for i := 0; i < n; i++ {
	//	D = append(D, decData[i])
	//}
	//recovered, err := pvss.RecoverSecret(suite, G, X, encShares, D, t, n)
	//if err != nil {
	//	fmt.Println("verify Err")
	//	fmt.Println(err.Error())
	//}
	//if suite.Point().Mul(secret, nil).Equal(recovered) == false {
	//	fmt.Println("result : ", suite.Point().Mul(secret, nil).Equal(recovered))
	//}
	//fmt.Println("다시 복원했다 슈벌 ", secret.String())
}
