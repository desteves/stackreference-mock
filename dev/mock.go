package dev

import "github.com/pulumi/pulumi/sdk/go/pulumi"

// MockStackReference is a mock implementation of StackReference for testing purposes.
type MockStackReference struct {
	// Include fields and methods needed for testing
}

// Implement the methods of StackReference interface for MockStackReference

// For example:
func (m *MockStackReference) Name() pulumi.StringOutput {
	// Mock implementation
	return pulumi.String("mocked-stack-name").ToStringOutput()
}

func (m *MockStackReference) GetOutput(name pulumi.StringInput) pulumi.AnyOutput {
	// Mock implementation
	return pulumi.Any("mocked-stack-output")
}
