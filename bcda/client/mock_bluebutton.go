package client

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	models "github.com/CMSgov/bcda-app/bcda/models/fhir"

	"github.com/stretchr/testify/mock"
)

type MockBlueButtonClient struct {
	mock.Mock
	HICN *string
	MBI  *string
}

func (bbc *MockBlueButtonClient) GetExplanationOfBenefit(patientID, jobID, cmsID, since string, transactionTime time.Time, serviceDate ClaimsWindow) (*models.Bundle, error) {
	args := bbc.Called(patientID, jobID, cmsID, since, transactionTime, serviceDate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Bundle), args.Error(1)
}

func (bbc *MockBlueButtonClient) GetPatientByIdentifierHash(hashedIdentifier string) (string, error) {
	args := bbc.Called(hashedIdentifier)
	return args.String(0), args.Error(1)
}

func (bbc *MockBlueButtonClient) GetPatient(patientID, jobID, cmsID, since string, transactionTime time.Time) (*models.Bundle, error) {
	args := bbc.Called(patientID, jobID, cmsID, since, transactionTime)
	return args.Get(0).(*models.Bundle), args.Error(1)
}

func (bbc *MockBlueButtonClient) GetCoverage(beneficiaryID, jobID, cmsID, since string, transactionTime time.Time) (*models.Bundle, error) {
	args := bbc.Called(beneficiaryID, jobID, cmsID, since, transactionTime)
	return args.Get(0).(*models.Bundle), args.Error(1)
}

func (bbc *MockBlueButtonClient) GetClaim(mbi, jobID, cmsID, since string, transactionTime time.Time, claimsWindow ClaimsWindow) (*models.Bundle, error) {
	args := bbc.Called(mbi, jobID, cmsID, since, transactionTime, claimsWindow)
	return args.Get(0).(*models.Bundle), args.Error(1)
}

func (bbc *MockBlueButtonClient) GetClaimResponse(mbi, jobID, cmsID, since string, transactionTime time.Time, claimsWindow ClaimsWindow) (*models.Bundle, error) {
	args := bbc.Called(mbi, jobID, cmsID, since, transactionTime, claimsWindow)
	return args.Get(0).(*models.Bundle), args.Error(1)
}

// Returns copy of a static json file (From Blue Button Sandbox originally) after replacing the patient ID of 20000000000001 with the requested identifier
// This is private in the real function and should remain so, but in the test client it makes maintenance easier to expose it.
func (bbc *MockBlueButtonClient) GetData(endpoint, patientID string) (string, error) {
	var fData []byte
	fData, err := ioutil.ReadFile(filepath.Join("../shared_files/synthetic_beneficiary_data/", filepath.Clean(endpoint)))
	if err != nil {
		fData, err = ioutil.ReadFile(filepath.Join("../../shared_files/synthetic_beneficiary_data/", filepath.Clean(endpoint)))
		if err != nil {
			return "", err
		}
	}
	cleanData := strings.Replace(string(fData), "20000000000001", patientID, -1)
	if bbc.MBI != nil {
		// no longer hashed, but this is only a test file with synthetic test data
		cleanData = strings.Replace(cleanData, "-1Q03Z002871", *bbc.MBI, -1)
	}
	return cleanData, err
}

func (bbc *MockBlueButtonClient) GetBundleData(endpoint, patientID string) (*models.Bundle, error) {
	payload, err := bbc.GetData(endpoint, patientID)
	if err != nil {
		return nil, err
	}

	var b models.Bundle
	err = json.Unmarshal([]byte(payload), &b)
	if err != nil {
		return nil, err
	}

	return &b, err
}
