// Copyright 2014 loolgame Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package basemodule

import (
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/rpc"
	"github.com/liangdas/mqant/registry"
	"github.com/liangdas/mqant/rpc/base"
)

func NewServerSession(app module.App, name string,node *registry.Node) (module.ServerSession,error){
	rpc,err:=defaultrpc.NewRPCClient(app,node)
	if err!=nil{
		return nil,err
	}
	session := &serverSession{
		Rpc:   rpc,
		name:name,
		node:node,
	}
	return session,err
}

type serverSession struct {
	node 	*registry.Node
	name 	string
	Rpc   	mqrpc.RPCClient
}

func (c *serverSession) GetId() string {
	return c.node.Id
}
func (c *serverSession) GetName() string {
	return c.name
}
func (c *serverSession) GetRpc() mqrpc.RPCClient {
	return c.Rpc
}

/**
消息请求 需要回复
*/
func (c *serverSession) Call(_func string, params ...interface{}) (interface{}, string) {
	return c.Rpc.Call(_func, params...)
}

/**
消息请求 不需要回复
*/
func (c *serverSession) CallNR(_func string, params ...interface{}) (err error) {
	return c.Rpc.CallNR(_func, params...)
}

/**
消息请求 需要回复
*/
func (c *serverSession) CallArgs(_func string, ArgsType []string, args [][]byte) (interface{}, string) {
	return c.Rpc.CallArgs(_func, ArgsType, args)
}

/**
消息请求 不需要回复
*/
func (c *serverSession) CallNRArgs(_func string, ArgsType []string, args [][]byte) (err error) {
	return c.Rpc.CallNRArgs(_func, ArgsType, args)
}
