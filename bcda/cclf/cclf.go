package cclf

import (
	"archive/zip"
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jackc/pgx/stdlib"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/CMSgov/bcda-app/bcda/cclf/metrics"
	"github.com/CMSgov/bcda-app/bcda/constants"
	"github.com/CMSgov/bcda-app/bcda/database"
	"github.com/CMSgov/bcda-app/bcda/models"
	"github.com/CMSgov/bcda-app/bcda/models/postgres"
	"github.com/CMSgov/bcda-app/bcda/utils"
	"github.com/CMSgov/bcda-app/conf"
	"github.com/CMSgov/bcda-app/log"
)

type cclfFileMetadata struct {
	name         string
	env          string
	acoID        string
	cclfNum      int
	perfYear     int
	timestamp    time.Time
	filePath     string
	imported     bool
	deliveryDate time.Time
	fileID       uint
	fileType     models.CCLFFileType
}

type cclfFileValidator struct {
	totalRecordCount int
	maxRecordLength  int
}

func importCCLF0(ctx context.Context, fileMetadata *cclfFileMetadata) (map[string]cclfFileValidator, error) {
	if fileMetadata == nil {
		fmt.Println("File CCLF0 not found.")
		err := errors.New("file CCLF0 not found")
		log.API.Error(err)
		return nil, err
	}

	fmt.Printf("Importing CCLF0 file %s...\n", fileMetadata)
	log.API.Infof("Importing CCLF0 file %s...", fileMetadata)

	r, err := zip.OpenReader(filepath.Clean(fileMetadata.filePath))
	if err != nil {
		fmt.Printf("Could not read CCLF0 archive %s.\n", fileMetadata)
		err := errors.Wrapf(err, "could not read CCLF0 archive %s", fileMetadata)
		log.API.Error(err)
		return nil, err
	}
	defer r.Close()

	const (
		fileNumStart, fileNumEnd           = 0, 7
		totalRecordStart, totalRecordEnd   = 52, 63
		recordLengthStart, recordLengthEnd = 64, 69
	)

	close := metrics.NewChild(ctx, "importCCLF0")
	defer close()

	var validator map[string]cclfFileValidator
	var rawFile *zip.File

	for _, f := range r.File {
		// iterate in this zipped folder until we find our cclf0 file
		if f.Name == fileMetadata.name {
			rawFile = f
			fmt.Printf(constants.ReadingFileN, fileMetadata.name, fileMetadata.filePath)
			log.API.Infof("Reading file %s from archive %s", fileMetadata.name, fileMetadata.filePath)
		}
	}

	if rawFile == nil {
		fmt.Printf(constants.FileNotFoundN, fileMetadata.name, fileMetadata.filePath)
		err = errors.Wrapf(err, constants.FileNotFound, fileMetadata.name, fileMetadata.filePath)
		log.API.Error(err)
		return nil, err
	}

	rc, err := rawFile.Open()
	if err != nil {
		fmt.Printf("Could not read file %s in CCLF0 archive %s.\n", fileMetadata.name, fileMetadata.filePath)
		err = errors.Wrapf(err, "could not read file %s in CCLF0 archive %s", fileMetadata.name, fileMetadata.filePath)
		log.API.Error(err)
		return nil, err
	}
	defer rc.Close()
	sc := bufio.NewScanner(rc)
	for sc.Scan() {
		b := sc.Bytes()
		if len(bytes.TrimSpace(b)) > 0 {
			filetype := string(bytes.TrimSpace(b[fileNumStart:fileNumEnd]))

			if filetype == "CCLF8" {
				if validator == nil {
					validator = make(map[string]cclfFileValidator)
				}

				if _, ok := validator[filetype]; ok {
					fmt.Printf("Duplicate %v file type found from CCLF0 file.\n", filetype)
					err := fmt.Errorf("duplicate %v file type found from CCLF0 file", filetype)
					log.API.Error(err)
					return nil, err
				}

				count, err := strconv.Atoi(string(bytes.TrimSpace(b[totalRecordStart:totalRecordEnd])))
				if err != nil {
					fmt.Printf("Failed to parse %s record count from CCLF0 file.\n", filetype)
					err = errors.Wrapf(err, "failed to parse %s record count from CCLF0 file", filetype)
					log.API.Error(err)
					return nil, err
				}
				length, err := strconv.Atoi(string(bytes.TrimSpace(b[recordLengthStart:recordLengthEnd])))
				if err != nil {
					fmt.Printf("Failed to parse %s record length from CCLF0 file.\n", filetype)
					err = errors.Wrapf(err, "failed to parse %s record length from CCLF0 file", filetype)
					log.API.Error(err)
					return nil, err
				}
				validator[filetype] = cclfFileValidator{totalRecordCount: count, maxRecordLength: length}
			}
		}
	}

	if _, ok := validator["CCLF8"]; !ok {
		fmt.Printf("Failed to parse CCLF8 from CCLF0 file %s.\n", fileMetadata)
		err := fmt.Errorf("failed to parse CCLF8 from CCLF0 file %s", fileMetadata)
		log.API.Error(err)
		return nil, err
	}
	fmt.Printf("Successfully imported CCLF0 file %s.\n", fileMetadata)
	log.API.Infof("Successfully imported CCLF0 file %s.", fileMetadata)

	return validator, nil
}

func importCCLF8(ctx context.Context, fileMetadata *cclfFileMetadata) (err error) {
	db := database.Connection
	repository := postgres.NewRepository(db)

	exists, err := repository.GetCCLFFileExistsByName(ctx, fileMetadata.name)
	if err != nil {
		fmt.Printf("failed to check existence of CCLF%d file.\n", fileMetadata.cclfNum)
		err = errors.Wrapf(err, "failed to check existence of CCLF%d file", fileMetadata.cclfNum)
		log.API.Error(err)
		return err
	}

	if exists {
		fmt.Printf("CCL%d file %s already exists in database, skipping import...", fileMetadata.cclfNum, fileMetadata)
		log.API.Infof("CCL%d file %s already exists in database, skipping import...", fileMetadata.cclfNum, fileMetadata)
		return nil
	}

	fmt.Printf("Importing CCLF%d file %s...\n", fileMetadata.cclfNum, fileMetadata)
	log.API.Infof("Importing CCLF%d file %s...", fileMetadata.cclfNum, fileMetadata)

	conn, err := stdlib.AcquireConn(db)
	defer utils.CloseAndLog(logrus.WarnLevel, func() error { return stdlib.ReleaseConn(db, conn) })

	tx, err := conn.BeginEx(ctx, nil)
	if err != nil {
		err = fmt.Errorf("failed to start transaction: %w", err)
		log.API.Error(err)

		return err
	}

	rtx := postgres.NewRepositoryPgxTx(tx)

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.API.Warnf("Failed to rollback transaction %s", err.Error())
			}
			return
		}
	}()

	r, err := zip.OpenReader(filepath.Clean(fileMetadata.filePath))
	if err != nil {
		fmt.Printf("Could not read CCLF%d archive %s.\n", fileMetadata.cclfNum, fileMetadata.filePath)
		err = errors.Wrapf(err, "could not read CCLF%d archive %s", fileMetadata.cclfNum, fileMetadata.filePath)
		log.API.Error(err)
		return err
	}
	defer r.Close()

	if len(r.File) < 1 {
		fmt.Printf("No files found in CCLF%d archive %s.\n", fileMetadata.cclfNum, fileMetadata.filePath)
		err = fmt.Errorf("no files found in CCLF%d archive %s", fileMetadata.cclfNum, fileMetadata.filePath)
		log.API.Error(err)
		return err
	}

	close := metrics.NewChild(ctx, fmt.Sprintf("importCCLF%d", fileMetadata.cclfNum))
	defer close()

	cclfFile := models.CCLFFile{
		CCLFNum:         fileMetadata.cclfNum,
		Name:            fileMetadata.name,
		ACOCMSID:        fileMetadata.acoID,
		Timestamp:       fileMetadata.timestamp,
		PerformanceYear: fileMetadata.perfYear,
		ImportStatus:    constants.ImportInprog,
		Type:            fileMetadata.fileType,
	}

	cclfFile.ID, err = rtx.CreateCCLFFile(ctx, cclfFile)
	if err != nil {
		fmt.Printf("Could not create CCLF%d file record.\n", fileMetadata.cclfNum)
		err = errors.Wrapf(err, "could not create CCLF%d file record", fileMetadata.cclfNum)
		log.API.Error(err)
		return err
	}

	fileMetadata.fileID = cclfFile.ID

	var rawFile *zip.File

	for _, f := range r.File {
		if f.Name == fileMetadata.name {
			rawFile = f
			fmt.Printf(constants.ReadingFileN, fileMetadata.name, fileMetadata.filePath)
			log.API.Infof("Reading file %s from archive %s", fileMetadata.name, fileMetadata.filePath)
		}
	}

	if rawFile == nil {
		fmt.Printf(constants.FileNotFoundN, fileMetadata.name, fileMetadata.filePath)
		err = fmt.Errorf(constants.FileNotFound, fileMetadata.name, fileMetadata.filePath)
		log.API.Error(err)
		return err
	}

	rc, err := rawFile.Open()
	if err != nil {
		fmt.Printf("Could not read file %s for CCLF%d in archive %s.\n", cclfFile.Name, fileMetadata.cclfNum, fileMetadata.filePath)
		err = errors.Wrapf(err, "could not read file %s for CCLF%d in archive %s", cclfFile.Name, fileMetadata.cclfNum, fileMetadata.filePath)
		log.API.Error(err)
		return err
	}
	defer rc.Close()
	sc := bufio.NewScanner(rc)

	importedCount, err := CopyFrom(ctx, tx, sc, cclfFile.ID, utils.GetEnvInt("CCLF_IMPORT_STATUS_RECORDS_INTERVAL", 10000))
	if err != nil {
		return errors.Wrap(err, "failed to copy data to beneficiaries table")
	}

	err = rtx.UpdateCCLFFileImportStatus(ctx, fileMetadata.fileID, constants.ImportComplete)
	if err != nil {
		fmt.Printf("Could not update cclf file record for file: %s. \n", fileMetadata)
		err = errors.Wrapf(err, "could not update cclf file record for file: %s.", fileMetadata)
		log.API.Error(err)
	}

	if err = tx.Commit(); err != nil {
		log.API.Error(err.Error())
		failMsg := fmt.Sprintf("failed to commit transaction for CCLF%d import file %s", fileMetadata.cclfNum, fileMetadata)
		return errors.Wrap(err, failMsg)
	}

	successMsg := fmt.Sprintf("Successfully imported %d records from CCLF%d file %s.", importedCount, fileMetadata.cclfNum, fileMetadata)
	fmt.Println(successMsg)
	log.API.Infof(successMsg)

	return nil
}

func ImportCCLFDirectory(filePath string) (success, failure, skipped int, err error) {
	t := metrics.GetTimer()
	defer t.Close()
	ctx := metrics.NewContext(context.Background(), t)

	// We are not going to create any children from this parent so we can
	// safely ignored the returned context.
	_, c := metrics.NewParent(ctx, "ImportCCLFDirectory#sortCCLFArchives")
	cclfMap, skipped, err := processCCLFArchives(filePath)
	c()

	if err != nil {
		return 0, 0, 0, err
	}

	if len(cclfMap) == 0 {
		log.API.Info("Failed to find any CCLF files in directory")
		return 0, 0, skipped, nil
	}

	for acoID := range cclfMap {
		func() {
			ctx, c := metrics.NewParent(ctx, "ImportCCLFDirectory#processACOs")
			defer c()
			for _, cclfFiles := range cclfMap[acoID] {
				var cclf0, cclf8 *cclfFileMetadata
				for _, cclf := range cclfFiles {
					if cclf.cclfNum == 0 {
						cclf0 = cclf
					} else if cclf.cclfNum == 8 {
						cclf8 = cclf
					}
				}
				cclfvalidator, err := importCCLF0(ctx, cclf0)
				if err != nil {
					fmt.Printf("Failed to import CCLF0 file: %s, Skipping CCLF8 file: %s.\n ", cclf0, cclf8)
					log.API.Errorf("Failed to import CCLF0 file: %s, Skipping CCLF8 file: %s ", cclf0, cclf8)
					failure++
					skipped += 2
					continue
				} else {
					success++
				}
				err = validate(ctx, cclf8, cclfvalidator)
				if err != nil {
					fmt.Printf("Failed to validate CCLF8 file: %s.\n", cclf8)
					log.API.Errorf("Failed to validate CCLF8 file: %s", cclf8)
					failure++
				} else {
					if err = importCCLF8(ctx, cclf8); err != nil {
						fmt.Printf("Failed to import CCLF8 file: %s %s.\n", cclf8, err)
						log.API.Errorf("Failed to import CCLF8 file: %s %s", cclf8, err)
						failure++
					} else {
						cclf8.imported = true
						success++
					}
				}
				cclf0.imported = cclf8 != nil && cclf8.imported
			}
		}()
	}

	if err = func() error {
		ctx, c := metrics.NewParent(ctx, "ImportCCLFDirectory#cleanupCCLF")
		defer c()
		return cleanUpCCLF(ctx, cclfMap)
	}(); err != nil {
		log.API.Error(err)
	}

	if failure > 0 {
		err = errors.New("one or more files failed to import correctly")
		log.API.Error(err)
	} else {
		err = nil
	}

	return success, failure, skipped, err
}

func validate(ctx context.Context, fileMetadata *cclfFileMetadata, cclfFileValidator map[string]cclfFileValidator) error {
	if fileMetadata == nil {
		fmt.Printf("File not found.\n")
		err := errors.New("file not found")
		log.API.Error(err)
		return err
	}

	fmt.Printf("Validating CCLF%d file %s...\n", fileMetadata.cclfNum, fileMetadata)
	log.API.Infof("Validating CCLF%d file %s...", fileMetadata.cclfNum, fileMetadata)

	var key string
	if fileMetadata.cclfNum == 8 {
		key = "CCLF8"
	} else {
		fmt.Printf("Unknown file type when validating file: %s.\n", fileMetadata)
		err := fmt.Errorf("unknown file type when validating file: %s", fileMetadata)
		log.API.Error(err)
		return err
	}

	r, err := zip.OpenReader(filepath.Clean(fileMetadata.filePath))
	if err != nil {
		fmt.Printf("Could not read archive %s.\n", fileMetadata.filePath)
		err := errors.Wrapf(err, "could not read archive %s", fileMetadata.filePath)
		log.API.Error(err)
		return err
	}
	defer r.Close()

	close := metrics.NewChild(ctx, "validate")
	defer close()

	count := 0
	validator := cclfFileValidator[key]
	var rawFile *zip.File

	for _, f := range r.File {
		if f.Name == fileMetadata.name {
			rawFile = f
			fmt.Printf(constants.ReadingFileN, fileMetadata.name, fileMetadata.filePath)
			log.API.Infof("Reading file %s from archive %s", fileMetadata.name, fileMetadata.filePath)
		}
	}

	if rawFile == nil {
		fmt.Printf(constants.FileNotFoundN, fileMetadata.name, fileMetadata.filePath)
		err = errors.Wrapf(err, constants.FileNotFound, fileMetadata.name, fileMetadata.filePath)
		log.API.Error(err)
		return err
	}

	rc, err := rawFile.Open()
	if err != nil {
		fmt.Printf("Could not read file %s in archive %s.\n", fileMetadata.name, fileMetadata.filePath)
		err = errors.Wrapf(err, "could not read file %s in archive %s", fileMetadata.name, fileMetadata.filePath)
		log.API.Error(err)
		return err
	}
	defer rc.Close()
	sc := bufio.NewScanner(rc)
	for sc.Scan() {
		b := sc.Bytes()
		bytelength := len(bytes.TrimSpace(b))
		if bytelength > 0 && bytelength <= validator.maxRecordLength {
			count++

			// currently only errors if there are more records than we expect.
			if count > validator.totalRecordCount {
				fmt.Printf("Maximum record count reached for file %s, Expected record count: %d, Actual record count: %d.\n", key, validator.totalRecordCount, count)
				err := fmt.Errorf("maximum record count reached for file %s (expected: %d, actual: %d)", key, validator.totalRecordCount, count)
				log.API.Error(err)
				return err
			}
		} else {
			fmt.Printf("Incorrect record length for file %s, Expected record length: %d, Actual record length: %d.\n", key, validator.maxRecordLength, bytelength)
			err := fmt.Errorf("incorrect record length for file %s (expected: %d, actual: %d)", key, validator.maxRecordLength, bytelength)
			log.API.Error(err)
			return err
		}
	}
	fmt.Printf("Successfully validated CCLF%d file %s.\n", fileMetadata.cclfNum, fileMetadata)
	log.API.Infof("Successfully validated CCLF%d file %s.", fileMetadata.cclfNum, fileMetadata)
	return nil
}

func cleanUpCCLF(ctx context.Context, cclfMap map[string]map[metadataKey][]*cclfFileMetadata) error {
	errCount := 0
	for _, cclfFileMap := range cclfMap {
		for _, cclfFileList := range cclfFileMap {
			for _, cclf := range cclfFileList {
				func() {
					close := metrics.NewChild(ctx, fmt.Sprintf("cleanUpCCLF%d", cclf.cclfNum))
					defer close()

					fmt.Printf("Cleaning up file %s.\n", cclf.filePath)
					log.API.Infof("Cleaning up file %s", cclf.filePath)
					folderName := filepath.Base(cclf.filePath)
					newpath := fmt.Sprintf("%s/%s", conf.GetEnv("PENDING_DELETION_DIR"), folderName)
					if !cclf.imported {
						// check the timestamp on the failed files
						elapsed := time.Since(cclf.deliveryDate).Hours()
						deleteThreshold := utils.GetEnvInt("BCDA_ETL_FILE_ARCHIVE_THRESHOLD_HR", 72)
						if int(elapsed) > deleteThreshold {
							if _, err := os.Stat(newpath); err == nil {
								return
							}
							// move the (un)successful files to the deletion dir
							err := os.Rename(cclf.filePath, newpath)
							if err != nil {
								errCount++
								errMsg := fmt.Sprintf("File %s failed to clean up properly: %v", cclf.filePath, err)
								fmt.Println(errMsg)
								log.API.Error(errMsg)
							} else {
								fmt.Printf("File %s never ingested, moved to the pending deletion dir.\n", cclf.filePath)
								log.API.Infof("File %s never ingested, moved to the pending deletion dir", cclf.filePath)
							}
						}
					} else {
						if _, err := os.Stat(newpath); err == nil {
							return
						}
						// move the successful files to the deletion dir
						err := os.Rename(cclf.filePath, newpath)
						if err != nil {
							errCount++
							errMsg := fmt.Sprintf("File %s failed to clean up properly: %v", cclf.filePath, err)
							fmt.Println(errMsg)
							log.API.Error(errMsg)
						} else {
							fmt.Printf("File %s successfully ingested, moved to the pending deletion dir.\n", cclf.filePath)
							log.API.Infof("File %s successfully ingested, moved to the pending deletion dir", cclf.filePath)
						}
					}
				}()
			}
		}
	}
	if errCount > 0 {
		return fmt.Errorf("%d files could not be cleaned up", errCount)
	}
	return nil
}

func (m cclfFileMetadata) String() string {
	if m.name != "" {
		return m.name
	}
	return m.filePath
}
