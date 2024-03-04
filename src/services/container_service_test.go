package services

/*
var testClient, _ = GetContainerConnection()

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setup() {
}

func TestContainerConnection(t *testing.T) {
	client, err := GetContainerConnection()
	assert.NotNil(t, client)
	assert.Nil(t, err)
}
func TestSaveBlob(t *testing.T) {
	location, err := SaveBlob(testClient)
	handleError(err)
	if location != "random" {
		t.Fatal("not random")
	}
}
func TestDeleteBlobByName(t *testing.T) {
	blobName, err := SaveBlob(testClient)
	handleError(err)
	blob, err := GetBlobByName(testClient, blobName)
	if err != nil || blob == nil {
		t.Fatal("Failed to get blob")
	}
	success, err := DeleteBlobByName(testClient, blobName)
	if err != nil || success != true {
		t.Fatal("Delete blob failed: ", success, err)
	}
}

func TestGetBlobByName(t *testing.T) {
	newBlobName, err := SaveBlob(testClient)
	handleError(err)
	_, err = GetBlobByName(testClient, newBlobName)
	if err != nil {
		t.Fatal("no blob")
	}
}
*/
