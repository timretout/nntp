// Copyright (C) 2015 Tim Retout <tim@retout.co.uk>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package nntp

// NNTP status codes as defined by RFC 3977.
const (
	StatusHelpText       = 100
	StatusCapabilityList = 101
	StatusServerDate     = 111

	StatusPostingAllowed       = 200
	StatusPostingProhibited    = 201
	StatusConnectionClosing    = 205
	StatusGroupSelected        = 211
	StatusInformationFollows   = 215
	StatusArticleFollows       = 220
	StatusArticleHeadersFollow = 221
	StatusArticleBodyFollows   = 222
	StatusArticleSelected      = 223
	StatusOverviewFollows      = 224
	StatusHeadersFollow        = 225
	StatusNewArticlesFollow    = 230
	StatusNewNewsgroupsFollow  = 231
	StatusArticleTransferredOK = 235
	StatusArticleReceivedOK    = 240

	StatusTransferArticle = 335
	StatusPostArticle     = 340

	StatusServiceNotAvailable   = 400
	StatusWrongServerMode       = 401
	StatusInternalFault         = 403
	StatusNoSuchGroup           = 411
	StatusNoGroupSelected       = 412
	StatusCurrentArticleInvalid = 420
	StatusNoNextArticle         = 421
	StatusNoPreviousArticle     = 422
	StatusNoSuchArticleNumber   = 423
	StatusNoSuchMessageID       = 430
	StatusArticleNotWanted      = 435
	StatusTransferFailed        = 436
	StatusTransferRejected      = 437
	StatusPostingNotPermitted   = 440
	StatusPostingFailed         = 441
	StatusNotAuthenticated      = 480
	StatusNotPrivate            = 483

	StatusUnknownCommand      = 500
	StatusSyntaxError         = 501
	StatusNotPermitted        = 502
	StatusFeatureNotSupported = 503
	StatusEncodingError       = 504
)
