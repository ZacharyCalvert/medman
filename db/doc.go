// Package db represents the database layer of our managed media.
// Note, this isn't some sophisticated event-driven, de-normalized dataset.  This represents
// someone's local home photos and movies.  Even if we were in the 10s of thousands,
// a modern laptop will handle that in 1-2 megs of ram and do processing/filtering in
// subsecond, given we have no real join requirements.
package db
