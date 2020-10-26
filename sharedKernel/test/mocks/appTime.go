package mocks

const (
	FakeTodayEpoch = 1577880000 // 2020-01-01 T 12:00:00 GMT PM
	FakeFutureEpoch = 1583841600 // 2020-03-10 T 12:00:00 GMT PN
	FakePastEpoch = 1573560000 // 2019-11-01 T 12:00:00 GMT PN
	FakeDateTime = "2020-01-01 T 12:00:00"
)

type AppTime struct {
	GetEpochFromDateFunc func(date string) (int64, error)
	AddToTimestampFunc func(years int, months int, days int) int64
	GetEpochFromDateAndTimeFunc func(datetime string) (int64, error)
}

func (d AppTime) NowDateTime() string {
	return FakeDateTime
}

func (d AppTime) GetEpochFromDateAndTime(datetime string) (int64, error) {
	return d.GetEpochFromDateAndTimeFunc(datetime)
}

func (d AppTime) NowTimestamp() int64 {
	return FakeTodayEpoch
}

func (d AppTime) AddToTimestamp(years int, months int, days int) int64 {
	return d.AddToTimestampFunc(years, months, days)
}

func (d AppTime) GetEpochFromDate(date string) (int64, error) {
	return d.GetEpochFromDateFunc(date)
}

