// Copyright (c) 2015 Matt Weagle <mweagle@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package sparta

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"net/http"
)

func lambdaHelloWorld(event *LambdaEvent, context *LambdaContext, w *http.ResponseWriter, logger *logrus.Logger) {
	fmt.Fprintf(*w, "Hello World!")
}

func ExampleNewLambda() {
	// IAM role is implicitly created to support the lambda execution.
	//
	roleDefinition := sparta.IAMRoleDefinition{}
	roleDefinition.Privileges = append(roleDefinition.Privileges, sparta.IAMRolePrivilege{
		Actions: []string{"s3:GetObject",
			"s3:PutObject"},
		Resource: "arn:aws:s3:::*",
	})
	helloWorldLambda := NewLambda(sparta.IAMRoleDefinition{}, lambdaHelloWorld, nil)
	if nil != helloWorldLambda {
		fmt.Printf("Failed to create new Lambda function")
	}
}