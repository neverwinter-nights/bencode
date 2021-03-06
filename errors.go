// errors.go.

//============================================================================//
//
// Copyright © 2018..2020 by McArcher.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
//============================================================================//
//
// Web Site:		'https://github.com/neverwinter-nights'.
// Author:			McArcher.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// Error Messages and cached Errors.

package bencode

import "errors"

// Error Messages and Formats.
const (
	ErrByteStringToIntError     = "Byte String to Integer Conversion Error"
	ErrDataTypeError            = "Unsupported Type"
	ErrSectionDoesNotExistError = "Section does not exist"
	ErrSelfCheckError           = "Self-Check Error"
	ErrTypeAssertionError       = "Type Assertion Error"
	ErrFileNotInitializedError  = "File is not initialized"
	//
	ErrfIntegerLengthError = "The Integer is too big: %v"
	ErrHeaderLengthError   = "The Length Header is too big: %v"
)

// Cached Errors.
var (
	ErrByteStringToInt     = errors.New(ErrByteStringToIntError)
	ErrDataType            = errors.New(ErrDataTypeError)
	ErrSectionDoesNotExist = errors.New(ErrSectionDoesNotExistError)
	ErrSelfCheck           = errors.New(ErrSelfCheckError)
	ErrTypeAssertion       = errors.New(ErrTypeAssertionError)
	ErrFileNotInitialized  = errors.New(ErrFileNotInitializedError)
)

// Formats of Error Messages.
const (
	ErrFmtSyntaxErrorAt = "Syntax Error At: '%v'."
	ErrCombinator       = ": "
)

// Combines two Errors.
func combineErrors(
	error1 error,
	error2 error,
) (result error) {
	if error1 == nil {
		if error2 == nil {
			return
		} else {
			return error2
		}
	} else {
		if error2 == nil {
			return error1
		} else {
			result = errors.New(error1.Error() + ErrCombinator + error2.Error())
			return
		}
	}
}
