package config

const BaseDataHost string = "horus.byrds:8080"
const DataFilePath string = "../vote_data/"

// GetFileFormat expects the first 4 charaacters of a
// filename to determine a canonicle file format of the
// voter data. It will return string of a known format
// for data file parsing
func GetFileFormat(f string) string {

	for _, val := range [...]string{"arap", "jeff"} {
		if val == f {
			return "fmt1"
		}
	}
	for _, val := range [...]string{"doug"} {
		if val == f {
			return "fmt2"
		}
	}
	for _, val := range [...]string{"denv"} {
		if val == f {
			return "fmt3"
		}
	}
	for _, val := range [...]string{"boul"} {
		if val == f {
			return "fmt4"
		}
	}
	return "no match"
}

type VoterDataRecord struct {
	VOTER_ID             uint64 `json:"voterid"`
	COUNTY               string `json:"county"`
	FIRST_NAME           string `json:"firstname"`
	MIDDLE_NAME          string `json:"middlename"`
	LAST_NAME            string `json:"lastname"`
	NAME_SUFFIX          string `json:"namesuffix"`
	REGISTRATION_DATE    string `json:"registrationdate"`
	EFFECTIVE_DATE       string `json:"effectivedate"`
	LAST_UPDATED_DATE    string `json:"lastupdateddate"`
	OLD_VOTER_ID         string `json:"oldvoterid"`
	PHONE_NUM            string `json:"phonenum"`
	HOUSE_NUM            string `json:"houseum"`
	HOUSE_SUFFIX         string `json:"housesuffix"`
	PRE_DIR              string `json:"predir"`
	STREET_NAME          string `json:"streetname"`
	STREET_TYPE          string `json:"streettype"`
	POST_DIR             string `json:"postdir"`
	UNIT_TYPE            string `json:"unittype"`
	UNIT_NUM             string `json:"unitnum"`
	RESIDENTIAL_ADDRESS  string `json:"residentialaddress"`
	RESIDENTIAL_CITY     string `json:"residentialcity"`
	RESIDENTIAL_STAT     string `json:"residentialstate"`
	RESIDENTIAL_ZIP_CODE string `json:"residentialzipcode"`
	RESIDENTIAL_ZIP_PLUS string `json:"redidentialzipplus"`
	MAILING_ADDRESS_1    string `json:"mailingaddress1"`
	MAILING_ADDRESS_2    string `json:"mailingaddress2"`
	MAILING_ADDRESS_3    string `json:"mailingaddress3"`
	MAILING_CITY         string `json:"mailingcity"`
	MAILING_STATE        string `json:"mailingstate"`
	MAILING_ZIP_CODE     string `json:"mailingzipcode"`
	MAILING_ZIP_PLUS     string `json:"mailingzipplus"`
	MAILING_COUNTRY      string `json:"mailingcountry"`
	VOTER_STATUS         string `json:"voterstatus"`
	STATUS_REASON        string `json:"statusreason"`
	PARTY                string `json:"party"`
	GENDER               string `json:"gender"`
	BIRTH_YEAR           string `json:"birthyear"`
	PRECINCT_CODE        string `json:"precinctcode"`
	PRECINCT_NAME        string `json:"precinctname"`
}

type VoterDataCounties struct {
	CountyPointer *VoterDataRecord
}
