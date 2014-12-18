package main

const (
	// CodeCredentialsMissing means a request that was required to be authenticated had no auth data.
	CodeCredentialsMissing = "ANONE"
	// CodeCredentialsIncorrect means auth data on a request was present, but incorrect.
	CodeCredentialsIncorrect = "AFAIL"

	// CodeInvalidJobJSON means a POST body to /jobs was not parseable JSON.
	CodeInvalidJobJSON = "JPRS"
	// CodeMissingCommand means a job is missing a "cmd" element.
	CodeMissingCommand = "JCMD"
	// CodeInvalidResultSource means a job has an invalid result source.
	CodeInvalidResultSource = "JRSRC"
	// CodeInvalidResultType means a job has an invalid result type.
	CodeInvalidResultType = "JRTYPE"
	// CodeEnqueueFailure means a job could not be enqueued in the storage engine.
	CodeEnqueueFailure = "JQUEUE"
)
