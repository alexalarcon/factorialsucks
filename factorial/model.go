package factorial

import "net/http"

type factorialClient struct {
	http.Client
	employee_id int
	period_id   int
	calendar    []calendarDay
	shifts      []shift
	year        int
	month       int
	clock_in    string
	clock_out   string
	today_only  bool
	until_today bool
}

type period struct {
	Id          int
	Employee_id int
	Year        int
	Month       int
}
type Period struct {
	Id                                          int       `json:"id"`
	EmployeeId                                  int       `json:"employee_id"`
	Year                                        int       `json:"year"`
	Month                                       int       `json:"month"`
	StartOn                                     string    `json:"start_on"`
	EndOn                                       string    `json:"end_on"`
	State                                       string    `json:"state"`
	TimeUnitsDistibution                        []string  `json:"time_units_distibution"`
	WorkedMinutes                               int       `json:"worked_minutes"`
	TrackedMinutes                              int       `json:"tracked_minutes"`
	TrackedMinutesDistribution                  []int     `json:"tracked_minutes_distribution"`
	Distribution                                []int     `json:"distribution"`
	WorkedMinutesNotApprovedDistribution        []int     `json:"worked_minutes_not_approved_distribution"`
	BalanceMinutes                              string    `json:"balance_minutes"`
	BalanceMinutesDistribution                  []int     `json:"balance_minutes_distribution"`
	EstimatedMinutes                            int       `json:"estimated_minutes"`
	EstimatedRegularMinutes                     int       `json:"estimated_regular_minutes"`
	EstimatedRegularMinutesDistribution         []float64 `json:"estimated_regular_minutes_distribution"`
	EstimatedOvertimeMinutes                    int       `json:"estimated_overtime_minutes"`
	EstimatedMinutesUntilToday                  int       `json:"estimated_minutes_until_today"`
	EstimatedMinutesDistribution                []int     `json:"estimated_minutes_distribution"`
	EstimatedByShiftsDistribution               []bool    `json:"estimated_by_shifts_distribution"`
	EstimatedOvertimeMinutesDistribution        []float64 `json:"estimated_overtime_minutes_distribution"`
	EstimatedOvertimeRequestMinutesDistribution []float64 `json:"estimated_overtime_request_minutes_distribution"`
	WorkedHalfDays                              int       `json:"worked_half_days"`
	Permissions                                 struct {
		Read    bool `json:"read"`
		Edit    bool `json:"edit"`
		Approve bool `json:"approve"`
		Delete  bool `json:"delete"`
	} `json:"permissions"`
	Reviews []interface{} `json:"reviews"`
}
type calendarDay struct {
	Id               string
	Day              int
	DayBeforeHoliday bool
	Date             string
	Is_laborable     bool
	Is_leave         bool
	Leave_name       string
	MinutesLeft      float64
}
type newShift struct {
	ClockIn                          string      `json:"clock_in"`
	ClockOut                         string      `json:"clock_out"`
	Day                              int         `json:"day"`
	EmployeeId                       int         `json:"employee_id"`
	Workable                         bool        `json:"workable"`
	LocationType                     string      `json:"location_type"`
	TimeSettingsBreakConfigurationId interface{} `json:"time_settings_break_configuration_id"`
	Minutes                          interface{} `json:"minutes"`
	Date                             string      `json:"date"`
	Source                           string      `json:"source"`
	ReferenceDate                    string      `json:"reference_date"`
}
type shift struct {
	Id           int64  `json:"id"`
	Period_id    int64  `json:"period_id"`
	Day          int    `json:"day"`
	Clock_in     string `json:"clock_in"`
	LocationType string `json:"location_type"`
	Source       string `json:"desktop"`
	Clock_out    string `json:"clock_out"`
	Minutes      int64  `json:"minutes"`
}