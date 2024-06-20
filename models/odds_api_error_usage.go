/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package models

type ErrorCode struct {
	Message    string `json:"message"`
	ErrorCode  string `json:"error_code"`
	DetailsURL string `json:"details_url"`
}
