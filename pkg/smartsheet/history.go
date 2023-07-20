/*
 * Copyright 2020 wfleming@grumpysysadm.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package smartsheet

import (
	"fmt"
	"time"
)

// Define a struct to represent the response body from the /sheets/{sheetID} endpoint
type HistoryResponse struct {
	PageNumber int           `json:"pageNumber"`
	PageSize   int           `json:"pageSize"`
	TotalPages int           `json:"totalPages"`
	TotalCount int           `json:"totalCount"`
	Data       []HistoryData `json:"data"`
}

type HistoryData struct {
	ColumnID     int        `json:"columnId"`
	Value        string     `json:"value"`
	DisplayValue string     `json:"displayValue"`
	Formula      string     `json:"formula"`
	ModifiedBy   time.Time  `json:"modifiedBy"`
	ModifiedAt   CellAuthor `json:"modifiedAt"`
}

type CellAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Return ResultObject object
func (c Client) GetHistory(sheetId int64, rowId int64, columnId int64) (*HistoryResponse, error) {
	var historyResponse HistoryResponse
	resp, err := c.get(fmt.Sprintf("%s/sheets/%d/rows/%d/columns/%d/history", apiEndpoint, sheetId, rowId, columnId))
	if err != nil {
		return nil, err
	}
	if dErr := c.decodeJSON(resp, &historyResponse); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	return &historyResponse, nil
}
