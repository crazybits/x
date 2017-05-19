// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"reflect"

	. "github.com/crazybits/x/blockchain"
	"github.com/crazybits/x/cmd"
)

func main() {
	cmd.Execute()

	receiver := StringToAddress("crazybit")

	withdrawSymbol := "Symbol"

	withdrawAmount := int64(1024)

	depositOp := NewDepositOperation()
	depositOp.Receiver = receiver
	depositOp.Symbol = withdrawSymbol
	depositOp.Amount = withdrawAmount

	data, _ := depositOp.Encode()

	op := NewOperation()
	op.Type = OperationType_Deposit
	op.Payload = data

	tx := NewTransaction()
	tx.AddOperation(op)

	block := NewBlock()

	block.AddTransaction(tx)

	bm := NewBlockManager()

	bm.ProcessBlock(block)

	id, err := block.Digest()
	if err != nil {

	}
	bm.PushBlock(block)

	newBlock := bm.GetBlockByID(id)

	if !reflect.DeepEqual(block, newBlock) {

	}
}
