package load_data

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/hraywilson/Vote-Base/config"
)

// ReadDataFile function expects a file name to read and process and
// the name of a slice of struct VoterDetails passed by ref. This function
// does not does not return a value, rather it appends directly
// to the slice passed into the second parameter.
func ReadDataFile(fileName string) []*config.VoterDataRecord {
	// func ReadDataFile(fileName string, voterDataSlice []config.VoterDataRecord) {

	// verify the file existance
	fullFilePathName, rErr := os.Readlink(config.DataFilePath + fileName)
	if rErr != nil {
		log.Println(rErr)
		return []*config.VoterDataRecord{}
	}

	// if can open the file determine which file type to read
	// data files are not in a consistant format so we need
	// to read the header to determine field order to determine
	// how to parse the file for the data struct
	fh, fErr := os.Open(fullFilePathName)
	if fErr != nil {
		log.Println(fErr)
		return []*config.VoterDataRecord{}
	}
	defer fh.Close()

	// Read the entire file into memory for processing
	csvReader := csv.NewReader(fh)
	csvReader.FieldsPerRecord = -1
	fullDataSet, dErr := csvReader.ReadAll()
	if dErr != nil {
		log.Println("cvs Reader", dErr)
		return []*config.VoterDataRecord{}
	}

	// Read the header record to determine field layout
	// TODO: figure out a clean approach to loop records
	// based on the record layout
	voterSlice := []*config.VoterDataRecord{}
	// var tval uint64
	fileFmt := config.GetFileFormat(fileName[0:4])
	skipHeaderOnce := true
	if fileFmt == "fmt1" {
		// Arapahoe, Jefferson
		for _, data := range fullDataSet {
			tval, parsErr := strconv.ParseUint(data[0], 10, 64)
			// toss any records without a voter id
			if parsErr != nil {
				continue
			}
			voter := &config.VoterDataRecord{
				VOTER_ID:             tval,
				COUNTY:               data[1],
				FIRST_NAME:           data[2],
				MIDDLE_NAME:          data[3],
				LAST_NAME:            data[4],
				NAME_SUFFIX:          data[5],
				REGISTRATION_DATE:    data[6],
				EFFECTIVE_DATE:       data[7],
				LAST_UPDATED_DATE:    data[8],
				OLD_VOTER_ID:         data[9],
				PHONE_NUM:            data[10],
				HOUSE_NUM:            data[11],
				HOUSE_SUFFIX:         data[12],
				PRE_DIR:              data[13],
				STREET_NAME:          data[14],
				STREET_TYPE:          data[15],
				POST_DIR:             data[16],
				UNIT_TYPE:            data[17],
				UNIT_NUM:             data[18],
				RESIDENTIAL_ADDRESS:  data[19],
				RESIDENTIAL_CITY:     data[20],
				RESIDENTIAL_STAT:     data[21],
				RESIDENTIAL_ZIP_CODE: data[22],
				RESIDENTIAL_ZIP_PLUS: data[23],
				MAILING_ADDRESS_1:    data[24],
				MAILING_ADDRESS_2:    data[25],
				MAILING_ADDRESS_3:    data[26],
				MAILING_CITY:         data[27],
				MAILING_STATE:        data[28],
				MAILING_ZIP_CODE:     data[29],
				MAILING_ZIP_PLUS:     data[30],
				MAILING_COUNTRY:      data[31],
				VOTER_STATUS:         data[32],
				STATUS_REASON:        data[33],
				PARTY:                data[34],
				GENDER:               data[35],
				BIRTH_YEAR:           data[36],
				PRECINCT_CODE:        data[37],
				PRECINCT_NAME:        data[38],
			}

			// if fErr != nil {
			// 	log.Println(fErr)
			// 	return
			// }

			// Append voter details to the voters slice.
			if !skipHeaderOnce {
				skipHeaderOnce = true
			} else {
				voterSlice = append(voterSlice, voter)
			}
		}
	} else if fileFmt == "fmt2" {
		// Douglas
		for _, data := range fullDataSet {
			tval, parsErr := strconv.ParseUint(data[0], 10, 64)
			// toss any records without a voter id
			if parsErr != nil {
				continue
			}
			voter := &config.VoterDataRecord{
				VOTER_ID:             tval,
				COUNTY:               data[1],
				FIRST_NAME:           data[2],
				MIDDLE_NAME:          data[3],
				LAST_NAME:            data[4],
				NAME_SUFFIX:          data[5],
				REGISTRATION_DATE:    data[6],
				EFFECTIVE_DATE:       data[7],
				LAST_UPDATED_DATE:    data[8],
				OLD_VOTER_ID:         data[9],
				HOUSE_NUM:            data[10],
				HOUSE_SUFFIX:         data[11],
				PRE_DIR:              data[12],
				STREET_NAME:          data[13],
				STREET_TYPE:          data[14],
				POST_DIR:             data[15],
				UNIT_TYPE:            data[16],
				UNIT_NUM:             data[17],
				RESIDENTIAL_ADDRESS:  data[18],
				RESIDENTIAL_CITY:     data[19],
				RESIDENTIAL_STAT:     data[20],
				RESIDENTIAL_ZIP_CODE: data[21],
				RESIDENTIAL_ZIP_PLUS: data[22],
				MAILING_ADDRESS_1:    data[23],
				MAILING_ADDRESS_2:    data[24],
				MAILING_ADDRESS_3:    data[25],
				MAILING_CITY:         data[26],
				MAILING_STATE:        data[27],
				MAILING_ZIP_CODE:     data[28],
				MAILING_ZIP_PLUS:     data[29],
				MAILING_COUNTRY:      data[30],
				VOTER_STATUS:         data[31],
				STATUS_REASON:        data[32],
				PARTY:                data[33],
				GENDER:               data[34],
				BIRTH_YEAR:           data[35],
				PRECINCT_CODE:        data[36],
				PRECINCT_NAME:        data[37],
			}

			// if fErr != nil {
			// 	log.Println(fErr)
			// 	return
			// }

			// Append voter details to the voters slice.
			if !skipHeaderOnce {
				skipHeaderOnce = true
			} else {
				voterSlice = append(voterSlice, voter)
			}
		}
	} else if fileFmt == "fmt3" {
		// Denver
		for _, data := range fullDataSet {
			tval, parsErr := strconv.ParseUint(data[1], 10, 64)
			// toss any records without a voter id
			if parsErr != nil {
				continue
			}
			voter := &config.VoterDataRecord{
				VOTER_ID:    tval,
				COUNTY:      data[0],
				FIRST_NAME:  data[3],
				MIDDLE_NAME: data[4],
				LAST_NAME:   data[2],
				NAME_SUFFIX: data[5],
				// REGISTRATION_DATE:    data[6],
				// EFFECTIVE_DATE:       data[7],
				// LAST_UPDATED_DATE:    data[8],
				// OLD_VOTER_ID:         data[9],
				// HOUSE_NUM:            data[10],
				// HOUSE_SUFFIX:         data[11],
				// PRE_DIR:              data[12],
				// STREET_NAME:          data[13],
				// STREET_TYPE:          data[14],
				// POST_DIR:             data[15],
				// UNIT_TYPE:            data[16],
				// UNIT_NUM:             data[17],
				RESIDENTIAL_ADDRESS:  data[15],
				RESIDENTIAL_CITY:     data[16],
				RESIDENTIAL_STAT:     data[17],
				RESIDENTIAL_ZIP_CODE: data[18],
				// RESIDENTIAL_ZIP_PLUS: data[22],
				MAILING_ADDRESS_1: data[19],
				// MAILING_ADDRESS_2:    data[24],
				// MAILING_ADDRESS_3:    data[25],
				MAILING_CITY:     data[20],
				MAILING_STATE:    data[21],
				MAILING_ZIP_CODE: data[22],
				// MAILING_ZIP_PLUS:     data[29],
				// MAILING_COUNTRY:      data[24],
				// VOTER_STATUS:         data[31],
				// STATUS_REASON:        data[32],
				PARTY:      data[8],
				GENDER:     data[7],
				BIRTH_YEAR: data[6],
				// PRECINCT_CODE:        data[36],
				PRECINCT_NAME: data[12],
				PHONE_NUM:     data[9],
			}

			// if fErr != nil {
			// 	log.Println(fErr)
			// 	return
			// }

			// Append voter details to the voters slice.
			if !skipHeaderOnce {
				skipHeaderOnce = true
			} else {
				voterSlice = append(voterSlice, voter)
			}
		}
	} else if fileFmt == "fmt4" {
		// Boulder
		for _, data := range fullDataSet {
			tval, parsErr := strconv.ParseUint(data[1], 10, 64)
			// toss any records without a voter id
			if parsErr != nil {
				continue
			}
			voter := &config.VoterDataRecord{
				VOTER_ID:    tval,
				COUNTY:      data[0],
				FIRST_NAME:  data[2],
				MIDDLE_NAME: data[3],
				LAST_NAME:   data[4],
				NAME_SUFFIX: data[5],
				// REGISTRATION_DATE:    data[6],
				// EFFECTIVE_DATE:       data[7],
				// LAST_UPDATED_DATE:    data[8],
				// OLD_VOTER_ID:         data[9],
				HOUSE_NUM:            data[21],
				HOUSE_SUFFIX:         data[22],
				PRE_DIR:              data[22],
				STREET_NAME:          data[23],
				STREET_TYPE:          data[24],
				POST_DIR:             data[25],
				UNIT_TYPE:            data[26],
				UNIT_NUM:             data[27],
				RESIDENTIAL_ADDRESS:  data[28],
				RESIDENTIAL_CITY:     data[29],
				RESIDENTIAL_STAT:     data[30],
				RESIDENTIAL_ZIP_CODE: data[31],
				RESIDENTIAL_ZIP_PLUS: data[32],
				// MAILING_ADDRESS_1:    data[33],
				// MAILING_ADDRESS_2:    data[34],
				// MAILING_ADDRESS_3:    data[35],
				// MAILING_CITY:         data[36],
				// MAILING_STATE:        data[37],
				// MAILING_ZIP_CODE:     data[38],
				// MAILING_ZIP_PLUS:     data[39],
				// MAILING_COUNTRY:      data[40],
				// VOTER_STATUS:         data[31],
				// STATUS_REASON:        data[32],
				PARTY:      data[9],
				GENDER:     data[8],
				BIRTH_YEAR: data[7],
				// PRECINCT_CODE:        data[36],
				PRECINCT_NAME: data[19],
				PHONE_NUM:     data[9],
			}

			// if fErr != nil {
			// 	log.Println(fErr)
			// 	return
			// }

			// Append voter details to the voters slice.
			if !skipHeaderOnce {
				skipHeaderOnce = true
			} else {
				voterSlice = append(voterSlice, voter)
			}
		}
	}
	return voterSlice
}
