package service

import (
	"context"
	"fmt"
	"regexp"
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

	acoID  = "SOME_ACO"
	acoCfg = ACOConfig{
		patternExp:    regexp.MustCompile(acoID),
		LookbackYears: 3,
		perfYear:      time.Now(),
	}
	acoCfgs = map[*regexp.Regexp]*ACOConfig{acoCfg.patternExp: &acoCfg}
)

type claimsWindow struct {
	LowerBound time.Time
	UpperBound time.Time
}

func TestClaims(t *testing.T) {
	cmsID := "CMS_ID_ClAIMS"
	uuid := uuid.NewUUID()

	historicalACO := models.ACO{
		UUID:  uuid,
		CMSID: &cmsID,
		TerminationDetails: &models.Termination{
			BlacklistType:   models.Limited,
			ClaimsStrategy:  models.ClaimsHistorical,
			TerminationDate: pastDate,
			CutoffDate:      futureDate,
		},
	}
	histClaimsWindow := claimsWindow{
		UpperBound: historicalACO.TerminationDetails.ClaimsDate(),
	}
	latestACO := models.ACO{
		UUID:  uuid,
		CMSID: &cmsID,
		TerminationDetails: &models.Termination{
			BlacklistType:   models.Limited,
			ClaimsStrategy:  models.ClaimsLatest,
			TerminationDate: pastDate,
			CutoffDate:      futureDate,
		},
	}
	latestClaimsWindow := claimsWindow{
		UpperBound: latestACO.TerminationDetails.ClaimsDate(),
	}
	futureACO := models.ACO{
		UUID:  uuid,
		CMSID: &cmsID,
		TerminationDetails: &models.Termination{
			BlacklistType:   models.Limited,
			ClaimsStrategy:  models.ClaimsHistorical,
			TerminationDate: pastDate,
			CutoffDate:      futureDate,
		},
	}
	futureClaimsWindow := claimsWindow{
		UpperBound: futureACO.TerminationDetails.ClaimsDate(),
	}

	tests := []struct {
		name            string
		aco             models.ACO
		expBenes        []*models.CCLFBeneficiary
		expBenesMBIs    []string
		expClaimsWindow claimsWindow
		since           time.Time
		terminationDate time.Time
	}{
		{"Historical Claims", historicalACO, oldBenes, oldBeneMBIs, histClaimsWindow, pastDate, historicalACO.TerminationDetails.TerminationDate},              // "_since" after term date should only return old benes
		{"Latest Claims", latestACO, allBenes, allBeneMBIs, latestClaimsWindow, pastDate, latestACO.TerminationDetails.TerminationDate},                        // should receive all benes
		{"Future Termination Historical Claims", futureACO, allBenes, allBeneMBIs, futureClaimsWindow, pastDate, futureACO.TerminationDetails.TerminationDate}, // should receive all benes since term date is in future
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			verifyRequest(t, tt.name, tt.aco, tt.expBenes, tt.expBenesMBIs, tt.since, tt.expClaimsWindow, tt.terminationDate)
		})
	}
}

func TestOptOut(t *testing.T) {

}

func TestAttribution(t *testing.T) {
	cmsID := "CMS_ID_ATTR"
	uuid := uuid.NewUUID()

	// historicalACO := models.ACO{
	// 	UUID:  uuid,
	// 	CMSID: &cmsID,
	// 	TerminationDetails: &models.Termination{
	// 		BlacklistType:       models.Limited,
	// 		AttributionStrategy: models.AttributionHistorical,
	// 		TerminationDate:     pastDate,
	// 		CutoffDate:          futureDate,
	// 	},
	// }
	// histAttrDate := historicalACO.TerminationDetails.AttributionDate()

	latestACO := models.ACO{
		UUID:  uuid,
		CMSID: &cmsID,
		TerminationDetails: &models.Termination{
			BlacklistType:       models.Limited,
			AttributionStrategy: models.AttributionLatest,
			TerminationDate:     pastDate,
			CutoffDate:          futureDate,
		},
	}
	latestAttrDate := latestACO.TerminationDetails.AttributionDate()
	fmt.Println("latest attr date:", latestAttrDate)
	// futureTerminationACO := models.ACO{
	// 	UUID:  uuid,
	// 	CMSID: &cmsID,
	// 	TerminationDetails: &models.Termination{
	// 		BlacklistType:       models.Limited,
	// 		AttributionStrategy: models.AttributionHistorical,
	// 		TerminationDate:     futureDate,
	// 		CutoffDate:          futureDate.Add(30 * 24 * time.Hour),
	// 	},
	// }
	// futureAttrDate := futureTerminationACO.TerminationDetails.AttributionDate()

	tests := []struct {
		name            string
		aco             models.ACO
		expBenes        []*models.CCLFBeneficiary
		expBenesMBIs    []string
		expClaimsWindow claimsWindow
		since           time.Time
		upperBound      time.Time
	}{
		// {"Historical Attribution", historicalACO, oldBenes, oldBeneMBIs, claimsWindow{UpperBound: histAttrDate}, histAttrDate.Add(1 * 24 * time.Hour), histAttrDate},   // "_since" after term date should only return old benes
		{"Latest Attribution", latestACO, allBenes, allBeneMBIs, claimsWindow{}, pastDate, latestAttrDate}, // should receive all benes
		// {"Future Termination Historical Attribution", futureTerminationACO, allBenes, allBeneMBIs, claimsWindow{UpperBound: futureAttrDate}, pastDate, futureAttrDate}, // should receive all benes since term date is in future
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("&&&", tt.since, tt.upperBound, "~*~*~")
			verifyRequest(t, tt.name, tt.aco, tt.expBenes, tt.expBenesMBIs, tt.since, tt.expClaimsWindow, tt.upperBound)
		})
	}
}

func verifyRequest(
	t *testing.T,
	jobID string,
	aco models.ACO,
	expBenes []*models.CCLFBeneficiary,
	expBenesMBIs []string,
	expSince time.Time, // since
	expClaimsWindow claimsWindow,
	upperBound time.Time, // termination date
) {
	rc := RequestConditions{
		ReqType:   RetrieveNewBeneHistData,
		Resources: []string{"ExplanationOfBenefit", "Patient", "Coverage"},
		CMSID:     jobID,
		ACOID:     aco.UUID,
		Since:     expSince,
	}
	fmt.Println("***", jobID, "since:", expSince, "upper:", upperBound)
	repository := &models.MockRepository{}
	svc := NewService(repository, &Config{cutoffDuration: 45}, "/v2/fhir")
	svc.(*service).acoConfig = acoCfgs

	repository.On("GetACOByCMSID", testUtils.CtxMatcher, rc.CMSID).
		Return(&models.ACO{UUID: rc.ACOID, TerminationDetails: aco.TerminationDetails}, nil)
	repository.On("GetLatestCCLFFile", testUtils.CtxMatcher, rc.CMSID, mock.Anything, mock.Anything, mock.Anything, upperBound, mock.Anything).
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
		assert.Equal(t, expClaimsWindow.LowerBound, job.ClaimsWindow.LowerBound)
		assert.Equal(t, expClaimsWindow.UpperBound, job.ClaimsWindow.UpperBound)
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
