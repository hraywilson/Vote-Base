package populate_meta_data

// type sliceHeader struct {
// 	Length        int
// 	Capacity      int
// 	ZerothElement *byte
// }
import (
	"sort"
	"strings"

	"github.com/hraywilson/Vote-Base/app/config"
)

func PopulateCounty(voterSlice []*config.VoterDataRecord, countyName string) []*config.VoterDataRecord {
	//
	countyToMatch := strings.ToUpper(countyName)
	sliceOfCountyPtr := []*config.VoterDataRecord{}

	for _, item := range voterSlice {
		if strings.ToUpper(item.COUNTY) == countyToMatch {
			sliceOfCountyPtr = append(sliceOfCountyPtr, item)
		}
	}

	return sliceOfCountyPtr
}

// Make4BigDataSet takes a multiplier value for increaseBy param that
// will be used to increase the total size of the base data slice.
// This function is used for load testing, but is also an example
// of a deep copy of a slice of pointers to struct data structure.
func Make4BigDataSet(increaseBy int, sourceVoterSlice []*config.VoterDataRecord) []*config.VoterDataRecord {

	destVoterSlice := []*config.VoterDataRecord{}

	for ; increaseBy >= 1; increaseBy-- {
		// deep copy before appending to destVoterSlice
		for _, voter := range sourceVoterSlice {
			voterRec := *voter
			destVoterSlice = append(destVoterSlice, &voterRec)
		}
		// Sort the destVoterSlice before bumping the voter id values
		sort.Slice(destVoterSlice, func(i, j int) bool {
			return destVoterSlice[i].VOTER_ID < destVoterSlice[j].VOTER_ID
		})
		// bump voter id by one so each voter remains "unique"
		// for idx, _ := range destVoterSlice {
		for idx := 0; idx < len(destVoterSlice); idx++ {
			if idx > 0 && destVoterSlice[idx].VOTER_ID == destVoterSlice[idx-1].VOTER_ID {
				destVoterSlice[idx].VOTER_ID++
			}
		}
	}
	return destVoterSlice
}
