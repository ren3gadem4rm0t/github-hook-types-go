// Package github provides strongly-typed definitions for GitHub webhook
// events and payloads. This library aims to provide complete coverage
// of all webhook event types to assist developers in building applications
// that respond to GitHub webhook events.
//
// Each webhook event type is represented by a corresponding Go struct
// with appropriate JSON tags to support direct unmarshaling of webhook
// payloads received from GitHub.
package github
