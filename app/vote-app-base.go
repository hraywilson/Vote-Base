package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hraywilson/Vote-Base/config"
	"github.com/hraywilson/Vote-Base/load_data"
	"github.com/hraywilson/Vote-Base/populate_meta_data"
	"github.com/hraywilson/Vote-Base/voter_search"
)

var voterBaseData = []*config.VoterDataRecord{}

func main() {
	// Define in-memory Slice that will hold all the data for query micro services

	// define logger settings
	log.SetPrefix("Voter JSON app: ")
	log.SetFlags(2)

	// load app data

	// declare app routes
	gin.SetMode(gin.ReleaseMode)
	w3Service := gin.Default()

	w3Service.GET("/get_base_data_rowcount", getBaseDataRowcount)

	w3Service.GET("/get_voter_deets/:voterid", getVoterbyId)

	w3Service.GET("/get_voter_deets_by_name/:lastname/:firstname", getVoterbyName)

	w3Service.GET("/get_voter_deets_by_name/:lastname/:firstname/:middlename", getVoterbyName)

	w3Service.GET("/get_voter_deets_by_name/:lastname", getVoterbyName)

	w3Service.GET("/load_base_data/:filename", loadDataFiles)

	w3Service.GET("/multiply_base_data/:multiplyer", multiplyBaseData)

	w3Service.GET("/reset_base_data", resetBaseData)

	w3Service.GET("/sort_base_data", sortBaseData)

	w3Service.GET("/load_meta_data/:county", loadMetaData)

	// start api server
	w3Service.Run(config.BaseDataHost)

}

// w3Service Receivers

// Report the row count of the Base data
func getBaseDataRowcount(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{"message": "base data row count: " + fmt.Sprintf("%d", len(voterBaseData))})
	log.Println("Get Base Data rowcount:", len(voterBaseData))
}

// Get voter details by voter-id is a call to a recursive search function
func getVoterbyId(c *gin.Context) {
	inputVoterId, _ := strconv.ParseUint(c.Param("voterid"), 10, 64)
	log.Println("Query Base data by Voter Id", inputVoterId)
	c.JSON(http.StatusOK, gin.H{"voter id: " + c.Param("voterid"): voter_search.SearchVoterRecordbyVoterId(inputVoterId, voterBaseData, 0, uint32(len(voterBaseData)-1))})
}

// Get voter details by voter-id is a call to a recursive search function
func getVoterbyName(c *gin.Context) {
	searchRecord := config.VoterDataRecord{LAST_NAME: strings.ToUpper(c.Param("lastname")),
		FIRST_NAME:  strings.ToUpper(c.Param("firstname")),
		MIDDLE_NAME: strings.ToUpper(c.Param("middlename"))}
	log.Println("Query Base data by Voter Name", searchRecord.LAST_NAME, searchRecord.FIRST_NAME)
	c.JSON(http.StatusOK, gin.H{"voter name: " + searchRecord.LAST_NAME + ", " + searchRecord.FIRST_NAME: voter_search.SearchVoterRecordbyName(&searchRecord, voterBaseData)})
}

// Load comma delimited data file passed on the url
func loadDataFiles(c *gin.Context) {
	inputFile := c.Param("filename")
	c.JSON(http.StatusOK,
		gin.H{"message": "input file name: " + inputFile})

	// Load the data file and if records returned are not nil
	// we'll append that data to the Base slice
	log.Println("Loading data file:", inputFile)
	sliceOfSourceData := load_data.ReadDataFile(inputFile)
	log.Println("input file", inputFile, ": row count:", len(sliceOfSourceData))

	if sliceOfSourceData != nil {
		voterBaseData = append(voterBaseData, sliceOfSourceData...)
	}
	log.Printf("Length of Base voter Slice: %d\n", len(voterBaseData))
}

// Multiply the base data size by the multiplyer value passed
// into the function
func multiplyBaseData(c *gin.Context) {
	loopCnt, _ := strconv.Atoi(c.Param("multiplyer"))

	tempVoterSlice := populate_meta_data.Make4BigDataSet(loopCnt, voterBaseData)
	voterBaseData = append(voterBaseData, tempVoterSlice...)
}

// Reset the Base slice to zero rows
func resetBaseData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "row count reset", "base data row count": len(voterBaseData)})
	voterBaseData = nil
}

// Sort Base dataset by voter id.
func sortBaseData(c *gin.Context) {
	log.Println("Sort Base Data set")
	sort.Slice(voterBaseData, func(i, j int) bool {
		return voterBaseData[i].VOTER_ID < voterBaseData[j].VOTER_ID
	})
	c.JSON(http.StatusOK, gin.H{"message": "Sort base data by voter id",
		"data deets": fmt.Sprintf(`[low val: %d, high val: %d, len %d]`, voterBaseData[0].VOTER_ID, voterBaseData[len(voterBaseData)-1].VOTER_ID, len(voterBaseData)),
	})
}

// Load meta data
func loadMetaData(c *gin.Context) {
	inputCounty := c.Param("county")

	log.Println("load meta data table for", inputCounty)
	countyMetaDataSlice := populate_meta_data.PopulateCounty(voterBaseData, inputCounty)
	c.JSON(http.StatusOK, gin.H{
		inputCounty: fmt.Sprintf(`[low val: %d, high val: %d, len %d]`, countyMetaDataSlice[0].VOTER_ID, countyMetaDataSlice[len(countyMetaDataSlice)-1].VOTER_ID, len(countyMetaDataSlice)),
	})
}
