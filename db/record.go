package db

import "time"

/*
252F2D6C07AE83BCD39C055288044B7624416A835C7F3856042051D80F35E5B8:
  sha256: 252F2D6C07AE83BCD39C055288044B7624416A835C7F3856042051D80F35E5B8
  extensions:
    - jpg
    - JPG
  paths:
    - >-
      /Volumes/Backup_2TB-Alpha/ToProcess7/Zach/2011/Anti-TSA Rally
      6-5-2011/103D7000/DSC_0891.JPG
    - /Volumes/PhotoUsb/Processed/photos/2011/05/103D7000/DSC_0891.JPG
  earliestDate: 1304830800000
  reviewDone: true
  ignore: true
37EFA83C36884C46D5B09514B86AD6063555253060678E1773E3E13883313F1E:
  sha256: 37EFA83C36884C46D5B09514B86AD6063555253060678E1773E3E13883313F1E
  extensions:
    - jpg
    - JPG
  paths:
    - >-
      /Volumes/Backup_2TB-Alpha/ToProcess8/Old2/Pictures/2011/Wedding/Photos/Wedding/CZ-251.jpg
    - /Volumes/PhotoUsb/Processed/photos/2011/10/Wedding/CZ-251.jpg
  earliestDate: 1316754000000
  storedAt: 3/7/37EFA83C36884C46D5B09514B86AD6063555253060678E1773E3E13883313F1E.JPG
  tags:
    - 'event:Zach and Crystal Wedding'
  reviewDone: true
*/

type MediaFile struct {
	ShaKey     string   `yaml:"sha256"`
	StoredAt   string   `yaml:"storedAt"`
	Extensions []string `yaml:"extensions"`
	Earliest   int64    `yaml:"earliestDate"`
	Paths      []string `yaml:"paths"`
	Ignore     *bool    `yaml:"ignore,omitempty"`
	Reviewed   *bool    `yaml:"reviewDone,omitempty"`
	Tags       []string `yaml:"tags"`
}

func (rec MediaFile) IsReviewed() bool {
	return rec.Reviewed != nil && *rec.Reviewed
}

func (rec MediaFile) GetDate() time.Time {
	return time.Unix(rec.Earliest/1000, 0)
}

func (rec MediaFile) IsIgnoredMedia() bool {
	return rec.Reviewed != nil && *rec.Reviewed && rec.Ignore != nil && *rec.Ignore
}
