package voter_search

import (
	"sort"

	"github.com/hraywilson/Vote-Base/config"
)

// Do a binary search for the searchRec parameter
// recursively thru the voters slice data passed
// *config.VoterDataRecord
func SearchVoterRecordbyVoterId(voterId uint64, votersSlice []*config.VoterDataRecord, searchLowIdx uint32, searchHighIdx uint32) config.VoterDataRecord {

	var returnRec config.VoterDataRecord

	if searchHighIdx >= searchLowIdx {
		searchMidIdx := searchLowIdx + (searchHighIdx-searchLowIdx)/2
		// If the votersSlice at the midpoint of the poulaton then
		// copy it to a local variable and return it.
		if uint64(votersSlice[searchMidIdx].VOTER_ID) == voterId {
			returnRec = *(votersSlice)[searchMidIdx]
			return returnRec
		}
		// If the votersSlice at the midpoint of the population sorts
		// higher than the voter we are searching for, than we'll recurse
		// to the lower half of the remaining population. Else, we'll
		// recurse to the upper half of the remaining population of votersSlice.
		if uint64(votersSlice[searchMidIdx].VOTER_ID) > voterId {
			// search votersSlice population LOWER half
			return SearchVoterRecordbyVoterId(voterId, votersSlice, searchLowIdx, searchMidIdx-1)
		} else {
			// search votersSlice population UPPER half
			return SearchVoterRecordbyVoterId(voterId, votersSlice, searchMidIdx+1, searchHighIdx)
		}
	}
	returnRec.MAILING_ADDRESS_1 = "SEARCH NOT FOUND"
	return returnRec
}

func SearchVoterRecordbyName(searchRecord *config.VoterDataRecord, voterBaseData []*config.VoterDataRecord) []config.VoterDataRecord {
	voterDataSlice := []config.VoterDataRecord{}
	// Make a shallow copy of the voterBaseData slice DataRecord{}
	searchBaseData := voterBaseData
	if searchRecord.LAST_NAME != "" {
		sort.Slice(searchBaseData, func(i, j int) bool {
			return searchBaseData[i].LAST_NAME+searchBaseData[i].FIRST_NAME+searchBaseData[i].MIDDLE_NAME < searchBaseData[j].LAST_NAME+searchBaseData[j].FIRST_NAME+searchBaseData[j].MIDDLE_NAME
		})
	}
	for _, dat := range searchBaseData {
		if searchRecord.LAST_NAME != "" && searchRecord.FIRST_NAME != "" && searchRecord.MIDDLE_NAME != "" {
			if dat.LAST_NAME == searchRecord.LAST_NAME && dat.FIRST_NAME == searchRecord.FIRST_NAME && dat.MIDDLE_NAME == searchRecord.MIDDLE_NAME {
				voterDataSlice = append(voterDataSlice, *dat)
			}
		} else if searchRecord.LAST_NAME != "" && searchRecord.FIRST_NAME != "" && searchRecord.MIDDLE_NAME == "" {
			if dat.LAST_NAME == searchRecord.LAST_NAME && dat.FIRST_NAME == searchRecord.FIRST_NAME {
				voterDataSlice = append(voterDataSlice, *dat)
			}
		} else if searchRecord.LAST_NAME != "" && searchRecord.FIRST_NAME == "" && searchRecord.MIDDLE_NAME == "" {
			if dat.LAST_NAME == searchRecord.LAST_NAME {
				voterDataSlice = append(voterDataSlice, *dat)
			}
		}
	}
	return voterDataSlice
}
