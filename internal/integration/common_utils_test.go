package integration_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/limpidchart/lc-integration-testing/internal/render/github.com/limpidchart/lc-proto/render/v0"

	"github.com/google/uuid"
)

const (
	createAreaRequestJSONPath = "fixtures/create_area_request.json"
	createAreaReplyData       = "fixtures/create_area_reply_data.svg"

	createHorizontalBarRequestJSONPath = "fixtures/create_horizontal_bar_request.json"
	createHorizontalBarReplyData       = "fixtures/create_horizontal_bar_reply_data.svg"

	createLineAndVerticalBarRequestJSONPath = "fixtures/create_line_and_vertical_bar_request.json"
	createLineAndVerticalBarReplyData       = "fixtures/create_line_and_vertical_bar_reply_data.svg"

	createLineRequestJSONPath = "fixtures/create_line_request.json"
	createLineReplyData       = "fixtures/create_line_reply_data.svg"

	createScatterRequestJSONPath = "fixtures/create_scatter_request.json"
	createScatterReplyData       = "fixtures/create_scatter_reply_data.svg"

	createStackedHorizontalBarRequestJSONPath = "fixtures/create_stacked_horizontal_bar_request.json"
	createStackedHorizontalBarReplyData       = "fixtures/create_stacked_horizontal_bar_reply_data.svg"

	createStackedVerticalBarRequestJSONPath = "fixtures/create_stacked_vertical_bar_request.json"
	createStackedVerticalBarReplyData       = "fixtures/create_stacked_vertical_bar_reply_data.svg"

	createTwoLinesRequestJSONPath = "fixtures/create_two_lines_request.json"
	createTwoLinesReplyData       = "fixtures/create_two_lines_reply_data.svg"

	createTwoScattersRequestJSONPath = "fixtures/create_two_scatters_request.json"
	createTwoScattersReplyData       = "fixtures/create_two_scatters_reply_data.svg"

	createVerticalBarRequestJSONPath = "fixtures/create_vertical_bar_request.json"
	createVerticalBarReplyData       = "fixtures/create_vertical_bar_reply_data.svg"
)

// ChartReply represents a reply from create or get requests.
type ChartReply struct {
	RequestID   string     `json:"request_id"`
	ChartID     string     `json:"chart_id"`
	ChartStatus string     `json:"chart_status"`
	CreatedAt   *time.Time `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	ChartData   string     `json:"chart_data"`
}

func chartReplyFromProtobuf(reply *render.ChartReply) *ChartReply {
	createdAt := reply.CreatedAt.AsTime()
	deletedAt := reply.DeletedAt.AsTime()

	return &ChartReply{
		RequestID:   reply.RequestId,
		ChartID:     reply.ChartId,
		ChartStatus: string(reply.ChartStatus),
		CreatedAt:   &createdAt,
		DeletedAt:   &deletedAt,
		ChartData:   string(reply.ChartData),
	}
}

func checkBasicCreateChartReplyFields(rep *ChartReply, testStart time.Time) error {
	if err := checkUUID(rep.RequestID); err != nil {
		return fmt.Errorf("unable to parse request ID: %s", err)
	}

	if err := checkUUID(rep.ChartID); err != nil {
		return fmt.Errorf("unable to parse chart ID: %s", err)
	}

	if rep.CreatedAt == nil {
		return errors.New("created at value is empty")
	}

	if rep.CreatedAt.Sub(testStart).Seconds() > LCAPIHTTPTimeout {
		return fmt.Errorf("created at %s is far in the future from the test start %s", rep.CreatedAt, testStart)
	}

	if testStart.Sub(*rep.CreatedAt).Seconds() > LCAPIHTTPTimeout {
		return fmt.Errorf("created at %s is far in the past from the test start %s", rep.CreatedAt, testStart)
	}

	return nil
}

func checkUUID(u string) error {
	if _, err := uuid.Parse(u); err != nil {
		return fmt.Errorf("unable to parse %s as UUID: %w", u, err)
	}

	return nil
}

func checkCreateChartReplyCreatedAtAndDeletedAtEqual(rep *ChartReply) error {
	if !rep.CreatedAt.Equal(*rep.DeletedAt) {
		return fmt.Errorf("created at: %s and deleted at: %s are not equal", rep.CreatedAt, rep.DeletedAt)
	}

	return nil
}

func compareExpectedAndActualChartLines(expectedChartDataPath string, actualChartData []byte) error {
	expectedChartData, err := ioutil.ReadFile(expectedChartDataPath)
	if err != nil {
		return fmt.Errorf("unable to read %s: %w", expectedChartDataPath, err)
	}

	if err = checkIfElementsMatch(
		strings.Split(string(expectedChartData), "\n"),
		strings.Split(string(actualChartData), "\n"),
	); err != nil {
		return fmt.Errorf("expected and actual lines don't match: %w", err)
	}

	return nil
}

func checkIfElementsMatch(listA, listB []string) error {
	if len(listA) == 0 && len(listB) == 0 {
		return nil
	}

	var (
		extraA []string
		extraB []string
	)

	visited := make([]bool, len(listB))

	for i := range listA {
		found := false
		for j := range listB {
			if visited[j] {
				continue
			}
			if listA[i] == listB[j] {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			extraA = append(extraA, listA[i])
		}
	}

	for j := range listB {
		if !visited[j] {
			extraB = append(extraB, listB[j])
		}
	}

	if len(extraA) == 0 && len(extraB) == 0 {
		return nil
	}

	return fmt.Errorf("extra elements in the first list: %s, extra elements in the second list: %s", extraA, extraB)
}

func TestCheckUUID(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name       string
		u          string
		shouldFail bool
	}{
		{
			"good_uuid",
			uuid.NewString(),
			false,
		},
		{
			"bad_uuid",
			"u",
			true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := checkUUID(tc.u)

			if !tc.shouldFail && err != nil {
				t.Fatalf("got unexpected err: %s", err)
			}

			if tc.shouldFail && err == nil {
				t.Fatal("expected err but it's empty")
			}
		})
	}
}

func TestCheckCreateChartReplyCreatedAtAndDeletedAtEqual(t *testing.T) {
	t.Parallel()

	t1 := time.Date(2021, 8, 1, 1, 1, 1, 0, time.UTC)
	t2 := time.Date(2021, 8, 2, 2, 2, 2, 0, time.UTC)

	tt := []struct {
		name       string
		reply      *ChartReply
		shouldFail bool
	}{
		{
			"equal",
			&ChartReply{
				CreatedAt: &t1,
				DeletedAt: &t1,
			},
			false,
		},
		{
			"not_equal",
			&ChartReply{
				CreatedAt: &t1,
				DeletedAt: &t2,
			},
			true,
		},
	}

	for _, tc := range tt {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := checkCreateChartReplyCreatedAtAndDeletedAtEqual(tc.reply)

			if !tc.shouldFail && err != nil {
				t.Fatalf("got unexpected err: %s", err)
			}

			if tc.shouldFail && err == nil {
				t.Fatal("expected err but it's empty")
			}
		})
	}
}

func TestCheckIfElementsMatch(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		listA       []string
		listB       []string
		expectedErr error
	}{
		{
			"empty",
			[]string{""},
			[]string{""},
			nil,
		},
		{
			"match",
			[]string{"a", "b", "c"},
			[]string{"b", "c", "a"},
			nil,
		},
		{
			"dont_match_equal_len",
			[]string{"a", "b", "c"},
			[]string{"z", "b", "c"},
			fmt.Errorf(
				"extra elements in the first list: %v, extra elements in the second list: %v",
				[]string{"a"},
				[]string{"z"},
			),
		},
		{
			"dont_match_diff_len",
			[]string{"a", "b", "c", "d"},
			[]string{"z", "b", "c"},
			fmt.Errorf(
				"extra elements in the first list: %v, extra elements in the second list: %v",
				[]string{"a", "d"},
				[]string{"z"},
			),
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actualErr := checkIfElementsMatch(tc.listA, tc.listB)

			if tc.expectedErr != nil && (tc.expectedErr.Error() != actualErr.Error()) {
				t.Fatalf("expected %s, but got %s", tc.expectedErr.Error(), actualErr.Error())
			}
		})
	}
}
