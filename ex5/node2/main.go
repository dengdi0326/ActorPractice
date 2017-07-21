/*
*
* MIT License
*
* Copyright (c) 2017 SmartestEE Inc.
*
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
*
* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*/

/*
* Revision History:
* Initial: 2017/07/20 spaxry
*/

package main

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	//"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/AsynkronIT/goconsole"
	//"golang.org/x/text/message"
	//"time"
	"actor_ex/ex5/messages"


	//"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/remote"
	"fmt"
)
type MyActor struct{
	count int
}

func (state *MyActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.Response:
		state.count++
		fmt.Println(state.count)
	default:
		fmt.Println(msg)
	}
}

func main() {
	remote.Start("localhost:8090")
        props:=actor.FromInstance(&MyActor{})
	pid := actor.Spawn(props)
	message := &messages.Echo{Message: "hej", Sender: pid}

	//this is the remote actor we want to communicate with
	remote := actor.NewPID("localhost:8091", "myactor")
	for i := 0; i < 10; i++ {
		remote.Tell(message)
	}

	console.ReadLine()
}