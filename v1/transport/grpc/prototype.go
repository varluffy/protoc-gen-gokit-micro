// The MIT License (MIT)
//
// Copyright (c) 2022 leewckk@126.com
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
//

package grpc

import (
	"github.com/leewckk/protoc-gen-gokit-micro/common"
	"google.golang.org/protobuf/compiler/protogen"
)

type Prototype struct {
	generators []Generator
}

func NewPrototype() *Prototype {
	p := &Prototype{
		generators: make([]Generator, 0, 0),
	}
	return p
}

func PrototypeMethodFieldName(methodName string) string {
	return "__" + methodName + "__"
}

func GetTransportName(serviceName string) string {
	return serviceName + "ServiceTransport"
}

func (this *Prototype) GenerateFile(fileUri string, gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) (*protogen.GeneratedFile, error) {
	serviceName := GetTransportName(svc.GoName)
	gfile.P()
	gfile.P("type ", serviceName, " struct {")
	gfile.P(common.GetPBPackagePath(file, options).Ident("Unimplemented" + svc.GoName + "Server"))
	for _, method := range svc.Methods {
		gfile.P(PrototypeMethodFieldName(method.GoName), " ", common.GoKitGRPC.Ident("Handler"))
	}
	gfile.P("}")
	gfile.P()
	return gfile, nil
}
