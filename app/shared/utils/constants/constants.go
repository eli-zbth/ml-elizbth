package constants

const (
	EnvDevFile = ".env"
	Local      = "local"
)


const (
	HalfSpanID        = 2
	TraceparentRegexp = `^[0-9a-fA-F]{2}-[0-9a-fA-F]{32}-[0-9a-fA-F]{16}-[0-9a-fA-F]{2}`
)


const (
	CorruptedBodyRequestMsg = "corrupted body request: %s"
	InvalidBodyRequestMsg = "invalid body request: %s"
	FailedToParseBodyRequestMsg = "failed to parse body request"
)

const(
	ShortUrlLength = 7
)

const(
    MongodatabaseName = "shortUrl"
    MongocollectionName = "urls"

	MongoDbDuplicateError = "write exception: write errors: [E11000 duplicate key error collection"
	DuplicateUrlError = "Cannot create short url, it already exists"
)


const(
	UpdateShortUrlField = "key"
	UpdateLongUrlField = "url"
	UpdateUrlSuccess = "url updated successfully"
	MongoDosentExistsError = "no documents in result"
	DosentExistsError ="Can't update the URL, it dosen't exists"

)


const(
	UpdateUrlStatusSucess = "Url status updated successfully"
	UpdateStatusField = "active"
)