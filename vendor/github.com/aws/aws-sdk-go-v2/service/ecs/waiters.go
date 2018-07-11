// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilServicesInactive uses the Amazon ECS API operation
// DescribeServices to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ECS) WaitUntilServicesInactive(input *DescribeServicesInput) error {
	return c.WaitUntilServicesInactiveWithContext(aws.BackgroundContext(), input)
}

// WaitUntilServicesInactiveWithContext is an extended version of WaitUntilServicesInactive.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) WaitUntilServicesInactiveWithContext(ctx aws.Context, input *DescribeServicesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilServicesInactive",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "failures[].reason",
				Expected: "MISSING",
			},
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "services[].status",
				Expected: "INACTIVE",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeServicesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeServicesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilServicesStable uses the Amazon ECS API operation
// DescribeServices to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ECS) WaitUntilServicesStable(input *DescribeServicesInput) error {
	return c.WaitUntilServicesStableWithContext(aws.BackgroundContext(), input)
}

// WaitUntilServicesStableWithContext is an extended version of WaitUntilServicesStable.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) WaitUntilServicesStableWithContext(ctx aws.Context, input *DescribeServicesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilServicesStable",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "failures[].reason",
				Expected: "MISSING",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "services[].status",
				Expected: "DRAINING",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "services[].status",
				Expected: "INACTIVE",
			},
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(services[?!(length(deployments) == `1` && runningCount == desiredCount)]) == `0`",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeServicesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeServicesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilTasksRunning uses the Amazon ECS API operation
// DescribeTasks to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ECS) WaitUntilTasksRunning(input *DescribeTasksInput) error {
	return c.WaitUntilTasksRunningWithContext(aws.BackgroundContext(), input)
}

// WaitUntilTasksRunningWithContext is an extended version of WaitUntilTasksRunning.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) WaitUntilTasksRunningWithContext(ctx aws.Context, input *DescribeTasksInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilTasksRunning",
		MaxAttempts: 100,
		Delay:       aws.ConstantWaiterDelay(6 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "tasks[].lastStatus",
				Expected: "STOPPED",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "failures[].reason",
				Expected: "MISSING",
			},
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "tasks[].lastStatus",
				Expected: "RUNNING",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeTasksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeTasksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilTasksStopped uses the Amazon ECS API operation
// DescribeTasks to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ECS) WaitUntilTasksStopped(input *DescribeTasksInput) error {
	return c.WaitUntilTasksStoppedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilTasksStoppedWithContext is an extended version of WaitUntilTasksStopped.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) WaitUntilTasksStoppedWithContext(ctx aws.Context, input *DescribeTasksInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilTasksStopped",
		MaxAttempts: 100,
		Delay:       aws.ConstantWaiterDelay(6 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "tasks[].lastStatus",
				Expected: "STOPPED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeTasksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeTasksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
