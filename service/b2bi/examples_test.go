// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package b2bi_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/b2bi"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// Sample CreateCapability call
//

func ExampleB2bi_CreateCapability_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.CreateCapabilityInput{
		ClientToken: aws.String("foo"),
		Configuration: &b2bi.CapabilityConfiguration{
			Edi: &b2bi.EdiConfiguration{
				InputLocation: &b2bi.S3Location{
					BucketName: aws.String("test-bucket"),
					Key:        aws.String("input/"),
				},
				OutputLocation: &b2bi.S3Location{
					BucketName: aws.String("test-bucket"),
					Key:        aws.String("output/"),
				},
				TransformerId: aws.String("tr-9a893cf536df4658b"),
				Type: &b2bi.EdiType{
					X12Details: &b2bi.X12Details{
						TransactionSet: aws.String("X12_110"),
						Version:        aws.String("VERSION_4010"),
					},
				},
			},
		},
		InstructionsDocuments: []*b2bi.S3Location{
			{
				BucketName: aws.String("test-bucket"),
				Key:        aws.String("instructiondoc.txt"),
			},
		},
		Name: aws.String("b2biexample"),
		Tags: []*b2bi.Tag{
			{
				Key:   aws.String("capabilityKey1"),
				Value: aws.String("capabilityValue1"),
			},
		},
		Type: aws.String("edi"),
	}

	result, err := svc.CreateCapability(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeConflictException:
				fmt.Println(b2bi.ErrCodeConflictException, aerr.Error())
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeServiceQuotaExceededException:
				fmt.Println(b2bi.ErrCodeServiceQuotaExceededException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample CreatePartnership call
//

func ExampleB2bi_CreatePartnership_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.CreatePartnershipInput{
		Capabilities: []*string{
			aws.String("ca-963a8121e4fc4e348"),
		},
		ClientToken: aws.String("foo"),
		Email:       aws.String("john@example.com"),
		Name:        aws.String("b2bipartner"),
		Phone:       aws.String("5555555555"),
		ProfileId:   aws.String("p-60fbc37c87f04fce9"),
		Tags: []*b2bi.Tag{
			{
				Key:   aws.String("sampleKey1"),
				Value: aws.String("sampleValue1"),
			},
		},
	}

	result, err := svc.CreatePartnership(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeConflictException:
				fmt.Println(b2bi.ErrCodeConflictException, aerr.Error())
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeServiceQuotaExceededException:
				fmt.Println(b2bi.ErrCodeServiceQuotaExceededException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample CreateProfile call
//

func ExampleB2bi_CreateProfile_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.CreateProfileInput{
		BusinessName: aws.String("John's Shipping"),
		ClientToken:  aws.String("foo"),
		Email:        aws.String("john@example.com"),
		Logging:      aws.String("ENABLED"),
		Name:         aws.String("Shipping Profile"),
		Phone:        aws.String("5555555555"),
		Tags: []*b2bi.Tag{
			{
				Key:   aws.String("sampleKey"),
				Value: aws.String("sampleValue"),
			},
		},
	}

	result, err := svc.CreateProfile(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeConflictException:
				fmt.Println(b2bi.ErrCodeConflictException, aerr.Error())
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeServiceQuotaExceededException:
				fmt.Println(b2bi.ErrCodeServiceQuotaExceededException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample CreateTransformer call
//

func ExampleB2bi_CreateTransformer_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.CreateTransformerInput{
		ClientToken: aws.String("foo"),
		EdiType: &b2bi.EdiType{
			X12Details: &b2bi.X12Details{
				TransactionSet: aws.String("X12_110"),
				Version:        aws.String("VERSION_4010"),
			},
		},
		FileFormat:      aws.String("JSON"),
		MappingTemplate: aws.String("{}"),
		Name:            aws.String("transformJSON"),
		SampleDocument:  aws.String("s3://test-bucket/sampleDoc.txt"),
		Tags: []*b2bi.Tag{
			{
				Key:   aws.String("sampleKey"),
				Value: aws.String("sampleValue"),
			},
		},
	}

	result, err := svc.CreateTransformer(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeConflictException:
				fmt.Println(b2bi.ErrCodeConflictException, aerr.Error())
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeServiceQuotaExceededException:
				fmt.Println(b2bi.ErrCodeServiceQuotaExceededException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample DeleteCapabilty call
//

func ExampleB2bi_DeleteCapability_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.DeleteCapabilityInput{
		CapabilityId: aws.String("ca-963a8121e4fc4e348"),
	}

	result, err := svc.DeleteCapability(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeConflictException:
				fmt.Println(b2bi.ErrCodeConflictException, aerr.Error())
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample DeletePartnership call
//

func ExampleB2bi_DeletePartnership_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.DeletePartnershipInput{
		PartnershipId: aws.String("ps-219fa02f5b4242af8"),
	}

	result, err := svc.DeletePartnership(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeConflictException:
				fmt.Println(b2bi.ErrCodeConflictException, aerr.Error())
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample DeleteProfile call
//

func ExampleB2bi_DeleteProfile_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.DeleteProfileInput{
		ProfileId: aws.String("p-60fbc37c87f04fce9"),
	}

	result, err := svc.DeleteProfile(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeConflictException:
				fmt.Println(b2bi.ErrCodeConflictException, aerr.Error())
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample DeleteTransformer call
//

func ExampleB2bi_DeleteTransformer_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.DeleteTransformerInput{
		TransformerId: aws.String("tr-974c129999f84d8c9"),
	}

	result, err := svc.DeleteTransformer(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeConflictException:
				fmt.Println(b2bi.ErrCodeConflictException, aerr.Error())
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample GetCapabilty call
//

func ExampleB2bi_GetCapability_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.GetCapabilityInput{
		CapabilityId: aws.String("ca-963a8121e4fc4e348"),
	}

	result, err := svc.GetCapability(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample GetPartnership call
//

func ExampleB2bi_GetPartnership_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.GetPartnershipInput{
		PartnershipId: aws.String("ps-219fa02f5b4242af8"),
	}

	result, err := svc.GetPartnership(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample GetProfile call
//

func ExampleB2bi_GetProfile_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.GetProfileInput{
		ProfileId: aws.String("p-60fbc37c87f04fce9"),
	}

	result, err := svc.GetProfile(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample GetTransformer call
//

func ExampleB2bi_GetTransformer_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.GetTransformerInput{
		TransformerId: aws.String("tr-974c129999f84d8c9"),
	}

	result, err := svc.GetTransformer(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample GetTransformerJob call
//

func ExampleB2bi_GetTransformerJob_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.GetTransformerJobInput{
		TransformerId:    aws.String("tr-974c129999f84d8c9"),
		TransformerJobId: aws.String("tj-vpYxfV7yQOqjMSYllEslLw"),
	}

	result, err := svc.GetTransformerJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample ListCapabilities call
//

func ExampleB2bi_ListCapabilities_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.ListCapabilitiesInput{
		MaxResults: aws.Int64(50),
		NextToken:  aws.String("foo"),
	}

	result, err := svc.ListCapabilities(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample ListPartnerships call
//

func ExampleB2bi_ListPartnerships_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.ListPartnershipsInput{
		MaxResults: aws.Int64(50),
		NextToken:  aws.String("foo"),
		ProfileId:  aws.String("p-60fbc37c87f04fce9"),
	}

	result, err := svc.ListPartnerships(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample ListProfiles call
//

func ExampleB2bi_ListProfiles_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.ListProfilesInput{
		MaxResults: aws.Int64(50),
		NextToken:  aws.String("foo"),
	}

	result, err := svc.ListProfiles(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample ListTagsForResources call
//

func ExampleB2bi_ListTagsForResource_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.ListTagsForResourceInput{
		ResourceARN: aws.String("arn:aws:b2bi:us-west-2:123456789012:profile/p-60fbc37c87f04fce9"),
	}

	result, err := svc.ListTagsForResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample ListTransformers call
//

func ExampleB2bi_ListTransformers_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.ListTransformersInput{
		MaxResults: aws.Int64(50),
		NextToken:  aws.String("foo"),
	}

	result, err := svc.ListTransformers(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample StartTransformerJob call
//

func ExampleB2bi_StartTransformerJob_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.StartTransformerJobInput{
		ClientToken: aws.String("foo"),
		InputFile: &b2bi.S3Location{
			BucketName: aws.String("test-bucket"),
			Key:        aws.String("input/inputFile.txt"),
		},
		OutputLocation: &b2bi.S3Location{
			BucketName: aws.String("test-bucket"),
			Key:        aws.String("output/"),
		},
		TransformerId: aws.String("tr-974c129999f84d8c9"),
	}

	result, err := svc.StartTransformerJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample TagResource call
//

func ExampleB2bi_TagResource_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.TagResourceInput{
		ResourceARN: aws.String("arn:aws:b2bi:us-west-2:123456789012:profile/p-60fbc37c87f04fce9"),
		Tags: []*b2bi.Tag{
			{
				Key:   aws.String("sampleKey"),
				Value: aws.String("SampleValue"),
			},
		},
	}

	result, err := svc.TagResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample TestMapping call
//

func ExampleB2bi_TestMapping_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.TestMappingInput{
		FileFormat:       aws.String("JSON"),
		InputFileContent: aws.String("Sample file content"),
		MappingTemplate:  aws.String("$"),
	}

	result, err := svc.TestMapping(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample TestParsing call
//

func ExampleB2bi_TestParsing_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.TestParsingInput{
		EdiType: &b2bi.EdiType{
			X12Details: &b2bi.X12Details{
				TransactionSet: aws.String("X12_110"),
				Version:        aws.String("VERSION_4010"),
			},
		},
		FileFormat: aws.String("JSON"),
		InputFile: &b2bi.S3Location{
			BucketName: aws.String("test-bucket"),
			Key:        aws.String("sampleFile.txt"),
		},
	}

	result, err := svc.TestParsing(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample UntagResource call
//

func ExampleB2bi_UntagResource_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.UntagResourceInput{
		ResourceARN: aws.String("arn:aws:b2bi:us-west-2:123456789012:profile/p-60fbc37c87f04fce9"),
		TagKeys: []*string{
			aws.String("sampleKey"),
		},
	}

	result, err := svc.UntagResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample UpdateCapability call
//

func ExampleB2bi_UpdateCapability_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.UpdateCapabilityInput{
		CapabilityId: aws.String("ca-963a8121e4fc4e348"),
		Configuration: &b2bi.CapabilityConfiguration{
			Edi: &b2bi.EdiConfiguration{
				InputLocation: &b2bi.S3Location{
					BucketName: aws.String("test-bucket"),
					Key:        aws.String("input/"),
				},
				OutputLocation: &b2bi.S3Location{
					BucketName: aws.String("test-bucket"),
					Key:        aws.String("output/"),
				},
				TransformerId: aws.String("tr-9a893cf536df4658b"),
				Type: &b2bi.EdiType{
					X12Details: &b2bi.X12Details{
						TransactionSet: aws.String("X12_110"),
						Version:        aws.String("VERSION_4010"),
					},
				},
			},
		},
		InstructionsDocuments: []*b2bi.S3Location{
			{
				BucketName: aws.String("test-bucket"),
				Key:        aws.String("instructiondoc.txt"),
			},
		},
		Name: aws.String("b2biexample"),
	}

	result, err := svc.UpdateCapability(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeConflictException:
				fmt.Println(b2bi.ErrCodeConflictException, aerr.Error())
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeServiceQuotaExceededException:
				fmt.Println(b2bi.ErrCodeServiceQuotaExceededException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample UpdatePartnership call
//

func ExampleB2bi_UpdatePartnership_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.UpdatePartnershipInput{
		Capabilities: []*string{
			aws.String("ca-963a8121e4fc4e348"),
		},
		Name:          aws.String("b2bipartner"),
		PartnershipId: aws.String("ps-219fa02f5b4242af8"),
	}

	result, err := svc.UpdatePartnership(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeConflictException:
				fmt.Println(b2bi.ErrCodeConflictException, aerr.Error())
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeServiceQuotaExceededException:
				fmt.Println(b2bi.ErrCodeServiceQuotaExceededException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample UpdateProfile call
//

func ExampleB2bi_UpdateProfile_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.UpdateProfileInput{
		BusinessName: aws.String("John's Shipping"),
		Email:        aws.String("john@example.com"),
		Name:         aws.String("Shipping Profile"),
		Phone:        aws.String("5555555555"),
		ProfileId:    aws.String("p-60fbc37c87f04fce9"),
	}

	result, err := svc.UpdateProfile(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeConflictException:
				fmt.Println(b2bi.ErrCodeConflictException, aerr.Error())
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeServiceQuotaExceededException:
				fmt.Println(b2bi.ErrCodeServiceQuotaExceededException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Sample UpdateTransformer call
//

func ExampleB2bi_UpdateTransformer_shared00() {
	svc := b2bi.New(session.New())
	input := &b2bi.UpdateTransformerInput{
		EdiType: &b2bi.EdiType{
			X12Details: &b2bi.X12Details{
				TransactionSet: aws.String("X12_110"),
				Version:        aws.String("VERSION_4010"),
			},
		},
		FileFormat:      aws.String("JSON"),
		MappingTemplate: aws.String("{}"),
		Name:            aws.String("transformJSON"),
		SampleDocument:  aws.String("s3://test-bucket/sampleDoc.txt"),
		Status:          aws.String("inactive"),
		TransformerId:   aws.String("tr-974c129999f84d8c9"),
	}

	result, err := svc.UpdateTransformer(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case b2bi.ErrCodeConflictException:
				fmt.Println(b2bi.ErrCodeConflictException, aerr.Error())
			case b2bi.ErrCodeAccessDeniedException:
				fmt.Println(b2bi.ErrCodeAccessDeniedException, aerr.Error())
			case b2bi.ErrCodeValidationException:
				fmt.Println(b2bi.ErrCodeValidationException, aerr.Error())
			case b2bi.ErrCodeThrottlingException:
				fmt.Println(b2bi.ErrCodeThrottlingException, aerr.Error())
			case b2bi.ErrCodeResourceNotFoundException:
				fmt.Println(b2bi.ErrCodeResourceNotFoundException, aerr.Error())
			case b2bi.ErrCodeServiceQuotaExceededException:
				fmt.Println(b2bi.ErrCodeServiceQuotaExceededException, aerr.Error())
			case b2bi.ErrCodeInternalServerException:
				fmt.Println(b2bi.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}