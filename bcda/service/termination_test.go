package service

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/CMSgov/bcda-app/bcda/models"
	"github.com/CMSgov/bcda-app/bcda/testUtils"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	pastDate   = time.Now().Add(-30 * 24 * time.Hour).Round(time.Millisecond).UTC()
	futureDate = time.Now().Add(30 * 24 * time.Hour).Round(time.Millisecond).UTC()

	allBenes, allBeneMBIs = generateBeneMBIs(15)
	oldBenes, oldBeneMBIs = allBenes[0:10], allBeneMBIs[0:10] // benes attributed before termination date
)

func TestClaims(t *testing.T) {

}

// func TestOptOut(t *testing.T) {

// }

func TestAttribution(t *testing.T) {
	cmsID := "CMS_ID_ATTR"
	uuid := uuid.NewUUID()

	historicalACO := models.ACO{
		UUID:  uuid,
		CMSID: &cmsID,
		TerminationDetails: &models.Termination{
			AttributionStrategy: models.AttributionHistorical,
			TerminationDate:     pastDate,
		},
	}
	histAttrDate := historicalACO.TerminationDetails.AttributionDate()

	latestACO := models.ACO{
		UUID:  uuid,
		CMSID: &cmsID,
		TerminationDetails: &models.Termination{
			AttributionStrategy: models.AttributionLatest,
			TerminationDate:     pastDate,
		},
	}
	latestAttrDate := latestACO.TerminationDetails.AttributionDate()

	futureTerminationACO := models.ACO{
		UUID:  uuid,
		CMSID: &cmsID,
		TerminationDetails: &models.Termination{
			AttributionStrategy: models.AttributionHistorical,
			TerminationDate:     futureDate,
		},
	}
	futureAttrDate := futureTerminationACO.TerminationDetails.AttributionDate()

	tests := []struct {
		name            string
		aco             models.ACO
		expBenes        []*models.CCLFBeneficiary
		expBenesMBIs    []string
		since           time.Time
		terminationDate time.Time
	}{
		{"Historical Attribution", historicalACO, oldBenes, oldBeneMBIs, histAttrDate.Add(1 * 24 * time.Hour), histAttrDate}, // "_since" after term date should only return old benes
		{"Latest Attribution", latestACO, allBenes, allBeneMBIs, pastDate, latestAttrDate},                                   // should receive all benes
		{"Future Termination Historical Attribution", futureTerminationACO, allBenes, allBeneMBIs, pastDate, futureAttrDate}, // should receive all benes since term date is in future
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			verifyRequest(t, tt.name, tt.aco, tt.expBenes, tt.expBenesMBIs, tt.since, tt.terminationDate)
		})
	}
}

func verifyRequest(
	t *testing.T,
	jobID string,
	aco models.ACO,
	expBenes []*models.CCLFBeneficiary,
	expBenesMBIs []string,
	expSince time.Time,
	attrUpperBound time.Time,
) {
	rc := RequestConditions{
		ReqType:   RetrieveNewBeneHistData,
		Resources: []string{"ExplanationOfBenefit", "Patient", "Coverage"},
		CMSID:     *aco.CMSID,
		ACOID:     aco.UUID,
		Since:     expSince,
	}

	repository := &models.MockRepository{}
	svc := NewService(repository, &Config{}, "/v2/fhir")
	repository.On("GetACOByCMSID", testUtils.CtxMatcher, rc.CMSID).
		Return(&models.ACO{UUID: rc.ACOID, TerminationDetails: aco.TerminationDetails}, nil)
	repository.On("GetLatestCCLFFile", testUtils.CtxMatcher, rc.CMSID, mock.Anything, mock.Anything, mock.Anything, attrUpperBound, mock.Anything).
		Return(getCCLFFile(1), nil)
	repository.On("GetSuppressedMBIs", testUtils.CtxMatcher, mock.Anything, mock.Anything).
		Return(nil, nil)
	repository.On("GetCCLFBeneficiaries", testUtils.CtxMatcher, mock.Anything, mock.Anything).
		Return(expBenes, nil)
	repository.On("GetCCLFBeneficiaryMBIs", testUtils.CtxMatcher, mock.Anything).
		Return(expBenesMBIs, nil)

	qj, err := svc.GetQueJobs(context.Background(), rc)

	assert.Equal(t, len(rc.Resources), len(qj))
	assert.Nil(t, err)

	for _, job := range qj {
		assert.Equal(t, expBenesMBIs, job.BeneficiaryIDs)
	}
}

func generateBeneMBIs(beneCount int) (benes []*models.CCLFBeneficiary, ids []string) {
	beneArray := make([]*models.CCLFBeneficiary, beneCount)
	beneIDs := make([]string, beneCount)
	for i := 0; i < len(beneArray); i++ {
		id := uint(i*1000 + i + 1)
		beneArray[i] = &models.CCLFBeneficiary{
			ID:  id,
			MBI: fmt.Sprintf("MBI%d", id),
		}
		beneIDs[i] = strconv.Itoa(int(id))
	}
	return beneArray, beneIDs
}
