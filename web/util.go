package web

import (
	"net/http"
	"strconv"
	"strings"

	pb "github.com/autograde/aguis/proto/_proto/aguis/library"
	"github.com/labstack/echo"
)

// GetEventsURL returns the event URL for a given base URL and a provider.
func GetEventsURL(baseURL, provider string) string {
	return GetProviderURL(baseURL, "hook", provider, "events")
}

// GetProviderURL returns a URL endpoint given a base URL and a provider.
func GetProviderURL(baseURL, route, provider, endpoint string) string {
	return "https://" + baseURL + "/" + route + "/" + provider + "/" + endpoint
}

// parseUint takes a string and returns the corresponding uint64. If the string
// parses to 0 or an error occurs, an error is returned.
func parseUint(s string) (uint64, error) {
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil || n == 0 {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "invalid identifier")
	}
	return n, nil
}

var errInvalidStatus = echo.NewHTTPError(http.StatusBadRequest, "invalid status query")

// parseEnrollmentStatus takes a string of comma separated status values
// and returns a slice of the corresponding status constants.
// TODO This function should be deleted and replaced with proper type safe handling without strings
// The RecordWithStatusRequest should pass []Enrollment_Status instead of State - to make it type safe.
// That is, we won't need to parse the requst.State object.
func parseEnrollmentStatus(s string) ([]pb.Enrollment_Status, error) {
	if s == "" {
		return []pb.Enrollment_Status{}, nil
	}

	ss := strings.Split(s, ",")
	if len(ss) > 4 {
		return []pb.Enrollment_Status{}, errInvalidStatus
	}
	var statuses []pb.Enrollment_Status
	for _, s := range ss {
		switch s {
		case "pending":
			statuses = append(statuses, pb.Enrollment_Pending)
		case "rejected":
			statuses = append(statuses, pb.Enrollment_Rejected)
		case "student":
			statuses = append(statuses, pb.Enrollment_Student)
		case "teacher":
			statuses = append(statuses, pb.Enrollment_Teacher)
		default:
			return []pb.Enrollment_Status{}, errInvalidStatus
		}
	}
	return statuses, nil
}
