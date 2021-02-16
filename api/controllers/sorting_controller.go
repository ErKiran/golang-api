package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"jikkosoft/api/responses"
	"math/rand"
	"net/http"
)

type SortingRequest struct {
	Unsorted []int `json:"unsorted"`
}

type SortingResponse struct {
	Unsorted []int `json:"unsorted"`
	Sorted   []int `json:"sorted"`
}

type Result struct {
	Success bool            `json:"success"`
	Data    SortingResponse `json:"data"`
}

// SortArray godoc
// @Summary Sort the Array of numbers accordingly to per-defined criteria
// @Description This API Endpoint is used to sort array of numbers
// @Tags sort
// @Accept  json
// @Produce  json
// @Success 200 {object} Result
// @Param ticket body SortingRequest true "request"
// @Router /sort [post]
func (server *Server) SortArray(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	var request SortingRequest
	err = json.Unmarshal(body, &request)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if request.Unsorted == nil {
		responses.ERROR(w, http.StatusNotAcceptable, errors.New("unsorted array not found"))
		return
	}

	uniqueNumbers, repeatedNumbers := arraySeperator(request.Unsorted)

	sortedArray := quickSort(uniqueNumbers)
	sortedArray = append(sortedArray, repeatedNumbers...)
	responses.JSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data": SortingResponse{
			Unsorted: request.Unsorted,
			Sorted:   sortedArray,
		}})
}

func arraySeperator(intSlice []int) ([]int, []int) {
	keys := make(map[int]bool)
	unique := []int{}
	repeated := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			unique = append(unique, entry)
		} else {
			repeated = append(repeated, entry)
		}
	}
	return unique, repeated
}

func quickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = quickSort(lowPart)
	highPart = quickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}
