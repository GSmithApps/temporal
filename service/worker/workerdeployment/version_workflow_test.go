// The MIT License
//
// Copyright (c) 2024 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2024 Uber Technologies, Inc.
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

package workerdeployment

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
	"go.uber.org/mock/gomock"
)

type deploymentSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	controller *gomock.Controller
	env        *testsuite.TestWorkflowEnvironment
}

func TestDeploymentSuite(t *testing.T) {
	suite.Run(t, new(deploymentSuite))
}

func (s *deploymentSuite) SetupTest() {
	s.controller = gomock.NewController(s.T())
	s.env = s.WorkflowTestSuite.NewTestWorkflowEnvironment()
	s.env.RegisterWorkflow(VersionWorkflow)
}

func (s *deploymentSuite) TearDownTest() {
	s.controller.Finish()
	s.env.AssertExpectations(s.T())
}

// TestRegisterTaskQueueInDeployment tests the case when a task-queue
// is registered in a deployment
// func (s *deploymentSuite) TestRegisterTaskQueueInDeployment() {
// )
// }

// TestRegisterTaskQueuesInDeployment tests the case when multiple task-queues
// are registered (non-concurrently) in a deployment
// func (s *deploymentSuite) TestRegisterTaskQueuesInDeployment() {
// )
// }

// TestRegisterTaskQueuesInDeploymentConcurrent tests the case when multiple task-queues
// are registered concurrently in a deployment
// func (s *deploymentSuite) TestRegisterTaskQueuesInDeploymentConcurrent() {
// )
// }

// TestRegisterTaskQueuesExceedLimit tests the case when the number of registered task-queues
// exceed the allowed per-deployment limit
// func (s *deploymentSuite) TestRegisterTaskQueuesExceedLimit() {
// )
// }

// TestStartDeploymentWorkflowExceedLimit tests the case when the number of
// deployment workflow executions exceed the allowed namespace limit
// func (s *deploymentSuite) TestRegisterTaskQueuesExceedLimit() {
// )
// }
