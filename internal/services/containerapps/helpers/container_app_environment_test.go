package helpers

import (
	"testing"
)

func TestExpandWorkloadProfiles_ConsumptionGPU(t *testing.T) {
	input := []WorkloadProfileModel{
		{
			Name:                "gpu-profile",
			WorkloadProfileType: "Consumption-GPU-NC24-A100",
			MaximumCount:        0,
			MinimumCount:        0,
		},
		{
			Name:                "gpu-profile-2",
			WorkloadProfileType: "Consumption-GPU-NC8as-T4",
			MaximumCount:        0,
			MinimumCount:        0,
		},
		{
			Name:                "dedicated-profile",
			WorkloadProfileType: "D4",
			MaximumCount:        1,
			MinimumCount:        1,
		},
	}

	result := ExpandWorkloadProfiles(input)

	if result == nil || len(*result) != 3 {
		t.Fatalf("Expected 3 profiles, got %v", result)
	}

	// Case 1: Consumption-GPU-NC24-A100
	profile1 := (*result)[0]
	if profile1.WorkloadProfileType != "Consumption-GPU-NC24-A100" {
		t.Errorf("Expected WorkloadProfileType to be Consumption-GPU-NC24-A100, got %s", profile1.WorkloadProfileType)
	}
	if profile1.MaximumCount != nil {
		t.Errorf("Expected MaximumCount to be nil for Consumption-GPU-NC24-A100, got %d", *profile1.MaximumCount)
	}
	if profile1.MinimumCount != nil {
		t.Errorf("Expected MinimumCount to be nil for Consumption-GPU-NC24-A100, got %d", *profile1.MinimumCount)
	}

	// Case 2: Consumption-GPU-NC8as-T4
	profile2 := (*result)[1]
	if profile2.WorkloadProfileType != "Consumption-GPU-NC8as-T4" {
		t.Errorf("Expected WorkloadProfileType to be Consumption-GPU-NC8as-T4, got %s", profile2.WorkloadProfileType)
	}
	if profile2.MaximumCount != nil {
		t.Errorf("Expected MaximumCount to be nil for Consumption-GPU-NC8as-T4, got %d", *profile2.MaximumCount)
	}
	if profile2.MinimumCount != nil {
		t.Errorf("Expected MinimumCount to be nil for Consumption-GPU-NC8as-T4, got %d", *profile2.MinimumCount)
	}

	// Case 3: Dedicated D4
	profile3 := (*result)[2]
	if profile3.WorkloadProfileType != "D4" {
		t.Errorf("Expected WorkloadProfileType to be D4, got %s", profile3.WorkloadProfileType)
	}
	if profile3.MaximumCount == nil || *profile3.MaximumCount != 1 {
		t.Errorf("Expected MaximumCount to be 1 for D4, got %v", profile3.MaximumCount)
	}
	if profile3.MinimumCount == nil || *profile3.MinimumCount != 1 {
		t.Errorf("Expected MinimumCount to be 1 for D4, got %v", profile3.MinimumCount)
	}
}
