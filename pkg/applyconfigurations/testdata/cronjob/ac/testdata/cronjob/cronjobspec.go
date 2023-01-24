// Code generated by __debug_bin. DO NOT EDIT.

package cronjob

import (
	v1beta1 "k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"
	cronjob "sigs.k8s.io/controller-tools/pkg/applyconfigurations/testdata/cronjob"
)

// CronJobSpecApplyConfiguration represents an declarative configuration of the CronJobSpec type for use
// with apply.
type CronJobSpecApplyConfiguration struct {
	Schedule                   *string                             `json:"schedule,omitempty"`
	StartingDeadlineSeconds    *int64                              `json:"startingDeadlineSeconds,omitempty"`
	ConcurrencyPolicy          *cronjob.ConcurrencyPolicy          `json:"concurrencyPolicy,omitempty"`
	Suspend                    *bool                               `json:"suspend,omitempty"`
	BinaryName                 []byte                              `json:"binaryName,omitempty"`
	CanBeNull                  *string                             `json:"canBeNull,omitempty"`
	JobTemplate                *v1beta1.JobTemplateSpec            `json:"jobTemplate,omitempty"`
	SuccessfulJobsHistoryLimit *int32                              `json:"successfulJobsHistoryLimit,omitempty"`
	FailedJobsHistoryLimit     *int32                              `json:"failedJobsHistoryLimit,omitempty"`
	StringSliceData            map[string][]string                 `json:"stringSliceData,omitempty"`
	PtrData                    map[string]*string                  `json:"ptrData,omitempty"`
	Slice                      []string                            `json:"slice,omitempty"`
	SlicePtr                   []*string                           `json:"slicePtr,omitempty"`
	SliceStruct                []*cronjob.ExampleStruct            `json:"sliceStruct,omitempty"`
	BuiltInReference           *v1.PodSpec                         `json:"builtInReference,omitempty"`
	Int                        *int                                `json:"int,omitempty"`
	AssociativeList            []AssociativeTypeApplyConfiguration `json:"associativeList,omitempty"`
}

// CronJobSpecApplyConfiguration constructs an declarative configuration of the CronJobSpec type for use with
// apply.
func CronJobSpec() *CronJobSpecApplyConfiguration {
	return &CronJobSpecApplyConfiguration{}
}

// WithSchedule sets the Schedule field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Schedule field is set to the value of the last call.
func (b *CronJobSpecApplyConfiguration) WithSchedule(value string) *CronJobSpecApplyConfiguration {
	b.Schedule = &value
	return b
}

// WithStartingDeadlineSeconds sets the StartingDeadlineSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartingDeadlineSeconds field is set to the value of the last call.
func (b *CronJobSpecApplyConfiguration) WithStartingDeadlineSeconds(value int64) *CronJobSpecApplyConfiguration {
	b.StartingDeadlineSeconds = &value
	return b
}

// WithConcurrencyPolicy sets the ConcurrencyPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConcurrencyPolicy field is set to the value of the last call.
func (b *CronJobSpecApplyConfiguration) WithConcurrencyPolicy(value cronjob.ConcurrencyPolicy) *CronJobSpecApplyConfiguration {
	b.ConcurrencyPolicy = &value
	return b
}

// WithSuspend sets the Suspend field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Suspend field is set to the value of the last call.
func (b *CronJobSpecApplyConfiguration) WithSuspend(value bool) *CronJobSpecApplyConfiguration {
	b.Suspend = &value
	return b
}

// WithBinaryName adds the given value to the BinaryName field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the BinaryName field.
func (b *CronJobSpecApplyConfiguration) WithBinaryName(values ...byte) *CronJobSpecApplyConfiguration {
	for i := range values {
		b.BinaryName = append(b.BinaryName, values[i])
	}
	return b
}

// WithCanBeNull sets the CanBeNull field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CanBeNull field is set to the value of the last call.
func (b *CronJobSpecApplyConfiguration) WithCanBeNull(value string) *CronJobSpecApplyConfiguration {
	b.CanBeNull = &value
	return b
}

// WithJobTemplate sets the JobTemplate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the JobTemplate field is set to the value of the last call.
func (b *CronJobSpecApplyConfiguration) WithJobTemplate(value v1beta1.JobTemplateSpec) *CronJobSpecApplyConfiguration {
	b.JobTemplate = &value
	return b
}

// WithSuccessfulJobsHistoryLimit sets the SuccessfulJobsHistoryLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SuccessfulJobsHistoryLimit field is set to the value of the last call.
func (b *CronJobSpecApplyConfiguration) WithSuccessfulJobsHistoryLimit(value int32) *CronJobSpecApplyConfiguration {
	b.SuccessfulJobsHistoryLimit = &value
	return b
}

// WithFailedJobsHistoryLimit sets the FailedJobsHistoryLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailedJobsHistoryLimit field is set to the value of the last call.
func (b *CronJobSpecApplyConfiguration) WithFailedJobsHistoryLimit(value int32) *CronJobSpecApplyConfiguration {
	b.FailedJobsHistoryLimit = &value
	return b
}

// WithStringSliceData puts the entries into the StringSliceData field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the StringSliceData field,
// overwriting an existing map entries in StringSliceData field with the same key.
func (b *CronJobSpecApplyConfiguration) WithStringSliceData(entries map[string][]string) *CronJobSpecApplyConfiguration {
	if b.StringSliceData == nil && len(entries) > 0 {
		b.StringSliceData = make(map[string][]string, len(entries))
	}
	for k, v := range entries {
		b.StringSliceData[k] = v
	}
	return b
}

// WithPtrData puts the entries into the PtrData field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the PtrData field,
// overwriting an existing map entries in PtrData field with the same key.
func (b *CronJobSpecApplyConfiguration) WithPtrData(entries map[string]*string) *CronJobSpecApplyConfiguration {
	if b.PtrData == nil && len(entries) > 0 {
		b.PtrData = make(map[string]*string, len(entries))
	}
	for k, v := range entries {
		b.PtrData[k] = v
	}
	return b
}

// WithSlice adds the given value to the Slice field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Slice field.
func (b *CronJobSpecApplyConfiguration) WithSlice(values ...string) *CronJobSpecApplyConfiguration {
	for i := range values {
		b.Slice = append(b.Slice, values[i])
	}
	return b
}

// WithSlicePtr adds the given value to the SlicePtr field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SlicePtr field.
func (b *CronJobSpecApplyConfiguration) WithSlicePtr(values ...*string) *CronJobSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSlicePtr")
		}
		b.SlicePtr = append(b.SlicePtr, *values[i])
	}
	return b
}

// WithSliceStruct adds the given value to the SliceStruct field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SliceStruct field.
func (b *CronJobSpecApplyConfiguration) WithSliceStruct(values ...**cronjob.ExampleStruct) *CronJobSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSliceStruct")
		}
		b.SliceStruct = append(b.SliceStruct, *values[i])
	}
	return b
}

// WithBuiltInReference sets the BuiltInReference field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BuiltInReference field is set to the value of the last call.
func (b *CronJobSpecApplyConfiguration) WithBuiltInReference(value v1.PodSpec) *CronJobSpecApplyConfiguration {
	b.BuiltInReference = &value
	return b
}

// WithInt sets the Int field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Int field is set to the value of the last call.
func (b *CronJobSpecApplyConfiguration) WithInt(value int) *CronJobSpecApplyConfiguration {
	b.Int = &value
	return b
}

// WithAssociativeList adds the given value to the AssociativeList field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AssociativeList field.
func (b *CronJobSpecApplyConfiguration) WithAssociativeList(values ...*AssociativeTypeApplyConfiguration) *CronJobSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAssociativeList")
		}
		b.AssociativeList = append(b.AssociativeList, *values[i])
	}
	return b
}
