package connection

import (
	"errors"
	"testing"
	"time"

	"github.com/influxdata/influxdb1-client/models"
	client "github.com/influxdata/influxdb1-client/v2"
)

type TestClient struct {
	ExpectedErrorLocation string
	ExpectedError         error
	ExpectedResponse      *client.Response
}

func (tc *TestClient) Ping(timeout time.Duration) (time.Duration, string, error) {
	if tc.ExpectedErrorLocation == "Ping" {
		return 0, "", tc.ExpectedError
	}
	return 0, "", nil
}

func (tc *TestClient) Write(bp client.BatchPoints) error {
	if tc.ExpectedErrorLocation == "Ping" {
		return tc.ExpectedError
	}
	return nil
}

func (tc *TestClient) Query(q client.Query) (*client.Response, error) {
	if tc.ExpectedErrorLocation == "Query" {
		return tc.ExpectedResponse, tc.ExpectedError
	}
	return tc.ExpectedResponse, nil
}

func (tc *TestClient) QueryAsChunk(q client.Query) (*client.ChunkedResponse, error) {
	if tc.ExpectedErrorLocation == "QueryAsChunk" {
		return nil, tc.ExpectedError
	}
	return nil, tc.ExpectedError
}

func (tc *TestClient) Close() error {
	if tc.ExpectedErrorLocation == "Ping" {
		return tc.ExpectedError
	}
	return nil
}

func TestQueryDB(t *testing.T) {
	var cInflux client.Client

	cInflux = &TestClient{
		ExpectedResponse: &client.Response{},
	}

	_, err := QueryDB("SELECT * FROM test", cInflux)

	if err != nil {
		t.Fatalf("Unexpected error when querying : %s", err)
	}

	cInflux = &TestClient{
		ExpectedErrorLocation: "Query",
		ExpectedError:         errors.New("test error"),
		ExpectedResponse:      &client.Response{},
	}

	_, err = QueryDB("SELECT * WHERE test", cInflux)

	if err == nil {
		t.Fatal("It should give an error message")
	}

	cInflux = &TestClient{
		ExpectedResponse: &client.Response{
			Err: "test error",
		},
	}

	_, err = QueryDB("SELECT * WHERE test", cInflux)

	if err == nil {
		t.Fatal("It should give an error message")
	}
	cInflux = &TestClient{
		ExpectedResponse: &client.Response{
			Results: []client.Result{
				{
					Err: "",
				},
				{
					Err: "test error",
				},
				{
					Err: "",
				},
			},
		},
	}

	_, err = QueryDB("SELECT * WHERE test", cInflux)

	if err == nil {
		t.Fatal("It should give an error message")
	}
}

func TestGetIndexColumn(t *testing.T) {
	res := []client.Result{
		{
			Series: []models.Row{
				{
					Columns: []string{"time", "test"},
				},
			},
		},
	}

	testIndex := GetIndexColumn("test", res[0].Series[0])

	if testIndex != 1 {
		t.Fatalf("Expected result index is %d but the result is %d", 1, testIndex)
	}

	testIndex = GetIndexColumn("failedTest", res[0].Series[0])

	if testIndex != -1 {
		t.Fatalf("Expected result index is %d but the result is %d", -1, testIndex)
	}
}

func TestCheckIndexColumn(t *testing.T) {
	res := []client.Result{
		{
			Series: []models.Row{
				{
					Columns: []string{"time", "test", "test2"},
				},
			},
		},
	}

	err := CheckIndexColumn(res[0].Series[0], "test", "test_failed", "test2")

	if err == nil || (err != nil && err.Error() != "test_failed column not found") {
		t.Fatalf("Expected error message %s but got %s", "test_failed column not found", err)
	}

	err = CheckIndexColumn(res[0].Series[0], "test_failed")

	if err == nil || (err != nil && err.Error() != "test_failed column not found") {
		t.Fatalf("Expected error message %s but got %s", "test_failed column not found", err)
	}

	err = CheckIndexColumn(res[0].Series[0], "test", "test2")

	if err != nil {
		t.Fatal(err)
	}
}
