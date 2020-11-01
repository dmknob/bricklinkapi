package bricklinkapi

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	brickLinkAPIBaseURL  = "https://api.bricklink.com/api/store/v1"
	oauthVersion         = "1.0"
	oauthSignatureMethod = "HMAC-SHA1"
)

var (
	itemTypes = []string{"MINIFIG", "PART", "SET", "BOOK", "GEAR", "CATALOG", "INSTRUCTION", "UNSORTED_LOT", "ORIGINAL_BOX"}
)

// Bricklink is the main handler for the Bricklink API requests
type Bricklink struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	TokenSecret    string
	request        RequestHandler
}

// New returns a Bricklink handler ready to use
func New(consumerKey, consumerSecret, token, tokenSecret string) *Bricklink {
	bl := &Bricklink{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		Token:          token,
		TokenSecret:    tokenSecret,
		request: &request{
			consumerKey:    consumerKey,
			consumerSecret: consumerSecret,
			token:          token,
			tokenSecret:    tokenSecret,
		},
	}

	return bl
}

// GetItem issues a GET request to the Bricklink API and querys for the specified item.
func (bl Bricklink) GetItem(itemType, itemNumber string) (response string, err error) {
	// validate itemType
	err = validateParam(itemType, itemTypes)
	if err != nil {
		return response, err
	}

	// validate itemNumber
	if itemNumber == "" {
		return response, errors.New("itemNumber is not specified")
	}

	// build uri
	uri := "/items/" + itemType + "/" + itemNumber

	body, err := bl.request.Request("GET", uri)
	if err != nil {
		return response, err
	}

	return string(body), nil
}

// GetItemImage issues a GET request to the Bricklink API and querys for the specified item image.
func (bl Bricklink) GetItemImage(itemType, itemNumber string, colorID int) (response string, err error) {
	// validate itemType
	err = validateParam(itemType, itemTypes)
	if err != nil {
		return response, err
	}

	// validate itemNumber
	if itemNumber == "" {
		return response, errors.New("itemNumber is not specified")
	}

	// build uri
	uri := "/items/" + itemType + "/" + itemNumber + "/images/" + strconv.Itoa(colorID)

	body, err := bl.request.Request("GET", uri)
	if err != nil {
		return response, err
	}

	return string(body), nil
}

// GetItemPrice issues a GET request to the Bricklink API and querys for the price of an item.
func (bl Bricklink) GetItemPrice(itemType, itemNumber string, params map[string]string) (response string, err error) {
	// validate itemType
	err = validateParam(itemType, itemTypes)
	if err != nil {
		return response, err
	}

	// validate itemNumber
	if itemNumber == "" {
		return response, errors.New("itemNumber is not specified")
	}

	// build uri
	uri := "/items/" + itemType + "/" + itemNumber + "/price"

	// validate and build params
	if len(params) != 0 {
		var paramString string
		for k, v := range params {
			if paramString != "" {
				paramString += "&"
			}
			paramString += k + "=" + v
		}
		uri += "?" + paramString
	}

	body, err := bl.request.Request("GET", uri)
	if err != nil {
		return response, err
	}

	return string(body), nil
}

// GetColorList issues a GET request to the Bricklink API and querys for a list of all colors.
func (bl Bricklink) GetColorList() (response string, err error) {
	// build uri
	uri := "/colors"

	body, err := bl.request.Request("GET", uri)
	if err != nil {
		return response, err
	}

	return string(body), nil
}

// GetColor issues a GET request to the Bricklink API and querys for the specified color.
func (bl Bricklink) GetColor(colorID int) (response string, err error) {
	// build uri
	uri := "/colors/" + strconv.Itoa(colorID)

	body, err := bl.request.Request("GET", uri)
	if err != nil {
		return response, err
	}

	return string(body), nil
}

// GetCategoryList issues a GET request to the Bricklink API and querys for a list of all categories.
func (bl Bricklink) GetCategoryList() (response string, err error) {
	// build uri
	uri := "/categories"

	body, err := bl.request.Request("GET", uri)
	if err != nil {
		return response, err
	}

	return string(body), nil
}

// GetCategory issues a GET request to the Bricklink API and querys for a specified category.
func (bl Bricklink) GetCategory(categoryID int) (response string, err error) {
	// build uri
	uri := "/categories/" + strconv.Itoa(categoryID)

	body, err := bl.request.Request("GET", uri)
	if err != nil {
		return response, err
	}

	return string(body), nil
}

// GetInventories issues a GET request to the Bricklink API and querys for user Inventories.
func (bl Bricklink) GetInventories(categoryID int) (response string, err error) {
	// build uri
	uri := "/inventories/" + strconv.Itoa(categoryID)

	body, err := bl.request.Request("GET", uri)
	if err != nil {
		return response, err
	}

	return string(body), nil
}

// helper function to validate a param
func validateParam(param string, list []string) (err error) {
	// parameter must be set
	if param == "" {
		return errors.New("param is empty")
	}

	// param must be valid
	if !stringInSlice(param, list) {
		return fmt.Errorf("param \"%v\" is not valid", param)
	}

	return nil
}

// helper function to check if a string is in a slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if strings.ToLower(b) == strings.ToLower(a) {
			return true
		}
	}
	return false
}
